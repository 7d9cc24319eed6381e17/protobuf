// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: iex.proto

package iex

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// Search
type StockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol       string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	SecurityName string `protobuf:"bytes,2,opt,name=securityName,proto3" json:"securityName,omitempty"`
	SecurityType string `protobuf:"bytes,3,opt,name=securityType,proto3" json:"securityType,omitempty"`
	Region       string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Exchange     string `protobuf:"bytes,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *StockInfo) Reset() {
	*x = StockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockInfo) ProtoMessage() {}

func (x *StockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockInfo.ProtoReflect.Descriptor instead.
func (*StockInfo) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{0}
}

func (x *StockInfo) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *StockInfo) GetSecurityName() string {
	if x != nil {
		return x.SecurityName
	}
	return ""
}

func (x *StockInfo) GetSecurityType() string {
	if x != nil {
		return x.SecurityType
	}
	return ""
}

func (x *StockInfo) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *StockInfo) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type StockSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fragment string `protobuf:"bytes,1,opt,name=fragment,proto3" json:"fragment,omitempty"`
}

func (x *StockSearchRequest) Reset() {
	*x = StockSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockSearchRequest) ProtoMessage() {}

func (x *StockSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockSearchRequest.ProtoReflect.Descriptor instead.
func (*StockSearchRequest) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{1}
}

func (x *StockSearchRequest) GetFragment() string {
	if x != nil {
		return x.Fragment
	}
	return ""
}

type StockSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockInfo []*StockInfo `protobuf:"bytes,1,rep,name=stockInfo,proto3" json:"stockInfo,omitempty"`
}

func (x *StockSearchResponse) Reset() {
	*x = StockSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockSearchResponse) ProtoMessage() {}

func (x *StockSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockSearchResponse.ProtoReflect.Descriptor instead.
func (*StockSearchResponse) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{2}
}

func (x *StockSearchResponse) GetStockInfo() []*StockInfo {
	if x != nil {
		return x.StockInfo
	}
	return nil
}

// Quotes
type StockQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *StockQuoteRequest) Reset() {
	*x = StockQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockQuoteRequest) ProtoMessage() {}

func (x *StockQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockQuoteRequest.ProtoReflect.Descriptor instead.
func (*StockQuoteRequest) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{3}
}

func (x *StockQuoteRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type StockQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price  float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *StockQuoteResponse) Reset() {
	*x = StockQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockQuoteResponse) ProtoMessage() {}

func (x *StockQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockQuoteResponse.ProtoReflect.Descriptor instead.
func (*StockQuoteResponse) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{4}
}

func (x *StockQuoteResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *StockQuoteResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type StockQuoteRequestBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbols []*StockQuoteRequest `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
}

func (x *StockQuoteRequestBatch) Reset() {
	*x = StockQuoteRequestBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockQuoteRequestBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockQuoteRequestBatch) ProtoMessage() {}

func (x *StockQuoteRequestBatch) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockQuoteRequestBatch.ProtoReflect.Descriptor instead.
func (*StockQuoteRequestBatch) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{5}
}

func (x *StockQuoteRequestBatch) GetSymbols() []*StockQuoteRequest {
	if x != nil {
		return x.Symbols
	}
	return nil
}

type StockQuoteResponseBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quotes []*StockQuoteResponse `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
}

func (x *StockQuoteResponseBatch) Reset() {
	*x = StockQuoteResponseBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockQuoteResponseBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockQuoteResponseBatch) ProtoMessage() {}

func (x *StockQuoteResponseBatch) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockQuoteResponseBatch.ProtoReflect.Descriptor instead.
func (*StockQuoteResponseBatch) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{6}
}

func (x *StockQuoteResponseBatch) GetQuotes() []*StockQuoteResponse {
	if x != nil {
		return x.Quotes
	}
	return nil
}

// NEWS
type StockNewsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datetime   int64  `protobuf:"varint,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Headline   string `protobuf:"bytes,2,opt,name=headline,proto3" json:"headline,omitempty"`
	Source     string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Url        string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Summary    string `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Related    string `protobuf:"bytes,6,opt,name=related,proto3" json:"related,omitempty"`
	Image      string `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	Lang       string `protobuf:"bytes,8,opt,name=lang,proto3" json:"lang,omitempty"`
	HasPaywall string `protobuf:"bytes,9,opt,name=hasPaywall,proto3" json:"hasPaywall,omitempty"`
}

