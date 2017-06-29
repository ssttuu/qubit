// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api-web/api-web.proto

/*
Package web is a generated protocol buffer package.

It is generated from these files:
	api-web/api-web.proto

It has these top-level messages:
*/
package web

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import image_sequences "github.com/stupschwartz/qubit/proto-gen/go/image_sequences"
import images "github.com/stupschwartz/qubit/proto-gen/go/images"
import projects "github.com/stupschwartz/qubit/proto-gen/go/projects"
import organizations "github.com/stupschwartz/qubit/proto-gen/go/organizations"
import render_parameters "github.com/stupschwartz/qubit/proto-gen/go/render_parameters"
import renders "github.com/stupschwartz/qubit/proto-gen/go/renders"
import scenes "github.com/stupschwartz/qubit/proto-gen/go/scenes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ImageSequence from public import image_sequences/image_sequences.proto
type ImageSequence image_sequences.ImageSequence

func (m *ImageSequence) Reset()         { (*image_sequences.ImageSequence)(m).Reset() }
func (m *ImageSequence) String() string { return (*image_sequences.ImageSequence)(m).String() }
func (*ImageSequence) ProtoMessage()    {}
func (m *ImageSequence) GetId() string  { return (*image_sequences.ImageSequence)(m).GetId() }
func (m *ImageSequence) GetProjectId() string {
	return (*image_sequences.ImageSequence)(m).GetProjectId()
}
func (m *ImageSequence) GetName() string { return (*image_sequences.ImageSequence)(m).GetName() }

// ListImageSequencesRequest from public import image_sequences/image_sequences.proto
type ListImageSequencesRequest image_sequences.ListImageSequencesRequest

func (m *ListImageSequencesRequest) Reset() { (*image_sequences.ListImageSequencesRequest)(m).Reset() }
func (m *ListImageSequencesRequest) String() string {
	return (*image_sequences.ListImageSequencesRequest)(m).String()
}
func (*ListImageSequencesRequest) ProtoMessage() {}
func (m *ListImageSequencesRequest) GetPageSize() int32 {
	return (*image_sequences.ListImageSequencesRequest)(m).GetPageSize()
}
func (m *ListImageSequencesRequest) GetPageToken() string {
	return (*image_sequences.ListImageSequencesRequest)(m).GetPageToken()
}

// ListImageSequencesResponse from public import image_sequences/image_sequences.proto
type ListImageSequencesResponse image_sequences.ListImageSequencesResponse

func (m *ListImageSequencesResponse) Reset() { (*image_sequences.ListImageSequencesResponse)(m).Reset() }
func (m *ListImageSequencesResponse) String() string {
	return (*image_sequences.ListImageSequencesResponse)(m).String()
}
func (*ListImageSequencesResponse) ProtoMessage() {}
func (m *ListImageSequencesResponse) GetImageSequences() []*ImageSequence {
	o := (*image_sequences.ListImageSequencesResponse)(m).GetImageSequences()
	if o == nil {
		return nil
	}
	s := make([]*ImageSequence, len(o))
	for i, x := range o {
		s[i] = (*ImageSequence)(x)
	}
	return s
}
func (m *ListImageSequencesResponse) GetNextPageToken() string {
	return (*image_sequences.ListImageSequencesResponse)(m).GetNextPageToken()
}

// GetImageSequenceRequest from public import image_sequences/image_sequences.proto
type GetImageSequenceRequest image_sequences.GetImageSequenceRequest

func (m *GetImageSequenceRequest) Reset() { (*image_sequences.GetImageSequenceRequest)(m).Reset() }
func (m *GetImageSequenceRequest) String() string {
	return (*image_sequences.GetImageSequenceRequest)(m).String()
}
func (*GetImageSequenceRequest) ProtoMessage() {}
func (m *GetImageSequenceRequest) GetId() string {
	return (*image_sequences.GetImageSequenceRequest)(m).GetId()
}

