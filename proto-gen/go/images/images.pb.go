// Code generated by protoc-gen-go. DO NOT EDIT.
// source: images/images.proto

/*
Package images is a generated protocol buffer package.

It is generated from these files:
	images/images.proto

It has these top-level messages:
	Frame
	Plane
	Channel
	Row
	ListImagesRequest
	ListImagesResponse
	GetImageRequest
	CreateImageRequest
	UpdateImageRequest
	DeleteImageRequest
*/
package images

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

type Frame struct {
	Name   string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Width  int32             `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height int32             `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Planes []*Plane          `protobuf:"bytes,5,rep,name=planes" json:"planes,omitempty"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Frame) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Frame) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Frame) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Frame) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Frame) GetPlanes() []*Plane {
	if m != nil {
		return m.Planes
	}
	return nil
}

type Plane struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Width    int32             `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height   int32             `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Labels   map[string]string `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Channels []*Channel        `protobuf:"bytes,5,rep,name=channels" json:"channels,omitempty"`
}

func (m *Plane) Reset()                    { *m = Plane{} }
func (m *Plane) String() string            { return proto.CompactTextString(m) }
func (*Plane) ProtoMessage()               {}
func (*Plane) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Plane) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plane) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Plane) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Plane) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Plane) GetChannels() []*Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

