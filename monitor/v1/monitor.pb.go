// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/monitor/v1/monitor.proto

package monitor_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type Monitor struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProtocolKind         string   `protobuf:"bytes,3,opt,name=protocol_kind,json=protocolKind,proto3" json:"protocol_kind,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Method               string   `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Url                  string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	RequestTime          string   `protobuf:"bytes,7,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	ServerLocation       string   `protobuf:"bytes,8,opt,name=server_location,json=serverLocation,proto3" json:"server_location,omitempty"`
	Status               string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Enabled              bool     `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Monitor) Reset()         { *m = Monitor{} }
func (m *Monitor) String() string { return proto.CompactTextString(m) }
func (*Monitor) ProtoMessage()    {}
func (*Monitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a17b25f00b0e9c, []int{0}
}

func (m *Monitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Monitor.Unmarshal(m, b)
}
func (m *Monitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Monitor.Marshal(b, m, deterministic)
}
func (m *Monitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Monitor.Merge(m, src)
}
func (m *Monitor) XXX_Size() int {
	return xxx_messageInfo_Monitor.Size(m)
}
func (m *Monitor) XXX_DiscardUnknown() {
	xxx_messageInfo_Monitor.DiscardUnknown(m)
}

var xxx_messageInfo_Monitor proto.InternalMessageInfo

func (m *Monitor) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Monitor) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Monitor) GetProtocolKind() string {
	if m != nil {
		return m.ProtocolKind
	}
	return ""
}

func (m *Monitor) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Monitor) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Monitor) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Monitor) GetRequestTime() string {
	if m != nil {
		return m.RequestTime
	}
	return ""
}

func (m *Monitor) GetServerLocation() string {
	if m != nil {
		return m.ServerLocation
	}
	return ""
}

func (m *Monitor) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Monitor) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type ListMonitorsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMonitorsRequest) Reset()         { *m = ListMonitorsRequest{} }
func (m *ListMonitorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitorsRequest) ProtoMessage()    {}
func (*ListMonitorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a17b25f00b0e9c, []int{1}
}

