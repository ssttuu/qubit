// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var scenes_scenes_pb = require('../scenes/scenes_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var operators_operators_pb = require('../operators/operators_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scenes_CreateSceneRequest(arg) {
  if (!(arg instanceof scenes_scenes_pb.CreateSceneRequest)) {
    throw new Error('Expected argument of type scenes.CreateSceneRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scenes_CreateSceneRequest(buffer_arg) {
  return scenes_scenes_pb.CreateSceneRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scenes_DeleteSceneRequest(arg) {
  if (!(arg instanceof scenes_scenes_pb.DeleteSceneRequest)) {
    throw new Error('Expected argument of type scenes.DeleteSceneRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scenes_DeleteSceneRequest(buffer_arg) {
  return scenes_scenes_pb.DeleteSceneRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scenes_GetSceneRequest(arg) {
  if (!(arg instanceof scenes_scenes_pb.GetSceneRequest)) {
    throw new Error('Expected argument of type scenes.GetSceneRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scenes_GetSceneRequest(buffer_arg) {
  return scenes_scenes_pb.GetSceneRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scenes_ListScenesRequest(arg) {
  if (!(arg instanceof scenes_scenes_pb.ListScenesRequest)) {
    throw new Error('Expected argument of type scenes.ListScenesRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scenes_ListScenesRequest(buffer_arg) {
  return scenes_scenes_pb.ListScenesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scenes_ListScenesResponse(arg) {
  if (!(arg instanceof scenes_scenes_pb.ListScenesResponse)) {
    throw new Error('Expected argument of type scenes.ListScenesResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scenes_ListScenesResponse(buffer_arg) {
  return scenes_scenes_pb.ListScenesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scenes_RenameSceneRequest(arg) {
  if (!(arg instanceof scenes_scenes_pb.RenameSceneRequest)) {
    throw new Error('Expected argument of type scenes.RenameSceneRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scenes_RenameSceneRequest(buffer_arg) {
  return scenes_scenes_pb.RenameSceneRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scenes_Scene(arg) {
  if (!(arg instanceof scenes_scenes_pb.Scene)) {
    throw new Error('Expected argument of type scenes.Scene');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scenes_Scene(buffer_arg) {
  return scenes_scenes_pb.Scene.deserializeBinary(new Uint8Array(buffer_arg));
}


var ScenesService = exports.ScenesService = {
  list: {
    path: '/scenes.Scenes/List',
    requestStream: false,
    responseStream: false,
    requestType: scenes_scenes_pb.ListScenesRequest,
    responseType: scenes_scenes_pb.ListScenesResponse,
    requestSerialize: serialize_scenes_ListScenesRequest,
    requestDeserialize: deserialize_scenes_ListScenesRequest,
    responseSerialize: serialize_scenes_ListScenesResponse,
    responseDeserialize: deserialize_scenes_ListScenesResponse,
  },
  get: {
    path: '/scenes.Scenes/Get',
    requestStream: false,
    responseStream: false,
    requestType: scenes_scenes_pb.GetSceneRequest,
    responseType: scenes_scenes_pb.Scene,
    requestSerialize: serialize_scenes_GetSceneRequest,
    requestDeserialize: deserialize_scenes_GetSceneRequest,
    responseSerialize: serialize_scenes_Scene,
    responseDeserialize: deserialize_scenes_Scene,
  },
  create: {
    path: '/scenes.Scenes/Create',
    requestStream: false,
    responseStream: false,
    requestType: scenes_scenes_pb.CreateSceneRequest,
    responseType: scenes_scenes_pb.Scene,
    requestSerialize: serialize_scenes_CreateSceneRequest,
    requestDeserialize: deserialize_scenes_CreateSceneRequest,
    responseSerialize: serialize_scenes_Scene,
    responseDeserialize: deserialize_scenes_Scene,
  },
  delete: {
    path: '/scenes.Scenes/Delete',
    requestStream: false,
    responseStream: false,
    requestType: scenes_scenes_pb.DeleteSceneRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_scenes_DeleteSceneRequest,
    requestDeserialize: deserialize_scenes_DeleteSceneRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  rename: {
    path: '/scenes.Scenes/Rename',
    requestStream: false,
    responseStream: false,
    requestType: scenes_scenes_pb.RenameSceneRequest,
    responseType: scenes_scenes_pb.Scene,
    requestSerialize: serialize_scenes_RenameSceneRequest,
    requestDeserialize: deserialize_scenes_RenameSceneRequest,
    responseSerialize: serialize_scenes_Scene,
    responseDeserialize: deserialize_scenes_Scene,
  },
};

exports.ScenesClient = grpc.makeGenericClientConstructor(ScenesService);
