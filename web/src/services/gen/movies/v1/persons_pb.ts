// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file movies/v1/persons.proto (package movies.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";
import { PaginationRequestParams, PaginationResponseMetadata } from "../../shared/v1/pagination_pb.js";

/**
 * @generated from message movies.v1.Person
 */
export class Person extends Message<Person> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string slug = 3;
   */
  slug = "";

  /**
   * @generated from field: repeated string nicknames = 4;
   */
  nicknames: string[] = [];

  /**
   * @generated from field: repeated string occupations = 5;
   */
  occupations: string[] = [];

  /**
   * @generated from field: string profile_picture_url = 6;
   */
  profilePictureUrl = "";

  /**
   * @generated from field: optional google.protobuf.Timestamp dob = 7;
   */
  dob?: Timestamp;

  /**
   * @generated from field: string about = 8;
   */
  about = "";

  /**
   * @generated from field: bool is_deleted = 9;
   */
  isDeleted = false;

  constructor(data?: PartialMessage<Person>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.Person";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "slug", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "nicknames", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "occupations", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "profile_picture_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "dob", kind: "message", T: Timestamp, opt: true },
    { no: 8, name: "about", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "is_deleted", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Person {
    return new Person().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Person {
    return new Person().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Person {
    return new Person().fromJsonString(jsonString, options);
  }

  static equals(a: Person | PlainMessage<Person> | undefined, b: Person | PlainMessage<Person> | undefined): boolean {
    return proto3.util.equals(Person, a, b);
  }
}

/**
 * @generated from message movies.v1.GetPersonRequest
 */
export class GetPersonRequest extends Message<GetPersonRequest> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  constructor(data?: PartialMessage<GetPersonRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.GetPersonRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonRequest {
    return new GetPersonRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonRequest {
    return new GetPersonRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonRequest {
    return new GetPersonRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetPersonRequest | PlainMessage<GetPersonRequest> | undefined, b: GetPersonRequest | PlainMessage<GetPersonRequest> | undefined): boolean {
    return proto3.util.equals(GetPersonRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.GetPersonResponse
 */
export class GetPersonResponse extends Message<GetPersonResponse> {
  /**
   * @generated from field: movies.v1.Person person = 1;
   */
  person?: Person;

  constructor(data?: PartialMessage<GetPersonResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.GetPersonResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "person", kind: "message", T: Person },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPersonResponse {
    return new GetPersonResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPersonResponse {
    return new GetPersonResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPersonResponse {
    return new GetPersonResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetPersonResponse | PlainMessage<GetPersonResponse> | undefined, b: GetPersonResponse | PlainMessage<GetPersonResponse> | undefined): boolean {
    return proto3.util.equals(GetPersonResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.CreatePersonRequest
 */
export class CreatePersonRequest extends Message<CreatePersonRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: repeated string nicknames = 2;
   */
  nicknames: string[] = [];

  /**
   * @generated from field: repeated string occupations = 3;
   */
  occupations: string[] = [];

  /**
   * @generated from field: optional string profile_picture_key = 4;
   */
  profilePictureKey?: string;

  /**
   * @generated from field: optional google.protobuf.Timestamp dob = 5;
   */
  dob?: Timestamp;

  /**
   * @generated from field: string about = 6;
   */
  about = "";

  constructor(data?: PartialMessage<CreatePersonRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.CreatePersonRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "nicknames", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "occupations", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "profile_picture_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "dob", kind: "message", T: Timestamp, opt: true },
    { no: 6, name: "about", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreatePersonRequest {
    return new CreatePersonRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreatePersonRequest {
    return new CreatePersonRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreatePersonRequest {
    return new CreatePersonRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreatePersonRequest | PlainMessage<CreatePersonRequest> | undefined, b: CreatePersonRequest | PlainMessage<CreatePersonRequest> | undefined): boolean {
    return proto3.util.equals(CreatePersonRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.CreatePersonResponse
 */
export class CreatePersonResponse extends Message<CreatePersonResponse> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  constructor(data?: PartialMessage<CreatePersonResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.CreatePersonResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreatePersonResponse {
    return new CreatePersonResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreatePersonResponse {
    return new CreatePersonResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreatePersonResponse {
    return new CreatePersonResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreatePersonResponse | PlainMessage<CreatePersonResponse> | undefined, b: CreatePersonResponse | PlainMessage<CreatePersonResponse> | undefined): boolean {
    return proto3.util.equals(CreatePersonResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonRequest
 */
export class UpdatePersonRequest extends Message<UpdatePersonRequest> {
  /**
   * @generated from field: optional string name = 1;
   */
  name?: string;

  /**
   * @generated from field: optional string profile_picture_key = 4;
   */
  profilePictureKey?: string;

  /**
   * @generated from field: optional google.protobuf.Timestamp dob = 5;
   */
  dob?: Timestamp;

  constructor(data?: PartialMessage<UpdatePersonRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "profile_picture_key", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "dob", kind: "message", T: Timestamp, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonRequest {
    return new UpdatePersonRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonRequest {
    return new UpdatePersonRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonRequest {
    return new UpdatePersonRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonRequest | PlainMessage<UpdatePersonRequest> | undefined, b: UpdatePersonRequest | PlainMessage<UpdatePersonRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePersonRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonResponse
 */
export class UpdatePersonResponse extends Message<UpdatePersonResponse> {
  /**
   * @generated from field: movies.v1.Person person = 1;
   */
  person?: Person;

  constructor(data?: PartialMessage<UpdatePersonResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "person", kind: "message", T: Person },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonResponse {
    return new UpdatePersonResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonResponse {
    return new UpdatePersonResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonResponse {
    return new UpdatePersonResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonResponse | PlainMessage<UpdatePersonResponse> | undefined, b: UpdatePersonResponse | PlainMessage<UpdatePersonResponse> | undefined): boolean {
    return proto3.util.equals(UpdatePersonResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonAddOccupationRequest
 */
export class UpdatePersonAddOccupationRequest extends Message<UpdatePersonAddOccupationRequest> {
  /**
   * @generated from field: repeated string occupations = 1;
   */
  occupations: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonAddOccupationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonAddOccupationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "occupations", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonAddOccupationRequest {
    return new UpdatePersonAddOccupationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonAddOccupationRequest {
    return new UpdatePersonAddOccupationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonAddOccupationRequest {
    return new UpdatePersonAddOccupationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonAddOccupationRequest | PlainMessage<UpdatePersonAddOccupationRequest> | undefined, b: UpdatePersonAddOccupationRequest | PlainMessage<UpdatePersonAddOccupationRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePersonAddOccupationRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonAddOccupationResponse
 */
export class UpdatePersonAddOccupationResponse extends Message<UpdatePersonAddOccupationResponse> {
  /**
   * @generated from field: repeated string occupations = 1;
   */
  occupations: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonAddOccupationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonAddOccupationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "occupations", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonAddOccupationResponse {
    return new UpdatePersonAddOccupationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonAddOccupationResponse {
    return new UpdatePersonAddOccupationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonAddOccupationResponse {
    return new UpdatePersonAddOccupationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonAddOccupationResponse | PlainMessage<UpdatePersonAddOccupationResponse> | undefined, b: UpdatePersonAddOccupationResponse | PlainMessage<UpdatePersonAddOccupationResponse> | undefined): boolean {
    return proto3.util.equals(UpdatePersonAddOccupationResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonRemoveOccupationRequest
 */
export class UpdatePersonRemoveOccupationRequest extends Message<UpdatePersonRemoveOccupationRequest> {
  /**
   * @generated from field: repeated string occupations = 1;
   */
  occupations: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonRemoveOccupationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonRemoveOccupationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "occupations", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonRemoveOccupationRequest {
    return new UpdatePersonRemoveOccupationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonRemoveOccupationRequest {
    return new UpdatePersonRemoveOccupationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonRemoveOccupationRequest {
    return new UpdatePersonRemoveOccupationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonRemoveOccupationRequest | PlainMessage<UpdatePersonRemoveOccupationRequest> | undefined, b: UpdatePersonRemoveOccupationRequest | PlainMessage<UpdatePersonRemoveOccupationRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePersonRemoveOccupationRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonRemoveOccupationResponse
 */
export class UpdatePersonRemoveOccupationResponse extends Message<UpdatePersonRemoveOccupationResponse> {
  /**
   * @generated from field: repeated string occupations = 1;
   */
  occupations: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonRemoveOccupationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonRemoveOccupationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "occupations", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonRemoveOccupationResponse {
    return new UpdatePersonRemoveOccupationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonRemoveOccupationResponse {
    return new UpdatePersonRemoveOccupationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonRemoveOccupationResponse {
    return new UpdatePersonRemoveOccupationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonRemoveOccupationResponse | PlainMessage<UpdatePersonRemoveOccupationResponse> | undefined, b: UpdatePersonRemoveOccupationResponse | PlainMessage<UpdatePersonRemoveOccupationResponse> | undefined): boolean {
    return proto3.util.equals(UpdatePersonRemoveOccupationResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonAddNicknameRequest
 */
export class UpdatePersonAddNicknameRequest extends Message<UpdatePersonAddNicknameRequest> {
  /**
   * @generated from field: repeated string nicknames = 1;
   */
  nicknames: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonAddNicknameRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonAddNicknameRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nicknames", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonAddNicknameRequest {
    return new UpdatePersonAddNicknameRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonAddNicknameRequest {
    return new UpdatePersonAddNicknameRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonAddNicknameRequest {
    return new UpdatePersonAddNicknameRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonAddNicknameRequest | PlainMessage<UpdatePersonAddNicknameRequest> | undefined, b: UpdatePersonAddNicknameRequest | PlainMessage<UpdatePersonAddNicknameRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePersonAddNicknameRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonAddNicknameResponse
 */
export class UpdatePersonAddNicknameResponse extends Message<UpdatePersonAddNicknameResponse> {
  /**
   * @generated from field: repeated string nicknames = 1;
   */
  nicknames: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonAddNicknameResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonAddNicknameResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nicknames", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonAddNicknameResponse {
    return new UpdatePersonAddNicknameResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonAddNicknameResponse {
    return new UpdatePersonAddNicknameResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonAddNicknameResponse {
    return new UpdatePersonAddNicknameResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonAddNicknameResponse | PlainMessage<UpdatePersonAddNicknameResponse> | undefined, b: UpdatePersonAddNicknameResponse | PlainMessage<UpdatePersonAddNicknameResponse> | undefined): boolean {
    return proto3.util.equals(UpdatePersonAddNicknameResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonRemoveNicknameRequest
 */
export class UpdatePersonRemoveNicknameRequest extends Message<UpdatePersonRemoveNicknameRequest> {
  /**
   * @generated from field: repeated string nicknames = 1;
   */
  nicknames: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonRemoveNicknameRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonRemoveNicknameRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nicknames", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonRemoveNicknameRequest {
    return new UpdatePersonRemoveNicknameRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonRemoveNicknameRequest {
    return new UpdatePersonRemoveNicknameRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonRemoveNicknameRequest {
    return new UpdatePersonRemoveNicknameRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonRemoveNicknameRequest | PlainMessage<UpdatePersonRemoveNicknameRequest> | undefined, b: UpdatePersonRemoveNicknameRequest | PlainMessage<UpdatePersonRemoveNicknameRequest> | undefined): boolean {
    return proto3.util.equals(UpdatePersonRemoveNicknameRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.UpdatePersonRemoveNicknameResponse
 */
export class UpdatePersonRemoveNicknameResponse extends Message<UpdatePersonRemoveNicknameResponse> {
  /**
   * @generated from field: repeated string nicknames = 1;
   */
  nicknames: string[] = [];

  constructor(data?: PartialMessage<UpdatePersonRemoveNicknameResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.UpdatePersonRemoveNicknameResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nicknames", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdatePersonRemoveNicknameResponse {
    return new UpdatePersonRemoveNicknameResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdatePersonRemoveNicknameResponse {
    return new UpdatePersonRemoveNicknameResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdatePersonRemoveNicknameResponse {
    return new UpdatePersonRemoveNicknameResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdatePersonRemoveNicknameResponse | PlainMessage<UpdatePersonRemoveNicknameResponse> | undefined, b: UpdatePersonRemoveNicknameResponse | PlainMessage<UpdatePersonRemoveNicknameResponse> | undefined): boolean {
    return proto3.util.equals(UpdatePersonRemoveNicknameResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.DeletePersonRequest
 */
export class DeletePersonRequest extends Message<DeletePersonRequest> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  constructor(data?: PartialMessage<DeletePersonRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.DeletePersonRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeletePersonRequest {
    return new DeletePersonRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeletePersonRequest {
    return new DeletePersonRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeletePersonRequest {
    return new DeletePersonRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeletePersonRequest | PlainMessage<DeletePersonRequest> | undefined, b: DeletePersonRequest | PlainMessage<DeletePersonRequest> | undefined): boolean {
    return proto3.util.equals(DeletePersonRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.DeletePersonResponse
 */
export class DeletePersonResponse extends Message<DeletePersonResponse> {
  constructor(data?: PartialMessage<DeletePersonResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.DeletePersonResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeletePersonResponse {
    return new DeletePersonResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeletePersonResponse {
    return new DeletePersonResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeletePersonResponse {
    return new DeletePersonResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeletePersonResponse | PlainMessage<DeletePersonResponse> | undefined, b: DeletePersonResponse | PlainMessage<DeletePersonResponse> | undefined): boolean {
    return proto3.util.equals(DeletePersonResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.ListPersonsRequest
 */
export class ListPersonsRequest extends Message<ListPersonsRequest> {
  /**
   * @generated from field: shared.v1.PaginationRequestParams pagination = 1;
   */
  pagination?: PaginationRequestParams;

  constructor(data?: PartialMessage<ListPersonsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.ListPersonsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PaginationRequestParams },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListPersonsRequest {
    return new ListPersonsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListPersonsRequest {
    return new ListPersonsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListPersonsRequest {
    return new ListPersonsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListPersonsRequest | PlainMessage<ListPersonsRequest> | undefined, b: ListPersonsRequest | PlainMessage<ListPersonsRequest> | undefined): boolean {
    return proto3.util.equals(ListPersonsRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.ListPersonsResponse
 */
export class ListPersonsResponse extends Message<ListPersonsResponse> {
  /**
   * @generated from field: shared.v1.PaginationResponseMetadata metadata = 1;
   */
  metadata?: PaginationResponseMetadata;

  /**
   * @generated from field: repeated movies.v1.Person persons = 2;
   */
  persons: Person[] = [];

  constructor(data?: PartialMessage<ListPersonsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.ListPersonsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: PaginationResponseMetadata },
    { no: 2, name: "persons", kind: "message", T: Person, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListPersonsResponse {
    return new ListPersonsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListPersonsResponse {
    return new ListPersonsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListPersonsResponse {
    return new ListPersonsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListPersonsResponse | PlainMessage<ListPersonsResponse> | undefined, b: ListPersonsResponse | PlainMessage<ListPersonsResponse> | undefined): boolean {
    return proto3.util.equals(ListPersonsResponse, a, b);
  }
}

/**
 * @generated from message movies.v1.SearchPersonRequest
 */
export class SearchPersonRequest extends Message<SearchPersonRequest> {
  /**
   * @generated from field: shared.v1.PaginationRequestParams pagination = 1;
   */
  pagination?: PaginationRequestParams;

  /**
   * @generated from field: string query = 2;
   */
  query = "";

  constructor(data?: PartialMessage<SearchPersonRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.SearchPersonRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PaginationRequestParams },
    { no: 2, name: "query", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchPersonRequest {
    return new SearchPersonRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchPersonRequest {
    return new SearchPersonRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchPersonRequest {
    return new SearchPersonRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SearchPersonRequest | PlainMessage<SearchPersonRequest> | undefined, b: SearchPersonRequest | PlainMessage<SearchPersonRequest> | undefined): boolean {
    return proto3.util.equals(SearchPersonRequest, a, b);
  }
}

/**
 * @generated from message movies.v1.SearchPersonResponse
 */
export class SearchPersonResponse extends Message<SearchPersonResponse> {
  /**
   * @generated from field: shared.v1.PaginationResponseMetadata metadata = 1;
   */
  metadata?: PaginationResponseMetadata;

  /**
   * @generated from field: repeated movies.v1.Person persons = 2;
   */
  persons: Person[] = [];

  constructor(data?: PartialMessage<SearchPersonResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "movies.v1.SearchPersonResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: PaginationResponseMetadata },
    { no: 2, name: "persons", kind: "message", T: Person, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchPersonResponse {
    return new SearchPersonResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchPersonResponse {
    return new SearchPersonResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchPersonResponse {
    return new SearchPersonResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SearchPersonResponse | PlainMessage<SearchPersonResponse> | undefined, b: SearchPersonResponse | PlainMessage<SearchPersonResponse> | undefined): boolean {
    return proto3.util.equals(SearchPersonResponse, a, b);
  }
}
