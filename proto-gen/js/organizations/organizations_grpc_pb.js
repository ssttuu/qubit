// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var organizations_organizations_pb = require('../organizations/organizations_pb.js');
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

function serialize_organizations_CreateOrganizationRequest(arg) {
  if (!(arg instanceof organizations_organizations_pb.CreateOrganizationRequest)) {
    throw new Error('Expected argument of type organizations.CreateOrganizationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_organizations_CreateOrganizationRequest(buffer_arg) {
  return organizations_organizations_pb.CreateOrganizationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_organizations_DeleteOrganizationRequest(arg) {
  if (!(arg instanceof organizations_organizations_pb.DeleteOrganizationRequest)) {
    throw new Error('Expected argument of type organizations.DeleteOrganizationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_organizations_DeleteOrganizationRequest(buffer_arg) {
  return organizations_organizations_pb.DeleteOrganizationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_organizations_GetOrganizationRequest(arg) {
  if (!(arg instanceof organizations_organizations_pb.GetOrganizationRequest)) {
    throw new Error('Expected argument of type organizations.GetOrganizationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_organizations_GetOrganizationRequest(buffer_arg) {
  return organizations_organizations_pb.GetOrganizationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_organizations_ListOrganizationsRequest(arg) {
  if (!(arg instanceof organizations_organizations_pb.ListOrganizationsRequest)) {
    throw new Error('Expected argument of type organizations.ListOrganizationsRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_organizations_ListOrganizationsRequest(buffer_arg) {
  return organizations_organizations_pb.ListOrganizationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_organizations_ListOrganizationsResponse(arg) {
  if (!(arg instanceof organizations_organizations_pb.ListOrganizationsResponse)) {
    throw new Error('Expected argument of type organizations.ListOrganizationsResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_organizations_ListOrganizationsResponse(buffer_arg) {
  return organizations_organizations_pb.ListOrganizationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_organizations_Organization(arg) {
  if (!(arg instanceof organizations_organizations_pb.Organization)) {
    throw new Error('Expected argument of type organizations.Organization');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_organizations_Organization(buffer_arg) {
  return organizations_organizations_pb.Organization.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_organizations_UpdateOrganizationRequest(arg) {
  if (!(arg instanceof organizations_organizations_pb.UpdateOrganizationRequest)) {
    throw new Error('Expected argument of type organizations.UpdateOrganizationRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_organizations_UpdateOrganizationRequest(buffer_arg) {
  return organizations_organizations_pb.UpdateOrganizationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var OrganizationsService = exports.OrganizationsService = {
  list: {
    path: '/organizations.Organizations/List',
    requestStream: false,
    responseStream: false,
    requestType: organizations_organizations_pb.ListOrganizationsRequest,
    responseType: organizations_organizations_pb.ListOrganizationsResponse,
    requestSerialize: serialize_organizations_ListOrganizationsRequest,
    requestDeserialize: deserialize_organizations_ListOrganizationsRequest,
    responseSerialize: serialize_organizations_ListOrganizationsResponse,
    responseDeserialize: deserialize_organizations_ListOrganizationsResponse,
  },
  get: {
    path: '/organizations.Organizations/Get',
    requestStream: false,
    responseStream: false,
    requestType: organizations_organizations_pb.GetOrganizationRequest,
    responseType: organizations_organizations_pb.Organization,
    requestSerialize: serialize_organizations_GetOrganizationRequest,
    requestDeserialize: deserialize_organizations_GetOrganizationRequest,
    responseSerialize: serialize_organizations_Organization,
    responseDeserialize: deserialize_organizations_Organization,
  },
  create: {
    path: '/organizations.Organizations/Create',
    requestStream: false,
    responseStream: false,
    requestType: organizations_organizations_pb.CreateOrganizationRequest,
    responseType: organizations_organizations_pb.Organization,
    requestSerialize: serialize_organizations_CreateOrganizationRequest,
    requestDeserialize: deserialize_organizations_CreateOrganizationRequest,
    responseSerialize: serialize_organizations_Organization,
    responseDeserialize: deserialize_organizations_Organization,
  },
  update: {
    path: '/organizations.Organizations/Update',
    requestStream: false,
    responseStream: false,
    requestType: organizations_organizations_pb.UpdateOrganizationRequest,
    responseType: organizations_organizations_pb.Organization,
    requestSerialize: serialize_organizations_UpdateOrganizationRequest,
    requestDeserialize: deserialize_organizations_UpdateOrganizationRequest,
    responseSerialize: serialize_organizations_Organization,
    responseDeserialize: deserialize_organizations_Organization,
  },
  delete: {
    path: '/organizations.Organizations/Delete',
    requestStream: false,
    responseStream: false,
    requestType: organizations_organizations_pb.DeleteOrganizationRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_organizations_DeleteOrganizationRequest,
    requestDeserialize: deserialize_organizations_DeleteOrganizationRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.OrganizationsClient = grpc.makeGenericClientConstructor(OrganizationsService);
