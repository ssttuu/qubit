// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var scene_events_scene_events_pb = require('../scene_events/scene_events_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');

function serialize_scene_events_CreateSceneEventRequest(arg) {
  if (!(arg instanceof scene_events_scene_events_pb.CreateSceneEventRequest)) {
    throw new Error('Expected argument of type scene_events.CreateSceneEventRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_events_CreateSceneEventRequest(buffer_arg) {
  return scene_events_scene_events_pb.CreateSceneEventRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scene_events_GetSceneEventRequest(arg) {
  if (!(arg instanceof scene_events_scene_events_pb.GetSceneEventRequest)) {
    throw new Error('Expected argument of type scene_events.GetSceneEventRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_events_GetSceneEventRequest(buffer_arg) {
  return scene_events_scene_events_pb.GetSceneEventRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scene_events_ListSceneEventsRequest(arg) {
  if (!(arg instanceof scene_events_scene_events_pb.ListSceneEventsRequest)) {
    throw new Error('Expected argument of type scene_events.ListSceneEventsRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_events_ListSceneEventsRequest(buffer_arg) {
  return scene_events_scene_events_pb.ListSceneEventsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scene_events_ListSceneEventsResponse(arg) {
  if (!(arg instanceof scene_events_scene_events_pb.ListSceneEventsResponse)) {
    throw new Error('Expected argument of type scene_events.ListSceneEventsResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_events_ListSceneEventsResponse(buffer_arg) {
  return scene_events_scene_events_pb.ListSceneEventsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_scene_events_SceneEvent(arg) {
  if (!(arg instanceof scene_events_scene_events_pb.SceneEvent)) {
    throw new Error('Expected argument of type scene_events.SceneEvent');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_scene_events_SceneEvent(buffer_arg) {
  return scene_events_scene_events_pb.SceneEvent.deserializeBinary(new Uint8Array(buffer_arg));
}


var SceneEventsService = exports.SceneEventsService = {
  list: {
    path: '/scene_events.SceneEvents/List',
    requestStream: false,
    responseStream: false,
    requestType: scene_events_scene_events_pb.ListSceneEventsRequest,
    responseType: scene_events_scene_events_pb.ListSceneEventsResponse,
    requestSerialize: serialize_scene_events_ListSceneEventsRequest,
    requestDeserialize: deserialize_scene_events_ListSceneEventsRequest,
    responseSerialize: serialize_scene_events_ListSceneEventsResponse,
    responseDeserialize: deserialize_scene_events_ListSceneEventsResponse,
  },
  get: {
    path: '/scene_events.SceneEvents/Get',
    requestStream: false,
    responseStream: false,
    requestType: scene_events_scene_events_pb.GetSceneEventRequest,
    responseType: scene_events_scene_events_pb.SceneEvent,
    requestSerialize: serialize_scene_events_GetSceneEventRequest,
    requestDeserialize: deserialize_scene_events_GetSceneEventRequest,
    responseSerialize: serialize_scene_events_SceneEvent,
    responseDeserialize: deserialize_scene_events_SceneEvent,
  },
  create: {
    path: '/scene_events.SceneEvents/Create',
    requestStream: false,
    responseStream: false,
    requestType: scene_events_scene_events_pb.CreateSceneEventRequest,
    responseType: scene_events_scene_events_pb.SceneEvent,
    requestSerialize: serialize_scene_events_CreateSceneEventRequest,
    requestDeserialize: deserialize_scene_events_CreateSceneEventRequest,
    responseSerialize: serialize_scene_events_SceneEvent,
    responseDeserialize: deserialize_scene_events_SceneEvent,
  },
};

exports.SceneEventsClient = grpc.makeGenericClientConstructor(SceneEventsService);
