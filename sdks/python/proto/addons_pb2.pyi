from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddAddonsRequest(_message.Message):
    __slots__ = ["addons"]
    ADDONS_FIELD_NUMBER: _ClassVar[int]
    addons: _containers.RepeatedCompositeFieldContainer[Addon]
    def __init__(self, addons: _Optional[_Iterable[_Union[Addon, _Mapping]]] = ...) -> None: ...

class Addon(_message.Message):
    __slots__ = ["ack_add_on"]
    ACK_ADD_ON_FIELD_NUMBER: _ClassVar[int]
    ack_add_on: AckAddOn
    def __init__(self, ack_add_on: _Optional[_Union[AckAddOn, _Mapping]] = ...) -> None: ...

class AckAddOn(_message.Message):
    __slots__ = ["id", "serviceName"]
    ID_FIELD_NUMBER: _ClassVar[int]
    SERVICENAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    serviceName: str
    def __init__(self, id: _Optional[str] = ..., serviceName: _Optional[str] = ...) -> None: ...
