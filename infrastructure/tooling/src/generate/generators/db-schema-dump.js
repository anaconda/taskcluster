const tcpg = require('taskcluster-lib-postgres');
const testing = require('taskcluster-lib-testing');
const parseDbURL = require('pg-connection-string').parse;
const { REPO_ROOT, writeRepoFile, execCommand } = require('../../utils');

// Generate a readable JSON version of the schema.
exports.tasks = [{
  title: 'DB Schema Dump',
  requires: ['db-schema-serializable'],
  provides: ['db-schema-dump'],
  run: async (requirements, utils) => {
    if (!process.env.TEST_DB_URL) {
      throw new Error("'yarn generate' requires $TEST_DB_URL to be set");
    }

    // reset the DB back to its empty state..
    utils.step({title: 'Reset Test Database'});
    await testing.resetDb();

    // .. and then upgrade to the latest version
    utils.step({title: 'Upgrade Test Database'});
    const schema = tcpg.Schema.fromSerializable(requirements['db-schema-serializable']);
    await tcpg.Database.upgrade({
      schema,
      showProgress: message => utils.status({message}),
      usernamePrefix: 'test',
      adminDbUrl: process.env.TEST_DB_URL,
    });

    // now dump the schema of the DB
    utils.step({title: 'Dump Test Database'});
    const { host, port, user, password, database } = parseDbURL(process.env.TEST_DB_URL);
    const pgdump = await execCommand({
      dir: REPO_ROOT,
      command: [
        'pg_dump',
        '--schema-only',
        '-h', host,
        '-p', (port || 5432).toString(),
        '-U', user,
        '-d', database,
      ],
      keepAllOutput: true,
      env: {
        ...process.env,
        PGPASSWORD: password,
      },
      utils,
    });

    /* Parse the output as separated by comments of the form:
     *
     * --
     * -- Name: access_token_table_entities; Type: TABLE; Schema: public; Owner: dustindev
     * --
     *
     * for primary keys:
     *
     * -- Name: access_token_table_entities access_token_table_entities_pkey; Type: CONSTRAINT;
     *    Schema: public; Owner: dustindev
     *   [^newline added]
     *
     * and for indexes:
     *
     * -- Name: azure_queue_messages_inserted; Type: INDEX; Schema: public; Owner: dustindev
     */

    const re = /--\n-- Name: ([^;]+); Type: ([^;]+); Schema: [^;]+; Owner: [^;]+\n--\n/;
    const statements = [];
    const parts = pgdump.split(re);
    parts.shift(); // throw out the header
    for (let i = 0; i < parts.length; i += 3) {
      const [name, type, body] = parts.slice(i, i + 3);
      statements.push({name, type, body});
    }

    const byTable = new Map();
    for (let {name, type, body} of statements) {
      switch (type) {
        case 'TABLE': {
          byTable.set(name, {...(byTable.get(name) || {}), create: body});
          break;
        }

        case 'CONSTRAINT': {
          const re = /ALTER TABLE ONLY public\.([0-9a-zA-Z_-]+)/;
          const [_, tableName] = re.exec(body);
          byTable.set(tableName, {...(byTable.get(tableName) || {}), constraint: body});
          break;
        }

        case 'INDEX': {
          const re = /CREATE INDEX [0-9a-zA-Z_-]+ ON public\.([0-9a-zA-Z_-]+)/;
          const [_, tableName] = re.exec(body);
          const table = byTable.get(tableName) || {};
          table.indexes = (table.indexes || []).concat([body]);
          byTable.set(tableName, table);
          break;
        }
      }
    }

    // reformat that schema so that it contains only what we want, and in
    // a format that doesn't change from run to run

    const output = [];
    output.push('# Database Schema\n');

    for (let tableName of [...byTable.keys()].sort()) {
      output.push(` * [\`${tableName}\`](#${tableName})`);
    }

    output.push('\n');

    for (let tableName of [...byTable.keys()].sort()) {
      const {create, constraint, indexes} = byTable.get(tableName);

      output.push(`## ${tableName}\n`);

      output.push("```sql");
      if (create) {
        const text = create
          // drop the unnecessary ownership of this table
          .replace(/ALTER TABLE [0-9a-zA-Z._-]+ OWNER TO [0-9a-zA-Z_-]+;/g, '')
          // drop the unnecesary 'public.' in the table name
          .replace(/CREATE TABLE public\./, 'CREATE TABLE ');
        output.push(text.trim());
      }

      if (constraint) {
        const text = constraint
          // drop the unnecesary 'public.' in the table name
          .replace(/ALTER TABLE ONLY public\./, 'ALTER TABLE ');
        output.push(text.trim());
      }

      if (indexes) {
        for (let statement of indexes) {
          const text = statement
            // drop the unnecesary 'public.' in the table name
            .replace(/ON public\./, 'ON ');
          output.push(text.trim());
        }
      }

      output.push("```\n");
    }

    await writeRepoFile('db/schema.md', output.join('\n'));
  },
}];
