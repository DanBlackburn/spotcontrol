// Code generated by protoc-gen-go.
// source: pubsub.proto
// DO NOT EDIT!

package Spotify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Subscription struct {
	Uri              *string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Expiry           *int32  `protobuf:"varint,2,opt,name=expiry" json:"expiry,omitempty"`
	StatusCode       *int32  `protobuf:"varint,3,opt,name=status_code" json:"status_code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Subscription) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Subscription) GetExpiry() int32 {
	if m != nil && m.Expiry != nil {
		return *m.Expiry
	}
	return 0
}

func (m *Subscription) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func init() {
	proto.RegisterType((*Subscription)(nil), "Spotify.Subscription")
}

var fileDescriptor6 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0x4d, 0x2a,
	0x2e, 0x4d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x0f, 0x2e, 0xc8, 0x2f, 0xc9, 0x4c,
	0xab, 0x54, 0x72, 0xe0, 0xe2, 0x09, 0x06, 0x4a, 0x24, 0x17, 0x65, 0x16, 0x94, 0x64, 0xe6, 0xe7,
	0x09, 0x71, 0x73, 0x31, 0x97, 0x16, 0x65, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x0a, 0xf1, 0x71,
	0xb1, 0xa5, 0x56, 0x14, 0x64, 0x16, 0x55, 0x4a, 0x30, 0x01, 0xf9, 0xac, 0x42, 0xc2, 0x5c, 0xdc,
	0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0xf1, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0xcc, 0x20, 0x41, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x24, 0x15, 0x70, 0x59, 0x00, 0x00, 0x00,
}