func (x *StockNewsData) Reset() {
	*x = StockNewsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockNewsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockNewsData) ProtoMessage() {}

func (x *StockNewsData) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockNewsData.ProtoReflect.Descriptor instead.
func (*StockNewsData) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{7}
}

func (x *StockNewsData) GetDatetime() int64 {
	if x != nil {
		return x.Datetime
	}
	return 0
}

func (x *StockNewsData) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *StockNewsData) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *StockNewsData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StockNewsData) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *StockNewsData) GetRelated() string {
	if x != nil {
		return x.Related
	}
	return ""
}

func (x *StockNewsData) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *StockNewsData) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *StockNewsData) GetHasPaywall() string {
	if x != nil {
		return x.HasPaywall
	}
	return ""
}

type StockNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *StockNewsRequest) Reset() {
	*x = StockNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockNewsRequest) ProtoMessage() {}

func (x *StockNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockNewsRequest.ProtoReflect.Descriptor instead.
func (*StockNewsRequest) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{8}
}

func (x *StockNewsRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

// Default request contains 10 latest news
type StockNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News []*StockNewsData `protobuf:"bytes,1,rep,name=news,proto3" json:"news,omitempty"`
}

func (x *StockNewsResponse) Reset() {
	*x = StockNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockNewsResponse) ProtoMessage() {}

func (x *StockNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockNewsResponse.ProtoReflect.Descriptor instead.
func (*StockNewsResponse) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{9}
}

func (x *StockNewsResponse) GetNews() []*StockNewsData {
	if x != nil {
		return x.News
	}
	return nil
}

// Return just latest
type StockNewsResponseLatest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News *StockNewsData `protobuf:"bytes,1,opt,name=news,proto3" json:"news,omitempty"`
}

func (x *StockNewsResponseLatest) Reset() {
	*x = StockNewsResponseLatest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iex_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockNewsResponseLatest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockNewsResponseLatest) ProtoMessage() {}

func (x *StockNewsResponseLatest) ProtoReflect() protoreflect.Message {
	mi := &file_iex_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockNewsResponseLatest.ProtoReflect.Descriptor instead.
func (*StockNewsResponseLatest) Descriptor() ([]byte, []int) {
	return file_iex_proto_rawDescGZIP(), []int{10}
}

func (x *StockNewsResponseLatest) GetNews() *StockNewsData {
	if x != nil {
		return x.News
	}
	return nil
}

var File_iex_proto protoreflect.FileDescriptor

var file_iex_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x65, 0x78,
	0x22, 0x9f, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x22, 0x30, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x16, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x22, 0x4a, 0x0a, 0x17, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x2f, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x22, 0xef, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x50, 0x61, 0x79, 0x77, 0x61,
	0x6c, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x50, 0x61, 0x79,
	0x77, 0x61, 0x6c, 0x6c, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x22, 0x3b, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e,
	0x65, 0x77, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x22, 0x41, 0x0a,
	0x17, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x4e, 0x65, 0x77, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73,
	0x32, 0xf2, 0x02, 0x0a, 0x03, 0x49, 0x65, 0x78, 0x12, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x65, 0x78, 0x2e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x1b, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x1a,
	0x1c, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x69,
	0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x15, 0x2e, 0x69,
	0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x15,
	0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x37, 0x64, 0x39, 0x63, 0x63, 0x32, 0x34, 0x33, 0x31, 0x39, 0x65, 0x65,
	0x64, 0x36, 0x33, 0x38, 0x31, 0x65, 0x31, 0x37, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x69, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iex_proto_rawDescOnce sync.Once
	file_iex_proto_rawDescData = file_iex_proto_rawDesc
)

func file_iex_proto_rawDescGZIP() []byte {
	file_iex_proto_rawDescOnce.Do(func() {
		file_iex_proto_rawDescData = protoimpl.X.CompressGZIP(file_iex_proto_rawDescData)
	})
	return file_iex_proto_rawDescData
}

