# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: teams.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bteams.proto\x12\x07\x63odegen\"E\n\x0f\x41\x64\x64TeamsRequest\x12\x14\n\x0c\x63luster_name\x18\x01 \x01(\t\x12\x1c\n\x05teams\x18\x02 \x03(\x0b\x32\r.codegen.Team\"\xa2\x01\n\x04Team\x12,\n\x0cgeneric_team\x18\x01 \x01(\x0b\x32\x14.codegen.GenericTeamH\x00\x12.\n\rplatform_team\x18\x02 \x01(\x0b\x32\x15.codegen.PlatformTeamH\x00\x12\x34\n\x10\x61pplication_team\x18\x03 \x01(\x0b\x32\x18.codegen.ApplicationTeamH\x00\x42\x06\n\x04team\"\x1b\n\x0bGenericTeam\x12\x0c\n\x04name\x18\x01 \x01(\t\"Q\n\x16\x41\x64\x64PlatformTeamRequest\x12\x14\n\x0c\x63luster_name\x18\x01 \x01(\t\x12!\n\x05props\x18\x02 \x01(\x0b\x32\x12.codegen.TeamProps\"\x1c\n\x0cPlatformTeam\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1f\n\x0f\x41pplicationTeam\x12\x0c\n\x04name\x18\x01 \x01(\t\"T\n\x19\x41\x64\x64\x41pplicationTeamRequest\x12\x14\n\x0c\x63luster_name\x18\x01 \x01(\t\x12!\n\x05props\x18\x02 \x01(\x0b\x32\x12.codegen.TeamProps\"\x19\n\tTeamProps\x12\x0c\n\x04name\x18\x01 \x01(\tB(Z&github.com/zjaco13/sdks/go-sdk/codegenb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'teams_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z&github.com/zjaco13/sdks/go-sdk/codegen'
  _globals['_ADDTEAMSREQUEST']._serialized_start=24
  _globals['_ADDTEAMSREQUEST']._serialized_end=93
  _globals['_TEAM']._serialized_start=96
  _globals['_TEAM']._serialized_end=258
  _globals['_GENERICTEAM']._serialized_start=260
  _globals['_GENERICTEAM']._serialized_end=287
  _globals['_ADDPLATFORMTEAMREQUEST']._serialized_start=289
  _globals['_ADDPLATFORMTEAMREQUEST']._serialized_end=370
  _globals['_PLATFORMTEAM']._serialized_start=372
  _globals['_PLATFORMTEAM']._serialized_end=400
  _globals['_APPLICATIONTEAM']._serialized_start=402
  _globals['_APPLICATIONTEAM']._serialized_end=433
  _globals['_ADDAPPLICATIONTEAMREQUEST']._serialized_start=435
  _globals['_ADDAPPLICATIONTEAMREQUEST']._serialized_end=519
  _globals['_TEAMPROPS']._serialized_start=521
  _globals['_TEAMPROPS']._serialized_end=546
# @@protoc_insertion_point(module_scope)