// CreateImageSequenceRequest from public import image_sequences/image_sequences.proto
type CreateImageSequenceRequest image_sequences.CreateImageSequenceRequest

func (m *CreateImageSequenceRequest) Reset() { (*image_sequences.CreateImageSequenceRequest)(m).Reset() }
func (m *CreateImageSequenceRequest) String() string {
	return (*image_sequences.CreateImageSequenceRequest)(m).String()
}
func (*CreateImageSequenceRequest) ProtoMessage() {}
func (m *CreateImageSequenceRequest) GetImageSequence() *ImageSequence {
	return (*ImageSequence)((*image_sequences.CreateImageSequenceRequest)(m).GetImageSequence())
}

// UpdateImageSequenceRequest from public import image_sequences/image_sequences.proto
type UpdateImageSequenceRequest image_sequences.UpdateImageSequenceRequest

func (m *UpdateImageSequenceRequest) Reset() { (*image_sequences.UpdateImageSequenceRequest)(m).Reset() }
func (m *UpdateImageSequenceRequest) String() string {
	return (*image_sequences.UpdateImageSequenceRequest)(m).String()
}
func (*UpdateImageSequenceRequest) ProtoMessage() {}
func (m *UpdateImageSequenceRequest) GetId() string {
	return (*image_sequences.UpdateImageSequenceRequest)(m).GetId()
}
func (m *UpdateImageSequenceRequest) GetImageSequence() *ImageSequence {
	return (*ImageSequence)((*image_sequences.UpdateImageSequenceRequest)(m).GetImageSequence())
}

// DeleteImageSequenceRequest from public import image_sequences/image_sequences.proto
type DeleteImageSequenceRequest image_sequences.DeleteImageSequenceRequest

func (m *DeleteImageSequenceRequest) Reset() { (*image_sequences.DeleteImageSequenceRequest)(m).Reset() }
func (m *DeleteImageSequenceRequest) String() string {
	return (*image_sequences.DeleteImageSequenceRequest)(m).String()
}
func (*DeleteImageSequenceRequest) ProtoMessage() {}
func (m *DeleteImageSequenceRequest) GetId() string {
	return (*image_sequences.DeleteImageSequenceRequest)(m).GetId()
}

// Image from public import images/images.proto
type Image images.Image

func (m *Image) Reset()                     { (*images.Image)(m).Reset() }
func (m *Image) String() string             { return (*images.Image)(m).String() }
func (*Image) ProtoMessage()                {}
func (m *Image) GetId() string              { return (*images.Image)(m).GetId() }
func (m *Image) GetImageSequenceId() string { return (*images.Image)(m).GetImageSequenceId() }
func (m *Image) GetName() string            { return (*images.Image)(m).GetName() }
func (m *Image) GetWidth() int32            { return (*images.Image)(m).GetWidth() }
func (m *Image) GetHeight() int32           { return (*images.Image)(m).GetHeight() }
func (m *Image) GetLabels() map[string]string {
	o := (*images.Image)(m).GetLabels()
	if o == nil {
		return nil
	}
	s := make(map[string]string, len(o))
	for k, v := range o {
		s[k] = (string)(v)
	}
	return s
}
func (m *Image) GetPlanes() []*Plane {
	o := (*images.Image)(m).GetPlanes()
	if o == nil {
		return nil
	}
	s := make([]*Plane, len(o))
	for i, x := range o {
		s[i] = (*Plane)(x)
	}
	return s
}

// Plane from public import images/images.proto
type Plane images.Plane

