/* eslint-disable */
import {
  CallOptions,
  ChannelCredentials,
  Client,
  ClientOptions,
  ClientUnaryCall,
  handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  ServiceError,
  UntypedServiceImplementation,
} from "@grpc/grpc-js";
import * as _m0 from "protobufjs/minimal";
import { AddAddonsRequest } from "./addons";
import { AddClusterProviderRequest } from "./cluster_provider";
import { AddResourceProviderRequest } from "./resource_provider";
import { AddTeamsRequest } from "./team";

export interface APIResponse {
  message: string;
}

export interface CreateClusterRequest {
  id: string;
  name?: string | undefined;
}

export interface BuildClusterRequest {
  clusterName: string;
  account?: string | undefined;
  region?: string | undefined;
}

function createBaseAPIResponse(): APIResponse {
  return { message: "" };
}

export const APIResponse = {
  encode(message: APIResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): APIResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAPIResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.message = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): APIResponse {
    return { message: isSet(object.message) ? String(object.message) : "" };
  },

  toJSON(message: APIResponse): unknown {
    const obj: any = {};
    message.message !== undefined && (obj.message = message.message);
    return obj;
  },

  create<I extends Exact<DeepPartial<APIResponse>, I>>(base?: I): APIResponse {
    return APIResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<APIResponse>, I>>(object: I): APIResponse {
    const message = createBaseAPIResponse();
    message.message = object.message ?? "";
    return message;
  },
};

function createBaseCreateClusterRequest(): CreateClusterRequest {
  return { id: "", name: undefined };
}

export const CreateClusterRequest = {
  encode(message: CreateClusterRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== undefined) {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateClusterRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateClusterRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateClusterRequest {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : undefined,
    };
  },

  toJSON(message: CreateClusterRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateClusterRequest>, I>>(base?: I): CreateClusterRequest {
    return CreateClusterRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateClusterRequest>, I>>(object: I): CreateClusterRequest {
    const message = createBaseCreateClusterRequest();
    message.id = object.id ?? "";
    message.name = object.name ?? undefined;
    return message;
  },
};

function createBaseBuildClusterRequest(): BuildClusterRequest {
  return { clusterName: "", account: undefined, region: undefined };
}

export const BuildClusterRequest = {
  encode(message: BuildClusterRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.clusterName !== "") {
      writer.uint32(10).string(message.clusterName);
    }
    if (message.account !== undefined) {
      writer.uint32(18).string(message.account);
    }
    if (message.region !== undefined) {
      writer.uint32(26).string(message.region);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BuildClusterRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBuildClusterRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.clusterName = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.account = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.region = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): BuildClusterRequest {
    return {
      clusterName: isSet(object.clusterName) ? String(object.clusterName) : "",
      account: isSet(object.account) ? String(object.account) : undefined,
      region: isSet(object.region) ? String(object.region) : undefined,
    };
  },

  toJSON(message: BuildClusterRequest): unknown {
    const obj: any = {};
    message.clusterName !== undefined && (obj.clusterName = message.clusterName);
    message.account !== undefined && (obj.account = message.account);
    message.region !== undefined && (obj.region = message.region);
    return obj;
  },

  create<I extends Exact<DeepPartial<BuildClusterRequest>, I>>(base?: I): BuildClusterRequest {
    return BuildClusterRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BuildClusterRequest>, I>>(object: I): BuildClusterRequest {
    const message = createBaseBuildClusterRequest();
    message.clusterName = object.clusterName ?? "";
    message.account = object.account ?? undefined;
    message.region = object.region ?? undefined;
    return message;
  },
};

export type ClusterServiceService = typeof ClusterServiceService;
export const ClusterServiceService = {
  createCluster: {
    path: "/proto.ClusterService/CreateCluster",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateClusterRequest) => Buffer.from(CreateClusterRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateClusterRequest.decode(value),
    responseSerialize: (value: APIResponse) => Buffer.from(APIResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => APIResponse.decode(value),
  },
  addTeams: {
    path: "/proto.ClusterService/AddTeams",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AddTeamsRequest) => Buffer.from(AddTeamsRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AddTeamsRequest.decode(value),
    responseSerialize: (value: APIResponse) => Buffer.from(APIResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => APIResponse.decode(value),
  },
  addClusterProvider: {
    path: "/proto.ClusterService/AddClusterProvider",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AddClusterProviderRequest) =>
      Buffer.from(AddClusterProviderRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AddClusterProviderRequest.decode(value),
    responseSerialize: (value: APIResponse) => Buffer.from(APIResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => APIResponse.decode(value),
  },
  addResourceProvider: {
    path: "/proto.ClusterService/AddResourceProvider",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AddResourceProviderRequest) =>
      Buffer.from(AddResourceProviderRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AddResourceProviderRequest.decode(value),
    responseSerialize: (value: APIResponse) => Buffer.from(APIResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => APIResponse.decode(value),
  },
  addAddons: {
    path: "/proto.ClusterService/AddAddons",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: AddAddonsRequest) => Buffer.from(AddAddonsRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => AddAddonsRequest.decode(value),
    responseSerialize: (value: APIResponse) => Buffer.from(APIResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => APIResponse.decode(value),
  },
  /** rpc SendCluster (EksBlueprint) returns (EksBlueprintResponse); */
  buildCluster: {
    path: "/proto.ClusterService/BuildCluster",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: BuildClusterRequest) => Buffer.from(BuildClusterRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => BuildClusterRequest.decode(value),
    responseSerialize: (value: APIResponse) => Buffer.from(APIResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => APIResponse.decode(value),
  },
} as const;

export interface ClusterServiceServer extends UntypedServiceImplementation {
  createCluster: handleUnaryCall<CreateClusterRequest, APIResponse>;
  addTeams: handleUnaryCall<AddTeamsRequest, APIResponse>;
  addClusterProvider: handleUnaryCall<AddClusterProviderRequest, APIResponse>;
  addResourceProvider: handleUnaryCall<AddResourceProviderRequest, APIResponse>;
  addAddons: handleUnaryCall<AddAddonsRequest, APIResponse>;
  /** rpc SendCluster (EksBlueprint) returns (EksBlueprintResponse); */
  buildCluster: handleUnaryCall<BuildClusterRequest, APIResponse>;
}

export interface ClusterServiceClient extends Client {
  createCluster(
    request: CreateClusterRequest,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  createCluster(
    request: CreateClusterRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  createCluster(
    request: CreateClusterRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addTeams(
    request: AddTeamsRequest,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addTeams(
    request: AddTeamsRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addTeams(
    request: AddTeamsRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addClusterProvider(
    request: AddClusterProviderRequest,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addClusterProvider(
    request: AddClusterProviderRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addClusterProvider(
    request: AddClusterProviderRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addResourceProvider(
    request: AddResourceProviderRequest,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addResourceProvider(
    request: AddResourceProviderRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addResourceProvider(
    request: AddResourceProviderRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addAddons(
    request: AddAddonsRequest,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addAddons(
    request: AddAddonsRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  addAddons(
    request: AddAddonsRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  /** rpc SendCluster (EksBlueprint) returns (EksBlueprintResponse); */
  buildCluster(
    request: BuildClusterRequest,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  buildCluster(
    request: BuildClusterRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
  buildCluster(
    request: BuildClusterRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: APIResponse) => void,
  ): ClientUnaryCall;
}

export const ClusterServiceClient = makeGenericClientConstructor(
  ClusterServiceService,
  "proto.ClusterService",
) as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): ClusterServiceClient;
  service: typeof ClusterServiceService;
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