type Channel struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Rows []*Row `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
}

func (m *Channel) Reset()                    { *m = Channel{} }
func (m *Channel) String() string            { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()               {}
func (*Channel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Channel) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Row struct {
	Data []float64 `protobuf:"fixed64,1,rep,packed,name=data" json:"data,omitempty"`
}

func (m *Row) Reset()                    { *m = Row{} }
func (m *Row) String() string            { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()               {}
func (*Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Row) GetData() []float64 {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListImagesRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	PageSize       int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken      string `protobuf:"bytes,5,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListImagesRequest) Reset()                    { *m = ListImagesRequest{} }
func (m *ListImagesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListImagesRequest) ProtoMessage()               {}
func (*ListImagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListImagesRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *ListImagesRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *ListImagesRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *ListImagesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListImagesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListImagesResponse struct {
	Images        []*Frame `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListImagesResponse) Reset()                    { *m = ListImagesResponse{} }
func (m *ListImagesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListImagesResponse) ProtoMessage()               {}
func (*ListImagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListImagesResponse) GetImages() []*Frame {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ListImagesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetImageRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	ImageId        string `protobuf:"bytes,4,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
}

func (m *GetImageRequest) Reset()                    { *m = GetImageRequest{} }
func (m *GetImageRequest) String() string            { return proto.CompactTextString(m) }
func (*GetImageRequest) ProtoMessage()               {}
func (*GetImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetImageRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *GetImageRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *GetImageRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *GetImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

type CreateImageRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	ImageId        string `protobuf:"bytes,4,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	Image          *Frame `protobuf:"bytes,5,opt,name=image" json:"image,omitempty"`
}

func (m *CreateImageRequest) Reset()                    { *m = CreateImageRequest{} }
func (m *CreateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateImageRequest) ProtoMessage()               {}
func (*CreateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateImageRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *CreateImageRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *CreateImageRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *CreateImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *CreateImageRequest) GetImage() *Frame {
	if m != nil {
		return m.Image
	}
	return nil
}

type UpdateImageRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	ImageId        string `protobuf:"bytes,4,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	Image          *Frame `protobuf:"bytes,5,opt,name=image" json:"image,omitempty"`
}

func (m *UpdateImageRequest) Reset()                    { *m = UpdateImageRequest{} }
func (m *UpdateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateImageRequest) ProtoMessage()               {}
func (*UpdateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateImageRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *UpdateImageRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *UpdateImageRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *UpdateImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *UpdateImageRequest) GetImage() *Frame {
	if m != nil {
		return m.Image
	}
	return nil
}

type DeleteImageRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	ImageId        string `protobuf:"bytes,4,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
}

func (m *DeleteImageRequest) Reset()                    { *m = DeleteImageRequest{} }
func (m *DeleteImageRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteImageRequest) ProtoMessage()               {}
func (*DeleteImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteImageRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *DeleteImageRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *DeleteImageRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *DeleteImageRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func init() {
	proto.RegisterType((*Frame)(nil), "images.Frame")
	proto.RegisterType((*Plane)(nil), "images.Plane")
	proto.RegisterType((*Channel)(nil), "images.Channel")
	proto.RegisterType((*Row)(nil), "images.Row")
	proto.RegisterType((*ListImagesRequest)(nil), "images.ListImagesRequest")
	proto.RegisterType((*ListImagesResponse)(nil), "images.ListImagesResponse")
	proto.RegisterType((*GetImageRequest)(nil), "images.GetImageRequest")
	proto.RegisterType((*CreateImageRequest)(nil), "images.CreateImageRequest")
	proto.RegisterType((*UpdateImageRequest)(nil), "images.UpdateImageRequest")
	proto.RegisterType((*DeleteImageRequest)(nil), "images.DeleteImageRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Images service

type ImagesClient interface {
	List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Frame, error)
	Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*Frame, error)
	Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*Frame, error)
	Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type imagesClient struct {
	cc *grpc.ClientConn
}

func NewImagesClient(cc *grpc.ClientConn) ImagesClient {
	return &imagesClient{cc}
}

func (c *imagesClient) List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := grpc.Invoke(ctx, "/images.Images/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Frame, error) {
	out := new(Frame)
	err := grpc.Invoke(ctx, "/images.Images/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*Frame, error) {
	out := new(Frame)
	err := grpc.Invoke(ctx, "/images.Images/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*Frame, error) {
	out := new(Frame)
	err := grpc.Invoke(ctx, "/images.Images/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/images.Images/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Images service

type ImagesServer interface {
	List(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	Get(context.Context, *GetImageRequest) (*Frame, error)
	Create(context.Context, *CreateImageRequest) (*Frame, error)
	Update(context.Context, *UpdateImageRequest) (*Frame, error)
	Delete(context.Context, *DeleteImageRequest) (*google_protobuf1.Empty, error)
}

func RegisterImagesServer(s *grpc.Server, srv ImagesServer) {
	s.RegisterService(&_Images_serviceDesc, srv)
}

func _Images_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).List(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).Get(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).Create(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).Update(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).Delete(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Images_serviceDesc = grpc.ServiceDesc{
	ServiceName: "images.Images",
	HandlerType: (*ImagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Images_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Images_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Images_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Images_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Images_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "images/images.proto",
}

func init() { proto.RegisterFile("images/images.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x96, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xe5, 0x38, 0x71, 0xdb, 0x89, 0xfa, 0xa5, 0xdf, 0x80, 0x4a, 0xea, 0x52, 0xb5, 0x32,
	0x02, 0x2a, 0x90, 0x62, 0xb5, 0x5c, 0xa0, 0x07, 0x0e, 0x94, 0x82, 0x2a, 0xf5, 0x50, 0x99, 0xc2,
	0x09, 0x14, 0x6d, 0x93, 0x25, 0xb1, 0x9a, 0x7a, 0x8d, 0xbd, 0x6d, 0x69, 0xab, 0x5c, 0x78, 0x03,
	0x04, 0x42, 0x48, 0x1c, 0x11, 0x12, 0x0f, 0xc0, 0x85, 0xe7, 0xe0, 0xce, 0xa9, 0x0f, 0xc2, 0xee,
	0xac, 0x9d, 0x86, 0x24, 0x37, 0x14, 0x54, 0x4e, 0xf1, 0xcc, 0x7f, 0xc6, 0xf3, 0x9b, 0xd9, 0x9d,
	0xd6, 0x70, 0x29, 0xdc, 0x67, 0x2d, 0x9e, 0xfa, 0xe6, 0xa7, 0x16, 0x27, 0x42, 0x0a, 0x74, 0x8c,
	0xe5, 0x5e, 0x6d, 0x09, 0xd1, 0xea, 0x70, 0x9f, 0xc5, 0xa1, 0xcf, 0xa2, 0x48, 0x48, 0x26, 0x43,
	0x11, 0x65, 0x51, 0xee, 0x7c, 0xa6, 0x92, 0xb5, 0x7b, 0xf0, 0xd2, 0xe7, 0xfb, 0xb1, 0x3c, 0x36,
	0xa2, 0xf7, 0xd3, 0x82, 0xd2, 0xa3, 0x84, 0xed, 0x73, 0x44, 0x28, 0x46, 0xea, 0xb7, 0x6a, 0x2d,
	0x59, 0xcb, 0x53, 0x01, 0x3d, 0xe3, 0x65, 0x28, 0x1d, 0x85, 0x4d, 0xd9, 0xae, 0x16, 0x94, 0xb3,
	0x14, 0x18, 0x03, 0x67, 0xc1, 0x69, 0xf3, 0xb0, 0xd5, 0x96, 0x55, 0x9b, 0xdc, 0x99, 0x85, 0x2b,
	0xe0, 0x74, 0xd8, 0x2e, 0xef, 0xa4, 0xd5, 0xe2, 0x92, 0xbd, 0x5c, 0x5e, 0x9d, 0xab, 0x65, 0xb4,
	0x54, 0xa0, 0xb6, 0x45, 0xda, 0x46, 0x24, 0x93, 0xe3, 0x20, 0x0b, 0xc4, 0xeb, 0xe0, 0xc4, 0x1d,
	0x16, 0xf1, 0xb4, 0x5a, 0xa2, 0x94, 0xe9, 0x3c, 0x65, 0x5b, 0x7b, 0x83, 0x4c, 0x74, 0xef, 0x41,
	0xb9, 0x2f, 0x1b, 0x67, 0xc0, 0xde, 0xe3, 0xc7, 0x19, 0xa9, 0x7e, 0xd4, 0xa0, 0x87, 0xac, 0x73,
	0xc0, 0x09, 0x74, 0x2a, 0x30, 0xc6, 0x5a, 0xe1, 0xae, 0xe5, 0x9d, 0xa9, 0x06, 0xe9, 0x65, 0xe3,
	0x6c, 0x90, 0x0a, 0x8c, 0x6c, 0xf0, 0x36, 0x4c, 0x36, 0xda, 0xea, 0x4c, 0x74, 0x92, 0x69, 0xb1,
	0x92, 0x27, 0xad, 0x1b, 0x7f, 0xd0, 0x0b, 0xf8, 0x93, 0x36, 0xef, 0xc3, 0x44, 0xf6, 0xbe, 0x91,
	0x7d, 0x2e, 0x42, 0x31, 0x11, 0x47, 0xa9, 0xca, 0xd3, 0x08, 0xe5, 0x1c, 0x21, 0x10, 0x47, 0x01,
	0x09, 0xde, 0x02, 0xd8, 0xca, 0x50, 0x9d, 0x17, 0x9b, 0x4c, 0x32, 0x95, 0x6b, 0x2f, 0x5b, 0x0f,
	0x0a, 0x33, 0x56, 0x40, 0xb6, 0xf7, 0xcd, 0x82, 0xff, 0xb7, 0xc2, 0x54, 0x6e, 0x52, 0x5e, 0xc0,
	0x5f, 0x1d, 0xf0, 0x54, 0xe2, 0x4d, 0xa8, 0x88, 0xa4, 0xc5, 0xa2, 0xf0, 0x84, 0x2e, 0x5c, 0x3d,
	0x6c, 0x66, 0x45, 0xff, 0xeb, 0x77, 0x6f, 0x36, 0x71, 0x0e, 0x26, 0xd3, 0x06, 0x8f, 0xb8, 0x8e,
	0x30, 0xe8, 0x13, 0x64, 0x2b, 0x69, 0x11, 0xca, 0x22, 0xe6, 0x09, 0x93, 0x22, 0xd1, 0xaa, 0x4d,
	0x2a, 0xe4, 0x2e, 0x15, 0x30, 0x0f, 0x53, 0xb1, 0x2a, 0x5a, 0x4f, 0xc3, 0x13, 0xae, 0xe6, 0xae,
	0xcf, 0x63, 0x52, 0x3b, 0x9e, 0x28, 0x1b, 0x17, 0x00, 0x48, 0x94, 0x62, 0x8f, 0x47, 0x6a, 0xc0,
	0x3a, 0x99, 0xc2, 0x77, 0xb4, 0xc3, 0x6b, 0x00, 0xf6, 0x53, 0xa7, 0xb1, 0xda, 0x0a, 0xae, 0x2f,
	0x9d, 0xe9, 0x9f, 0xda, 0xec, 0xbb, 0x74, 0x74, 0x4f, 0x83, 0x4c, 0xc4, 0x1b, 0x50, 0x89, 0xf8,
	0x6b, 0x59, 0xef, 0x2b, 0x60, 0xd8, 0xa7, 0xb5, 0x7b, 0xbb, 0x57, 0xe4, 0xad, 0x05, 0x95, 0xc7,
	0xdc, 0x14, 0xf9, 0xab, 0x93, 0x51, 0xb9, 0x84, 0xaa, 0xd5, 0xa2, 0xc9, 0x25, 0x7b, 0xb3, 0xe9,
	0x7d, 0xb7, 0x00, 0xd7, 0x13, 0xce, 0x24, 0xbf, 0x48, 0x58, 0x78, 0x0d, 0x4a, 0xf4, 0x48, 0x27,
	0x35, 0x34, 0x78, 0xa3, 0x11, 0xfb, 0xd3, 0xb8, 0xf9, 0x4f, 0xb2, 0xbf, 0x57, 0xec, 0x0f, 0x79,
	0x87, 0x5f, 0x2c, 0xf6, 0xd5, 0xaf, 0x0e, 0x38, 0x66, 0x09, 0xf0, 0x93, 0x05, 0x45, 0xbd, 0x13,
	0xd8, 0xfb, 0xe3, 0x35, 0xb4, 0xd7, 0xae, 0x3b, 0x4a, 0x32, 0xcb, 0xe3, 0x3d, 0x7f, 0xf3, 0xe3,
	0xec, 0x5d, 0xe1, 0x19, 0xee, 0xf8, 0x87, 0x2b, 0x7e, 0x3f, 0x7d, 0xea, 0x9f, 0x0e, 0xf4, 0xd8,
	0xf5, 0x89, 0x5d, 0x09, 0x79, 0x4f, 0x5d, 0x3f, 0xe7, 0xd5, 0xd1, 0xe7, 0xdd, 0x74, 0xb3, 0xff,
	0x6b, 0xf8, 0xc1, 0x02, 0x5b, 0xed, 0x12, 0x5e, 0xc9, 0x09, 0x06, 0x16, 0xcb, 0xfd, 0x7d, 0xec,
	0x5e, 0x8b, 0x68, 0x18, 0xd6, 0xc7, 0x41, 0xe3, 0x9f, 0xe6, 0xf3, 0xec, 0xe2, 0x47, 0x0b, 0x1c,
	0xb3, 0x50, 0xd8, 0x9b, 0xce, 0xf0, 0x82, 0x0d, 0xe2, 0x35, 0x08, 0xef, 0x85, 0x37, 0x96, 0x61,
	0xad, 0x99, 0x3b, 0x87, 0x9f, 0x15, 0x9a, 0xd9, 0x97, 0x73, 0xb4, 0xe1, 0xfd, 0x19, 0x44, 0x8b,
	0x08, 0xad, 0xed, 0x8e, 0x7b, 0x72, 0x39, 0xe5, 0x17, 0x45, 0x69, 0x36, 0xe3, 0x9c, 0x72, 0x78,
	0x53, 0xdc, 0xd9, 0x9a, 0xf9, 0x5a, 0xa9, 0xe5, 0x5f, 0x2b, 0xb5, 0x0d, 0xfd, 0xb5, 0x92, 0x1f,
	0xf4, 0xad, 0x71, 0xe3, 0xee, 0x3a, 0x54, 0xf8, 0xce, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89,
	0x40, 0x83, 0x1a, 0x71, 0x09, 0x00, 0x00,
}