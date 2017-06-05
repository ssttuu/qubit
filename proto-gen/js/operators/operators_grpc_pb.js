// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var operators_operators_pb = require('../operators/operators_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');
var parameters_parameters_pb = require('../parameters/parameters_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_CreateOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.CreateOperatorRequest)) {
    throw new Error('Expected argument of type operators.CreateOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_CreateOperatorRequest(buffer_arg) {
  return operators_operators_pb.CreateOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_DeleteOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.DeleteOperatorRequest)) {
    throw new Error('Expected argument of type operators.DeleteOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_DeleteOperatorRequest(buffer_arg) {
  return operators_operators_pb.DeleteOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_GetOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.GetOperatorRequest)) {
    throw new Error('Expected argument of type operators.GetOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_GetOperatorRequest(buffer_arg) {
  return operators_operators_pb.GetOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_ListOperatorsRequest(arg) {
  if (!(arg instanceof operators_operators_pb.ListOperatorsRequest)) {
    throw new Error('Expected argument of type operators.ListOperatorsRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_ListOperatorsRequest(buffer_arg) {
  return operators_operators_pb.ListOperatorsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_ListOperatorsResponse(arg) {
  if (!(arg instanceof operators_operators_pb.ListOperatorsResponse)) {
    throw new Error('Expected argument of type operators.ListOperatorsResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_ListOperatorsResponse(buffer_arg) {
  return operators_operators_pb.ListOperatorsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_Operator(arg) {
  if (!(arg instanceof operators_operators_pb.Operator)) {
    throw new Error('Expected argument of type operators.Operator');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_Operator(buffer_arg) {
  return operators_operators_pb.Operator.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_RenderOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.RenderOperatorRequest)) {
    throw new Error('Expected argument of type operators.RenderOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_RenderOperatorRequest(buffer_arg) {
  return operators_operators_pb.RenderOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_RenderOperatorResponse(arg) {
  if (!(arg instanceof operators_operators_pb.RenderOperatorResponse)) {
    throw new Error('Expected argument of type operators.RenderOperatorResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_RenderOperatorResponse(buffer_arg) {
  return operators_operators_pb.RenderOperatorResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_UpdateOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.UpdateOperatorRequest)) {
    throw new Error('Expected argument of type operators.UpdateOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_UpdateOperatorRequest(buffer_arg) {
  return operators_operators_pb.UpdateOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var OperatorsService = exports.OperatorsService = {
  list: {
    path: '/operators.Operators/List',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.ListOperatorsRequest,
    responseType: operators_operators_pb.ListOperatorsResponse,
    requestSerialize: serialize_operators_ListOperatorsRequest,
    requestDeserialize: deserialize_operators_ListOperatorsRequest,
    responseSerialize: serialize_operators_ListOperatorsResponse,
    responseDeserialize: deserialize_operators_ListOperatorsResponse,
  },
  get: {
    path: '/operators.Operators/Get',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.GetOperatorRequest,
    responseType: operators_operators_pb.Operator,
    requestSerialize: serialize_operators_GetOperatorRequest,
    requestDeserialize: deserialize_operators_GetOperatorRequest,
    responseSerialize: serialize_operators_Operator,
    responseDeserialize: deserialize_operators_Operator,
  },
  create: {
    path: '/operators.Operators/Create',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.CreateOperatorRequest,
    responseType: operators_operators_pb.Operator,
    requestSerialize: serialize_operators_CreateOperatorRequest,
    requestDeserialize: deserialize_operators_CreateOperatorRequest,
    responseSerialize: serialize_operators_Operator,
    responseDeserialize: deserialize_operators_Operator,
  },
  update: {
    path: '/operators.Operators/Update',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.UpdateOperatorRequest,
    responseType: operators_operators_pb.Operator,
    requestSerialize: serialize_operators_UpdateOperatorRequest,
    requestDeserialize: deserialize_operators_UpdateOperatorRequest,
    responseSerialize: serialize_operators_Operator,
    responseDeserialize: deserialize_operators_Operator,
  },
  delete: {
    path: '/operators.Operators/Delete',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.DeleteOperatorRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_operators_DeleteOperatorRequest,
    requestDeserialize: deserialize_operators_DeleteOperatorRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  render: {
    path: '/operators.Operators/Render',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.RenderOperatorRequest,
    responseType: operators_operators_pb.RenderOperatorResponse,
    requestSerialize: serialize_operators_RenderOperatorRequest,
    requestDeserialize: deserialize_operators_RenderOperatorRequest,
    responseSerialize: serialize_operators_RenderOperatorResponse,
    responseDeserialize: deserialize_operators_RenderOperatorResponse,
  },
};

exports.OperatorsClient = grpc.makeGenericClientConstructor(OperatorsService);
