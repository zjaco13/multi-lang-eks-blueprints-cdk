from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddTeamsRequest(_message.Message):
    __slots__ = ["teams"]
    TEAMS_FIELD_NUMBER: _ClassVar[int]
    teams: _containers.RepeatedCompositeFieldContainer[Team]
    def __init__(self, teams: _Optional[_Iterable[_Union[Team, _Mapping]]] = ...) -> None: ...

class Team(_message.Message):
    __slots__ = ["generic_team", "platform_team", "application_team"]
    GENERIC_TEAM_FIELD_NUMBER: _ClassVar[int]
    PLATFORM_TEAM_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_TEAM_FIELD_NUMBER: _ClassVar[int]
    generic_team: GenericTeam
    platform_team: PlatformTeam
    application_team: ApplicationTeam
    def __init__(self, generic_team: _Optional[_Union[GenericTeam, _Mapping]] = ..., platform_team: _Optional[_Union[PlatformTeam, _Mapping]] = ..., application_team: _Optional[_Union[ApplicationTeam, _Mapping]] = ...) -> None: ...

class GenericTeam(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class PlatformTeam(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class ApplicationTeam(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...
