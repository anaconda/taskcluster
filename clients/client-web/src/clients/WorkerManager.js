// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.

import Client from '../Client';

export default class WorkerManager extends Client {
  constructor(options = {}) {
    super({
      serviceName: 'worker-manager',
      serviceVersion: 'v1',
      exchangePrefix: '',
      ...options,
    });
    this.ping.entry = {"args":[],"category":"Ping Server","method":"get","name":"ping","query":[],"route":"/ping","stability":"stable","type":"function"}; // eslint-disable-line
    this.lbheartbeat.entry = {"args":[],"category":"Load Balancer Heartbeat","method":"get","name":"lbheartbeat","query":[],"route":"/__lbheartbeat__","stability":"stable","type":"function"}; // eslint-disable-line
    this.version.entry = {"args":[],"category":"Taskcluster Version","method":"get","name":"version","query":[],"route":"/__version__","stability":"stable","type":"function"}; // eslint-disable-line
    this.listProviders.entry = {"args":[],"category":"Providers","method":"get","name":"listProviders","output":true,"query":["continuationToken","limit"],"route":"/providers","scopes":"worker-manager:list-providers","stability":"stable","type":"function"}; // eslint-disable-line
    this.createWorkerPool.entry = {"args":["workerPoolId"],"category":"Worker Pools","input":true,"method":"put","name":"createWorkerPool","output":true,"query":[],"route":"/worker-pool/<workerPoolId>","scopes":{"AllOf":["worker-manager:manage-worker-pool:<workerPoolId>","worker-manager:provider:<providerId>"]},"stability":"stable","type":"function"}; // eslint-disable-line
    this.updateWorkerPool.entry = {"args":["workerPoolId"],"category":"Worker Pools","input":true,"method":"post","name":"updateWorkerPool","output":true,"query":[],"route":"/worker-pool/<workerPoolId>","scopes":{"AllOf":["worker-manager:manage-worker-pool:<workerPoolId>","worker-manager:provider:<providerId>"]},"stability":"experimental","type":"function"}; // eslint-disable-line
    this.deleteWorkerPool.entry = {"args":["workerPoolId"],"category":"Worker Pools","method":"delete","name":"deleteWorkerPool","output":true,"query":[],"route":"/worker-pool/<workerPoolId>","scopes":"worker-manager:manage-worker-pool:<workerPoolId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.workerPool.entry = {"args":["workerPoolId"],"category":"Worker Pools","method":"get","name":"workerPool","output":true,"query":[],"route":"/worker-pool/<workerPoolId>","scopes":"worker-manager:get-worker-pool:<workerPoolId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.listWorkerPools.entry = {"args":[],"category":"Worker Pools","method":"get","name":"listWorkerPools","output":true,"query":["continuationToken","limit"],"route":"/worker-pools","scopes":"worker-manager:list-worker-pools","stability":"stable","type":"function"}; // eslint-disable-line
    this.reportWorkerError.entry = {"args":["workerPoolId"],"category":"Worker Interface","input":true,"method":"post","name":"reportWorkerError","output":true,"query":[],"route":"/worker-pool-errors/<workerPoolId>","scopes":{"AllOf":["assume:worker-pool:<workerPoolId>","assume:worker-id:<workerGroup>/<workerId>"]},"stability":"stable","type":"function"}; // eslint-disable-line
    this.listWorkerPoolErrors.entry = {"args":["workerPoolId"],"category":"Worker Pools","method":"get","name":"listWorkerPoolErrors","output":true,"query":["continuationToken","limit"],"route":"/worker-pool-errors/<workerPoolId>","scopes":"worker-manager:list-worker-pool-errors:<workerPoolId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.listWorkersForWorkerGroup.entry = {"args":["workerPoolId","workerGroup"],"category":"Workers","method":"get","name":"listWorkersForWorkerGroup","output":true,"query":["continuationToken","limit"],"route":"/workers/<workerPoolId>:/<workerGroup>","scopes":"worker-manager:list-workers:<workerPoolId>/<workerGroup>","stability":"stable","type":"function"}; // eslint-disable-line
    this.worker.entry = {"args":["workerPoolId","workerGroup","workerId"],"category":"Workers","method":"get","name":"worker","output":true,"query":[],"route":"/workers/<workerPoolId>:/<workerGroup>/<workerId>","scopes":"worker-manager:get-worker:<workerPoolId>/<workerGroup>/<workerId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.createWorker.entry = {"args":["workerPoolId","workerGroup","workerId"],"category":"Workers","input":true,"method":"put","name":"createWorker","output":true,"query":[],"route":"/workers/<workerPoolId>:/<workerGroup>/<workerId>","scopes":"worker-manager:create-worker:<workerPoolId>/<workerGroup>/<workerId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.updateWorker.entry = {"args":["workerPoolId","workerGroup","workerId"],"category":"Workers","input":true,"method":"post","name":"updateWorker","output":true,"query":[],"route":"/workers/<workerPoolId>:/<workerGroup>/<workerId>","scopes":"worker-manager:update-worker:<workerPoolId>/<workerGroup>/<workerId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.removeWorker.entry = {"args":["workerPoolId","workerGroup","workerId"],"category":"Workers","method":"delete","name":"removeWorker","query":[],"route":"/workers/<workerPoolId>/<workerGroup>/<workerId>","scopes":"worker-manager:remove-worker:<workerPoolId>/<workerGroup>/<workerId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.listWorkersForWorkerPool.entry = {"args":["workerPoolId"],"category":"Workers","method":"get","name":"listWorkersForWorkerPool","output":true,"query":["continuationToken","limit"],"route":"/workers/<workerPoolId>","scopes":"worker-manager:list-workers:<workerPoolId>","stability":"stable","type":"function"}; // eslint-disable-line
    this.registerWorker.entry = {"args":[],"category":"Worker Interface","input":true,"method":"post","name":"registerWorker","output":true,"query":[],"route":"/worker/register","stability":"stable","type":"function"}; // eslint-disable-line
    this.reregisterWorker.entry = {"args":[],"category":"Workers","input":true,"method":"post","name":"reregisterWorker","output":true,"query":[],"route":"/worker/reregister","scopes":"worker-manager:reregister-worker:<workerPoolId>/<workerGroup>/<workerId>","stability":"experimental","type":"function"}; // eslint-disable-line
  }
  /* eslint-disable max-len */
  // Respond without doing anything.
  // This endpoint is used to check that the service is up.
  /* eslint-enable max-len */
  ping(...args) {
    this.validate(this.ping.entry, args);

    return this.request(this.ping.entry, args);
  }
  /* eslint-disable max-len */
  // Respond without doing anything.
  // This endpoint is used to check that the service is up.
  /* eslint-enable max-len */
  lbheartbeat(...args) {
    this.validate(this.lbheartbeat.entry, args);

    return this.request(this.lbheartbeat.entry, args);
  }
  /* eslint-disable max-len */
  // Respond with the JSON version object.
  // https://github.com/mozilla-services/Dockerflow/blob/main/docs/version_object.md
  /* eslint-enable max-len */
  version(...args) {
    this.validate(this.version.entry, args);

    return this.request(this.version.entry, args);
  }
  /* eslint-disable max-len */
  // Retrieve a list of providers that are available for worker pools.
  /* eslint-enable max-len */
  listProviders(...args) {
    this.validate(this.listProviders.entry, args);

    return this.request(this.listProviders.entry, args);
  }
  /* eslint-disable max-len */
  // Create a new worker pool. If the worker pool already exists, this will throw an error.
  /* eslint-enable max-len */
  createWorkerPool(...args) {
    this.validate(this.createWorkerPool.entry, args);

    return this.request(this.createWorkerPool.entry, args);
  }
  /* eslint-disable max-len */
  // Given an existing worker pool definition, this will modify it and return
  // the new definition.
  // To delete a worker pool, set its `providerId` to `"null-provider"`.
  // After any existing workers have exited, a cleanup job will remove the
  // worker pool.  During that time, the worker pool can be updated again, such
  // as to set its `providerId` to a real provider.
  /* eslint-enable max-len */
  updateWorkerPool(...args) {
    this.validate(this.updateWorkerPool.entry, args);

    return this.request(this.updateWorkerPool.entry, args);
  }
  /* eslint-disable max-len */
  // Mark a worker pool for deletion.  This is the same as updating the pool to
  // set its providerId to `"null-provider"`, but does not require scope
  // `worker-manager:provider:null-provider`.
  /* eslint-enable max-len */
  deleteWorkerPool(...args) {
    this.validate(this.deleteWorkerPool.entry, args);

    return this.request(this.deleteWorkerPool.entry, args);
  }
  /* eslint-disable max-len */
  // Fetch an existing worker pool defition.
  /* eslint-enable max-len */
  workerPool(...args) {
    this.validate(this.workerPool.entry, args);

    return this.request(this.workerPool.entry, args);
  }
  /* eslint-disable max-len */
  // Get the list of all the existing worker pools.
  /* eslint-enable max-len */
  listWorkerPools(...args) {
    this.validate(this.listWorkerPools.entry, args);

    return this.request(this.listWorkerPools.entry, args);
  }
  /* eslint-disable max-len */
  // Report an error that occurred on a worker.  This error will be included
  // with the other errors in `listWorkerPoolErrors(workerPoolId)`.
  // Workers can use this endpoint to report startup or configuration errors
  // that might be associated with the worker pool configuration and thus of
  // interest to a worker-pool administrator.
  // NOTE: errors are publicly visible.  Ensure that none of the content
  // contains secrets or other sensitive information.
  /* eslint-enable max-len */
  reportWorkerError(...args) {
    this.validate(this.reportWorkerError.entry, args);

    return this.request(this.reportWorkerError.entry, args);
  }
  /* eslint-disable max-len */
  // Get the list of worker pool errors.
  /* eslint-enable max-len */
  listWorkerPoolErrors(...args) {
    this.validate(this.listWorkerPoolErrors.entry, args);

    return this.request(this.listWorkerPoolErrors.entry, args);
  }
  /* eslint-disable max-len */
  // Get the list of all the existing workers in a given group in a given worker pool.
  /* eslint-enable max-len */
  listWorkersForWorkerGroup(...args) {
    this.validate(this.listWorkersForWorkerGroup.entry, args);

    return this.request(this.listWorkersForWorkerGroup.entry, args);
  }
  /* eslint-disable max-len */
  // Get a single worker.
  /* eslint-enable max-len */
  worker(...args) {
    this.validate(this.worker.entry, args);

    return this.request(this.worker.entry, args);
  }
  /* eslint-disable max-len */
  // Create a new worker.  This is only useful for worker pools where the provider
  // does not create workers automatically, such as those with a `static` provider
  // type.  Providers that do not support creating workers will return a 400 error.
  // See the documentation for the individual providers, and in particular the
  // [static provider](https://docs.taskcluster.net/docs/reference/core/worker-manager/)
  // for more information.
  /* eslint-enable max-len */
  createWorker(...args) {
    this.validate(this.createWorker.entry, args);

    return this.request(this.createWorker.entry, args);
  }
  /* eslint-disable max-len */
  // Update an existing worker in-place.  Like `createWorker`, this is only useful for
  // worker pools where the provider does not create workers automatically.
  // This method allows updating all fields in the schema unless otherwise indicated
  // in the provider documentation.
  // See the documentation for the individual providers, and in particular the
  // [static provider](https://docs.taskcluster.net/docs/reference/core/worker-manager/)
  // for more information.
  /* eslint-enable max-len */
  updateWorker(...args) {
    this.validate(this.updateWorker.entry, args);

    return this.request(this.updateWorker.entry, args);
  }
  /* eslint-disable max-len */
  // Remove an existing worker.  The precise behavior of this method depends
  // on the provider implementing the given worker.  Some providers
  // do not support removing workers at all, and will return a 400 error.
  // Others may begin removing the worker, but it may remain available via
  // the API (perhaps even in state RUNNING) afterward.
  /* eslint-enable max-len */
  removeWorker(...args) {
    this.validate(this.removeWorker.entry, args);

    return this.request(this.removeWorker.entry, args);
  }
  /* eslint-disable max-len */
  // Get the list of all the existing workers in a given worker pool.
  /* eslint-enable max-len */
  listWorkersForWorkerPool(...args) {
    this.validate(this.listWorkersForWorkerPool.entry, args);

    return this.request(this.listWorkersForWorkerPool.entry, args);
  }
  /* eslint-disable max-len */
  // Register a running worker.  Workers call this method on worker start-up.
  // This call both marks the worker as running and returns the credentials
  // the worker will require to perform its work.  The worker must provide
  // some proof of its identity, and that proof varies by provider type.
  /* eslint-enable max-len */
  registerWorker(...args) {
    this.validate(this.registerWorker.entry, args);

    return this.request(this.registerWorker.entry, args);
  }
  /* eslint-disable max-len */
  // Reregister a running worker.
  // This will generate and return new Taskcluster credentials for the worker
  // on that instance to use. The credentials will not live longer the
  // `registrationTimeout` for that worker. The endpoint will update `terminateAfter`
  // for the worker so that worker-manager does not terminate the instance.
  /* eslint-enable max-len */
  reregisterWorker(...args) {
    this.validate(this.reregisterWorker.entry, args);

    return this.request(this.reregisterWorker.entry, args);
  }
}
