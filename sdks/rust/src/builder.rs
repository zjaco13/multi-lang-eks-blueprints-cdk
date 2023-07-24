use core::fmt;
use std::{future::Future, error::Error};

use futures::{stream::FuturesOrdered, StreamExt};
use tonic::transport::Channel;

use crate::proto::{cluster_service_client::ClusterServiceClient, self};
use async_trait::async_trait;


pub async fn await_build<F, Fut>(f: F) -> Result<(), Box<dyn Error + Send + Sync + 'static>>
where
    F: FnOnce(&mut FuturesOrdered<Fut>, &mut ClusterServiceClient<Channel>) -> Result<tonic::Response<proto::ApiResponse>, tonic::Status> + Send +'static,
    Fut: Future<Output = Result<tonic::Response<proto::ApiResponse>, tonic::Status>>,
{
    let channel = tonic::transport::Channel::from_static("http://[::1]:50051").connect().await?;
    let mut futures_list = FuturesOrdered::new();
    let mut client = ClusterServiceClient::new(channel);

    let res = f(&mut futures_list, &mut client);
    res?;
    let r = futures_list.collect::<Vec<Result<tonic::Response<proto::ApiResponse>, tonic::Status>>>().await;
    for result in r {
        result?;
    }

    Ok(())
}
