// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/api.proto

package proto

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
	Tts_Tts_FullMethodName = "/playht.v1.Tts/Tts"
)

// TtsClient is the client API for Tts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TtsClient interface {
	Tts(ctx context.Context, in chan *TtsRequest, opts ...grpc.CallOption) (Tts_TtsClient, error)
}

type ttsClient struct {
	cc grpc.ClientConnInterface
}

func NewTtsClient(cc grpc.ClientConnInterface) TtsClient {
	return &ttsClient{cc}
}

func (c *ttsClient) Tts(ctx context.Context, in chan *TtsRequest, opts ...grpc.CallOption) (Tts_TtsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tts_ServiceDesc.Streams[0], Tts_Tts_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &ttsTtsClient{stream}

	go func() {
		for elem := range in {
			if err := x.ClientStream.SendMsg(elem); err != nil {
				//do nothing right now...
				break
			}
		}
		if err := x.ClientStream.CloseSend(); err != nil {
			//do nothing right now
		}
	}()
	return x, nil
}

type Tts_TtsClient interface {
	Recv() (*TtsResponse, error)
	grpc.ClientStream
}

type ttsTtsClient struct {
	grpc.ClientStream
}

func (x *ttsTtsClient) Recv() (*TtsResponse, error) {
	m := new(TtsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TtsServer is the server API for Tts service.
// All implementations must embed UnimplementedTtsServer
// for forward compatibility
type TtsServer interface {
	Tts(*TtsRequest, Tts_TtsServer) error
	mustEmbedUnimplementedTtsServer()
}

// UnimplementedTtsServer must be embedded to have forward compatible implementations.
type UnimplementedTtsServer struct {
}

func (UnimplementedTtsServer) Tts(*TtsRequest, Tts_TtsServer) error {
	return status.Errorf(codes.Unimplemented, "method Tts not implemented")
}
func (UnimplementedTtsServer) mustEmbedUnimplementedTtsServer() {}

// UnsafeTtsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TtsServer will
// result in compilation errors.
type UnsafeTtsServer interface {
	mustEmbedUnimplementedTtsServer()
}

func RegisterTtsServer(s grpc.ServiceRegistrar, srv TtsServer) {
	s.RegisterService(&Tts_ServiceDesc, srv)
}

func _Tts_Tts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TtsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TtsServer).Tts(m, &ttsTtsServer{stream})
}

type Tts_TtsServer interface {
	Send(*TtsResponse) error
	grpc.ServerStream
}

type ttsTtsServer struct {
	grpc.ServerStream
}

func (x *ttsTtsServer) Send(m *TtsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Tts_ServiceDesc is the grpc.ServiceDesc for Tts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "playht.v1.Tts",
	HandlerType: (*TtsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tts",
			Handler:       _Tts_Tts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api.proto",
}
