use eks_blueprints_rust_sdk::{
    builder::{self},
    codegen::{self, cluster_service_client::ClusterServiceClient},
};
use tonic::transport::Channel;

#[tokio::main]
async fn main() {
    builder::run(build).await.expect("Error in build");
}

async fn build(client: &mut ClusterServiceClient<Channel>) -> Result<(), tonic::Status> {
    let res = client
        .create_cluster(codegen::CreateClusterRequest {
            id: String::from("test-from-rust"),
            name: None,
        })
        .await?;
    println!("{:?}", res);
    let res = client
        .add_application_team(codegen::AddApplicationTeamRequest {
            cluster_name: String::from("test-from-rust"),
            props: Some(codegen::TeamProps {
                name: String::from("app-team"),
            }),
        })
        .await?;
    println!("{:?}", res);
    let res = client
        .add_mng_cluster_provider(codegen::AddMngClusterProviderRequest {
            cluster_name: String::from("test-from-rust"),
            mng_cluster_provider: Some(codegen::MngClusterProvider {
                name: Some(String::from("cl_p")),
                version: String::from("1.27"),
            }),
        })
        .await?;
    println!("{:?}", res);
    let res = client
        .add_kube_proxy_add_on(codegen::AddKubeProxyAddOnRequest {
            cluster_name: String::from("test-from-rust"),
            kube_proxy_add_on: Some(codegen::KubeProxyAddOn { version: None }),
        })
        .await?;
    println!("{:?}", res);

    let res = client
        .add_vpc_provider(codegen::AddVpcProviderRequest {
            cluster_name: String::from("test-from-rust"),
            name: String::from("vpc"),
            vpc_provider: Some(codegen::VpcProvider { vpc_id: None }),
        })
        .await?;
    println!("{:?}", res);

    let res = client
        .build_cluster(codegen::BuildClusterRequest {
            cluster_name: String::from("test-from-rust"),
            account: None,
            region: None,
        })
        .await?;
    println!("{:?}", res);
    Ok(())
}