func (m *Plane) Reset()           { (*images.Plane)(m).Reset() }
func (m *Plane) String() string   { return (*images.Plane)(m).String() }
func (*Plane) ProtoMessage()      {}
func (m *Plane) GetName() string  { return (*images.Plane)(m).GetName() }
func (m *Plane) GetWidth() int32  { return (*images.Plane)(m).GetWidth() }
func (m *Plane) GetHeight() int32 { return (*images.Plane)(m).GetHeight() }
func (m *Plane) GetLabels() map[string]string {
	o := (*images.Plane)(m).GetLabels()
	if o == nil {
		return nil
	}
	s := make(map[string]string, len(o))
	for k, v := range o {
		s[k] = (string)(v)
	}
	return s
}
func (m *Plane) GetChannels() []*Channel {
	o := (*images.Plane)(m).GetChannels()
	if o == nil {
		return nil
	}
	s := make([]*Channel, len(o))
	for i, x := range o {
		s[i] = (*Channel)(x)
	}
	return s
}

// Channel from public import images/images.proto
type Channel images.Channel

func (m *Channel) Reset()          { (*images.Channel)(m).Reset() }
func (m *Channel) String() string  { return (*images.Channel)(m).String() }
func (*Channel) ProtoMessage()     {}
func (m *Channel) GetName() string { return (*images.Channel)(m).GetName() }
func (m *Channel) GetRows() []*Row {
	o := (*images.Channel)(m).GetRows()
	if o == nil {
		return nil
	}
	s := make([]*Row, len(o))
	for i, x := range o {
		s[i] = (*Row)(x)
	}
	return s
}

// Row from public import images/images.proto
type Row images.Row

func (m *Row) Reset()             { (*images.Row)(m).Reset() }
func (m *Row) String() string     { return (*images.Row)(m).String() }
func (*Row) ProtoMessage()        {}
func (m *Row) GetData() []float64 { return (*images.Row)(m).GetData() }

// ListImagesRequest from public import images/images.proto
type ListImagesRequest images.ListImagesRequest

func (m *ListImagesRequest) Reset()             { (*images.ListImagesRequest)(m).Reset() }
func (m *ListImagesRequest) String() string     { return (*images.ListImagesRequest)(m).String() }
func (*ListImagesRequest) ProtoMessage()        {}
func (m *ListImagesRequest) GetPageSize() int32 { return (*images.ListImagesRequest)(m).GetPageSize() }
func (m *ListImagesRequest) GetPageToken() string {
	return (*images.ListImagesRequest)(m).GetPageToken()
}

// ListImagesResponse from public import images/images.proto
type ListImagesResponse images.ListImagesResponse

func (m *ListImagesResponse) Reset()         { (*images.ListImagesResponse)(m).Reset() }
func (m *ListImagesResponse) String() string { return (*images.ListImagesResponse)(m).String() }
func (*ListImagesResponse) ProtoMessage()    {}
func (m *ListImagesResponse) GetImages() []*Image {
	o := (*images.ListImagesResponse)(m).GetImages()
	if o == nil {
		return nil
	}
	s := make([]*Image, len(o))
	for i, x := range o {
		s[i] = (*Image)(x)
	}
	return s
}
func (m *ListImagesResponse) GetNextPageToken() string {
	return (*images.ListImagesResponse)(m).GetNextPageToken()
}

// GetImageRequest from public import images/images.proto
type GetImageRequest images.GetImageRequest

func (m *GetImageRequest) Reset()         { (*images.GetImageRequest)(m).Reset() }
func (m *GetImageRequest) String() string { return (*images.GetImageRequest)(m).String() }
func (*GetImageRequest) ProtoMessage()    {}
func (m *GetImageRequest) GetId() string  { return (*images.GetImageRequest)(m).GetId() }

// CreateImageRequest from public import images/images.proto
type CreateImageRequest images.CreateImageRequest

func (m *CreateImageRequest) Reset()         { (*images.CreateImageRequest)(m).Reset() }
func (m *CreateImageRequest) String() string { return (*images.CreateImageRequest)(m).String() }
func (*CreateImageRequest) ProtoMessage()    {}
func (m *CreateImageRequest) GetImage() *Image {
	return (*Image)((*images.CreateImageRequest)(m).GetImage())
}

