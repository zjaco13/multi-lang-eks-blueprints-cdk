import team_pb2 as _team_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class APIResponse(_message.Message):
    __slots__ = ["message"]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class CreateClusterRequest(_message.Message):
    __slots__ = ["id", "name"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class AddClusterProviderRequest(_message.Message):
    __slots__ = ["clusterName"]
    CLUSTERNAME_FIELD_NUMBER: _ClassVar[int]
    clusterName: str
    def __init__(self, clusterName: _Optional[str] = ...) -> None: ...

class AddResourceProviderRequest(_message.Message):
    __slots__ = ["clusterName"]
    CLUSTERNAME_FIELD_NUMBER: _ClassVar[int]
    clusterName: str
    def __init__(self, clusterName: _Optional[str] = ...) -> None: ...

class BuildClusterRequest(_message.Message):
    __slots__ = ["clusterName", "account", "region"]
    CLUSTERNAME_FIELD_NUMBER: _ClassVar[int]
    ACCOUNT_FIELD_NUMBER: _ClassVar[int]
    REGION_FIELD_NUMBER: _ClassVar[int]
    clusterName: str
    account: str
    region: str
    def __init__(self, clusterName: _Optional[str] = ..., account: _Optional[str] = ..., region: _Optional[str] = ...) -> None: ...
