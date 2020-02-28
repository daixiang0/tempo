// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frigg.proto

/*
Package friggpb is a generated protocol buffer package.

It is generated from these files:
	frigg.proto

It has these top-level messages:
	TraceByIDRequest
	TraceByIDResponse
	Trace
	PushRequest
	PushResponse
*/
package friggpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import opentelemetry_proto_trace_v1 "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TraceByIDRequest struct {
	TraceID []byte `protobuf:"bytes,1,opt,name=traceID,proto3" json:"traceID,omitempty"`
}

func (m *TraceByIDRequest) Reset()                    { *m = TraceByIDRequest{} }
func (m *TraceByIDRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceByIDRequest) ProtoMessage()               {}
func (*TraceByIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TraceByIDRequest) GetTraceID() []byte {
	if m != nil {
		return m.TraceID
	}
	return nil
}

type TraceByIDResponse struct {
	Trace *Trace `protobuf:"bytes,1,opt,name=trace" json:"trace,omitempty"`
}

func (m *TraceByIDResponse) Reset()                    { *m = TraceByIDResponse{} }
func (m *TraceByIDResponse) String() string            { return proto.CompactTextString(m) }
func (*TraceByIDResponse) ProtoMessage()               {}
func (*TraceByIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TraceByIDResponse) GetTrace() *Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

