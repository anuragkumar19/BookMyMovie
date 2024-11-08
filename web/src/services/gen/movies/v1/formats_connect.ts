// @generated by protoc-gen-connect-es v1.5.0 with parameter "target=ts"
// @generated from file movies/v1/formats.proto (package movies.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateFormatRequest, CreateFormatResponse, DeleteFormatRequest, DeleteFormatResponse, GetFormatRequest, GetFormatResponse, GetFormatsRequest, GetFormatsResponse, UpdateFormatRequest, UpdateFormatResponse } from "./formats_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service movies.v1.MoviesFormatsService
 */
export const MoviesFormatsService = {
  typeName: "movies.v1.MoviesFormatsService",
  methods: {
    /**
     * @generated from rpc movies.v1.MoviesFormatsService.GetFormat
     */
    getFormat: {
      name: "GetFormat",
      I: GetFormatRequest,
      O: GetFormatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc movies.v1.MoviesFormatsService.GetFormats
     */
    getFormats: {
      name: "GetFormats",
      I: GetFormatsRequest,
      O: GetFormatsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc movies.v1.MoviesFormatsService.CreateFormat
     */
    createFormat: {
      name: "CreateFormat",
      I: CreateFormatRequest,
      O: CreateFormatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc movies.v1.MoviesFormatsService.UpdateFormat
     */
    updateFormat: {
      name: "UpdateFormat",
      I: UpdateFormatRequest,
      O: UpdateFormatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc movies.v1.MoviesFormatsService.DeleteFormat
     */
    deleteFormat: {
      name: "DeleteFormat",
      I: DeleteFormatRequest,
      O: DeleteFormatResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

