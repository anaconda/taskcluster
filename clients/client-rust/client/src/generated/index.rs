#![allow(unused_imports)]
#![cfg_attr(rustfmt, rustfmt_skip)]
/* THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT */
use crate::{Client, ClientBuilder, Credentials, Retry};
use anyhow::Error;
use serde_json::Value;
use std::time::Duration;
use crate::util::urlencode;

/// Index Service
///
/// The index service is responsible for indexing tasks. The service ensures that
/// tasks can be located by user-defined names.
///
/// As described in the service documentation, tasks are typically indexed via Pulse
/// messages, so the most common use of API methods is to read from the index.
///
/// Slashes (`/`) aren't allowed in index paths.
pub struct Index {
    /// The underlying client used to make API calls for this service.
    pub client: Client
}

#[allow(non_snake_case)]
impl Index {
    /// Create a new Index instance, based on the given client builder
    pub fn new<CB: Into<ClientBuilder>>(client_builder: CB) -> Result<Self, Error> {
        Ok(Self{
            client: client_builder
                .into()
                .path_prefix("api/index/v1/")
                .build()?,
        })
    }

    /// Ping Server
    ///
    /// Respond without doing anything.
    /// This endpoint is used to check that the service is up.
    pub async fn ping(&self) -> Result<(), Error> {
        let method = "GET";
        let (path, query) = Self::ping_details();
        let body = None;
        let resp = self.client.request(method, path, query, body).await?;
        resp.bytes().await?;
        Ok(())
    }

    /// Generate an unsigned URL for the ping endpoint
    pub fn ping_url(&self) -> Result<String, Error> {
        let (path, query) = Self::ping_details();
        self.client.make_url(path, query)
    }

    /// Generate a signed URL for the ping endpoint
    pub fn ping_signed_url(&self, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::ping_details();
        self.client.make_signed_url(path, query, ttl)
    }

