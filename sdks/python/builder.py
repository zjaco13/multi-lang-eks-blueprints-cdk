import grpc
import proto 

def run(build_func):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = proto.ClusterServiceStub(channel)
        build_func(stub)
