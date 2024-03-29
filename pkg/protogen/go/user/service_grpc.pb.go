// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/proto/user/service.proto

package user

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

const (
	UserService_UserRegistration_FullMethodName     = "/user.UserService/UserRegistration"
	UserService_Login_FullMethodName                = "/user.UserService/Login"
	UserService_Logout_FullMethodName               = "/user.UserService/Logout"
	UserService_RenewToken_FullMethodName           = "/user.UserService/RenewToken"
	UserService_PasswordReset_FullMethodName        = "/user.UserService/PasswordReset"
	UserService_SendVerificationCode_FullMethodName = "/user.UserService/SendVerificationCode"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutResquest, opts ...grpc.CallOption) (*LogoutResponse, error)
	RenewToken(ctx context.Context, in *RenewTokenRequest, opts ...grpc.CallOption) (*RenewTokenResponse, error)
	PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*PasswordResetResponse, error)
	SendVerificationCode(ctx context.Context, in *CodeVerficationRequest, opts ...grpc.CallOption) (*CodeVerficationResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, UserService_UserRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *LogoutResquest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, UserService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RenewToken(ctx context.Context, in *RenewTokenRequest, opts ...grpc.CallOption) (*RenewTokenResponse, error) {
	out := new(RenewTokenResponse)
	err := c.cc.Invoke(ctx, UserService_RenewToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*PasswordResetResponse, error) {
	out := new(PasswordResetResponse)
	err := c.cc.Invoke(ctx, UserService_PasswordReset_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendVerificationCode(ctx context.Context, in *CodeVerficationRequest, opts ...grpc.CallOption) (*CodeVerficationResponse, error) {
	out := new(CodeVerficationResponse)
	err := c.cc.Invoke(ctx, UserService_SendVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutResquest) (*LogoutResponse, error)
	RenewToken(context.Context, *RenewTokenRequest) (*RenewTokenResponse, error)
	PasswordReset(context.Context, *PasswordResetRequest) (*PasswordResetResponse, error)
	SendVerificationCode(context.Context, *CodeVerficationRequest) (*CodeVerficationResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegistration not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) Logout(context.Context, *LogoutResquest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServiceServer) RenewToken(context.Context, *RenewTokenRequest) (*RenewTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewToken not implemented")
}
func (UnimplementedUserServiceServer) PasswordReset(context.Context, *PasswordResetRequest) (*PasswordResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordReset not implemented")
}
func (UnimplementedUserServiceServer) SendVerificationCode(context.Context, *CodeVerficationRequest) (*CodeVerficationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerificationCode not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegistration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*LogoutResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RenewToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RenewToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RenewToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RenewToken(ctx, req.(*RenewTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PasswordReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PasswordReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_PasswordReset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PasswordReset(ctx, req.(*PasswordResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeVerficationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SendVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendVerificationCode(ctx, req.(*CodeVerficationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegistration",
			Handler:    _UserService_UserRegistration_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "RenewToken",
			Handler:    _UserService_RenewToken_Handler,
		},
		{
			MethodName: "PasswordReset",
			Handler:    _UserService_PasswordReset_Handler,
		},
		{
			MethodName: "SendVerificationCode",
			Handler:    _UserService_SendVerificationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/user/service.proto",
}
