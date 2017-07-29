// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var computations_computations_pb = require('../computations/computations_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');
var scenes_scenes_pb = require('../scenes/scenes_pb.js');

function serialize_computations_ComputationStatus(arg) {
  if (!(arg instanceof computations_computations_pb.ComputationStatus)) {
    throw new Error('Expected argument of type computations.ComputationStatus');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computations_ComputationStatus(buffer_arg) {
  return computations_computations_pb.ComputationStatus.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_computations_CreateComputationRequest(arg) {
  if (!(arg instanceof computations_computations_pb.CreateComputationRequest)) {
    throw new Error('Expected argument of type computations.CreateComputationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computations_CreateComputationRequest(buffer_arg) {
  return computations_computations_pb.CreateComputationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_computations_GetComputationStatusRequest(arg) {
  if (!(arg instanceof computations_computations_pb.GetComputationStatusRequest)) {
    throw new Error('Expected argument of type computations.GetComputationStatusRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computations_GetComputationStatusRequest(buffer_arg) {
  return computations_computations_pb.GetComputationStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputationsService = exports.ComputationsService = {
  createComputation: {
    path: '/computations.Computations/CreateComputation',
    requestStream: false,
    responseStream: false,
    requestType: computations_computations_pb.CreateComputationRequest,
    responseType: computations_computations_pb.ComputationStatus,
    requestSerialize: serialize_computations_CreateComputationRequest,
    requestDeserialize: deserialize_computations_CreateComputationRequest,
    responseSerialize: serialize_computations_ComputationStatus,
    responseDeserialize: deserialize_computations_ComputationStatus,
  },
  getComputationStatus: {
    path: '/computations.Computations/GetComputationStatus',
    requestStream: false,
    responseStream: false,
    requestType: computations_computations_pb.GetComputationStatusRequest,
    responseType: computations_computations_pb.ComputationStatus,
    requestSerialize: serialize_computations_GetComputationStatusRequest,
    requestDeserialize: deserialize_computations_GetComputationStatusRequest,
    responseSerialize: serialize_computations_ComputationStatus,
    responseDeserialize: deserialize_computations_ComputationStatus,
  },
};

exports.ComputationsClient = grpc.makeGenericClientConstructor(ComputationsService);
