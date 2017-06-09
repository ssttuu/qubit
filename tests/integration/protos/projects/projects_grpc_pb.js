// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var projects_projects_pb = require('../projects/projects_pb.js');
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

function serialize_projects_CreateProjectRequest(arg) {
  if (!(arg instanceof projects_projects_pb.CreateProjectRequest)) {
    throw new Error('Expected argument of type projects.CreateProjectRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_projects_CreateProjectRequest(buffer_arg) {
  return projects_projects_pb.CreateProjectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_projects_DeleteProjectRequest(arg) {
  if (!(arg instanceof projects_projects_pb.DeleteProjectRequest)) {
    throw new Error('Expected argument of type projects.DeleteProjectRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_projects_DeleteProjectRequest(buffer_arg) {
  return projects_projects_pb.DeleteProjectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_projects_GetProjectRequest(arg) {
  if (!(arg instanceof projects_projects_pb.GetProjectRequest)) {
    throw new Error('Expected argument of type projects.GetProjectRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_projects_GetProjectRequest(buffer_arg) {
  return projects_projects_pb.GetProjectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_projects_ListProjectsRequest(arg) {
  if (!(arg instanceof projects_projects_pb.ListProjectsRequest)) {
    throw new Error('Expected argument of type projects.ListProjectsRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_projects_ListProjectsRequest(buffer_arg) {
  return projects_projects_pb.ListProjectsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_projects_ListProjectsResponse(arg) {
  if (!(arg instanceof projects_projects_pb.ListProjectsResponse)) {
    throw new Error('Expected argument of type projects.ListProjectsResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_projects_ListProjectsResponse(buffer_arg) {
  return projects_projects_pb.ListProjectsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_projects_Project(arg) {
  if (!(arg instanceof projects_projects_pb.Project)) {
    throw new Error('Expected argument of type projects.Project');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_projects_Project(buffer_arg) {
  return projects_projects_pb.Project.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_projects_UpdateProjectRequest(arg) {
  if (!(arg instanceof projects_projects_pb.UpdateProjectRequest)) {
    throw new Error('Expected argument of type projects.UpdateProjectRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_projects_UpdateProjectRequest(buffer_arg) {
  return projects_projects_pb.UpdateProjectRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var ProjectsService = exports.ProjectsService = {
  list: {
    path: '/projects.Projects/List',
    requestStream: false,
    responseStream: false,
    requestType: projects_projects_pb.ListProjectsRequest,
    responseType: projects_projects_pb.ListProjectsResponse,
    requestSerialize: serialize_projects_ListProjectsRequest,
    requestDeserialize: deserialize_projects_ListProjectsRequest,
    responseSerialize: serialize_projects_ListProjectsResponse,
    responseDeserialize: deserialize_projects_ListProjectsResponse,
  },
  get: {
    path: '/projects.Projects/Get',
    requestStream: false,
    responseStream: false,
    requestType: projects_projects_pb.GetProjectRequest,
    responseType: projects_projects_pb.Project,
    requestSerialize: serialize_projects_GetProjectRequest,
    requestDeserialize: deserialize_projects_GetProjectRequest,
    responseSerialize: serialize_projects_Project,
    responseDeserialize: deserialize_projects_Project,
  },
  create: {
    path: '/projects.Projects/Create',
    requestStream: false,
    responseStream: false,
    requestType: projects_projects_pb.CreateProjectRequest,
    responseType: projects_projects_pb.Project,
    requestSerialize: serialize_projects_CreateProjectRequest,
    requestDeserialize: deserialize_projects_CreateProjectRequest,
    responseSerialize: serialize_projects_Project,
    responseDeserialize: deserialize_projects_Project,
  },
  update: {
    path: '/projects.Projects/Update',
    requestStream: false,
    responseStream: false,
    requestType: projects_projects_pb.UpdateProjectRequest,
    responseType: projects_projects_pb.Project,
    requestSerialize: serialize_projects_UpdateProjectRequest,
    requestDeserialize: deserialize_projects_UpdateProjectRequest,
    responseSerialize: serialize_projects_Project,
    responseDeserialize: deserialize_projects_Project,
  },
  delete: {
    path: '/projects.Projects/Delete',
    requestStream: false,
    responseStream: false,
    requestType: projects_projects_pb.DeleteProjectRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_projects_DeleteProjectRequest,
    requestDeserialize: deserialize_projects_DeleteProjectRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.ProjectsClient = grpc.makeGenericClientConstructor(ProjectsService);
