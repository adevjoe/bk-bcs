// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: cache_service.proto

package pbcs

import (
	context "context"
	base "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/base"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Cache_GetAppID_FullMethodName                 = "/pbcs.Cache/GetAppID"
	Cache_GetAppMeta_FullMethodName               = "/pbcs.Cache/GetAppMeta"
	Cache_ListApps_FullMethodName                 = "/pbcs.Cache/ListApps"
	Cache_GetReleasedCI_FullMethodName            = "/pbcs.Cache/GetReleasedCI"
	Cache_GetReleasedHook_FullMethodName          = "/pbcs.Cache/GetReleasedHook"
	Cache_ListAppReleasedGroups_FullMethodName    = "/pbcs.Cache/ListAppReleasedGroups"
	Cache_GetCurrentCursorReminder_FullMethodName = "/pbcs.Cache/GetCurrentCursorReminder"
	Cache_ListEventsMeta_FullMethodName           = "/pbcs.Cache/ListEventsMeta"
	Cache_GetCredential_FullMethodName            = "/pbcs.Cache/GetCredential"
	Cache_BenchAppMeta_FullMethodName             = "/pbcs.Cache/BenchAppMeta"
	Cache_BenchReleasedCI_FullMethodName          = "/pbcs.Cache/BenchReleasedCI"
	Cache_GetReleasedKv_FullMethodName            = "/pbcs.Cache/GetReleasedKv"
	Cache_GetReleasedKvValue_FullMethodName       = "/pbcs.Cache/GetReleasedKvValue"
)

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheClient interface {
	GetAppID(ctx context.Context, in *GetAppIDReq, opts ...grpc.CallOption) (*GetAppIDResp, error)
	GetAppMeta(ctx context.Context, in *GetAppMetaReq, opts ...grpc.CallOption) (*JsonRawResp, error)
	ListApps(ctx context.Context, in *ListAppsReq, opts ...grpc.CallOption) (*ListAppsResp, error)
	GetReleasedCI(ctx context.Context, in *GetReleasedCIReq, opts ...grpc.CallOption) (*JsonRawResp, error)
	GetReleasedHook(ctx context.Context, in *GetReleasedHookReq, opts ...grpc.CallOption) (*JsonRawResp, error)
	ListAppReleasedGroups(ctx context.Context, in *ListAppReleasedGroupsReq, opts ...grpc.CallOption) (*JsonRawResp, error)
	GetCurrentCursorReminder(ctx context.Context, in *base.EmptyReq, opts ...grpc.CallOption) (*CurrentCursorReminderResp, error)
	ListEventsMeta(ctx context.Context, in *ListEventsReq, opts ...grpc.CallOption) (*ListEventsResp, error)
	GetCredential(ctx context.Context, in *GetCredentialReq, opts ...grpc.CallOption) (*JsonRawResp, error)
	// only stress test use.
	BenchAppMeta(ctx context.Context, in *BenchAppMetaReq, opts ...grpc.CallOption) (*BenchAppMetaResp, error)
	BenchReleasedCI(ctx context.Context, in *BenchReleasedCIReq, opts ...grpc.CallOption) (*BenchReleasedCIResp, error)
	GetReleasedKv(ctx context.Context, in *GetReleasedKvReq, opts ...grpc.CallOption) (*JsonRawResp, error)
	GetReleasedKvValue(ctx context.Context, in *GetReleasedKvValueReq, opts ...grpc.CallOption) (*JsonRawResp, error)
}

type cacheClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheClient(cc grpc.ClientConnInterface) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) GetAppID(ctx context.Context, in *GetAppIDReq, opts ...grpc.CallOption) (*GetAppIDResp, error) {
	out := new(GetAppIDResp)
	err := c.cc.Invoke(ctx, Cache_GetAppID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetAppMeta(ctx context.Context, in *GetAppMetaReq, opts ...grpc.CallOption) (*JsonRawResp, error) {
	out := new(JsonRawResp)
	err := c.cc.Invoke(ctx, Cache_GetAppMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) ListApps(ctx context.Context, in *ListAppsReq, opts ...grpc.CallOption) (*ListAppsResp, error) {
	out := new(ListAppsResp)
	err := c.cc.Invoke(ctx, Cache_ListApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetReleasedCI(ctx context.Context, in *GetReleasedCIReq, opts ...grpc.CallOption) (*JsonRawResp, error) {
	out := new(JsonRawResp)
	err := c.cc.Invoke(ctx, Cache_GetReleasedCI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetReleasedHook(ctx context.Context, in *GetReleasedHookReq, opts ...grpc.CallOption) (*JsonRawResp, error) {
	out := new(JsonRawResp)
	err := c.cc.Invoke(ctx, Cache_GetReleasedHook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) ListAppReleasedGroups(ctx context.Context, in *ListAppReleasedGroupsReq, opts ...grpc.CallOption) (*JsonRawResp, error) {
	out := new(JsonRawResp)
	err := c.cc.Invoke(ctx, Cache_ListAppReleasedGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetCurrentCursorReminder(ctx context.Context, in *base.EmptyReq, opts ...grpc.CallOption) (*CurrentCursorReminderResp, error) {
	out := new(CurrentCursorReminderResp)
	err := c.cc.Invoke(ctx, Cache_GetCurrentCursorReminder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) ListEventsMeta(ctx context.Context, in *ListEventsReq, opts ...grpc.CallOption) (*ListEventsResp, error) {
	out := new(ListEventsResp)
	err := c.cc.Invoke(ctx, Cache_ListEventsMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetCredential(ctx context.Context, in *GetCredentialReq, opts ...grpc.CallOption) (*JsonRawResp, error) {
	out := new(JsonRawResp)
	err := c.cc.Invoke(ctx, Cache_GetCredential_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) BenchAppMeta(ctx context.Context, in *BenchAppMetaReq, opts ...grpc.CallOption) (*BenchAppMetaResp, error) {
	out := new(BenchAppMetaResp)
	err := c.cc.Invoke(ctx, Cache_BenchAppMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) BenchReleasedCI(ctx context.Context, in *BenchReleasedCIReq, opts ...grpc.CallOption) (*BenchReleasedCIResp, error) {
	out := new(BenchReleasedCIResp)
	err := c.cc.Invoke(ctx, Cache_BenchReleasedCI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetReleasedKv(ctx context.Context, in *GetReleasedKvReq, opts ...grpc.CallOption) (*JsonRawResp, error) {
	out := new(JsonRawResp)
	err := c.cc.Invoke(ctx, Cache_GetReleasedKv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) GetReleasedKvValue(ctx context.Context, in *GetReleasedKvValueReq, opts ...grpc.CallOption) (*JsonRawResp, error) {
	out := new(JsonRawResp)
	err := c.cc.Invoke(ctx, Cache_GetReleasedKvValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
// All implementations should embed UnimplementedCacheServer
// for forward compatibility
type CacheServer interface {
	GetAppID(context.Context, *GetAppIDReq) (*GetAppIDResp, error)
	GetAppMeta(context.Context, *GetAppMetaReq) (*JsonRawResp, error)
	ListApps(context.Context, *ListAppsReq) (*ListAppsResp, error)
	GetReleasedCI(context.Context, *GetReleasedCIReq) (*JsonRawResp, error)
	GetReleasedHook(context.Context, *GetReleasedHookReq) (*JsonRawResp, error)
	ListAppReleasedGroups(context.Context, *ListAppReleasedGroupsReq) (*JsonRawResp, error)
	GetCurrentCursorReminder(context.Context, *base.EmptyReq) (*CurrentCursorReminderResp, error)
	ListEventsMeta(context.Context, *ListEventsReq) (*ListEventsResp, error)
	GetCredential(context.Context, *GetCredentialReq) (*JsonRawResp, error)
	// only stress test use.
	BenchAppMeta(context.Context, *BenchAppMetaReq) (*BenchAppMetaResp, error)
	BenchReleasedCI(context.Context, *BenchReleasedCIReq) (*BenchReleasedCIResp, error)
	GetReleasedKv(context.Context, *GetReleasedKvReq) (*JsonRawResp, error)
	GetReleasedKvValue(context.Context, *GetReleasedKvValueReq) (*JsonRawResp, error)
}

// UnimplementedCacheServer should be embedded to have forward compatible implementations.
type UnimplementedCacheServer struct {
}

func (UnimplementedCacheServer) GetAppID(context.Context, *GetAppIDReq) (*GetAppIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppID not implemented")
}
func (UnimplementedCacheServer) GetAppMeta(context.Context, *GetAppMetaReq) (*JsonRawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppMeta not implemented")
}
func (UnimplementedCacheServer) ListApps(context.Context, *ListAppsReq) (*ListAppsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApps not implemented")
}
func (UnimplementedCacheServer) GetReleasedCI(context.Context, *GetReleasedCIReq) (*JsonRawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReleasedCI not implemented")
}
func (UnimplementedCacheServer) GetReleasedHook(context.Context, *GetReleasedHookReq) (*JsonRawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReleasedHook not implemented")
}
func (UnimplementedCacheServer) ListAppReleasedGroups(context.Context, *ListAppReleasedGroupsReq) (*JsonRawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppReleasedGroups not implemented")
}
func (UnimplementedCacheServer) GetCurrentCursorReminder(context.Context, *base.EmptyReq) (*CurrentCursorReminderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentCursorReminder not implemented")
}
func (UnimplementedCacheServer) ListEventsMeta(context.Context, *ListEventsReq) (*ListEventsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventsMeta not implemented")
}
func (UnimplementedCacheServer) GetCredential(context.Context, *GetCredentialReq) (*JsonRawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredential not implemented")
}
func (UnimplementedCacheServer) BenchAppMeta(context.Context, *BenchAppMetaReq) (*BenchAppMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BenchAppMeta not implemented")
}
func (UnimplementedCacheServer) BenchReleasedCI(context.Context, *BenchReleasedCIReq) (*BenchReleasedCIResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BenchReleasedCI not implemented")
}
func (UnimplementedCacheServer) GetReleasedKv(context.Context, *GetReleasedKvReq) (*JsonRawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReleasedKv not implemented")
}
func (UnimplementedCacheServer) GetReleasedKvValue(context.Context, *GetReleasedKvValueReq) (*JsonRawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReleasedKvValue not implemented")
}

// UnsafeCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServer will
// result in compilation errors.
type UnsafeCacheServer interface {
	mustEmbedUnimplementedCacheServer()
}

func RegisterCacheServer(s grpc.ServiceRegistrar, srv CacheServer) {
	s.RegisterService(&Cache_ServiceDesc, srv)
}

func _Cache_GetAppID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetAppID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetAppID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetAppID(ctx, req.(*GetAppIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetAppMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetAppMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetAppMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetAppMeta(ctx, req.(*GetAppMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_ListApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).ListApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_ListApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).ListApps(ctx, req.(*ListAppsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetReleasedCI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleasedCIReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetReleasedCI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetReleasedCI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetReleasedCI(ctx, req.(*GetReleasedCIReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetReleasedHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleasedHookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetReleasedHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetReleasedHook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetReleasedHook(ctx, req.(*GetReleasedHookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_ListAppReleasedGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppReleasedGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).ListAppReleasedGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_ListAppReleasedGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).ListAppReleasedGroups(ctx, req.(*ListAppReleasedGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetCurrentCursorReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(base.EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetCurrentCursorReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetCurrentCursorReminder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetCurrentCursorReminder(ctx, req.(*base.EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_ListEventsMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).ListEventsMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_ListEventsMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).ListEventsMeta(ctx, req.(*ListEventsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCredentialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetCredential(ctx, req.(*GetCredentialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_BenchAppMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BenchAppMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).BenchAppMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_BenchAppMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).BenchAppMeta(ctx, req.(*BenchAppMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_BenchReleasedCI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BenchReleasedCIReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).BenchReleasedCI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_BenchReleasedCI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).BenchReleasedCI(ctx, req.(*BenchReleasedCIReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetReleasedKv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleasedKvReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetReleasedKv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetReleasedKv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetReleasedKv(ctx, req.(*GetReleasedKvReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_GetReleasedKvValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleasedKvValueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).GetReleasedKvValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cache_GetReleasedKvValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).GetReleasedKvValue(ctx, req.(*GetReleasedKvValueReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cache_ServiceDesc is the grpc.ServiceDesc for Cache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbcs.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppID",
			Handler:    _Cache_GetAppID_Handler,
		},
		{
			MethodName: "GetAppMeta",
			Handler:    _Cache_GetAppMeta_Handler,
		},
		{
			MethodName: "ListApps",
			Handler:    _Cache_ListApps_Handler,
		},
		{
			MethodName: "GetReleasedCI",
			Handler:    _Cache_GetReleasedCI_Handler,
		},
		{
			MethodName: "GetReleasedHook",
			Handler:    _Cache_GetReleasedHook_Handler,
		},
		{
			MethodName: "ListAppReleasedGroups",
			Handler:    _Cache_ListAppReleasedGroups_Handler,
		},
		{
			MethodName: "GetCurrentCursorReminder",
			Handler:    _Cache_GetCurrentCursorReminder_Handler,
		},
		{
			MethodName: "ListEventsMeta",
			Handler:    _Cache_ListEventsMeta_Handler,
		},
		{
			MethodName: "GetCredential",
			Handler:    _Cache_GetCredential_Handler,
		},
		{
			MethodName: "BenchAppMeta",
			Handler:    _Cache_BenchAppMeta_Handler,
		},
		{
			MethodName: "BenchReleasedCI",
			Handler:    _Cache_BenchReleasedCI_Handler,
		},
		{
			MethodName: "GetReleasedKv",
			Handler:    _Cache_GetReleasedKv_Handler,
		},
		{
			MethodName: "GetReleasedKvValue",
			Handler:    _Cache_GetReleasedKvValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache_service.proto",
}
