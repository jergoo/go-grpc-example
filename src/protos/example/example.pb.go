// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: protos/example/example.proto

package example

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

// Status 枚举
type Status int32

const (
	Status_OK   Status = 0
	Status_FAIL Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_example_example_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_protos_example_example_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_protos_example_example_proto_rawDescGZIP(), []int{0}
}

// Request 请求结构
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *string `protobuf:"bytes,1,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_protos_example_example_proto_msgTypes[0]
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
	return file_protos_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

// Response 响应结构
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valuee string `protobuf:"bytes,1,opt,name=valuee,proto3" json:"valuee,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protos_example_example_proto_msgTypes[1]
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
	return file_protos_example_example_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetValuee() string {
	if x != nil {
		return x.Valuee
	}
	return ""
}

// Msg message 数据类型示例
type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I32     int32             `protobuf:"varint,1,opt,name=i32,proto3" json:"i32,omitempty"`
	I64     int64             `protobuf:"varint,2,opt,name=i64,proto3" json:"i64,omitempty"`
	F32     float32           `protobuf:"fixed32,3,opt,name=f32,proto3" json:"f32,omitempty"`
	F64     float64           `protobuf:"fixed64,4,opt,name=f64,proto3" json:"f64,omitempty"`
	Str     string            `protobuf:"bytes,5,opt,name=str,proto3" json:"str,omitempty"`
	Boolean bool              `protobuf:"varint,6,opt,name=boolean,proto3" json:"boolean,omitempty"`
	ByteArr []byte            `protobuf:"bytes,7,opt,name=byteArr,proto3" json:"byteArr,omitempty"`
	Dict    map[string]string `protobuf:"bytes,8,rep,name=dict,proto3" json:"dict,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status  Status            `protobuf:"varint,9,opt,name=status,proto3,enum=example.Status" json:"status,omitempty"`
	EmbMsg  *EmbMsg           `protobuf:"bytes,10,opt,name=embMsg,proto3" json:"embMsg,omitempty"`
	IntArr  []int64           `protobuf:"varint,11,rep,packed,name=intArr,proto3" json:"intArr,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_example_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_protos_example_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_protos_example_example_proto_rawDescGZIP(), []int{2}
}

func (x *Msg) GetI32() int32 {
	if x != nil {
		return x.I32
	}
	return 0
}

func (x *Msg) GetI64() int64 {
	if x != nil {
		return x.I64
	}
	return 0
}

func (x *Msg) GetF32() float32 {
	if x != nil {
		return x.F32
	}
	return 0
}

func (x *Msg) GetF64() float64 {
	if x != nil {
		return x.F64
	}
	return 0
}

func (x *Msg) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Msg) GetBoolean() bool {
	if x != nil {
		return x.Boolean
	}
	return false
}

func (x *Msg) GetByteArr() []byte {
	if x != nil {
		return x.ByteArr
	}
	return nil
}

func (x *Msg) GetDict() map[string]string {
	if x != nil {
		return x.Dict
	}
	return nil
}

func (x *Msg) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *Msg) GetEmbMsg() *EmbMsg {
	if x != nil {
		return x.EmbMsg
	}
	return nil
}

func (x *Msg) GetIntArr() []int64 {
	if x != nil {
		return x.IntArr
	}
	return nil
}

type EmbMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EmbMsg) Reset() {
	*x = EmbMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_example_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbMsg) ProtoMessage() {}

func (x *EmbMsg) ProtoReflect() protoreflect.Message {
	mi := &file_protos_example_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbMsg.ProtoReflect.Descriptor instead.
func (*EmbMsg) Descriptor() ([]byte, []int) {
	return file_protos_example_example_proto_rawDescGZIP(), []int{3}
}

func (x *EmbMsg) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_protos_example_example_proto protoreflect.FileDescriptor

var file_protos_example_example_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x22, 0xe2, 0x02, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x33, 0x32, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x66, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x36, 0x34,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x79, 0x74, 0x65, 0x41,
	0x72, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x79, 0x74, 0x65, 0x41, 0x72,
	0x72, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x69, 0x63, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x2e, 0x44, 0x69,
	0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x69, 0x63, 0x74, 0x12, 0x27, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x6d, 0x62, 0x4d, 0x73, 0x67,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x45, 0x6d, 0x62, 0x4d, 0x73, 0x67, 0x52, 0x06, 0x65, 0x6d, 0x62, 0x4d, 0x73, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x06, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x69, 0x63, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x1e, 0x0a, 0x06, 0x45, 0x6d, 0x62, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2a, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x32, 0xe2, 0x01, 0x0a,
	0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x33, 0x0a, 0x08,
	0x42, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_example_example_proto_rawDescOnce sync.Once
	file_protos_example_example_proto_rawDescData = file_protos_example_example_proto_rawDesc
)

func file_protos_example_example_proto_rawDescGZIP() []byte {
	file_protos_example_example_proto_rawDescOnce.Do(func() {
		file_protos_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_example_example_proto_rawDescData)
	})
	return file_protos_example_example_proto_rawDescData
}

var file_protos_example_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_example_example_proto_goTypes = []interface{}{
	(Status)(0),      // 0: example.Status
	(*Request)(nil),  // 1: example.Request
	(*Response)(nil), // 2: example.Response
	(*Msg)(nil),      // 3: example.Msg
	(*EmbMsg)(nil),   // 4: example.EmbMsg
	nil,              // 5: example.Msg.DictEntry
}
var file_protos_example_example_proto_depIdxs = []int32{
	5, // 0: example.Msg.dict:type_name -> example.Msg.DictEntry
	0, // 1: example.Msg.status:type_name -> example.Status
	4, // 2: example.Msg.embMsg:type_name -> example.EmbMsg
	1, // 3: example.ExampleService.Single:input_type -> example.Request
	1, // 4: example.ExampleService.ServerStream:input_type -> example.Request
	1, // 5: example.ExampleService.ClientStream:input_type -> example.Request
	1, // 6: example.ExampleService.BiStream:input_type -> example.Request
	2, // 7: example.ExampleService.Single:output_type -> example.Response
	2, // 8: example.ExampleService.ServerStream:output_type -> example.Response
	2, // 9: example.ExampleService.ClientStream:output_type -> example.Response
	2, // 10: example.ExampleService.BiStream:output_type -> example.Response
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_example_example_proto_init() }
func file_protos_example_example_proto_init() {
	if File_protos_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_example_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_protos_example_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbMsg); i {
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
	file_protos_example_example_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_example_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_example_example_proto_goTypes,
		DependencyIndexes: file_protos_example_example_proto_depIdxs,
		EnumInfos:         file_protos_example_example_proto_enumTypes,
		MessageInfos:      file_protos_example_example_proto_msgTypes,
	}.Build()
	File_protos_example_example_proto = out.File
	file_protos_example_example_proto_rawDesc = nil
	file_protos_example_example_proto_goTypes = nil
	file_protos_example_example_proto_depIdxs = nil
}
