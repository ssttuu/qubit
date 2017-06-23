// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var parameters_parameters_pb = require('../parameters/parameters_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_CreateParameterRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.CreateParameterRequest)) {
    throw new Error('Expected argument of type parameters.CreateParameterRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_CreateParameterRequest(buffer_arg) {
  return parameters_parameters_pb.CreateParameterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_DeleteParameterRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.DeleteParameterRequest)) {
    throw new Error('Expected argument of type parameters.DeleteParameterRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_DeleteParameterRequest(buffer_arg) {
  return parameters_parameters_pb.DeleteParameterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_GetParameterRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.GetParameterRequest)) {
    throw new Error('Expected argument of type parameters.GetParameterRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_GetParameterRequest(buffer_arg) {
  return parameters_parameters_pb.GetParameterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_ListParametersRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.ListParametersRequest)) {
    throw new Error('Expected argument of type parameters.ListParametersRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_ListParametersRequest(buffer_arg) {
  return parameters_parameters_pb.ListParametersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_ListParametersResponse(arg) {
  if (!(arg instanceof parameters_parameters_pb.ListParametersResponse)) {
    throw new Error('Expected argument of type parameters.ListParametersResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_ListParametersResponse(buffer_arg) {
  return parameters_parameters_pb.ListParametersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_Parameter(arg) {
  if (!(arg instanceof parameters_parameters_pb.Parameter)) {
    throw new Error('Expected argument of type parameters.Parameter');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_Parameter(buffer_arg) {
  return parameters_parameters_pb.Parameter.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_RenderParameterRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.RenderParameterRequest)) {
    throw new Error('Expected argument of type parameters.RenderParameterRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_RenderParameterRequest(buffer_arg) {
  return parameters_parameters_pb.RenderParameterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_RenderParameterResponse(arg) {
  if (!(arg instanceof parameters_parameters_pb.RenderParameterResponse)) {
    throw new Error('Expected argument of type parameters.RenderParameterResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_RenderParameterResponse(buffer_arg) {
  return parameters_parameters_pb.RenderParameterResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_parameters_UpdateParameterRequest(arg) {
  if (!(arg instanceof parameters_parameters_pb.UpdateParameterRequest)) {
    throw new Error('Expected argument of type parameters.UpdateParameterRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_parameters_UpdateParameterRequest(buffer_arg) {
  return parameters_parameters_pb.UpdateParameterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ParametersService = exports.ParametersService = {
  list: {
    path: '/parameters.Parameters/List',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.ListParametersRequest,
    responseType: parameters_parameters_pb.ListParametersResponse,
    requestSerialize: serialize_parameters_ListParametersRequest,
    requestDeserialize: deserialize_parameters_ListParametersRequest,
    responseSerialize: serialize_parameters_ListParametersResponse,
    responseDeserialize: deserialize_parameters_ListParametersResponse,
  },
  get: {
    path: '/parameters.Parameters/Get',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.GetParameterRequest,
    responseType: parameters_parameters_pb.Parameter,
    requestSerialize: serialize_parameters_GetParameterRequest,
    requestDeserialize: deserialize_parameters_GetParameterRequest,
    responseSerialize: serialize_parameters_Parameter,
    responseDeserialize: deserialize_parameters_Parameter,
  },
  create: {
    path: '/parameters.Parameters/Create',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.CreateParameterRequest,
    responseType: parameters_parameters_pb.Parameter,
    requestSerialize: serialize_parameters_CreateParameterRequest,
    requestDeserialize: deserialize_parameters_CreateParameterRequest,
    responseSerialize: serialize_parameters_Parameter,
    responseDeserialize: deserialize_parameters_Parameter,
  },
  update: {
    path: '/parameters.Parameters/Update',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.UpdateParameterRequest,
    responseType: parameters_parameters_pb.Parameter,
    requestSerialize: serialize_parameters_UpdateParameterRequest,
    requestDeserialize: deserialize_parameters_UpdateParameterRequest,
    responseSerialize: serialize_parameters_Parameter,
    responseDeserialize: deserialize_parameters_Parameter,
  },
  delete: {
    path: '/parameters.Parameters/Delete',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.DeleteParameterRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_parameters_DeleteParameterRequest,
    requestDeserialize: deserialize_parameters_DeleteParameterRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  render: {
    path: '/parameters.Parameters/Render',
    requestStream: false,
    responseStream: false,
    requestType: parameters_parameters_pb.RenderParameterRequest,
    responseType: parameters_parameters_pb.RenderParameterResponse,
    requestSerialize: serialize_parameters_RenderParameterRequest,
    requestDeserialize: deserialize_parameters_RenderParameterRequest,
    responseSerialize: serialize_parameters_RenderParameterResponse,
    responseDeserialize: deserialize_parameters_RenderParameterResponse,
  },
};

exports.ParametersClient = grpc.makeGenericClientConstructor(ParametersService);
