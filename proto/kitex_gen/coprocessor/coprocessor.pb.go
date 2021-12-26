// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: coprocessor.proto

package coprocessor

import (
	context "context"
	errorpb "github.com/pingcap/kvproto/proto/kitex_gen/errorpb"
	_ "github.com/pingcap/kvproto/proto/kitex_gen/github.com/gogo/protobuf/gogoproto"
	kvrpcpb "github.com/pingcap/kvproto/proto/kitex_gen/kvrpcpb"
	metapb "github.com/pingcap/kvproto/proto/kitex_gen/metapb"
	_ "github.com/pingcap/kvproto/proto/kitex_gen/rustproto"
	span "github.com/pingcap/kvproto/proto/kitex_gen/span"
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

// [start, end)
type KeyRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *KeyRange) Reset() {
	*x = KeyRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocessor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRange) ProtoMessage() {}

func (x *KeyRange) ProtoReflect() protoreflect.Message {
	mi := &file_coprocessor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRange.ProtoReflect.Descriptor instead.
func (*KeyRange) Descriptor() ([]byte, []int) {
	return file_coprocessor_proto_rawDescGZIP(), []int{0}
}

func (x *KeyRange) GetStart() []byte {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *KeyRange) GetEnd() []byte {
	if x != nil {
		return x.End
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *kvrpcpb.Context `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Tp      int64            `protobuf:"varint,2,opt,name=tp,proto3" json:"tp,omitempty"`
	Data    []byte           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	StartTs uint64           `protobuf:"varint,7,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	Ranges  []*KeyRange      `protobuf:"bytes,4,rep,name=ranges,proto3" json:"ranges,omitempty"`
	// If cache is enabled, TiKV returns cache hit instead of data if
	// its last version matches this `cache_if_match_version`.
	IsCacheEnabled      bool   `protobuf:"varint,5,opt,name=is_cache_enabled,json=isCacheEnabled,proto3" json:"is_cache_enabled,omitempty"`
	CacheIfMatchVersion uint64 `protobuf:"varint,6,opt,name=cache_if_match_version,json=cacheIfMatchVersion,proto3" json:"cache_if_match_version,omitempty"`
	// Any schema-ful storage to validate schema correctness if necessary.
	SchemaVer      int64 `protobuf:"varint,8,opt,name=schema_ver,json=schemaVer,proto3" json:"schema_ver,omitempty"`
	IsTraceEnabled bool  `protobuf:"varint,9,opt,name=is_trace_enabled,json=isTraceEnabled,proto3" json:"is_trace_enabled,omitempty"`
	// paging_size is 0 when it's disabled, otherwise, it should be a positive number.
	PagingSize uint64 `protobuf:"varint,10,opt,name=paging_size,json=pagingSize,proto3" json:"paging_size,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocessor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_coprocessor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_coprocessor_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetContext() *kvrpcpb.Context {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Request) GetTp() int64 {
	if x != nil {
		return x.Tp
	}
	return 0
}

func (x *Request) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Request) GetStartTs() uint64 {
	if x != nil {
		return x.StartTs
	}
	return 0
}

func (x *Request) GetRanges() []*KeyRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

func (x *Request) GetIsCacheEnabled() bool {
	if x != nil {
		return x.IsCacheEnabled
	}
	return false
}

func (x *Request) GetCacheIfMatchVersion() uint64 {
	if x != nil {
		return x.CacheIfMatchVersion
	}
	return 0
}

func (x *Request) GetSchemaVer() int64 {
	if x != nil {
		return x.SchemaVer
	}
	return 0
}

func (x *Request) GetIsTraceEnabled() bool {
	if x != nil {
		return x.IsTraceEnabled
	}
	return false
}

func (x *Request) GetPagingSize() uint64 {
	if x != nil {
		return x.PagingSize
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        []byte            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	RegionError *errorpb.Error    `protobuf:"bytes,2,opt,name=region_error,json=regionError,proto3" json:"region_error,omitempty"`
	Locked      *kvrpcpb.LockInfo `protobuf:"bytes,3,opt,name=locked,proto3" json:"locked,omitempty"`
	OtherError  string            `protobuf:"bytes,4,opt,name=other_error,json=otherError,proto3" json:"other_error,omitempty"`
	Range       *KeyRange         `protobuf:"bytes,5,opt,name=range,proto3" json:"range,omitempty"`
	// This field is always filled for compatibility consideration. However
	// newer TiDB should respect `exec_details_v2` field instead.
	ExecDetails *kvrpcpb.ExecDetails `protobuf:"bytes,6,opt,name=exec_details,json=execDetails,proto3" json:"exec_details,omitempty"`
	// This field is provided in later versions, containing more detailed
	// information.
	ExecDetailsV2    *kvrpcpb.ExecDetailsV2 `protobuf:"bytes,11,opt,name=exec_details_v2,json=execDetailsV2,proto3" json:"exec_details_v2,omitempty"`
	IsCacheHit       bool                   `protobuf:"varint,7,opt,name=is_cache_hit,json=isCacheHit,proto3" json:"is_cache_hit,omitempty"`
	CacheLastVersion uint64                 `protobuf:"varint,8,opt,name=cache_last_version,json=cacheLastVersion,proto3" json:"cache_last_version,omitempty"`
	CanBeCached      bool                   `protobuf:"varint,9,opt,name=can_be_cached,json=canBeCached,proto3" json:"can_be_cached,omitempty"`
	Spans            []*span.SpanSet        `protobuf:"bytes,10,rep,name=spans,proto3" json:"spans,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocessor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_coprocessor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_coprocessor_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetRegionError() *errorpb.Error {
	if x != nil {
		return x.RegionError
	}
	return nil
}

func (x *Response) GetLocked() *kvrpcpb.LockInfo {
	if x != nil {
		return x.Locked
	}
	return nil
}

func (x *Response) GetOtherError() string {
	if x != nil {
		return x.OtherError
	}
	return ""
}

func (x *Response) GetRange() *KeyRange {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *Response) GetExecDetails() *kvrpcpb.ExecDetails {
	if x != nil {
		return x.ExecDetails
	}
	return nil
}

func (x *Response) GetExecDetailsV2() *kvrpcpb.ExecDetailsV2 {
	if x != nil {
		return x.ExecDetailsV2
	}
	return nil
}

func (x *Response) GetIsCacheHit() bool {
	if x != nil {
		return x.IsCacheHit
	}
	return false
}

func (x *Response) GetCacheLastVersion() uint64 {
	if x != nil {
		return x.CacheLastVersion
	}
	return 0
}

func (x *Response) GetCanBeCached() bool {
	if x != nil {
		return x.CanBeCached
	}
	return false
}

func (x *Response) GetSpans() []*span.SpanSet {
	if x != nil {
		return x.Spans
	}
	return nil
}

type RegionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId    uint64              `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	RegionEpoch *metapb.RegionEpoch `protobuf:"bytes,2,opt,name=region_epoch,json=regionEpoch,proto3" json:"region_epoch,omitempty"`
	Ranges      []*KeyRange         `protobuf:"bytes,3,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *RegionInfo) Reset() {
	*x = RegionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocessor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionInfo) ProtoMessage() {}

func (x *RegionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_coprocessor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionInfo.ProtoReflect.Descriptor instead.
func (*RegionInfo) Descriptor() ([]byte, []int) {
	return file_coprocessor_proto_rawDescGZIP(), []int{3}
}

func (x *RegionInfo) GetRegionId() uint64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *RegionInfo) GetRegionEpoch() *metapb.RegionEpoch {
	if x != nil {
		return x.RegionEpoch
	}
	return nil
}

func (x *RegionInfo) GetRanges() []*KeyRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type BatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *kvrpcpb.Context `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Tp      int64            `protobuf:"varint,2,opt,name=tp,proto3" json:"tp,omitempty"`
	Data    []byte           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Regions []*RegionInfo    `protobuf:"bytes,4,rep,name=regions,proto3" json:"regions,omitempty"`
	StartTs uint64           `protobuf:"varint,5,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	// Any schema-ful storage to validate schema correctness if necessary.
	SchemaVer int64 `protobuf:"varint,6,opt,name=schema_ver,json=schemaVer,proto3" json:"schema_ver,omitempty"`
}

