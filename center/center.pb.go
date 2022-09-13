// Code generated by protoc-gen-go. DO NOT EDIT.
// source: center/center.proto

package center

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// registe
type RegisterRequest struct {
	ServerName           string   `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Functions            []string `protobuf:"bytes,3,rep,name=functions,proto3" json:"functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6400d989d59d2c95, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *RegisterRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *RegisterRequest) GetFunctions() []string {
	if m != nil {
		return m.Functions
	}
	return nil
}

type RegisterResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6400d989d59d2c95, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// discovery
type DiscoveryRequest struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6400d989d59d2c95, []int{2}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type DiscoveryResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6400d989d59d2c95, []int{3}
}

func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "center.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "center.RegisterResponse")
	proto.RegisterType((*DiscoveryRequest)(nil), "center.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "center.DiscoveryResponse")
}

func init() { proto.RegisterFile("center/center.proto", fileDescriptor_6400d989d59d2c95) }

var fileDescriptor_6400d989d59d2c95 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4e, 0xcd, 0x2b,
	0x49, 0x2d, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x52,
	0x32, 0x17, 0x7f, 0x50, 0x6a, 0x7a, 0x66, 0x71, 0x49, 0x6a, 0x51, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x89, 0x90, 0x1c, 0x17, 0x57, 0x71, 0x6a, 0x51, 0x59, 0x6a, 0x91, 0x5f, 0x62, 0x6e, 0xaa,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x92, 0x88, 0x90, 0x10, 0x17, 0x4b, 0x62, 0x4a, 0x4a,
	0x91, 0x04, 0x13, 0x58, 0x06, 0xcc, 0x16, 0x92, 0xe1, 0xe2, 0x4c, 0x2b, 0xcd, 0x4b, 0x2e, 0xc9,
	0xcc, 0xcf, 0x2b, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x0c, 0x42, 0x08, 0x28, 0xa9, 0x70, 0x09,
	0x20, 0x2c, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe0, 0x62, 0x4e, 0x2d, 0x2a, 0x82,
	0x1a, 0x0f, 0x62, 0x2a, 0x69, 0x71, 0x09, 0xb8, 0x64, 0x16, 0x27, 0xe7, 0x97, 0xa5, 0x16, 0x55,
	0xc2, 0xdc, 0x22, 0xc6, 0xc5, 0x06, 0xb1, 0x19, 0xaa, 0x10, 0xca, 0x53, 0x52, 0xe5, 0x12, 0x44,
	0x52, 0x8b, 0xcb, 0x48, 0x27, 0xae, 0x28, 0x0e, 0x3d, 0xa8, 0xbf, 0x93, 0xd8, 0xc0, 0x1e, 0x37,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x46, 0x35, 0x77, 0xd1, 0x0f, 0x01, 0x00, 0x00,
}