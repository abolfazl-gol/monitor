// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/monitor/v1/service.proto

package monitor_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("proto/monitor/v1/service.proto", fileDescriptor_83dceab02d1ecd67) }

var fileDescriptor_83dceab02d1ecd67 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x77, 0x26, 0xde, 0xc8, 0x66, 0x88, 0x2c, 0xf0, 0x77, 0x65, 0x60, 0x33, 0x0d, 0xfa,
	0x02, 0x26, 0xfe, 0x10, 0x13, 0x35, 0xd1, 0xc8, 0x8a, 0x85, 0x01, 0xbc, 0x34, 0x93, 0x40, 0xa7,
	0xce, 0x0c, 0x24, 0xbe, 0xaa, 0x4f, 0x63, 0xe8, 0xf4, 0x42, 0xa7, 0x33, 0x2d, 0x8d, 0xcb, 0x9e,
	0xef, 0xf4, 0x5c, 0x38, 0x27, 0x85, 0xf3, 0x54, 0x49, 0x23, 0xa3, 0xa5, 0x4c, 0x84, 0x91, 0x2a,
	0x5a, 0x0f, 0x22, 0x8d, 0x6a, 0x2d, 0x66, 0xc8, 0x33, 0xc0, 0x20, 0x27, 0x7c, 0x3d, 0xe8, 0xfa,
	0x5e, 0x82, 0x19, 0x08, 0x70, 0xa3, 0x44, 0x1c, 0x23, 0xf1, 0x9e, 0xc7, 0x57, 0x1a, 0xd5, 0x67,
	0x22, 0x8d, 0x98, 0x8b, 0xd9, 0xc4, 0x08, 0x99, 0xe4, 0xce, 0x93, 0x58, 0xca, 0x78, 0x81, 0x51,
	0xf6, 0x34, 0x5d, 0xcd, 0x23, 0x5c, 0xa6, 0xe6, 0xc7, 0xc2, 0xeb, 0xdf, 0x43, 0x80, 0x17, 0x9b,
	0x21, 0x92, 0x98, 0xbd, 0xc1, 0xd1, 0xb3, 0xd0, 0x26, 0x57, 0x34, 0xbb, 0xe0, 0xbb, 0x9f, 0xcc,
	0x8b, 0xe4, 0x1d, 0xbf, 0x57, 0xa8, 0x4d, 0xf7, 0xb2, 0xda, 0xa0, 0x53, 0x99, 0x68, 0x64, 0xb7,
	0x00, 0x43, 0x24, 0x99, 0x9d, 0x15, 0xfd, 0x3b, 0x9d, 0xe2, 0xda, 0x45, 0x4c, 0xef, 0x3c, 0x42,
	0xeb, 0x4e, 0xe1, 0xc4, 0x20, 0x09, 0xce, 0x51, 0x07, 0xed, 0xcb, 0x19, 0xa5, 0x5f, 0x55, 0x39,
	0x0e, 0xaa, 0xcd, 0x79, 0x82, 0xd6, 0x3d, 0x2e, 0xb0, 0x22, 0xc7, 0x41, 0x94, 0xd3, 0xe1, 0x76,
	0x04, 0x4e, 0x23, 0xf0, 0x87, 0xcd, 0x08, 0xd4, 0xf7, 0x87, 0x9d, 0x36, 0xd0, 0x37, 0x91, 0xca,
	0xbe, 0x77, 0x06, 0xa7, 0xef, 0x5c, 0xf6, 0xfa, 0xce, 0xf5, 0xe0, 0xff, 0xa3, 0x77, 0xb6, 0x7d,
	0x93, 0x10, 0xe8, 0xbb, 0x61, 0x8e, 0x2d, 0x35, 0x98, 0xe3, 0xa0, 0xda, 0x9c, 0x6d, 0xdf, 0xc1,
	0x1c, 0x07, 0xed, 0xeb, 0x7b, 0x01, 0xc7, 0x9b, 0xd2, 0x46, 0x1a, 0xd5, 0x6b, 0xe1, 0x4b, 0xd1,
	0xac, 0x57, 0xee, 0xd5, 0xb3, 0x50, 0x74, 0xbf, 0x81, 0x33, 0x9f, 0x62, 0x0c, 0xed, 0x21, 0x7a,
	0x9c, 0x5d, 0x95, 0x36, 0x29, 0x1b, 0xe8, 0xd2, 0xa9, 0x53, 0x57, 0x39, 0x65, 0x02, 0x1d, 0x3b,
	0x85, 0x47, 0xfa, 0xfe, 0x5c, 0xff, 0x3e, 0x61, 0x57, 0xaa, 0x3f, 0x11, 0xf6, 0x34, 0x3b, 0x31,
	0x86, 0x8e, 0x1d, 0xb0, 0xfe, 0x44, 0xd8, 0xb3, 0x67, 0xed, 0xe9, 0x41, 0xf6, 0x7c, 0xf3, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x89, 0xbb, 0x6b, 0x20, 0x98, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitoringClient is the client API for Monitoring service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitoringClient interface {
	ListMonitors(ctx context.Context, in *ListMonitorsRequest, opts ...grpc.CallOption) (*ListMonitorsResponse, error)
	GetMonitor(ctx context.Context, in *GetMonitorRequest, opts ...grpc.CallOption) (*Monitor, error)
	CreateMonitor(ctx context.Context, in *CreateMonitorRequest, opts ...grpc.CallOption) (*Monitor, error)
	UpdateMonitor(ctx context.Context, in *UpdateMonitorRequest, opts ...grpc.CallOption) (*Monitor, error)
	DeleteMonitor(ctx context.Context, in *DeleteMonitorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListTriggers(ctx context.Context, in *ListTriggersRequest, opts ...grpc.CallOption) (*ListTriggersResponse, error)
	GetTrigger(ctx context.Context, in *GetTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	CreateTrigger(ctx context.Context, in *CreateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	UpdateTrigger(ctx context.Context, in *UpdateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	DeleteTrigger(ctx context.Context, in *DeleteTriggerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListUserNotifications(ctx context.Context, in *ListUserNotificationsRequest, opts ...grpc.CallOption) (*ListUserNotificationsResponse, error)
	GetUserNotification(ctx context.Context, in *GetUserNotificationRequest, opts ...grpc.CallOption) (*UserNotification, error)
	CreateUserNotification(ctx context.Context, in *CreateUserNotificationRequest, opts ...grpc.CallOption) (*UserNotification, error)
	UpdateUserNotification(ctx context.Context, in *UpdateUserNotificationRequest, opts ...grpc.CallOption) (*UserNotification, error)
	DeleteUserNotification(ctx context.Context, in *DeleteUserNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type monitoringClient struct {
	cc *grpc.ClientConn
}

func NewMonitoringClient(cc *grpc.ClientConn) MonitoringClient {
	return &monitoringClient{cc}
}

func (c *monitoringClient) ListMonitors(ctx context.Context, in *ListMonitorsRequest, opts ...grpc.CallOption) (*ListMonitorsResponse, error) {
	out := new(ListMonitorsResponse)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/ListMonitors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetMonitor(ctx context.Context, in *GetMonitorRequest, opts ...grpc.CallOption) (*Monitor, error) {
	out := new(Monitor)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/GetMonitor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) CreateMonitor(ctx context.Context, in *CreateMonitorRequest, opts ...grpc.CallOption) (*Monitor, error) {
	out := new(Monitor)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/CreateMonitor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) UpdateMonitor(ctx context.Context, in *UpdateMonitorRequest, opts ...grpc.CallOption) (*Monitor, error) {
	out := new(Monitor)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/UpdateMonitor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) DeleteMonitor(ctx context.Context, in *DeleteMonitorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/DeleteMonitor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) ListTriggers(ctx context.Context, in *ListTriggersRequest, opts ...grpc.CallOption) (*ListTriggersResponse, error) {
	out := new(ListTriggersResponse)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/ListTriggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetTrigger(ctx context.Context, in *GetTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/GetTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) CreateTrigger(ctx context.Context, in *CreateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/CreateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) UpdateTrigger(ctx context.Context, in *UpdateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/UpdateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) DeleteTrigger(ctx context.Context, in *DeleteTriggerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/DeleteTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) ListUserNotifications(ctx context.Context, in *ListUserNotificationsRequest, opts ...grpc.CallOption) (*ListUserNotificationsResponse, error) {
	out := new(ListUserNotificationsResponse)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/ListUserNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetUserNotification(ctx context.Context, in *GetUserNotificationRequest, opts ...grpc.CallOption) (*UserNotification, error) {
	out := new(UserNotification)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/GetUserNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) CreateUserNotification(ctx context.Context, in *CreateUserNotificationRequest, opts ...grpc.CallOption) (*UserNotification, error) {
	out := new(UserNotification)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/CreateUserNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) UpdateUserNotification(ctx context.Context, in *UpdateUserNotificationRequest, opts ...grpc.CallOption) (*UserNotification, error) {
	out := new(UserNotification)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/UpdateUserNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) DeleteUserNotification(ctx context.Context, in *DeleteUserNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/monitor.v1.Monitoring/DeleteUserNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringServer is the server API for Monitoring service.
type MonitoringServer interface {
	ListMonitors(context.Context, *ListMonitorsRequest) (*ListMonitorsResponse, error)
	GetMonitor(context.Context, *GetMonitorRequest) (*Monitor, error)
	CreateMonitor(context.Context, *CreateMonitorRequest) (*Monitor, error)
	UpdateMonitor(context.Context, *UpdateMonitorRequest) (*Monitor, error)
	DeleteMonitor(context.Context, *DeleteMonitorRequest) (*empty.Empty, error)
	ListTriggers(context.Context, *ListTriggersRequest) (*ListTriggersResponse, error)
	GetTrigger(context.Context, *GetTriggerRequest) (*Trigger, error)
	CreateTrigger(context.Context, *CreateTriggerRequest) (*Trigger, error)
	UpdateTrigger(context.Context, *UpdateTriggerRequest) (*Trigger, error)
	DeleteTrigger(context.Context, *DeleteTriggerRequest) (*empty.Empty, error)
	ListUserNotifications(context.Context, *ListUserNotificationsRequest) (*ListUserNotificationsResponse, error)
	GetUserNotification(context.Context, *GetUserNotificationRequest) (*UserNotification, error)
	CreateUserNotification(context.Context, *CreateUserNotificationRequest) (*UserNotification, error)
	UpdateUserNotification(context.Context, *UpdateUserNotificationRequest) (*UserNotification, error)
	DeleteUserNotification(context.Context, *DeleteUserNotificationRequest) (*empty.Empty, error)
}

// UnimplementedMonitoringServer can be embedded to have forward compatible implementations.
type UnimplementedMonitoringServer struct {
}

func (*UnimplementedMonitoringServer) ListMonitors(ctx context.Context, req *ListMonitorsRequest) (*ListMonitorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMonitors not implemented")
}
func (*UnimplementedMonitoringServer) GetMonitor(ctx context.Context, req *GetMonitorRequest) (*Monitor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonitor not implemented")
}
func (*UnimplementedMonitoringServer) CreateMonitor(ctx context.Context, req *CreateMonitorRequest) (*Monitor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMonitor not implemented")
}
func (*UnimplementedMonitoringServer) UpdateMonitor(ctx context.Context, req *UpdateMonitorRequest) (*Monitor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMonitor not implemented")
}
func (*UnimplementedMonitoringServer) DeleteMonitor(ctx context.Context, req *DeleteMonitorRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMonitor not implemented")
}
func (*UnimplementedMonitoringServer) ListTriggers(ctx context.Context, req *ListTriggersRequest) (*ListTriggersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTriggers not implemented")
}
func (*UnimplementedMonitoringServer) GetTrigger(ctx context.Context, req *GetTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrigger not implemented")
}
func (*UnimplementedMonitoringServer) CreateTrigger(ctx context.Context, req *CreateTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrigger not implemented")
}
func (*UnimplementedMonitoringServer) UpdateTrigger(ctx context.Context, req *UpdateTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrigger not implemented")
}
func (*UnimplementedMonitoringServer) DeleteTrigger(ctx context.Context, req *DeleteTriggerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrigger not implemented")
}
func (*UnimplementedMonitoringServer) ListUserNotifications(ctx context.Context, req *ListUserNotificationsRequest) (*ListUserNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserNotifications not implemented")
}
func (*UnimplementedMonitoringServer) GetUserNotification(ctx context.Context, req *GetUserNotificationRequest) (*UserNotification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserNotification not implemented")
}
func (*UnimplementedMonitoringServer) CreateUserNotification(ctx context.Context, req *CreateUserNotificationRequest) (*UserNotification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserNotification not implemented")
}
func (*UnimplementedMonitoringServer) UpdateUserNotification(ctx context.Context, req *UpdateUserNotificationRequest) (*UserNotification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserNotification not implemented")
}
func (*UnimplementedMonitoringServer) DeleteUserNotification(ctx context.Context, req *DeleteUserNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserNotification not implemented")
}

func RegisterMonitoringServer(s *grpc.Server, srv MonitoringServer) {
	s.RegisterService(&_Monitoring_serviceDesc, srv)
}

func _Monitoring_ListMonitors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).ListMonitors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/ListMonitors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).ListMonitors(ctx, req.(*ListMonitorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/GetMonitor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetMonitor(ctx, req.(*GetMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_CreateMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).CreateMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/CreateMonitor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).CreateMonitor(ctx, req.(*CreateMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_UpdateMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).UpdateMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/UpdateMonitor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).UpdateMonitor(ctx, req.(*UpdateMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_DeleteMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).DeleteMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/DeleteMonitor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).DeleteMonitor(ctx, req.(*DeleteMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_ListTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTriggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).ListTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/ListTriggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).ListTriggers(ctx, req.(*ListTriggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/GetTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetTrigger(ctx, req.(*GetTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_CreateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).CreateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/CreateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).CreateTrigger(ctx, req.(*CreateTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_UpdateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).UpdateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/UpdateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).UpdateTrigger(ctx, req.(*UpdateTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_DeleteTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).DeleteTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/DeleteTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).DeleteTrigger(ctx, req.(*DeleteTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_ListUserNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).ListUserNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/ListUserNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).ListUserNotifications(ctx, req.(*ListUserNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetUserNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetUserNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/GetUserNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetUserNotification(ctx, req.(*GetUserNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_CreateUserNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).CreateUserNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/CreateUserNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).CreateUserNotification(ctx, req.(*CreateUserNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_UpdateUserNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).UpdateUserNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/UpdateUserNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).UpdateUserNotification(ctx, req.(*UpdateUserNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_DeleteUserNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).DeleteUserNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.Monitoring/DeleteUserNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).DeleteUserNotification(ctx, req.(*DeleteUserNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitoring_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.v1.Monitoring",
	HandlerType: (*MonitoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMonitors",
			Handler:    _Monitoring_ListMonitors_Handler,
		},
		{
			MethodName: "GetMonitor",
			Handler:    _Monitoring_GetMonitor_Handler,
		},
		{
			MethodName: "CreateMonitor",
			Handler:    _Monitoring_CreateMonitor_Handler,
		},
		{
			MethodName: "UpdateMonitor",
			Handler:    _Monitoring_UpdateMonitor_Handler,
		},
		{
			MethodName: "DeleteMonitor",
			Handler:    _Monitoring_DeleteMonitor_Handler,
		},
		{
			MethodName: "ListTriggers",
			Handler:    _Monitoring_ListTriggers_Handler,
		},
		{
			MethodName: "GetTrigger",
			Handler:    _Monitoring_GetTrigger_Handler,
		},
		{
			MethodName: "CreateTrigger",
			Handler:    _Monitoring_CreateTrigger_Handler,
		},
		{
			MethodName: "UpdateTrigger",
			Handler:    _Monitoring_UpdateTrigger_Handler,
		},
		{
			MethodName: "DeleteTrigger",
			Handler:    _Monitoring_DeleteTrigger_Handler,
		},
		{
			MethodName: "ListUserNotifications",
			Handler:    _Monitoring_ListUserNotifications_Handler,
		},
		{
			MethodName: "GetUserNotification",
			Handler:    _Monitoring_GetUserNotification_Handler,
		},
		{
			MethodName: "CreateUserNotification",
			Handler:    _Monitoring_CreateUserNotification_Handler,
		},
		{
			MethodName: "UpdateUserNotification",
			Handler:    _Monitoring_UpdateUserNotification_Handler,
		},
		{
			MethodName: "DeleteUserNotification",
			Handler:    _Monitoring_DeleteUserNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/monitor/v1/service.proto",
}
