const helper = require('./helper');
const { Schema } = require('taskcluster-lib-postgres');
const { Entity } = require('taskcluster-lib-entities');
const path = require('path');
const assert = require('assert').strict;

helper.dbSuite(path.basename(__filename), function() {
  let db;

  teardown(async function() {
    if (db) {
      try {
        await db.close();
      } finally {
        db = null;
      }
    }
  });

  const schema = Schema.fromDbDirectory(path.join(__dirname, 'db'));
  const properties = {
    taskId: 'string',
    provisionerId: 'string',
    workerType: 'string',
  };
  const entity = Entity.configure({
    partitionKey: 'taskId',
    rowKey: 'task',
    properties,
  });
  const serviceName = 'test-entities';

  suite('ensure table', function() {
    test('ensure table', async function() {
      db = await helper.withDb({ schema, serviceName });
      entity.setup({ tableName: 'test_entities', db, serviceName });

      await entity.ensureTable();
    });
    test('ensure table (again)', async function() {
      db = await helper.withDb({ schema, serviceName });
      entity.setup({ tableName: 'test_entities', db, serviceName });

      await entity.ensureTable();
      await entity.ensureTable();
    });
    test('ensure table creates table when it doesn\'t exist', async function() {
      db = await helper.withDb({ schema, serviceName });
      entity.setup({ tableName: 'test_entities', db, serviceName });

      await entity.removeTable();
      await entity.ensureTable();
    });
  });
});
