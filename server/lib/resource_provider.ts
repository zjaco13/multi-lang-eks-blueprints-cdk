/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export interface AddResourceProviderRequest {
  name: string;
  resourceProvider: ResourceProvider | undefined;
}

export interface ResourceProvider {
  vpcProvider?: VpcProvider | undefined;
}

export interface VpcProvider {
  vpcId?: string | undefined;
}

function createBaseAddResourceProviderRequest(): AddResourceProviderRequest {
  return { name: "", resourceProvider: undefined };
}

export const AddResourceProviderRequest = {
  encode(message: AddResourceProviderRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.resourceProvider !== undefined) {
      ResourceProvider.encode(message.resourceProvider, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddResourceProviderRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddResourceProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.resourceProvider = ResourceProvider.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AddResourceProviderRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      resourceProvider: isSet(object.resourceProvider) ? ResourceProvider.fromJSON(object.resourceProvider) : undefined,
    };
  },

  toJSON(message: AddResourceProviderRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.resourceProvider !== undefined &&
      (obj.resourceProvider = message.resourceProvider ? ResourceProvider.toJSON(message.resourceProvider) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<AddResourceProviderRequest>, I>>(base?: I): AddResourceProviderRequest {
    return AddResourceProviderRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AddResourceProviderRequest>, I>>(object: I): AddResourceProviderRequest {
    const message = createBaseAddResourceProviderRequest();
    message.name = object.name ?? "";
    message.resourceProvider = (object.resourceProvider !== undefined && object.resourceProvider !== null)
      ? ResourceProvider.fromPartial(object.resourceProvider)
      : undefined;
    return message;
  },
};

function createBaseResourceProvider(): ResourceProvider {
  return { vpcProvider: undefined };
}

export const ResourceProvider = {
  encode(message: ResourceProvider, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vpcProvider !== undefined) {
      VpcProvider.encode(message.vpcProvider, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResourceProvider {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResourceProvider();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.vpcProvider = VpcProvider.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ResourceProvider {
    return { vpcProvider: isSet(object.vpcProvider) ? VpcProvider.fromJSON(object.vpcProvider) : undefined };
  },

  toJSON(message: ResourceProvider): unknown {
    const obj: any = {};
    message.vpcProvider !== undefined &&
      (obj.vpcProvider = message.vpcProvider ? VpcProvider.toJSON(message.vpcProvider) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ResourceProvider>, I>>(base?: I): ResourceProvider {
    return ResourceProvider.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ResourceProvider>, I>>(object: I): ResourceProvider {
    const message = createBaseResourceProvider();
    message.vpcProvider = (object.vpcProvider !== undefined && object.vpcProvider !== null)
      ? VpcProvider.fromPartial(object.vpcProvider)
      : undefined;
    return message;
  },
};

function createBaseVpcProvider(): VpcProvider {
  return { vpcId: undefined };
}

export const VpcProvider = {
  encode(message: VpcProvider, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vpcId !== undefined) {
      writer.uint32(10).string(message.vpcId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VpcProvider {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVpcProvider();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.vpcId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): VpcProvider {
    return { vpcId: isSet(object.vpcId) ? String(object.vpcId) : undefined };
  },

  toJSON(message: VpcProvider): unknown {
    const obj: any = {};
    message.vpcId !== undefined && (obj.vpcId = message.vpcId);
    return obj;
  },

  create<I extends Exact<DeepPartial<VpcProvider>, I>>(base?: I): VpcProvider {
    return VpcProvider.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<VpcProvider>, I>>(object: I): VpcProvider {
    const message = createBaseVpcProvider();
    message.vpcId = object.vpcId ?? undefined;
    return message;
  },
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
