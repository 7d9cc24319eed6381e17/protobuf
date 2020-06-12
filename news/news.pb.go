// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: news.proto

package news

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	ClientId   string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Symbol     string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *SubscribeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SubscribeRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	ClientId   string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Symbol     string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{1}
}

func (x *UnsubscribeRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *UnsubscribeRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UnsubscribeRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type ShowSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId string `protobuf:"bytes,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	ClientId   string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ShowSubscriptionsRequest) Reset() {
	*x = ShowSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowSubscriptionsRequest) ProtoMessage() {}

func (x *ShowSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*ShowSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{2}
}

func (x *ShowSubscriptionsRequest) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *ShowSubscriptionsRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ShowSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []string `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *ShowSubscriptionsResponse) Reset() {
	*x = ShowSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowSubscriptionsResponse) ProtoMessage() {}

func (x *ShowSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*ShowSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{3}
}

func (x *ShowSubscriptionsResponse) GetSubscriptions() []string {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

var File_news_proto protoreflect.FileDescriptor

var file_news_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x65,
	0x77, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x6a, 0x0a, 0x12, 0x55, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x58, 0x0a, 0x18, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x19, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0xe0, 0x01, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x3d, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x73,
	0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x11, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x37, 0x64, 0x39, 0x63, 0x63, 0x32, 0x34, 0x33, 0x31, 0x39, 0x65, 0x65,
	0x64, 0x36, 0x33, 0x38, 0x31, 0x65, 0x31, 0x37, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_news_proto_rawDescOnce sync.Once
	file_news_proto_rawDescData = file_news_proto_rawDesc
)

func file_news_proto_rawDescGZIP() []byte {
	file_news_proto_rawDescOnce.Do(func() {
		file_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_news_proto_rawDescData)
	})
	return file_news_proto_rawDescData
}

var file_news_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_news_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),          // 0: news.SubscribeRequest
	(*UnsubscribeRequest)(nil),        // 1: news.UnsubscribeRequest
	(*ShowSubscriptionsRequest)(nil),  // 2: news.ShowSubscriptionsRequest
	(*ShowSubscriptionsResponse)(nil), // 3: news.ShowSubscriptionsResponse
	(*empty.Empty)(nil),               // 4: google.protobuf.Empty
}
var file_news_proto_depIdxs = []int32{
	0, // 0: news.News.Subscribe:input_type -> news.SubscribeRequest
	1, // 1: news.News.Unsubscribe:input_type -> news.UnsubscribeRequest
	2, // 2: news.News.ShowSubscriptions:input_type -> news.ShowSubscriptionsRequest
	4, // 3: news.News.Subscribe:output_type -> google.protobuf.Empty
	4, // 4: news.News.Unsubscribe:output_type -> google.protobuf.Empty
	3, // 5: news.News.ShowSubscriptions:output_type -> news.ShowSubscriptionsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_news_proto_init() }
func file_news_proto_init() {
	if File_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowSubscriptionsRequest); i {
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
		file_news_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowSubscriptionsResponse); i {
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
			RawDescriptor: file_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_news_proto_goTypes,
		DependencyIndexes: file_news_proto_depIdxs,
		MessageInfos:      file_news_proto_msgTypes,
	}.Build()
	File_news_proto = out.File
	file_news_proto_rawDesc = nil
	file_news_proto_goTypes = nil
	file_news_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NewsClient is the client API for News service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NewsClient interface {
	// Subscriptions
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ShowSubscriptions(ctx context.Context, in *ShowSubscriptionsRequest, opts ...grpc.CallOption) (*ShowSubscriptionsResponse, error)
}

type newsClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsClient(cc grpc.ClientConnInterface) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/news.News/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/news.News/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) ShowSubscriptions(ctx context.Context, in *ShowSubscriptionsRequest, opts ...grpc.CallOption) (*ShowSubscriptionsResponse, error) {
	out := new(ShowSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/news.News/ShowSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServer is the server API for News service.
type NewsServer interface {
	// Subscriptions
	Subscribe(context.Context, *SubscribeRequest) (*empty.Empty, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*empty.Empty, error)
	ShowSubscriptions(context.Context, *ShowSubscriptionsRequest) (*ShowSubscriptionsResponse, error)
}

// UnimplementedNewsServer can be embedded to have forward compatible implementations.
type UnimplementedNewsServer struct {
}

func (*UnimplementedNewsServer) Subscribe(context.Context, *SubscribeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedNewsServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (*UnimplementedNewsServer) ShowSubscriptions(context.Context, *ShowSubscriptionsRequest) (*ShowSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowSubscriptions not implemented")
}

func RegisterNewsServer(s *grpc.Server, srv NewsServer) {
	s.RegisterService(&_News_serviceDesc, srv)
}

func _News_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_ShowSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).ShowSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.News/ShowSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).ShowSubscriptions(ctx, req.(*ShowSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _News_serviceDesc = grpc.ServiceDesc{
	ServiceName: "news.News",
	HandlerType: (*NewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _News_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _News_Unsubscribe_Handler,
		},
		{
			MethodName: "ShowSubscriptions",
			Handler:    _News_ShowSubscriptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}
