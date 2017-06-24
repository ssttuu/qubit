// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var computation_computation_pb = require('../computation/computation_pb.js');
var computation_status_computation_status_pb = require('../computation_status/computation_status_pb.js');
var operators_operators_pb = require('../operators/operators_pb.js');

function serialize_computation_ComputationStatus(arg) {
  if (!(arg instanceof computation_status_computation_status_pb.ComputationStatus)) {
    throw new Error('Expected argument of type computation.ComputationStatus');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_ComputationStatus(buffer_arg) {
  return computation_status_computation_status_pb.ComputationStatus.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_computation_CreateComputationRequest(arg) {
  if (!(arg instanceof computation_computation_pb.CreateComputationRequest)) {
    throw new Error('Expected argument of type computation.CreateComputationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_CreateComputationRequest(buffer_arg) {
  return computation_computation_pb.CreateComputationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputationsService = exports.ComputationsService = {
  createComputation: {
    path: '/computation.Computations/CreateComputation',
    requestStream: false,
    responseStream: false,
    requestType: computation_computation_pb.CreateComputationRequest,
    responseType: computation_status_computation_status_pb.ComputationStatus,
    requestSerialize: serialize_computation_CreateComputationRequest,
    requestDeserialize: deserialize_computation_CreateComputationRequest,
    responseSerialize: serialize_computation_ComputationStatus,
    responseDeserialize: deserialize_computation_ComputationStatus,
  },
};

exports.ComputationsClient = grpc.makeGenericClientConstructor(ComputationsService);
