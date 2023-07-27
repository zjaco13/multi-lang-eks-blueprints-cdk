from src import builder
from src.codegen import addons_pb2
from src.codegen import teams_pb2 
from src.codegen import cluster_pb2 
from src.codegen import cluster_providers_pb2 

def build(stub):
    print(stub.CreateCluster(cluster_pb2.CreateClusterRequest(id="test-from-python")))

    print(stub.AddApplicationTeam(teams_pb2.AddApplicationTeamRequest(cluster_name="test-from-python", props=teams_pb2.TeamProps(name="app-team"))))

    print(stub.AddMngClusterProvider(cluster_providers_pb2.AddMngClusterProviderRequest(cluster_name="test-from-python", mng_cluster_provider=cluster_providers_pb2.MngClusterProvider(name="provider", version="1.27"))))

    print(stub.AddKubeProxyAddOn(addons_pb2.AddKubeProxyAddOnRequest(cluster_name="test-from-python", kube_proxy_add_on=addons_pb2.KubeProxyAddOn())))

    print(stub.BuildCluster(cluster_pb2.BuildClusterRequest(cluster_name="test-from-python")))

if __name__ == "__main__":
   builder.run(build) 
