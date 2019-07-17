// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package com_lzqysoft_bussinessenv_srv_filesys

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

type Response struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Mime                 string   `protobuf:"bytes,2,opt,name=mime,proto3" json:"mime,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Msg                  string   `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UploadFileRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	File                 []byte   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileRequest) Reset()         { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()    {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *UploadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest.Unmarshal(m, b)
}
func (m *UploadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest.Merge(m, src)
}
func (m *UploadFileRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest.Size(m)
}
func (m *UploadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest proto.InternalMessageInfo

func (m *UploadFileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadFileRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type DeleteFileRequest struct {
	Fid                  string   `protobuf:"bytes,1,opt,name=fid,proto3" json:"fid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFileRequest) Reset()         { *m = DeleteFileRequest{} }
func (m *DeleteFileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFileRequest) ProtoMessage()    {}
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{2}
}

func (m *DeleteFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFileRequest.Unmarshal(m, b)
}
func (m *DeleteFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFileRequest.Merge(m, src)
}
func (m *DeleteFileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFileRequest.Size(m)
}
func (m *DeleteFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFileRequest proto.InternalMessageInfo

func (m *DeleteFileRequest) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

type UpdateFileRequest struct {
	Fid                  string   `protobuf:"bytes,1,opt,name=fid,proto3" json:"fid,omitempty"`
	File                 []byte   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFileRequest) Reset()         { *m = UpdateFileRequest{} }
func (m *UpdateFileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFileRequest) ProtoMessage()    {}
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{3}
}

func (m *UpdateFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFileRequest.Unmarshal(m, b)
}
func (m *UpdateFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFileRequest.Merge(m, src)
}
func (m *UpdateFileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFileRequest.Size(m)
}
func (m *UpdateFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFileRequest proto.InternalMessageInfo

func (m *UpdateFileRequest) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *UpdateFileRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "com.lzqysoft.bussinessenv.srv.file.Response")
	proto.RegisterType((*UploadFileRequest)(nil), "com.lzqysoft.bussinessenv.srv.file.UploadFileRequest")
	proto.RegisterType((*DeleteFileRequest)(nil), "com.lzqysoft.bussinessenv.srv.file.DeleteFileRequest")
	proto.RegisterType((*UpdateFileRequest)(nil), "com.lzqysoft.bussinessenv.srv.file.UpdateFileRequest")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x13, 0x8b, 0x0e, 0x1e, 0xda, 0x3d, 0x48, 0xf0, 0x54, 0x16, 0x0a, 0x3d, 0xad,
	0xa0, 0x17, 0xc5, 0xab, 0xf8, 0x00, 0x0b, 0x3e, 0x40, 0x6a, 0x26, 0xb2, 0xb0, 0x9b, 0x4d, 0x3b,
	0x9b, 0x42, 0xfa, 0xe2, 0x5e, 0x65, 0xd6, 0x6a, 0xd4, 0x0a, 0xa6, 0xbd, 0xfd, 0x19, 0xe6, 0x9f,
	0x6f, 0xe6, 0xcf, 0x02, 0x54, 0xc6, 0xa2, 0x6a, 0xd6, 0x3e, 0x78, 0x31, 0x7f, 0xf1, 0x4e, 0xd9,
	0xed, 0xaa, 0x23, 0x5f, 0x05, 0xb5, 0x6c, 0x89, 0x4c, 0x8d, 0x44, 0x58, 0x6f, 0x14, 0xad, 0x37,
	0x8a, 0x3b, 0xa9, 0x23, 0xd9, 0xc0, 0x99, 0x46, 0x6a, 0x7c, 0x4d, 0x28, 0x2e, 0x61, 0x4c, 0xa1,
	0x08, 0x2d, 0xe5, 0xc9, 0x2c, 0x59, 0xa4, 0x7a, 0xf7, 0x25, 0x04, 0x64, 0xce, 0x38, 0xcc, 0x47,
	0xb3, 0x64, 0x71, 0xae, 0xa3, 0xe6, 0x5a, 0x5d, 0x38, 0xcc, 0xd3, 0x8f, 0x1a, 0x6b, 0xae, 0x91,
	0xd9, 0x62, 0x9e, 0x45, 0x77, 0xd4, 0x62, 0x02, 0xa9, 0xa3, 0xd7, 0xfc, 0x34, 0xb6, 0xb1, 0x94,
	0x0f, 0x30, 0x7d, 0x6e, 0xac, 0x2f, 0xca, 0x27, 0x63, 0x51, 0xe3, 0xaa, 0x45, 0x0a, 0x5f, 0xe3,
	0x92, 0x9f, 0xe3, 0x78, 0xcb, 0x88, 0xbd, 0xd0, 0x51, 0xcb, 0x39, 0x4c, 0x1f, 0xd1, 0x62, 0xc0,
	0xef, 0xe6, 0x09, 0xa4, 0x95, 0x29, 0x77, 0x5e, 0x96, 0xf2, 0x9e, 0x19, 0x65, 0xf1, 0x4f, 0xdb,
	0x5f, 0x84, 0x9b, 0xb7, 0x11, 0x64, 0xec, 0x12, 0x1d, 0x40, 0xbf, 0xa7, 0xb8, 0x53, 0x83, 0xf2,
	0x54, 0x7b, 0xa7, 0x5d, 0x5d, 0x0f, 0x74, 0x7e, 0xfe, 0x06, 0x79, 0xc2, 0xe8, 0xfe, 0xca, 0xc1,
	0xe8, 0xbd, 0x60, 0x8e, 0x44, 0xf7, 0xc9, 0x1d, 0x70, 0xf5, 0xaf, 0xb0, 0x8f, 0x40, 0x2f, 0xc7,
	0xf1, 0xe1, 0xde, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x54, 0xdd, 0x02, 0x45, 0xc6, 0x02, 0x00,
	0x00,
}