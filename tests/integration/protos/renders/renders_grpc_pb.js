// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var renders_renders_pb = require('../renders/renders_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');

function serialize_renders_RenderRequest(arg) {
  if (!(arg instanceof renders_renders_pb.RenderRequest)) {
    throw new Error('Expected argument of type renders.RenderRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_renders_RenderRequest(buffer_arg) {
  return renders_renders_pb.RenderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_renders_RenderResponse(arg) {
  if (!(arg instanceof renders_renders_pb.RenderResponse)) {
    throw new Error('Expected argument of type renders.RenderResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_renders_RenderResponse(buffer_arg) {
  return renders_renders_pb.RenderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var RendersService = exports.RendersService = {
  doRender: {
    path: '/renders.Renders/DoRender',
    requestStream: false,
    responseStream: false,
    requestType: renders_renders_pb.RenderRequest,
    responseType: renders_renders_pb.RenderResponse,
    requestSerialize: serialize_renders_RenderRequest,
    requestDeserialize: deserialize_renders_RenderRequest,
    responseSerialize: serialize_renders_RenderResponse,
    responseDeserialize: deserialize_renders_RenderResponse,
  },
};

exports.RendersClient = grpc.makeGenericClientConstructor(RendersService);
