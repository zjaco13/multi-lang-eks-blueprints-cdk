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
git clone https://github.com/zjaco13/multi-lang-eks-blueprints-cdk.git
cd multi-lang-eks-blueprints-cdk
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

Add this to a `main.rs` file
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

Add function calls to the gRPC server in the build function to build your EKS Blueprint

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
make deploy-server
```
## Contributing

### Publishing

Whenever changes are made to the protobufs, a new package will need to be released

To publish the package, first make an account on [crates.io](https://crates.io/), and you must be added to the [eks-blueprints-rust-sdk](https://crates.io/crates/eks-blueprints-rust-sdk) project

You will need an API Token.  Once you have this token, use it to login locally with this command.

```bash
cargo login
```

Run this command to build and publish the rust sdk crate (NOTE: It is recommended to run `cargo package` first to ensure the `.crate` file isn’t over 10MB)

```bash
cargo publish
```

Changes to packaging and publishing can be made by editing the `Cargo.toml` file




