import builder
from proto import cluster_pb2_grpc
from proto import cluster_pb2

def build(stub):
    stub.CreateCluster(cluster_pb2.CreateClusterRequest("test-from-python"))

if __name__ == "__main__":
   builder.run(build()) 