// UpdateImageRequest from public import images/images.proto
type UpdateImageRequest images.UpdateImageRequest

func (m *UpdateImageRequest) Reset()         { (*images.UpdateImageRequest)(m).Reset() }
func (m *UpdateImageRequest) String() string { return (*images.UpdateImageRequest)(m).String() }
func (*UpdateImageRequest) ProtoMessage()    {}
func (m *UpdateImageRequest) GetId() string  { return (*images.UpdateImageRequest)(m).GetId() }
func (m *UpdateImageRequest) GetImage() *Image {
	return (*Image)((*images.UpdateImageRequest)(m).GetImage())
}

// DeleteImageRequest from public import images/images.proto
type DeleteImageRequest images.DeleteImageRequest

func (m *DeleteImageRequest) Reset()         { (*images.DeleteImageRequest)(m).Reset() }
func (m *DeleteImageRequest) String() string { return (*images.DeleteImageRequest)(m).String() }
func (*DeleteImageRequest) ProtoMessage()    {}
func (m *DeleteImageRequest) GetId() string  { return (*images.DeleteImageRequest)(m).GetId() }

// Project from public import projects/projects.proto
type Project projects.Project

func (m *Project) Reset()                    { (*projects.Project)(m).Reset() }
func (m *Project) String() string            { return (*projects.Project)(m).String() }
func (*Project) ProtoMessage()               {}
func (m *Project) GetId() string             { return (*projects.Project)(m).GetId() }
func (m *Project) GetOrganizationId() string { return (*projects.Project)(m).GetOrganizationId() }
func (m *Project) GetName() string           { return (*projects.Project)(m).GetName() }

// ListProjectsRequest from public import projects/projects.proto
type ListProjectsRequest projects.ListProjectsRequest

func (m *ListProjectsRequest) Reset()         { (*projects.ListProjectsRequest)(m).Reset() }
func (m *ListProjectsRequest) String() string { return (*projects.ListProjectsRequest)(m).String() }
func (*ListProjectsRequest) ProtoMessage()    {}
func (m *ListProjectsRequest) GetPageSize() int32 {
	return (*projects.ListProjectsRequest)(m).GetPageSize()
}
func (m *ListProjectsRequest) GetPageToken() string {
	return (*projects.ListProjectsRequest)(m).GetPageToken()
}

// ListProjectsResponse from public import projects/projects.proto
type ListProjectsResponse projects.ListProjectsResponse

func (m *ListProjectsResponse) Reset()         { (*projects.ListProjectsResponse)(m).Reset() }
func (m *ListProjectsResponse) String() string { return (*projects.ListProjectsResponse)(m).String() }
func (*ListProjectsResponse) ProtoMessage()    {}
func (m *ListProjectsResponse) GetProjects() []*Project {
	o := (*projects.ListProjectsResponse)(m).GetProjects()
	if o == nil {
		return nil
	}
	s := make([]*Project, len(o))
	for i, x := range o {
		s[i] = (*Project)(x)
	}
	return s
}
func (m *ListProjectsResponse) GetNextPageToken() string {
	return (*projects.ListProjectsResponse)(m).GetNextPageToken()
}

// GetProjectRequest from public import projects/projects.proto
type GetProjectRequest projects.GetProjectRequest

func (m *GetProjectRequest) Reset()         { (*projects.GetProjectRequest)(m).Reset() }
func (m *GetProjectRequest) String() string { return (*projects.GetProjectRequest)(m).String() }
func (*GetProjectRequest) ProtoMessage()    {}
func (m *GetProjectRequest) GetId() string  { return (*projects.GetProjectRequest)(m).GetId() }

// CreateProjectRequest from public import projects/projects.proto
type CreateProjectRequest projects.CreateProjectRequest

