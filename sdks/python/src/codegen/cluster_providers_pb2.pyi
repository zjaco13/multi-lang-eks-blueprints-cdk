from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddClusterProviderRequest(_message.Message):
    __slots__ = ["cluster_name", "cluster_provider"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    CLUSTER_PROVIDER_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    cluster_provider: ClusterProvider
    def __init__(self, cluster_name: _Optional[str] = ..., cluster_provider: _Optional[_Union[ClusterProvider, _Mapping]] = ...) -> None: ...

class ClusterProvider(_message.Message):
    __slots__ = ["asg_cluster_provider", "mng_cluster_provider"]
    ASG_CLUSTER_PROVIDER_FIELD_NUMBER: _ClassVar[int]
    MNG_CLUSTER_PROVIDER_FIELD_NUMBER: _ClassVar[int]
    asg_cluster_provider: AsgClusterProvider
    mng_cluster_provider: MngClusterProvider
    def __init__(self, asg_cluster_provider: _Optional[_Union[AsgClusterProvider, _Mapping]] = ..., mng_cluster_provider: _Optional[_Union[MngClusterProvider, _Mapping]] = ...) -> None: ...

class AddAsgClusterProviderRequest(_message.Message):
    __slots__ = ["cluster_name", "asg_cluster_provider"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    ASG_CLUSTER_PROVIDER_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    asg_cluster_provider: AsgClusterProvider
    def __init__(self, cluster_name: _Optional[str] = ..., asg_cluster_provider: _Optional[_Union[AsgClusterProvider, _Mapping]] = ...) -> None: ...

class AsgClusterProvider(_message.Message):
    __slots__ = ["name", "version", "id"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    name: str
    version: str
    id: str
    def __init__(self, name: _Optional[str] = ..., version: _Optional[str] = ..., id: _Optional[str] = ...) -> None: ...

class AddMngClusterProviderRequest(_message.Message):
    __slots__ = ["cluster_name", "mng_cluster_provider"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    MNG_CLUSTER_PROVIDER_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    mng_cluster_provider: MngClusterProvider
    def __init__(self, cluster_name: _Optional[str] = ..., mng_cluster_provider: _Optional[_Union[MngClusterProvider, _Mapping]] = ...) -> None: ...

class MngClusterProvider(_message.Message):
    __slots__ = ["name", "version"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    name: str
    version: str
    def __init__(self, name: _Optional[str] = ..., version: _Optional[str] = ...) -> None: ...
