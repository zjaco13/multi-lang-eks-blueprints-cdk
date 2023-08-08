// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: addons.proto

package codegen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddAddonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string   `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Addons      []*Addon `protobuf:"bytes,2,rep,name=addons,proto3" json:"addons,omitempty"`
}

func (x *AddAddonsRequest) Reset() {
	*x = AddAddonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAddonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAddonsRequest) ProtoMessage() {}

func (x *AddAddonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAddonsRequest.ProtoReflect.Descriptor instead.
func (*AddAddonsRequest) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{0}
}

func (x *AddAddonsRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *AddAddonsRequest) GetAddons() []*Addon {
	if x != nil {
		return x.Addons
	}
	return nil
}

type Addon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Addon:
	//
	//	*Addon_AckAddOn
	//	*Addon_KubeProxyAddOn
	Addon isAddon_Addon `protobuf_oneof:"addon"`
}

func (x *Addon) Reset() {
	*x = Addon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addon) ProtoMessage() {}

func (x *Addon) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addon.ProtoReflect.Descriptor instead.
func (*Addon) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{1}
}

func (m *Addon) GetAddon() isAddon_Addon {
	if m != nil {
		return m.Addon
	}
	return nil
}

func (x *Addon) GetAckAddOn() *AckAddOn {
	if x, ok := x.GetAddon().(*Addon_AckAddOn); ok {
		return x.AckAddOn
	}
	return nil
}

func (x *Addon) GetKubeProxyAddOn() *KubeProxyAddOn {
	if x, ok := x.GetAddon().(*Addon_KubeProxyAddOn); ok {
		return x.KubeProxyAddOn
	}
	return nil
}

type isAddon_Addon interface {
	isAddon_Addon()
}

type Addon_AckAddOn struct {
	AckAddOn *AckAddOn `protobuf:"bytes,1,opt,name=ack_add_on,json=ackAddOn,proto3,oneof"`
}

type Addon_KubeProxyAddOn struct {
	KubeProxyAddOn *KubeProxyAddOn `protobuf:"bytes,2,opt,name=kube_proxy_add_on,json=kubeProxyAddOn,proto3,oneof"`
}

func (*Addon_AckAddOn) isAddon_Addon() {}

func (*Addon_KubeProxyAddOn) isAddon_Addon() {}

type AddAckAddOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string    `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	AckAddOn    *AckAddOn `protobuf:"bytes,2,opt,name=ack_add_on,json=ackAddOn,proto3" json:"ack_add_on,omitempty"`
}

func (x *AddAckAddOnRequest) Reset() {
	*x = AddAckAddOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAckAddOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAckAddOnRequest) ProtoMessage() {}

func (x *AddAckAddOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAckAddOnRequest.ProtoReflect.Descriptor instead.
func (*AddAckAddOnRequest) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{2}
}

func (x *AddAckAddOnRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *AddAckAddOnRequest) GetAckAddOn() *AckAddOn {
	if x != nil {
		return x.AckAddOn
	}
	return nil
}

type AckAddOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	ServiceName string  `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
}

func (x *AckAddOn) Reset() {
	*x = AckAddOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckAddOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckAddOn) ProtoMessage() {}

func (x *AckAddOn) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckAddOn.ProtoReflect.Descriptor instead.
func (*AckAddOn) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{3}
}

func (x *AckAddOn) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *AckAddOn) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type AddKubeProxyAddOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName    string          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	KubeProxyAddOn *KubeProxyAddOn `protobuf:"bytes,2,opt,name=kube_proxy_add_on,json=kubeProxyAddOn,proto3" json:"kube_proxy_add_on,omitempty"`
}

func (x *AddKubeProxyAddOnRequest) Reset() {
	*x = AddKubeProxyAddOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKubeProxyAddOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKubeProxyAddOnRequest) ProtoMessage() {}

func (x *AddKubeProxyAddOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKubeProxyAddOnRequest.ProtoReflect.Descriptor instead.
func (*AddKubeProxyAddOnRequest) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{4}
}

func (x *AddKubeProxyAddOnRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *AddKubeProxyAddOnRequest) GetKubeProxyAddOn() *KubeProxyAddOn {
	if x != nil {
		return x.KubeProxyAddOn
	}
	return nil
}

type KubeProxyAddOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *string `protobuf:"bytes,1,opt,name=version,proto3,oneof" json:"version,omitempty"`
}

func (x *KubeProxyAddOn) Reset() {
	*x = KubeProxyAddOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeProxyAddOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeProxyAddOn) ProtoMessage() {}

func (x *KubeProxyAddOn) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeProxyAddOn.ProtoReflect.Descriptor instead.
func (*KubeProxyAddOn) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{5}
}

func (x *KubeProxyAddOn) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type AddCoreDNSAddOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName  string        `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	CoreDnsAddOn *CoreDNSAddOn `protobuf:"bytes,2,opt,name=core_dns_add_on,json=coreDnsAddOn,proto3" json:"core_dns_add_on,omitempty"`
}

