// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geometry/geometry.proto

/*
Package geometry is a generated protocol buffer package.

It is generated from these files:
	geometry/geometry.proto

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
	StartX int32 `protobuf:"varint,1,opt,name=start_x,json=startX" json:"start_x,omitempty"`
	StartY int32 `protobuf:"varint,2,opt,name=start_y,json=startY" json:"start_y,omitempty"`
	EndX   int32 `protobuf:"varint,3,opt,name=end_x,json=endX" json:"end_x,omitempty"`
	EndY   int32 `protobuf:"varint,4,opt,name=end_y,json=endY" json:"end_y,omitempty"`
}

func (m *BoundingBox2D) Reset()                    { *m = BoundingBox2D{} }
func (m *BoundingBox2D) String() string            { return proto.CompactTextString(m) }
func (*BoundingBox2D) ProtoMessage()               {}
func (*BoundingBox2D) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BoundingBox2D) GetStartX() int32 {
	if m != nil {
		return m.StartX
	}
	return 0
}

func (m *BoundingBox2D) GetStartY() int32 {
	if m != nil {
		return m.StartY
	}
	return 0
}

func (m *BoundingBox2D) GetEndX() int32 {
	if m != nil {
		return m.EndX
	}
	return 0
}

func (m *BoundingBox2D) GetEndY() int32 {
	if m != nil {
		return m.EndY
	}
	return 0
}

func init() {
	proto.RegisterType((*BoundingBox2D)(nil), "geometry.BoundingBox2D")
}

func init() { proto.RegisterFile("geometry/geometry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4f, 0xcd, 0xcf,
	0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0xa5, 0x6c, 0x2e, 0x5e, 0xa7, 0xfc, 0xd2, 0xbc, 0x94, 0xcc, 0xbc, 0x74, 0xa7, 0xfc, 0x0a,
	0x23, 0x17, 0x21, 0x71, 0x2e, 0xf6, 0xe2, 0x92, 0xc4, 0xa2, 0x92, 0xf8, 0x0a, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xd6, 0x20, 0x36, 0x30, 0x37, 0x02, 0x21, 0x51, 0x29, 0xc1, 0x84, 0x24, 0x11, 0x29,
	0x24, 0xcc, 0xc5, 0x9a, 0x9a, 0x97, 0x02, 0x54, 0xcf, 0x0c, 0x16, 0x66, 0x01, 0x72, 0x22, 0x60,
	0x82, 0x95, 0x12, 0x2c, 0x70, 0xc1, 0xc8, 0x24, 0x36, 0xb0, 0xed, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0x8b, 0x11, 0xf0, 0x98, 0x00, 0x00, 0x00,
}