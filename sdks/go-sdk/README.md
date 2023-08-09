# Go Client for Multi Language EKS Blueprints for cdk

## Run the example

To run the example, first start the server
```bash
git clone https://github.com/aws-quickstart/cdk-eks-blueprints.git
cd cdk-eks-blueprints
git checkout blueprints-api-support
make run-server
```

Then run the go example by cloning this repository and running with make, go is required
```bash
git clone https://github.com/zjaco13/multi-lang-eks-blueprints-cdk.git
cd multi-lang-eks-blueprints-cdk
make go-example
```

## How to use as a package

Create a new go project
```bash
mkdir example
cd example
go mod init example.com
```

Add this sdk as a library
```bash
go get github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk
```

Add this to a `main.go` file
```go
package main

import (
	builder "github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk"
	pb "github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk/codegen"
)

func main() {
	builder.RunBuild(build)
}

func build(client pb.ClusterServiceClient, ctx context.Context) {
    ...
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
go run main.go
```

From the server, deploy your EKS Blueprint
```bash
make deploy-server
```

## Contributing

#### Publishing

A new version should be released whenever the protobufs are changed

To release a new version for the go sdk, push a new tagged commit to the origin repository e.g

```bash
git tag v0.1.2
git push origin v0.1.2
```

Make the module available on the proxy server using `go list`

```bash
GOPROXY=proxy.golang.org go list -m github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk@v0.1.2
```
