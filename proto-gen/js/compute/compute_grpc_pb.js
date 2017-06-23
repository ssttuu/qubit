// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var compute_compute_pb = require('../compute/compute_pb.js');
var operators_operators_pb = require('../operators/operators_pb.js');

function serialize_compute_Computation(arg) {
  if (!(arg instanceof compute_compute_pb.Computation)) {
    throw new Error('Expected argument of type compute.Computation');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_compute_Computation(buffer_arg) {
  return compute_compute_pb.Computation.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_compute_GetComputationRequest(arg) {
  if (!(arg instanceof compute_compute_pb.GetComputationRequest)) {
    throw new Error('Expected argument of type compute.GetComputationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_compute_GetComputationRequest(buffer_arg) {
  return compute_compute_pb.GetComputationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputeService = exports.ComputeService = {
  createComputation: {
    path: '/compute.Compute/CreateComputation',
    requestStream: false,
    responseStream: false,
    requestType: compute_compute_pb.CreateComputationRequest,
    responseType: compute_compute_pb.Computation,
    requestSerialize: serialize_compute_CreateComputationRequest,
    requestDeserialize: deserialize_compute_CreateComputationRequest,
    responseSerialize: serialize_compute_Computation,
    responseDeserialize: deserialize_compute_Computation,
  },
  getComputation: {
    path: '/compute.Compute/GetComputation',
    requestStream: false,
    responseStream: false,
    requestType: compute_compute_pb.GetComputationRequest,
    responseType: compute_compute_pb.Computation,
    requestSerialize: serialize_compute_GetComputationRequest,
    requestDeserialize: deserialize_compute_GetComputationRequest,
    responseSerialize: serialize_compute_Computation,
    responseDeserialize: deserialize_compute_Computation,
  },
};

exports.ComputeClient = grpc.makeGenericClientConstructor(ComputeService);
