use rust_sdk::{builder::{ self}, proto::{cluster_service_client::ClusterServiceClient, self}};
use std::{future::Future, process::Output};
use tonic::transport::Channel;
use futures::{stream::FuturesOrdered, TryFutureExt};

#[tokio::main]
async fn main() {
    builder::await_build(build);
}

fn build<Fut>(futs: &mut FuturesOrdered<Fut>, client: &mut ClusterServiceClient<Channel>) -> Result<tonic::Response<proto::ApiResponse>, tonic::Status>
where
    Fut: Future<Output = Result<tonic::Response<proto::ApiResponse>, tonic::Status>>,
{
    let fut1 = client.create_cluster(proto::CreateClusterRequest{id: String::from("test-from-rust"), name: None});
    futs.push_back(fut1);
    let fut2 = fut1.and_then(|_| client.add_teams(proto::AddTeamsRequest{teams: vec![proto::Team{team: Some(proto::team::Team::GenericTeam(proto::GenericTeam{name: String::from("hello")}))}]}));
    let fut3 = fut2.and_then(|_| client.add_addons(proto::AddAddonsRequest{addons: vec![proto::Addon{addon: Some(proto::addon::Addon::AckAddOn(proto::AckAddOn{id: Some(String::from("hi")), service_name: String::from("eks")}))}]}));
    let fut4 = fut3.and_then(|_| client.build_cluster(proto::BuildClusterRequest{cluster_name: String::from("test-from-rust"), account: None, region:None}));
    fut4.await
}
