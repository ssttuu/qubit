// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var images_images_pb = require('../images/images_pb.js');
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

function serialize_images_CreateImageRequest(arg) {
  if (!(arg instanceof images_images_pb.CreateImageRequest)) {
    throw new Error('Expected argument of type images.CreateImageRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_images_CreateImageRequest(buffer_arg) {
  return images_images_pb.CreateImageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_images_DeleteImageRequest(arg) {
  if (!(arg instanceof images_images_pb.DeleteImageRequest)) {
    throw new Error('Expected argument of type images.DeleteImageRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_images_DeleteImageRequest(buffer_arg) {
  return images_images_pb.DeleteImageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_images_Frame(arg) {
  if (!(arg instanceof images_images_pb.Frame)) {
    throw new Error('Expected argument of type images.Frame');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_images_Frame(buffer_arg) {
  return images_images_pb.Frame.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_images_GetImageRequest(arg) {
  if (!(arg instanceof images_images_pb.GetImageRequest)) {
    throw new Error('Expected argument of type images.GetImageRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_images_GetImageRequest(buffer_arg) {
  return images_images_pb.GetImageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_images_ListImagesRequest(arg) {
  if (!(arg instanceof images_images_pb.ListImagesRequest)) {
    throw new Error('Expected argument of type images.ListImagesRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_images_ListImagesRequest(buffer_arg) {
  return images_images_pb.ListImagesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_images_ListImagesResponse(arg) {
  if (!(arg instanceof images_images_pb.ListImagesResponse)) {
    throw new Error('Expected argument of type images.ListImagesResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_images_ListImagesResponse(buffer_arg) {
  return images_images_pb.ListImagesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_images_UpdateImageRequest(arg) {
  if (!(arg instanceof images_images_pb.UpdateImageRequest)) {
    throw new Error('Expected argument of type images.UpdateImageRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_images_UpdateImageRequest(buffer_arg) {
  return images_images_pb.UpdateImageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ImagesService = exports.ImagesService = {
  list: {
    path: '/images.Images/List',
    requestStream: false,
    responseStream: false,
    requestType: images_images_pb.ListImagesRequest,
    responseType: images_images_pb.ListImagesResponse,
    requestSerialize: serialize_images_ListImagesRequest,
    requestDeserialize: deserialize_images_ListImagesRequest,
    responseSerialize: serialize_images_ListImagesResponse,
    responseDeserialize: deserialize_images_ListImagesResponse,
  },
  get: {
    path: '/images.Images/Get',
    requestStream: false,
    responseStream: false,
    requestType: images_images_pb.GetImageRequest,
    responseType: images_images_pb.Frame,
    requestSerialize: serialize_images_GetImageRequest,
    requestDeserialize: deserialize_images_GetImageRequest,
    responseSerialize: serialize_images_Frame,
    responseDeserialize: deserialize_images_Frame,
  },
  create: {
    path: '/images.Images/Create',
    requestStream: false,
    responseStream: false,
    requestType: images_images_pb.CreateImageRequest,
    responseType: images_images_pb.Frame,
    requestSerialize: serialize_images_CreateImageRequest,
    requestDeserialize: deserialize_images_CreateImageRequest,
    responseSerialize: serialize_images_Frame,
    responseDeserialize: deserialize_images_Frame,
  },
  update: {
    path: '/images.Images/Update',
    requestStream: false,
    responseStream: false,
    requestType: images_images_pb.UpdateImageRequest,
    responseType: images_images_pb.Frame,
    requestSerialize: serialize_images_UpdateImageRequest,
    requestDeserialize: deserialize_images_UpdateImageRequest,
    responseSerialize: serialize_images_Frame,
    responseDeserialize: deserialize_images_Frame,
  },
  delete: {
    path: '/images.Images/Delete',
    requestStream: false,
    responseStream: false,
    requestType: images_images_pb.DeleteImageRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_images_DeleteImageRequest,
    requestDeserialize: deserialize_images_DeleteImageRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.ImagesClient = grpc.makeGenericClientConstructor(ImagesService);
