// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var image_sequences_image_sequences_pb = require('../image_sequences/image_sequences_pb.js');
var google_api_annotations_pb = require('../google/api/annotations_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_image_sequences_CreateImageSequenceRequest(arg) {
  if (!(arg instanceof image_sequences_image_sequences_pb.CreateImageSequenceRequest)) {
    throw new Error('Expected argument of type image_sequences.CreateImageSequenceRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_image_sequences_CreateImageSequenceRequest(buffer_arg) {
  return image_sequences_image_sequences_pb.CreateImageSequenceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_image_sequences_DeleteImageSequenceRequest(arg) {
  if (!(arg instanceof image_sequences_image_sequences_pb.DeleteImageSequenceRequest)) {
    throw new Error('Expected argument of type image_sequences.DeleteImageSequenceRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_image_sequences_DeleteImageSequenceRequest(buffer_arg) {
  return image_sequences_image_sequences_pb.DeleteImageSequenceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_image_sequences_GetImageSequenceRequest(arg) {
  if (!(arg instanceof image_sequences_image_sequences_pb.GetImageSequenceRequest)) {
    throw new Error('Expected argument of type image_sequences.GetImageSequenceRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_image_sequences_GetImageSequenceRequest(buffer_arg) {
  return image_sequences_image_sequences_pb.GetImageSequenceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_image_sequences_ImageSequence(arg) {
  if (!(arg instanceof image_sequences_image_sequences_pb.ImageSequence)) {
    throw new Error('Expected argument of type image_sequences.ImageSequence');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_image_sequences_ImageSequence(buffer_arg) {
  return image_sequences_image_sequences_pb.ImageSequence.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_image_sequences_ListImageSequencesRequest(arg) {
  if (!(arg instanceof image_sequences_image_sequences_pb.ListImageSequencesRequest)) {
    throw new Error('Expected argument of type image_sequences.ListImageSequencesRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_image_sequences_ListImageSequencesRequest(buffer_arg) {
  return image_sequences_image_sequences_pb.ListImageSequencesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_image_sequences_ListImageSequencesResponse(arg) {
  if (!(arg instanceof image_sequences_image_sequences_pb.ListImageSequencesResponse)) {
    throw new Error('Expected argument of type image_sequences.ListImageSequencesResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_image_sequences_ListImageSequencesResponse(buffer_arg) {
  return image_sequences_image_sequences_pb.ListImageSequencesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_image_sequences_UpdateImageSequenceRequest(arg) {
  if (!(arg instanceof image_sequences_image_sequences_pb.UpdateImageSequenceRequest)) {
    throw new Error('Expected argument of type image_sequences.UpdateImageSequenceRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_image_sequences_UpdateImageSequenceRequest(buffer_arg) {
  return image_sequences_image_sequences_pb.UpdateImageSequenceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ImageSequencesService = exports.ImageSequencesService = {
  list: {
    path: '/image_sequences.ImageSequences/List',
    requestStream: false,
    responseStream: false,
    requestType: image_sequences_image_sequences_pb.ListImageSequencesRequest,
    responseType: image_sequences_image_sequences_pb.ListImageSequencesResponse,
    requestSerialize: serialize_image_sequences_ListImageSequencesRequest,
    requestDeserialize: deserialize_image_sequences_ListImageSequencesRequest,
    responseSerialize: serialize_image_sequences_ListImageSequencesResponse,
    responseDeserialize: deserialize_image_sequences_ListImageSequencesResponse,
  },
  get: {
    path: '/image_sequences.ImageSequences/Get',
    requestStream: false,
    responseStream: false,
    requestType: image_sequences_image_sequences_pb.GetImageSequenceRequest,
    responseType: image_sequences_image_sequences_pb.ImageSequence,
    requestSerialize: serialize_image_sequences_GetImageSequenceRequest,
    requestDeserialize: deserialize_image_sequences_GetImageSequenceRequest,
    responseSerialize: serialize_image_sequences_ImageSequence,
    responseDeserialize: deserialize_image_sequences_ImageSequence,
  },
  create: {
    path: '/image_sequences.ImageSequences/Create',
    requestStream: false,
    responseStream: false,
    requestType: image_sequences_image_sequences_pb.CreateImageSequenceRequest,
    responseType: image_sequences_image_sequences_pb.ImageSequence,
    requestSerialize: serialize_image_sequences_CreateImageSequenceRequest,
    requestDeserialize: deserialize_image_sequences_CreateImageSequenceRequest,
    responseSerialize: serialize_image_sequences_ImageSequence,
    responseDeserialize: deserialize_image_sequences_ImageSequence,
  },
  update: {
    path: '/image_sequences.ImageSequences/Update',
    requestStream: false,
    responseStream: false,
    requestType: image_sequences_image_sequences_pb.UpdateImageSequenceRequest,
    responseType: image_sequences_image_sequences_pb.ImageSequence,
    requestSerialize: serialize_image_sequences_UpdateImageSequenceRequest,
    requestDeserialize: deserialize_image_sequences_UpdateImageSequenceRequest,
    responseSerialize: serialize_image_sequences_ImageSequence,
    responseDeserialize: deserialize_image_sequences_ImageSequence,
  },
  delete: {
    path: '/image_sequences.ImageSequences/Delete',
    requestStream: false,
    responseStream: false,
    requestType: image_sequences_image_sequences_pb.DeleteImageSequenceRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_image_sequences_DeleteImageSequenceRequest,
    requestDeserialize: deserialize_image_sequences_DeleteImageSequenceRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.ImageSequencesClient = grpc.makeGenericClientConstructor(ImageSequencesService);
