// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recommendations.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	recommendations.proto

It has these top-level messages:
	RecommendArticles
	Recommendations
*/
package model

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

type RecommendArticles struct {
	Details []*RecommendArticles_Details `protobuf:"bytes,1,rep,name=details" json:"details,omitempty"`
}

func (m *RecommendArticles) Reset()                    { *m = RecommendArticles{} }
func (m *RecommendArticles) String() string            { return proto.CompactTextString(m) }
func (*RecommendArticles) ProtoMessage()               {}
func (*RecommendArticles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RecommendArticles) GetDetails() []*RecommendArticles_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

type RecommendArticles_Details struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *RecommendArticles_Details) Reset()                    { *m = RecommendArticles_Details{} }
func (m *RecommendArticles_Details) String() string            { return proto.CompactTextString(m) }
func (*RecommendArticles_Details) ProtoMessage()               {}
func (*RecommendArticles_Details) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *RecommendArticles_Details) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Recommendations struct {
	Articles *RecommendArticles `protobuf:"bytes,1,opt,name=articles" json:"articles,omitempty"`
}

func (m *Recommendations) Reset()                    { *m = Recommendations{} }
func (m *Recommendations) String() string            { return proto.CompactTextString(m) }
func (*Recommendations) ProtoMessage()               {}
func (*Recommendations) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Recommendations) GetArticles() *RecommendArticles {
	if m != nil {
		return m.Articles
	}
	return nil
}

func init() {
	proto.RegisterType((*RecommendArticles)(nil), "model.RecommendArticles")
	proto.RegisterType((*RecommendArticles_Details)(nil), "model.RecommendArticles.Details")
	proto.RegisterType((*Recommendations)(nil), "model.Recommendations")
}

func init() { proto.RegisterFile("recommendations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0x4d, 0xce,
	0xcf, 0xcd, 0x4d, 0xcd, 0x4b, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x51, 0xca, 0xe2, 0x12, 0x0c, 0x82, 0xc9, 0x3b,
	0x16, 0x95, 0x64, 0x26, 0xe7, 0xa4, 0x16, 0x0b, 0x59, 0x71, 0xb1, 0xa7, 0xa4, 0x96, 0x24, 0x66,
	0xe6, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x29, 0xe8, 0x81, 0x55, 0xeb, 0x61, 0x28,
	0xd5, 0x73, 0x81, 0xa8, 0x0b, 0x82, 0x69, 0x90, 0x92, 0xe4, 0x62, 0x87, 0x8a, 0x09, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x28, 0xb9, 0x73,
	0xf1, 0x07, 0xa1, 0xba, 0x45, 0xc8, 0x84, 0x8b, 0x23, 0x11, 0x6a, 0x14, 0x58, 0x21, 0xb7, 0x91,
	0x04, 0x2e, 0xab, 0x82, 0xe0, 0x2a, 0x93, 0xd8, 0xc0, 0x5e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x32, 0x72, 0xdc, 0xeb, 0xdb, 0x00, 0x00, 0x00,
}
