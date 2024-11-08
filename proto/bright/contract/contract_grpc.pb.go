// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: bright/contract/contract.proto

package contract

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	GetContractCode(ctx context.Context, in *GetContractCodeRequest, opts ...grpc.CallOption) (*GetContractCodeResponse, error)
	CompileContractCode(ctx context.Context, in *CompileContractCodeRequest, opts ...grpc.CallOption) (*CompileContractCodeResponse, error)
	CreateContractWithAccount(ctx context.Context, in *CreateContractWithAccountRequest, opts ...grpc.CallOption) (*CreateContractWithAccountResponse, error)
	GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error)
	DeleteContract(ctx context.Context, in *DeleteContractRequest, opts ...grpc.CallOption) (*DeleteContractResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetContractCode(ctx context.Context, in *GetContractCodeRequest, opts ...grpc.CallOption) (*GetContractCodeResponse, error) {
	out := new(GetContractCodeResponse)
	err := c.cc.Invoke(ctx, "/bright.contract.Manager/GetContractCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CompileContractCode(ctx context.Context, in *CompileContractCodeRequest, opts ...grpc.CallOption) (*CompileContractCodeResponse, error) {
	out := new(CompileContractCodeResponse)
	err := c.cc.Invoke(ctx, "/bright.contract.Manager/CompileContractCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateContractWithAccount(ctx context.Context, in *CreateContractWithAccountRequest, opts ...grpc.CallOption) (*CreateContractWithAccountResponse, error) {
	out := new(CreateContractWithAccountResponse)
	err := c.cc.Invoke(ctx, "/bright.contract.Manager/CreateContractWithAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error) {
	out := new(GetContractResponse)
	err := c.cc.Invoke(ctx, "/bright.contract.Manager/GetContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteContract(ctx context.Context, in *DeleteContractRequest, opts ...grpc.CallOption) (*DeleteContractResponse, error) {
	out := new(DeleteContractResponse)
	err := c.cc.Invoke(ctx, "/bright.contract.Manager/DeleteContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	GetContractCode(context.Context, *GetContractCodeRequest) (*GetContractCodeResponse, error)
	CompileContractCode(context.Context, *CompileContractCodeRequest) (*CompileContractCodeResponse, error)
	CreateContractWithAccount(context.Context, *CreateContractWithAccountRequest) (*CreateContractWithAccountResponse, error)
	GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error)
	DeleteContract(context.Context, *DeleteContractRequest) (*DeleteContractResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) GetContractCode(context.Context, *GetContractCodeRequest) (*GetContractCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContractCode not implemented")
}
func (UnimplementedManagerServer) CompileContractCode(context.Context, *CompileContractCodeRequest) (*CompileContractCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompileContractCode not implemented")
}
func (UnimplementedManagerServer) CreateContractWithAccount(context.Context, *CreateContractWithAccountRequest) (*CreateContractWithAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContractWithAccount not implemented")
}
func (UnimplementedManagerServer) GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContract not implemented")
}
func (UnimplementedManagerServer) DeleteContract(context.Context, *DeleteContractRequest) (*DeleteContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContract not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_GetContractCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetContractCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bright.contract.Manager/GetContractCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetContractCode(ctx, req.(*GetContractCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CompileContractCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileContractCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CompileContractCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bright.contract.Manager/CompileContractCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CompileContractCode(ctx, req.(*CompileContractCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateContractWithAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContractWithAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateContractWithAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bright.contract.Manager/CreateContractWithAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateContractWithAccount(ctx, req.(*CreateContractWithAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bright.contract.Manager/GetContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetContract(ctx, req.(*GetContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bright.contract.Manager/DeleteContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteContract(ctx, req.(*DeleteContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bright.contract.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContractCode",
			Handler:    _Manager_GetContractCode_Handler,
		},
		{
			MethodName: "CompileContractCode",
			Handler:    _Manager_CompileContractCode_Handler,
		},
		{
			MethodName: "CreateContractWithAccount",
			Handler:    _Manager_CreateContractWithAccount_Handler,
		},
		{
			MethodName: "GetContract",
			Handler:    _Manager_GetContract_Handler,
		},
		{
			MethodName: "DeleteContract",
			Handler:    _Manager_DeleteContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bright/contract/contract.proto",
}
