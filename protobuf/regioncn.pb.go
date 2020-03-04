// Code generated by protoc-gen-go. DO NOT EDIT.
// source: regioncn.proto

package protobuf

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

type Model struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ShortName            string   `protobuf:"bytes,4,opt,name=shortName,proto3" json:"shortName,omitempty"`
	Lat                  float64  `protobuf:"fixed64,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  float64  `protobuf:"fixed64,6,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccbb38e2d0c8ca27, []int{0}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Model) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *Model) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Model) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type Models struct {
	List                 []*Model `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Models) Reset()         { *m = Models{} }
func (m *Models) String() string { return proto.CompactTextString(m) }
func (*Models) ProtoMessage()    {}
func (*Models) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccbb38e2d0c8ca27, []int{1}
}

func (m *Models) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Models.Unmarshal(m, b)
}
func (m *Models) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Models.Marshal(b, m, deterministic)
}
func (m *Models) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Models.Merge(m, src)
}
func (m *Models) XXX_Size() int {
	return xxx_messageInfo_Models.Size(m)
}
func (m *Models) XXX_DiscardUnknown() {
	xxx_messageInfo_Models.DiscardUnknown(m)
}

var xxx_messageInfo_Models proto.InternalMessageInfo

func (m *Models) GetList() []*Model {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Model)(nil), "regioncn.Model")
	proto.RegisterType((*Models)(nil), "regioncn.Models")
}

func init() {
	proto.RegisterFile("regioncn.proto", fileDescriptor_ccbb38e2d0c8ca27)
}

var fileDescriptor_ccbb38e2d0c8ca27 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xbd, 0x0e, 0x82, 0x30,
	0x14, 0x46, 0x53, 0x0a, 0x04, 0xae, 0x09, 0x9a, 0x4e, 0x1d, 0x1c, 0x1a, 0x5c, 0xba, 0xc8, 0xa0,
	0x6f, 0xe0, 0xae, 0x43, 0x47, 0x37, 0xa0, 0x15, 0x9b, 0x60, 0x6b, 0xa0, 0x3e, 0x80, 0x6f, 0x6e,
	0x7a, 0xfd, 0xdb, 0xce, 0x3d, 0x67, 0xf9, 0x2e, 0x54, 0x93, 0x19, 0xac, 0x77, 0xbd, 0x6b, 0xee,
	0x93, 0x0f, 0x9e, 0x15, 0xdf, 0xbb, 0x7e, 0x12, 0xc8, 0x8e, 0x5e, 0x9b, 0x91, 0x55, 0x90, 0x58,
	0xcd, 0x89, 0x20, 0x92, 0xaa, 0xc4, 0x6a, 0xc6, 0x20, 0xed, 0xbd, 0x36, 0x3c, 0x11, 0x44, 0x96,
	0x0a, 0x39, 0x3a, 0xd7, 0xde, 0x0c, 0xa7, 0x6f, 0x17, 0x99, 0xad, 0xa1, 0x9c, 0xaf, 0x7e, 0x0a,
	0xa7, 0x18, 0x52, 0x0c, 0x7f, 0xc1, 0x56, 0x40, 0xc7, 0x36, 0xf0, 0x4c, 0x10, 0x49, 0x54, 0x44,
	0x34, 0x6e, 0xe0, 0xf9, 0xc7, 0xb8, 0xa1, 0xde, 0x42, 0x8e, 0x13, 0x66, 0xb6, 0x81, 0x74, 0xb4,
	0x73, 0xe0, 0x44, 0x50, 0xb9, 0xd8, 0x2d, 0x9b, 0xdf, 0x6c, 0xec, 0x0a, 0xe3, 0x01, 0xce, 0x05,
	0x7e, 0xd1, 0x3d, 0x2e, 0x5d, 0x8e, 0xb4, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xdf, 0xa7,
	0x88, 0xe1, 0x00, 0x00, 0x00,
}
