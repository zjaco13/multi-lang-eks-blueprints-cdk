import { Server, ServerUnaryCall, sendUnaryData, ServerCredentials, handleUnaryCall } from "@grpc/grpc-js";
import * as lib from '../lib';
import * as blueprints from "@aws-quickstart/eks-blueprints";
import * as cdk from "aws-cdk-lib";

const server = new Server();
const app = new cdk.App();
const builder = blueprints.EksBlueprint.builder();

class ClusterServer implements lib.proto.ClusterServiceServer {
    [name: string]: import("@grpc/grpc-js").UntypedHandleCall;
    createCluster: handleUnaryCall<lib.proto.CreateClusterRequest, lib.proto.APIResponse> = (call, callback) => {
        const response = lib.proto.APIResponse.create();
        builder.id(call.request.id);
        const name = call.request.name ?? call.request.id;
        builder.name(name);

        response.message = `Cluster Created: ${name}`;
        callback(null, response);
    };
    addTeams: handleUnaryCall<lib.proto.AddTeamsRequest, lib.proto.APIResponse> = (call, callback) => {
        const response = lib.proto.APIResponse.create();
        const teams: Array<blueprints.Team> = [];
        call.request.teams.forEach(team => {
            if(team.applicationTeam){
                teams.push(new blueprints.ApplicationTeam({name: team.applicationTeam.name}));
            }
            if(team.platformTeam){
                teams.push(new blueprints.PlatformTeam({name: team.platformTeam.name}));
            }
            if(team.genericTeam) {

            }
        });

        response.message = 'Added teams to cluster';
        builder.teams(...teams);
        callback(null, response);
    };
    addClusterProvider: handleUnaryCall<lib.proto.AddClusterProviderRequest, lib.proto.APIResponse> = (call, callback) => {
        throw new Error("unimplemented");
    };
    addResourceProvider: handleUnaryCall<lib.proto.AddResourceProviderRequest, lib.proto.APIResponse> = (call, callback) => {
        throw new Error("unimplemented")
    };
    buildCluster: handleUnaryCall<lib.proto.BuildClusterRequest, lib.proto.APIResponse> = (call, callback) => {
        const response = lib.proto.APIResponse.create();
        const name = call.request.clusterName;
        const account = call.request.account ?? process.env.CDK_DEFAULT_ACCOUNT!;
        const region =  call.request.region ?? process.env.CDK_DEFAULT_REGION!;

        builder
            .account(account)
            .region(region)
            .build(app, builder.props.id!);
        app.synth();

        response.message = `Cluster ${name} Built: ${account}:${region}`;

        callback(null, response);

    }

}

server.addService(lib.proto.ClusterServiceService, new ClusterServer());
server.bindAsync('0.0.0.0:50051', ServerCredentials.createInsecure(), () => {
    server.start()
    console.log('server is running on 0.0.0.0:50051');
});