func (x *BatchRequest) Reset() {
	*x = BatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocessor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRequest) ProtoMessage() {}

func (x *BatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coprocessor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRequest.ProtoReflect.Descriptor instead.
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return file_coprocessor_proto_rawDescGZIP(), []int{4}
}

func (x *BatchRequest) GetContext() *kvrpcpb.Context {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *BatchRequest) GetTp() int64 {
	if x != nil {
		return x.Tp
	}
	return 0
}

func (x *BatchRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BatchRequest) GetRegions() []*RegionInfo {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *BatchRequest) GetStartTs() uint64 {
	if x != nil {
		return x.StartTs
	}
	return 0
}

func (x *BatchRequest) GetSchemaVer() int64 {
	if x != nil {
		return x.SchemaVer
	}
	return 0
}

type BatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         []byte               `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	OtherError   string               `protobuf:"bytes,2,opt,name=other_error,json=otherError,proto3" json:"other_error,omitempty"`
	ExecDetails  *kvrpcpb.ExecDetails `protobuf:"bytes,3,opt,name=exec_details,json=execDetails,proto3" json:"exec_details,omitempty"`
	RetryRegions []*metapb.Region     `protobuf:"bytes,4,rep,name=retry_regions,json=retryRegions,proto3" json:"retry_regions,omitempty"`
}

func (x *BatchResponse) Reset() {
	*x = BatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coprocessor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResponse) ProtoMessage() {}

func (x *BatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coprocessor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResponse.ProtoReflect.Descriptor instead.
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return file_coprocessor_proto_rawDescGZIP(), []int{5}
}

func (x *BatchResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BatchResponse) GetOtherError() string {
	if x != nil {
		return x.OtherError
	}
	return ""
}

func (x *BatchResponse) GetExecDetails() *kvrpcpb.ExecDetails {
	if x != nil {
		return x.ExecDetails
	}
	return nil
}

func (x *BatchResponse) GetRetryRegions() []*metapb.Region {
	if x != nil {
		return x.RetryRegions
	}
	return nil
}

var File_coprocessor_proto protoreflect.FileDescriptor

var file_coprocessor_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x1a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x6b, 0x76, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x75, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0xec, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6b, 0x76, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x69, 0x66, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x13, 0x63, 0x61, 0x63, 0x68, 0x65, 0x49, 0x66, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x69, 0x73, 0x54, 0x72, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x9c, 0x04, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x3e, 0xda,
	0xde, 0x1f, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x63, 0x61, 0x70, 0x2f, 0x6b, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x76, 0x72, 0x70, 0x63, 0x70, 0x62,
	0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e,
	0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x37, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x76, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x65, 0x78, 0x65,
	0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3e, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5f, 0x76, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x76, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x56, 0x32, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x56, 0x32, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x5f, 0x68, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x61, 0x73,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x5f,
	0x62, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x63, 0x61, 0x6e, 0x42, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x05,
	0x73, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x05, 0x73, 0x70, 0x61, 0x6e,
	0x73, 0x22, 0x90, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x76, 0x72, 0x70, 0x63, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56,
	0x65, 0x72, 0x22, 0xf2, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x3e, 0xda, 0xde, 0x1f, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x70, 0x2f, 0x6b, 0x76, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x65, 0x78, 0x65,
	0x63, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6b, 0x76, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x5a, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x2e, 0x74,
	0x69, 0x6b, 0x76, 0x2e, 0x6b, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x61, 0x70, 0x2f,
	0x6b, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x69,
	0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0xc8, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xd8,
	0xa8, 0x08, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coprocessor_proto_rawDescOnce sync.Once
	file_coprocessor_proto_rawDescData = file_coprocessor_proto_rawDesc
)

func file_coprocessor_proto_rawDescGZIP() []byte {
	file_coprocessor_proto_rawDescOnce.Do(func() {
		file_coprocessor_proto_rawDescData = protoimpl.X.CompressGZIP(file_coprocessor_proto_rawDescData)
	})
	return file_coprocessor_proto_rawDescData
}

var file_coprocessor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_coprocessor_proto_goTypes = []interface{}{
	(*KeyRange)(nil),              // 0: coprocessor.KeyRange
	(*Request)(nil),               // 1: coprocessor.Request
	(*Response)(nil),              // 2: coprocessor.Response
	(*RegionInfo)(nil),            // 3: coprocessor.RegionInfo
	(*BatchRequest)(nil),          // 4: coprocessor.BatchRequest
	(*BatchResponse)(nil),         // 5: coprocessor.BatchResponse
	(*kvrpcpb.Context)(nil),       // 6: kvrpcpb.Context
	(*errorpb.Error)(nil),         // 7: errorpb.Error
	(*kvrpcpb.LockInfo)(nil),      // 8: kvrpcpb.LockInfo
	(*kvrpcpb.ExecDetails)(nil),   // 9: kvrpcpb.ExecDetails
	(*kvrpcpb.ExecDetailsV2)(nil), // 10: kvrpcpb.ExecDetailsV2
	(*span.SpanSet)(nil),          // 11: span.SpanSet
	(*metapb.RegionEpoch)(nil),    // 12: metapb.RegionEpoch
	(*metapb.Region)(nil),         // 13: metapb.Region
}
var file_coprocessor_proto_depIdxs = []int32{
	6,  // 0: coprocessor.Request.context:type_name -> kvrpcpb.Context
	0,  // 1: coprocessor.Request.ranges:type_name -> coprocessor.KeyRange
	7,  // 2: coprocessor.Response.region_error:type_name -> errorpb.Error
	8,  // 3: coprocessor.Response.locked:type_name -> kvrpcpb.LockInfo
	0,  // 4: coprocessor.Response.range:type_name -> coprocessor.KeyRange
	9,  // 5: coprocessor.Response.exec_details:type_name -> kvrpcpb.ExecDetails
	10, // 6: coprocessor.Response.exec_details_v2:type_name -> kvrpcpb.ExecDetailsV2
	11, // 7: coprocessor.Response.spans:type_name -> span.SpanSet
	12, // 8: coprocessor.RegionInfo.region_epoch:type_name -> metapb.RegionEpoch
	0,  // 9: coprocessor.RegionInfo.ranges:type_name -> coprocessor.KeyRange
	6,  // 10: coprocessor.BatchRequest.context:type_name -> kvrpcpb.Context
	3,  // 11: coprocessor.BatchRequest.regions:type_name -> coprocessor.RegionInfo
	9,  // 12: coprocessor.BatchResponse.exec_details:type_name -> kvrpcpb.ExecDetails
	13, // 13: coprocessor.BatchResponse.retry_regions:type_name -> metapb.Region
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

// func init() { file_coprocessor_proto_init() }
func file_coprocessor_proto_init() {
	if File_coprocessor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coprocessor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRange); i {
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
		file_coprocessor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_coprocessor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_coprocessor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionInfo); i {
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
		file_coprocessor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchRequest); i {
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
		file_coprocessor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchResponse); i {
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
			RawDescriptor: file_coprocessor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coprocessor_proto_goTypes,
		DependencyIndexes: file_coprocessor_proto_depIdxs,
		MessageInfos:      file_coprocessor_proto_msgTypes,
	}.Build()
	File_coprocessor_proto = out.File
	file_coprocessor_proto_rawDesc = nil
	file_coprocessor_proto_goTypes = nil
	file_coprocessor_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.1.2. DO NOT EDIT.
