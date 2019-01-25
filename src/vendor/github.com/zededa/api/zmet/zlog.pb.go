// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zlog.proto

/*
Package zmet is a generated protocol buffer package.

It is generated from these files:
	zlog.proto
	zmet.proto
	zregister.proto

It has these top-level messages:
	LogEntry
	LogBundle
	ZInfoManufacturer
	ZInfoNetwork
	GeoLoc
	ZInfoDNS
	ZinfoPeripheral
	ZInfoSW
	ErrorInfo
	ZInfoDevice
	ZInfoDevSW
	ZInfoStorage
	ZInfoApp
	ZInfoVpnLinkInfo
	ZInfoVpnLink
	ZInfoVpnEndPoint
	ZInfoVpnConn
	ZInfoVpn
	RlocState
	MapCacheEntry
	DatabaseMap
	DecapKey
	ZInfoLisp
	ZInfoService
	ZInfoNetworkObject
	ZInfoNetworkInstance
	ZmetIPAssignmentEntry
	ZmetVifInfo
	ZInfoMsg
	ZioBundle
	MemoryMetric
	NetworkMetric
	ZedcloudMetric
	UrlcloudMetric
	AppCpuMetric
	DeviceMetric
	MetricItem
	DiskMetric
	AppDiskMetric
	AppMetric
	PktStat
	RlocStats
	EidStats
	ZMetricLisp
	ZMetricConn
	ZMetricVpn
	ZMetricNone
	ZMetricFlowLink
	ZMetricFlowEndPoint
	ZMetricFlow
	ZMetricLispGlobal
	ZMetricService
	ZMetricNetworkInstance
	ZMetricMsg
	ZRegisterResp
	ZRegisterMsg
*/
package zmet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogEntry struct {
	Severity  string                     `protobuf:"bytes,1,opt,name=severity" json:"severity,omitempty"`
	Source    string                     `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Iid       string                     `protobuf:"bytes,3,opt,name=iid" json:"iid,omitempty"`
	Content   string                     `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	Msgid     uint64                     `protobuf:"varint,5,opt,name=msgid" json:"msgid,omitempty"`
	Tags      map[string]string          `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogEntry) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *LogEntry) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *LogEntry) GetIid() string {
	if m != nil {
		return m.Iid
	}
	return ""
}

func (m *LogEntry) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *LogEntry) GetMsgid() uint64 {
	if m != nil {
		return m.Msgid
	}
	return 0
}

func (m *LogEntry) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *LogEntry) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type LogBundle struct {
	DevID     string                     `protobuf:"bytes,1,opt,name=devID" json:"devID,omitempty"`
	Image     string                     `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Log       []*LogEntry                `protobuf:"bytes,3,rep,name=log" json:"log,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *LogBundle) Reset()                    { *m = LogBundle{} }
func (m *LogBundle) String() string            { return proto.CompactTextString(m) }
func (*LogBundle) ProtoMessage()               {}
func (*LogBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogBundle) GetDevID() string {
	if m != nil {
		return m.DevID
	}
	return ""
}

func (m *LogBundle) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *LogBundle) GetLog() []*LogEntry {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *LogBundle) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "LogEntry")
	proto.RegisterType((*LogBundle)(nil), "LogBundle")
}

func init() { proto.RegisterFile("zlog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xbf, 0x6a, 0xeb, 0x30,
	0x14, 0xc6, 0x71, 0xec, 0xfc, 0xf1, 0xc9, 0x72, 0xd1, 0xbd, 0x5c, 0x84, 0xef, 0x10, 0x93, 0xe5,
	0x7a, 0x92, 0x21, 0x1d, 0x1a, 0x0a, 0x5d, 0x42, 0x3b, 0x14, 0x32, 0x99, 0x4c, 0xdd, 0x1c, 0x4b,
	0x55, 0x45, 0x2d, 0x2b, 0x58, 0x92, 0x21, 0x79, 0x88, 0x3e, 0x41, 0x1f, 0xb6, 0x48, 0x8e, 0x5d,
	0x3a, 0x76, 0x3b, 0xbf, 0xa3, 0xef, 0x88, 0xef, 0x07, 0x70, 0xa9, 0x15, 0x27, 0xa7, 0x56, 0x19,
	0x95, 0xac, 0xb8, 0x52, 0xbc, 0x66, 0xb9, 0xa7, 0xa3, 0x7d, 0xc9, 0x8d, 0x90, 0x4c, 0x9b, 0x52,
	0x9e, 0xfa, 0xc0, 0xfa, 0x63, 0x02, 0x8b, 0xbd, 0xe2, 0x8f, 0x8d, 0x69, 0xcf, 0x28, 0x81, 0x85,
	0x66, 0x1d, 0x6b, 0x85, 0x39, 0xe3, 0x20, 0x0d, 0xb2, 0xb8, 0x18, 0x19, 0xfd, 0x85, 0x99, 0x56,
	0xb6, 0xad, 0x18, 0x9e, 0xf8, 0x97, 0x2b, 0xa1, 0x5f, 0x10, 0x0a, 0x41, 0x71, 0xe8, 0x97, 0x6e,
	0x44, 0x18, 0xe6, 0x95, 0x6a, 0x0c, 0x6b, 0x0c, 0x8e, 0xfc, 0x76, 0x40, 0xf4, 0x07, 0xa6, 0x52,
	0x73, 0x41, 0xf1, 0x34, 0x0d, 0xb2, 0xa8, 0xe8, 0x01, 0xfd, 0x87, 0xc8, 0x94, 0x5c, 0xe3, 0x59,
	0x1a, 0x66, 0xcb, 0xcd, 0x6f, 0x32, 0xd4, 0x21, 0x87, 0x92, 0x6b, 0x3f, 0x15, 0x3e, 0x80, 0xb6,
	0x10, 0x8f, 0xf5, 0xf1, 0x3c, 0x0d, 0xb2, 0xe5, 0x26, 0x21, 0xbd, 0x20, 0x19, 0x04, 0xc9, 0x61,
	0x48, 0x14, 0x5f, 0xe1, 0xe4, 0x16, 0xe2, 0xf1, 0x33, 0xd7, 0xf8, 0x8d, 0x0d, 0x82, 0x6e, 0x74,
	0xbd, 0xba, 0xb2, 0xb6, 0x83, 0x5a, 0x0f, 0x77, 0x93, 0x6d, 0xb0, 0x7e, 0x0f, 0x20, 0xde, 0x2b,
	0xbe, 0xb3, 0x0d, 0xad, 0x99, 0xcb, 0x51, 0xd6, 0x3d, 0x3d, 0x5c, 0x6f, 0x7b, 0x70, 0x5b, 0x21,
	0x4b, 0x3e, 0x5e, 0x7b, 0x40, 0xff, 0x20, 0xac, 0x15, 0xc7, 0xa1, 0x97, 0x8a, 0x47, 0xa9, 0xc2,
	0x6d, 0xbf, 0x9b, 0x44, 0x3f, 0x30, 0xd9, 0xdd, 0xc3, 0xaa, 0x52, 0x92, 0x5c, 0x18, 0x65, 0xb4,
	0x24, 0x55, 0xad, 0x2c, 0x25, 0x56, 0xb3, 0xb6, 0x13, 0xd5, 0xf5, 0xf4, 0x39, 0xe1, 0xc2, 0xbc,
	0xda, 0x23, 0xa9, 0x94, 0xcc, 0xfb, 0x5c, 0x5e, 0x9e, 0x44, 0x7e, 0x91, 0xcc, 0x1c, 0x67, 0x3e,
	0x72, 0xf3, 0x19, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xce, 0xd4, 0xcd, 0x24, 0x02, 0x00, 0x00,
}
