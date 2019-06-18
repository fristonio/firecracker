// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysql.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type None struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *None) Reset()         { *m = None{} }
func (m *None) String() string { return proto.CompactTextString(m) }
func (*None) ProtoMessage()    {}
func (*None) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysql_40a914aca32fed04, []int{0}
}
func (m *None) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_None.Unmarshal(m, b)
}
func (m *None) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_None.Marshal(b, m, deterministic)
}
func (dst *None) XXX_Merge(src proto.Message) {
	xxx_messageInfo_None.Merge(dst, src)
}
func (m *None) XXX_Size() int {
	return xxx_messageInfo_None.Size(m)
}
func (m *None) XXX_DiscardUnknown() {
	xxx_messageInfo_None.DiscardUnknown(m)
}

var xxx_messageInfo_None proto.InternalMessageInfo

// Message respresenting a MySQL instance
type MySQLInstance struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Database             string   `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLInstance) Reset()         { *m = MySQLInstance{} }
func (m *MySQLInstance) String() string { return proto.CompactTextString(m) }
func (*MySQLInstance) ProtoMessage()    {}
func (*MySQLInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_mysql_40a914aca32fed04, []int{1}
}
func (m *MySQLInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLInstance.Unmarshal(m, b)
}
func (m *MySQLInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLInstance.Marshal(b, m, deterministic)
}
func (dst *MySQLInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLInstance.Merge(dst, src)
}
func (m *MySQLInstance) XXX_Size() int {
	return xxx_messageInfo_MySQLInstance.Size(m)
}
func (m *MySQLInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLInstance.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLInstance proto.InternalMessageInfo

func (m *MySQLInstance) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MySQLInstance) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *MySQLInstance) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*None)(nil), "protobuf.None")
	proto.RegisterType((*MySQLInstance)(nil), "protobuf.MySQLInstance")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MySQLServiceClient is the client API for MySQLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MySQLServiceClient interface {
	// RPC to create a new MySQL instance, an instance is represented by a
	// user and a database. So this RPC creates a new user along with a database
	// which the user owns. After creating the database it returns the instance
	// created.
	CreateMySQLInstance(ctx context.Context, in *None, opts ...grpc.CallOption) (*MySQLInstance, error)
	// This RPC deletes an existing MySQL instance, it deletes the user and the database
	// user is associated with. As a return value it results RPCResult depicting if the
	// operation was successful or not.
	DeleteMySQLInstance(ctx context.Context, in *MySQLInstance, opts ...grpc.CallOption) (*None, error)
}

type mySQLServiceClient struct {
	cc *grpc.ClientConn
}

func NewMySQLServiceClient(cc *grpc.ClientConn) MySQLServiceClient {
	return &mySQLServiceClient{cc}
}

func (c *mySQLServiceClient) CreateMySQLInstance(ctx context.Context, in *None, opts ...grpc.CallOption) (*MySQLInstance, error) {
	out := new(MySQLInstance)
	err := c.cc.Invoke(ctx, "/protobuf.MySQLService/CreateMySQLInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySQLServiceClient) DeleteMySQLInstance(ctx context.Context, in *MySQLInstance, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/protobuf.MySQLService/DeleteMySQLInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MySQLServiceServer is the server API for MySQLService service.
type MySQLServiceServer interface {
	// RPC to create a new MySQL instance, an instance is represented by a
	// user and a database. So this RPC creates a new user along with a database
	// which the user owns. After creating the database it returns the instance
	// created.
	CreateMySQLInstance(context.Context, *None) (*MySQLInstance, error)
	// This RPC deletes an existing MySQL instance, it deletes the user and the database
	// user is associated with. As a return value it results RPCResult depicting if the
	// operation was successful or not.
	DeleteMySQLInstance(context.Context, *MySQLInstance) (*None, error)
}

func RegisterMySQLServiceServer(s *grpc.Server, srv MySQLServiceServer) {
	s.RegisterService(&_MySQLService_serviceDesc, srv)
}

func _MySQLService_CreateMySQLInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServiceServer).CreateMySQLInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MySQLService/CreateMySQLInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServiceServer).CreateMySQLInstance(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySQLService_DeleteMySQLInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MySQLInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServiceServer).DeleteMySQLInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MySQLService/DeleteMySQLInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServiceServer).DeleteMySQLInstance(ctx, req.(*MySQLInstance))
	}
	return interceptor(ctx, in, info, handler)
}

var _MySQLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.MySQLService",
	HandlerType: (*MySQLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMySQLInstance",
			Handler:    _MySQLService_CreateMySQLInstance_Handler,
		},
		{
			MethodName: "DeleteMySQLInstance",
			Handler:    _MySQLService_DeleteMySQLInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mysql.proto",
}

func init() { proto.RegisterFile("mysql.proto", fileDescriptor_mysql_40a914aca32fed04) }

var fileDescriptor_mysql_40a914aca32fed04 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xad, 0x2c, 0x2e,
	0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0x6c,
	0x5c, 0x2c, 0x7e, 0xf9, 0x79, 0xa9, 0x4a, 0xc9, 0x5c, 0xbc, 0xbe, 0x95, 0xc1, 0x81, 0x3e, 0x9e,
	0x79, 0xc5, 0x25, 0x89, 0x79, 0xc9, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0xa5, 0xc5, 0xa9, 0x45, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0x3e, 0x48, 0x2e, 0x25, 0xb1,
	0x24, 0x31, 0x29, 0xb1, 0x38, 0x55, 0x82, 0x09, 0x22, 0x07, 0xe3, 0x83, 0xe4, 0x0a, 0x12, 0x8b,
	0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x98, 0x21, 0x72, 0x30, 0xbe, 0xd1, 0x24, 0x46, 0x2e, 0x1e,
	0xb0, 0x2d, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x0e, 0x5c, 0xc2, 0xce, 0x45, 0xa9,
	0x89, 0x25, 0xa9, 0xa8, 0x76, 0xf3, 0xe9, 0xc1, 0xdc, 0xa7, 0x07, 0x72, 0x9c, 0x94, 0x38, 0x82,
	0x8f, 0xa2, 0x50, 0x89, 0x01, 0x64, 0x82, 0x4b, 0x6a, 0x4e, 0x2a, 0xba, 0x09, 0xb8, 0x74, 0x48,
	0xa1, 0x19, 0xad, 0xc4, 0x90, 0xc4, 0x06, 0x16, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0x87, 0xbc, 0x06, 0x21, 0x01, 0x00, 0x00,
}