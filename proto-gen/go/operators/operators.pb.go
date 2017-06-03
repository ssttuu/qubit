// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operators/operators.proto

/*
Package operators is a generated protocol buffer package.

It is generated from these files:
	operators/operators.proto

It has these top-level messages:
	Operator
	ListOperatorsRequest
	ListOperatorsResponse
	GetOperatorRequest
	CreateOperatorRequest
	UpdateOperatorRequest
	DeleteOperatorRequest
	RenderOperatorRequest
	RenderOperatorResponse
*/
package operators

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import geometry "github.com/stupschwartz/qubit/proto-gen/go/geometry"

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

type Operator struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Context string `protobuf:"bytes,3,opt,name=context" json:"context,omitempty"`
	Type    string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
}

func (m *Operator) Reset()                    { *m = Operator{} }
func (m *Operator) String() string            { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()               {}
func (*Operator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Operator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Operator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operator) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *Operator) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ListOperatorsRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	PageSize       int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken      string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListOperatorsRequest) Reset()                    { *m = ListOperatorsRequest{} }
func (m *ListOperatorsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOperatorsRequest) ProtoMessage()               {}
func (*ListOperatorsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListOperatorsRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *ListOperatorsRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *ListOperatorsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListOperatorsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListOperatorsResponse struct {
	Operators     []*Operator `protobuf:"bytes,1,rep,name=operators" json:"operators,omitempty"`
	NextPageToken string      `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListOperatorsResponse) Reset()                    { *m = ListOperatorsResponse{} }
func (m *ListOperatorsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOperatorsResponse) ProtoMessage()               {}
func (*ListOperatorsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListOperatorsResponse) GetOperators() []*Operator {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *ListOperatorsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetOperatorRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
}

func (m *GetOperatorRequest) Reset()                    { *m = GetOperatorRequest{} }
func (m *GetOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOperatorRequest) ProtoMessage()               {}
func (*GetOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetOperatorRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *GetOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *GetOperatorRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

type CreateOperatorRequest struct {
	OrganizationId string    `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string    `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Operator       *Operator `protobuf:"bytes,3,opt,name=operator" json:"operator,omitempty"`
}

func (m *CreateOperatorRequest) Reset()                    { *m = CreateOperatorRequest{} }
func (m *CreateOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOperatorRequest) ProtoMessage()               {}
func (*CreateOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateOperatorRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *CreateOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *CreateOperatorRequest) GetOperator() *Operator {
	if m != nil {
		return m.Operator
	}
	return nil
}

type UpdateOperatorRequest struct {
	OrganizationId string    `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string    `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string    `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	Operator       *Operator `protobuf:"bytes,4,opt,name=operator" json:"operator,omitempty"`
}

func (m *UpdateOperatorRequest) Reset()                    { *m = UpdateOperatorRequest{} }
func (m *UpdateOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateOperatorRequest) ProtoMessage()               {}
func (*UpdateOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateOperatorRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *UpdateOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *UpdateOperatorRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *UpdateOperatorRequest) GetOperator() *Operator {
	if m != nil {
		return m.Operator
	}
	return nil
}

type DeleteOperatorRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
}

