// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var computation_statuses_computation_statuses_pb = require('../computation_statuses/computation_statuses_pb.js');

function serialize_computation_statuses_ComputationStatus(arg) {
  if (!(arg instanceof computation_statuses_computation_statuses_pb.ComputationStatus)) {
    throw new Error('Expected argument of type computation_statuses.ComputationStatus');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_statuses_ComputationStatus(buffer_arg) {
  return computation_statuses_computation_statuses_pb.ComputationStatus.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_computation_statuses_GetComputationStatusRequest(arg) {
  if (!(arg instanceof computation_statuses_computation_statuses_pb.GetComputationStatusRequest)) {
    throw new Error('Expected argument of type computation_statuses.GetComputationStatusRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_computation_statuses_GetComputationStatusRequest(buffer_arg) {
  return computation_statuses_computation_statuses_pb.GetComputationStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputationStatusesService = exports.ComputationStatusesService = {
  getComputationStatus: {
    path: '/computation_statuses.ComputationStatuses/GetComputationStatus',
    requestStream: false,
    responseStream: false,
    requestType: computation_statuses_computation_statuses_pb.GetComputationStatusRequest,
    responseType: computation_statuses_computation_statuses_pb.ComputationStatus,
    requestSerialize: serialize_computation_statuses_GetComputationStatusRequest,
    requestDeserialize: deserialize_computation_statuses_GetComputationStatusRequest,
    responseSerialize: serialize_computation_statuses_ComputationStatus,
    responseDeserialize: deserialize_computation_statuses_ComputationStatus,
  },
};

exports.ComputationStatusesClient = grpc.makeGenericClientConstructor(ComputationStatusesService);
