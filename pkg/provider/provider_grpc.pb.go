// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package provider

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProviderPluginClient is the client API for ProviderPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderPluginClient interface {
	Init(ctx context.Context, in *ProviderInit_Request, opts ...grpc.CallOption) (*ProviderInit_Response, error)
	Name(ctx context.Context, in *ProviderName_Request, opts ...grpc.CallOption) (*ProviderName_Response, error)
	GetInfo(ctx context.Context, in *GetInfo_Request, opts ...grpc.CallOption) (*GetInfo_Response, error)
	GetCommits(ctx context.Context, in *GetCommits_Request, opts ...grpc.CallOption) (*GetCommits_Response, error)
	GetReleases(ctx context.Context, in *GetReleases_Request, opts ...grpc.CallOption) (*GetReleases_Response, error)
	CreateRelease(ctx context.Context, in *CreateRelease_Request, opts ...grpc.CallOption) (*CreateRelease_Response, error)
}

type providerPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderPluginClient(cc grpc.ClientConnInterface) ProviderPluginClient {
	return &providerPluginClient{cc}
}

func (c *providerPluginClient) Init(ctx context.Context, in *ProviderInit_Request, opts ...grpc.CallOption) (*ProviderInit_Response, error) {
	out := new(ProviderInit_Response)
	err := c.cc.Invoke(ctx, "/ProviderPlugin/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) Name(ctx context.Context, in *ProviderName_Request, opts ...grpc.CallOption) (*ProviderName_Response, error) {
	out := new(ProviderName_Response)
	err := c.cc.Invoke(ctx, "/ProviderPlugin/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) GetInfo(ctx context.Context, in *GetInfo_Request, opts ...grpc.CallOption) (*GetInfo_Response, error) {
	out := new(GetInfo_Response)
	err := c.cc.Invoke(ctx, "/ProviderPlugin/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) GetCommits(ctx context.Context, in *GetCommits_Request, opts ...grpc.CallOption) (*GetCommits_Response, error) {
	out := new(GetCommits_Response)
	err := c.cc.Invoke(ctx, "/ProviderPlugin/GetCommits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) GetReleases(ctx context.Context, in *GetReleases_Request, opts ...grpc.CallOption) (*GetReleases_Response, error) {
	out := new(GetReleases_Response)
	err := c.cc.Invoke(ctx, "/ProviderPlugin/GetReleases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) CreateRelease(ctx context.Context, in *CreateRelease_Request, opts ...grpc.CallOption) (*CreateRelease_Response, error) {
	out := new(CreateRelease_Response)
	err := c.cc.Invoke(ctx, "/ProviderPlugin/CreateRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderPluginServer is the server API for ProviderPlugin service.
// All implementations must embed UnimplementedProviderPluginServer
// for forward compatibility
type ProviderPluginServer interface {
	Init(context.Context, *ProviderInit_Request) (*ProviderInit_Response, error)
	Name(context.Context, *ProviderName_Request) (*ProviderName_Response, error)
	GetInfo(context.Context, *GetInfo_Request) (*GetInfo_Response, error)
	GetCommits(context.Context, *GetCommits_Request) (*GetCommits_Response, error)
	GetReleases(context.Context, *GetReleases_Request) (*GetReleases_Response, error)
	CreateRelease(context.Context, *CreateRelease_Request) (*CreateRelease_Response, error)
	mustEmbedUnimplementedProviderPluginServer()
}

// UnimplementedProviderPluginServer must be embedded to have forward compatible implementations.
type UnimplementedProviderPluginServer struct {
}

func (*UnimplementedProviderPluginServer) Init(context.Context, *ProviderInit_Request) (*ProviderInit_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedProviderPluginServer) Name(context.Context, *ProviderName_Request) (*ProviderName_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (*UnimplementedProviderPluginServer) GetInfo(context.Context, *GetInfo_Request) (*GetInfo_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (*UnimplementedProviderPluginServer) GetCommits(context.Context, *GetCommits_Request) (*GetCommits_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommits not implemented")
}
func (*UnimplementedProviderPluginServer) GetReleases(context.Context, *GetReleases_Request) (*GetReleases_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReleases not implemented")
}
func (*UnimplementedProviderPluginServer) CreateRelease(context.Context, *CreateRelease_Request) (*CreateRelease_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelease not implemented")
}
func (*UnimplementedProviderPluginServer) mustEmbedUnimplementedProviderPluginServer() {}

func RegisterProviderPluginServer(s *grpc.Server, srv ProviderPluginServer) {
	s.RegisterService(&_ProviderPlugin_serviceDesc, srv)
}

func _ProviderPlugin_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderInit_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProviderPlugin/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).Init(ctx, req.(*ProviderInit_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderName_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProviderPlugin/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).Name(ctx, req.(*ProviderName_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfo_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProviderPlugin/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).GetInfo(ctx, req.(*GetInfo_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_GetCommits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommits_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).GetCommits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProviderPlugin/GetCommits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).GetCommits(ctx, req.(*GetCommits_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_GetReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleases_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).GetReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProviderPlugin/GetReleases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).GetReleases(ctx, req.(*GetReleases_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_CreateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelease_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).CreateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProviderPlugin/CreateRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).CreateRelease(ctx, req.(*CreateRelease_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProviderPlugin",
	HandlerType: (*ProviderPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ProviderPlugin_Init_Handler,
		},
		{
			MethodName: "Name",
			Handler:    _ProviderPlugin_Name_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _ProviderPlugin_GetInfo_Handler,
		},
		{
			MethodName: "GetCommits",
			Handler:    _ProviderPlugin_GetCommits_Handler,
		},
		{
			MethodName: "GetReleases",
			Handler:    _ProviderPlugin_GetReleases_Handler,
		},
		{
			MethodName: "CreateRelease",
			Handler:    _ProviderPlugin_CreateRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/provider/provider.proto",
}