func (m *DeleteOperatorRequest) Reset()                    { *m = DeleteOperatorRequest{} }
func (m *DeleteOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOperatorRequest) ProtoMessage()               {}
func (*DeleteOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteOperatorRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *DeleteOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *DeleteOperatorRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

type RenderOperatorRequest struct {
	OrganizationId string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	SceneId        string `protobuf:"bytes,2,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId     string `protobuf:"bytes,3,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	// TODO: use double?
	Frame       int32                   `protobuf:"varint,4,opt,name=frame" json:"frame,omitempty"`
	BoundingBox *geometry.BoundingBox2D `protobuf:"bytes,5,opt,name=bounding_box,json=boundingBox" json:"bounding_box,omitempty"`
}

func (m *RenderOperatorRequest) Reset()                    { *m = RenderOperatorRequest{} }
func (m *RenderOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*RenderOperatorRequest) ProtoMessage()               {}
func (*RenderOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RenderOperatorRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *RenderOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *RenderOperatorRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *RenderOperatorRequest) GetFrame() int32 {
	if m != nil {
		return m.Frame
	}
	return 0
}

func (m *RenderOperatorRequest) GetBoundingBox() *geometry.BoundingBox2D {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

type RenderOperatorResponse struct {
	ResultUrl  string `protobuf:"bytes,1,opt,name=result_url,json=resultUrl" json:"result_url,omitempty"`
	ResultType string `protobuf:"bytes,2,opt,name=result_type,json=resultType" json:"result_type,omitempty"`
}

func (m *RenderOperatorResponse) Reset()                    { *m = RenderOperatorResponse{} }
func (m *RenderOperatorResponse) String() string            { return proto.CompactTextString(m) }
func (*RenderOperatorResponse) ProtoMessage()               {}
func (*RenderOperatorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RenderOperatorResponse) GetResultUrl() string {
	if m != nil {
		return m.ResultUrl
	}
	return ""
}

func (m *RenderOperatorResponse) GetResultType() string {
	if m != nil {
		return m.ResultType
	}
	return ""
}

func init() {
	proto.RegisterType((*Operator)(nil), "operators.Operator")
	proto.RegisterType((*ListOperatorsRequest)(nil), "operators.ListOperatorsRequest")
	proto.RegisterType((*ListOperatorsResponse)(nil), "operators.ListOperatorsResponse")
	proto.RegisterType((*GetOperatorRequest)(nil), "operators.GetOperatorRequest")
	proto.RegisterType((*CreateOperatorRequest)(nil), "operators.CreateOperatorRequest")
	proto.RegisterType((*UpdateOperatorRequest)(nil), "operators.UpdateOperatorRequest")
	proto.RegisterType((*DeleteOperatorRequest)(nil), "operators.DeleteOperatorRequest")
	proto.RegisterType((*RenderOperatorRequest)(nil), "operators.RenderOperatorRequest")
	proto.RegisterType((*RenderOperatorResponse)(nil), "operators.RenderOperatorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Operators service

type OperatorsClient interface {
	List(ctx context.Context, in *ListOperatorsRequest, opts ...grpc.CallOption) (*ListOperatorsResponse, error)
	Get(ctx context.Context, in *GetOperatorRequest, opts ...grpc.CallOption) (*Operator, error)
	Create(ctx context.Context, in *CreateOperatorRequest, opts ...grpc.CallOption) (*Operator, error)
	Update(ctx context.Context, in *UpdateOperatorRequest, opts ...grpc.CallOption) (*Operator, error)
	Delete(ctx context.Context, in *DeleteOperatorRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Render(ctx context.Context, in *RenderOperatorRequest, opts ...grpc.CallOption) (*RenderOperatorResponse, error)
}

type operatorsClient struct {
	cc *grpc.ClientConn
}

func NewOperatorsClient(cc *grpc.ClientConn) OperatorsClient {
	return &operatorsClient{cc}
}

func (c *operatorsClient) List(ctx context.Context, in *ListOperatorsRequest, opts ...grpc.CallOption) (*ListOperatorsResponse, error) {
	out := new(ListOperatorsResponse)
	err := grpc.Invoke(ctx, "/operators.Operators/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Get(ctx context.Context, in *GetOperatorRequest, opts ...grpc.CallOption) (*Operator, error) {
	out := new(Operator)
	err := grpc.Invoke(ctx, "/operators.Operators/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Create(ctx context.Context, in *CreateOperatorRequest, opts ...grpc.CallOption) (*Operator, error) {
	out := new(Operator)
	err := grpc.Invoke(ctx, "/operators.Operators/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Update(ctx context.Context, in *UpdateOperatorRequest, opts ...grpc.CallOption) (*Operator, error) {
	out := new(Operator)
	err := grpc.Invoke(ctx, "/operators.Operators/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Delete(ctx context.Context, in *DeleteOperatorRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/operators.Operators/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Render(ctx context.Context, in *RenderOperatorRequest, opts ...grpc.CallOption) (*RenderOperatorResponse, error) {
	out := new(RenderOperatorResponse)
	err := grpc.Invoke(ctx, "/operators.Operators/Render", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Operators service

type OperatorsServer interface {
	List(context.Context, *ListOperatorsRequest) (*ListOperatorsResponse, error)
	Get(context.Context, *GetOperatorRequest) (*Operator, error)
	Create(context.Context, *CreateOperatorRequest) (*Operator, error)
	Update(context.Context, *UpdateOperatorRequest) (*Operator, error)
	Delete(context.Context, *DeleteOperatorRequest) (*google_protobuf1.Empty, error)
	Render(context.Context, *RenderOperatorRequest) (*RenderOperatorResponse, error)
}

func RegisterOperatorsServer(s *grpc.Server, srv OperatorsServer) {
	s.RegisterService(&_Operators_serviceDesc, srv)
}

func _Operators_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).List(ctx, req.(*ListOperatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Get(ctx, req.(*GetOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Create(ctx, req.(*CreateOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Update(ctx, req.(*UpdateOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Delete(ctx, req.(*DeleteOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_Render_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Render(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Render",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Render(ctx, req.(*RenderOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operators_serviceDesc = grpc.ServiceDesc{
	ServiceName: "operators.Operators",
	HandlerType: (*OperatorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Operators_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Operators_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Operators_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Operators_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Operators_Delete_Handler,
		},
		{
			MethodName: "Render",
			Handler:    _Operators_Render_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operators/operators.proto",
}

func init() { proto.RegisterFile("operators/operators.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x95, 0xcf, 0x4f, 0x13, 0x41,
	0x14, 0xc7, 0x53, 0x28, 0xd0, 0x3e, 0x14, 0x92, 0x91, 0x1f, 0x65, 0x91, 0x80, 0x7b, 0x50, 0xe2,
	0xa1, 0x1b, 0xf0, 0xc6, 0xc5, 0x88, 0x18, 0x42, 0xd4, 0xa0, 0x15, 0x8c, 0x26, 0x24, 0xcd, 0x96,
	0x7d, 0x6c, 0x36, 0x96, 0x99, 0x75, 0x76, 0x6a, 0x28, 0xca, 0xc5, 0xb3, 0x07, 0xa3, 0x89, 0x89,
	0x37, 0xef, 0xc6, 0xa3, 0xff, 0x86, 0x17, 0xff, 0x05, 0xff, 0x10, 0xe7, 0xc7, 0xfe, 0x6a, 0x5d,
	0xd4, 0x84, 0x12, 0x6e, 0x33, 0xef, 0xbd, 0xe6, 0x7d, 0xe6, 0xfb, 0xfa, 0xbe, 0x0b, 0x73, 0x2c,
	0x44, 0xee, 0x0a, 0xc6, 0x23, 0x27, 0x3d, 0xd5, 0x43, 0xce, 0x04, 0x23, 0xd5, 0x34, 0x60, 0x5d,
	0xf5, 0x19, 0xf3, 0xdb, 0xe8, 0xb8, 0x61, 0xe0, 0xb8, 0x94, 0x32, 0xe1, 0x8a, 0x80, 0xd1, 0xb8,
	0xd0, 0x9a, 0x8f, 0xb3, 0xfa, 0xd6, 0xea, 0x1c, 0x38, 0x78, 0x18, 0x8a, 0x6e, 0x9c, 0x9c, 0xf5,
	0x91, 0x1d, 0xa2, 0xe0, 0x5d, 0x27, 0x39, 0x98, 0x84, 0xbd, 0x07, 0x95, 0xed, 0xb8, 0x01, 0x99,
	0x80, 0xa1, 0xc0, 0xab, 0x95, 0x96, 0x4a, 0xcb, 0xd5, 0x86, 0x3c, 0x11, 0x02, 0x65, 0xea, 0x1e,
	0x62, 0x6d, 0x48, 0x47, 0xf4, 0x99, 0xd4, 0x60, 0x6c, 0x9f, 0x51, 0x81, 0x47, 0xa2, 0x36, 0xac,
	0xc3, 0xc9, 0x55, 0x55, 0x8b, 0x6e, 0x88, 0xb5, 0xb2, 0xa9, 0x56, 0x67, 0xfb, 0x53, 0x09, 0xa6,
	0x1e, 0x04, 0x91, 0x48, 0x5a, 0x44, 0x0d, 0x7c, 0xd9, 0xc1, 0x48, 0x90, 0x1b, 0x30, 0xc9, 0xb8,
	0xef, 0xd2, 0xe0, 0x58, 0xbf, 0xa1, 0x99, 0xf6, 0x9d, 0xc8, 0x87, 0xb7, 0x3c, 0x32, 0x07, 0x95,
	0x68, 0x1f, 0x29, 0xaa, 0x0a, 0xc3, 0x31, 0xa6, 0xef, 0x32, 0x35, 0x0f, 0xd5, 0xd0, 0xf5, 0xb1,
	0x19, 0x05, 0xc7, 0xa8, 0x61, 0x46, 0x1a, 0x15, 0x15, 0x78, 0x22, 0xef, 0x64, 0x01, 0x40, 0x27,
	0x05, 0x7b, 0x81, 0x34, 0x66, 0xd2, 0xe5, 0x3b, 0x2a, 0x60, 0x73, 0x98, 0xee, 0xe3, 0x8a, 0x42,
	0x29, 0x25, 0x92, 0x15, 0xc8, 0x04, 0x97, 0x48, 0xc3, 0xcb, 0xe3, 0xab, 0x57, 0xea, 0xd9, 0x4c,
	0x92, 0x1f, 0x34, 0xb2, 0x2a, 0x72, 0x1d, 0x26, 0xa9, 0x14, 0xa0, 0x99, 0xeb, 0x67, 0x48, 0x2f,
	0xab, 0xf0, 0xa3, 0xb4, 0x67, 0x17, 0xc8, 0x26, 0xa6, 0x2d, 0x07, 0xa9, 0xc4, 0x22, 0x8c, 0x27,
	0x38, 0x2a, 0x6b, 0x06, 0x03, 0x49, 0x68, 0xcb, 0xb3, 0xdf, 0x95, 0x60, 0xfa, 0x2e, 0x47, 0x57,
	0xe0, 0x79, 0xb4, 0x77, 0xa0, 0x92, 0xf4, 0xd2, 0xbd, 0x4f, 0x91, 0x2c, 0x2d, 0xb2, 0xbf, 0x49,
	0x9c, 0xdd, 0xd0, 0x3b, 0x27, 0x9c, 0x7f, 0xa9, 0xd1, 0xc3, 0x5b, 0xfe, 0x1f, 0xde, 0x37, 0x30,
	0xbd, 0x81, 0x6d, 0xbc, 0x18, 0x5c, 0xfb, 0x87, 0x54, 0xab, 0x81, 0xd4, 0x43, 0x7e, 0x21, 0x6a,
	0x4d, 0xc1, 0xc8, 0x01, 0x57, 0x36, 0x50, 0xd6, 0x2b, 0x66, 0x2e, 0x64, 0x0d, 0x2e, 0xb5, 0x58,
	0x87, 0x7a, 0x01, 0xf5, 0x9b, 0x2d, 0x76, 0x54, 0x1b, 0xd1, 0x3a, 0xce, 0xd6, 0x53, 0x7b, 0x59,
	0x8f, 0xb3, 0xeb, 0xec, 0x68, 0x75, 0xa3, 0x31, 0xde, 0xca, 0xae, 0xf6, 0x33, 0x98, 0xe9, 0x7f,
	0x4f, 0xbc, 0x7d, 0x72, 0x6b, 0x39, 0x46, 0x9d, 0xb6, 0x68, 0x76, 0x78, 0x3b, 0x7e, 0x4b, 0xd5,
	0x44, 0x76, 0x79, 0x5b, 0xb1, 0xc6, 0x69, 0xed, 0x34, 0xe6, 0x25, 0xf1, 0x2f, 0x76, 0x64, 0x64,
	0xf5, 0xfb, 0x18, 0x54, 0xd3, 0x9d, 0x26, 0x1f, 0x4a, 0x50, 0x56, 0x5b, 0x4e, 0x16, 0x73, 0xe3,
	0x2d, 0xb2, 0x23, 0x6b, 0xe9, 0xf4, 0x02, 0x43, 0x66, 0x6f, 0xbe, 0xfd, 0xf9, 0xeb, 0xe3, 0xd0,
	0x1d, 0x72, 0xdb, 0x79, 0xb5, 0xe2, 0xe4, 0xd5, 0x8d, 0x9c, 0xd7, 0x7d, 0x33, 0x38, 0x71, 0xb4,
	0xb6, 0x32, 0x91, 0x68, 0x7e, 0x92, 0xb9, 0x3a, 0x79, 0x5f, 0x82, 0x61, 0x69, 0x03, 0x64, 0x21,
	0xd7, 0xf2, 0x4f, 0x5b, 0xb0, 0x8a, 0xfe, 0x91, 0xf6, 0xae, 0x86, 0xd8, 0x26, 0x0f, 0xcf, 0x08,
	0x21, 0xab, 0xb3, 0x91, 0x9f, 0x28, 0x9d, 0x46, 0x8d, 0x3b, 0x90, 0xbc, 0x10, 0x85, 0x86, 0x51,
	0x0c, 0xf6, 0x58, 0x83, 0xdd, 0xb7, 0xcf, 0xaa, 0xce, 0x5a, 0xba, 0x73, 0xe4, 0x8b, 0x84, 0x32,
	0x1e, 0xd1, 0x03, 0x55, 0x68, 0x1b, 0xc5, 0x50, 0x4d, 0x0d, 0xf5, 0xdc, 0x1a, 0xac, 0x5a, 0x39,
	0xc4, 0xcf, 0x12, 0xd1, 0xf8, 0x42, 0x0f, 0x62, 0xa1, 0x55, 0x58, 0x33, 0x75, 0xf3, 0x7d, 0xae,
	0x27, 0xdf, 0xe7, 0xfa, 0x3d, 0xf5, 0x7d, 0x4e, 0x66, 0x7a, 0x73, 0xc0, 0x33, 0xfd, 0x2a, 0xd9,
	0xcc, 0x92, 0xf5, 0xb0, 0x15, 0xfa, 0x88, 0x75, 0xed, 0x2f, 0x15, 0xf1, 0xff, 0x7f, 0x4f, 0x63,
	0x3e, 0xb5, 0x77, 0x06, 0x2b, 0x26, 0xd7, 0xdd, 0x5a, 0xa3, 0x5a, 0x94, 0x5b, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x0b, 0x19, 0xf7, 0x6a, 0x08, 0x09, 0x00, 0x00,
}