func (m *ListMonitorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMonitorsRequest.Unmarshal(m, b)
}
func (m *ListMonitorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMonitorsRequest.Marshal(b, m, deterministic)
}
func (m *ListMonitorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMonitorsRequest.Merge(m, src)
}
func (m *ListMonitorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListMonitorsRequest.Size(m)
}
func (m *ListMonitorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMonitorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMonitorsRequest proto.InternalMessageInfo

type ListMonitorsResponse struct {
	Monitor              []*Monitor `protobuf:"bytes,1,rep,name=monitor,proto3" json:"monitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMonitorsResponse) Reset()         { *m = ListMonitorsResponse{} }
func (m *ListMonitorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitorsResponse) ProtoMessage()    {}
func (*ListMonitorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a17b25f00b0e9c, []int{2}
}

func (m *ListMonitorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMonitorsResponse.Unmarshal(m, b)
}
func (m *ListMonitorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMonitorsResponse.Marshal(b, m, deterministic)
}
func (m *ListMonitorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMonitorsResponse.Merge(m, src)
}
func (m *ListMonitorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListMonitorsResponse.Size(m)
}
func (m *ListMonitorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMonitorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMonitorsResponse proto.InternalMessageInfo

func (m *ListMonitorsResponse) GetMonitor() []*Monitor {
	if m != nil {
		return m.Monitor
	}
	return nil
}

type GetMonitorRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMonitorRequest) Reset()         { *m = GetMonitorRequest{} }
func (m *GetMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*GetMonitorRequest) ProtoMessage()    {}
func (*GetMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a17b25f00b0e9c, []int{3}
}

func (m *GetMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMonitorRequest.Unmarshal(m, b)
}
func (m *GetMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMonitorRequest.Marshal(b, m, deterministic)
}
func (m *GetMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMonitorRequest.Merge(m, src)
}
func (m *GetMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_GetMonitorRequest.Size(m)
}
func (m *GetMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMonitorRequest proto.InternalMessageInfo

func (m *GetMonitorRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateMonitorRequest struct {
	Monitor              *Monitor `protobuf:"bytes,1,opt,name=monitor,proto3" json:"monitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMonitorRequest) Reset()         { *m = CreateMonitorRequest{} }
func (m *CreateMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMonitorRequest) ProtoMessage()    {}
func (*CreateMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a17b25f00b0e9c, []int{4}
}

func (m *CreateMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMonitorRequest.Unmarshal(m, b)
}
func (m *CreateMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMonitorRequest.Marshal(b, m, deterministic)
}
func (m *CreateMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMonitorRequest.Merge(m, src)
}
func (m *CreateMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMonitorRequest.Size(m)
}
func (m *CreateMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMonitorRequest proto.InternalMessageInfo

func (m *CreateMonitorRequest) GetMonitor() *Monitor {
	if m != nil {
		return m.Monitor
	}
	return nil
}

type UpdateMonitorRequest struct {
	Monitor              *Monitor              `protobuf:"bytes,1,opt,name=monitor,proto3" json:"monitor,omitempty"`
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateMonitorRequest) Reset()         { *m = UpdateMonitorRequest{} }
func (m *UpdateMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMonitorRequest) ProtoMessage()    {}
func (*UpdateMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a17b25f00b0e9c, []int{5}
}

func (m *UpdateMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMonitorRequest.Unmarshal(m, b)
}
func (m *UpdateMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMonitorRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMonitorRequest.Merge(m, src)
}
func (m *UpdateMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMonitorRequest.Size(m)
}
func (m *UpdateMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMonitorRequest proto.InternalMessageInfo

func (m *UpdateMonitorRequest) GetMonitor() *Monitor {
	if m != nil {
		return m.Monitor
	}
	return nil
}

func (m *UpdateMonitorRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteMonitorRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMonitorRequest) Reset()         { *m = DeleteMonitorRequest{} }
func (m *DeleteMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMonitorRequest) ProtoMessage()    {}
func (*DeleteMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a17b25f00b0e9c, []int{6}
}

func (m *DeleteMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMonitorRequest.Unmarshal(m, b)
}
func (m *DeleteMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMonitorRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMonitorRequest.Merge(m, src)
}
func (m *DeleteMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMonitorRequest.Size(m)
}
func (m *DeleteMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMonitorRequest proto.InternalMessageInfo

func (m *DeleteMonitorRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Monitor)(nil), "monitor.v1.Monitor")
	proto.RegisterType((*ListMonitorsRequest)(nil), "monitor.v1.ListMonitorsRequest")
	proto.RegisterType((*ListMonitorsResponse)(nil), "monitor.v1.ListMonitorsResponse")
	proto.RegisterType((*GetMonitorRequest)(nil), "monitor.v1.GetMonitorRequest")
	proto.RegisterType((*CreateMonitorRequest)(nil), "monitor.v1.CreateMonitorRequest")
	proto.RegisterType((*UpdateMonitorRequest)(nil), "monitor.v1.UpdateMonitorRequest")
	proto.RegisterType((*DeleteMonitorRequest)(nil), "monitor.v1.DeleteMonitorRequest")
}

func init() { proto.RegisterFile("proto/monitor/v1/monitor.proto", fileDescriptor_30a17b25f00b0e9c) }

var fileDescriptor_30a17b25f00b0e9c = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0x54, 0x12, 0x36, 0xd9, 0x7d, 0x59, 0x16, 0xf0, 0x06, 0xb0, 0xf6, 0x80, 0x42, 0x2a, 0x41,
	0x2e, 0x24, 0x6a, 0x39, 0x72, 0xe4, 0x4b, 0x88, 0xf6, 0x12, 0xc1, 0x39, 0x4a, 0xeb, 0xd7, 0x62,
	0xd5, 0x89, 0x8b, 0xed, 0xf4, 0x07, 0xf0, 0x1f, 0xf8, 0xbf, 0x28, 0x76, 0x22, 0xd4, 0x72, 0xe0,
	0xb0, 0xb7, 0x37, 0xf3, 0xc6, 0x33, 0xed, 0xcb, 0xc0, 0x8b, 0x83, 0x92, 0x46, 0x96, 0xad, 0xec,
	0xb8, 0x91, 0xaa, 0x3c, 0xce, 0xa7, 0xb1, 0xb0, 0x0b, 0x02, 0x13, 0x3c, 0xce, 0xef, 0xd2, 0x9d,
	0x94, 0x3b, 0x81, 0xa5, 0xdd, 0xac, 0xfb, 0x6d, 0xb9, 0xe5, 0x28, 0x58, 0xdd, 0x36, 0x7a, 0xef,
	0xd4, 0xd9, 0x6f, 0x1f, 0xa2, 0x95, 0x7b, 0x40, 0x6e, 0xc0, 0xe7, 0x8c, 0x7a, 0xa9, 0x97, 0x07,
	0x95, 0xcf, 0x19, 0x79, 0x0e, 0x51, 0xaf, 0x51, 0xd5, 0x9c, 0x51, 0xdf, 0x92, 0xe1, 0x00, 0xbf,
	0x30, 0x32, 0x83, 0x87, 0xf6, 0xf5, 0x46, 0x8a, 0x7a, 0xcf, 0x3b, 0x46, 0x83, 0xd4, 0xcb, 0xaf,
	0xaa, 0xeb, 0x89, 0xfc, 0xca, 0x3b, 0x46, 0x12, 0xb8, 0x30, 0xdc, 0x08, 0xa4, 0x0f, 0xec, 0xd2,
	0x01, 0xf2, 0x0c, 0xc2, 0x16, 0xcd, 0x0f, 0xc9, 0xe8, 0x85, 0xa5, 0x47, 0x44, 0x1e, 0x43, 0xd0,
	0x2b, 0x41, 0x43, 0x4b, 0x0e, 0x23, 0x79, 0x09, 0xd7, 0x0a, 0x7f, 0xf6, 0xa8, 0x4d, 0x6d, 0x78,
	0x8b, 0x34, 0xb2, 0xab, 0x78, 0xe4, 0xbe, 0xf1, 0x16, 0xc9, 0x6b, 0x78, 0xa4, 0x51, 0x1d, 0x51,
	0xd5, 0x42, 0x6e, 0x1a, 0xc3, 0x65, 0x47, 0x2f, 0xad, 0xea, 0xc6, 0xd1, 0xcb, 0x91, 0x1d, 0x52,
	0xb5, 0x69, 0x4c, 0xaf, 0xe9, 0x95, 0x4b, 0x75, 0x88, 0x50, 0x88, 0xb0, 0x6b, 0xd6, 0x02, 0x19,
	0x85, 0xd4, 0xcb, 0x2f, 0xab, 0x09, 0x66, 0x4f, 0xe1, 0x76, 0xc9, 0xb5, 0x19, 0x4f, 0xa3, 0x2b,
	0x97, 0x9a, 0x7d, 0x84, 0xe4, 0x94, 0xd6, 0x07, 0xd9, 0x69, 0x24, 0x6f, 0x20, 0x1a, 0xcf, 0x4e,
	0xbd, 0x34, 0xc8, 0xe3, 0xc5, 0x6d, 0xf1, 0xf7, 0x33, 0x14, 0xa3, 0xbc, 0x9a, 0x34, 0xd9, 0x0c,
	0x9e, 0x7c, 0xc6, 0xc9, 0x65, 0xf4, 0x3e, 0x3f, 0xff, 0x90, 0xf5, 0x5e, 0x61, 0x63, 0xf0, 0x4c,
	0x77, 0x92, 0xe5, 0xfd, 0x37, 0xeb, 0x97, 0x07, 0xc9, 0xf7, 0x03, 0xbb, 0xaf, 0x0f, 0x79, 0x07,
	0x71, 0x6f, 0x6d, 0x6c, 0x7d, 0x6c, 0x23, 0xe2, 0xc5, 0x5d, 0xe1, 0x1a, 0x56, 0x4c, 0x0d, 0x2b,
	0x3e, 0x0d, 0x0d, 0x5b, 0x35, 0x7a, 0x5f, 0x81, 0x93, 0x0f, 0x73, 0xf6, 0x0a, 0x92, 0x0f, 0x28,
	0xf0, 0x9f, 0xdf, 0x70, 0xf6, 0x9f, 0xd7, 0xa1, 0xf5, 0x79, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0x61, 0x08, 0x11, 0xe5, 0x02, 0x00, 0x00,
}
