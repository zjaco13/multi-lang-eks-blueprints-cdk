import { Server, ServerUnaryCall, sendUnaryData, ServerCredentials } from "@grpc/grpc-js";
import { proto_test } from "../lib/cluster";
import * as blueprints from "@aws-quickstart/eks-blueprints"
import * as cdk from "aws-cdk-lib"
import UnimplementedClusterServiceService = proto_test.UnimplementedClusterServiceService

const server = new Server();
const app = new cdk.App();
const builder = blueprints.EksBlueprint.builder();

class ClusterServer extends UnimplementedClusterServiceService {
    CreateCluster(call: ServerUnaryCall<proto_test.CreateClusterRequest, proto_test.APIResponse>, callback: sendUnaryData<proto_test.APIResponse>): void {
        const response = new proto_test.APIResponse;
        builder.id(call.request.id);

        const name = call.request.has_name ? call.request.name : call.request.id;
        builder.name(name);
        response.message = `Cluster Created: ${name}`;
        callback(null, response);
    }
    AddTeams(call: ServerUnaryCall<proto_test.AddTeamsRequest, proto_test.APIResponse>, callback: sendUnaryData<proto_test.APIResponse>): void {
        throw new Error("Method not implemented.");
    }
    AddClusterProvider(call: ServerUnaryCall<proto_test.AddClusterProviderRequest, proto_test.APIResponse>, callback: sendUnaryData<proto_test.APIResponse>): void {
        throw new Error("Method not implemented.");
    }
    AddResourceProvider(call: ServerUnaryCall<proto_test.AddResourceProviderRequest, proto_test.APIResponse>, callback: sendUnaryData<proto_test.APIResponse>): void {
        throw new Error("Method not implemented.");
    }
    BuildCluster(call: ServerUnaryCall<proto_test.BuildClusterRequest, proto_test.APIResponse>, callback: sendUnaryData<proto_test.APIResponse>): void {
        const response = new proto_test.APIResponse;
        const name = call.request.clusterName;
        const account = call.request.has_account ? call.request.account : process.env.CDK_DEFAULT_ACCOUNT!;
        const region = call.request.has_region ? call.request.region : process.env.CDK_DEFAULT_REGION!;

        builder
            .account(account)
            .region(region)
            .build(app, builder.props.id!);
        app.synth()

        response.message = `Cluster ${name} Built: ${account}:${region}`;

        callback(null, response);
    }

}

server.addService(UnimplementedClusterServiceService.definition, new ClusterServer())
server.bindAsync('0.0.0.0:50051', ServerCredentials.createInsecure(), () => {
    server.start()
    console.log('server is running on 0.0.0.0:50051');
})


