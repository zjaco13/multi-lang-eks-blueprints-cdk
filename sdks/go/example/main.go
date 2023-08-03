package main

import (
	"context"
	"fmt"

	builder "github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk"
	pb "github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk/codegen"
)

func main() {
	builder.RunBuild(build)
}

func build(client pb.ClusterServiceClient, ctx context.Context) {
	fmt.Println(client.CreateCluster(ctx, &pb.CreateClusterRequest{Id: "test-from-go"}))

	fmt.Println(client.AddApplicationTeam(ctx, &pb.AddApplicationTeamRequest{
		ClusterName: "test-from-go",
		Props:       &pb.TeamProps{Name: "app-team"},
	}))

	fmt.Println(client.AddKubeProxyAddOn(ctx, &pb.AddKubeProxyAddOnRequest{
		ClusterName:    "test-from-go",
		KubeProxyAddOn: &pb.KubeProxyAddOn{},
	}))

	name := "test-cluster-provider"
	fmt.Println(client.AddMngClusterProvider(ctx, &pb.AddMngClusterProviderRequest{
		ClusterName:        "test-from-go",
		MngClusterProvider: &pb.MngClusterProvider{Name: &name, Version: "1.27"},
	}))

	fmt.Println(client.AddVpcProvider(ctx, &pb.AddVpcProviderRequest{
		ClusterName: "test-from-go",
		Name:        "vpc",
		VpcProvider: &pb.VpcProvider{},
	}))

	fmt.Println(client.BuildCluster(ctx, &pb.BuildClusterRequest{ClusterName: "test-from-go"}))
}
