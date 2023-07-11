/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 4.23.3
 * source: cluster.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as pb_1 from "google-protobuf";
import * as grpc_1 from "@grpc/grpc-js";
export namespace proto_test {
    export class APIResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            message?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("message" in data && data.message != undefined) {
                    this.message = data.message;
                }
            }
        }
        get message() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set message(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            message?: string;
        }): APIResponse {
            const message = new APIResponse({});
            if (data.message != null) {
                message.message = data.message;
            }
            return message;
        }
        toObject() {
            const data: {
                message?: string;
            } = {};
            if (this.message != null) {
                data.message = this.message;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.message.length)
                writer.writeString(1, this.message);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): APIResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new APIResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.message = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): APIResponse {
            return APIResponse.deserialize(bytes);
        }
    }
    export class CreateClusterRequest extends pb_1.Message {
        #one_of_decls: number[][] = [[2]];
        constructor(data?: any[] | ({
            id?: string;
        } & (({
            name?: string;
        })))) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("id" in data && data.id != undefined) {
                    this.id = data.id;
                }
                if ("name" in data && data.name != undefined) {
                    this.name = data.name;
                }
            }
        }
        get id() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set id(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get name() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set name(value: string) {
            pb_1.Message.setOneofField(this, 2, this.#one_of_decls[0], value);
        }
        get has_name() {
            return pb_1.Message.getField(this, 2) != null;
        }
        get _name() {
            const cases: {
                [index: number]: "none" | "name";
            } = {
                0: "none",
                2: "name"
            };
            return cases[pb_1.Message.computeOneofCase(this, [2])];
        }
        static fromObject(data: {
            id?: string;
            name?: string;
        }): CreateClusterRequest {
            const message = new CreateClusterRequest({});
            if (data.id != null) {
                message.id = data.id;
            }
            if (data.name != null) {
                message.name = data.name;
            }
            return message;
        }
        toObject() {
            const data: {
                id?: string;
                name?: string;
            } = {};
            if (this.id != null) {
                data.id = this.id;
            }
            if (this.name != null) {
                data.name = this.name;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.id.length)
                writer.writeString(1, this.id);
            if (this.has_name)
                writer.writeString(2, this.name);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreateClusterRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreateClusterRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.id = reader.readString();
                        break;
                    case 2:
                        message.name = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): CreateClusterRequest {
            return CreateClusterRequest.deserialize(bytes);
        }
    }
    export class AddTeamsRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            clusterName?: string;
            teamName?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("clusterName" in data && data.clusterName != undefined) {
                    this.clusterName = data.clusterName;
                }
                if ("teamName" in data && data.teamName != undefined) {
                    this.teamName = data.teamName;
                }
            }
        }
        get clusterName() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set clusterName(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get teamName() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set teamName(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            clusterName?: string;
            teamName?: string;
        }): AddTeamsRequest {
            const message = new AddTeamsRequest({});
            if (data.clusterName != null) {
                message.clusterName = data.clusterName;
            }
            if (data.teamName != null) {
                message.teamName = data.teamName;
            }
            return message;
        }
        toObject() {
            const data: {
                clusterName?: string;
                teamName?: string;
            } = {};
            if (this.clusterName != null) {
                data.clusterName = this.clusterName;
            }
            if (this.teamName != null) {
                data.teamName = this.teamName;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.clusterName.length)
                writer.writeString(1, this.clusterName);
            if (this.teamName.length)
                writer.writeString(2, this.teamName);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AddTeamsRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AddTeamsRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.clusterName = reader.readString();
                        break;
                    case 2:
                        message.teamName = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AddTeamsRequest {
            return AddTeamsRequest.deserialize(bytes);
        }
    }
    export class AddClusterProviderRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            clusterName?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("clusterName" in data && data.clusterName != undefined) {
                    this.clusterName = data.clusterName;
                }
            }
        }
        get clusterName() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set clusterName(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            clusterName?: string;
        }): AddClusterProviderRequest {
            const message = new AddClusterProviderRequest({});
            if (data.clusterName != null) {
                message.clusterName = data.clusterName;
            }
            return message;
        }
        toObject() {
            const data: {
                clusterName?: string;
            } = {};
            if (this.clusterName != null) {
                data.clusterName = this.clusterName;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.clusterName.length)
                writer.writeString(1, this.clusterName);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AddClusterProviderRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AddClusterProviderRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.clusterName = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AddClusterProviderRequest {
            return AddClusterProviderRequest.deserialize(bytes);
        }
    }
    export class AddResourceProviderRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            clusterName?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("clusterName" in data && data.clusterName != undefined) {
                    this.clusterName = data.clusterName;
                }
            }
        }
        get clusterName() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set clusterName(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            clusterName?: string;
        }): AddResourceProviderRequest {
            const message = new AddResourceProviderRequest({});
            if (data.clusterName != null) {
                message.clusterName = data.clusterName;
            }
            return message;
        }
        toObject() {
            const data: {
                clusterName?: string;
            } = {};
            if (this.clusterName != null) {
                data.clusterName = this.clusterName;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.clusterName.length)
                writer.writeString(1, this.clusterName);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AddResourceProviderRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AddResourceProviderRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.clusterName = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AddResourceProviderRequest {
            return AddResourceProviderRequest.deserialize(bytes);
        }
    }
    export class BuildClusterRequest extends pb_1.Message {
        #one_of_decls: number[][] = [[2], [3]];
        constructor(data?: any[] | ({
            clusterName?: string;
        } & (({
            account?: string;
        }) | ({
            region?: string;
        })))) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("clusterName" in data && data.clusterName != undefined) {
                    this.clusterName = data.clusterName;
                }
                if ("account" in data && data.account != undefined) {
                    this.account = data.account;
                }
                if ("region" in data && data.region != undefined) {
                    this.region = data.region;
                }
            }
        }
        get clusterName() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set clusterName(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get account() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set account(value: string) {
            pb_1.Message.setOneofField(this, 2, this.#one_of_decls[0], value);
        }
        get has_account() {
            return pb_1.Message.getField(this, 2) != null;
        }
        get region() {
            return pb_1.Message.getFieldWithDefault(this, 3, "") as string;
        }
        set region(value: string) {
            pb_1.Message.setOneofField(this, 3, this.#one_of_decls[1], value);
        }
        get has_region() {
            return pb_1.Message.getField(this, 3) != null;
        }
        get _account() {
            const cases: {
                [index: number]: "none" | "account";
            } = {
                0: "none",
                2: "account"
            };
            return cases[pb_1.Message.computeOneofCase(this, [2])];
        }
        get _region() {
            const cases: {
                [index: number]: "none" | "region";
            } = {
                0: "none",
                3: "region"
            };
            return cases[pb_1.Message.computeOneofCase(this, [3])];
        }
        static fromObject(data: {
            clusterName?: string;
            account?: string;
            region?: string;
        }): BuildClusterRequest {
            const message = new BuildClusterRequest({});
            if (data.clusterName != null) {
                message.clusterName = data.clusterName;
            }
            if (data.account != null) {
                message.account = data.account;
            }
            if (data.region != null) {
                message.region = data.region;
            }
            return message;
        }
        toObject() {
            const data: {
                clusterName?: string;
                account?: string;
                region?: string;
            } = {};
            if (this.clusterName != null) {
                data.clusterName = this.clusterName;
            }
            if (this.account != null) {
                data.account = this.account;
            }
            if (this.region != null) {
                data.region = this.region;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.clusterName.length)
                writer.writeString(1, this.clusterName);
            if (this.has_account)
                writer.writeString(2, this.account);
            if (this.has_region)
                writer.writeString(3, this.region);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): BuildClusterRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new BuildClusterRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.clusterName = reader.readString();
                        break;
                    case 2:
                        message.account = reader.readString();
                        break;
                    case 3:
                        message.region = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): BuildClusterRequest {
            return BuildClusterRequest.deserialize(bytes);
        }
    }
    interface GrpcUnaryServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
    }
    interface GrpcStreamServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
        (message: P, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
    }
    interface GrpWritableServiceInterface<P, R> {
        (metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
    }
    interface GrpcChunkServiceInterface<P, R> {
        (metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
        (options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
    }
    interface GrpcPromiseServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): Promise<R>;
        (message: P, options?: grpc_1.CallOptions): Promise<R>;
    }
    export abstract class UnimplementedClusterServiceService {
        static definition = {
            CreateCluster: {
                path: "/proto_test.ClusterService/CreateCluster",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: CreateClusterRequest) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => CreateClusterRequest.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: APIResponse) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => APIResponse.deserialize(new Uint8Array(bytes))
            },
            AddTeams: {
                path: "/proto_test.ClusterService/AddTeams",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: AddTeamsRequest) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => AddTeamsRequest.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: APIResponse) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => APIResponse.deserialize(new Uint8Array(bytes))
            },
            AddClusterProvider: {
                path: "/proto_test.ClusterService/AddClusterProvider",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: AddClusterProviderRequest) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => AddClusterProviderRequest.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: APIResponse) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => APIResponse.deserialize(new Uint8Array(bytes))
            },
            AddResourceProvider: {
                path: "/proto_test.ClusterService/AddResourceProvider",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: AddResourceProviderRequest) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => AddResourceProviderRequest.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: APIResponse) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => APIResponse.deserialize(new Uint8Array(bytes))
            },
            BuildCluster: {
                path: "/proto_test.ClusterService/BuildCluster",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: BuildClusterRequest) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => BuildClusterRequest.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: APIResponse) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => APIResponse.deserialize(new Uint8Array(bytes))
            }
        };
        [method: string]: grpc_1.UntypedHandleCall;
        abstract CreateCluster(call: grpc_1.ServerUnaryCall<CreateClusterRequest, APIResponse>, callback: grpc_1.sendUnaryData<APIResponse>): void;
        abstract AddTeams(call: grpc_1.ServerUnaryCall<AddTeamsRequest, APIResponse>, callback: grpc_1.sendUnaryData<APIResponse>): void;
        abstract AddClusterProvider(call: grpc_1.ServerUnaryCall<AddClusterProviderRequest, APIResponse>, callback: grpc_1.sendUnaryData<APIResponse>): void;
        abstract AddResourceProvider(call: grpc_1.ServerUnaryCall<AddResourceProviderRequest, APIResponse>, callback: grpc_1.sendUnaryData<APIResponse>): void;
        abstract BuildCluster(call: grpc_1.ServerUnaryCall<BuildClusterRequest, APIResponse>, callback: grpc_1.sendUnaryData<APIResponse>): void;
    }
    export class ClusterServiceClient extends grpc_1.makeGenericClientConstructor(UnimplementedClusterServiceService.definition, "ClusterService", {}) {
        constructor(address: string, credentials: grpc_1.ChannelCredentials, options?: Partial<grpc_1.ChannelOptions>) {
            super(address, credentials, options);
        }
        CreateCluster: GrpcUnaryServiceInterface<CreateClusterRequest, APIResponse> = (message: CreateClusterRequest, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, options?: grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, callback?: grpc_1.requestCallback<APIResponse>): grpc_1.ClientUnaryCall => {
            return super.CreateCluster(message, metadata, options, callback);
        };
        AddTeams: GrpcUnaryServiceInterface<AddTeamsRequest, APIResponse> = (message: AddTeamsRequest, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, options?: grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, callback?: grpc_1.requestCallback<APIResponse>): grpc_1.ClientUnaryCall => {
            return super.AddTeams(message, metadata, options, callback);
        };
        AddClusterProvider: GrpcUnaryServiceInterface<AddClusterProviderRequest, APIResponse> = (message: AddClusterProviderRequest, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, options?: grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, callback?: grpc_1.requestCallback<APIResponse>): grpc_1.ClientUnaryCall => {
            return super.AddClusterProvider(message, metadata, options, callback);
        };
        AddResourceProvider: GrpcUnaryServiceInterface<AddResourceProviderRequest, APIResponse> = (message: AddResourceProviderRequest, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, options?: grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, callback?: grpc_1.requestCallback<APIResponse>): grpc_1.ClientUnaryCall => {
            return super.AddResourceProvider(message, metadata, options, callback);
        };
        BuildCluster: GrpcUnaryServiceInterface<BuildClusterRequest, APIResponse> = (message: BuildClusterRequest, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, options?: grpc_1.CallOptions | grpc_1.requestCallback<APIResponse>, callback?: grpc_1.requestCallback<APIResponse>): grpc_1.ClientUnaryCall => {
            return super.BuildCluster(message, metadata, options, callback);
        };
    }
}
