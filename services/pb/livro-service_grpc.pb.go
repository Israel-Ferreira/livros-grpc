// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: livro-service.proto

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

// LivroServiceClient is the client API for LivroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LivroServiceClient interface {
	FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Livro, error)
	DeleteById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Livro, error)
	FindAll(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*LivroList, error)
	Create(ctx context.Context, in *LivroRequest, opts ...grpc.CallOption) (*Livro, error)
}

type livroServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLivroServiceClient(cc grpc.ClientConnInterface) LivroServiceClient {
	return &livroServiceClient{cc}
}

func (c *livroServiceClient) FindById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Livro, error) {
	out := new(Livro)
	err := c.cc.Invoke(ctx, "/livro.LivroService/FindById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livroServiceClient) DeleteById(ctx context.Context, in *FindByIdRequest, opts ...grpc.CallOption) (*Livro, error) {
	out := new(Livro)
	err := c.cc.Invoke(ctx, "/livro.LivroService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livroServiceClient) FindAll(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*LivroList, error) {
	out := new(LivroList)
	err := c.cc.Invoke(ctx, "/livro.LivroService/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *livroServiceClient) Create(ctx context.Context, in *LivroRequest, opts ...grpc.CallOption) (*Livro, error) {
	out := new(Livro)
	err := c.cc.Invoke(ctx, "/livro.LivroService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LivroServiceServer is the server API for LivroService service.
// All implementations must embed UnimplementedLivroServiceServer
// for forward compatibility
type LivroServiceServer interface {
	FindById(context.Context, *FindByIdRequest) (*Livro, error)
	DeleteById(context.Context, *FindByIdRequest) (*Livro, error)
	FindAll(context.Context, *EmptyRequest) (*LivroList, error)
	Create(context.Context, *LivroRequest) (*Livro, error)
	mustEmbedUnimplementedLivroServiceServer()
}

// UnimplementedLivroServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLivroServiceServer struct {
}

func (UnimplementedLivroServiceServer) FindById(context.Context, *FindByIdRequest) (*Livro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedLivroServiceServer) DeleteById(context.Context, *FindByIdRequest) (*Livro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedLivroServiceServer) FindAll(context.Context, *EmptyRequest) (*LivroList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedLivroServiceServer) Create(context.Context, *LivroRequest) (*Livro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLivroServiceServer) mustEmbedUnimplementedLivroServiceServer() {}

// UnsafeLivroServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LivroServiceServer will
// result in compilation errors.
type UnsafeLivroServiceServer interface {
	mustEmbedUnimplementedLivroServiceServer()
}

func RegisterLivroServiceServer(s grpc.ServiceRegistrar, srv LivroServiceServer) {
	s.RegisterService(&LivroService_ServiceDesc, srv)
}

func _LivroService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivroServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livro.LivroService/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivroServiceServer).FindById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivroService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivroServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livro.LivroService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivroServiceServer).DeleteById(ctx, req.(*FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivroService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivroServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livro.LivroService/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivroServiceServer).FindAll(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LivroService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivroServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livro.LivroService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivroServiceServer).Create(ctx, req.(*LivroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LivroService_ServiceDesc is the grpc.ServiceDesc for LivroService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LivroService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "livro.LivroService",
	HandlerType: (*LivroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindById",
			Handler:    _LivroService_FindById_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _LivroService_DeleteById_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _LivroService_FindAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LivroService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "livro-service.proto",
}
