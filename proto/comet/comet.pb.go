// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: proto/comet/comet.proto

package comet

import (
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content   []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UserId    int32  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_comet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_comet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_comet_comet_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PushMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushMessageReq) Reset() {
	*x = PushMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_comet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageReq) ProtoMessage() {}

func (x *PushMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_comet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageReq.ProtoReflect.Descriptor instead.
func (*PushMessageReq) Descriptor() ([]byte, []int) {
	return file_proto_comet_comet_proto_rawDescGZIP(), []int{1}
}

func (x *PushMessageReq) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type PushMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushMessageResp) Reset() {
	*x = PushMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_comet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageResp) ProtoMessage() {}

func (x *PushMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_comet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageResp.ProtoReflect.Descriptor instead.
func (*PushMessageResp) Descriptor() ([]byte, []int) {
	return file_proto_comet_comet_proto_rawDescGZIP(), []int{2}
}

var File_proto_comet_comet_proto protoreflect.FileDescriptor

var file_proto_comet_comet_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x63, 0x6f,
	0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x6d, 0x65, 0x74,
	0x22, 0x6e, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x3a, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x11, 0x0a, 0x0f,
	0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32,
	0x45, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6b, 0x64, 0x73, 0x2f, 0x6b, 0x69, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x3b, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_comet_comet_proto_rawDescOnce sync.Once
	file_proto_comet_comet_proto_rawDescData = file_proto_comet_comet_proto_rawDesc
)

func file_proto_comet_comet_proto_rawDescGZIP() []byte {
	file_proto_comet_comet_proto_rawDescOnce.Do(func() {
		file_proto_comet_comet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comet_comet_proto_rawDescData)
	})
	return file_proto_comet_comet_proto_rawDescData
}

var file_proto_comet_comet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_comet_comet_proto_goTypes = []interface{}{
	(*Message)(nil),         // 0: comet.Message
	(*PushMessageReq)(nil),  // 1: comet.PushMessageReq
	(*PushMessageResp)(nil), // 2: comet.PushMessageResp
}
var file_proto_comet_comet_proto_depIdxs = []int32{
	0, // 0: comet.PushMessageReq.message:type_name -> comet.Message
	1, // 1: comet.Comet.PushMessage:input_type -> comet.PushMessageReq
	2, // 2: comet.Comet.PushMessage:output_type -> comet.PushMessageResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_comet_comet_proto_init() }
func file_proto_comet_comet_proto_init() {
	if File_proto_comet_comet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comet_comet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_comet_comet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessageReq); i {
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
		file_proto_comet_comet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessageResp); i {
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
			RawDescriptor: file_proto_comet_comet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_comet_comet_proto_goTypes,
		DependencyIndexes: file_proto_comet_comet_proto_depIdxs,
		MessageInfos:      file_proto_comet_comet_proto_msgTypes,
	}.Build()
	File_proto_comet_comet_proto = out.File
	file_proto_comet_comet_proto_rawDesc = nil
	file_proto_comet_comet_proto_goTypes = nil
	file_proto_comet_comet_proto_depIdxs = nil
}