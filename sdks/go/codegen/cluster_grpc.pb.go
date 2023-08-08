// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: cluster.proto

package codegen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClusterService_CreateCluster_FullMethodName         = "/codegen.ClusterService/CreateCluster"
	ClusterService_BuildCluster_FullMethodName          = "/codegen.ClusterService/BuildCluster"
	ClusterService_CloneCluster_FullMethodName          = "/codegen.ClusterService/CloneCluster"
	ClusterService_AddPlatformTeam_FullMethodName       = "/codegen.ClusterService/AddPlatformTeam"
	ClusterService_AddApplicationTeam_FullMethodName    = "/codegen.ClusterService/AddApplicationTeam"
	ClusterService_AddTeams_FullMethodName              = "/codegen.ClusterService/AddTeams"
	ClusterService_AddMngClusterProvider_FullMethodName = "/codegen.ClusterService/AddMngClusterProvider"
	ClusterService_AddAsgClusterProvider_FullMethodName = "/codegen.ClusterService/AddAsgClusterProvider"
	ClusterService_AddClusterProvider_FullMethodName    = "/codegen.ClusterService/AddClusterProvider"
	ClusterService_AddVpcProvider_FullMethodName        = "/codegen.ClusterService/AddVpcProvider"
	ClusterService_AddResourceProvider_FullMethodName   = "/codegen.ClusterService/AddResourceProvider"
	ClusterService_AddAckAddOn_FullMethodName           = "/codegen.ClusterService/AddAckAddOn"
	ClusterService_AddKubeProxyAddOn_FullMethodName     = "/codegen.ClusterService/AddKubeProxyAddOn"
	ClusterService_AddCoreDNSAddOn_FullMethodName       = "/codegen.ClusterService/AddCoreDNSAddOn"
	ClusterService_AddMetricsServerAddOn_FullMethodName = "/codegen.ClusterService/AddMetricsServerAddOn"
	ClusterService_AddAddons_FullMethodName             = "/codegen.ClusterService/AddAddons"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*APIResponse, error)
	BuildCluster(ctx context.Context, in *BuildClusterRequest, opts ...grpc.CallOption) (*APIResponse, error)
	CloneCluster(ctx context.Context, in *CloneClusterRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddPlatformTeam(ctx context.Context, in *AddPlatformTeamRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddApplicationTeam(ctx context.Context, in *AddApplicationTeamRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddTeams(ctx context.Context, in *AddTeamsRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddMngClusterProvider(ctx context.Context, in *AddMngClusterProviderRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddAsgClusterProvider(ctx context.Context, in *AddAsgClusterProviderRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddClusterProvider(ctx context.Context, in *AddClusterProviderRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddVpcProvider(ctx context.Context, in *AddVpcProviderRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddResourceProvider(ctx context.Context, in *AddResourceProviderRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddAckAddOn(ctx context.Context, in *AddAckAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddKubeProxyAddOn(ctx context.Context, in *AddKubeProxyAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddCoreDNSAddOn(ctx context.Context, in *AddCoreDNSAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddMetricsServerAddOn(ctx context.Context, in *AddMetricsServerAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error)
	AddAddons(ctx context.Context, in *AddAddonsRequest, opts ...grpc.CallOption) (*APIResponse, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_CreateCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) BuildCluster(ctx context.Context, in *BuildClusterRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_BuildCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) CloneCluster(ctx context.Context, in *CloneClusterRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_CloneCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddPlatformTeam(ctx context.Context, in *AddPlatformTeamRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddPlatformTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddApplicationTeam(ctx context.Context, in *AddApplicationTeamRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddApplicationTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddTeams(ctx context.Context, in *AddTeamsRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddTeams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddMngClusterProvider(ctx context.Context, in *AddMngClusterProviderRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddMngClusterProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddAsgClusterProvider(ctx context.Context, in *AddAsgClusterProviderRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddAsgClusterProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddClusterProvider(ctx context.Context, in *AddClusterProviderRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddClusterProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddVpcProvider(ctx context.Context, in *AddVpcProviderRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddVpcProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddResourceProvider(ctx context.Context, in *AddResourceProviderRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddResourceProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddAckAddOn(ctx context.Context, in *AddAckAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddAckAddOn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddKubeProxyAddOn(ctx context.Context, in *AddKubeProxyAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddKubeProxyAddOn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddCoreDNSAddOn(ctx context.Context, in *AddCoreDNSAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddCoreDNSAddOn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddMetricsServerAddOn(ctx context.Context, in *AddMetricsServerAddOnRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddMetricsServerAddOn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) AddAddons(ctx context.Context, in *AddAddonsRequest, opts ...grpc.CallOption) (*APIResponse, error) {
	out := new(APIResponse)
	err := c.cc.Invoke(ctx, ClusterService_AddAddons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations must embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	CreateCluster(context.Context, *CreateClusterRequest) (*APIResponse, error)
	BuildCluster(context.Context, *BuildClusterRequest) (*APIResponse, error)
	CloneCluster(context.Context, *CloneClusterRequest) (*APIResponse, error)
	AddPlatformTeam(context.Context, *AddPlatformTeamRequest) (*APIResponse, error)
	AddApplicationTeam(context.Context, *AddApplicationTeamRequest) (*APIResponse, error)
	AddTeams(context.Context, *AddTeamsRequest) (*APIResponse, error)
	AddMngClusterProvider(context.Context, *AddMngClusterProviderRequest) (*APIResponse, error)
	AddAsgClusterProvider(context.Context, *AddAsgClusterProviderRequest) (*APIResponse, error)
	AddClusterProvider(context.Context, *AddClusterProviderRequest) (*APIResponse, error)
	AddVpcProvider(context.Context, *AddVpcProviderRequest) (*APIResponse, error)
	AddResourceProvider(context.Context, *AddResourceProviderRequest) (*APIResponse, error)
	AddAckAddOn(context.Context, *AddAckAddOnRequest) (*APIResponse, error)
	AddKubeProxyAddOn(context.Context, *AddKubeProxyAddOnRequest) (*APIResponse, error)
	AddCoreDNSAddOn(context.Context, *AddCoreDNSAddOnRequest) (*APIResponse, error)
	AddMetricsServerAddOn(context.Context, *AddMetricsServerAddOnRequest) (*APIResponse, error)
	AddAddons(context.Context, *AddAddonsRequest) (*APIResponse, error)
	mustEmbedUnimplementedClusterServiceServer()
}

// UnimplementedClusterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) CreateCluster(context.Context, *CreateClusterRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedClusterServiceServer) BuildCluster(context.Context, *BuildClusterRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildCluster not implemented")
}
func (UnimplementedClusterServiceServer) CloneCluster(context.Context, *CloneClusterRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneCluster not implemented")
}
func (UnimplementedClusterServiceServer) AddPlatformTeam(context.Context, *AddPlatformTeamRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlatformTeam not implemented")
}
func (UnimplementedClusterServiceServer) AddApplicationTeam(context.Context, *AddApplicationTeamRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddApplicationTeam not implemented")
}
func (UnimplementedClusterServiceServer) AddTeams(context.Context, *AddTeamsRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeams not implemented")
}
func (UnimplementedClusterServiceServer) AddMngClusterProvider(context.Context, *AddMngClusterProviderRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMngClusterProvider not implemented")
}
func (UnimplementedClusterServiceServer) AddAsgClusterProvider(context.Context, *AddAsgClusterProviderRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAsgClusterProvider not implemented")
}
func (UnimplementedClusterServiceServer) AddClusterProvider(context.Context, *AddClusterProviderRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClusterProvider not implemented")
}
func (UnimplementedClusterServiceServer) AddVpcProvider(context.Context, *AddVpcProviderRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVpcProvider not implemented")
}
func (UnimplementedClusterServiceServer) AddResourceProvider(context.Context, *AddResourceProviderRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddResourceProvider not implemented")
}
func (UnimplementedClusterServiceServer) AddAckAddOn(context.Context, *AddAckAddOnRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAckAddOn not implemented")
}
func (UnimplementedClusterServiceServer) AddKubeProxyAddOn(context.Context, *AddKubeProxyAddOnRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKubeProxyAddOn not implemented")
}
func (UnimplementedClusterServiceServer) AddCoreDNSAddOn(context.Context, *AddCoreDNSAddOnRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoreDNSAddOn not implemented")
}
func (UnimplementedClusterServiceServer) AddMetricsServerAddOn(context.Context, *AddMetricsServerAddOnRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetricsServerAddOn not implemented")
}
func (UnimplementedClusterServiceServer) AddAddons(context.Context, *AddAddonsRequest) (*APIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddons not implemented")
}
func (UnimplementedClusterServiceServer) mustEmbedUnimplementedClusterServiceServer() {}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_CreateCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_BuildCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).BuildCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_BuildCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).BuildCluster(ctx, req.(*BuildClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_CloneCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CloneCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_CloneCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CloneCluster(ctx, req.(*CloneClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddPlatformTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlatformTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddPlatformTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddPlatformTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddPlatformTeam(ctx, req.(*AddPlatformTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddApplicationTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddApplicationTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddApplicationTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddApplicationTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddApplicationTeam(ctx, req.(*AddApplicationTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddTeams(ctx, req.(*AddTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddMngClusterProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMngClusterProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddMngClusterProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddMngClusterProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddMngClusterProvider(ctx, req.(*AddMngClusterProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddAsgClusterProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAsgClusterProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddAsgClusterProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddAsgClusterProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddAsgClusterProvider(ctx, req.(*AddAsgClusterProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddClusterProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClusterProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddClusterProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddClusterProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddClusterProvider(ctx, req.(*AddClusterProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddVpcProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVpcProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddVpcProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddVpcProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddVpcProvider(ctx, req.(*AddVpcProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddResourceProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddResourceProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddResourceProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddResourceProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddResourceProvider(ctx, req.(*AddResourceProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddAckAddOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAckAddOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddAckAddOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddAckAddOn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddAckAddOn(ctx, req.(*AddAckAddOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddKubeProxyAddOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKubeProxyAddOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddKubeProxyAddOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddKubeProxyAddOn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddKubeProxyAddOn(ctx, req.(*AddKubeProxyAddOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddCoreDNSAddOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoreDNSAddOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddCoreDNSAddOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddCoreDNSAddOn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddCoreDNSAddOn(ctx, req.(*AddCoreDNSAddOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddMetricsServerAddOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMetricsServerAddOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddMetricsServerAddOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddMetricsServerAddOn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddMetricsServerAddOn(ctx, req.(*AddMetricsServerAddOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_AddAddons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAddonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).AddAddons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_AddAddons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).AddAddons(ctx, req.(*AddAddonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codegen.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _ClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "BuildCluster",
			Handler:    _ClusterService_BuildCluster_Handler,
		},
		{
			MethodName: "CloneCluster",
			Handler:    _ClusterService_CloneCluster_Handler,
		},
		{
			MethodName: "AddPlatformTeam",
			Handler:    _ClusterService_AddPlatformTeam_Handler,
		},
		{
			MethodName: "AddApplicationTeam",
			Handler:    _ClusterService_AddApplicationTeam_Handler,
		},
		{
			MethodName: "AddTeams",
			Handler:    _ClusterService_AddTeams_Handler,
		},
		{
			MethodName: "AddMngClusterProvider",
			Handler:    _ClusterService_AddMngClusterProvider_Handler,
		},
		{
			MethodName: "AddAsgClusterProvider",
			Handler:    _ClusterService_AddAsgClusterProvider_Handler,
		},
		{
			MethodName: "AddClusterProvider",
			Handler:    _ClusterService_AddClusterProvider_Handler,
		},
		{
			MethodName: "AddVpcProvider",
			Handler:    _ClusterService_AddVpcProvider_Handler,
		},
		{
			MethodName: "AddResourceProvider",
			Handler:    _ClusterService_AddResourceProvider_Handler,
		},
		{
			MethodName: "AddAckAddOn",
			Handler:    _ClusterService_AddAckAddOn_Handler,
		},
		{
			MethodName: "AddKubeProxyAddOn",
			Handler:    _ClusterService_AddKubeProxyAddOn_Handler,
		},
		{
			MethodName: "AddCoreDNSAddOn",
			Handler:    _ClusterService_AddCoreDNSAddOn_Handler,
		},
		{
			MethodName: "AddMetricsServerAddOn",
			Handler:    _ClusterService_AddMetricsServerAddOn_Handler,
		},
		{
			MethodName: "AddAddons",
			Handler:    _ClusterService_AddAddons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}
