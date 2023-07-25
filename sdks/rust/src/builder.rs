use async_fn_traits::AsyncFnOnce1;
use tonic::transport::Channel;

use crate::proto::{cluster_service_client::ClusterServiceClient, self};


pub async fn await_build<F>(f: F) -> Result<tonic::Response<proto::ApiResponse>,tonic::Status> 
where
    F: for<'a> AsyncFnOnce1<&'a mut ClusterServiceClient<Channel>, Output = Result<tonic::Response<proto::ApiResponse>, tonic::Status>>
{
    let channel = tonic::transport::Channel::from_static("http://[::1]:50051").connect().await.expect("no connect");
    let mut client = ClusterServiceClient::new(channel);

    let res = f(&mut client).await;
    res
}
