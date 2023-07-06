from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class EksBlueprint(_message.Message):
    __slots__ = ["id", "name"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class EksBlueprintResponse(_message.Message):
    __slots__ = ["resp"]
    RESP_FIELD_NUMBER: _ClassVar[int]
    resp: str
    def __init__(self, resp: _Optional[str] = ...) -> None: ...
