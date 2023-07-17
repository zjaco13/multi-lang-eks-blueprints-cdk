from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddResourceProviderRequest(_message.Message):
    __slots__ = ["name", "resource_provider"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_PROVIDER_FIELD_NUMBER: _ClassVar[int]
    name: str
    resource_provider: ResourceProvider
    def __init__(self, name: _Optional[str] = ..., resource_provider: _Optional[_Union[ResourceProvider, _Mapping]] = ...) -> None: ...

class ResourceProvider(_message.Message):
    __slots__ = ["vpc_provider"]
    VPC_PROVIDER_FIELD_NUMBER: _ClassVar[int]
    vpc_provider: VpcProvider
    def __init__(self, vpc_provider: _Optional[_Union[VpcProvider, _Mapping]] = ...) -> None: ...

class VpcProvider(_message.Message):
    __slots__ = ["vpcId"]
    VPCID_FIELD_NUMBER: _ClassVar[int]
    vpcId: str
    def __init__(self, vpcId: _Optional[str] = ...) -> None: ...
