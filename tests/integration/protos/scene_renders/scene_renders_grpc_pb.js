// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var scene_renders_scene_renders_pb = require('../scene_renders/scene_renders_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var geometry_geometry_pb = require('../geometry/geometry_pb.js');

function serialize_scene_renders_SceneRenderRequest(arg) {
  if (!(arg instanceof scene_renders_scene_renders_pb.SceneRenderRequest)) {
    throw new Error('Expected argument of type scene_renders.SceneRenderRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_renders_SceneRenderRequest(buffer_arg) {
  return scene_renders_scene_renders_pb.SceneRenderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scene_renders_SceneRenderStatus(arg) {
  if (!(arg instanceof scene_renders_scene_renders_pb.SceneRenderStatus)) {
    throw new Error('Expected argument of type scene_renders.SceneRenderStatus');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_renders_SceneRenderStatus(buffer_arg) {
  return scene_renders_scene_renders_pb.SceneRenderStatus.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scene_renders_SceneRenderStatusRequest(arg) {
  if (!(arg instanceof scene_renders_scene_renders_pb.SceneRenderStatusRequest)) {
    throw new Error('Expected argument of type scene_renders.SceneRenderStatusRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_renders_SceneRenderStatusRequest(buffer_arg) {
  return scene_renders_scene_renders_pb.SceneRenderStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var SceneRendersService = exports.SceneRendersService = {
  create: {
    path: '/scene_renders.SceneRenders/Create',
    requestStream: false,
    responseStream: false,
    requestType: scene_renders_scene_renders_pb.SceneRenderRequest,
    responseType: scene_renders_scene_renders_pb.SceneRenderStatus,
    requestSerialize: serialize_scene_renders_SceneRenderRequest,
    requestDeserialize: deserialize_scene_renders_SceneRenderRequest,
    responseSerialize: serialize_scene_renders_SceneRenderStatus,
    responseDeserialize: deserialize_scene_renders_SceneRenderStatus,
  },
  get: {
    path: '/scene_renders.SceneRenders/Get',
    requestStream: false,
    responseStream: false,
    requestType: scene_renders_scene_renders_pb.SceneRenderStatusRequest,
    responseType: scene_renders_scene_renders_pb.SceneRenderStatus,
    requestSerialize: serialize_scene_renders_SceneRenderStatusRequest,
    requestDeserialize: deserialize_scene_renders_SceneRenderStatusRequest,
    responseSerialize: serialize_scene_renders_SceneRenderStatus,
    responseDeserialize: deserialize_scene_renders_SceneRenderStatus,
  },
};

exports.SceneRendersClient = grpc.makeGenericClientConstructor(SceneRendersService);
