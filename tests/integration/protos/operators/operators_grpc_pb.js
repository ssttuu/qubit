// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var operators_operators_pb = require('../operators/operators_pb.js');
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

function serialize_operators_ConnectOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.ConnectOperatorRequest)) {
    throw new Error('Expected argument of type operators.ConnectOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_ConnectOperatorRequest(buffer_arg) {
  return operators_operators_pb.ConnectOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_Connection(arg) {
  if (!(arg instanceof operators_operators_pb.Connection)) {
    throw new Error('Expected argument of type operators.Connection');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_Connection(buffer_arg) {
  return operators_operators_pb.Connection.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_operators_DisconnectOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.DisconnectOperatorRequest)) {
    throw new Error('Expected argument of type operators.DisconnectOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_DisconnectOperatorRequest(buffer_arg) {
  return operators_operators_pb.DisconnectOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_operators_RenameOperatorRequest(arg) {
  if (!(arg instanceof operators_operators_pb.RenameOperatorRequest)) {
    throw new Error('Expected argument of type operators.RenameOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_RenameOperatorRequest(buffer_arg) {
  return operators_operators_pb.RenameOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_SetExpressionRequest(arg) {
  if (!(arg instanceof operators_operators_pb.SetExpressionRequest)) {
    throw new Error('Expected argument of type operators.SetExpressionRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_SetExpressionRequest(buffer_arg) {
  return operators_operators_pb.SetExpressionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_SetKeyFrameRequest(arg) {
  if (!(arg instanceof operators_operators_pb.SetKeyFrameRequest)) {
    throw new Error('Expected argument of type operators.SetKeyFrameRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_SetKeyFrameRequest(buffer_arg) {
  return operators_operators_pb.SetKeyFrameRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_operators_SetValueRequest(arg) {
  if (!(arg instanceof operators_operators_pb.SetValueRequest)) {
    throw new Error('Expected argument of type operators.SetValueRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_operators_SetValueRequest(buffer_arg) {
  return operators_operators_pb.SetValueRequest.deserializeBinary(new Uint8Array(buffer_arg));
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
  rename: {
    path: '/operators.Operators/Rename',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.RenameOperatorRequest,
    responseType: operators_operators_pb.Operator,
    requestSerialize: serialize_operators_RenameOperatorRequest,
    requestDeserialize: deserialize_operators_RenameOperatorRequest,
    responseSerialize: serialize_operators_Operator,
    responseDeserialize: deserialize_operators_Operator,
  },
  // Connections API
  connect: {
    path: '/operators.Operators/Connect',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.ConnectOperatorRequest,
    responseType: operators_operators_pb.Connection,
    requestSerialize: serialize_operators_ConnectOperatorRequest,
    requestDeserialize: deserialize_operators_ConnectOperatorRequest,
    responseSerialize: serialize_operators_Connection,
    responseDeserialize: deserialize_operators_Connection,
  },
  disconnect: {
    path: '/operators.Operators/Disconnect',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.DisconnectOperatorRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_operators_DisconnectOperatorRequest,
    requestDeserialize: deserialize_operators_DisconnectOperatorRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Parameters API
  // TODO: does anything need to be returned or are status codes sufficient?
  // TODO: this will be a chatty API, ideally status codes are sufficient.
  setValue: {
    path: '/operators.Operators/SetValue',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.SetValueRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_operators_SetValueRequest,
    requestDeserialize: deserialize_operators_SetValueRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  setKeyFrame: {
    path: '/operators.Operators/SetKeyFrame',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.SetKeyFrameRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_operators_SetKeyFrameRequest,
    requestDeserialize: deserialize_operators_SetKeyFrameRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  setExpression: {
    path: '/operators.Operators/SetExpression',
    requestStream: false,
    responseStream: false,
    requestType: operators_operators_pb.SetExpressionRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_operators_SetExpressionRequest,
    requestDeserialize: deserialize_operators_SetExpressionRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.OperatorsClient = grpc.makeGenericClientConstructor(OperatorsService);