var file_iex_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_iex_proto_goTypes = []interface{}{
	(*StockInfo)(nil),               // 0: iex.StockInfo
	(*StockSearchRequest)(nil),      // 1: iex.StockSearchRequest
	(*StockSearchResponse)(nil),     // 2: iex.StockSearchResponse
	(*StockQuoteRequest)(nil),       // 3: iex.StockQuoteRequest
	(*StockQuoteResponse)(nil),      // 4: iex.StockQuoteResponse
	(*StockQuoteRequestBatch)(nil),  // 5: iex.StockQuoteRequestBatch
	(*StockQuoteResponseBatch)(nil), // 6: iex.StockQuoteResponseBatch
	(*StockNewsData)(nil),           // 7: iex.StockNewsData
	(*StockNewsRequest)(nil),        // 8: iex.StockNewsRequest
	(*StockNewsResponse)(nil),       // 9: iex.StockNewsResponse
	(*StockNewsResponseLatest)(nil), // 10: iex.StockNewsResponseLatest
}
var file_iex_proto_depIdxs = []int32{
	0,  // 0: iex.StockSearchResponse.stockInfo:type_name -> iex.StockInfo
	3,  // 1: iex.StockQuoteRequestBatch.symbols:type_name -> iex.StockQuoteRequest
	4,  // 2: iex.StockQuoteResponseBatch.quotes:type_name -> iex.StockQuoteResponse
	7,  // 3: iex.StockNewsResponse.news:type_name -> iex.StockNewsData
	7,  // 4: iex.StockNewsResponseLatest.news:type_name -> iex.StockNewsData
	3,  // 5: iex.Iex.GetRealTimeStockQuote:input_type -> iex.StockQuoteRequest
	5,  // 6: iex.Iex.GetRealTimeStockQuoteBatch:input_type -> iex.StockQuoteRequestBatch
	1,  // 7: iex.Iex.FindStock:input_type -> iex.StockSearchRequest
	8,  // 8: iex.Iex.GetNews:input_type -> iex.StockNewsRequest
	8,  // 9: iex.Iex.GetNewsLatest:input_type -> iex.StockNewsRequest
	4,  // 10: iex.Iex.GetRealTimeStockQuote:output_type -> iex.StockQuoteResponse
	6,  // 11: iex.Iex.GetRealTimeStockQuoteBatch:output_type -> iex.StockQuoteResponseBatch
	2,  // 12: iex.Iex.FindStock:output_type -> iex.StockSearchResponse
	9,  // 13: iex.Iex.GetNews:output_type -> iex.StockNewsResponse
	10, // 14: iex.Iex.GetNewsLatest:output_type -> iex.StockNewsResponseLatest
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_iex_proto_init() }
func file_iex_proto_init() {
	if File_iex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockInfo); i {
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
		file_iex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockSearchRequest); i {
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
		file_iex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockSearchResponse); i {
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
		file_iex_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockQuoteRequest); i {
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
		file_iex_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockQuoteResponse); i {
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
		file_iex_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockQuoteRequestBatch); i {
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
		file_iex_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockQuoteResponseBatch); i {
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
		file_iex_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockNewsData); i {
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
		file_iex_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockNewsRequest); i {
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
		file_iex_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockNewsResponse); i {
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
		file_iex_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockNewsResponseLatest); i {
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
			RawDescriptor: file_iex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iex_proto_goTypes,
		DependencyIndexes: file_iex_proto_depIdxs,
		MessageInfos:      file_iex_proto_msgTypes,
	}.Build()
	File_iex_proto = out.File
	file_iex_proto_rawDesc = nil
	file_iex_proto_goTypes = nil
	file_iex_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IexClient is the client API for Iex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IexClient interface {
	// Quotes
	GetRealTimeStockQuote(ctx context.Context, in *StockQuoteRequest, opts ...grpc.CallOption) (*StockQuoteResponse, error)
	GetRealTimeStockQuoteBatch(ctx context.Context, in *StockQuoteRequestBatch, opts ...grpc.CallOption) (*StockQuoteResponseBatch, error)
	// Search
	FindStock(ctx context.Context, in *StockSearchRequest, opts ...grpc.CallOption) (*StockSearchResponse, error)
	// News
	GetNews(ctx context.Context, in *StockNewsRequest, opts ...grpc.CallOption) (*StockNewsResponse, error)
	GetNewsLatest(ctx context.Context, in *StockNewsRequest, opts ...grpc.CallOption) (*StockNewsResponseLatest, error)
}

type iexClient struct {
	cc grpc.ClientConnInterface
}

func NewIexClient(cc grpc.ClientConnInterface) IexClient {
	return &iexClient{cc}
}

