// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: service_bank_of_asia.proto

package pb

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

// BankOfAsiaClient is the client API for BankOfAsia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankOfAsiaClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type bankOfAsiaClient struct {
	cc grpc.ClientConnInterface
}

func NewBankOfAsiaClient(cc grpc.ClientConnInterface) BankOfAsiaClient {
	return &bankOfAsiaClient{cc}
}

func (c *bankOfAsiaClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.BankOfAsia/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankOfAsiaClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.BankOfAsia/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankOfAsiaServer is the server API for BankOfAsia service.
// All implementations must embed UnimplementedBankOfAsiaServer
// for forward compatibility
type BankOfAsiaServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedBankOfAsiaServer()
}

// UnimplementedBankOfAsiaServer must be embedded to have forward compatible implementations.
type UnimplementedBankOfAsiaServer struct {
}

func (UnimplementedBankOfAsiaServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedBankOfAsiaServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedBankOfAsiaServer) mustEmbedUnimplementedBankOfAsiaServer() {}

// UnsafeBankOfAsiaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankOfAsiaServer will
// result in compilation errors.
type UnsafeBankOfAsiaServer interface {
	mustEmbedUnimplementedBankOfAsiaServer()
}

func RegisterBankOfAsiaServer(s grpc.ServiceRegistrar, srv BankOfAsiaServer) {
	s.RegisterService(&BankOfAsia_ServiceDesc, srv)
}

func _BankOfAsia_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankOfAsiaServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BankOfAsia/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankOfAsiaServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankOfAsia_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankOfAsiaServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BankOfAsia/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankOfAsiaServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankOfAsia_ServiceDesc is the grpc.ServiceDesc for BankOfAsia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankOfAsia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BankOfAsia",
	HandlerType: (*BankOfAsiaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _BankOfAsia_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _BankOfAsia_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_bank_of_asia.proto",
}
