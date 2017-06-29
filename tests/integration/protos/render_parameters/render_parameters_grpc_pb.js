// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var render_parameters_render_parameters_pb = require('../render_parameters/render_parameters_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');

function serialize_render_parameters_RenderParameter(arg) {
  if (!(arg instanceof render_parameters_render_parameters_pb.RenderParameter)) {
    throw new Error('Expected argument of type render_parameters.RenderParameter');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_render_parameters_RenderParameter(buffer_arg) {
  return render_parameters_render_parameters_pb.RenderParameter.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_render_parameters_RenderParameterRequest(arg) {
  if (!(arg instanceof render_parameters_render_parameters_pb.RenderParameterRequest)) {
    throw new Error('Expected argument of type render_parameters.RenderParameterRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_render_parameters_RenderParameterRequest(buffer_arg) {
  return render_parameters_render_parameters_pb.RenderParameterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var RenderParametersService = exports.RenderParametersService = {
  getRenderParameters: {
    path: '/render_parameters.RenderParameters/GetRenderParameters',
    requestStream: false,
    responseStream: false,
    requestType: render_parameters_render_parameters_pb.RenderParameterRequest,
    responseType: render_parameters_render_parameters_pb.RenderParameter,
    requestSerialize: serialize_render_parameters_RenderParameterRequest,
    requestDeserialize: deserialize_render_parameters_RenderParameterRequest,
    responseSerialize: serialize_render_parameters_RenderParameter,
    responseDeserialize: deserialize_render_parameters_RenderParameter,
  },
};

exports.RenderParametersClient = grpc.makeGenericClientConstructor(RenderParametersService);