func (m *CreateProjectRequest) Reset()         { (*projects.CreateProjectRequest)(m).Reset() }
func (m *CreateProjectRequest) String() string { return (*projects.CreateProjectRequest)(m).String() }
func (*CreateProjectRequest) ProtoMessage()    {}
func (m *CreateProjectRequest) GetProject() *Project {
	return (*Project)((*projects.CreateProjectRequest)(m).GetProject())
}

// UpdateProjectRequest from public import projects/projects.proto
type UpdateProjectRequest projects.UpdateProjectRequest

func (m *UpdateProjectRequest) Reset()         { (*projects.UpdateProjectRequest)(m).Reset() }
func (m *UpdateProjectRequest) String() string { return (*projects.UpdateProjectRequest)(m).String() }
func (*UpdateProjectRequest) ProtoMessage()    {}
func (m *UpdateProjectRequest) GetId() string  { return (*projects.UpdateProjectRequest)(m).GetId() }
func (m *UpdateProjectRequest) GetProject() *Project {
	return (*Project)((*projects.UpdateProjectRequest)(m).GetProject())
}

// DeleteProjectRequest from public import projects/projects.proto
type DeleteProjectRequest projects.DeleteProjectRequest

func (m *DeleteProjectRequest) Reset()         { (*projects.DeleteProjectRequest)(m).Reset() }
func (m *DeleteProjectRequest) String() string { return (*projects.DeleteProjectRequest)(m).String() }
func (*DeleteProjectRequest) ProtoMessage()    {}
func (m *DeleteProjectRequest) GetId() string  { return (*projects.DeleteProjectRequest)(m).GetId() }

// Organization from public import organizations/organizations.proto
type Organization organizations.Organization

func (m *Organization) Reset()          { (*organizations.Organization)(m).Reset() }
func (m *Organization) String() string  { return (*organizations.Organization)(m).String() }
func (*Organization) ProtoMessage()     {}
func (m *Organization) GetId() string   { return (*organizations.Organization)(m).GetId() }
func (m *Organization) GetName() string { return (*organizations.Organization)(m).GetName() }

// ListOrganizationsRequest from public import organizations/organizations.proto
type ListOrganizationsRequest organizations.ListOrganizationsRequest

func (m *ListOrganizationsRequest) Reset() { (*organizations.ListOrganizationsRequest)(m).Reset() }
func (m *ListOrganizationsRequest) String() string {
	return (*organizations.ListOrganizationsRequest)(m).String()
}
func (*ListOrganizationsRequest) ProtoMessage() {}
func (m *ListOrganizationsRequest) GetPageSize() int32 {
	return (*organizations.ListOrganizationsRequest)(m).GetPageSize()
}
func (m *ListOrganizationsRequest) GetPageToken() string {
	return (*organizations.ListOrganizationsRequest)(m).GetPageToken()
}

// ListOrganizationsResponse from public import organizations/organizations.proto
type ListOrganizationsResponse organizations.ListOrganizationsResponse

func (m *ListOrganizationsResponse) Reset() { (*organizations.ListOrganizationsResponse)(m).Reset() }
func (m *ListOrganizationsResponse) String() string {
	return (*organizations.ListOrganizationsResponse)(m).String()
}
func (*ListOrganizationsResponse) ProtoMessage() {}
func (m *ListOrganizationsResponse) GetOrganizations() []*Organization {
	o := (*organizations.ListOrganizationsResponse)(m).GetOrganizations()
	if o == nil {
		return nil
	}
	s := make([]*Organization, len(o))
	for i, x := range o {
		s[i] = (*Organization)(x)
	}
	return s
}
func (m *ListOrganizationsResponse) GetNextPageToken() string {
	return (*organizations.ListOrganizationsResponse)(m).GetNextPageToken()
}

// GetOrganizationRequest from public import organizations/organizations.proto
type GetOrganizationRequest organizations.GetOrganizationRequest

