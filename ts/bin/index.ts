import { Server, ServerUnaryCall, sendUnaryData, ServerCredentials } from "@grpc/grpc-js";
import { proto_test } from "../lib/cluster";
import * as blueprints from "@aws-quickstart/eks-blueprints"
import * as cdk from "aws-cdk-lib"
import UnimplementedClusterServiceService = proto_test.UnimplementedClusterServiceService

const server = new Server();

class ClusterServer extends UnimplementedClusterServiceService {
    SendCluster(call: ServerUnaryCall<proto_test.EksBlueprint, proto_test.EksBlueprintResponse>, callback: sendUnaryData<proto_test.EksBlueprintResponse>): void {
        const response = new proto_test.EksBlueprintResponse();
        const app = new cdk.App();
        const blueprint = blueprints.EksBlueprint.builder()
            .account(process.env.CDK_DEFAULT_ACCOUNT!)
            .region(process.env.CDK_DEFAULT_REGION!)
            .build(app, call.request.id);
        app.synth()
        console.log(blueprint)
        response.resp = `Created stack with id ${call.request.id}`
        callback(null, response)

    }
}

server.addService(UnimplementedClusterServiceService.definition, new ClusterServer())
server.bindAsync('0.0.0.0:50051', ServerCredentials.createInsecure(), () => {
    server.start()
    console.log('server is running on 0.0.0.0:50051');
})


