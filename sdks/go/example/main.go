package main

import (
	"context"
	"fmt"

	builder "github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk"
	pb "github.com/zjaco13/multi-lang-eks-blueprints-cdk/sdks/go-sdk/proto"
)

func main() {
	builder.RunBuild(build)
}

func build(client pb.ClusterServiceClient, ctx context.Context) {
	fmt.Println(client.CreateCluster(ctx, &pb.CreateClusterRequest{Id: "test-from-go"}))

	teams := []*pb.Team{}
	teams = append(teams, &pb.Team{Team: &pb.Team_ApplicationTeam{ApplicationTeam: &pb.ApplicationTeam{Name: "team1"}}})
	fmt.Println(client.AddTeams(ctx, &pb.AddTeamsRequest{
		Teams: teams,
	}))

	addons := []*pb.Addon{}
	addons = append(addons, &pb.Addon{Addon: &pb.Addon_KubeProxyAddOn{KubeProxyAddOn: &pb.KubeProxyAddOn{}}})
	fmt.Println(client.AddAddons(ctx, &pb.AddAddonsRequest{
		Addons: addons,
	}))

	name := "test-cluster-provider"
	fmt.Println(client.AddClusterProvider(ctx, &pb.AddClusterProviderRequest{ClusterProvider: &pb.ClusterProvider{ClusterProvider: &pb.ClusterProvider_MngClusterProvider{
		MngClusterProvider: &pb.MngClusterProvider{Name: &name, Version: "1.27"},
	}}}))

	id := "vpc"
	fmt.Println(client.AddResourceProvider(ctx, &pb.AddResourceProviderRequest{
		Name: id,
		ResourceProvider: &pb.ResourceProvider{ResourceProvider: &pb.ResourceProvider_VpcProvider{
			VpcProvider: &pb.VpcProvider{},
		},
		}}))

	fmt.Println(client.BuildCluster(ctx, &pb.BuildClusterRequest{ClusterName: "test-from-go"}))
}
