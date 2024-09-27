// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: protos/metrics/metrics.proto

package metrics

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Counters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Total   int64  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Sent    int64  `protobuf:"varint,3,opt,name=sent,proto3" json:"sent,omitempty"`
	Success int64  `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
	Failed  int64  `protobuf:"varint,5,opt,name=failed,proto3" json:"failed,omitempty"`
	Timeout int64  `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Invalid int64  `protobuf:"varint,7,opt,name=invalid,proto3" json:"invalid,omitempty"`
	Latency int64  `protobuf:"varint,8,opt,name=latency,proto3" json:"latency,omitempty"`
}

func (x *Counters) Reset() {
	*x = Counters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_metrics_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counters) ProtoMessage() {}

func (x *Counters) ProtoReflect() protoreflect.Message {
	mi := &file_protos_metrics_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counters.ProtoReflect.Descriptor instead.
func (*Counters) Descriptor() ([]byte, []int) {
	return file_protos_metrics_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Counters) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Counters) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Counters) GetSent() int64 {
	if x != nil {
		return x.Sent
	}
	return 0
}

func (x *Counters) GetSuccess() int64 {
	if x != nil {
		return x.Success
	}
	return 0
}

func (x *Counters) GetFailed() int64 {
	if x != nil {
		return x.Failed
	}
	return 0
}

func (x *Counters) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Counters) GetInvalid() int64 {
	if x != nil {
		return x.Invalid
	}
	return 0
}

func (x *Counters) GetLatency() int64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_metrics_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_protos_metrics_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_protos_metrics_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_protos_metrics_metrics_proto protoreflect.FileDescriptor

var file_protos_metrics_metrics_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x1a,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x3f, 0x0a, 0x07, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x0f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6d, 0x65, 0x61, 0x75,
	0x78, 0x2f, 0x62, 0x72, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_metrics_metrics_proto_rawDescOnce sync.Once
	file_protos_metrics_metrics_proto_rawDescData = file_protos_metrics_metrics_proto_rawDesc
)

func file_protos_metrics_metrics_proto_rawDescGZIP() []byte {
	file_protos_metrics_metrics_proto_rawDescOnce.Do(func() {
		file_protos_metrics_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_metrics_metrics_proto_rawDescData)
	})
	return file_protos_metrics_metrics_proto_rawDescData
}

var (
	file_protos_metrics_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_protos_metrics_metrics_proto_goTypes  = []any{
		(*Counters)(nil), // 0: metrics.Counters
		(*Result)(nil),   // 1: metrics.Result
	}
)
var file_protos_metrics_metrics_proto_depIdxs = []int32{
	0, // 0: metrics.Metrics.sendCounters:input_type -> metrics.Counters
	1, // 1: metrics.Metrics.sendCounters:output_type -> metrics.Result
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_metrics_metrics_proto_init() }
func file_protos_metrics_metrics_proto_init() {
	if File_protos_metrics_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_metrics_metrics_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Counters); i {
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
		file_protos_metrics_metrics_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_protos_metrics_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_metrics_metrics_proto_goTypes,
		DependencyIndexes: file_protos_metrics_metrics_proto_depIdxs,
		MessageInfos:      file_protos_metrics_metrics_proto_msgTypes,
	}.Build()
	File_protos_metrics_metrics_proto = out.File
	file_protos_metrics_metrics_proto_rawDesc = nil
	file_protos_metrics_metrics_proto_goTypes = nil
	file_protos_metrics_metrics_proto_depIdxs = nil
}
