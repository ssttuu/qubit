// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var parameters_parameters_pb = require('../parameters/parameters_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_CreateParameterRootRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.CreateParameterRootRequest)) {
    throw new Error('Expected argument of type parameters.CreateParameterRootRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_CreateParameterRootRequest(buffer_arg) {
  return parameters_parameters_pb.CreateParameterRootRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_DeleteParameterRootRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.DeleteParameterRootRequest)) {
    throw new Error('Expected argument of type parameters.DeleteParameterRootRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_DeleteParameterRootRequest(buffer_arg) {
  return parameters_parameters_pb.DeleteParameterRootRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_GetParameterRootRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.GetParameterRootRequest)) {
    throw new Error('Expected argument of type parameters.GetParameterRootRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_GetParameterRootRequest(buffer_arg) {
  return parameters_parameters_pb.GetParameterRootRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_ParameterRoot(arg) {
  if (!(arg instanceof parameters_parameters_pb.ParameterRoot)) {
    throw new Error('Expected argument of type parameters.ParameterRoot');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_ParameterRoot(buffer_arg) {
  return parameters_parameters_pb.ParameterRoot.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_UpdateParameterRootRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.UpdateParameterRootRequest)) {
    throw new Error('Expected argument of type parameters.UpdateParameterRootRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_UpdateParameterRootRequest(buffer_arg) {
  return parameters_parameters_pb.UpdateParameterRootRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ParametersService = exports.ParametersService = {
  get: {
    path: '/parameters.Parameters/Get',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.GetParameterRootRequest,
    responseType: parameters_parameters_pb.ParameterRoot,
    requestSerialize: serialize_parameters_GetParameterRootRequest,
    requestDeserialize: deserialize_parameters_GetParameterRootRequest,
    responseSerialize: serialize_parameters_ParameterRoot,
    responseDeserialize: deserialize_parameters_ParameterRoot,
  },
  create: {
    path: '/parameters.Parameters/Create',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.CreateParameterRootRequest,
    responseType: parameters_parameters_pb.ParameterRoot,
    requestSerialize: serialize_parameters_CreateParameterRootRequest,
    requestDeserialize: deserialize_parameters_CreateParameterRootRequest,
    responseSerialize: serialize_parameters_ParameterRoot,
    responseDeserialize: deserialize_parameters_ParameterRoot,
  },
  update: {
    path: '/parameters.Parameters/Update',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.UpdateParameterRootRequest,
    responseType: parameters_parameters_pb.ParameterRoot,
    requestSerialize: serialize_parameters_UpdateParameterRootRequest,
    requestDeserialize: deserialize_parameters_UpdateParameterRootRequest,
    responseSerialize: serialize_parameters_ParameterRoot,
    responseDeserialize: deserialize_parameters_ParameterRoot,
  },
  delete: {
    path: '/parameters.Parameters/Delete',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.DeleteParameterRootRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_parameters_DeleteParameterRootRequest,
    requestDeserialize: deserialize_parameters_DeleteParameterRootRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.ParametersClient = grpc.makeGenericClientConstructor(ParametersService);
