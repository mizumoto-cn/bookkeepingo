// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: v1/interface.proto

package interfacev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SiteAdminServiceClient is the client API for SiteAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiteAdminServiceClient interface {
	Login(ctx context.Context, in *LoginRequire, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequire, opts ...grpc.CallOption) (*LogoutResponse, error)
	ListAccount(ctx context.Context, in *ListAccountRequire, opts ...grpc.CallOption) (*ListAccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequire, opts ...grpc.CallOption) (*GetAccountResponse, error)
}

type siteAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSiteAdminServiceClient(cc grpc.ClientConnInterface) SiteAdminServiceClient {
	return &siteAdminServiceClient{cc}
}

func (c *siteAdminServiceClient) Login(ctx context.Context, in *LoginRequire, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteAdminServiceClient) Logout(ctx context.Context, in *LogoutRequire, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteAdminServiceClient) ListAccount(ctx context.Context, in *ListAccountRequire, opts ...grpc.CallOption) (*ListAccountResponse, error) {
	out := new(ListAccountResponse)
	err := c.cc.Invoke(ctx, "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/ListAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteAdminServiceClient) GetAccount(ctx context.Context, in *GetAccountRequire, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiteAdminServiceServer is the server API for SiteAdminService service.
// All implementations must embed UnimplementedSiteAdminServiceServer
// for forward compatibility
type SiteAdminServiceServer interface {
	Login(context.Context, *LoginRequire) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequire) (*LogoutResponse, error)
	ListAccount(context.Context, *ListAccountRequire) (*ListAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequire) (*GetAccountResponse, error)
	mustEmbedUnimplementedSiteAdminServiceServer()
}

// UnimplementedSiteAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSiteAdminServiceServer struct {
}

func (UnimplementedSiteAdminServiceServer) Login(context.Context, *LoginRequire) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSiteAdminServiceServer) Logout(context.Context, *LogoutRequire) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSiteAdminServiceServer) ListAccount(context.Context, *ListAccountRequire) (*ListAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
}
func (UnimplementedSiteAdminServiceServer) GetAccount(context.Context, *GetAccountRequire) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedSiteAdminServiceServer) mustEmbedUnimplementedSiteAdminServiceServer() {}

// UnsafeSiteAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiteAdminServiceServer will
// result in compilation errors.
type UnsafeSiteAdminServiceServer interface {
	mustEmbedUnimplementedSiteAdminServiceServer()
}

func RegisterSiteAdminServiceServer(s grpc.ServiceRegistrar, srv SiteAdminServiceServer) {
	s.RegisterService(&SiteAdminService_ServiceDesc, srv)
}

func _SiteAdminService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequire)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteAdminServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteAdminServiceServer).Login(ctx, req.(*LoginRequire))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteAdminService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequire)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteAdminServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteAdminServiceServer).Logout(ctx, req.(*LogoutRequire))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteAdminService_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountRequire)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteAdminServiceServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/ListAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteAdminServiceServer).ListAccount(ctx, req.(*ListAccountRequire))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteAdminService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequire)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteAdminServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteAdminServiceServer).GetAccount(ctx, req.(*GetAccountRequire))
	}
	return interceptor(ctx, in, info, handler)
}

// SiteAdminService_ServiceDesc is the grpc.ServiceDesc for SiteAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SiteAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookkeepingo.mizumoto.tech.site.interface.v1.SiteAdminService",
	HandlerType: (*SiteAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SiteAdminService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SiteAdminService_Logout_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _SiteAdminService_ListAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _SiteAdminService_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/interface.proto",
}
