// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operators/operators.proto

/*
Package operators is a generated protocol buffer package.

It is generated from these files:
	operators/operators.proto

It has these top-level messages:
	Operator
	Connection
	ListOperatorsRequest
	ListOperatorsResponse
	GetOperatorRequest
	CreateOperatorRequest
	DeleteOperatorRequest
	RenameOperatorRequest
	ConnectOperatorRequest
	DisconnectOperatorRequest
	Value
	KeyFrameValue
	ExpressionValue
	SetValueRequest
	SetKeyFrameRequest
	SetExpressionRequest
*/
package operators

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

type Operator struct {
	SceneId      string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Id           string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Version      int32  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	OperatorData []byte `protobuf:"bytes,5,opt,name=operator_data,json=operatorData,proto3" json:"operator_data,omitempty"`
}

func (m *Operator) Reset()                    { *m = Operator{} }
func (m *Operator) String() string            { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()               {}
func (*Operator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Operator) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *Operator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Operator) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Operator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operator) GetOperatorData() []byte {
	if m != nil {
		return m.OperatorData
	}
	return nil
}

type Connection struct {
	SceneId     string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	InputId     string `protobuf:"bytes,3,opt,name=input_id,json=inputId" json:"input_id,omitempty"`
	InputIndex  int32  `protobuf:"varint,4,opt,name=input_index,json=inputIndex" json:"input_index,omitempty"`
	OutputId    string `protobuf:"bytes,5,opt,name=output_id,json=outputId" json:"output_id,omitempty"`
	OutputIndex int32  `protobuf:"varint,6,opt,name=output_index,json=outputIndex" json:"output_index,omitempty"`
}

func (m *Connection) Reset()                    { *m = Connection{} }
func (m *Connection) String() string            { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()               {}
func (*Connection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Connection) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *Connection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Connection) GetInputId() string {
	if m != nil {
		return m.InputId
	}
	return ""
}

func (m *Connection) GetInputIndex() int32 {
	if m != nil {
		return m.InputIndex
	}
	return 0
}

func (m *Connection) GetOutputId() string {
	if m != nil {
		return m.OutputId
	}
	return ""
}

func (m *Connection) GetOutputIndex() int32 {
	if m != nil {
		return m.OutputIndex
	}
	return 0
}

// Operator Requests
type ListOperatorsRequest struct {
	SceneId   string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListOperatorsRequest) Reset()                    { *m = ListOperatorsRequest{} }
