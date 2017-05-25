// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geometry.proto

/*
Package geometry is a generated protocol buffer package.

It is generated from these files:
	geometry.proto

It has these top-level messages:
	BoundingBox2D
*/
package geometry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BoundingBox2D struct {
	StartX int64 `protobuf:"varint,1,opt,name=start_x,json=startX" json:"start_x,omitempty"`
	StartY int64 `protobuf:"varint,2,opt,name=start_y,json=startY" json:"start_y,omitempty"`
	EndX   int64 `protobuf:"varint,3,opt,name=end_x,json=endX" json:"end_x,omitempty"`
	EndY   int64 `protobuf:"varint,4,opt,name=end_y,json=endY" json:"end_y,omitempty"`
}

func (m *BoundingBox2D) Reset()                    { *m = BoundingBox2D{} }
func (m *BoundingBox2D) String() string            { return proto.CompactTextString(m) }
func (*BoundingBox2D) ProtoMessage()               {}
func (*BoundingBox2D) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BoundingBox2D) GetStartX() int64 {
	if m != nil {
		return m.StartX
	}
	return 0
}

func (m *BoundingBox2D) GetStartY() int64 {
	if m != nil {
		return m.StartY
	}
	return 0
}

func (m *BoundingBox2D) GetEndX() int64 {
	if m != nil {
		return m.EndX
	}
	return 0
}

func (m *BoundingBox2D) GetEndY() int64 {
	if m != nil {
		return m.EndY
	}
	return 0
}

func init() {
	proto.RegisterType((*BoundingBox2D)(nil), "geometry.BoundingBox2D")
}

func init() { proto.RegisterFile("geometry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0xcd, 0xcf,
	0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xb2,
	0xb9, 0x78, 0x9d, 0xf2, 0x4b, 0xf3, 0x52, 0x32, 0xf3, 0xd2, 0x9d, 0xf2, 0x2b, 0x8c, 0x5c, 0x84,
	0xc4, 0xb9, 0xd8, 0x8b, 0x4b, 0x12, 0x8b, 0x4a, 0xe2, 0x2b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98,
	0x83, 0xd8, 0xc0, 0xdc, 0x08, 0x84, 0x44, 0xa5, 0x04, 0x13, 0x92, 0x44, 0xa4, 0x90, 0x30, 0x17,
	0x6b, 0x6a, 0x5e, 0x0a, 0x50, 0x3d, 0x33, 0x58, 0x98, 0x05, 0xc8, 0x89, 0x80, 0x09, 0x56, 0x4a,
	0xb0, 0xc0, 0x05, 0x23, 0x93, 0xd8, 0xc0, 0xb6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0f,
	0x90, 0xa8, 0xc0, 0x8f, 0x00, 0x00, 0x00,
}
