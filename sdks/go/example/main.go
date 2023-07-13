package main

import (
	"context"
	"fmt"

	builder "github.com/zjaco13/proto-test/sdks/go-sdk"
	pb "github.com/zjaco13/proto-test/sdks/go-sdk/proto"
)

func main() {
	builder.RunBuild(build)
}

func build(client pb.ClusterServiceClient, ctx context.Context) {
	fmt.Println(client.CreateCluster(ctx, &pb.CreateClusterRequest{Id: "testy"}))
	teams := []*pb.Team{}
	teams = append(teams, &pb.Team{Team: &pb.Team_ApplicationTeam{ApplicationTeam: &pb.ApplicationTeam{Name: "team1"}}})
	fmt.Println(client.AddTeams(ctx, &pb.AddTeamsRequest{
		Teams: teams,
	}))
	fmt.Println(client.BuildCluster(ctx, &pb.BuildClusterRequest{ClusterName: "testy"}))
}