func (m *ListOperatorsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOperatorsRequest) ProtoMessage()               {}
func (*ListOperatorsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*ListOperatorsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	SceneId string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *GetOperatorRequest) Reset()                    { *m = GetOperatorRequest{} }
func (m *GetOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOperatorRequest) ProtoMessage()               {}
func (*GetOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *GetOperatorRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateOperatorRequest struct {
	SceneId  string    `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Operator *Operator `protobuf:"bytes,2,opt,name=operator" json:"operator,omitempty"`
}

func (m *CreateOperatorRequest) Reset()                    { *m = CreateOperatorRequest{} }
func (m *CreateOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOperatorRequest) ProtoMessage()               {}
func (*CreateOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

type DeleteOperatorRequest struct {
	SceneId string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteOperatorRequest) Reset()                    { *m = DeleteOperatorRequest{} }
func (m *DeleteOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOperatorRequest) ProtoMessage()               {}
func (*DeleteOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *DeleteOperatorRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RenameOperatorRequest struct {
	SceneId string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *RenameOperatorRequest) Reset()                    { *m = RenameOperatorRequest{} }
func (m *RenameOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*RenameOperatorRequest) ProtoMessage()               {}
func (*RenameOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RenameOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *RenameOperatorRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RenameOperatorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Connection Requests
type ConnectOperatorRequest struct {
	SceneId    string      `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Connection *Connection `protobuf:"bytes,2,opt,name=connection" json:"connection,omitempty"`
}

func (m *ConnectOperatorRequest) Reset()                    { *m = ConnectOperatorRequest{} }
func (m *ConnectOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*ConnectOperatorRequest) ProtoMessage()               {}
func (*ConnectOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ConnectOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *ConnectOperatorRequest) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

type DisconnectOperatorRequest struct {
	SceneId      string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId" json:"connection_id,omitempty"`
}

func (m *DisconnectOperatorRequest) Reset()                    { *m = DisconnectOperatorRequest{} }
func (m *DisconnectOperatorRequest) String() string            { return proto.CompactTextString(m) }
func (*DisconnectOperatorRequest) ProtoMessage()               {}
func (*DisconnectOperatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DisconnectOperatorRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *DisconnectOperatorRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

// Parameters Requests
type Value struct {
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type KeyFrameValue struct {
	Time  float64 `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	Value *Value  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyFrameValue) Reset()                    { *m = KeyFrameValue{} }
func (m *KeyFrameValue) String() string            { return proto.CompactTextString(m) }
func (*KeyFrameValue) ProtoMessage()               {}
func (*KeyFrameValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *KeyFrameValue) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *KeyFrameValue) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type ExpressionValue struct {
	Expression string `protobuf:"bytes,1,opt,name=expression" json:"expression,omitempty"`
}

func (m *ExpressionValue) Reset()                    { *m = ExpressionValue{} }
func (m *ExpressionValue) String() string            { return proto.CompactTextString(m) }
func (*ExpressionValue) ProtoMessage()               {}
func (*ExpressionValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ExpressionValue) GetExpression() string {
	if m != nil {
		return m.Expression
	}
	return ""
}

type SetValueRequest struct {
	SceneId     string `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId  string `protobuf:"bytes,2,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	ParameterId string `protobuf:"bytes,3,opt,name=parameter_id,json=parameterId" json:"parameter_id,omitempty"`
	Value       *Value `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *SetValueRequest) Reset()                    { *m = SetValueRequest{} }
func (m *SetValueRequest) String() string            { return proto.CompactTextString(m) }
func (*SetValueRequest) ProtoMessage()               {}
func (*SetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SetValueRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *SetValueRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *SetValueRequest) GetParameterId() string {
	if m != nil {
		return m.ParameterId
	}
	return ""
}

func (m *SetValueRequest) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetKeyFrameRequest struct {
	SceneId     string         `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId  string         `protobuf:"bytes,2,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	ParameterId string         `protobuf:"bytes,3,opt,name=parameter_id,json=parameterId" json:"parameter_id,omitempty"`
	Keyframe    *KeyFrameValue `protobuf:"bytes,4,opt,name=keyframe" json:"keyframe,omitempty"`
}

func (m *SetKeyFrameRequest) Reset()                    { *m = SetKeyFrameRequest{} }
func (m *SetKeyFrameRequest) String() string            { return proto.CompactTextString(m) }
func (*SetKeyFrameRequest) ProtoMessage()               {}
func (*SetKeyFrameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SetKeyFrameRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *SetKeyFrameRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *SetKeyFrameRequest) GetParameterId() string {
	if m != nil {
		return m.ParameterId
	}
	return ""
}

func (m *SetKeyFrameRequest) GetKeyframe() *KeyFrameValue {
	if m != nil {
		return m.Keyframe
	}
	return nil
}

type SetExpressionRequest struct {
	SceneId     string           `protobuf:"bytes,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	OperatorId  string           `protobuf:"bytes,2,opt,name=operator_id,json=operatorId" json:"operator_id,omitempty"`
	ParameterId string           `protobuf:"bytes,3,opt,name=parameter_id,json=parameterId" json:"parameter_id,omitempty"`
	Expression  *ExpressionValue `protobuf:"bytes,4,opt,name=expression" json:"expression,omitempty"`
}

func (m *SetExpressionRequest) Reset()                    { *m = SetExpressionRequest{} }
func (m *SetExpressionRequest) String() string            { return proto.CompactTextString(m) }
func (*SetExpressionRequest) ProtoMessage()               {}
func (*SetExpressionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *SetExpressionRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

func (m *SetExpressionRequest) GetOperatorId() string {
	if m != nil {
		return m.OperatorId
	}
	return ""
}

func (m *SetExpressionRequest) GetParameterId() string {
	if m != nil {
		return m.ParameterId
	}
	return ""
}

func (m *SetExpressionRequest) GetExpression() *ExpressionValue {
	if m != nil {
		return m.Expression
	}
	return nil
}

func init() {
	proto.RegisterType((*Operator)(nil), "operators.Operator")
	proto.RegisterType((*Connection)(nil), "operators.Connection")
	proto.RegisterType((*ListOperatorsRequest)(nil), "operators.ListOperatorsRequest")
	proto.RegisterType((*ListOperatorsResponse)(nil), "operators.ListOperatorsResponse")
	proto.RegisterType((*GetOperatorRequest)(nil), "operators.GetOperatorRequest")
	proto.RegisterType((*CreateOperatorRequest)(nil), "operators.CreateOperatorRequest")
	proto.RegisterType((*DeleteOperatorRequest)(nil), "operators.DeleteOperatorRequest")
	proto.RegisterType((*RenameOperatorRequest)(nil), "operators.RenameOperatorRequest")
	proto.RegisterType((*ConnectOperatorRequest)(nil), "operators.ConnectOperatorRequest")
	proto.RegisterType((*DisconnectOperatorRequest)(nil), "operators.DisconnectOperatorRequest")
	proto.RegisterType((*Value)(nil), "operators.Value")
	proto.RegisterType((*KeyFrameValue)(nil), "operators.KeyFrameValue")
	proto.RegisterType((*ExpressionValue)(nil), "operators.ExpressionValue")
	proto.RegisterType((*SetValueRequest)(nil), "operators.SetValueRequest")
	proto.RegisterType((*SetKeyFrameRequest)(nil), "operators.SetKeyFrameRequest")
	proto.RegisterType((*SetExpressionRequest)(nil), "operators.SetExpressionRequest")
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
	Delete(ctx context.Context, in *DeleteOperatorRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Rename(ctx context.Context, in *RenameOperatorRequest, opts ...grpc.CallOption) (*Operator, error)
	// Connections API
	Connect(ctx context.Context, in *ConnectOperatorRequest, opts ...grpc.CallOption) (*Connection, error)
	Disconnect(ctx context.Context, in *DisconnectOperatorRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Parameters API
	// TODO: does anything need to be returned or are status codes sufficient?
	// TODO: this will be a chatty API, ideally status codes are sufficient.
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	SetKeyFrame(ctx context.Context, in *SetKeyFrameRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	SetExpression(ctx context.Context, in *SetExpressionRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
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

func (c *operatorsClient) Delete(ctx context.Context, in *DeleteOperatorRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/operators.Operators/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Rename(ctx context.Context, in *RenameOperatorRequest, opts ...grpc.CallOption) (*Operator, error) {
	out := new(Operator)
	err := grpc.Invoke(ctx, "/operators.Operators/Rename", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Connect(ctx context.Context, in *ConnectOperatorRequest, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := grpc.Invoke(ctx, "/operators.Operators/Connect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) Disconnect(ctx context.Context, in *DisconnectOperatorRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/operators.Operators/Disconnect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/operators.Operators/SetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) SetKeyFrame(ctx context.Context, in *SetKeyFrameRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/operators.Operators/SetKeyFrame", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorsClient) SetExpression(ctx context.Context, in *SetExpressionRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/operators.Operators/SetExpression", in, out, c.cc, opts...)
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
	Delete(context.Context, *DeleteOperatorRequest) (*google_protobuf1.Empty, error)
	Rename(context.Context, *RenameOperatorRequest) (*Operator, error)
	// Connections API
	Connect(context.Context, *ConnectOperatorRequest) (*Connection, error)
	Disconnect(context.Context, *DisconnectOperatorRequest) (*google_protobuf1.Empty, error)
	// Parameters API
	// TODO: does anything need to be returned or are status codes sufficient?
	// TODO: this will be a chatty API, ideally status codes are sufficient.
	SetValue(context.Context, *SetValueRequest) (*google_protobuf1.Empty, error)
	SetKeyFrame(context.Context, *SetKeyFrameRequest) (*google_protobuf1.Empty, error)
	SetExpression(context.Context, *SetExpressionRequest) (*google_protobuf1.Empty, error)
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

func _Operators_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Rename(ctx, req.(*RenameOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Connect(ctx, req.(*ConnectOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).Disconnect(ctx, req.(*DisconnectOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_SetKeyFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKeyFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).SetKeyFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/SetKeyFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).SetKeyFrame(ctx, req.(*SetKeyFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operators_SetExpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorsServer).SetExpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operators.Operators/SetExpression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorsServer).SetExpression(ctx, req.(*SetExpressionRequest))
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
			MethodName: "Delete",
			Handler:    _Operators_Delete_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _Operators_Rename_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Operators_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Operators_Disconnect_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _Operators_SetValue_Handler,
		},
		{
			MethodName: "SetKeyFrame",
			Handler:    _Operators_SetKeyFrame_Handler,
		},
		{
			MethodName: "SetExpression",
			Handler:    _Operators_SetExpression_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operators/operators.proto",
}

func init() { proto.RegisterFile("operators/operators.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 941 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x96, 0x5d, 0x6b, 0x33, 0x45,
	0x14, 0xc7, 0xd9, 0xe6, 0xfd, 0x24, 0xb1, 0x32, 0x3e, 0x79, 0x48, 0xf3, 0xf8, 0xd8, 0x66, 0x2c,
	0xb5, 0xb4, 0x92, 0x25, 0xad, 0x45, 0x0c, 0x88, 0x60, 0x5b, 0x4b, 0xa9, 0x60, 0xd9, 0x6a, 0x41,
	0x7a, 0x11, 0xb6, 0xc9, 0x34, 0xac, 0x6d, 0x77, 0xd7, 0xdd, 0x49, 0xe9, 0x0b, 0xf1, 0x42, 0x10,
	0xbc, 0x17, 0xc4, 0x3b, 0xaf, 0xbc, 0xf0, 0x52, 0xc4, 0x4f, 0xe2, 0x57, 0xf0, 0x83, 0x38, 0x3b,
	0xb3, 0x3b, 0x3b, 0x9b, 0x6e, 0x9a, 0xe6, 0xa1, 0xf4, 0x6e, 0xe7, 0x9c, 0x33, 0xe7, 0xfc, 0xcf,
	0xbc, 0xfc, 0x76, 0x60, 0xc1, 0x71, 0x89, 0x67, 0x52, 0xc7, 0xf3, 0x75, 0xf9, 0xd5, 0x72, 0x3d,
	0x87, 0x3a, 0xa8, 0x24, 0x0d, 0x8d, 0x77, 0x07, 0x8e, 0x33, 0xb8, 0x20, 0xba, 0xe9, 0x5a, 0xba,
	0x69, 0xdb, 0x0e, 0x35, 0xa9, 0xe5, 0xd8, 0x61, 0x60, 0xe3, 0x55, 0xe8, 0xe5, 0xa3, 0xd3, 0xe1,
	0x99, 0x4e, 0x2e, 0x5d, 0x7a, 0x23, 0x9c, 0xf8, 0x67, 0x0d, 0x8a, 0x5f, 0x85, 0x89, 0xd0, 0x02,
	0x14, 0xfd, 0x1e, 0xb1, 0x49, 0xd7, 0xea, 0xd7, 0xb5, 0x25, 0x6d, 0xb5, 0x64, 0x14, 0xf8, 0x78,
	0xbf, 0x8f, 0xde, 0x82, 0x39, 0x66, 0x9c, 0xe3, 0x46, 0xf6, 0x85, 0xea, 0x50, 0xb8, 0x22, 0x9e,
	0xcf, 0xca, 0xd4, 0x33, 0xcc, 0x98, 0x33, 0xa2, 0x21, 0x42, 0x90, 0xb5, 0xcd, 0x4b, 0x52, 0xcf,
	0xf2, 0x58, 0xfe, 0x8d, 0xde, 0x87, 0x6a, 0xa4, 0xb6, 0xdb, 0x37, 0xa9, 0x59, 0xcf, 0x31, 0x67,
	0xc5, 0xa8, 0x44, 0xc6, 0x1d, 0x66, 0xc3, 0x7f, 0x6b, 0x00, 0xdb, 0x8e, 0x6d, 0x93, 0x5e, 0xa0,
	0x7e, 0x16, 0x31, 0x2c, 0xd4, 0xb2, 0xdd, 0x21, 0x0d, 0x42, 0x33, 0x22, 0x94, 0x8f, 0x59, 0xe8,
	0x22, 0x94, 0x43, 0x97, 0xdd, 0x27, 0xd7, 0x5c, 0x54, 0xce, 0x00, 0xe1, 0x0d, 0x2c, 0xe8, 0x15,
	0x94, 0x9c, 0x21, 0x0d, 0x27, 0xe7, 0xf8, 0xe4, 0xa2, 0x30, 0xb0, 0xd9, 0x4d, 0xa8, 0x44, 0x4e,
	0x3e, 0x3d, 0xcf, 0xa7, 0x97, 0x43, 0x7f, 0x60, 0xc2, 0x97, 0xf0, 0xe2, 0x4b, 0xcb, 0xa7, 0xd1,
	0x1a, 0xfa, 0x06, 0xf9, 0x7e, 0x48, 0x7c, 0xfa, 0x90, 0x7c, 0x56, 0xd2, 0x35, 0x07, 0xa4, 0xeb,
	0x5b, 0xb7, 0x84, 0x77, 0x91, 0x33, 0x8a, 0x81, 0xe1, 0x88, 0x8d, 0xd1, 0x6b, 0x00, 0xee, 0xa4,
	0xce, 0x39, 0xb1, 0xc3, 0x6e, 0x78, 0xf8, 0xd7, 0x81, 0x01, 0x7b, 0x50, 0x1b, 0x2b, 0xe7, 0xbb,
	0x6c, 0xab, 0x09, 0x6a, 0x43, 0x7c, 0x20, 0x58, 0xc1, 0xcc, 0x6a, 0x79, 0xe3, 0x9d, 0x56, 0x7c,
	0x66, 0xa2, 0x09, 0x46, 0x1c, 0x85, 0x56, 0x60, 0xde, 0x26, 0xd7, 0xb4, 0xab, 0xd4, 0x13, 0x6b,
	0x5a, 0x0d, 0xcc, 0x87, 0xb2, 0xe6, 0x67, 0x80, 0xf6, 0x88, 0x2c, 0xf9, 0x88, 0x06, 0xc7, 0xf6,
	0x07, 0xf7, 0xa0, 0xb6, 0xed, 0x11, 0x93, 0x92, 0x19, 0x72, 0xe8, 0x50, 0x8c, 0x94, 0xf2, 0x4c,
	0x13, 0xda, 0x91, 0x41, 0xf8, 0x73, 0xa8, 0xed, 0x90, 0x0b, 0x32, 0x53, 0x91, 0x71, 0xa1, 0xc7,
	0x50, 0x33, 0x48, 0x70, 0x62, 0xdf, 0x3c, 0x87, 0x3c, 0xff, 0x99, 0xf8, 0xfc, 0xe3, 0xef, 0xe0,
	0x65, 0x78, 0xb2, 0x67, 0x48, 0xbc, 0x05, 0xd0, 0x93, 0xd7, 0x21, 0x5c, 0x83, 0x9a, 0xb2, 0x06,
	0xf1, 0x5d, 0x31, 0x94, 0x40, 0x7c, 0x02, 0x0b, 0x3b, 0x96, 0xdf, 0x9b, 0xb9, 0x1c, 0xbb, 0xa3,
	0x71, 0x96, 0xae, 0x6c, 0xa9, 0x12, 0x1b, 0xf7, 0xfb, 0xb8, 0x00, 0xb9, 0x63, 0xf3, 0x62, 0x48,
	0xf0, 0x01, 0x54, 0x0f, 0xc8, 0xcd, 0x17, 0x1e, 0xeb, 0x8e, 0x1b, 0x82, 0xb6, 0xa9, 0xc5, 0xda,
	0x0e, 0xb2, 0x6a, 0x06, 0xff, 0x66, 0x07, 0x2c, 0x77, 0x15, 0x38, 0x43, 0xf1, 0x6f, 0x2b, 0xe2,
	0xf9, 0x24, 0x43, 0xb8, 0x71, 0x1b, 0xe6, 0x77, 0xaf, 0x5d, 0x8f, 0xf8, 0x01, 0x40, 0x44, 0xba,
	0xf7, 0x00, 0x88, 0x34, 0x85, 0x52, 0x15, 0x0b, 0xfe, 0x4d, 0x83, 0xf9, 0x23, 0x42, 0x45, 0x9a,
	0xe9, 0xcd, 0x31, 0x0c, 0x48, 0x00, 0xc9, 0xd6, 0x20, 0x32, 0x89, 0x9b, 0xee, 0x9a, 0x41, 0x37,
	0x94, 0x78, 0x31, 0x46, 0xca, 0xd2, 0xc6, 0x42, 0x64, 0x37, 0xd9, 0x87, 0xbb, 0xf9, 0x53, 0x03,
	0xc4, 0xa4, 0x45, 0xcb, 0xf3, 0x4c, 0xea, 0x3e, 0x82, 0xe2, 0x39, 0xb9, 0x39, 0xf3, 0x22, 0xf4,
	0x96, 0x37, 0xea, 0x8a, 0xc0, 0xc4, 0x5e, 0x19, 0x32, 0x12, 0xff, 0xa5, 0xc1, 0x0b, 0xa6, 0x35,
	0x5e, 0xfd, 0x67, 0x52, 0xdb, 0x49, 0x6c, 0xaf, 0xd0, 0xdb, 0x50, 0xf4, 0x8e, 0x1d, 0x07, 0x75,
	0xeb, 0x37, 0xfe, 0x01, 0x28, 0x49, 0xfe, 0x21, 0x0f, 0xb2, 0x01, 0x10, 0xd1, 0xa2, 0x32, 0x3b,
	0x0d, 0xc8, 0x8d, 0xa5, 0xc9, 0x01, 0x02, 0xa1, 0xf8, 0x83, 0x1f, 0xff, 0xfd, 0xef, 0x97, 0xb9,
	0x26, 0x5a, 0xd4, 0xaf, 0xda, 0x3a, 0xef, 0xd6, 0xd7, 0xef, 0xa2, 0x55, 0x18, 0xc5, 0x3f, 0x60,
	0x34, 0x80, 0x0c, 0x03, 0x22, 0x7a, 0xad, 0x64, 0xbc, 0x0f, 0xc8, 0x46, 0x1a, 0xaf, 0xf0, 0x87,
	0xbc, 0xc6, 0x0a, 0x5a, 0x9e, 0x52, 0x43, 0xbf, 0x63, 0x43, 0xd6, 0x5c, 0x5e, 0x80, 0x13, 0xa9,
	0xea, 0x53, 0x59, 0x9a, 0x5e, 0xae, 0xcd, 0xcb, 0xad, 0xe3, 0x69, 0x2d, 0x75, 0x24, 0x47, 0x91,
	0x0b, 0x79, 0xc1, 0xd1, 0x44, 0xcd, 0x54, 0xb4, 0x36, 0x5e, 0xb6, 0xc4, 0xdb, 0xa2, 0x15, 0xbd,
	0x2d, 0x5a, 0xbb, 0xc1, 0xdb, 0x22, 0xea, 0x72, 0xed, 0x71, 0x5d, 0xfe, 0x00, 0x79, 0x41, 0xdd,
	0x44, 0xc5, 0x54, 0x10, 0xa7, 0x77, 0xf9, 0x29, 0x2f, 0xf7, 0x31, 0x5e, 0x7f, 0x4c, 0x39, 0xdd,
	0xe3, 0x89, 0x95, 0x8e, 0x6f, 0xa1, 0x10, 0xb2, 0x14, 0x35, 0xef, 0xf3, 0x75, 0x5c, 0x41, 0x3a,
	0x82, 0xf1, 0x16, 0xd7, 0xa0, 0xe3, 0x66, 0xba, 0x86, 0x98, 0xa2, 0x7e, 0x47, 0xa1, 0x35, 0xfa,
	0x89, 0x3d, 0x7a, 0x62, 0x5c, 0xa3, 0x65, 0x75, 0xc9, 0x27, 0x51, 0x7c, 0xe2, 0xb2, 0x7f, 0xc2,
	0x35, 0x6c, 0xae, 0xb5, 0xa7, 0x6a, 0xd0, 0xef, 0x12, 0xac, 0x1f, 0xa1, 0x5f, 0xd9, 0x3b, 0x30,
	0xe2, 0x29, 0x52, 0x6f, 0xe2, 0x18, 0x64, 0x27, 0xd6, 0x3e, 0xe6, 0xb5, 0x0f, 0xf1, 0xde, 0xd4,
	0x3d, 0x50, 0xe0, 0x31, 0xd2, 0x25, 0x14, 0x98, 0x43, 0x85, 0xc6, 0xa8, 0x23, 0x68, 0x8a, 0x7e,
	0xd7, 0xa0, 0xac, 0xd0, 0x34, 0x71, 0xe9, 0xee, 0x53, 0x76, 0xa2, 0xbc, 0x6f, 0xb9, 0xbc, 0xa3,
	0xa7, 0x93, 0x27, 0x19, 0x8a, 0xfe, 0xd0, 0xa0, 0x9a, 0x60, 0x68, 0x82, 0x45, 0x69, 0x74, 0x9d,
	0xa8, 0xf2, 0x84, 0xab, 0xfc, 0xe6, 0xe9, 0x54, 0x2a, 0xdc, 0x3c, 0xcd, 0xf3, 0x62, 0x9b, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x47, 0x4d, 0xeb, 0x54, 0x0c, 0x00, 0x00,
}
