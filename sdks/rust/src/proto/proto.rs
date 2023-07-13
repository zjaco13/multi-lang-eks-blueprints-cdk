// @generated
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddTeamsRequest {
    #[prost(message, repeated, tag="1")]
    pub teams: ::prost::alloc::vec::Vec<Team>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Team {
    #[prost(oneof="team::Team", tags="1, 2, 3")]
    pub team: ::core::option::Option<team::Team>,
}
/// Nested message and enum types in `Team`.
pub mod team {
    #[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Team {
        #[prost(message, tag="1")]
        GenericTeam(super::GenericTeam),
        #[prost(message, tag="2")]
        PlatformTeam(super::PlatformTeam),
        #[prost(message, tag="3")]
        ApplicationTeam(super::ApplicationTeam),
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GenericTeam {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PlatformTeam {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ApplicationTeam {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ApiResponse {
    #[prost(string, tag="1")]
    pub message: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateClusterRequest {
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
    #[prost(string, optional, tag="2")]
    pub name: ::core::option::Option<::prost::alloc::string::String>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddClusterProviderRequest {
    #[prost(string, tag="1")]
    pub cluster_name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AddResourceProviderRequest {
    #[prost(string, tag="1")]
    pub cluster_name: ::prost::alloc::string::String,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BuildClusterRequest {
    #[prost(string, tag="1")]
    pub cluster_name: ::prost::alloc::string::String,
    #[prost(string, optional, tag="2")]
    pub account: ::core::option::Option<::prost::alloc::string::String>,
    #[prost(string, optional, tag="3")]
    pub region: ::core::option::Option<::prost::alloc::string::String>,
}
include!("proto.tonic.rs");
// @@protoc_insertion_point(module)
