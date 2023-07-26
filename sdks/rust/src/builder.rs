use async_fn_traits::AsyncFnOnce1;
use tonic::transport::Channel;

use crate::codegen::cluster_service_client::ClusterServiceClient;


pub async fn await_build<F>(f: F) -> Result<(), tonic::Status> 
where
    F: for<'a> AsyncFnOnce1<&'a mut ClusterServiceClient<Channel>, Output = Result<(), tonic::Status>>
{
    let channel = tonic::transport::Channel::from_static("http://0.0.0.0:50051").connect().await.expect("connection failed");
    let mut client = ClusterServiceClient::new(channel);

    let res = f(&mut client).await;
    println!("{:?}", res);
    res
}