func (m *GetOrganizationRequest) Reset() { (*organizations.GetOrganizationRequest)(m).Reset() }
func (m *GetOrganizationRequest) String() string {
	return (*organizations.GetOrganizationRequest)(m).String()
}
func (*GetOrganizationRequest) ProtoMessage() {}
func (m *GetOrganizationRequest) GetId() string {
	return (*organizations.GetOrganizationRequest)(m).GetId()
}

// CreateOrganizationRequest from public import organizations/organizations.proto
type CreateOrganizationRequest organizations.CreateOrganizationRequest

func (m *CreateOrganizationRequest) Reset() { (*organizations.CreateOrganizationRequest)(m).Reset() }
func (m *CreateOrganizationRequest) String() string {
	return (*organizations.CreateOrganizationRequest)(m).String()
}
func (*CreateOrganizationRequest) ProtoMessage() {}
func (m *CreateOrganizationRequest) GetOrganization() *Organization {
	return (*Organization)((*organizations.CreateOrganizationRequest)(m).GetOrganization())
}

// UpdateOrganizationRequest from public import organizations/organizations.proto
type UpdateOrganizationRequest organizations.UpdateOrganizationRequest

func (m *UpdateOrganizationRequest) Reset() { (*organizations.UpdateOrganizationRequest)(m).Reset() }
func (m *UpdateOrganizationRequest) String() string {
	return (*organizations.UpdateOrganizationRequest)(m).String()
}
func (*UpdateOrganizationRequest) ProtoMessage() {}
func (m *UpdateOrganizationRequest) GetId() string {
	return (*organizations.UpdateOrganizationRequest)(m).GetId()
}
func (m *UpdateOrganizationRequest) GetOrganization() *Organization {
	return (*Organization)((*organizations.UpdateOrganizationRequest)(m).GetOrganization())
}

// DeleteOrganizationRequest from public import organizations/organizations.proto
type DeleteOrganizationRequest organizations.DeleteOrganizationRequest

func (m *DeleteOrganizationRequest) Reset() { (*organizations.DeleteOrganizationRequest)(m).Reset() }
func (m *DeleteOrganizationRequest) String() string {
	return (*organizations.DeleteOrganizationRequest)(m).String()
}
func (*DeleteOrganizationRequest) ProtoMessage() {}
func (m *DeleteOrganizationRequest) GetId() string {
	return (*organizations.DeleteOrganizationRequest)(m).GetId()
}

// RenderParameterRequest from public import render_parameters/render_parameters.proto
type RenderParameterRequest render_parameters.RenderParameterRequest

func (m *RenderParameterRequest) Reset() { (*render_parameters.RenderParameterRequest)(m).Reset() }
func (m *RenderParameterRequest) String() string {
	return (*render_parameters.RenderParameterRequest)(m).String()
}
func (*RenderParameterRequest) ProtoMessage() {}
func (m *RenderParameterRequest) GetOperatorKey() string {
	return (*render_parameters.RenderParameterRequest)(m).GetOperatorKey()
}
func (m *RenderParameterRequest) GetTime() float64 {
	return (*render_parameters.RenderParameterRequest)(m).GetTime()
}

// RenderParameter from public import render_parameters/render_parameters.proto
type RenderParameter render_parameters.RenderParameter

func (m *RenderParameter) Reset()          { (*render_parameters.RenderParameter)(m).Reset() }
func (m *RenderParameter) String() string  { return (*render_parameters.RenderParameter)(m).String() }
func (*RenderParameter) ProtoMessage()     {}
func (m *RenderParameter) GetType() string { return (*render_parameters.RenderParameter)(m).GetType() }
func (m *RenderParameter) GetConfiguration() []byte {
	return (*render_parameters.RenderParameter)(m).GetConfiguration()
}
func (m *RenderParameter) GetInputs() []*RenderParameter {
	o := (*render_parameters.RenderParameter)(m).GetInputs()
	if o == nil {
		return nil
	}
	s := make([]*RenderParameter, len(o))
	for i, x := range o {
		s[i] = (*RenderParameter)(x)
	}
	return s
}

