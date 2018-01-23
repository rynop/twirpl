// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/blog/service.proto

/*
Package blog is a generated protocol buffer package.

It is generated from these files:
	rpc/blog/service.proto

It has these top-level messages:
	User
*/
package blog

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

// User subscribing to my blog
type User struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "com.rynop.coolapi.blog.User")
}

func init() { proto.RegisterFile("rpc/blog/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x2a, 0x48, 0xd6,
	0x4f, 0xca, 0xc9, 0x4f, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x4b, 0xce, 0xcf, 0xd5, 0x2b, 0xaa, 0xcc, 0xcb, 0x2f, 0xd0, 0x4b, 0xce,
	0xcf, 0xcf, 0x49, 0x2c, 0xc8, 0xd4, 0x03, 0xa9, 0x52, 0x72, 0xe0, 0x62, 0x09, 0x2d, 0x4e, 0x2d,
	0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c,
	0x11, 0x12, 0xe1, 0x62, 0x4d, 0xcd, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x02, 0x0b, 0x41, 0x38, 0x42,
	0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0xcc, 0x60, 0x41, 0x30, 0xdb, 0xc8, 0x9f, 0x8b,
	0xc5, 0x29, 0x27, 0x3f, 0x5d, 0xc8, 0x9d, 0x8b, 0x33, 0xb8, 0x34, 0xa9, 0x38, 0xb9, 0x28, 0x33,
	0x29, 0x55, 0x48, 0x46, 0x0f, 0xbb, 0x7d, 0x7a, 0x20, 0xcb, 0xa4, 0xf0, 0xca, 0x3a, 0xb1, 0x45,
	0xb1, 0x80, 0x38, 0x49, 0x6c, 0x60, 0x97, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x3f,
	0x20, 0x1a, 0xd3, 0x00, 0x00, 0x00,
}
