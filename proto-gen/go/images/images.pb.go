// Code generated by protoc-gen-go. DO NOT EDIT.
// source: images/images.proto

/*
Package images is a generated protocol buffer package.

It is generated from these files:
	images/images.proto

It has these top-level messages:
	Image
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

type Image struct {
	Id              string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ImageSequenceId string            `protobuf:"bytes,2,opt,name=image_sequence_id,json=imageSequenceId" json:"image_sequence_id,omitempty"`
	Name            string            `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Width           int32             `protobuf:"varint,4,opt,name=width" json:"width,omitempty"`
	Height          int32             `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Labels          map[string]string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Planes          []*Plane          `protobuf:"bytes,7,rep,name=planes" json:"planes,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetImageSequenceId() string {
	if m != nil {
		return m.ImageSequenceId
	}
	return ""
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Image) GetPlanes() []*Plane {
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
	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListImagesRequest) Reset()                    { *m = ListImagesRequest{} }
func (m *ListImagesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListImagesRequest) ProtoMessage()               {}
func (*ListImagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	Images        []*Image `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListImagesResponse) Reset()                    { *m = ListImagesResponse{} }
func (m *ListImagesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListImagesResponse) ProtoMessage()               {}
func (*ListImagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListImagesResponse) GetImages() []*Image {
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
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetImageRequest) Reset()                    { *m = GetImageRequest{} }
func (m *GetImageRequest) String() string            { return proto.CompactTextString(m) }
func (*GetImageRequest) ProtoMessage()               {}
func (*GetImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetImageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateImageRequest struct {
	Image *Image `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
}

func (m *CreateImageRequest) Reset()                    { *m = CreateImageRequest{} }
func (m *CreateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateImageRequest) ProtoMessage()               {}
func (*CreateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateImageRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type UpdateImageRequest struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Image *Image `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
}

func (m *UpdateImageRequest) Reset()                    { *m = UpdateImageRequest{} }
func (m *UpdateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateImageRequest) ProtoMessage()               {}
func (*UpdateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateImageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateImageRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type DeleteImageRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteImageRequest) Reset()                    { *m = DeleteImageRequest{} }
func (m *DeleteImageRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteImageRequest) ProtoMessage()               {}
func (*DeleteImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteImageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Image)(nil), "images.Image")
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
	Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error)
	Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*Image, error)
	Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*Image, error)
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

func (c *imagesClient) Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/images.Images/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/images.Images/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
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
	Get(context.Context, *GetImageRequest) (*Image, error)
	Create(context.Context, *CreateImageRequest) (*Image, error)
	Update(context.Context, *UpdateImageRequest) (*Image, error)
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
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x52, 0xd4, 0x4e,
	0x10, 0xae, 0x24, 0x9b, 0x00, 0xcd, 0x8f, 0xdf, 0x42, 0x63, 0x41, 0x08, 0xa2, 0x18, 0xff, 0x14,
	0x85, 0x55, 0x49, 0x81, 0x17, 0xe1, 0xe0, 0x41, 0x44, 0x8b, 0x2a, 0x4a, 0x31, 0xc8, 0x99, 0x0a,
	0x64, 0xdc, 0x4d, 0x11, 0x92, 0xb8, 0x19, 0x40, 0xb4, 0xbc, 0xf8, 0x0a, 0xde, 0xbc, 0xf8, 0x50,
	0xbe, 0x02, 0x0f, 0x62, 0xa6, 0x67, 0x12, 0x96, 0x64, 0xe5, 0xe2, 0x69, 0xd3, 0xfd, 0x75, 0x7f,
	0xf3, 0x7d, 0x3d, 0xbd, 0x03, 0xb3, 0xf1, 0x69, 0xd8, 0x63, 0x85, 0x2f, 0x7f, 0xbc, 0x7c, 0x90,
	0xf1, 0x0c, 0x2d, 0x19, 0x39, 0x77, 0x7b, 0x59, 0xd6, 0x4b, 0x98, 0x1f, 0xe6, 0xb1, 0x1f, 0xa6,
	0x69, 0xc6, 0x43, 0x1e, 0x67, 0xa9, 0xaa, 0x72, 0x16, 0x15, 0x4a, 0xd1, 0xd1, 0xd9, 0x47, 0x9f,
	0x9d, 0xe6, 0xfc, 0x52, 0x82, 0xee, 0x4f, 0x1d, 0xcc, 0x1d, 0xc1, 0x82, 0xff, 0x83, 0x1e, 0x47,
	0xb6, 0xb6, 0xac, 0xad, 0x4c, 0x04, 0xe5, 0x17, 0xae, 0xc2, 0x0c, 0xd1, 0x1f, 0x16, 0xec, 0xd3,
	0x19, 0x4b, 0x8f, 0xd9, 0x61, 0x09, 0xeb, 0x04, 0x77, 0x09, 0xd8, 0x57, 0xf9, 0x9d, 0x08, 0x11,
	0x3a, 0x69, 0x78, 0xca, 0x6c, 0x83, 0x60, 0xfa, 0xc6, 0x3b, 0x60, 0x5e, 0xc4, 0x11, 0xef, 0xdb,
	0x9d, 0x32, 0x69, 0x06, 0x32, 0xc0, 0x39, 0xb0, 0xfa, 0x2c, 0xee, 0xf5, 0xb9, 0x6d, 0x52, 0x5a,
	0x45, 0xb8, 0x06, 0x56, 0x12, 0x1e, 0xb1, 0xa4, 0xb0, 0xad, 0x65, 0x63, 0x65, 0x72, 0x7d, 0xc1,
	0x53, 0x4e, 0x49, 0x9c, 0xb7, 0x4b, 0xd8, 0x76, 0xca, 0x07, 0x97, 0x81, 0x2a, 0xc4, 0xc7, 0x60,
	0xe5, 0x49, 0x98, 0xb2, 0xc2, 0x1e, 0xa3, 0x96, 0xa9, 0xaa, 0x65, 0x4f, 0x64, 0x03, 0x05, 0x3a,
	0x1b, 0x30, 0x39, 0xd4, 0x8d, 0xd3, 0x60, 0x9c, 0xb0, 0x4b, 0xe5, 0x53, 0x7c, 0x0a, 0xa1, 0xe7,
	0x61, 0x72, 0xc6, 0x94, 0x39, 0x19, 0x6c, 0xea, 0xcf, 0x35, 0xf7, 0x4a, 0x03, 0x93, 0xc8, 0x6a,
	0x83, 0xda, 0x28, 0x83, 0xfa, 0x68, 0x83, 0xc6, 0x5f, 0x0c, 0x76, 0x6e, 0x1a, 0xa4, 0x03, 0x46,
	0x1a, 0x7c, 0x0a, 0xe3, 0xc7, 0xfd, 0xf2, 0x3e, 0x45, 0x93, 0x49, 0x4d, 0xdd, 0xaa, 0x69, 0x4b,
	0xe6, 0x83, 0xba, 0xe0, 0x5f, 0x6c, 0xbe, 0x80, 0x31, 0xc5, 0x37, 0xd2, 0xe7, 0x7d, 0xe8, 0x0c,
	0xb2, 0x8b, 0xa2, 0xec, 0x13, 0x12, 0x26, 0x2b, 0x09, 0x41, 0x76, 0x11, 0x10, 0xe0, 0x2e, 0x81,
	0x51, 0x06, 0xa5, 0xf3, 0x4e, 0x14, 0xf2, 0xb0, 0xec, 0x35, 0x56, 0xb4, 0x97, 0xfa, 0xb4, 0x16,
	0x50, 0xec, 0xbe, 0x83, 0x99, 0xdd, 0xb8, 0xe0, 0x74, 0x91, 0x45, 0x20, 0x96, 0xa6, 0xe0, 0xb8,
	0x08, 0x13, 0x39, 0x2d, 0x57, 0xfc, 0x45, 0x9e, 0x66, 0x06, 0xe3, 0x22, 0xb1, 0x5f, 0xc6, 0xb8,
	0x04, 0x40, 0x20, 0xcf, 0x4e, 0x58, 0xaa, 0xf4, 0x52, 0xf9, 0x07, 0x91, 0x70, 0x8f, 0x01, 0x87,
	0x09, 0x8b, 0xbc, 0xdc, 0x75, 0x26, 0xd6, 0x41, 0x2a, 0x23, 0x01, 0x43, 0xeb, 0x40, 0x75, 0x81,
	0x02, 0xf1, 0x09, 0x74, 0x53, 0xf6, 0x99, 0x1f, 0xb6, 0x0e, 0x98, 0x12, 0xe9, 0xbd, 0xfa, 0x90,
	0x07, 0xd0, 0x7d, 0xc3, 0xe4, 0x19, 0x95, 0xe6, 0xc6, 0x3f, 0xc4, 0xdd, 0x00, 0xdc, 0x1a, 0xb0,
	0x90, 0xb3, 0x1b, 0x55, 0x0f, 0xc1, 0xa4, 0xa3, 0xa8, 0xb0, 0x25, 0x43, 0x62, 0xee, 0x0e, 0xe0,
	0x41, 0x1e, 0x35, 0x5b, 0x9b, 0x7f, 0xc1, 0x9a, 0x4a, 0xbf, 0x85, 0xea, 0x11, 0xe0, 0x2b, 0x96,
	0xb0, 0xdb, 0xa9, 0xd6, 0x7f, 0x19, 0x60, 0xc9, 0x81, 0xe1, 0x7b, 0xe8, 0x88, 0xf1, 0x61, 0xbd,
	0x81, 0xad, 0xdb, 0x71, 0x9c, 0x51, 0x90, 0x9c, 0xb3, 0x8b, 0xdf, 0x7f, 0x5f, 0xfd, 0xd0, 0xff,
	0x43, 0xf0, 0xcf, 0xd7, 0xd4, 0x73, 0x84, 0xaf, 0xc1, 0x28, 0x87, 0x85, 0xf3, 0x55, 0x5b, 0x63,
	0x72, 0xce, 0x4d, 0xe5, 0xee, 0x3c, 0x51, 0xcc, 0x60, 0xf7, 0x9a, 0xc2, 0xff, 0x1a, 0x47, 0xdf,
	0xf0, 0x2d, 0x58, 0x72, 0xa2, 0x58, 0x2b, 0x68, 0x4f, 0xb8, 0xc9, 0xb6, 0x40, 0x6c, 0xb3, 0xee,
	0x90, 0xa0, 0x4d, 0x39, 0x1b, 0xdc, 0x07, 0x4b, 0x8e, 0xf9, 0x9a, 0xaf, 0x3d, 0xf6, 0x26, 0xdf,
	0x3d, 0xe2, 0xb3, 0x9d, 0xa6, 0xba, 0x8a, 0xf4, 0x00, 0x2c, 0x39, 0xf0, 0x6b, 0xd2, 0xf6, 0x05,
	0x38, 0x73, 0x9e, 0x7c, 0x76, 0xbd, 0xea, 0xd9, 0xf5, 0xb6, 0xc5, 0xb3, 0x5b, 0x79, 0x5f, 0x6d,
	0xb2, 0x1f, 0x59, 0x54, 0xf8, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x9c, 0x47, 0x40,
	0xea, 0x05, 0x00, 0x00,
}
