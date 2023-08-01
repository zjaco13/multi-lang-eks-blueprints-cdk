import grpc
from eks_blueprints_python_sdk.codegen import cluster_pb2_grpc

def run(build_func):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = cluster_pb2_grpc.ClusterServiceStub(channel)
        build_func(stub)
