// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/monitor/v1/trigger.proto

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

type Trigger struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MonitorId            int64     `protobuf:"varint,3,opt,name=monitor_id,json=monitorId,proto3" json:"monitor_id,omitempty"`
	TriggerKind          string    `protobuf:"bytes,4,opt,name=trigger_kind,json=triggerKind,proto3" json:"trigger_kind,omitempty"`
	NotificationKind     string    `protobuf:"bytes,5,opt,name=notification_kind,json=notificationKind,proto3" json:"notification_kind,omitempty"`
	ErrTolerance         string    `protobuf:"bytes,6,opt,name=err_tolerance,json=errTolerance,proto3" json:"err_tolerance,omitempty"`
	Warning              string    `protobuf:"bytes,7,opt,name=warning,proto3" json:"warning,omitempty"`
	Status               string    `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Enabled              bool      `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	HttpStat             *HttpStat `protobuf:"bytes,10,opt,name=httpStat,proto3" json:"httpStat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{0}
}

func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (m *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(m, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

func (m *Trigger) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Trigger) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Trigger) GetMonitorId() int64 {
	if m != nil {
		return m.MonitorId
	}
	return 0
}

func (m *Trigger) GetTriggerKind() string {
	if m != nil {
		return m.TriggerKind
	}
	return ""
}

func (m *Trigger) GetNotificationKind() string {
	if m != nil {
		return m.NotificationKind
	}
	return ""
}

func (m *Trigger) GetErrTolerance() string {
	if m != nil {
		return m.ErrTolerance
	}
	return ""
}

func (m *Trigger) GetWarning() string {
	if m != nil {
		return m.Warning
	}
	return ""
}

func (m *Trigger) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Trigger) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Trigger) GetHttpStat() *HttpStat {
	if m != nil {
		return m.HttpStat
	}
	return nil
}

type HttpStat struct {
	Dns                  int32    `protobuf:"varint,1,opt,name=dns,proto3" json:"dns,omitempty"`
	Tcp                  int32    `protobuf:"varint,2,opt,name=tcp,proto3" json:"tcp,omitempty"`
	Tls                  int32    `protobuf:"varint,3,opt,name=tls,proto3" json:"tls,omitempty"`
	Processing           int32    `protobuf:"varint,4,opt,name=processing,proto3" json:"processing,omitempty"`
	Transfer             int32    `protobuf:"varint,5,opt,name=transfer,proto3" json:"transfer,omitempty"`
	StatusCode           int32    `protobuf:"varint,6,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Status               string   `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpStat) Reset()         { *m = HttpStat{} }
func (m *HttpStat) String() string { return proto.CompactTextString(m) }
func (*HttpStat) ProtoMessage()    {}
func (*HttpStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{1}
}

func (m *HttpStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpStat.Unmarshal(m, b)
}
func (m *HttpStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpStat.Marshal(b, m, deterministic)
}
func (m *HttpStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStat.Merge(m, src)
}
func (m *HttpStat) XXX_Size() int {
	return xxx_messageInfo_HttpStat.Size(m)
}
func (m *HttpStat) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStat.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStat proto.InternalMessageInfo

func (m *HttpStat) GetDns() int32 {
	if m != nil {
		return m.Dns
	}
	return 0
}

func (m *HttpStat) GetTcp() int32 {
	if m != nil {
		return m.Tcp
	}
	return 0
}

func (m *HttpStat) GetTls() int32 {
	if m != nil {
		return m.Tls
	}
	return 0
}

func (m *HttpStat) GetProcessing() int32 {
	if m != nil {
		return m.Processing
	}
	return 0
}

func (m *HttpStat) GetTransfer() int32 {
	if m != nil {
		return m.Transfer
	}
	return 0
}

func (m *HttpStat) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *HttpStat) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ListTriggersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTriggersRequest) Reset()         { *m = ListTriggersRequest{} }
func (m *ListTriggersRequest) String() string { return proto.CompactTextString(m) }
func (*ListTriggersRequest) ProtoMessage()    {}
func (*ListTriggersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{2}
}