func (c *iexClient) GetRealTimeStockQuote(ctx context.Context, in *StockQuoteRequest, opts ...grpc.CallOption) (*StockQuoteResponse, error) {
	out := new(StockQuoteResponse)
	err := c.cc.Invoke(ctx, "/iex.Iex/GetRealTimeStockQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iexClient) GetRealTimeStockQuoteBatch(ctx context.Context, in *StockQuoteRequestBatch, opts ...grpc.CallOption) (*StockQuoteResponseBatch, error) {
	out := new(StockQuoteResponseBatch)
	err := c.cc.Invoke(ctx, "/iex.Iex/GetRealTimeStockQuoteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iexClient) FindStock(ctx context.Context, in *StockSearchRequest, opts ...grpc.CallOption) (*StockSearchResponse, error) {
	out := new(StockSearchResponse)
	err := c.cc.Invoke(ctx, "/iex.Iex/FindStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iexClient) GetNews(ctx context.Context, in *StockNewsRequest, opts ...grpc.CallOption) (*StockNewsResponse, error) {
	out := new(StockNewsResponse)
	err := c.cc.Invoke(ctx, "/iex.Iex/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iexClient) GetNewsLatest(ctx context.Context, in *StockNewsRequest, opts ...grpc.CallOption) (*StockNewsResponseLatest, error) {
	out := new(StockNewsResponseLatest)
	err := c.cc.Invoke(ctx, "/iex.Iex/GetNewsLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IexServer is the server API for Iex service.
type IexServer interface {
	// Quotes
	GetRealTimeStockQuote(context.Context, *StockQuoteRequest) (*StockQuoteResponse, error)
	GetRealTimeStockQuoteBatch(context.Context, *StockQuoteRequestBatch) (*StockQuoteResponseBatch, error)
	// Search
	FindStock(context.Context, *StockSearchRequest) (*StockSearchResponse, error)
	// News
	GetNews(context.Context, *StockNewsRequest) (*StockNewsResponse, error)
	GetNewsLatest(context.Context, *StockNewsRequest) (*StockNewsResponseLatest, error)
}

// UnimplementedIexServer can be embedded to have forward compatible implementations.
type UnimplementedIexServer struct {
}

func (*UnimplementedIexServer) GetRealTimeStockQuote(context.Context, *StockQuoteRequest) (*StockQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealTimeStockQuote not implemented")
}
func (*UnimplementedIexServer) GetRealTimeStockQuoteBatch(context.Context, *StockQuoteRequestBatch) (*StockQuoteResponseBatch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealTimeStockQuoteBatch not implemented")
}
func (*UnimplementedIexServer) FindStock(context.Context, *StockSearchRequest) (*StockSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStock not implemented")
}
func (*UnimplementedIexServer) GetNews(context.Context, *StockNewsRequest) (*StockNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (*UnimplementedIexServer) GetNewsLatest(context.Context, *StockNewsRequest) (*StockNewsResponseLatest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsLatest not implemented")
}

func RegisterIexServer(s *grpc.Server, srv IexServer) {
	s.RegisterService(&_Iex_serviceDesc, srv)
}

func _Iex_GetRealTimeStockQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IexServer).GetRealTimeStockQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iex.Iex/GetRealTimeStockQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IexServer).GetRealTimeStockQuote(ctx, req.(*StockQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iex_GetRealTimeStockQuoteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockQuoteRequestBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IexServer).GetRealTimeStockQuoteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iex.Iex/GetRealTimeStockQuoteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IexServer).GetRealTimeStockQuoteBatch(ctx, req.(*StockQuoteRequestBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iex_FindStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IexServer).FindStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iex.Iex/FindStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IexServer).FindStock(ctx, req.(*StockSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iex_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IexServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iex.Iex/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IexServer).GetNews(ctx, req.(*StockNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iex_GetNewsLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IexServer).GetNewsLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iex.Iex/GetNewsLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IexServer).GetNewsLatest(ctx, req.(*StockNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Iex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iex.Iex",
	HandlerType: (*IexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRealTimeStockQuote",
			Handler:    _Iex_GetRealTimeStockQuote_Handler,
		},
		{
			MethodName: "GetRealTimeStockQuoteBatch",
			Handler:    _Iex_GetRealTimeStockQuoteBatch_Handler,
		},
		{
			MethodName: "FindStock",
			Handler:    _Iex_FindStock_Handler,
		},
		{
			MethodName: "GetNews",
			Handler:    _Iex_GetNews_Handler,
		},
		{
			MethodName: "GetNewsLatest",
			Handler:    _Iex_GetNewsLatest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iex.proto",
}
