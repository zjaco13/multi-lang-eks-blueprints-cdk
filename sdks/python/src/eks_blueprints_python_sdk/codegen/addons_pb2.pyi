from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddAddonsRequest(_message.Message):
    __slots__ = ["cluster_name", "addons"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    ADDONS_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    addons: _containers.RepeatedCompositeFieldContainer[Addon]
    def __init__(self, cluster_name: _Optional[str] = ..., addons: _Optional[_Iterable[_Union[Addon, _Mapping]]] = ...) -> None: ...

class Addon(_message.Message):
    __slots__ = ["ack_add_on", "kube_proxy_add_on"]
    ACK_ADD_ON_FIELD_NUMBER: _ClassVar[int]
    KUBE_PROXY_ADD_ON_FIELD_NUMBER: _ClassVar[int]
    ack_add_on: AckAddOn
    kube_proxy_add_on: KubeProxyAddOn
    def __init__(self, ack_add_on: _Optional[_Union[AckAddOn, _Mapping]] = ..., kube_proxy_add_on: _Optional[_Union[KubeProxyAddOn, _Mapping]] = ...) -> None: ...

class AddAckAddOnRequest(_message.Message):
    __slots__ = ["cluster_name", "ack_add_on"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    ACK_ADD_ON_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    ack_add_on: AckAddOn
    def __init__(self, cluster_name: _Optional[str] = ..., ack_add_on: _Optional[_Union[AckAddOn, _Mapping]] = ...) -> None: ...

class AckAddOn(_message.Message):
    __slots__ = ["id", "serviceName"]
    ID_FIELD_NUMBER: _ClassVar[int]
    SERVICENAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    serviceName: str
    def __init__(self, id: _Optional[str] = ..., serviceName: _Optional[str] = ...) -> None: ...

class AddKubeProxyAddOnRequest(_message.Message):
    __slots__ = ["cluster_name", "kube_proxy_add_on"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    KUBE_PROXY_ADD_ON_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    kube_proxy_add_on: KubeProxyAddOn
    def __init__(self, cluster_name: _Optional[str] = ..., kube_proxy_add_on: _Optional[_Union[KubeProxyAddOn, _Mapping]] = ...) -> None: ...

class KubeProxyAddOn(_message.Message):
    __slots__ = ["version"]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    version: str
    def __init__(self, version: _Optional[str] = ...) -> None: ...

class AddCoreDNSAddOnRequest(_message.Message):
    __slots__ = ["cluster_name", "core_dns_add_on"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    CORE_DNS_ADD_ON_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    core_dns_add_on: CoreDNSAddOn
    def __init__(self, cluster_name: _Optional[str] = ..., core_dns_add_on: _Optional[_Union[CoreDNSAddOn, _Mapping]] = ...) -> None: ...

class CoreDNSAddOn(_message.Message):
    __slots__ = ["version"]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    version: str
    def __init__(self, version: _Optional[str] = ...) -> None: ...

class MetricsServerAddOn(_message.Message):
    __slots__ = ["createNamespace"]
    CREATENAMESPACE_FIELD_NUMBER: _ClassVar[int]
    createNamespace: bool
    def __init__(self, createNamespace: bool = ...) -> None: ...

class AddMetricsServerAddOnRequest(_message.Message):
    __slots__ = ["cluster_name", "metrics_server_add_on"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    METRICS_SERVER_ADD_ON_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    metrics_server_add_on: MetricsServerAddOn
    def __init__(self, cluster_name: _Optional[str] = ..., metrics_server_add_on: _Optional[_Union[MetricsServerAddOn, _Mapping]] = ...) -> None: ...
