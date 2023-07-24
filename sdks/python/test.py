import builder
from proto import cluster_pb2_grpc
from proto import cluster_pb2
from proto import team_pb2
from proto import cluster_provider_pb2
from proto import addons_pb2

def build(stub):
    stub.CreateCluster(cluster_pb2.CreateClusterRequest(id="test-from-python", name="testy"))
    stub.AddApplicationTeam(team_pb2.AddApplicationTeamRequest("testy", team_pb2.TeamProps("app-team")))
    stub.AddMngClusterProvider(cluster_provider_pb2.AddMngClusterProviderRequest("testy", cluster_provider_pb2.MngClusterProvider("provider", "1.27")))
    stub.AddKubeProxyAddOn(addons_pb2.AddKubeProxyAddOnRequest("testy", addons_pb2.KubeProxyAddOn()))
    stub.BuildCluster(cluster_pb2.BuildClusterRequest(clusterName="testy"))

if __name__ == "__main__":
   builder.run(build) 