type Trace struct {
	Batches []*opentelemetry_proto_trace_v1.ResourceSpans `protobuf:"bytes,1,rep,name=batches" json:"batches,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Trace) GetBatches() []*opentelemetry_proto_trace_v1.ResourceSpans {
	if m != nil {
		return m.Batches
	}
	return nil
}

type PushRequest struct {
	Batch *opentelemetry_proto_trace_v1.ResourceSpans `protobuf:"bytes,1,opt,name=batch" json:"batch,omitempty"`
}

func (m *PushRequest) Reset()                    { *m = PushRequest{} }
func (m *PushRequest) String() string            { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()               {}
func (*PushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PushRequest) GetBatch() *opentelemetry_proto_trace_v1.ResourceSpans {
	if m != nil {
		return m.Batch
	}
	return nil
}

type PushResponse struct {
}

func (m *PushResponse) Reset()                    { *m = PushResponse{} }
func (m *PushResponse) String() string            { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()               {}
func (*PushResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*TraceByIDRequest)(nil), "friggpb.TraceByIDRequest")
	proto.RegisterType((*TraceByIDResponse)(nil), "friggpb.TraceByIDResponse")
	proto.RegisterType((*Trace)(nil), "friggpb.Trace")
	proto.RegisterType((*PushRequest)(nil), "friggpb.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "friggpb.PushResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pusher service

type PusherClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type pusherClient struct {
	cc *grpc.ClientConn
}

func NewPusherClient(cc *grpc.ClientConn) PusherClient {
	return &pusherClient{cc}
}

func (c *pusherClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := grpc.Invoke(ctx, "/friggpb.Pusher/Push", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pusher service

type PusherServer interface {
	Push(context.Context, *PushRequest) (*PushResponse, error)
}

func RegisterPusherServer(s *grpc.Server, srv PusherServer) {
	s.RegisterService(&_Pusher_serviceDesc, srv)
}

func _Pusher_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friggpb.Pusher/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pusher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "friggpb.Pusher",
	HandlerType: (*PusherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Pusher_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "frigg.proto",
}

// Client API for Querier service

type QuerierClient interface {
	FindTraceByID(ctx context.Context, in *TraceByIDRequest, opts ...grpc.CallOption) (*TraceByIDResponse, error)
}

type querierClient struct {
	cc *grpc.ClientConn
}

func NewQuerierClient(cc *grpc.ClientConn) QuerierClient {
	return &querierClient{cc}
}

func (c *querierClient) FindTraceByID(ctx context.Context, in *TraceByIDRequest, opts ...grpc.CallOption) (*TraceByIDResponse, error) {
	out := new(TraceByIDResponse)
	err := grpc.Invoke(ctx, "/friggpb.Querier/FindTraceByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Querier service

type QuerierServer interface {
	FindTraceByID(context.Context, *TraceByIDRequest) (*TraceByIDResponse, error)
}

func RegisterQuerierServer(s *grpc.Server, srv QuerierServer) {
	s.RegisterService(&_Querier_serviceDesc, srv)
}

func _Querier_FindTraceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServer).FindTraceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friggpb.Querier/FindTraceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServer).FindTraceByID(ctx, req.(*TraceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Querier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "friggpb.Querier",
	HandlerType: (*QuerierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindTraceByID",
			Handler:    _Querier_FindTraceByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "frigg.proto",
}

func init() { proto.RegisterFile("frigg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0xfd, 0xca, 0x67, 0x57, 0xb8, 0x9d, 0x43, 0x83, 0x42, 0xed, 0xd3, 0x08, 0x3e, 0x14, 0x94,
	0x94, 0x55, 0x7c, 0xf0, 0x49, 0x94, 0x29, 0xee, 0x45, 0x66, 0xe6, 0x1f, 0x68, 0xeb, 0x75, 0x2b,
	0x68, 0x13, 0x93, 0x74, 0xb0, 0x7f, 0x2f, 0x4d, 0xd6, 0xd1, 0x89, 0x2f, 0xbe, 0xdd, 0x93, 0x73,
	0xce, 0xcd, 0xb9, 0x07, 0xc2, 0x77, 0x55, 0x2d, 0x97, 0x4c, 0x2a, 0x61, 0x04, 0x09, 0x2c, 0x90,
	0x45, 0x9c, 0x08, 0x89, 0xb5, 0xc1, 0x0f, 0xfc, 0x44, 0xa3, 0x36, 0xa9, 0x65, 0x53, 0xa3, 0xf2,
	0x12, 0xd3, 0xf5, 0xc4, 0x0d, 0xce, 0x42, 0x2f, 0xe1, 0xe8, 0xb5, 0x85, 0xf7, 0x9b, 0xd9, 0x94,
	0xe3, 0x57, 0x83, 0xda, 0x90, 0x08, 0x02, 0x2b, 0x99, 0x4d, 0x23, 0x6f, 0xec, 0x25, 0x43, 0xde,
	0x41, 0x7a, 0x03, 0xc7, 0x3d, 0xb5, 0x96, 0xa2, 0xd6, 0x48, 0xce, 0xc1, 0xb7, 0xbc, 0x15, 0x87,
	0xd9, 0x88, 0x6d, 0x53, 0x30, 0x2b, 0xe5, 0x8e, 0xa4, 0xcf, 0xe0, 0x5b, 0x4c, 0x1e, 0x20, 0x28,
	0x72, 0x53, 0xae, 0x50, 0x47, 0xde, 0xf8, 0x7f, 0x12, 0x66, 0x17, 0x6c, 0x2f, 0xad, 0x0b, 0xc6,
	0x5c, 0xc8, 0xf5, 0x84, 0x71, 0xd4, 0xa2, 0x51, 0x25, 0x2e, 0x64, 0x5e, 0x6b, 0xde, 0x79, 0xe9,
	0x1c, 0xc2, 0x79, 0xa3, 0x57, 0x5d, 0xe6, 0x3b, 0xf0, 0x2d, 0xb3, 0x0d, 0xf1, 0xa7, 0x9d, 0xce,
	0x49, 0x47, 0x30, 0x74, 0x1b, 0xdd, 0x5d, 0xd9, 0x2d, 0x0c, 0x5a, 0x8c, 0x8a, 0x5c, 0xc3, 0x41,
	0x3b, 0x91, 0x93, 0xdd, 0x69, 0xbd, 0xaf, 0xe3, 0xd3, 0x1f, 0xaf, 0xce, 0x4e, 0xff, 0x65, 0x0b,
	0x08, 0x5e, 0x1a, 0x54, 0x15, 0x2a, 0xf2, 0x04, 0x87, 0x8f, 0x55, 0xfd, 0xb6, 0x2b, 0x8f, 0x9c,
	0xed, 0xb7, 0xd4, 0xab, 0x3f, 0x8e, 0x7f, 0xa3, 0xba, 0xa5, 0xc5, 0xc0, 0x9e, 0x72, 0xf5, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xc2, 0x65, 0xd3, 0x39, 0xf9, 0x01, 0x00, 0x00,
}
