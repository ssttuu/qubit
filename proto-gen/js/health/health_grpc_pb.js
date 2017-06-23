// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var health_health_pb = require('../health/health_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');

function serialize_health_HealthCheckRequest(arg) {
  if (!(arg instanceof health_health_pb.HealthCheckRequest)) {
    throw new Error('Expected argument of type health.HealthCheckRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_health_HealthCheckRequest(buffer_arg) {
  return health_health_pb.HealthCheckRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_health_HealthCheckResponse(arg) {
  if (!(arg instanceof health_health_pb.HealthCheckResponse)) {
    throw new Error('Expected argument of type health.HealthCheckResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_health_HealthCheckResponse(buffer_arg) {
  return health_health_pb.HealthCheckResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var HealthService = exports.HealthService = {
  check: {
    path: '/health.Health/Check',
    requestStream: false,
    responseStream: false,
    requestType: health_health_pb.HealthCheckRequest,
    responseType: health_health_pb.HealthCheckResponse,
    requestSerialize: serialize_health_HealthCheckRequest,
    requestDeserialize: deserialize_health_HealthCheckRequest,
    responseSerialize: serialize_health_HealthCheckResponse,
    responseDeserialize: deserialize_health_HealthCheckResponse,
  },
};

exports.HealthClient = grpc.makeGenericClientConstructor(HealthService);