func (m *ListTriggersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTriggersRequest.Unmarshal(m, b)
}
func (m *ListTriggersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTriggersRequest.Marshal(b, m, deterministic)
}
func (m *ListTriggersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTriggersRequest.Merge(m, src)
}
func (m *ListTriggersRequest) XXX_Size() int {
	return xxx_messageInfo_ListTriggersRequest.Size(m)
}
func (m *ListTriggersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTriggersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTriggersRequest proto.InternalMessageInfo

type ListTriggersResponse struct {
	Triggers             []*Trigger `protobuf:"bytes,1,rep,name=triggers,proto3" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListTriggersResponse) Reset()         { *m = ListTriggersResponse{} }
func (m *ListTriggersResponse) String() string { return proto.CompactTextString(m) }
func (*ListTriggersResponse) ProtoMessage()    {}
func (*ListTriggersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{3}
}

func (m *ListTriggersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTriggersResponse.Unmarshal(m, b)
}
func (m *ListTriggersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTriggersResponse.Marshal(b, m, deterministic)
}
func (m *ListTriggersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTriggersResponse.Merge(m, src)
}
func (m *ListTriggersResponse) XXX_Size() int {
	return xxx_messageInfo_ListTriggersResponse.Size(m)
}
func (m *ListTriggersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTriggersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTriggersResponse proto.InternalMessageInfo

func (m *ListTriggersResponse) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type GetTriggerRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTriggerRequest) Reset()         { *m = GetTriggerRequest{} }
func (m *GetTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*GetTriggerRequest) ProtoMessage()    {}
func (*GetTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{4}
}

func (m *GetTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTriggerRequest.Unmarshal(m, b)
}
func (m *GetTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTriggerRequest.Marshal(b, m, deterministic)
}
func (m *GetTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTriggerRequest.Merge(m, src)
}
func (m *GetTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_GetTriggerRequest.Size(m)
}
func (m *GetTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTriggerRequest proto.InternalMessageInfo

func (m *GetTriggerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateTriggerRequest struct {
	Trigger              *Trigger `protobuf:"bytes,1,opt,name=trigger,proto3" json:"trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTriggerRequest) Reset()         { *m = CreateTriggerRequest{} }
func (m *CreateTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTriggerRequest) ProtoMessage()    {}
func (*CreateTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{5}
}

func (m *CreateTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTriggerRequest.Unmarshal(m, b)
}
func (m *CreateTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTriggerRequest.Marshal(b, m, deterministic)
}
func (m *CreateTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTriggerRequest.Merge(m, src)
}
func (m *CreateTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTriggerRequest.Size(m)
}
func (m *CreateTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTriggerRequest proto.InternalMessageInfo

func (m *CreateTriggerRequest) GetTrigger() *Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

type UpdateTriggerRequest struct {
	Trigger              *Trigger              `protobuf:"bytes,1,opt,name=trigger,proto3" json:"trigger,omitempty"`
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateTriggerRequest) Reset()         { *m = UpdateTriggerRequest{} }
func (m *UpdateTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTriggerRequest) ProtoMessage()    {}
func (*UpdateTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{6}
}

func (m *UpdateTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTriggerRequest.Unmarshal(m, b)
}
func (m *UpdateTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTriggerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTriggerRequest.Merge(m, src)
}
func (m *UpdateTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTriggerRequest.Size(m)
}
func (m *UpdateTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTriggerRequest proto.InternalMessageInfo

func (m *UpdateTriggerRequest) GetTrigger() *Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *UpdateTriggerRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteTriggerRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTriggerRequest) Reset()         { *m = DeleteTriggerRequest{} }
func (m *DeleteTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTriggerRequest) ProtoMessage()    {}
func (*DeleteTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4797f3ebc25aa022, []int{7}
}

func (m *DeleteTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTriggerRequest.Unmarshal(m, b)
}
func (m *DeleteTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTriggerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTriggerRequest.Merge(m, src)
}
func (m *DeleteTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTriggerRequest.Size(m)
}
func (m *DeleteTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTriggerRequest proto.InternalMessageInfo

func (m *DeleteTriggerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Trigger)(nil), "monitor.v1.Trigger")
	proto.RegisterType((*HttpStat)(nil), "monitor.v1.HttpStat")
	proto.RegisterType((*ListTriggersRequest)(nil), "monitor.v1.ListTriggersRequest")
	proto.RegisterType((*ListTriggersResponse)(nil), "monitor.v1.ListTriggersResponse")
	proto.RegisterType((*GetTriggerRequest)(nil), "monitor.v1.GetTriggerRequest")
	proto.RegisterType((*CreateTriggerRequest)(nil), "monitor.v1.CreateTriggerRequest")
	proto.RegisterType((*UpdateTriggerRequest)(nil), "monitor.v1.UpdateTriggerRequest")
	proto.RegisterType((*DeleteTriggerRequest)(nil), "monitor.v1.DeleteTriggerRequest")
}

func init() { proto.RegisterFile("proto/monitor/v1/trigger.proto", fileDescriptor_4797f3ebc25aa022) }

var fileDescriptor_4797f3ebc25aa022 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x76, 0x6d, 0xda, 0x97, 0x81, 0x36, 0xaf, 0x40, 0x54, 0x89, 0x11, 0x32, 0x09,
	0x55, 0x42, 0x24, 0xac, 0x1c, 0x39, 0x0e, 0x18, 0x13, 0x70, 0x31, 0xe3, 0x5c, 0xa5, 0xf5, 0x6b,
	0xb1, 0x9a, 0xd9, 0xc1, 0x76, 0xc7, 0x9d, 0xef, 0xc4, 0x37, 0xe0, 0x83, 0x21, 0xbf, 0x24, 0xa5,
	0x14, 0x71, 0xda, 0xcd, 0xef, 0xff, 0xff, 0xe5, 0xd9, 0xfe, 0xc7, 0x0f, 0x4e, 0x2b, 0xa3, 0x9d,
	0xce, 0x6f, 0xb4, 0x92, 0x4e, 0x9b, 0xfc, 0xf6, 0x3c, 0x77, 0x46, 0xae, 0x56, 0x68, 0x32, 0x32,
	0x18, 0x34, 0x4e, 0x76, 0x7b, 0x3e, 0x4e, 0x56, 0x5a, 0xaf, 0x4a, 0xcc, 0xc9, 0x99, 0x6f, 0x96,
	0xf9, 0x52, 0x62, 0x29, 0x66, 0x37, 0x85, 0x5d, 0xd7, 0x74, 0xfa, 0xab, 0x03, 0xe1, 0x75, 0xfd,
	0x3d, 0xbb, 0x0f, 0x1d, 0x29, 0xe2, 0x20, 0x09, 0x26, 0x5d, 0xde, 0x91, 0x82, 0x3d, 0x82, 0x70,
	0x63, 0xd1, 0xcc, 0xa4, 0x88, 0x3b, 0x24, 0xf6, 0x7d, 0x79, 0x25, 0xd8, 0x63, 0x68, 0x37, 0xf1,
	0x5e, 0x97, 0xbc, 0x61, 0xa3, 0x5c, 0x09, 0xf6, 0x14, 0x0e, 0x9b, 0x23, 0xcd, 0xd6, 0x52, 0x89,
	0xf8, 0x20, 0x09, 0x26, 0x43, 0x1e, 0x35, 0xda, 0x07, 0xa9, 0x04, 0x7b, 0x0e, 0xc7, 0x4a, 0x3b,
	0xb9, 0x94, 0x8b, 0xc2, 0x49, 0xad, 0x6a, 0xae, 0x47, 0xdc, 0xd1, 0xae, 0x41, 0xf0, 0x19, 0xdc,
	0x43, 0x63, 0x66, 0x4e, 0x97, 0x68, 0x0a, 0xb5, 0xc0, 0xb8, 0x4f, 0xe0, 0x21, 0x1a, 0x73, 0xdd,
	0x6a, 0x2c, 0x86, 0xf0, 0x7b, 0x61, 0x94, 0x54, 0xab, 0x38, 0x24, 0xbb, 0x2d, 0xd9, 0x43, 0xe8,
	0x5b, 0x57, 0xb8, 0x8d, 0x8d, 0x07, 0x64, 0x34, 0x95, 0xff, 0x02, 0x55, 0x31, 0x2f, 0x51, 0xc4,
	0xc3, 0x24, 0x98, 0x0c, 0x78, 0x5b, 0xb2, 0x97, 0x30, 0xf8, 0xea, 0x5c, 0xf5, 0xd9, 0x15, 0x2e,
	0x86, 0x24, 0x98, 0x44, 0xd3, 0x51, 0xf6, 0x27, 0xd5, 0xec, 0x7d, 0xe3, 0xf1, 0x2d, 0x95, 0xfe,
	0x0c, 0x60, 0xd0, 0xca, 0xec, 0x08, 0xba, 0x42, 0x59, 0x0a, 0xb2, 0xc7, 0xfd, 0xd2, 0x2b, 0x6e,
	0x51, 0x51, 0x8a, 0x3d, 0xee, 0x97, 0xa4, 0x94, 0x96, 0xb2, 0xf3, 0x4a, 0x69, 0xd9, 0x29, 0x40,
	0x65, 0xf4, 0x02, 0xad, 0xf5, 0x77, 0x38, 0x20, 0x63, 0x47, 0x61, 0x63, 0x18, 0x38, 0x53, 0x28,
	0xbb, 0x44, 0x43, 0x49, 0xf5, 0xf8, 0xb6, 0x66, 0x4f, 0x20, 0xaa, 0x2f, 0x35, 0x5b, 0x68, 0x51,
	0xe7, 0xd3, 0xe3, 0x50, 0x4b, 0x17, 0x5a, 0xe0, 0x4e, 0x06, 0xe1, 0x6e, 0x06, 0xe9, 0x03, 0x38,
	0xf9, 0x28, 0xad, 0x6b, 0x5e, 0x80, 0xe5, 0xf8, 0x6d, 0x83, 0xd6, 0xa5, 0x97, 0x30, 0xfa, 0x5b,
	0xb6, 0x95, 0x56, 0x16, 0x59, 0xee, 0xcf, 0x50, 0x6b, 0x71, 0x90, 0x74, 0x27, 0xd1, 0xf4, 0x64,
	0x37, 0x98, 0x86, 0xe7, 0x5b, 0x28, 0x3d, 0x83, 0xe3, 0x4b, 0x6c, 0xfb, 0x34, 0xdd, 0xf7, 0xdf,
	0x59, 0xfa, 0x16, 0x46, 0x17, 0x06, 0x0b, 0x87, 0x7b, 0xdc, 0x0b, 0x08, 0x9b, 0x46, 0x04, 0xff,
	0x67, 0xb3, 0x96, 0x49, 0x7f, 0x04, 0x30, 0xfa, 0x52, 0x89, 0xbb, 0xf6, 0x61, 0xaf, 0x21, 0xda,
	0x50, 0x1b, 0x9a, 0x13, 0xfa, 0x69, 0xd1, 0x74, 0x9c, 0xd5, 0xa3, 0x94, 0xb5, 0xa3, 0x94, 0xbd,
	0xf3, 0xa3, 0xf4, 0xa9, 0xb0, 0x6b, 0x0e, 0x35, 0xee, 0xd7, 0xe9, 0x33, 0x18, 0xbd, 0xc1, 0x12,
	0xff, 0x39, 0xc3, 0xde, 0x9d, 0xe7, 0x7d, 0xea, 0xf3, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x40, 0x07, 0xe3, 0x5f, 0xce, 0x03, 0x00, 0x00,
}
