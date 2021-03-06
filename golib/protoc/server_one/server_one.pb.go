// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go_micro/golib/protoc/server_one/server_one.proto

package pbserverone

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

type GetUserByUserNameReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByUserNameReq) Reset()         { *m = GetUserByUserNameReq{} }
func (m *GetUserByUserNameReq) String() string { return proto.CompactTextString(m) }
func (*GetUserByUserNameReq) ProtoMessage()    {}
func (*GetUserByUserNameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc6ef0ae1e83d450, []int{0}
}

func (m *GetUserByUserNameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByUserNameReq.Unmarshal(m, b)
}
func (m *GetUserByUserNameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByUserNameReq.Marshal(b, m, deterministic)
}
func (m *GetUserByUserNameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByUserNameReq.Merge(m, src)
}
func (m *GetUserByUserNameReq) XXX_Size() int {
	return xxx_messageInfo_GetUserByUserNameReq.Size(m)
}
func (m *GetUserByUserNameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByUserNameReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByUserNameReq proto.InternalMessageInfo

func (m *GetUserByUserNameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetUserByUserNameResp struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	CreateAt             string   `protobuf:"bytes,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt             string   `protobuf:"bytes,5,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByUserNameResp) Reset()         { *m = GetUserByUserNameResp{} }
func (m *GetUserByUserNameResp) String() string { return proto.CompactTextString(m) }
func (*GetUserByUserNameResp) ProtoMessage()    {}
func (*GetUserByUserNameResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc6ef0ae1e83d450, []int{1}
}

func (m *GetUserByUserNameResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByUserNameResp.Unmarshal(m, b)
}
func (m *GetUserByUserNameResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByUserNameResp.Marshal(b, m, deterministic)
}
func (m *GetUserByUserNameResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByUserNameResp.Merge(m, src)
}
func (m *GetUserByUserNameResp) XXX_Size() int {
	return xxx_messageInfo_GetUserByUserNameResp.Size(m)
}
func (m *GetUserByUserNameResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByUserNameResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByUserNameResp proto.InternalMessageInfo

func (m *GetUserByUserNameResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUserByUserNameResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserByUserNameResp) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *GetUserByUserNameResp) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *GetUserByUserNameResp) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserByUserNameReq)(nil), "pbserverone.GetUserByUserNameReq")
	proto.RegisterType((*GetUserByUserNameResp)(nil), "pbserverone.GetUserByUserNameResp")
}

func init() {
	proto.RegisterFile("go_micro/golib/protoc/server_one/server_one.proto", fileDescriptor_cc6ef0ae1e83d450)
}

var fileDescriptor_cc6ef0ae1e83d450 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0x8f, 0xcf,
	0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xcf, 0xcf, 0xc9, 0x4c, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9,
	0x4f, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0xcf, 0xcf, 0x4b, 0x45, 0x62, 0xea, 0x81,
	0x25, 0x85, 0xb8, 0x0b, 0x92, 0x20, 0x62, 0xf9, 0x79, 0xa9, 0x4a, 0x5a, 0x5c, 0x22, 0xee, 0xa9,
	0x25, 0xa1, 0xc5, 0xa9, 0x45, 0x4e, 0x95, 0x20, 0xd2, 0x2f, 0x31, 0x37, 0x35, 0x28, 0xb5, 0x50,
	0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc,
	0x56, 0x6a, 0x67, 0xe4, 0x12, 0xc5, 0xa2, 0xb8, 0xb8, 0x40, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x05,
	0xac, 0x96, 0x35, 0x88, 0x29, 0x33, 0x05, 0xae, 0x9b, 0x09, 0xa1, 0x5b, 0x48, 0x80, 0x8b, 0x39,
	0x31, 0x3d, 0x55, 0x82, 0x19, 0xac, 0x08, 0xc4, 0x14, 0x92, 0xe6, 0xe2, 0x4c, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x8d, 0x4f, 0x2c, 0x91, 0x60, 0x01, 0x2b, 0xe5, 0x80, 0x08, 0x38, 0x96, 0x80, 0x24,
	0x4b, 0x0b, 0x52, 0xa0, 0x92, 0xac, 0x10, 0x49, 0x88, 0x80, 0x63, 0x89, 0x51, 0x0a, 0x97, 0x40,
	0x30, 0xd8, 0x0b, 0xfe, 0x79, 0xa9, 0x20, 0x46, 0x66, 0x72, 0xaa, 0x50, 0x00, 0x17, 0x3b, 0xd4,
	0x71, 0x42, 0x8a, 0x7a, 0x48, 0x5e, 0xd4, 0xc3, 0xe6, 0x3f, 0x29, 0x25, 0x42, 0x4a, 0x8a, 0x0b,
	0x92, 0xd8, 0xc0, 0xe1, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x66, 0x25, 0xd5, 0x3e, 0x64,
	0x01, 0x00, 0x00,
}
