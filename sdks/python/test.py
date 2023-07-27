from src import builder
from src.codegen import addons_pb2
from src.codegen import teams_pb2 
from src.codegen import cluster_pb2 
from src.codegen import cluster_providers_pb2 

def build(stub):
    stub.CreateCluster(cluster_pb2.CreateClusterRequest(id="test-from-python", name="testy"))
    stub.AddApplicationTeam(teams_pb2.AddApplicationTeamRequest("testy", teams_pb2.TeamProps("app-team")))
    stub.AddMngClusterProvider(cluster_providers_pb2.AddMngClusterProviderRequest("testy", cluster_providers_pb2.MngClusterProvider("provider", "1.27")))
    stub.AddKubeProxyAddOn(addons_pb2.AddKubeProxyAddOnRequest("testy", addons_pb2.KubeProxyAddOn()))
    stub.BuildCluster(cluster_pb2.BuildClusterRequest(cluster_name="testy"))

if __name__ == "__main__":
   builder.run(build) 
