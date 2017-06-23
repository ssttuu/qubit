// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var compute_compute_pb = require('../compute/compute_pb.js');
var operators_operators_pb = require('../operators/operators_pb.js');

function serialize_compute_ComputationStatusRequest(arg) {
  if (!(arg instanceof compute_compute_pb.ComputationStatusRequest)) {
    throw new Error('Expected argument of type compute.ComputationStatusRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_compute_ComputationStatusRequest(buffer_arg) {
  return compute_compute_pb.ComputationStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_compute_ComputationStatusResponse(arg) {
  if (!(arg instanceof compute_compute_pb.ComputationStatusResponse)) {
    throw new Error('Expected argument of type compute.ComputationStatusResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_compute_ComputationStatusResponse(buffer_arg) {
  return compute_compute_pb.ComputationStatusResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_compute_CreateComputationRequest(arg) {
  if (!(arg instanceof compute_compute_pb.CreateComputationRequest)) {
    throw new Error('Expected argument of type compute.CreateComputationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_compute_CreateComputationRequest(buffer_arg) {
  return compute_compute_pb.CreateComputationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputeService = exports.ComputeService = {
  createComputation: {
    path: '/compute.Compute/CreateComputation',
    requestStream: false,
    responseStream: false,
    requestType: compute_compute_pb.CreateComputationRequest,
    responseType: compute_compute_pb.ComputationStatusResponse,
    requestSerialize: serialize_compute_CreateComputationRequest,
    requestDeserialize: deserialize_compute_CreateComputationRequest,
    responseSerialize: serialize_compute_ComputationStatusResponse,
    responseDeserialize: deserialize_compute_ComputationStatusResponse,
  },
  getComputationStatus: {
    path: '/compute.Compute/GetComputationStatus',
    requestStream: false,
    responseStream: false,
    requestType: compute_compute_pb.ComputationStatusRequest,
    responseType: compute_compute_pb.ComputationStatusResponse,
    requestSerialize: serialize_compute_ComputationStatusRequest,
    requestDeserialize: deserialize_compute_ComputationStatusRequest,
    responseSerialize: serialize_compute_ComputationStatusResponse,
    responseDeserialize: deserialize_compute_ComputationStatusResponse,
  },
};

exports.ComputeClient = grpc.makeGenericClientConstructor(ComputeService);
