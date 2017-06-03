// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health/health.proto

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health/health.proto

It has these top-level messages:
	HealthCheckRequest
	HealthCheckResponse
*/
package health

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=health.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "health.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "health.HealthCheckResponse")
	proto.RegisterEnum("health.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/health.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health/health.proto",
}

func init() { proto.RegisterFile("health/health.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94,
	0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e,
	0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x95, 0x92, 0x1e, 0x97, 0x90, 0x07, 0x58, 0x9d,
	0x73, 0x46, 0x6a, 0x72, 0x76, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x04, 0x17, 0x7b,
	0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab,
	0x34, 0x85, 0x91, 0x4b, 0x18, 0x45, 0x43, 0x71, 0x01, 0xd0, 0xb0, 0x54, 0x21, 0x47, 0x2e, 0xb6,
	0x62, 0xa0, 0xc9, 0xa5, 0xc5, 0x60, 0x0d, 0x7c, 0x46, 0x9a, 0x7a, 0x50, 0xc7, 0x60, 0x51, 0xac,
	0x17, 0x0c, 0x32, 0x2c, 0x2f, 0x3d, 0x18, 0xac, 0x21, 0x08, 0xaa, 0x51, 0xc9, 0x8a, 0x8b, 0x17,
	0x45, 0x42, 0x88, 0x9b, 0x8b, 0x3d, 0xd4, 0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80, 0x01, 0xc4,
	0x09, 0x76, 0x0d, 0x0a, 0xf3, 0xf4, 0x73, 0x17, 0x60, 0x14, 0xe2, 0xe7, 0xe2, 0xf6, 0xf3, 0x0f,
	0x89, 0x87, 0x09, 0x30, 0x19, 0x65, 0x70, 0xb1, 0x41, 0x2c, 0x12, 0x8a, 0xe3, 0x62, 0x05, 0x5b,
	0x26, 0x24, 0x85, 0xd5, 0x05, 0x60, 0xff, 0x49, 0x49, 0xe3, 0x71, 0x9d, 0x92, 0x4c, 0xd3, 0xe5,
	0x27, 0x93, 0x99, 0xc4, 0x84, 0x44, 0xf4, 0xcb, 0x0c, 0xa1, 0x41, 0xaa, 0x5f, 0x0d, 0xf5, 0x7f,
	0x6d, 0x12, 0x1b, 0x38, 0xdc, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0x91, 0xbd, 0x45,
	0x74, 0x01, 0x00, 0x00,
}