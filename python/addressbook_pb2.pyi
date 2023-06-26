from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddressRequest(_message.Message):
    __slots__ = ["areaCode"]
    AREACODE_FIELD_NUMBER: _ClassVar[int]
    areaCode: int
    def __init__(self, areaCode: _Optional[int] = ...) -> None: ...

class Person(_message.Message):
    __slots__ = ["name", "age", "id", "email", "phones"]
    class PhoneType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
        MOBILE: _ClassVar[Person.PhoneType]
        HOME: _ClassVar[Person.PhoneType]
        WORK: _ClassVar[Person.PhoneType]
    MOBILE: Person.PhoneType
    HOME: Person.PhoneType
    WORK: Person.PhoneType
    class PhoneNumber(_message.Message):
        __slots__ = ["number", "type"]
        NUMBER_FIELD_NUMBER: _ClassVar[int]
        TYPE_FIELD_NUMBER: _ClassVar[int]
        number: str
        type: Person.PhoneType
        def __init__(self, number: _Optional[str] = ..., type: _Optional[_Union[Person.PhoneType, str]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    AGE_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PHONES_FIELD_NUMBER: _ClassVar[int]
    name: str
    age: int
    id: int
    email: str
    phones: _containers.RepeatedCompositeFieldContainer[Person.PhoneNumber]
    def __init__(self, name: _Optional[str] = ..., age: _Optional[int] = ..., id: _Optional[int] = ..., email: _Optional[str] = ..., phones: _Optional[_Iterable[_Union[Person.PhoneNumber, _Mapping]]] = ...) -> None: ...

class AddressBook(_message.Message):
    __slots__ = ["areaCode", "people"]
    AREACODE_FIELD_NUMBER: _ClassVar[int]
    PEOPLE_FIELD_NUMBER: _ClassVar[int]
    areaCode: int
    people: _containers.RepeatedCompositeFieldContainer[Person]
    def __init__(self, areaCode: _Optional[int] = ..., people: _Optional[_Iterable[_Union[Person, _Mapping]]] = ...) -> None: ...
