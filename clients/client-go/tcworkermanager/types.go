// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcworkermanager

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster/v50/clients/client-go"
)

type (
	// Proof that this call is coming from the worker identified by the other fields.
	// The form of this proof varies depending on the provider type.
	AwsProviderType struct {

		// Instance identity document that is obtained by
		// curl http://169.254.169.254/latest/dynamic/instance-identity/document on the instance
		Document string `json:"document"`

		// The signature for instance identity document. Can be obtained by
		// curl http://169.254.169.254/latest/dynamic/instance-identity/signature on the instance
		Signature string `json:"signature"`
	}

	// Proof that this call is coming from the worker identified by the other fields.
	// The form of this proof varies depending on the provider type.
	AzureProviderType struct {

		// Attested data document that is obtained by
		// curl http://169.254.169.254/metadata/attested/document on the instance
		Document string `json:"document"`
	}

	// The credentials the worker
	// will need to perform its work.  Specifically, credentials with scopes
	// * `assume:worker-pool:<workerPoolId>`
	// * `assume:worker-id:<workerGroup>/<workerId>`
	// * `queue:worker-id:<workerGroup>/<workerId>`
	// * `secrets:get:worker-pool:<workerPoolId>`
	// * `queue:claim-work:<workerPoolId>`
	// * `worker-manager:remove-worker:<workerPoolId>/<workerGroup>/<workerId>`
	Credentials struct {
		AccessToken string `json:"accessToken"`

		// Note that a certificate may not be provided, if the credentials are not temporary.
		Certificate string `json:"certificate,omitempty"`

		ClientID string `json:"clientId"`
	}

	// The credentials the worker
	// will need to perform its work. Specifically, credentials with scopes
	// * `assume:worker-pool:<workerPoolId>`
	// * `assume:worker-id:<workerGroup>/<workerId>`
	// * `queue:worker-id:<workerGroup>/<workerId>`
	// * `secrets:get:worker-pool:<workerPoolId>`
	// * `queue:claim-work:<workerPoolId>`
	// * `worker-manager:remove-worker:<workerPoolId>/<workerGroup>/<workerId>`
	// * `worker-manager:reregister-worker:<workerPoolId>/<workerGroup>/<workerId>`
	Credentials1 struct {
		AccessToken string `json:"accessToken"`

		// Note that a certificate may not be provided, if the credentials are not temporary.
		Certificate string `json:"certificate,omitempty"`

		ClientID string `json:"clientId"`
	}

	// Proof that this call is coming from the worker identified by the other fields.
	// The form of this proof varies depending on the provider type.
	GoogleProviderType struct {

		// A JWT token as defined in [this google documentation](https://cloud.google.com/compute/docs/instances/verifying-instance-identity)
		Token string `json:"token"`
	}

	// Response from a `listWorkers` request.
	ListWorkersResponse struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of workers in the worker-type.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerTypes` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of workers in this worker-type.
		Workers []Worker `json:"workers"`
	}

	// A list of providers
	ProviderList struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of workers in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of all providers
		Providers []Var `json:"providers"`
	}

	// Information about when and why a worker was quarantined.
	QuarantineDetails struct {

		// The clientId of the client that made the request to quarantine the worker.
		ClientID string `json:"clientId"`

		// Usually a reason for the quarantine.
		QuarantineInfo string `json:"quarantineInfo"`

		// Value of the worker's quarantineUntil property at the moment of the quarantine.
		QuarantineUntil tcclient.Time `json:"quarantineUntil"`

		// Time when the quarantine was updated.
		UpdatedAt tcclient.Time `json:"updatedAt"`
	}

	// Request body to `registerWorker`.
	RegisterWorkerRequest struct {

		// The provider that had started the worker and responsible for managing it.
		// Can be different from the provider that's currently in the worker pool config.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// Worker group to which this worker belongs
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Worker ID
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`

		// Proof that this call is coming from the worker identified by the other fields.
		// The form of this proof varies depending on the provider type.
		//
		// One of:
		//   * GoogleProviderType
		//   * StaticProviderType1
		//   * AwsProviderType
		//   * AzureProviderType
		WorkerIdentityProof json.RawMessage `json:"workerIdentityProof"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId"`
	}

	// Response body to `registerWorker`.
	RegisterWorkerResponse struct {

		// The credentials the worker
		// will need to perform its work.  Specifically, credentials with scopes
		// * `assume:worker-pool:<workerPoolId>`
		// * `assume:worker-id:<workerGroup>/<workerId>`
		// * `queue:worker-id:<workerGroup>/<workerId>`
		// * `secrets:get:worker-pool:<workerPoolId>`
		// * `queue:claim-work:<workerPoolId>`
		// * `worker-manager:remove-worker:<workerPoolId>/<workerGroup>/<workerId>`
		Credentials Credentials `json:"credentials"`

		// Time at which the included credentials will expire.  Workers must either
		// re-register (for static workers) or terminate (for dynamically
		// provisioned workers) before this time.
		Expires tcclient.Time `json:"expires"`

		// A secret value generated by worker-manager that can be used in the call to `reregisterWorker`.
		// For more information, refer to https://docs.taskcluster.net/docs/reference/core/worker-manager#reregistration.
		//
		// Syntax:     ^[a-zA-Z0-9_-]{44}$
		Secret string `json:"secret"`

		// This value is supplied unchanged to the worker from the worker-pool configuration.
		// The expectation is that the worker will merge this information with configuration from other sources,
		// and this is precisely what [worker-runner](https://docs.taskcluster.net/docs/reference/workers/worker-runner) does.
		// This property must not be used for secret configuration, as it is visible both in the worker pool configuration and in the worker instance's metadata.
		// Instead, put secret configuration in the [secrets service](https://docs.taskcluster.net/docs/reference/workers/worker-runner).
		//
		// Additional properties allowed
		WorkerConfig json.RawMessage `json:"workerConfig"`
	}

	// Request body to `reregisterWorker`.
	ReregisterWorkerRequest struct {

		// The secret value that was last configured in `registerWorker` (in the case of a newly registerd worker) or
		// `reregisterWorker`.
		// For more information, refer to https://docs.taskcluster.net/docs/reference/core/worker-manager#reregistration.
		//
		// Syntax:     ^[a-zA-Z0-9_-]{44}$
		Secret string `json:"secret"`

		// Worker group to which this worker belongs
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Worker ID
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId"`
	}

	// Response body to `reregisterWorker`.
	ReregisterWorkerResponse struct {

		// The credentials the worker
		// will need to perform its work. Specifically, credentials with scopes
		// * `assume:worker-pool:<workerPoolId>`
		// * `assume:worker-id:<workerGroup>/<workerId>`
		// * `queue:worker-id:<workerGroup>/<workerId>`
		// * `secrets:get:worker-pool:<workerPoolId>`
		// * `queue:claim-work:<workerPoolId>`
		// * `worker-manager:remove-worker:<workerPoolId>/<workerGroup>/<workerId>`
		// * `worker-manager:reregister-worker:<workerPoolId>/<workerGroup>/<workerId>`
		Credentials Credentials1 `json:"credentials"`

		// Time at which the included credentials will expire. Workers must
		// re-register before this time.
		Expires tcclient.Time `json:"expires"`

		// The next secret value needed to reregister the worker (in `reregisterWorker).
		// For more information, refer to https://docs.taskcluster.net/docs/reference/core/worker-manager#reregistration.
		//
		// Syntax:     ^[a-zA-Z0-9_-]{44}$
		Secret string `json:"secret"`
	}

	// Provider-specific information
	StaticProviderType struct {

		// A secret value shared with the worker.  This value must be passed in the `workerIdentityProof` of the `registerWorker` method.
		// The ideal way to generate a secret of this form is `slugid() + slugid()`.
		//
		// Secrets are traded for Taskcluster credentials, and should be treated with similar care.
		// Each worker should have a distinct secret.
		//
		// Syntax:     ^[a-zA-Z0-9_-]{44}$
		StaticSecret string `json:"staticSecret"`
	}

	// Proof that this call is coming from the worker identified by the other fields.
	// The form of this proof varies depending on the provider type.
	StaticProviderType1 struct {

		// The secret value that was configured when the worker was created (in `createWorker`).
		//
		// Syntax:     ^[a-zA-Z0-9_-]{44}$
		StaticSecret string `json:"staticSecret"`
	}

	// Required task metadata
	TaskMetadata struct {

		// Human readable description of the task, please **explain** what the
		// task does. A few lines of documentation is not going to hurt you.
		//
		// Max length: 32768
		Description string `json:"description"`

		// Human readable name of task, used to very briefly given an idea about
		// what the task does.
		//
		// Max length: 255
		Name string `json:"name"`

		// Entity who caused this task, not necessarily a person with email who did
		// `hg push` as it could be automation bots as well. The entity we should
		// contact to ask why this task is here.
		//
		// Max length: 255
		Owner string `json:"owner"`

		// Link to source of this task, should specify a file, revision and
		// repository. This should be place someone can go an do a git/hg blame
		// to who came up with recipe for this task.
		//
		// Syntax:     ^(https?|ssh)://
		// Max length: 4096
		Source string `json:"source"`
	}

	// A run of a task.
	TaskRun struct {

		// Id of this task run, `run-id`s always starts from `0`
		//
		// Mininum:    0
		// Maximum:    1000
		RunID int64 `json:"runId"`

		// Unique task identifier, this is UUID encoded as
		// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
		// stripped of `=` padding.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`
	}

	Var struct {

		// The id of this provider
		ProviderID string `json:"providerId"`

		// The provider implementation underlying this provider
		ProviderType string `json:"providerType"`
	}

	Worker struct {

		// Number of tasks this worker can handle at once. A worker capacity of 0 means
		// the worker is not managed by worker manager and is only known to the queue, the
		// true capacity is not known.
		//
		// Mininum:    0
		Capacity int64 `json:"capacity,omitempty"`

		// Date of the first time this worker claimed a task.
		FirstClaim tcclient.Time `json:"firstClaim"`

		// Date of the last time this worker was seen active. Updated each time a worker calls
		// `queue.claimWork`, `queue.reclaimTask`, and `queue.declareWorker` for this task queue.
		// `lastDateActive` is updated every half hour but may be off by up-to half an hour.
		// Nonetheless, `lastDateActive` is a good indicator of when the worker was last seen active.
		// This defaults to null in the database, and is set to the current time when the worker
		// is first seen.
		LastDateActive tcclient.Time `json:"lastDateActive,omitempty"`

		// A run of a task.
		LatestTask TaskRun `json:"latestTask,omitempty"`

		// The provider that had started the worker and responsible for managing it.
		// Can be different from the provider that's currently in the worker pool config.
		// A providerId of "none" is used when the worker is not managed by worker manager.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId,omitempty"`

		// Quarantining a worker allows the machine to remain alive but not accept jobs.
		// Once the quarantineUntil time has elapsed, the worker resumes accepting jobs.
		// Note that a quarantine can be lifted by setting `quarantineUntil` to the present time (or
		// somewhere in the past).
		QuarantineUntil tcclient.Time `json:"quarantineUntil,omitempty"`

		// A string specifying the state this worker is in so far as worker-manager knows.
		// A "requested" worker is in the process of starting up, and if successful will enter
		// the "running" state once it has registered with the `registerWorker` API method.  A
		// "stopping" worker is in the process of shutting down and deleting resources, while
		// a "stopped" worker is completely stopped.  Stopped workers are kept for historical
		// purposes and are purged when they expire.  Note that some providers transition workers
		// directly from "running" to "stopped".
		// An "standalone" worker is a worker that is not managed by worker-manager, these workers
		// are only known by the queue.
		//
		// Possible values:
		//   * "requested"
		//   * "running"
		//   * "stopping"
		//   * "stopped"
		//   * "standalone"
		State string `json:"state,omitempty"`

		// Identifier for the worker group containing this worker.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Identifier for this worker (unique within this worker group).
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`

		// Unique identifier for a worker pool
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId,omitempty"`
	}

	// Actions provide a generic mechanism to expose additional features of a
	// provisioner, worker type, or worker to Taskcluster clients.
	//
	// An action is comprised of metadata describing the feature it exposes,
	// together with a webhook for triggering it.
	//
	// The Taskcluster tools site, for example, retrieves actions when displaying
	// provisioners, worker types and workers. It presents the provisioner/worker
	// type/worker specific actions to the user. When the user triggers an action,
	// the web client takes the registered webhook, substitutes parameters into the
	// URL (see `url`), signs the requests with the Taskcluster credentials of the
	// user operating the web interface, and issues the HTTP request.
	//
	// The level to which the action relates (provisioner, worker type, worker) is
	// called the action context. All actions, regardless of the action contexts,
	// are registered against the provisioner when calling
	// `queue.declareProvisioner`.
	//
	// The action context is used by the web client to determine where in the web
	// interface to present the action to the user as follows:
	//
	// | `context`   | Tool where action is displayed |
	// |-------------|--------------------------------|
	// | provisioner | Provisioner Explorer           |
	// | worker-type | Workers Explorer               |
	// | worker      | Worker Explorer                |
	//
	// See [actions docs](/docs/reference/platform/taskcluster-queue/docs/actions)
	// for more information.
	WorkerAction struct {

		// Only actions with the context `worker` are included.
		//
		// Possible values:
		//   * "worker"
		Context string `json:"context"`

		// Description of the provisioner.
		Description string `json:"description"`

		// Method to indicate the desired action to be performed for a given resource.
		//
		// Possible values:
		//   * "POST"
		//   * "PUT"
		//   * "DELETE"
		//   * "PATCH"
		Method string `json:"method"`

		// Short names for things like logging/error messages.
		Name string `json:"name"`

		// Appropriate title for any sort of Modal prompt.
		Title json.RawMessage `json:"title"`

		// When an action is triggered, a request is made using the `url` and `method`.
		// Depending on the `context`, the following parameters will be substituted in the url:
		//
		// | `context`   | Path parameters                                          |
		// |-------------|----------------------------------------------------------|
		// | provisioner | <provisionerId>                                          |
		// | worker-type | <provisionerId>, <workerType>                            |
		// | worker      | <provisionerId>, <workerType>, <workerGroup>, <workerId> |
		//
		// _Note: The request needs to be signed with the user's Taskcluster credentials._
		URL string `json:"url"`
	}

	// Request to create or update a worker. Capacity will default to 1 if not specified.
	WorkerCreationUpdateRequest struct {

		// Number of tasks this worker can handle at once
		//
		// Mininum:    1
		Capacity int64 `json:"capacity,omitempty"`

		// Date and time when this worker will be deleted from the DB
		Expires tcclient.Time `json:"expires"`

		// Provider-specific information
		//
		// One of:
		//   * StaticProviderType
		ProviderInfo json.RawMessage `json:"providerInfo,omitempty"`
	}

	// A report of an error from a worker.  This will be recorded with kind
	// `worker-error`.
	//
	// The worker's `workerGroup` and `workerId` will be added to `extra`.
	WorkerErrorReport struct {

		// A longer description of what occured in the error.
		//
		// Max length: 10240
		Description string `json:"description"`

		// Any extra structured information about this error
		//
		// Additional properties allowed
		Extra json.RawMessage `json:"extra"`

		// A general machine-readable way to identify this sort of error.
		//
		// Syntax:     [-a-z0-9]+
		// Max length: 128
		Kind string `json:"kind"`

		// A human-readable version of `kind`.
		//
		// Max length: 128
		Title string `json:"title"`

		// Worker group to which this worker belongs
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Worker ID
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`
	}

	// A complete worker definition.
	WorkerFullDefinition struct {

		// Number of tasks this worker can handle at once
		//
		// Mininum:    1
		Capacity int64 `json:"capacity"`

		// Date and time when this worker was created
		Created tcclient.Time `json:"created"`

		// Date and time when this worker will be deleted from the DB
		Expires tcclient.Time `json:"expires"`

		// Date and time when the state of this worker was verified with a cloud api.
		// For providers with nothing to check, this will just be permanently set to the
		// time the worker was created.
		LastChecked tcclient.Time `json:"lastChecked"`

		// Date and time when this worker last changed state
		LastModified tcclient.Time `json:"lastModified"`

		// The provider that had started the worker and responsible for managing it.
		// Can be different from the provider that's currently in the worker pool config.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// A string specifying the state this worker is in so far as worker-manager knows.
		// A "requested" worker is in the process of starting up, and if successful will enter
		// the "running" state once it has registered with the `registerWorker` API method.  A
		// "stopping" worker is in the process of shutting down and deleting resources, while
		// a "stopped" worker is completely stopped.  Stopped workers are kept for historical
		// purposes and are purged when they expire.  Note that some providers transition workers
		// directly from "running" to "stopped".
		//
		// Possible values:
		//   * "requested"
		//   * "running"
		//   * "stopping"
		//   * "stopped"
		State string `json:"state"`

		// Worker group to which this worker belongs
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Worker ID
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId"`
	}

	// A list of workers in a given worker pool
	WorkerListInAGivenWorkerPool struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of workers in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of all workers in a given worker pool
		Workers []WorkerFullDefinition `json:"workers"`
	}

	// Fields that are defined by a user for a worker pool.
	// Used to create worker-pool definitions. There is a larger
	// set of fields for viewing since some parts are generated
	// by the service.
	WorkerPoolDefinition struct {

		// Additional properties allowed
		Config json.RawMessage `json:"config"`

		// A description of this worker pool.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, the owner should be emailed on provisioning errors
		EmailOnError bool `json:"emailOnError"`

		// An email address to notify when there are provisioning errors for this
		// worker pool.
		Owner string `json:"owner"`

		// The provider responsible for managing this worker pool.
		//
		// If this value is `"null-provider"`, then the worker pool is pending deletion
		// once all existing workers have terminated.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`
	}

	// Fields that are defined by a user for a worker pool.
	// Used to modify worker-pool definitions.
	//
	// The `workerPoolId`, `created`, and `lastModified` fields are optional and
	// allowed only to ease the common practice of getting a worker pool definition
	// with `workerPool(..)`, modifying it, and writing it back with
	// `updateWorkerPool(..).  `workerPoolId` must be correct if
	// supplied, and the values of `created` and `lastModified` are ignored.
	WorkerPoolDefinition1 struct {

		// Additional properties allowed
		Config json.RawMessage `json:"config"`

		// Ignored on update
		Created tcclient.Time `json:"created,omitempty"`

		// A description of this worker pool.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, the owner should be emailed on provisioning errors
		EmailOnError bool `json:"emailOnError"`

		// Ignored on update
		LastModified tcclient.Time `json:"lastModified,omitempty"`

		// An email address to notify when there are provisioning errors for this
		// worker pool.
		Owner string `json:"owner"`

		// The provider responsible for managing this worker pool.
		//
		// If this value is `"null-provider"`, then the worker pool is pending deletion
		// once all existing workers have terminated.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId,omitempty"`
	}

	// A complete worker pool error definition.
	WorkerPoolError struct {

		// A longer description of what occured in the error.
		//
		// Max length: 10240
		Description string `json:"description"`

		// An arbitary unique identifier for this error
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		ErrorID string `json:"errorId"`

		// Any extra structured information about this error
		//
		// Additional properties allowed
		Extra json.RawMessage `json:"extra"`

		// A general machine-readable way to identify this sort of error.
		//
		// Syntax:     [-a-z0-9]+
		// Max length: 128
		Kind string `json:"kind"`

		// Date and time when this error was reported
		Reported tcclient.Time `json:"reported"`

		// A human-readable version of `kind`.
		//
		// Max length: 128
		Title string `json:"title"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId"`
	}

	// A list of worker pool errors
	WorkerPoolErrorList struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of worker-types in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of worker pool errors
		WorkerPoolErrors []WorkerPoolError `json:"workerPoolErrors"`
	}

	// A complete worker pool definition.
	WorkerPoolFullDefinition struct {

		// Additional properties allowed
		Config json.RawMessage `json:"config"`

		// Date and time when this worker pool was created
		Created tcclient.Time `json:"created"`

		// Total capacity available across all workers for this worker pool that are currently not "stopped"
		//
		// Mininum:    0
		CurrentCapacity int64 `json:"currentCapacity"`

		// A description of this worker pool.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, the owner should be emailed on provisioning errors
		EmailOnError bool `json:"emailOnError"`

		// Date and time when this worker pool was last updated
		LastModified tcclient.Time `json:"lastModified"`

		// An email address to notify when there are provisioning errors for this
		// worker pool.
		Owner string `json:"owner"`

		// The provider responsible for managing this worker pool.
		//
		// If this value is `"null-provider"`, then the worker pool is pending deletion
		// once all existing workers have terminated.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// Total capacity available across all workers for this worker pool with state "requested"
		//
		// Mininum:    0
		RequestedCapacity int64 `json:"requestedCapacity"`

		// Total worker count in "requested" state for this worker pool
		//
		// Mininum:    0
		RequestedCount int64 `json:"requestedCount"`

		// Total capacity available across all workers for this worker pool with state "running"
		//
		// Mininum:    0
		RunningCapacity int64 `json:"runningCapacity"`

		// Total worker count in "running" state for this worker pool
		//
		// Mininum:    0
		RunningCount int64 `json:"runningCount"`

		// Total capacity available across all workers for this worker pool with state "stopped"
		//
		// Mininum:    0
		StoppedCapacity int64 `json:"stoppedCapacity"`

		// Total worker count in "stopped" state for this worker pool
		//
		// Mininum:    0
		StoppedCount int64 `json:"stoppedCount"`

		// Total capacity available across all workers for this worker pool with state "stopping"
		//
		// Mininum:    0
		StoppingCapacity int64 `json:"stoppingCapacity"`

		// Total worker count in "stopping" state for this worker pool
		//
		// Mininum:    0
		StoppingCount int64 `json:"stoppingCount"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId,omitempty"`
	}

	// A list of worker pools
	WorkerPoolList struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of worker-types in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of all worker pools
		WorkerPools []WorkerPoolFullDefinition `json:"workerPools"`
	}

	// Response containing information about a worker.
	WorkerResponse struct {
		Actions []WorkerAction `json:"actions"`

		// Number of tasks this worker can handle at once. A worker capacity of 0 means
		// the worker is not managed by worker manager and is only known to the queue, the
		// true capacity is not known.
		//
		// Mininum:    0
		Capacity int64 `json:"capacity,omitempty"`

		// Date and time after which the worker will be automatically
		// deleted by the queue.
		Expires tcclient.Time `json:"expires"`

		// Date of the first time this worker claimed a task.
		FirstClaim tcclient.Time `json:"firstClaim"`

		// Date of the last time this worker was seen active. Updated each time a worker calls
		// `queue.claimWork`, `queue.reclaimTask`, and `queue.declareWorker` for this task queue.
		// `lastDateActive` is updated every half hour but may be off by up-to half an hour.
		// Nonetheless, `lastDateActive` is a good indicator of when the worker was last seen active.
		// This defaults to null in the database, and is set to the current time when the worker
		// is first seen.
		LastDateActive tcclient.Time `json:"lastDateActive,omitempty"`

		// The provider that had started the worker and responsible for managing it.
		// Can be different from the provider that's currently in the worker pool config.
		// A providerId of "none" is used when the worker is not managed by worker manager.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId,omitempty"`

		// Unique identifier for a provisioner, that can supply specified
		// `workerType`. Deprecation is planned for this property as it
		// will be replaced, together with `workerType`, by the new
		// identifier `workerPoolId`.
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}$
		ProvisionerID string `json:"provisionerId"`

		// This is a list of changes to the worker's quarantine status. Each entry is an object
		// containing information about the time, clientId and reason for the change.
		QuarantineDetails []QuarantineDetails `json:"quarantineDetails,omitempty"`

		// Quarantining a worker allows the machine to remain alive but not accept jobs.
		// Once the quarantineUntil time has elapsed, the worker resumes accepting jobs.
		// Note that a quarantine can be lifted by setting `quarantineUntil` to the present time (or
		// somewhere in the past).
		QuarantineUntil tcclient.Time `json:"quarantineUntil,omitempty"`

		// List of 20 most recent tasks claimed by the worker.
		RecentTasks []TaskRun `json:"recentTasks"`

		// A string specifying the state this worker is in so far as worker-manager knows.
		// A "requested" worker is in the process of starting up, and if successful will enter
		// the "running" state once it has registered with the `registerWorker` API method.  A
		// "stopping" worker is in the process of shutting down and deleting resources, while
		// a "stopped" worker is completely stopped.  Stopped workers are kept for historical
		// purposes and are purged when they expire.  Note that some providers transition workers
		// directly from "running" to "stopped".
		// An "standalone" worker is a worker that is not managed by worker-manager, these workers
		// are only known by the queue.
		//
		// Possible values:
		//   * "requested"
		//   * "running"
		//   * "stopping"
		//   * "stopped"
		//   * "standalone"
		State string `json:"state,omitempty"`

		// Identifier for group that worker who executes this run is a part of,
		// this identifier is mainly used for efficient routing.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Identifier for worker evaluating this run within given
		// `workerGroup`.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`

		// Unique identifier for a worker pool
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId,omitempty"`

		// Unique identifier for a worker-type within a specific
		// provisioner. Deprecation is planned for this property as it will
		// be replaced, together with `provisionerId`, by the new
		// identifier `workerPoolId`.
		//
		// Syntax:     ^[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerType string `json:"workerType"`
	}
)
