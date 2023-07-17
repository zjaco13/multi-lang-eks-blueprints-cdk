/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export interface AddAddonsRequest {
  addons: Addon[];
}

export interface Addon {
  ackAddOn?: AckAddOn | undefined;
}

export interface AckAddOn {
  id?: string | undefined;
  serviceName: string;
}

function createBaseAddAddonsRequest(): AddAddonsRequest {
  return { addons: [] };
}

export const AddAddonsRequest = {
  encode(message: AddAddonsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.addons) {
      Addon.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddAddonsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddAddonsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.addons.push(Addon.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AddAddonsRequest {
    return { addons: Array.isArray(object?.addons) ? object.addons.map((e: any) => Addon.fromJSON(e)) : [] };
  },

  toJSON(message: AddAddonsRequest): unknown {
    const obj: any = {};
    if (message.addons) {
      obj.addons = message.addons.map((e) => e ? Addon.toJSON(e) : undefined);
    } else {
      obj.addons = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AddAddonsRequest>, I>>(base?: I): AddAddonsRequest {
    return AddAddonsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AddAddonsRequest>, I>>(object: I): AddAddonsRequest {
    const message = createBaseAddAddonsRequest();
    message.addons = object.addons?.map((e) => Addon.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAddon(): Addon {
  return { ackAddOn: undefined };
}

export const Addon = {
  encode(message: Addon, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ackAddOn !== undefined) {
      AckAddOn.encode(message.ackAddOn, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Addon {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddon();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.ackAddOn = AckAddOn.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Addon {
    return { ackAddOn: isSet(object.ackAddOn) ? AckAddOn.fromJSON(object.ackAddOn) : undefined };
  },

  toJSON(message: Addon): unknown {
    const obj: any = {};
    message.ackAddOn !== undefined && (obj.ackAddOn = message.ackAddOn ? AckAddOn.toJSON(message.ackAddOn) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<Addon>, I>>(base?: I): Addon {
    return Addon.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Addon>, I>>(object: I): Addon {
    const message = createBaseAddon();
    message.ackAddOn = (object.ackAddOn !== undefined && object.ackAddOn !== null)
      ? AckAddOn.fromPartial(object.ackAddOn)
      : undefined;
    return message;
  },
};

function createBaseAckAddOn(): AckAddOn {
  return { id: undefined, serviceName: "" };
}

export const AckAddOn = {
  encode(message: AckAddOn, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== undefined) {
      writer.uint32(10).string(message.id);
    }
    if (message.serviceName !== "") {
      writer.uint32(18).string(message.serviceName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AckAddOn {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAckAddOn();
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

          message.serviceName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AckAddOn {
    return {
      id: isSet(object.id) ? String(object.id) : undefined,
      serviceName: isSet(object.serviceName) ? String(object.serviceName) : "",
    };
  },

  toJSON(message: AckAddOn): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.serviceName !== undefined && (obj.serviceName = message.serviceName);
    return obj;
  },

  create<I extends Exact<DeepPartial<AckAddOn>, I>>(base?: I): AckAddOn {
    return AckAddOn.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AckAddOn>, I>>(object: I): AckAddOn {
    const message = createBaseAckAddOn();
    message.id = object.id ?? undefined;
    message.serviceName = object.serviceName ?? "";
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
