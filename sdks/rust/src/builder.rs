use std::future::Future;

use tonic::transport::Channel;

use crate::proto::cluster_service_client::ClusterServiceClient;

pub async fn run_build<F, Fut>(build: F) -> ()
where for<'a >
    F: Fn(&mut ClusterServiceClient<Channel>) -> Fut,
    Fut: Future<Output = Result<(), Box<dyn std::error::Error>>>
{
    let mut client = ClusterServiceClient::connect("http://localhost:50051").await.expect("unable to connect to client");
    build(&mut client).await.expect("Build Fail");
} 
