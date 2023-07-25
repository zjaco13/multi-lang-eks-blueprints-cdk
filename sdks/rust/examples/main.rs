use rust_sdk::{builder::{ self}, proto::{cluster_service_client::ClusterServiceClient, self}};
use tonic::transport::Channel;


#[tokio::main]
async fn main() {
    builder::await_build(build).await.expect("Error in build");
}

async fn build(client: &mut ClusterServiceClient<Channel>) -> Result<tonic::Response<proto::ApiResponse>, tonic::Status> {
    client.create_cluster(proto::CreateClusterRequest{id: String::from("test-from-rust"), name: None}).await?;
    client.add_teams(proto::AddTeamsRequest{teams: vec![proto::Team{team: Some(proto::team::Team::GenericTeam(proto::GenericTeam{name: String::from("hello")}))}]}).await?;
    client.add_addons(proto::AddAddonsRequest{addons: vec![proto::Addon{addon: Some(proto::addon::Addon::AckAddOn(proto::AckAddOn{id: Some(String::from("hi")), service_name: String::from("eks")}))}]}).await?;
    client.build_cluster(proto::BuildClusterRequest{cluster_name: String::from("test-from-rust"), account: None, region:None}).await

}
