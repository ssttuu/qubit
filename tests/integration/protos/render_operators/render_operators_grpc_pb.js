// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var render_operators_render_operators_pb = require('../render_operators/render_operators_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');

function serialize_render_operators_RenderOperator(arg) {
  if (!(arg instanceof render_operators_render_operators_pb.RenderOperator)) {
    throw new Error('Expected argument of type render_operators.RenderOperator');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_render_operators_RenderOperator(buffer_arg) {
  return render_operators_render_operators_pb.RenderOperator.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_render_operators_RenderOperatorRequest(arg) {
  if (!(arg instanceof render_operators_render_operators_pb.RenderOperatorRequest)) {
    throw new Error('Expected argument of type render_operators.RenderOperatorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_render_operators_RenderOperatorRequest(buffer_arg) {
  return render_operators_render_operators_pb.RenderOperatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var RenderOperatorsService = exports.RenderOperatorsService = {
  getRenderOperators: {
    path: '/render_operators.RenderOperators/GetRenderOperators',
    requestStream: false,
    responseStream: false,
    requestType: render_operators_render_operators_pb.RenderOperatorRequest,
    responseType: render_operators_render_operators_pb.RenderOperator,
    requestSerialize: serialize_render_operators_RenderOperatorRequest,
    requestDeserialize: deserialize_render_operators_RenderOperatorRequest,
    responseSerialize: serialize_render_operators_RenderOperator,
    responseDeserialize: deserialize_render_operators_RenderOperator,
  },
};

exports.RenderOperatorsClient = grpc.makeGenericClientConstructor(RenderOperatorsService);