// RenderRequest from public import renders/renders.proto
type RenderRequest renders.RenderRequest

func (m *RenderRequest) Reset()                 { (*renders.RenderRequest)(m).Reset() }
func (m *RenderRequest) String() string         { return (*renders.RenderRequest)(m).String() }
func (*RenderRequest) ProtoMessage()            {}
func (m *RenderRequest) GetOperatorKey() string { return (*renders.RenderRequest)(m).GetOperatorKey() }
func (m *RenderRequest) GetTime() float64       { return (*renders.RenderRequest)(m).GetTime() }

// RenderResponse from public import renders/renders.proto
type RenderResponse renders.RenderResponse

func (m *RenderResponse) Reset()                { (*renders.RenderResponse)(m).Reset() }
func (m *RenderResponse) String() string        { return (*renders.RenderResponse)(m).String() }
func (*RenderResponse) ProtoMessage()           {}
func (m *RenderResponse) GetResourceId() string { return (*renders.RenderResponse)(m).GetResourceId() }

// Scene from public import scenes/scenes.proto
type Scene scenes.Scene

func (m *Scene) Reset()               { (*scenes.Scene)(m).Reset() }
func (m *Scene) String() string       { return (*scenes.Scene)(m).String() }
func (*Scene) ProtoMessage()          {}
func (m *Scene) GetId() string        { return (*scenes.Scene)(m).GetId() }
func (m *Scene) GetProjectId() string { return (*scenes.Scene)(m).GetProjectId() }
func (m *Scene) GetVersion() string   { return (*scenes.Scene)(m).GetVersion() }
func (m *Scene) GetName() string      { return (*scenes.Scene)(m).GetName() }
func (m *Scene) GetOperators() []byte { return (*scenes.Scene)(m).GetOperators() }

// ListScenesRequest from public import scenes/scenes.proto
type ListScenesRequest scenes.ListScenesRequest

func (m *ListScenesRequest) Reset()             { (*scenes.ListScenesRequest)(m).Reset() }
func (m *ListScenesRequest) String() string     { return (*scenes.ListScenesRequest)(m).String() }
func (*ListScenesRequest) ProtoMessage()        {}
func (m *ListScenesRequest) GetPageSize() int32 { return (*scenes.ListScenesRequest)(m).GetPageSize() }
func (m *ListScenesRequest) GetPageToken() string {
	return (*scenes.ListScenesRequest)(m).GetPageToken()
}

// ListScenesResponse from public import scenes/scenes.proto
type ListScenesResponse scenes.ListScenesResponse

func (m *ListScenesResponse) Reset()         { (*scenes.ListScenesResponse)(m).Reset() }
func (m *ListScenesResponse) String() string { return (*scenes.ListScenesResponse)(m).String() }
func (*ListScenesResponse) ProtoMessage()    {}
func (m *ListScenesResponse) GetScenes() []*Scene {
	o := (*scenes.ListScenesResponse)(m).GetScenes()
	if o == nil {
		return nil
	}
	s := make([]*Scene, len(o))
	for i, x := range o {
		s[i] = (*Scene)(x)
	}
	return s
}
func (m *ListScenesResponse) GetNextPageToken() string {
	return (*scenes.ListScenesResponse)(m).GetNextPageToken()
}

// GetSceneRequest from public import scenes/scenes.proto
type GetSceneRequest scenes.GetSceneRequest

func (m *GetSceneRequest) Reset()         { (*scenes.GetSceneRequest)(m).Reset() }
func (m *GetSceneRequest) String() string { return (*scenes.GetSceneRequest)(m).String() }
func (*GetSceneRequest) ProtoMessage()    {}
func (m *GetSceneRequest) GetId() string  { return (*scenes.GetSceneRequest)(m).GetId() }

