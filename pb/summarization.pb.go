// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: summarization.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *SummaryRequest) Reset() {
	*x = SummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summarization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryRequest) ProtoMessage() {}

func (x *SummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_summarization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryRequest.ProtoReflect.Descriptor instead.
func (*SummaryRequest) Descriptor() ([]byte, []int) {
	return file_summarization_proto_rawDescGZIP(), []int{0}
}

func (x *SummaryRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type SummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *SummaryResponse) Reset() {
	*x = SummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_summarization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryResponse) ProtoMessage() {}

func (x *SummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_summarization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryResponse.ProtoReflect.Descriptor instead.
func (*SummaryResponse) Descriptor() ([]byte, []int) {
	return file_summarization_proto_rawDescGZIP(), []int{1}
}

func (x *SummaryResponse) GetResponse() []string {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_summarization_proto protoreflect.FileDescriptor

var file_summarization_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x46, 0x0a, 0x0d, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_summarization_proto_rawDescOnce sync.Once
	file_summarization_proto_rawDescData = file_summarization_proto_rawDesc
)

func file_summarization_proto_rawDescGZIP() []byte {
	file_summarization_proto_rawDescOnce.Do(func() {
		file_summarization_proto_rawDescData = protoimpl.X.CompressGZIP(file_summarization_proto_rawDescData)
	})
	return file_summarization_proto_rawDescData
}

var file_summarization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_summarization_proto_goTypes = []interface{}{
	(*SummaryRequest)(nil),  // 0: pb.SummaryRequest
	(*SummaryResponse)(nil), // 1: pb.SummaryResponse
}
var file_summarization_proto_depIdxs = []int32{
	0, // 0: pb.Summarization.GetSummary:input_type -> pb.SummaryRequest
	1, // 1: pb.Summarization.GetSummary:output_type -> pb.SummaryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_summarization_proto_init() }
func file_summarization_proto_init() {
	if File_summarization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_summarization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_summarization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_summarization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_summarization_proto_goTypes,
		DependencyIndexes: file_summarization_proto_depIdxs,
		MessageInfos:      file_summarization_proto_msgTypes,
	}.Build()
	File_summarization_proto = out.File
	file_summarization_proto_rawDesc = nil
	file_summarization_proto_goTypes = nil
	file_summarization_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SummarizationClient is the client API for Summarization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SummarizationClient interface {
	GetSummary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (*SummaryResponse, error)
}

type summarizationClient struct {
	cc grpc.ClientConnInterface
}

func NewSummarizationClient(cc grpc.ClientConnInterface) SummarizationClient {
	return &summarizationClient{cc}
}

func (c *summarizationClient) GetSummary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (*SummaryResponse, error) {
	out := new(SummaryResponse)
	err := c.cc.Invoke(ctx, "/pb.Summarization/GetSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummarizationServer is the server API for Summarization service.
type SummarizationServer interface {
	GetSummary(context.Context, *SummaryRequest) (*SummaryResponse, error)
}

// UnimplementedSummarizationServer can be embedded to have forward compatible implementations.
type UnimplementedSummarizationServer struct {
}

func (*UnimplementedSummarizationServer) GetSummary(context.Context, *SummaryRequest) (*SummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummary not implemented")
}

func RegisterSummarizationServer(s *grpc.Server, srv SummarizationServer) {
	s.RegisterService(&_Summarization_serviceDesc, srv)
}

func _Summarization_GetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummarizationServer).GetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Summarization/GetSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummarizationServer).GetSummary(ctx, req.(*SummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Summarization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Summarization",
	HandlerType: (*SummarizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSummary",
			Handler:    _Summarization_GetSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "summarization.proto",
}
