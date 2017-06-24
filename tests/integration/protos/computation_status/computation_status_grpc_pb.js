// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var computation_status_computation_status_pb = require('../computation_status/computation_status_pb.js');

function serialize_computation_ComputationStatus(arg) {
  if (!(arg instanceof computation_status_computation_status_pb.ComputationStatus)) {
    throw new Error('Expected argument of type computation.ComputationStatus');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_ComputationStatus(buffer_arg) {
  return computation_status_computation_status_pb.ComputationStatus.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_computation_GetComputationStatusRequest(arg) {
  if (!(arg instanceof computation_status_computation_status_pb.GetComputationStatusRequest)) {
    throw new Error('Expected argument of type computation.GetComputationStatusRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_GetComputationStatusRequest(buffer_arg) {
  return computation_status_computation_status_pb.GetComputationStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputationStatusesService = exports.ComputationStatusesService = {
  getComputationStatus: {
    path: '/computation.ComputationStatuses/GetComputationStatus',
    requestStream: false,
    responseStream: false,
    requestType: computation_status_computation_status_pb.GetComputationStatusRequest,
    responseType: computation_status_computation_status_pb.ComputationStatus,
    requestSerialize: serialize_computation_GetComputationStatusRequest,
    requestDeserialize: deserialize_computation_GetComputationStatusRequest,
    responseSerialize: serialize_computation_ComputationStatus,
    responseDeserialize: deserialize_computation_ComputationStatus,
  },
};

exports.ComputationStatusesClient = grpc.makeGenericClientConstructor(ComputationStatusesService);
