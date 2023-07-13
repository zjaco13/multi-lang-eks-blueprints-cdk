package main

import (
	"context"
	"fmt"

	builder "github.com/zjaco13/proto-test/go"
	pb "github.com/zjaco13/proto-test/go/codegen"
)

func main() {
	builder.RunBuild(build)
}

func build(client pb.ClusterServiceClient, ctx context.Context) {
	fmt.Println(client.CreateCluster(ctx, &pb.CreateClusterRequest{Id: "testy"}))
	fmt.Println(client.BuildCluster(ctx, &pb.BuildClusterRequest{ClusterName: "testy"}))

}
