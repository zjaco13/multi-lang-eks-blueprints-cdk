use async_fn_traits::AsyncFnOnce1;
use tonic::transport::Channel;

use crate::codegen::cluster_service_client::ClusterServiceClient;

// Pass in a build function that builds an EKS Blueprint and run will connect a client to the grpc server
// and run the build function on the client
pub async fn run<F>(build: F) -> Result<(), tonic::Status>
where
    F: for<'a> AsyncFnOnce1<
        &'a mut ClusterServiceClient<Channel>,
        Output = Result<(), tonic::Status>,
    >,
{
    let channel = tonic::transport::Channel::from_static("http://0.0.0.0:50051")
        .connect()
        .await
        .expect("connection failed");
    let mut client = ClusterServiceClient::new(channel);

    build(&mut client).await?;
    Ok(())
}
