// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var compute_compute_pb = require('../compute/compute_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');
var images_images_pb = require('../images/images_pb.js');
var operators_operators_pb = require('../operators/operators_pb.js');

function serialize_compute_RenderImageRequest(arg) {
  if (!(arg instanceof compute_compute_pb.RenderImageRequest)) {
    throw new Error('Expected argument of type compute.RenderImageRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_compute_RenderImageRequest(buffer_arg) {
  return compute_compute_pb.RenderImageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_compute_RenderImageResponse(arg) {
  if (!(arg instanceof compute_compute_pb.RenderImageResponse)) {
    throw new Error('Expected argument of type compute.RenderImageResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_compute_RenderImageResponse(buffer_arg) {
  return compute_compute_pb.RenderImageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputeService = exports.ComputeService = {
  renderImage: {
    path: '/compute.Compute/RenderImage',
    requestStream: false,
    responseStream: false,
    requestType: compute_compute_pb.RenderImageRequest,
    responseType: compute_compute_pb.RenderImageResponse,
    requestSerialize: serialize_compute_RenderImageRequest,
    requestDeserialize: deserialize_compute_RenderImageRequest,
    responseSerialize: serialize_compute_RenderImageResponse,
    responseDeserialize: deserialize_compute_RenderImageResponse,
  },
};

exports.ComputeClient = grpc.makeGenericClientConstructor(ComputeService);
