use rust_sdk::builder;
use rust_sdk::proto;
use tonic::transport::Channel;

#[tokio::main]
async fn main() {
    builder::run_build(build).await;
}

async fn build(client: &mut proto::cluster_service_client::ClusterServiceClient<Channel>) -> Result<(), Box<dyn std::error::Error>>{
    let response = client.create_cluster(proto::CreateClusterRequest{id: String::from("test-from-rust"), name: None}).await?;
    println!("Response={:?}", response);
    let response = client.add_teams(proto::AddTeamsRequest{teams: vec![proto::Team{team: Some(proto::team::Team::GenericTeam(proto::GenericTeam{name: String::from("hello")}))}]}).await?;
    println!("Response={:?}", response);
    let response = client.add_addons(proto::AddAddonsRequest{addons: vec![proto::Addon{addon: Some(proto::addon::Addon::AckAddOn(proto::AckAddOn{id: Some(String::from("hi")), service_name: String::from("eks")}))}]}).await?;
    println!("Response={:?}", response);
    Ok(())
}
