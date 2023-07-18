import { Server, ServerCredentials, handleUnaryCall } from "@grpc/grpc-js";
import * as lib from '../lib';
import * as blueprints from "@aws-quickstart/eks-blueprints";
import * as cdk from "aws-cdk-lib";
import * as eks from 'aws-cdk-lib/aws-eks';

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
    addAddons: handleUnaryCall<lib.proto.AddAddonsRequest, lib.proto.APIResponse> = (call, callback) => {
        const response = lib.proto.APIResponse.create();
        const addons: Array<blueprints.ClusterAddOn> = [];
        call.request.addons.forEach(addon => {
            if(addon.ackAddOn) {
                addons.push(new blueprints.addons.AckAddOn({
                    ...addon.ackAddOn,
                    serviceName: addon.ackAddOn.serviceName as blueprints.AckServiceName
                }));
            }

        });
        callback(null, response);
    }
    addTeams: handleUnaryCall<lib.proto.AddTeamsRequest, lib.proto.APIResponse> = (call, callback) => {
        const response = lib.proto.APIResponse.create();
        const teams: Array<blueprints.Team> = [];
        call.request.teams.forEach(team => {
            if(team.applicationTeam){
                teams.push(new blueprints.ApplicationTeam(team.applicationTeam));
            }
            if(team.platformTeam){
                teams.push(new blueprints.PlatformTeam(team.platformTeam));
            }
            if(team.genericTeam) {

            }
        });

        response.message = 'Added teams to cluster: ' + teams.map(team => team.name).join(" ");
        builder.teams(...teams);
        callback(null, response);
    };
    addClusterProvider: handleUnaryCall<lib.proto.AddClusterProviderRequest, lib.proto.APIResponse> = (call, callback) => {
        const response = lib.proto.APIResponse.create();
        const reqClusterProvider = call.request.clusterProvider;
        let clusterProvider: blueprints.ClusterProvider | undefined;
        let type: string = "";
        const version = reqClusterProvider?.asgClusterProvider?.version ?? reqClusterProvider?.mngClusterProvider?.version ?? "undefined"
        if(reqClusterProvider?.asgClusterProvider) {
            clusterProvider = new blueprints.AsgClusterProvider(
                {
                    ...reqClusterProvider.asgClusterProvider,
                    version: eks.KubernetesVersion.of(version),
                }
            );
            type = "asg";
        }
        if(reqClusterProvider?.mngClusterProvider) {
            clusterProvider = new blueprints.MngClusterProvider(
                {...reqClusterProvider.mngClusterProvider, 
                    version: eks.KubernetesVersion.of(version),
                },
            );
            type = "mng";
        }
        builder.clusterProvider(clusterProvider!);
        response.message = `Added ClusterProvider to cluster: ${type}`;
        callback(null, response);
    };
    addResourceProvider: handleUnaryCall<lib.proto.AddResourceProviderRequest, lib.proto.APIResponse> = (call, callback) => {
        const response = lib.proto.APIResponse.create();
        const reqResourceProvider = call.request.resourceProvider; 
        let resourceProvider: blueprints.ResourceProvider | undefined;
        let name = call.request.name
        if(reqResourceProvider?.vpcProvider) {
            resourceProvider = new blueprints.VpcProvider(reqResourceProvider.vpcProvider.vpcId);
        }

        builder.resourceProvider(name, resourceProvider!)
        response.message = `Added Resource provider to cluster: ${name}`
        callback(null, response)
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
