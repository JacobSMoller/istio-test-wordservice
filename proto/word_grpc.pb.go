// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wordservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WordserviceClient is the client API for Wordservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WordserviceClient interface {
	// Sends a greeting
	UrbandictionaryLookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupReply, error)
}

type wordserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewWordserviceClient(cc grpc.ClientConnInterface) WordserviceClient {
	return &wordserviceClient{cc}
}

func (c *wordserviceClient) UrbandictionaryLookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupReply, error) {
	out := new(LookupReply)
	err := c.cc.Invoke(ctx, "/wordservice.Wordservice/UrbandictionaryLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordserviceServer is the server API for Wordservice service.
// All implementations must embed UnimplementedWordserviceServer
// for forward compatibility
type WordserviceServer interface {
	// Sends a greeting
	UrbandictionaryLookup(context.Context, *LookupRequest) (*LookupReply, error)
	mustEmbedUnimplementedWordserviceServer()
}

// UnimplementedWordserviceServer must be embedded to have forward compatible implementations.
type UnimplementedWordserviceServer struct {
}

func (*UnimplementedWordserviceServer) UrbandictionaryLookup(context.Context, *LookupRequest) (*LookupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UrbandictionaryLookup not implemented")
}
func (*UnimplementedWordserviceServer) mustEmbedUnimplementedWordserviceServer() {}

func RegisterWordserviceServer(s *grpc.Server, srv WordserviceServer) {
	s.RegisterService(&_Wordservice_serviceDesc, srv)
}

func _Wordservice_UrbandictionaryLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordserviceServer).UrbandictionaryLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordservice.Wordservice/UrbandictionaryLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordserviceServer).UrbandictionaryLookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wordservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wordservice.Wordservice",
	HandlerType: (*WordserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UrbandictionaryLookup",
			Handler:    _Wordservice_UrbandictionaryLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "word.proto",
}
