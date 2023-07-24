# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import addons_pb2 as addons__pb2
from . import cluster_pb2 as cluster__pb2
from . import cluster_provider_pb2 as cluster__provider__pb2
from . import resource_provider_pb2 as resource__provider__pb2
from . import team_pb2 as team__pb2


class ClusterServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateCluster = channel.unary_unary(
                '/codegen.ClusterService/CreateCluster',
                request_serializer=cluster__pb2.CreateClusterRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.BuildCluster = channel.unary_unary(
                '/codegen.ClusterService/BuildCluster',
                request_serializer=cluster__pb2.BuildClusterRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.CloneCluster = channel.unary_unary(
                '/codegen.ClusterService/CloneCluster',
                request_serializer=cluster__pb2.CloneClusterRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddPlatformTeam = channel.unary_unary(
                '/codegen.ClusterService/AddPlatformTeam',
                request_serializer=team__pb2.AddPlatformTeamRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddApplicationTeam = channel.unary_unary(
                '/codegen.ClusterService/AddApplicationTeam',
                request_serializer=team__pb2.AddApplicationTeamRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddTeams = channel.unary_unary(
                '/codegen.ClusterService/AddTeams',
                request_serializer=team__pb2.AddTeamsRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddMngClusterProvider = channel.unary_unary(
                '/codegen.ClusterService/AddMngClusterProvider',
                request_serializer=cluster__provider__pb2.AddMngClusterProviderRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddAsgClusterProvider = channel.unary_unary(
                '/codegen.ClusterService/AddAsgClusterProvider',
                request_serializer=cluster__provider__pb2.AddAsgClusterProviderRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddClusterProvider = channel.unary_unary(
                '/codegen.ClusterService/AddClusterProvider',
                request_serializer=cluster__provider__pb2.AddClusterProviderRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddVpcProvider = channel.unary_unary(
                '/codegen.ClusterService/AddVpcProvider',
                request_serializer=resource__provider__pb2.AddVpcProviderRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddResourceProvider = channel.unary_unary(
                '/codegen.ClusterService/AddResourceProvider',
                request_serializer=resource__provider__pb2.AddResourceProviderRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddAckAddOn = channel.unary_unary(
                '/codegen.ClusterService/AddAckAddOn',
                request_serializer=addons__pb2.AddAckAddOnRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddKubeProxyAddOn = channel.unary_unary(
                '/codegen.ClusterService/AddKubeProxyAddOn',
                request_serializer=addons__pb2.AddKubeProxyAddOnRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )
        self.AddAddons = channel.unary_unary(
                '/codegen.ClusterService/AddAddons',
                request_serializer=addons__pb2.AddAddonsRequest.SerializeToString,
                response_deserializer=cluster__pb2.APIResponse.FromString,
                )


class ClusterServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateCluster(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BuildCluster(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CloneCluster(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddPlatformTeam(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddApplicationTeam(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddTeams(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddMngClusterProvider(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAsgClusterProvider(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddClusterProvider(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddVpcProvider(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddResourceProvider(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAckAddOn(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddKubeProxyAddOn(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddAddons(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ClusterServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateCluster': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateCluster,
                    request_deserializer=cluster__pb2.CreateClusterRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'BuildCluster': grpc.unary_unary_rpc_method_handler(
                    servicer.BuildCluster,
                    request_deserializer=cluster__pb2.BuildClusterRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'CloneCluster': grpc.unary_unary_rpc_method_handler(
                    servicer.CloneCluster,
                    request_deserializer=cluster__pb2.CloneClusterRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddPlatformTeam': grpc.unary_unary_rpc_method_handler(
                    servicer.AddPlatformTeam,
                    request_deserializer=team__pb2.AddPlatformTeamRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddApplicationTeam': grpc.unary_unary_rpc_method_handler(
                    servicer.AddApplicationTeam,
                    request_deserializer=team__pb2.AddApplicationTeamRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddTeams': grpc.unary_unary_rpc_method_handler(
                    servicer.AddTeams,
                    request_deserializer=team__pb2.AddTeamsRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddMngClusterProvider': grpc.unary_unary_rpc_method_handler(
                    servicer.AddMngClusterProvider,
                    request_deserializer=cluster__provider__pb2.AddMngClusterProviderRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddAsgClusterProvider': grpc.unary_unary_rpc_method_handler(
                    servicer.AddAsgClusterProvider,
                    request_deserializer=cluster__provider__pb2.AddAsgClusterProviderRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddClusterProvider': grpc.unary_unary_rpc_method_handler(
                    servicer.AddClusterProvider,
                    request_deserializer=cluster__provider__pb2.AddClusterProviderRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddVpcProvider': grpc.unary_unary_rpc_method_handler(
                    servicer.AddVpcProvider,
                    request_deserializer=resource__provider__pb2.AddVpcProviderRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddResourceProvider': grpc.unary_unary_rpc_method_handler(
                    servicer.AddResourceProvider,
                    request_deserializer=resource__provider__pb2.AddResourceProviderRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddAckAddOn': grpc.unary_unary_rpc_method_handler(
                    servicer.AddAckAddOn,
                    request_deserializer=addons__pb2.AddAckAddOnRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddKubeProxyAddOn': grpc.unary_unary_rpc_method_handler(
                    servicer.AddKubeProxyAddOn,
                    request_deserializer=addons__pb2.AddKubeProxyAddOnRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
            'AddAddons': grpc.unary_unary_rpc_method_handler(
                    servicer.AddAddons,
                    request_deserializer=addons__pb2.AddAddonsRequest.FromString,
                    response_serializer=cluster__pb2.APIResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'codegen.ClusterService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ClusterService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateCluster(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/CreateCluster',
            cluster__pb2.CreateClusterRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BuildCluster(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/BuildCluster',
            cluster__pb2.BuildClusterRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CloneCluster(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/CloneCluster',
            cluster__pb2.CloneClusterRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddPlatformTeam(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddPlatformTeam',
            team__pb2.AddPlatformTeamRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddApplicationTeam(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddApplicationTeam',
            team__pb2.AddApplicationTeamRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddTeams(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddTeams',
            team__pb2.AddTeamsRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddMngClusterProvider(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddMngClusterProvider',
            cluster__provider__pb2.AddMngClusterProviderRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddAsgClusterProvider(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddAsgClusterProvider',
            cluster__provider__pb2.AddAsgClusterProviderRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddClusterProvider(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddClusterProvider',
            cluster__provider__pb2.AddClusterProviderRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddVpcProvider(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddVpcProvider',
            resource__provider__pb2.AddVpcProviderRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddResourceProvider(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddResourceProvider',
            resource__provider__pb2.AddResourceProviderRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddAckAddOn(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddAckAddOn',
            addons__pb2.AddAckAddOnRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddKubeProxyAddOn(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddKubeProxyAddOn',
            addons__pb2.AddKubeProxyAddOnRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddAddons(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/codegen.ClusterService/AddAddons',
            addons__pb2.AddAddonsRequest.SerializeToString,
            cluster__pb2.APIResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
