// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hostapd.proto

/*
Package hostapd is a generated protocol buffer package.

It is generated from these files:
	hostapd.proto

It has these top-level messages:
	PingRequest
	PongResponse
	ListRequest
	Client
	ListResponse
*/
package hostapd

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

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PongResponse struct {
}

func (m *PongResponse) Reset()                    { *m = PongResponse{} }
func (m *PongResponse) String() string            { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()               {}
func (*PongResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Client struct {
	Addr          string   `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Flag          []string `protobuf:"bytes,2,rep,name=flag" json:"flag,omitempty"`
	ConnectedTime uint32   `protobuf:"varint,3,opt,name=connected_time,json=connectedTime" json:"connected_time,omitempty"`
	IdleMsec      uint32   `protobuf:"varint,4,opt,name=idle_msec,json=idleMsec" json:"idle_msec,omitempty"`
	RxPackets     uint32   `protobuf:"varint,5,opt,name=rx_packets,json=rxPackets" json:"rx_packets,omitempty"`
	TxPackets     uint32   `protobuf:"varint,6,opt,name=tx_packets,json=txPackets" json:"tx_packets,omitempty"`
	RxBytes       uint32   `protobuf:"varint,7,opt,name=rx_bytes,json=rxBytes" json:"rx_bytes,omitempty"`
	TxBytes       uint32   `protobuf:"varint,8,opt,name=tx_bytes,json=txBytes" json:"tx_bytes,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Client) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Client) GetFlag() []string {
	if m != nil {
		return m.Flag
	}
	return nil
}

func (m *Client) GetConnectedTime() uint32 {
	if m != nil {
		return m.ConnectedTime
	}
	return 0
}

func (m *Client) GetIdleMsec() uint32 {
	if m != nil {
		return m.IdleMsec
	}
	return 0
}

func (m *Client) GetRxPackets() uint32 {
	if m != nil {
		return m.RxPackets
	}
	return 0
}

func (m *Client) GetTxPackets() uint32 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *Client) GetRxBytes() uint32 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *Client) GetTxBytes() uint32 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

type ListResponse struct {
	Client []*Client `protobuf:"bytes,1,rep,name=client" json:"client,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListResponse) GetClient() []*Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "hostapd.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "hostapd.PongResponse")
	proto.RegisterType((*ListRequest)(nil), "hostapd.ListRequest")
	proto.RegisterType((*Client)(nil), "hostapd.Client")
	proto.RegisterType((*ListResponse)(nil), "hostapd.ListResponse")
}

func init() { proto.RegisterFile("hostapd.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x9b, 0xaf, 0xfd, 0xd2, 0x76, 0xda, 0x54, 0x58, 0x14, 0x56, 0x45, 0x08, 0x01, 0xb1,
	0xa7, 0x1e, 0x2a, 0xe2, 0xc5, 0x93, 0xbd, 0x78, 0x50, 0x28, 0xc1, 0x7b, 0x48, 0x93, 0xb1, 0x06,
	0xd3, 0xdd, 0xb8, 0x3b, 0x42, 0x7c, 0x00, 0xdf, 0xd5, 0xc7, 0x90, 0xdd, 0x6d, 0xd3, 0xbd, 0xcd,
	0xfc, 0x7f, 0x99, 0xb0, 0xf3, 0x1b, 0x88, 0xde, 0xa5, 0xa6, 0xbc, 0x29, 0x17, 0x8d, 0x92, 0x24,
	0xd9, 0x70, 0xdf, 0x26, 0x11, 0x4c, 0xd6, 0x95, 0xd8, 0xa6, 0xf8, 0xf9, 0x85, 0x9a, 0x92, 0x19,
	0x4c, 0xd7, 0xd2, 0xb4, 0xba, 0x91, 0x42, 0xa3, 0xc1, 0xcf, 0x95, 0xa6, 0x03, 0xfe, 0x0d, 0x20,
	0x5c, 0xd5, 0x15, 0x0a, 0x62, 0x0c, 0x06, 0x79, 0x59, 0x2a, 0x1e, 0xc4, 0xc1, 0x7c, 0x9c, 0xda,
	0xda, 0x64, 0x6f, 0x75, 0xbe, 0xe5, 0xff, 0xe2, 0xbe, 0xc9, 0x4c, 0xcd, 0xae, 0x61, 0x56, 0x48,
	0x21, 0xb0, 0x20, 0x2c, 0x33, 0xaa, 0x76, 0xc8, 0xfb, 0x71, 0x30, 0x8f, 0xd2, 0xa8, 0x4b, 0x5f,
	0xab, 0x1d, 0xb2, 0x4b, 0x18, 0x57, 0x65, 0x8d, 0xd9, 0x4e, 0x63, 0xc1, 0x07, 0xf6, 0x8b, 0x91,
	0x09, 0x5e, 0x34, 0x16, 0xec, 0x0a, 0x40, 0xb5, 0x59, 0x93, 0x17, 0x1f, 0x48, 0x9a, 0xff, 0xb7,
	0x74, 0xac, 0xda, 0xb5, 0x0b, 0x0c, 0xa6, 0x23, 0x0e, 0x1d, 0xa6, 0x0e, 0x9f, 0xc3, 0x48, 0xb5,
	0xd9, 0xe6, 0x9b, 0x50, 0xf3, 0xa1, 0x85, 0x43, 0xd5, 0x3e, 0x9a, 0xd6, 0x20, 0x3a, 0xa0, 0x91,
	0x43, 0xe4, 0x50, 0x72, 0x0f, 0x53, 0xb7, 0xb9, 0x33, 0xc1, 0x6e, 0x20, 0x2c, 0xec, 0xe6, 0x3c,
	0x88, 0xfb, 0xf3, 0xc9, 0xf2, 0x64, 0x71, 0x30, 0xea, 0x84, 0xa4, 0x7b, 0xbc, 0xfc, 0x09, 0x60,
	0xf6, 0xe4, 0xd0, 0x4a, 0x0a, 0x52, 0xb2, 0x66, 0x77, 0x30, 0x30, 0x92, 0xd9, 0x69, 0x37, 0xe3,
	0x39, 0xbf, 0x38, 0x3b, 0xa6, 0xbe, 0xfa, 0x1e, 0x7b, 0x70, 0xf2, 0xdd, 0xff, 0xb5, 0x37, 0xed,
	0x9d, 0xc4, 0x9b, 0xf6, 0x9f, 0x9b, 0xf4, 0x36, 0xa1, 0xbd, 0xf4, 0xed, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9c, 0x41, 0x95, 0xe6, 0xfa, 0x01, 0x00, 0x00,
}
