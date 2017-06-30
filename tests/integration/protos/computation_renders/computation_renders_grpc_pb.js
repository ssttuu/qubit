// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var computation_renders_computation_renders_pb = require('../computation_renders/computation_renders_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');

function serialize_computation_renders_ComputationRenderRequest(arg) {
  if (!(arg instanceof computation_renders_computation_renders_pb.ComputationRenderRequest)) {
    throw new Error('Expected argument of type computation_renders.ComputationRenderRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_renders_ComputationRenderRequest(buffer_arg) {
  return computation_renders_computation_renders_pb.ComputationRenderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_computation_renders_ComputationRenderResponse(arg) {
  if (!(arg instanceof computation_renders_computation_renders_pb.ComputationRenderResponse)) {
    throw new Error('Expected argument of type computation_renders.ComputationRenderResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_renders_ComputationRenderResponse(buffer_arg) {
  return computation_renders_computation_renders_pb.ComputationRenderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputationRendersService = exports.ComputationRendersService = {
  render: {
    path: '/computation_renders.ComputationRenders/Render',
    requestStream: false,
    responseStream: false,
    requestType: computation_renders_computation_renders_pb.ComputationRenderRequest,
    responseType: computation_renders_computation_renders_pb.ComputationRenderResponse,
    requestSerialize: serialize_computation_renders_ComputationRenderRequest,
    requestDeserialize: deserialize_computation_renders_ComputationRenderRequest,
    responseSerialize: serialize_computation_renders_ComputationRenderResponse,
    responseDeserialize: deserialize_computation_renders_ComputationRenderResponse,
  },
};

exports.ComputationRendersClient = grpc.makeGenericClientConstructor(ComputationRendersService);
