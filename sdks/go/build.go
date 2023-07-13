package builder

import (
	"context"
	"log"
	"time"

	pb "github.com/zjaco13/proto-test/sdks/go-sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunBuild(build func(pb.ClusterServiceClient, context.Context)) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := pb.NewClusterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	build(client, ctx)
}
