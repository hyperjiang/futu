// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: Verification.proto

package verification

import (
	_ "github.com/hyperjiang/futu/pb/common"
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

type VerificationType int32

const (
	VerificationType_VerificationType_Unknow  VerificationType = 0 //未知操作
	VerificationType_VerificationType_Picture VerificationType = 1 // 图形验证码
	VerificationType_VerificationType_Phone   VerificationType = 2 // 手机验证码
)

// Enum value maps for VerificationType.
var (
	VerificationType_name = map[int32]string{
		0: "VerificationType_Unknow",
		1: "VerificationType_Picture",
		2: "VerificationType_Phone",
	}
	VerificationType_value = map[string]int32{
		"VerificationType_Unknow":  0,
		"VerificationType_Picture": 1,
		"VerificationType_Phone":   2,
	}
)

func (x VerificationType) Enum() *VerificationType {
	p := new(VerificationType)
	*p = x
	return p
}

func (x VerificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_Verification_proto_enumTypes[0].Descriptor()
}

func (VerificationType) Type() protoreflect.EnumType {
	return &file_Verification_proto_enumTypes[0]
}

func (x VerificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *VerificationType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = VerificationType(num)
	return nil
}

// Deprecated: Use VerificationType.Descriptor instead.
func (VerificationType) EnumDescriptor() ([]byte, []int) {
	return file_Verification_proto_rawDescGZIP(), []int{0}
}

type VerificationOp int32

const (
	VerificationOp_VerificationOp_Unknow        VerificationOp = 0 //未知操作
	VerificationOp_VerificationOp_Request       VerificationOp = 1 //请求验证码
	VerificationOp_VerificationOp_InputAndLogin VerificationOp = 2 //输入验证码并继续登录操作
)

// Enum value maps for VerificationOp.
var (
	VerificationOp_name = map[int32]string{
		0: "VerificationOp_Unknow",
		1: "VerificationOp_Request",
		2: "VerificationOp_InputAndLogin",
	}
	VerificationOp_value = map[string]int32{
		"VerificationOp_Unknow":        0,
		"VerificationOp_Request":       1,
		"VerificationOp_InputAndLogin": 2,
	}
)

func (x VerificationOp) Enum() *VerificationOp {
	p := new(VerificationOp)
	*p = x
	return p
}

func (x VerificationOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerificationOp) Descriptor() protoreflect.EnumDescriptor {
	return file_Verification_proto_enumTypes[1].Descriptor()
}

func (VerificationOp) Type() protoreflect.EnumType {
	return &file_Verification_proto_enumTypes[1]
}

func (x VerificationOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *VerificationOp) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = VerificationOp(num)
	return nil
}

// Deprecated: Use VerificationOp.Descriptor instead.
func (VerificationOp) EnumDescriptor() ([]byte, []int) {
	return file_Verification_proto_rawDescGZIP(), []int{1}
}

// 图形验证码下载之后会将其存至固定路径，请到该路径下查看验证码
// Windows平台：%appdata%/com.futunn.FutuOpenD/F3CNN/PicVerifyCode.png
// 非Windows平台：~/.com.futunn.FutuOpenD/F3CNN/PicVerifyCode.png
// 注意：只有最后一次请求验证码会生效，重复请求只有最后一次的验证码有效
type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *int32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"` //验证码类型, VerificationType
	Op   *int32  `protobuf:"varint,2,req,name=op" json:"op,omitempty"`     //操作, VerificationOp
	Code *string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`  //验证码，请求验证码时忽略该字段，输入时必填
}

func (x *C2S) Reset() {
	*x = C2S{}
	mi := &file_Verification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Verification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Verification_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *C2S) GetOp() int32 {
	if x != nil && x.Op != nil {
		return *x.Op
	}
	return 0
}

func (x *C2S) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2C) Reset() {
	*x = S2C{}
	mi := &file_Verification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Verification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Verification_proto_rawDescGZIP(), []int{1}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_Verification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Verification_proto_msgTypes[2]
	if x != nil {
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
	return file_Verification_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //返回结果，参见Common.RetType的枚举定义
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`             //返回结果描述
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`          //错误码，客户端一般通过retType和retMsg来判断结果和详情，errCode只做日志记录，仅在个别协议失败时对账用
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_Verification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Verification_proto_msgTypes[3]
	if x != nil {
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
	return file_Verification_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Verification_proto protoreflect.FileDescriptor

var file_Verification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x05, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x22, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x32,
	0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0x69, 0x0a, 0x10, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x10, 0x02, 0x2a, 0x69, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x12, 0x19, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x01, 0x12, 0x20,
	0x0a, 0x1c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x5f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x02,
	0x42, 0x41, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x66,
	0x75, 0x74, 0x75, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e,
}

var (
	file_Verification_proto_rawDescOnce sync.Once
	file_Verification_proto_rawDescData = file_Verification_proto_rawDesc
)

func file_Verification_proto_rawDescGZIP() []byte {
	file_Verification_proto_rawDescOnce.Do(func() {
		file_Verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_Verification_proto_rawDescData)
	})
	return file_Verification_proto_rawDescData
}

var file_Verification_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Verification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Verification_proto_goTypes = []any{
	(VerificationType)(0), // 0: Verification.VerificationType
	(VerificationOp)(0),   // 1: Verification.VerificationOp
	(*C2S)(nil),           // 2: Verification.C2S
	(*S2C)(nil),           // 3: Verification.S2C
	(*Request)(nil),       // 4: Verification.Request
	(*Response)(nil),      // 5: Verification.Response
}
var file_Verification_proto_depIdxs = []int32{
	2, // 0: Verification.Request.c2s:type_name -> Verification.C2S
	3, // 1: Verification.Response.s2c:type_name -> Verification.S2C
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Verification_proto_init() }
func file_Verification_proto_init() {
	if File_Verification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Verification_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Verification_proto_goTypes,
		DependencyIndexes: file_Verification_proto_depIdxs,
		EnumInfos:         file_Verification_proto_enumTypes,
		MessageInfos:      file_Verification_proto_msgTypes,
	}.Build()
	File_Verification_proto = out.File
	file_Verification_proto_rawDesc = nil
	file_Verification_proto_goTypes = nil
	file_Verification_proto_depIdxs = nil
}
