// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scenes/scenes.proto

/*
Package scenes is a generated protocol buffer package.

It is generated from these files:
	scenes/scenes.proto

It has these top-level messages:
	Scene
	ListScenesRequest
	ListScenesResponse
	GetSceneRequest
	CreateSceneRequest
	UpdateSceneRequest
	DeleteSceneRequest
*/
package scenes

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

type Scene struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Scene) Reset()                    { *m = Scene{} }
func (m *Scene) String() string            { return proto.CompactTextString(m) }
func (*Scene) ProtoMessage()               {}
func (*Scene) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Scene) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Scene) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Scene) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListScenesRequest struct {
	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListScenesRequest) Reset()                    { *m = ListScenesRequest{} }
func (m *ListScenesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListScenesRequest) ProtoMessage()               {}
func (*ListScenesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListScenesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListScenesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListScenesResponse struct {
	Scenes        []*Scene `protobuf:"bytes,1,rep,name=scenes" json:"scenes,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListScenesResponse) Reset()                    { *m = ListScenesResponse{} }
func (m *ListScenesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListScenesResponse) ProtoMessage()               {}
func (*ListScenesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListScenesResponse) GetScenes() []*Scene {
	if m != nil {
		return m.Scenes
	}
	return nil
}

func (m *ListScenesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetSceneRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetSceneRequest) Reset()                    { *m = GetSceneRequest{} }
func (m *GetSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSceneRequest) ProtoMessage()               {}
func (*GetSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetSceneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateSceneRequest struct {
	Scene *Scene `protobuf:"bytes,1,opt,name=scene" json:"scene,omitempty"`
}

func (m *CreateSceneRequest) Reset()                    { *m = CreateSceneRequest{} }
func (m *CreateSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSceneRequest) ProtoMessage()               {}
func (*CreateSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateSceneRequest) GetScene() *Scene {
	if m != nil {
		return m.Scene
	}
	return nil
}

type UpdateSceneRequest struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Scene *Scene `protobuf:"bytes,2,opt,name=scene" json:"scene,omitempty"`
}

func (m *UpdateSceneRequest) Reset()                    { *m = UpdateSceneRequest{} }
func (m *UpdateSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSceneRequest) ProtoMessage()               {}
func (*UpdateSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateSceneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateSceneRequest) GetScene() *Scene {
	if m != nil {
		return m.Scene
	}
	return nil
}

type DeleteSceneRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteSceneRequest) Reset()                    { *m = DeleteSceneRequest{} }
func (m *DeleteSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSceneRequest) ProtoMessage()               {}
func (*DeleteSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteSceneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Scene)(nil), "scenes.Scene")
	proto.RegisterType((*ListScenesRequest)(nil), "scenes.ListScenesRequest")
	proto.RegisterType((*ListScenesResponse)(nil), "scenes.ListScenesResponse")
	proto.RegisterType((*GetSceneRequest)(nil), "scenes.GetSceneRequest")
	proto.RegisterType((*CreateSceneRequest)(nil), "scenes.CreateSceneRequest")
	proto.RegisterType((*UpdateSceneRequest)(nil), "scenes.UpdateSceneRequest")
	proto.RegisterType((*DeleteSceneRequest)(nil), "scenes.DeleteSceneRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scenes service

type ScenesClient interface {
	List(ctx context.Context, in *ListScenesRequest, opts ...grpc.CallOption) (*ListScenesResponse, error)
	Get(ctx context.Context, in *GetSceneRequest, opts ...grpc.CallOption) (*Scene, error)
	Create(ctx context.Context, in *CreateSceneRequest, opts ...grpc.CallOption) (*Scene, error)
	Update(ctx context.Context, in *UpdateSceneRequest, opts ...grpc.CallOption) (*Scene, error)
	Delete(ctx context.Context, in *DeleteSceneRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type scenesClient struct {
	cc *grpc.ClientConn
}

func NewScenesClient(cc *grpc.ClientConn) ScenesClient {
	return &scenesClient{cc}
}

func (c *scenesClient) List(ctx context.Context, in *ListScenesRequest, opts ...grpc.CallOption) (*ListScenesResponse, error) {
	out := new(ListScenesResponse)
	err := grpc.Invoke(ctx, "/scenes.Scenes/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Get(ctx context.Context, in *GetSceneRequest, opts ...grpc.CallOption) (*Scene, error) {
	out := new(Scene)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Create(ctx context.Context, in *CreateSceneRequest, opts ...grpc.CallOption) (*Scene, error) {
	out := new(Scene)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Update(ctx context.Context, in *UpdateSceneRequest, opts ...grpc.CallOption) (*Scene, error) {
	out := new(Scene)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Delete(ctx context.Context, in *DeleteSceneRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scenes service

type ScenesServer interface {
	List(context.Context, *ListScenesRequest) (*ListScenesResponse, error)
	Get(context.Context, *GetSceneRequest) (*Scene, error)
	Create(context.Context, *CreateSceneRequest) (*Scene, error)
	Update(context.Context, *UpdateSceneRequest) (*Scene, error)
	Delete(context.Context, *DeleteSceneRequest) (*google_protobuf1.Empty, error)
}

func RegisterScenesServer(s *grpc.Server, srv ScenesServer) {
	s.RegisterService(&_Scenes_serviceDesc, srv)
}

func _Scenes_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScenesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).List(ctx, req.(*ListScenesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Get(ctx, req.(*GetSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Create(ctx, req.(*CreateSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Update(ctx, req.(*UpdateSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Delete(ctx, req.(*DeleteSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scenes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scenes.Scenes",
	HandlerType: (*ScenesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Scenes_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Scenes_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Scenes_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Scenes_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Scenes_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scenes/scenes.proto",
}

func init() { proto.RegisterFile("scenes/scenes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xdb, 0xce, 0xd2, 0x40,
	0x10, 0x80, 0x43, 0x0b, 0x8d, 0x8c, 0x22, 0x61, 0x48, 0x04, 0x8a, 0x1a, 0x5d, 0x0f, 0x21, 0x5e,
	0xb4, 0x11, 0xaf, 0xf4, 0xd6, 0x53, 0x30, 0xc6, 0x43, 0x91, 0x6b, 0x52, 0xe8, 0x48, 0xaa, 0xd0,
	0x56, 0xba, 0x18, 0x0f, 0xf1, 0xc6, 0x57, 0xf0, 0x01, 0x7c, 0xa8, 0xff, 0x15, 0xfe, 0x07, 0xf9,
	0xb7, 0xbb, 0x6d, 0x81, 0xed, 0x1f, 0xae, 0xe8, 0xce, 0xcc, 0x7e, 0x33, 0xfb, 0x2d, 0x0b, 0xdd,
	0x74, 0x49, 0x11, 0xa5, 0xae, 0xfa, 0x71, 0x92, 0x6d, 0xcc, 0x63, 0xb4, 0xd4, 0xca, 0xbe, 0xb9,
	0x8a, 0xe3, 0xd5, 0x9a, 0x5c, 0x3f, 0x09, 0x5d, 0x3f, 0x8a, 0x62, 0xee, 0xf3, 0x30, 0x8e, 0xf2,
	0x2a, 0x7b, 0x98, 0x67, 0xe5, 0x6a, 0xb1, 0xfb, 0xec, 0xd2, 0x26, 0xe1, 0x3f, 0x55, 0x92, 0xbd,
	0x81, 0xc6, 0x34, 0x83, 0xe0, 0x75, 0x30, 0xc2, 0xa0, 0x5f, 0xbb, 0x53, 0x1b, 0x35, 0x3d, 0xf1,
	0x85, 0xb7, 0x00, 0x44, 0xc5, 0x17, 0x5a, 0xf2, 0xb9, 0x88, 0x1b, 0x32, 0xde, 0xcc, 0x23, 0x93,
	0x00, 0x11, 0xea, 0x91, 0xbf, 0xa1, 0xbe, 0x29, 0x13, 0xf2, 0x9b, 0xbd, 0x87, 0xce, 0xdb, 0x30,
	0xe5, 0x92, 0x97, 0x7a, 0xf4, 0x6d, 0x47, 0x29, 0xc7, 0x21, 0x34, 0x13, 0x7f, 0x45, 0xf3, 0x34,
	0xfc, 0x45, 0x12, 0xdf, 0xf0, 0xae, 0x64, 0x81, 0xa9, 0x58, 0xcb, 0x26, 0x59, 0x92, 0xc7, 0x5f,
	0x29, 0x2a, 0x9b, 0x88, 0xc8, 0xa7, 0x2c, 0xc0, 0x96, 0x80, 0x87, 0xc0, 0x34, 0x11, 0x87, 0x22,
	0x7c, 0x00, 0xf9, 0xb9, 0x05, 0xce, 0x1c, 0x5d, 0x1d, 0xb7, 0x9c, 0x5c, 0x8a, 0xac, 0xf3, 0xf2,
	0x24, 0x3e, 0x84, 0x76, 0x44, 0x3f, 0xf8, 0xbc, 0xd2, 0xa0, 0x95, 0x85, 0x3f, 0x94, 0x4d, 0xee,
	0x42, 0xfb, 0x35, 0xa9, 0x1e, 0xc5, 0xcc, 0x9a, 0x0b, 0xf6, 0x14, 0xf0, 0xf9, 0x96, 0x7c, 0x4e,
	0x47, 0x55, 0xf7, 0xa0, 0x21, 0x5b, 0xc9, 0xc2, 0xca, 0x18, 0x2a, 0xc7, 0x26, 0x80, 0xb3, 0x24,
	0xd0, 0xb7, 0xea, 0xb2, 0x4b, 0x94, 0x71, 0x02, 0x75, 0x1f, 0xf0, 0x05, 0xad, 0xe9, 0x34, 0x6a,
	0xfc, 0xdf, 0x04, 0x4b, 0x09, 0xc3, 0x8f, 0x50, 0xcf, 0xf4, 0xe1, 0xa0, 0xc0, 0x55, 0x6e, 0xc7,
	0xb6, 0x2f, 0x4b, 0x29, 0xcf, 0x0c, 0xff, 0x9e, 0x9d, 0xff, 0x33, 0xae, 0x21, 0xb8, 0xdf, 0x1f,
	0xe7, 0xff, 0x3b, 0x7c, 0x05, 0xa6, 0x90, 0x85, 0xbd, 0x62, 0x9b, 0x66, 0xce, 0x3e, 0x9e, 0x9c,
	0xf5, 0x24, 0xa2, 0x83, 0xed, 0x3d, 0xc2, 0xfd, 0x1d, 0x06, 0x7f, 0xf0, 0x1d, 0x58, 0xca, 0x28,
	0x96, 0x13, 0x54, 0x0d, 0xeb, 0xb4, 0x81, 0xa4, 0x75, 0xd9, 0xc1, 0x40, 0xcf, 0x94, 0x1b, 0x9c,
	0x82, 0xa5, 0x34, 0xef, 0x79, 0x55, 0xed, 0x3a, 0xef, 0xb6, 0xe4, 0xf5, 0x6d, 0x7d, 0xba, 0x02,
	0x3a, 0x03, 0x4b, 0x09, 0xdf, 0x43, 0xab, 0x17, 0x60, 0xdf, 0x70, 0xd4, 0xfb, 0x72, 0x8a, 0xf7,
	0xe5, 0xbc, 0xcc, 0xde, 0x57, 0x71, 0xf6, 0x47, 0x3a, 0x7d, 0x61, 0xc9, 0xc2, 0x27, 0x17, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x67, 0x49, 0xcd, 0xd3, 0x03, 0x00, 0x00,
}
