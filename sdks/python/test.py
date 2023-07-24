import builder
from proto import cluster_pb2_grpc
from proto import cluster_pb2

def build(stub):
    stub.CreateCluster(cluster_pb2.CreateClusterRequest(id="test-from-python", name="testy"))
    stub.BuildCluster(cluster_pb2.BuildClusterRequest(clusterName="testy"))

if __name__ == "__main__":
   builder.run(build) 
