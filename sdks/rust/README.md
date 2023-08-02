# Rust Client for Multi Language EKS Blueprints for cdk

## Run the example
To run the example, first start the server from https://github.com/aws-quickstart/cdk-eks-blueprints/tree/blueprints-api-support
```bash
git clone https://github.com/aws-quickstart/cdk-eks-blueprints.git
cd cdk-eks-blueprints
git checkout blueprints-api-support
make run-server
```

Then run the rust example by cloning this repository and running with make, cargo is required
```bash
make rust-example
```

## How to use as a package

Create a new rust project
```bash
cargo new example
cd example
```

Import the sdk and other dependencies
```bash
cargo add eks-blueprints-rust-sdk
cargo add tonic
cargo add tokio
```

Add this to a main.rs file
```rust
use multi_lang_eks_blueprints_rust_sdk::{
    builder::{self},
    codegen::{self, cluster_service_client::ClusterServiceClient},
};
use tonic::transport::Channel;

#[tokio::main]
async fn main() {
    builder::run(build).await.expect("Error in build");
}

async fn build(client: &mut ClusterServiceClient<Channel>) -> Result<(), tonic::Status> {
    ...
    Ok(())
}
```

Add function calls to the grpc server in the build function to build your EKS Blueprint

Run the server
```bash
git clone https://github.com/aws-quickstart/cdk-eks-blueprints.git
cd cdk-eks-blueprints
git checkout blueprints-api-support
make run-server
```

Run your code
```bash
cargo run
```

From the server, deploy your EKS Blueprint
```bash
npx cdk -a cdk.out/ deploy
```