// CreateSceneRequest from public import scenes/scenes.proto
type CreateSceneRequest scenes.CreateSceneRequest

func (m *CreateSceneRequest) Reset()         { (*scenes.CreateSceneRequest)(m).Reset() }
func (m *CreateSceneRequest) String() string { return (*scenes.CreateSceneRequest)(m).String() }
func (*CreateSceneRequest) ProtoMessage()    {}
func (m *CreateSceneRequest) GetScene() *Scene {
	return (*Scene)((*scenes.CreateSceneRequest)(m).GetScene())
}

// UpdateSceneRequest from public import scenes/scenes.proto
type UpdateSceneRequest scenes.UpdateSceneRequest

func (m *UpdateSceneRequest) Reset()         { (*scenes.UpdateSceneRequest)(m).Reset() }
func (m *UpdateSceneRequest) String() string { return (*scenes.UpdateSceneRequest)(m).String() }
func (*UpdateSceneRequest) ProtoMessage()    {}
func (m *UpdateSceneRequest) GetId() string  { return (*scenes.UpdateSceneRequest)(m).GetId() }
func (m *UpdateSceneRequest) GetScene() *Scene {
	return (*Scene)((*scenes.UpdateSceneRequest)(m).GetScene())
}

// DeleteSceneRequest from public import scenes/scenes.proto
type DeleteSceneRequest scenes.DeleteSceneRequest

func (m *DeleteSceneRequest) Reset()         { (*scenes.DeleteSceneRequest)(m).Reset() }
func (m *DeleteSceneRequest) String() string { return (*scenes.DeleteSceneRequest)(m).String() }
func (*DeleteSceneRequest) ProtoMessage()    {}
func (m *DeleteSceneRequest) GetId() string  { return (*scenes.DeleteSceneRequest)(m).GetId() }

func init() { proto.RegisterFile("api-web/api-web.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8e, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x86, 0xd5, 0x6a, 0x07, 0x47, 0xa5, 0x08, 0xdd, 0x1c, 0x1c, 0x1c, 0x34, 0x83, 0x2f, 0x72,
	0x6f, 0x50, 0xae, 0xf1, 0x28, 0x11, 0x9a, 0xc4, 0xbb, 0x88, 0xe0, 0xd3, 0x1b, 0x9a, 0x44, 0x50,
	0xa7, 0xff, 0xee, 0xfb, 0xbf, 0xe1, 0x5f, 0x37, 0xe8, 0xcd, 0xe9, 0x49, 0xbd, 0xca, 0x79, 0xf6,
	0xec, 0x82, 0xdb, 0x54, 0xf1, 0x6c, 0x0f, 0x66, 0xc4, 0x81, 0x3a, 0xa1, 0xfb, 0x83, 0xac, 0x26,
	0x51, 0x3f, 0x7f, 0x72, 0xdb, 0xed, 0x84, 0x73, 0x5b, 0xe0, 0x2e, 0xc6, 0x8d, 0x74, 0x10, 0x55,
	0x8e, 0x5c, 0xec, 0x1d, 0x0f, 0x68, 0xcd, 0x0b, 0x83, 0x71, 0x56, 0xd4, 0xd7, 0x97, 0x95, 0x23,
	0x93, 0xbd, 0x12, 0x77, 0x1e, 0x19, 0x47, 0x0a, 0xc4, 0xa2, 0xfe, 0x48, 0x56, 0x9b, 0x54, 0x14,
	0xe1, 0x33, 0x49, 0x34, 0xd9, 0x38, 0x29, 0x45, 0x82, 0x30, 0x83, 0x39, 0x2c, 0xa0, 0x82, 0x25,
	0xac, 0xa0, 0xee, 0xeb, 0x09, 0x5e, 0xde, 0x01, 0x00, 0x00, 0xff, 0xff, 0x95, 0x48, 0x93, 0x0e,
	0x01, 0x01, 0x00, 0x00,
}
