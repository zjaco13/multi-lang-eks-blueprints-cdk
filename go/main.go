package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/zjaco13/proto-test/go/build"
	"github.com/zjaco13/proto-test/go/codegen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddr = flag.String("addr", "localhost:50051", "Server address in format host:port")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := codegen.NewClusterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SendCluster(ctx, build.Build())
	log.Println(resp)
}
