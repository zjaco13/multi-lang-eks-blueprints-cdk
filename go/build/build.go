package build

import (
	"context"
	"fmt"
	"log"

	pb "github.com/zjaco13/proto-test/go/codegen"
)

func Build(client pb.ClusterServiceClient, ctx context.Context) {
	resp, err := client.CreateCluster(ctx, &pb.CreateClusterRequest{Id: "test-cluster"})
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println(resp)
	resp, err = client.BuildCluster(ctx, &pb.BuildClusterRequest{ClusterName: "test-cluster"})
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println(resp)
}
