// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image_sequences/image_sequences.proto

/*
Package image_sequences is a generated protocol buffer package.

It is generated from these files:
	image_sequences/image_sequences.proto

It has these top-level messages:
	ImageSequence
	ListImageSequencesRequest
	ListImageSequencesResponse
	GetImageSequenceRequest
	CreateImageSequenceRequest
	UpdateImageSequenceRequest
	DeleteImageSequenceRequest
*/
package image_sequences

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ImageSequence struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ImageSequence) Reset()                    { *m = ImageSequence{} }
func (m *ImageSequence) String() string            { return proto.CompactTextString(m) }
func (*ImageSequence) ProtoMessage()               {}
func (*ImageSequence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImageSequence) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ImageSequence) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ImageSequence) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListImageSequencesRequest struct {
	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListImageSequencesRequest) Reset()                    { *m = ListImageSequencesRequest{} }
func (m *ListImageSequencesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListImageSequencesRequest) ProtoMessage()               {}
func (*ListImageSequencesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListImageSequencesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListImageSequencesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListImageSequencesResponse struct {
	ImageSequences []*ImageSequence `protobuf:"bytes,1,rep,name=image_sequences,json=imageSequences" json:"image_sequences,omitempty"`
	NextPageToken  string           `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListImageSequencesResponse) Reset()                    { *m = ListImageSequencesResponse{} }
func (m *ListImageSequencesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListImageSequencesResponse) ProtoMessage()               {}
func (*ListImageSequencesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListImageSequencesResponse) GetImageSequences() []*ImageSequence {
	if m != nil {
		return m.ImageSequences
	}
	return nil
}

func (m *ListImageSequencesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetImageSequenceRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetImageSequenceRequest) Reset()                    { *m = GetImageSequenceRequest{} }
func (m *GetImageSequenceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetImageSequenceRequest) ProtoMessage()               {}
func (*GetImageSequenceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetImageSequenceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateImageSequenceRequest struct {
	ImageSequence *ImageSequence `protobuf:"bytes,1,opt,name=image_sequence,json=imageSequence" json:"image_sequence,omitempty"`
}

func (m *CreateImageSequenceRequest) Reset()                    { *m = CreateImageSequenceRequest{} }
func (m *CreateImageSequenceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateImageSequenceRequest) ProtoMessage()               {}
func (*CreateImageSequenceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateImageSequenceRequest) GetImageSequence() *ImageSequence {
	if m != nil {
		return m.ImageSequence
	}
	return nil
}

type UpdateImageSequenceRequest struct {
	Id            string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ImageSequence *ImageSequence `protobuf:"bytes,2,opt,name=image_sequence,json=imageSequence" json:"image_sequence,omitempty"`
}

func (m *UpdateImageSequenceRequest) Reset()                    { *m = UpdateImageSequenceRequest{} }
func (m *UpdateImageSequenceRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateImageSequenceRequest) ProtoMessage()               {}
func (*UpdateImageSequenceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateImageSequenceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateImageSequenceRequest) GetImageSequence() *ImageSequence {
	if m != nil {
		return m.ImageSequence
	}
	return nil
}

type DeleteImageSequenceRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteImageSequenceRequest) Reset()                    { *m = DeleteImageSequenceRequest{} }
func (m *DeleteImageSequenceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteImageSequenceRequest) ProtoMessage()               {}
func (*DeleteImageSequenceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteImageSequenceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageSequence)(nil), "image_sequences.ImageSequence")
	proto.RegisterType((*ListImageSequencesRequest)(nil), "image_sequences.ListImageSequencesRequest")
	proto.RegisterType((*ListImageSequencesResponse)(nil), "image_sequences.ListImageSequencesResponse")
	proto.RegisterType((*GetImageSequenceRequest)(nil), "image_sequences.GetImageSequenceRequest")
	proto.RegisterType((*CreateImageSequenceRequest)(nil), "image_sequences.CreateImageSequenceRequest")
	proto.RegisterType((*UpdateImageSequenceRequest)(nil), "image_sequences.UpdateImageSequenceRequest")
	proto.RegisterType((*DeleteImageSequenceRequest)(nil), "image_sequences.DeleteImageSequenceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ImageSequences service

type ImageSequencesClient interface {
	List(ctx context.Context, in *ListImageSequencesRequest, opts ...grpc.CallOption) (*ListImageSequencesResponse, error)
	Get(ctx context.Context, in *GetImageSequenceRequest, opts ...grpc.CallOption) (*ImageSequence, error)
	Create(ctx context.Context, in *CreateImageSequenceRequest, opts ...grpc.CallOption) (*ImageSequence, error)
	Update(ctx context.Context, in *UpdateImageSequenceRequest, opts ...grpc.CallOption) (*ImageSequence, error)
	Delete(ctx context.Context, in *DeleteImageSequenceRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type imageSequencesClient struct {
	cc *grpc.ClientConn
}

func NewImageSequencesClient(cc *grpc.ClientConn) ImageSequencesClient {
	return &imageSequencesClient{cc}
}

func (c *imageSequencesClient) List(ctx context.Context, in *ListImageSequencesRequest, opts ...grpc.CallOption) (*ListImageSequencesResponse, error) {
	out := new(ListImageSequencesResponse)
	err := grpc.Invoke(ctx, "/image_sequences.ImageSequences/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageSequencesClient) Get(ctx context.Context, in *GetImageSequenceRequest, opts ...grpc.CallOption) (*ImageSequence, error) {
	out := new(ImageSequence)
	err := grpc.Invoke(ctx, "/image_sequences.ImageSequences/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageSequencesClient) Create(ctx context.Context, in *CreateImageSequenceRequest, opts ...grpc.CallOption) (*ImageSequence, error) {
	out := new(ImageSequence)
	err := grpc.Invoke(ctx, "/image_sequences.ImageSequences/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageSequencesClient) Update(ctx context.Context, in *UpdateImageSequenceRequest, opts ...grpc.CallOption) (*ImageSequence, error) {
	out := new(ImageSequence)
	err := grpc.Invoke(ctx, "/image_sequences.ImageSequences/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageSequencesClient) Delete(ctx context.Context, in *DeleteImageSequenceRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/image_sequences.ImageSequences/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageSequences service

type ImageSequencesServer interface {
	List(context.Context, *ListImageSequencesRequest) (*ListImageSequencesResponse, error)
	Get(context.Context, *GetImageSequenceRequest) (*ImageSequence, error)
	Create(context.Context, *CreateImageSequenceRequest) (*ImageSequence, error)
	Update(context.Context, *UpdateImageSequenceRequest) (*ImageSequence, error)
	Delete(context.Context, *DeleteImageSequenceRequest) (*google_protobuf1.Empty, error)
}

func RegisterImageSequencesServer(s *grpc.Server, srv ImageSequencesServer) {
	s.RegisterService(&_ImageSequences_serviceDesc, srv)
}

func _ImageSequences_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImageSequencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageSequencesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_sequences.ImageSequences/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageSequencesServer).List(ctx, req.(*ListImageSequencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageSequences_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageSequencesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_sequences.ImageSequences/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageSequencesServer).Get(ctx, req.(*GetImageSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageSequences_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageSequencesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_sequences.ImageSequences/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageSequencesServer).Create(ctx, req.(*CreateImageSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageSequences_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageSequencesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_sequences.ImageSequences/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageSequencesServer).Update(ctx, req.(*UpdateImageSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageSequences_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageSequencesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_sequences.ImageSequences/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageSequencesServer).Delete(ctx, req.(*DeleteImageSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageSequences_serviceDesc = grpc.ServiceDesc{
	ServiceName: "image_sequences.ImageSequences",
	HandlerType: (*ImageSequencesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ImageSequences_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ImageSequences_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ImageSequences_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ImageSequences_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ImageSequences_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image_sequences/image_sequences.proto",
}

func init() { proto.RegisterFile("image_sequences/image_sequences.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x92, 0x36, 0x22, 0x83, 0x92, 0x4a, 0x83, 0x80, 0xb0, 0x01, 0x54, 0xad, 0x54, 0x14,
	0x52, 0x64, 0x43, 0xb9, 0x71, 0x85, 0xaa, 0xaa, 0xc4, 0x01, 0xb9, 0x20, 0x8e, 0x91, 0x1b, 0x0f,
	0xd1, 0x42, 0xe3, 0x75, 0xeb, 0x2d, 0xe2, 0xf3, 0x40, 0x2f, 0x9c, 0x38, 0xf1, 0xd3, 0xf8, 0x0b,
	0xfc, 0x10, 0x76, 0xd7, 0x4e, 0x25, 0xaf, 0xbd, 0x72, 0xc5, 0x6d, 0xf7, 0xcd, 0x78, 0xde, 0x9b,
	0x99, 0xe7, 0x85, 0x1d, 0xb1, 0x8a, 0x97, 0x34, 0xcf, 0xe9, 0xf4, 0x9c, 0xd2, 0x05, 0xe5, 0xa1,
	0x73, 0x0f, 0xb2, 0x33, 0xa9, 0x24, 0x6e, 0x39, 0x30, 0xbb, 0xbb, 0x94, 0x72, 0x79, 0x42, 0x61,
	0x9c, 0x89, 0x30, 0x4e, 0x53, 0xa9, 0x62, 0x25, 0x64, 0x5a, 0xa6, 0xb3, 0x49, 0x19, 0xb5, 0xb7,
	0xe3, 0xf3, 0x77, 0x21, 0xad, 0x32, 0xf5, 0xb9, 0x08, 0xf2, 0x08, 0x86, 0x87, 0xa6, 0xda, 0x51,
	0x59, 0x0c, 0x47, 0xd0, 0x15, 0xc9, 0xb8, 0xb3, 0xdd, 0x99, 0x0e, 0x22, 0x7d, 0xc2, 0x7b, 0x00,
	0x3a, 0xf3, 0x3d, 0x2d, 0xd4, 0x5c, 0xe3, 0x5d, 0x8b, 0x0f, 0x4a, 0xe4, 0x30, 0x41, 0x84, 0x8d,
	0x34, 0x5e, 0xd1, 0xb8, 0x67, 0x03, 0xf6, 0xcc, 0xdf, 0xc2, 0x9d, 0x97, 0x22, 0x57, 0x95, 0xba,
	0x79, 0x64, 0x0e, 0xb9, 0xc2, 0x09, 0x0c, 0x32, 0xab, 0x5e, 0x7c, 0x21, 0x4b, 0xb3, 0x19, 0x5d,
	0x33, 0xc0, 0x91, 0xbe, 0x5b, 0x32, 0x13, 0x54, 0xf2, 0x03, 0xa5, 0x97, 0x64, 0x1a, 0x79, 0x6d,
	0x00, 0xfe, 0xab, 0x03, 0xac, 0xa9, 0x72, 0x9e, 0xe9, 0x6e, 0x09, 0x0f, 0xc0, 0x9d, 0x8c, 0x26,
	0xe8, 0x4d, 0xaf, 0xef, 0xdd, 0x0f, 0xdc, 0x41, 0x56, 0x2a, 0x44, 0x23, 0x51, 0x29, 0x88, 0x0f,
	0x60, 0x2b, 0xa5, 0x4f, 0x6a, 0x5e, 0xd3, 0x32, 0x34, 0xf0, 0xab, 0x4b, 0x3d, 0x0f, 0xe1, 0xf6,
	0x01, 0x55, 0xd5, 0xac, 0xdb, 0x74, 0xc6, 0xc8, 0x17, 0xc0, 0x9e, 0x9f, 0x51, 0xac, 0xa8, 0x31,
	0x7b, 0x1f, 0x46, 0x55, 0x85, 0xf6, 0xcb, 0x76, 0xe1, 0xc3, 0x8a, 0x70, 0x9e, 0x03, 0x7b, 0x93,
	0x25, 0x3e, 0x12, 0x77, 0xb3, 0x75, 0xd2, 0xee, 0xff, 0x90, 0x3e, 0x02, 0xf6, 0x82, 0x4e, 0xe8,
	0x6a, 0xa4, 0x7b, 0x3f, 0x36, 0x61, 0x54, 0x5d, 0x1f, 0x7e, 0x83, 0x0d, 0xb3, 0x54, 0x9c, 0xd5,
	0x78, 0xbd, 0x2e, 0x62, 0xbb, 0x57, 0xca, 0x2d, 0x7c, 0xc1, 0x27, 0x17, 0x7f, 0xfe, 0xfe, 0xee,
	0xde, 0xc4, 0x1b, 0xe1, 0xc7, 0x27, 0xee, 0x2f, 0x85, 0xa7, 0xd0, 0xd3, 0x3b, 0xc4, 0x69, 0xad,
	0xa0, 0x67, 0xb3, 0xac, 0x65, 0x3c, 0x7c, 0xdb, 0xb2, 0x31, 0x1c, 0x37, 0xb0, 0x85, 0x5f, 0x45,
	0xf2, 0x1d, 0x2f, 0x3a, 0xd0, 0x2f, 0xcc, 0x80, 0xf5, 0x3e, 0xfc, 0x2e, 0x69, 0x65, 0xde, 0xb5,
	0xcc, 0x3b, 0xbc, 0xa9, 0xcf, 0x67, 0xce, 0xae, 0xf1, 0xa7, 0x16, 0x51, 0x98, 0xa5, 0x41, 0x84,
	0xdf, 0x45, 0xad, 0x22, 0x1e, 0x5b, 0x11, 0x33, 0xe6, 0x6d, 0xbf, 0xa6, 0x44, 0x42, 0xbf, 0x30,
	0x50, 0x83, 0x10, 0xbf, 0xb3, 0xd8, 0xad, 0xa0, 0x78, 0xd7, 0x82, 0xf5, 0xbb, 0x16, 0xec, 0x9b,
	0x77, 0x6d, 0x3d, 0xff, 0x99, 0x57, 0xc0, 0x71, 0xdf, 0x7e, 0xf1, 0xf4, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x24, 0x13, 0xc9, 0x1b, 0x6f, 0x05, 0x00, 0x00,
}
