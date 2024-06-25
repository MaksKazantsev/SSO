// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: messages.proto

package pkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token           string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	ReceiverID      string `protobuf:"bytes,2,opt,name=ReceiverID,proto3" json:"ReceiverID,omitempty"`
	SenderID        string `protobuf:"bytes,3,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	ChatID          string `protobuf:"bytes,4,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	Value           string `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
	ReceiverOffline bool   `protobuf:"varint,6,opt,name=ReceiverOffline,proto3" json:"ReceiverOffline,omitempty"`
}

func (x *CreateMessageReq) Reset() {
	*x = CreateMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageReq) ProtoMessage() {}

func (x *CreateMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageReq.ProtoReflect.Descriptor instead.
func (*CreateMessageReq) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateMessageReq) GetReceiverID() string {
	if x != nil {
		return x.ReceiverID
	}
	return ""
}

func (x *CreateMessageReq) GetSenderID() string {
	if x != nil {
		return x.SenderID
	}
	return ""
}

func (x *CreateMessageReq) GetChatID() string {
	if x != nil {
		return x.ChatID
	}
	return ""
}

func (x *CreateMessageReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CreateMessageReq) GetReceiverOffline() bool {
	if x != nil {
		return x.ReceiverOffline
	}
	return false
}

type DeleteMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID string `protobuf:"bytes,1,opt,name=MessageID,proto3" json:"MessageID,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *DeleteMessageReq) Reset() {
	*x = DeleteMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageReq) ProtoMessage() {}

func (x *DeleteMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageReq.ProtoReflect.Descriptor instead.
func (*DeleteMessageReq) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteMessageReq) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *DeleteMessageReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x46, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x82, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x11,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x6b, 0x73, 0x4b, 0x61, 0x7a, 0x61,
	0x6e, 0x74, 0x73, 0x65, 0x76, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messages_proto_goTypes = []interface{}{
	(*CreateMessageReq)(nil), // 0: CreateMessageReq
	(*DeleteMessageReq)(nil), // 1: DeleteMessageReq
	(*emptypb.Empty)(nil),    // 2: google.protobuf.Empty
}
var file_messages_proto_depIdxs = []int32{
	0, // 0: Messages.CreateMessage:input_type -> CreateMessageReq
	1, // 1: Messages.DeleteMessage:input_type -> DeleteMessageReq
	2, // 2: Messages.CreateMessage:output_type -> google.protobuf.Empty
	2, // 3: Messages.DeleteMessage:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageReq); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageReq); i {
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
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
