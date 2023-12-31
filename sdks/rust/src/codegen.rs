#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddTeamsRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, repeated, tag = "2")]
    pub teams: ::prost::alloc::vec::Vec<Team>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Team {
    #[prost(oneof = "team::Team", tags = "1, 2, 3")]
    pub team: ::core::option::Option<team::Team>,
}
/// Nested message and enum types in `Team`.
pub mod team {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Team {
        #[prost(message, tag = "1")]
        GenericTeam(super::GenericTeam),
        #[prost(message, tag = "2")]
        PlatformTeam(super::PlatformTeam),
        #[prost(message, tag = "3")]
        ApplicationTeam(super::ApplicationTeam),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GenericTeam {
    #[prost(string, tag = "1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddPlatformTeamRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub props: ::core::option::Option<TeamProps>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PlatformTeam {
    #[prost(string, tag = "1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ApplicationTeam {
    #[prost(string, tag = "1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddApplicationTeamRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub props: ::core::option::Option<TeamProps>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TeamProps {
    #[prost(string, tag = "1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddClusterProviderRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub cluster_provider: ::core::option::Option<ClusterProvider>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ClusterProvider {
    #[prost(oneof = "cluster_provider::ClusterProvider", tags = "1, 2")]
    pub cluster_provider: ::core::option::Option<cluster_provider::ClusterProvider>,
}
/// Nested message and enum types in `ClusterProvider`.
pub mod cluster_provider {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum ClusterProvider {
        #[prost(message, tag = "1")]
        AsgClusterProvider(super::AsgClusterProvider),
        #[prost(message, tag = "2")]
        MngClusterProvider(super::MngClusterProvider),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddAsgClusterProviderRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub asg_cluster_provider: ::core::option::Option<AsgClusterProvider>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AsgClusterProvider {
    #[prost(string, optional, tag = "1")]
    pub name: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, tag = "2")]
    pub version: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub id: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddMngClusterProviderRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub mng_cluster_provider: ::core::option::Option<MngClusterProvider>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MngClusterProvider {
    #[prost(string, optional, tag = "1")]
    pub name: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, tag = "2")]
    pub version: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddResourceProviderRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "3")]
    pub resource_provider: ::core::option::Option<ResourceProvider>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ResourceProvider {
    #[prost(oneof = "resource_provider::ResourceProvider", tags = "1")]
    pub resource_provider: ::core::option::Option<resource_provider::ResourceProvider>,
}
/// Nested message and enum types in `ResourceProvider`.
pub mod resource_provider {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum ResourceProvider {
        #[prost(message, tag = "1")]
        VpcProvider(super::VpcProvider),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddVpcProviderRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "3")]
    pub vpc_provider: ::core::option::Option<VpcProvider>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct VpcProvider {
    #[prost(string, optional, tag = "1")]
    pub vpc_id: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddAddonsRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, repeated, tag = "2")]
    pub addons: ::prost::alloc::vec::Vec<Addon>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Addon {
    #[prost(oneof = "addon::Addon", tags = "1, 2")]
    pub addon: ::core::option::Option<addon::Addon>,
}
/// Nested message and enum types in `Addon`.
pub mod addon {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Addon {
        #[prost(message, tag = "1")]
        AckAddOn(super::AckAddOn),
        #[prost(message, tag = "2")]
        KubeProxyAddOn(super::KubeProxyAddOn),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddAckAddOnRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub ack_add_on: ::core::option::Option<AckAddOn>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AckAddOn {
    #[prost(string, optional, tag = "1")]
    pub id: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, tag = "2")]
    pub service_name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddKubeProxyAddOnRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub kube_proxy_add_on: ::core::option::Option<KubeProxyAddOn>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct KubeProxyAddOn {
    #[prost(string, optional, tag = "1")]
    pub version: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddCoreDnsAddOnRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub core_dns_add_on: ::core::option::Option<CoreDnsAddOn>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CoreDnsAddOn {
    #[prost(string, optional, tag = "1")]
    pub version: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MetricsServerAddOn {
    #[prost(bool, optional, tag = "1")]
    pub create_namespace: ::core::option::Option<bool>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddMetricsServerAddOnRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub metrics_server_add_on: ::core::option::Option<MetricsServerAddOn>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ApiResponse {
    #[prost(string, tag = "1")]
    pub message: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateClusterRequest {
    #[prost(string, tag = "1")]
    pub id: ::prost::alloc::string::String,
    #[prost(string, optional, tag = "2")]
    pub name: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BuildClusterRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(string, optional, tag = "2")]
    pub account: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, optional, tag = "3")]
    pub region: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CloneClusterRequest {
    #[prost(string, tag = "1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub new_cluster_id: ::prost::alloc::string::String,
    #[prost(string, optional, tag = "3")]
    pub new_cluster_name: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, optional, tag = "4")]
    pub region: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, optional, tag = "5")]
    pub account: ::core::option::Option<::prost::alloc::string::String>,
}
/// Generated client implementations.
pub mod cluster_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct ClusterServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl ClusterServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> ClusterServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> ClusterServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            ClusterServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn create_cluster(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateClusterRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/CreateCluster",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "CreateCluster"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn build_cluster(
            &mut self,
            request: impl tonic::IntoRequest<super::BuildClusterRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/BuildCluster",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "BuildCluster"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn clone_cluster(
            &mut self,
            request: impl tonic::IntoRequest<super::CloneClusterRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/CloneCluster",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "CloneCluster"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_platform_team(
            &mut self,
            request: impl tonic::IntoRequest<super::AddPlatformTeamRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddPlatformTeam",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddPlatformTeam"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_application_team(
            &mut self,
            request: impl tonic::IntoRequest<super::AddApplicationTeamRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddApplicationTeam",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddApplicationTeam"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_teams(
            &mut self,
            request: impl tonic::IntoRequest<super::AddTeamsRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddTeams",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddTeams"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_mng_cluster_provider(
            &mut self,
            request: impl tonic::IntoRequest<super::AddMngClusterProviderRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddMngClusterProvider",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("codegen.ClusterService", "AddMngClusterProvider"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_asg_cluster_provider(
            &mut self,
            request: impl tonic::IntoRequest<super::AddAsgClusterProviderRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddAsgClusterProvider",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("codegen.ClusterService", "AddAsgClusterProvider"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_cluster_provider(
            &mut self,
            request: impl tonic::IntoRequest<super::AddClusterProviderRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddClusterProvider",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddClusterProvider"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_vpc_provider(
            &mut self,
            request: impl tonic::IntoRequest<super::AddVpcProviderRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddVpcProvider",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddVpcProvider"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_resource_provider(
            &mut self,
            request: impl tonic::IntoRequest<super::AddResourceProviderRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddResourceProvider",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("codegen.ClusterService", "AddResourceProvider"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_ack_add_on(
            &mut self,
            request: impl tonic::IntoRequest<super::AddAckAddOnRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddAckAddOn",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddAckAddOn"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_kube_proxy_add_on(
            &mut self,
            request: impl tonic::IntoRequest<super::AddKubeProxyAddOnRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddKubeProxyAddOn",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddKubeProxyAddOn"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_core_dns_add_on(
            &mut self,
            request: impl tonic::IntoRequest<super::AddCoreDnsAddOnRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddCoreDNSAddOn",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddCoreDNSAddOn"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_metrics_server_add_on(
            &mut self,
            request: impl tonic::IntoRequest<super::AddMetricsServerAddOnRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddMetricsServerAddOn",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("codegen.ClusterService", "AddMetricsServerAddOn"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn add_addons(
            &mut self,
            request: impl tonic::IntoRequest<super::AddAddonsRequest>,
        ) -> std::result::Result<tonic::Response<super::ApiResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/codegen.ClusterService/AddAddons",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("codegen.ClusterService", "AddAddons"));
            self.inner.unary(req, path, codec).await
        }
    }
}