func (x *AddCoreDNSAddOnRequest) Reset() {
	*x = AddCoreDNSAddOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCoreDNSAddOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCoreDNSAddOnRequest) ProtoMessage() {}

func (x *AddCoreDNSAddOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCoreDNSAddOnRequest.ProtoReflect.Descriptor instead.
func (*AddCoreDNSAddOnRequest) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{6}
}

func (x *AddCoreDNSAddOnRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *AddCoreDNSAddOnRequest) GetCoreDnsAddOn() *CoreDNSAddOn {
	if x != nil {
		return x.CoreDnsAddOn
	}
	return nil
}

type CoreDNSAddOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *string `protobuf:"bytes,1,opt,name=version,proto3,oneof" json:"version,omitempty"`
}

func (x *CoreDNSAddOn) Reset() {
	*x = CoreDNSAddOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreDNSAddOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreDNSAddOn) ProtoMessage() {}

func (x *CoreDNSAddOn) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreDNSAddOn.ProtoReflect.Descriptor instead.
func (*CoreDNSAddOn) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{7}
}

func (x *CoreDNSAddOn) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type MetricsServerAddOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateNamespace *bool `protobuf:"varint,1,opt,name=createNamespace,proto3,oneof" json:"createNamespace,omitempty"`
}

func (x *MetricsServerAddOn) Reset() {
	*x = MetricsServerAddOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsServerAddOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsServerAddOn) ProtoMessage() {}

func (x *MetricsServerAddOn) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsServerAddOn.ProtoReflect.Descriptor instead.
func (*MetricsServerAddOn) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{8}
}

func (x *MetricsServerAddOn) GetCreateNamespace() bool {
	if x != nil && x.CreateNamespace != nil {
		return *x.CreateNamespace
	}
	return false
}

type AddMetricsServerAddOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName        string              `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	MetricsServerAddOn *MetricsServerAddOn `protobuf:"bytes,2,opt,name=metrics_server_add_on,json=metricsServerAddOn,proto3" json:"metrics_server_add_on,omitempty"`
}

func (x *AddMetricsServerAddOnRequest) Reset() {
	*x = AddMetricsServerAddOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_addons_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetricsServerAddOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetricsServerAddOnRequest) ProtoMessage() {}

func (x *AddMetricsServerAddOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_addons_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetricsServerAddOnRequest.ProtoReflect.Descriptor instead.
func (*AddMetricsServerAddOnRequest) Descriptor() ([]byte, []int) {
	return file_addons_proto_rawDescGZIP(), []int{9}
}

func (x *AddMetricsServerAddOnRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *AddMetricsServerAddOnRequest) GetMetricsServerAddOn() *MetricsServerAddOn {
	if x != nil {
		return x.MetricsServerAddOn
	}
	return nil
}

var File_addons_proto protoreflect.FileDescriptor

var file_addons_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x22, 0x5d, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x41, 0x64,
	0x64, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x6f, 0x6e,
	0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x41,
	0x63, 0x6b, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x61, 0x63, 0x6b, 0x41, 0x64,
	0x64, 0x4f, 0x6e, 0x12, 0x44, 0x0a, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x61, 0x64, 0x64,
	0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x41, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x4f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x61,
	0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x41, 0x64, 0x64,
	0x4f, 0x6e, 0x52, 0x08, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x22, 0x48, 0x0a, 0x08,
	0x41, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x4b, 0x75,
	0x62, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x22, 0x3b, 0x0a, 0x0e, 0x4b, 0x75,
	0x62, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x72, 0x65, 0x44, 0x4e, 0x53, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x6e, 0x73,
	0x5f, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x44, 0x4e, 0x53, 0x41,
	0x64, 0x64, 0x4f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x72, 0x65, 0x44, 0x6e, 0x73, 0x41, 0x64, 0x64,
	0x4f, 0x6e, 0x22, 0x39, 0x0a, 0x0c, 0x43, 0x6f, 0x72, 0x65, 0x44, 0x4e, 0x53, 0x41, 0x64, 0x64,
	0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x57, 0x0a,
	0x12, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x4f, 0x6e, 0x12, 0x2d, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x1c, 0x41, 0x64, 0x64, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x4f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x52, 0x12, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6a, 0x61, 0x63, 0x6f, 0x31, 0x33,
	0x2f, 0x73, 0x64, 0x6b, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_addons_proto_rawDescOnce sync.Once
	file_addons_proto_rawDescData = file_addons_proto_rawDesc
)

func file_addons_proto_rawDescGZIP() []byte {
	file_addons_proto_rawDescOnce.Do(func() {
		file_addons_proto_rawDescData = protoimpl.X.CompressGZIP(file_addons_proto_rawDescData)
	})
	return file_addons_proto_rawDescData
}

var file_addons_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_addons_proto_goTypes = []interface{}{
	(*AddAddonsRequest)(nil),             // 0: codegen.AddAddonsRequest
	(*Addon)(nil),                        // 1: codegen.Addon
	(*AddAckAddOnRequest)(nil),           // 2: codegen.AddAckAddOnRequest
	(*AckAddOn)(nil),                     // 3: codegen.AckAddOn
	(*AddKubeProxyAddOnRequest)(nil),     // 4: codegen.AddKubeProxyAddOnRequest
	(*KubeProxyAddOn)(nil),               // 5: codegen.KubeProxyAddOn
	(*AddCoreDNSAddOnRequest)(nil),       // 6: codegen.AddCoreDNSAddOnRequest
	(*CoreDNSAddOn)(nil),                 // 7: codegen.CoreDNSAddOn
	(*MetricsServerAddOn)(nil),           // 8: codegen.MetricsServerAddOn
	(*AddMetricsServerAddOnRequest)(nil), // 9: codegen.AddMetricsServerAddOnRequest
}
var file_addons_proto_depIdxs = []int32{
	1, // 0: codegen.AddAddonsRequest.addons:type_name -> codegen.Addon
	3, // 1: codegen.Addon.ack_add_on:type_name -> codegen.AckAddOn
	5, // 2: codegen.Addon.kube_proxy_add_on:type_name -> codegen.KubeProxyAddOn
	3, // 3: codegen.AddAckAddOnRequest.ack_add_on:type_name -> codegen.AckAddOn
	5, // 4: codegen.AddKubeProxyAddOnRequest.kube_proxy_add_on:type_name -> codegen.KubeProxyAddOn
	7, // 5: codegen.AddCoreDNSAddOnRequest.core_dns_add_on:type_name -> codegen.CoreDNSAddOn
	8, // 6: codegen.AddMetricsServerAddOnRequest.metrics_server_add_on:type_name -> codegen.MetricsServerAddOn
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_addons_proto_init() }
func file_addons_proto_init() {
	if File_addons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_addons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAddonsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addon); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAckAddOnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckAddOn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKubeProxyAddOnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeProxyAddOn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCoreDNSAddOnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreDNSAddOn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsServerAddOn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_addons_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMetricsServerAddOnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_addons_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Addon_AckAddOn)(nil),
		(*Addon_KubeProxyAddOn)(nil),
	}
	file_addons_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_addons_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_addons_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_addons_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_addons_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_addons_proto_goTypes,
		DependencyIndexes: file_addons_proto_depIdxs,
		MessageInfos:      file_addons_proto_msgTypes,
	}.Build()
	File_addons_proto = out.File
	file_addons_proto_rawDesc = nil
	file_addons_proto_goTypes = nil
	file_addons_proto_depIdxs = nil
}