    /// Determine the HTTP request details for ping
    fn ping_details<'a>() -> (&'static str, Option<Vec<(&'static str, &'a str)>>) {
        let path = "ping";
        let query = None;

        (path, query)
    }

    /// Load Balancer Heartbeat
    ///
    /// Respond without doing anything.
    /// This endpoint is used to check that the service is up.
    pub async fn lbheartbeat(&self) -> Result<(), Error> {
        let method = "GET";
        let (path, query) = Self::lbheartbeat_details();
        let body = None;
        let resp = self.client.request(method, path, query, body).await?;
        resp.bytes().await?;
        Ok(())
    }

    /// Generate an unsigned URL for the lbheartbeat endpoint
    pub fn lbheartbeat_url(&self) -> Result<String, Error> {
        let (path, query) = Self::lbheartbeat_details();
        self.client.make_url(path, query)
    }

    /// Generate a signed URL for the lbheartbeat endpoint
    pub fn lbheartbeat_signed_url(&self, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::lbheartbeat_details();
        self.client.make_signed_url(path, query, ttl)
    }

    /// Determine the HTTP request details for lbheartbeat
    fn lbheartbeat_details<'a>() -> (&'static str, Option<Vec<(&'static str, &'a str)>>) {
        let path = "__lbheartbeat__";
        let query = None;

        (path, query)
    }

    /// Taskcluster Version
    ///
    /// Respond with the JSON version object.
    /// https://github.com/mozilla-services/Dockerflow/blob/main/docs/version_object.md
    pub async fn version(&self) -> Result<(), Error> {
        let method = "GET";
        let (path, query) = Self::version_details();
        let body = None;
        let resp = self.client.request(method, path, query, body).await?;
        resp.bytes().await?;
        Ok(())
    }

    /// Generate an unsigned URL for the version endpoint
    pub fn version_url(&self) -> Result<String, Error> {
        let (path, query) = Self::version_details();
        self.client.make_url(path, query)
    }

    /// Generate a signed URL for the version endpoint
    pub fn version_signed_url(&self, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::version_details();
        self.client.make_signed_url(path, query, ttl)
    }

    /// Determine the HTTP request details for version
    fn version_details<'a>() -> (&'static str, Option<Vec<(&'static str, &'a str)>>) {
        let path = "__version__";
        let query = None;

        (path, query)
    }

    /// Find Indexed Task
    ///
    /// Find a task by index path, returning the highest-rank task with that path. If no
    /// task exists for the given path, this API end-point will respond with a 404 status.
    pub async fn findTask(&self, indexPath: &str) -> Result<Value, Error> {
        let method = "GET";
        let (path, query) = Self::findTask_details(indexPath);
        let body = None;
        let resp = self.client.request(method, &path, query, body).await?;
        Ok(resp.json().await?)
    }

    /// Generate an unsigned URL for the findTask endpoint
    pub fn findTask_url(&self, indexPath: &str) -> Result<String, Error> {
        let (path, query) = Self::findTask_details(indexPath);
        self.client.make_url(&path, query)
    }

    /// Generate a signed URL for the findTask endpoint
    pub fn findTask_signed_url(&self, indexPath: &str, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::findTask_details(indexPath);
        self.client.make_signed_url(&path, query, ttl)
    }

    /// Determine the HTTP request details for findTask
    fn findTask_details<'a>(indexPath: &'a str) -> (String, Option<Vec<(&'static str, &'a str)>>) {
        let path = format!("task/{}", urlencode(indexPath));
        let query = None;

        (path, query)
    }

    /// List Namespaces
    ///
    /// List the namespaces immediately under a given namespace.
    ///
    /// This endpoint
    /// lists up to 1000 namespaces. If more namespaces are present, a
    /// `continuationToken` will be returned, which can be given in the next
    /// request. For the initial request, the payload should be an empty JSON
    /// object.
    pub async fn listNamespaces(&self, namespace: &str, continuationToken: Option<&str>, limit: Option<&str>) -> Result<Value, Error> {
        let method = "GET";
        let (path, query) = Self::listNamespaces_details(namespace, continuationToken, limit);
        let body = None;
        let resp = self.client.request(method, &path, query, body).await?;
        Ok(resp.json().await?)
    }

    /// Generate an unsigned URL for the listNamespaces endpoint
    pub fn listNamespaces_url(&self, namespace: &str, continuationToken: Option<&str>, limit: Option<&str>) -> Result<String, Error> {
        let (path, query) = Self::listNamespaces_details(namespace, continuationToken, limit);
        self.client.make_url(&path, query)
    }

    /// Generate a signed URL for the listNamespaces endpoint
    pub fn listNamespaces_signed_url(&self, namespace: &str, continuationToken: Option<&str>, limit: Option<&str>, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::listNamespaces_details(namespace, continuationToken, limit);
        self.client.make_signed_url(&path, query, ttl)
    }

    /// Determine the HTTP request details for listNamespaces
    fn listNamespaces_details<'a>(namespace: &'a str, continuationToken: Option<&'a str>, limit: Option<&'a str>) -> (String, Option<Vec<(&'static str, &'a str)>>) {
        let path = format!("namespaces/{}", urlencode(namespace));
        let mut query = None;
        if let Some(q) = continuationToken {
            query.get_or_insert_with(Vec::new).push(("continuationToken", q));
        }
        if let Some(q) = limit {
            query.get_or_insert_with(Vec::new).push(("limit", q));
        }

        (path, query)
    }

    /// List Tasks
    ///
    /// List the tasks immediately under a given namespace.
    ///
    /// This endpoint
    /// lists up to 1000 tasks. If more tasks are present, a
    /// `continuationToken` will be returned, which can be given in the next
    /// request. For the initial request, the payload should be an empty JSON
    /// object.
    ///
    /// **Remark**, this end-point is designed for humans browsing for tasks, not
    /// services, as that makes little sense.
    pub async fn listTasks(&self, namespace: &str, continuationToken: Option<&str>, limit: Option<&str>) -> Result<Value, Error> {
        let method = "GET";
        let (path, query) = Self::listTasks_details(namespace, continuationToken, limit);
        let body = None;
        let resp = self.client.request(method, &path, query, body).await?;
        Ok(resp.json().await?)
    }

    /// Generate an unsigned URL for the listTasks endpoint
    pub fn listTasks_url(&self, namespace: &str, continuationToken: Option<&str>, limit: Option<&str>) -> Result<String, Error> {
        let (path, query) = Self::listTasks_details(namespace, continuationToken, limit);
        self.client.make_url(&path, query)
    }

    /// Generate a signed URL for the listTasks endpoint
    pub fn listTasks_signed_url(&self, namespace: &str, continuationToken: Option<&str>, limit: Option<&str>, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::listTasks_details(namespace, continuationToken, limit);
        self.client.make_signed_url(&path, query, ttl)
    }

    /// Determine the HTTP request details for listTasks
    fn listTasks_details<'a>(namespace: &'a str, continuationToken: Option<&'a str>, limit: Option<&'a str>) -> (String, Option<Vec<(&'static str, &'a str)>>) {
        let path = format!("tasks/{}", urlencode(namespace));
        let mut query = None;
        if let Some(q) = continuationToken {
            query.get_or_insert_with(Vec::new).push(("continuationToken", q));
        }
        if let Some(q) = limit {
            query.get_or_insert_with(Vec::new).push(("limit", q));
        }

        (path, query)
    }

    /// Insert Task into Index
    ///
    /// Insert a task into the index.  If the new rank is less than the existing rank
    /// at the given index path, the task is not indexed but the response is still 200 OK.
    ///
    /// Please see the introduction above for information
    /// about indexing successfully completed tasks automatically using custom routes.
    pub async fn insertTask(&self, namespace: &str, payload: &Value) -> Result<Value, Error> {
        let method = "PUT";
        let (path, query) = Self::insertTask_details(namespace);
        let body = Some(payload);
        let resp = self.client.request(method, &path, query, body).await?;
        Ok(resp.json().await?)
    }

    /// Determine the HTTP request details for insertTask
    fn insertTask_details<'a>(namespace: &'a str) -> (String, Option<Vec<(&'static str, &'a str)>>) {
        let path = format!("task/{}", urlencode(namespace));
        let query = None;

        (path, query)
    }

    /// Remove Task from Index
    ///
    /// Remove a task from the index.  This is intended for administrative use,
    /// where an index entry is no longer appropriate.  The parent namespace is
    /// not automatically deleted.  Index entries with lower rank that were
    /// previously inserted will not re-appear, as they were never stored.
    pub async fn deleteTask(&self, namespace: &str) -> Result<(), Error> {
        let method = "DELETE";
        let (path, query) = Self::deleteTask_details(namespace);
        let body = None;
        let resp = self.client.request(method, &path, query, body).await?;
        resp.bytes().await?;
        Ok(())
    }

    /// Determine the HTTP request details for deleteTask
    fn deleteTask_details<'a>(namespace: &'a str) -> (String, Option<Vec<(&'static str, &'a str)>>) {
        let path = format!("task/{}", urlencode(namespace));
        let query = None;

        (path, query)
    }

    /// Get Artifact From Indexed Task
    ///
    /// Find a task by index path and redirect to the artifact on the most recent
    /// run with the given `name`.
    ///
    /// Note that multiple calls to this endpoint may return artifacts from differen tasks
    /// if a new task is inserted into the index between calls. Avoid using this method as
    /// a stable link to multiple, connected files if the index path does not contain a
    /// unique identifier.  For example, the following two links may return unrelated files:
    /// * https://tc.example.com/api/index/v1/task/some-app.win64.latest.installer/artifacts/public/installer.exe`
    /// * https://tc.example.com/api/index/v1/task/some-app.win64.latest.installer/artifacts/public/debug-symbols.zip`
    ///
    /// This problem be remedied by including the revision in the index path or by bundling both
    /// installer and debug symbols into a single artifact.
    ///
    /// If no task exists for the given index path, this API end-point responds with 404.
    pub async fn findArtifactFromTask(&self, indexPath: &str, name: &str) -> Result<(), Error> {
        let method = "GET";
        let (path, query) = Self::findArtifactFromTask_details(indexPath, name);
        let body = None;
        let resp = self.client.request(method, &path, query, body).await?;
        resp.bytes().await?;
        Ok(())
    }

    /// Generate an unsigned URL for the findArtifactFromTask endpoint
    pub fn findArtifactFromTask_url(&self, indexPath: &str, name: &str) -> Result<String, Error> {
        let (path, query) = Self::findArtifactFromTask_details(indexPath, name);
        self.client.make_url(&path, query)
    }

    /// Generate a signed URL for the findArtifactFromTask endpoint
    pub fn findArtifactFromTask_signed_url(&self, indexPath: &str, name: &str, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::findArtifactFromTask_details(indexPath, name);
        self.client.make_signed_url(&path, query, ttl)
    }

    /// Determine the HTTP request details for findArtifactFromTask
    fn findArtifactFromTask_details<'a>(indexPath: &'a str, name: &'a str) -> (String, Option<Vec<(&'static str, &'a str)>>) {
        let path = format!("task/{}/artifacts/{}", urlencode(indexPath), urlencode(name));
        let query = None;

        (path, query)
    }

    /// Heartbeat
    ///
    /// Respond with a service heartbeat.
    ///
    /// This endpoint is used to check on backing services this service
    /// depends on.
    pub async fn heartbeat(&self) -> Result<(), Error> {
        let method = "GET";
        let (path, query) = Self::heartbeat_details();
        let body = None;
        let resp = self.client.request(method, path, query, body).await?;
        resp.bytes().await?;
        Ok(())
    }

    /// Generate an unsigned URL for the heartbeat endpoint
    pub fn heartbeat_url(&self) -> Result<String, Error> {
        let (path, query) = Self::heartbeat_details();
        self.client.make_url(path, query)
    }

    /// Generate a signed URL for the heartbeat endpoint
    pub fn heartbeat_signed_url(&self, ttl: Duration) -> Result<String, Error> {
        let (path, query) = Self::heartbeat_details();
        self.client.make_signed_url(path, query, ttl)
    }

    /// Determine the HTTP request details for heartbeat
    fn heartbeat_details<'a>() -> (&'static str, Option<Vec<(&'static str, &'a str)>>) {
        let path = "__heartbeat__";
        let query = None;

        (path, query)
    }
}
