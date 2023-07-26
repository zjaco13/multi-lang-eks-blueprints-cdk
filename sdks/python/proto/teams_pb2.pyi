from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddTeamsRequest(_message.Message):
    __slots__ = ["cluster_name", "teams"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    TEAMS_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    teams: _containers.RepeatedCompositeFieldContainer[Team]
    def __init__(self, cluster_name: _Optional[str] = ..., teams: _Optional[_Iterable[_Union[Team, _Mapping]]] = ...) -> None: ...

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

class AddPlatformTeamRequest(_message.Message):
    __slots__ = ["cluster_name", "props"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    PROPS_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    props: TeamProps
    def __init__(self, cluster_name: _Optional[str] = ..., props: _Optional[_Union[TeamProps, _Mapping]] = ...) -> None: ...

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

class AddApplicationTeamRequest(_message.Message):
    __slots__ = ["cluster_name", "props"]
    CLUSTER_NAME_FIELD_NUMBER: _ClassVar[int]
    PROPS_FIELD_NUMBER: _ClassVar[int]
    cluster_name: str
    props: TeamProps
    def __init__(self, cluster_name: _Optional[str] = ..., props: _Optional[_Union[TeamProps, _Mapping]] = ...) -> None: ...

class TeamProps(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...
