// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: auth_server.proto

package pbas

import (
	context "context"
	base "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/base"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Auth_InitAuthCenter_FullMethodName             = "/pbas.Auth/InitAuthCenter"
	Auth_GetUserInfo_FullMethodName                = "/pbas.Auth/GetUserInfo"
	Auth_ListUserSpace_FullMethodName              = "/pbas.Auth/ListUserSpace"
	Auth_QuerySpaceByAppID_FullMethodName          = "/pbas.Auth/QuerySpaceByAppID"
	Auth_PullResource_FullMethodName               = "/pbas.Auth/PullResource"
	Auth_CheckPermission_FullMethodName            = "/pbas.Auth/CheckPermission"
	Auth_AuthorizeBatch_FullMethodName             = "/pbas.Auth/AuthorizeBatch"
	Auth_GetPermissionToApply_FullMethodName       = "/pbas.Auth/GetPermissionToApply"
	Auth_QuerySpace_FullMethodName                 = "/pbas.Auth/QuerySpace"
	Auth_GetAuthLoginConf_FullMethodName           = "/pbas.Auth/GetAuthLoginConf"
	Auth_GrantResourceCreatorAction_FullMethodName = "/pbas.Auth/GrantResourceCreatorAction"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// init auth center's auth model.
	InitAuthCenter(ctx context.Context, in *InitAuthCenterReq, opts ...grpc.CallOption) (*InitAuthCenterResp, error)
	// 获取用户鉴权信息
	GetUserInfo(ctx context.Context, in *UserCredentialReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	ListUserSpace(ctx context.Context, in *ListUserSpaceReq, opts ...grpc.CallOption) (*ListUserSpaceResp, error)
	// 通过 AppID 查询 Space 信息
	QuerySpaceByAppID(ctx context.Context, in *QuerySpaceByAppIDReq, opts ...grpc.CallOption) (*Space, error)
	// iam pull resource callback.
	PullResource(ctx context.Context, in *PullResourceReq, opts ...grpc.CallOption) (*structpb.Struct, error)
	// 权限中心权限检测
	CheckPermission(ctx context.Context, in *CheckPermissionReq, opts ...grpc.CallOption) (*CheckPermissionResp, error)
	// authorize resource batch.
	AuthorizeBatch(ctx context.Context, in *AuthorizeBatchReq, opts ...grpc.CallOption) (*AuthorizeBatchResp, error)
	// get iam permission to apply.
	GetPermissionToApply(ctx context.Context, in *GetPermissionToApplyReq, opts ...grpc.CallOption) (*GetPermissionToApplyResp, error)
	QuerySpace(ctx context.Context, in *QuerySpaceReq, opts ...grpc.CallOption) (*QuerySpaceResp, error)
	// auth login conf
	GetAuthLoginConf(ctx context.Context, in *GetAuthLoginConfReq, opts ...grpc.CallOption) (*GetAuthLoginConfResp, error)
	GrantResourceCreatorAction(ctx context.Context, in *GrantResourceCreatorActionReq, opts ...grpc.CallOption) (*base.EmptyResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) InitAuthCenter(ctx context.Context, in *InitAuthCenterReq, opts ...grpc.CallOption) (*InitAuthCenterResp, error) {
	out := new(InitAuthCenterResp)
	err := c.cc.Invoke(ctx, Auth_InitAuthCenter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserInfo(ctx context.Context, in *UserCredentialReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, Auth_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListUserSpace(ctx context.Context, in *ListUserSpaceReq, opts ...grpc.CallOption) (*ListUserSpaceResp, error) {
	out := new(ListUserSpaceResp)
	err := c.cc.Invoke(ctx, Auth_ListUserSpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) QuerySpaceByAppID(ctx context.Context, in *QuerySpaceByAppIDReq, opts ...grpc.CallOption) (*Space, error) {
	out := new(Space)
	err := c.cc.Invoke(ctx, Auth_QuerySpaceByAppID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) PullResource(ctx context.Context, in *PullResourceReq, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, Auth_PullResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckPermission(ctx context.Context, in *CheckPermissionReq, opts ...grpc.CallOption) (*CheckPermissionResp, error) {
	out := new(CheckPermissionResp)
	err := c.cc.Invoke(ctx, Auth_CheckPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthorizeBatch(ctx context.Context, in *AuthorizeBatchReq, opts ...grpc.CallOption) (*AuthorizeBatchResp, error) {
	out := new(AuthorizeBatchResp)
	err := c.cc.Invoke(ctx, Auth_AuthorizeBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetPermissionToApply(ctx context.Context, in *GetPermissionToApplyReq, opts ...grpc.CallOption) (*GetPermissionToApplyResp, error) {
	out := new(GetPermissionToApplyResp)
	err := c.cc.Invoke(ctx, Auth_GetPermissionToApply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) QuerySpace(ctx context.Context, in *QuerySpaceReq, opts ...grpc.CallOption) (*QuerySpaceResp, error) {
	out := new(QuerySpaceResp)
	err := c.cc.Invoke(ctx, Auth_QuerySpace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAuthLoginConf(ctx context.Context, in *GetAuthLoginConfReq, opts ...grpc.CallOption) (*GetAuthLoginConfResp, error) {
	out := new(GetAuthLoginConfResp)
	err := c.cc.Invoke(ctx, Auth_GetAuthLoginConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GrantResourceCreatorAction(ctx context.Context, in *GrantResourceCreatorActionReq, opts ...grpc.CallOption) (*base.EmptyResp, error) {
	out := new(base.EmptyResp)
	err := c.cc.Invoke(ctx, Auth_GrantResourceCreatorAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// init auth center's auth model.
	InitAuthCenter(context.Context, *InitAuthCenterReq) (*InitAuthCenterResp, error)
	// 获取用户鉴权信息
	GetUserInfo(context.Context, *UserCredentialReq) (*UserInfoResp, error)
	ListUserSpace(context.Context, *ListUserSpaceReq) (*ListUserSpaceResp, error)
	// 通过 AppID 查询 Space 信息
	QuerySpaceByAppID(context.Context, *QuerySpaceByAppIDReq) (*Space, error)
	// iam pull resource callback.
	PullResource(context.Context, *PullResourceReq) (*structpb.Struct, error)
	// 权限中心权限检测
	CheckPermission(context.Context, *CheckPermissionReq) (*CheckPermissionResp, error)
	// authorize resource batch.
	AuthorizeBatch(context.Context, *AuthorizeBatchReq) (*AuthorizeBatchResp, error)
	// get iam permission to apply.
	GetPermissionToApply(context.Context, *GetPermissionToApplyReq) (*GetPermissionToApplyResp, error)
	QuerySpace(context.Context, *QuerySpaceReq) (*QuerySpaceResp, error)
	// auth login conf
	GetAuthLoginConf(context.Context, *GetAuthLoginConfReq) (*GetAuthLoginConfResp, error)
	GrantResourceCreatorAction(context.Context, *GrantResourceCreatorActionReq) (*base.EmptyResp, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) InitAuthCenter(context.Context, *InitAuthCenterReq) (*InitAuthCenterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitAuthCenter not implemented")
}
func (UnimplementedAuthServer) GetUserInfo(context.Context, *UserCredentialReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedAuthServer) ListUserSpace(context.Context, *ListUserSpaceReq) (*ListUserSpaceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserSpace not implemented")
}
func (UnimplementedAuthServer) QuerySpaceByAppID(context.Context, *QuerySpaceByAppIDReq) (*Space, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySpaceByAppID not implemented")
}
func (UnimplementedAuthServer) PullResource(context.Context, *PullResourceReq) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullResource not implemented")
}
func (UnimplementedAuthServer) CheckPermission(context.Context, *CheckPermissionReq) (*CheckPermissionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedAuthServer) AuthorizeBatch(context.Context, *AuthorizeBatchReq) (*AuthorizeBatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeBatch not implemented")
}
func (UnimplementedAuthServer) GetPermissionToApply(context.Context, *GetPermissionToApplyReq) (*GetPermissionToApplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionToApply not implemented")
}
func (UnimplementedAuthServer) QuerySpace(context.Context, *QuerySpaceReq) (*QuerySpaceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySpace not implemented")
}
func (UnimplementedAuthServer) GetAuthLoginConf(context.Context, *GetAuthLoginConfReq) (*GetAuthLoginConfResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthLoginConf not implemented")
}
func (UnimplementedAuthServer) GrantResourceCreatorAction(context.Context, *GrantResourceCreatorActionReq) (*base.EmptyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantResourceCreatorAction not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_InitAuthCenter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitAuthCenterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).InitAuthCenter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_InitAuthCenter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).InitAuthCenter(ctx, req.(*InitAuthCenterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserInfo(ctx, req.(*UserCredentialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListUserSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserSpaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListUserSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ListUserSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListUserSpace(ctx, req.(*ListUserSpaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_QuerySpaceByAppID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpaceByAppIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).QuerySpaceByAppID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_QuerySpaceByAppID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).QuerySpaceByAppID(ctx, req.(*QuerySpaceByAppIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_PullResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullResourceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).PullResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_PullResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).PullResource(ctx, req.(*PullResourceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckPermission(ctx, req.(*CheckPermissionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthorizeBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthorizeBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AuthorizeBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthorizeBatch(ctx, req.(*AuthorizeBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetPermissionToApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionToApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetPermissionToApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetPermissionToApply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetPermissionToApply(ctx, req.(*GetPermissionToApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_QuerySpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).QuerySpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_QuerySpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).QuerySpace(ctx, req.(*QuerySpaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAuthLoginConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthLoginConfReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAuthLoginConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetAuthLoginConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAuthLoginConf(ctx, req.(*GetAuthLoginConfReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GrantResourceCreatorAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantResourceCreatorActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GrantResourceCreatorAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GrantResourceCreatorAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GrantResourceCreatorAction(ctx, req.(*GrantResourceCreatorActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbas.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitAuthCenter",
			Handler:    _Auth_InitAuthCenter_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _Auth_GetUserInfo_Handler,
		},
		{
			MethodName: "ListUserSpace",
			Handler:    _Auth_ListUserSpace_Handler,
		},
		{
			MethodName: "QuerySpaceByAppID",
			Handler:    _Auth_QuerySpaceByAppID_Handler,
		},
		{
			MethodName: "PullResource",
			Handler:    _Auth_PullResource_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _Auth_CheckPermission_Handler,
		},
		{
			MethodName: "AuthorizeBatch",
			Handler:    _Auth_AuthorizeBatch_Handler,
		},
		{
			MethodName: "GetPermissionToApply",
			Handler:    _Auth_GetPermissionToApply_Handler,
		},
		{
			MethodName: "QuerySpace",
			Handler:    _Auth_QuerySpace_Handler,
		},
		{
			MethodName: "GetAuthLoginConf",
			Handler:    _Auth_GetAuthLoginConf_Handler,
		},
		{
			MethodName: "GrantResourceCreatorAction",
			Handler:    _Auth_GrantResourceCreatorAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_server.proto",
}
