// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: Common.proto

package common

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

// 返回结果
type RetType int32

const (
	RetType_RetType_Succeed    RetType = 0    //成功
	RetType_RetType_Failed     RetType = -1   //失败
	RetType_RetType_TimeOut    RetType = -100 //超时
	RetType_RetType_DisConnect RetType = -200 //连接断开
	RetType_RetType_Unknown    RetType = -400 //未知结果
	RetType_RetType_Invalid    RetType = -500 //包内容非法
)

// Enum value maps for RetType.
var (
	RetType_name = map[int32]string{
		0:    "RetType_Succeed",
		-1:   "RetType_Failed",
		-100: "RetType_TimeOut",
		-200: "RetType_DisConnect",
		-400: "RetType_Unknown",
		-500: "RetType_Invalid",
	}
	RetType_value = map[string]int32{
		"RetType_Succeed":    0,
		"RetType_Failed":     -1,
		"RetType_TimeOut":    -100,
		"RetType_DisConnect": -200,
		"RetType_Unknown":    -400,
		"RetType_Invalid":    -500,
	}
)

func (x RetType) Enum() *RetType {
	p := new(RetType)
	*p = x
	return p
}

func (x RetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetType) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[0].Descriptor()
}

func (RetType) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[0]
}

func (x RetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RetType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RetType(num)
	return nil
}

// Deprecated: Use RetType.Descriptor instead.
func (RetType) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{0}
}

// 包加密算法
type PacketEncAlgo int32

const (
	PacketEncAlgo_PacketEncAlgo_FTAES_ECB PacketEncAlgo = 0  //富途修改过的AES的ECB加密模式
	PacketEncAlgo_PacketEncAlgo_None      PacketEncAlgo = -1 //不加密
	PacketEncAlgo_PacketEncAlgo_AES_ECB   PacketEncAlgo = 1  //标准的AES的ECB加密模式
	PacketEncAlgo_PacketEncAlgo_AES_CBC   PacketEncAlgo = 2  //标准的AES的CBC加密模式
)

// Enum value maps for PacketEncAlgo.
var (
	PacketEncAlgo_name = map[int32]string{
		0:  "PacketEncAlgo_FTAES_ECB",
		-1: "PacketEncAlgo_None",
		1:  "PacketEncAlgo_AES_ECB",
		2:  "PacketEncAlgo_AES_CBC",
	}
	PacketEncAlgo_value = map[string]int32{
		"PacketEncAlgo_FTAES_ECB": 0,
		"PacketEncAlgo_None":      -1,
		"PacketEncAlgo_AES_ECB":   1,
		"PacketEncAlgo_AES_CBC":   2,
	}
)

func (x PacketEncAlgo) Enum() *PacketEncAlgo {
	p := new(PacketEncAlgo)
	*p = x
	return p
}

func (x PacketEncAlgo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketEncAlgo) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[1].Descriptor()
}

func (PacketEncAlgo) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[1]
}

func (x PacketEncAlgo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *PacketEncAlgo) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = PacketEncAlgo(num)
	return nil
}

// Deprecated: Use PacketEncAlgo.Descriptor instead.
func (PacketEncAlgo) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{1}
}

// 协议格式，请求协议在请求头中指定，推送协议在Init时指定
type ProtoFmt int32

const (
	ProtoFmt_ProtoFmt_Protobuf ProtoFmt = 0 //Google Protobuf格式
	ProtoFmt_ProtoFmt_Json     ProtoFmt = 1 //Json格式
)

// Enum value maps for ProtoFmt.
var (
	ProtoFmt_name = map[int32]string{
		0: "ProtoFmt_Protobuf",
		1: "ProtoFmt_Json",
	}
	ProtoFmt_value = map[string]int32{
		"ProtoFmt_Protobuf": 0,
		"ProtoFmt_Json":     1,
	}
)

func (x ProtoFmt) Enum() *ProtoFmt {
	p := new(ProtoFmt)
	*p = x
	return p
}

func (x ProtoFmt) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoFmt) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[2].Descriptor()
}

func (ProtoFmt) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[2]
}

func (x ProtoFmt) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProtoFmt) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProtoFmt(num)
	return nil
}

// Deprecated: Use ProtoFmt.Descriptor instead.
func (ProtoFmt) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{2}
}

// 用户注册归属地
type UserAttribution int32

const (
	UserAttribution_UserAttribution_Unknown UserAttribution = 0 //
	UserAttribution_UserAttribution_NN      UserAttribution = 1 //大陆
	UserAttribution_UserAttribution_MM      UserAttribution = 2 //MooMoo
	UserAttribution_UserAttribution_SG      UserAttribution = 3 //新加坡
	UserAttribution_UserAttribution_AU      UserAttribution = 4 //澳洲
	UserAttribution_UserAttribution_JP      UserAttribution = 5 //日本
	UserAttribution_UserAttribution_HK      UserAttribution = 6 //香港
)

// Enum value maps for UserAttribution.
var (
	UserAttribution_name = map[int32]string{
		0: "UserAttribution_Unknown",
		1: "UserAttribution_NN",
		2: "UserAttribution_MM",
		3: "UserAttribution_SG",
		4: "UserAttribution_AU",
		5: "UserAttribution_JP",
		6: "UserAttribution_HK",
	}
	UserAttribution_value = map[string]int32{
		"UserAttribution_Unknown": 0,
		"UserAttribution_NN":      1,
		"UserAttribution_MM":      2,
		"UserAttribution_SG":      3,
		"UserAttribution_AU":      4,
		"UserAttribution_JP":      5,
		"UserAttribution_HK":      6,
	}
)

func (x UserAttribution) Enum() *UserAttribution {
	p := new(UserAttribution)
	*p = x
	return p
}

func (x UserAttribution) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserAttribution) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[3].Descriptor()
}

func (UserAttribution) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[3]
}

func (x UserAttribution) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UserAttribution) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UserAttribution(num)
	return nil
}

// Deprecated: Use UserAttribution.Descriptor instead.
func (UserAttribution) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{3}
}

type ProgramStatusType int32

const (
	ProgramStatusType_ProgramStatusType_None                 ProgramStatusType = 0
	ProgramStatusType_ProgramStatusType_Loaded               ProgramStatusType = 1  //已完成类似加载配置,启动服务器等操作,服务器启动之前的状态无需返回
	ProgramStatusType_ProgramStatusType_Loging               ProgramStatusType = 2  //登录中
	ProgramStatusType_ProgramStatusType_NeedPicVerifyCode    ProgramStatusType = 3  //需要图形验证码
	ProgramStatusType_ProgramStatusType_NeedPhoneVerifyCode  ProgramStatusType = 4  //需要手机验证码
	ProgramStatusType_ProgramStatusType_LoginFailed          ProgramStatusType = 5  //登录失败,详细原因在描述返回
	ProgramStatusType_ProgramStatusType_ForceUpdate          ProgramStatusType = 6  //客户端版本过低
	ProgramStatusType_ProgramStatusType_NessaryDataPreparing ProgramStatusType = 7  //正在拉取类似免责声明等一些必要信息
	ProgramStatusType_ProgramStatusType_NessaryDataMissing   ProgramStatusType = 8  //缺少必要信息
	ProgramStatusType_ProgramStatusType_UnAgreeDisclaimer    ProgramStatusType = 9  //未同意免责声明
	ProgramStatusType_ProgramStatusType_Ready                ProgramStatusType = 10 //可以接收业务协议收发,正常可用状态
	// OpenD登录后被强制退出登录，会导致连接全部断开,需要重连后才能得到以下该状态（并且需要在ui模式下）
	ProgramStatusType_ProgramStatusType_ForceLogout          ProgramStatusType = 11 //被强制退出登录,例如修改了登录密码,中途打开设备锁等,详细原因在描述返回
	ProgramStatusType_ProgramStatusType_DisclaimerPullFailed ProgramStatusType = 12 //拉取免责声明标志失败
)

// Enum value maps for ProgramStatusType.
var (
	ProgramStatusType_name = map[int32]string{
		0:  "ProgramStatusType_None",
		1:  "ProgramStatusType_Loaded",
		2:  "ProgramStatusType_Loging",
		3:  "ProgramStatusType_NeedPicVerifyCode",
		4:  "ProgramStatusType_NeedPhoneVerifyCode",
		5:  "ProgramStatusType_LoginFailed",
		6:  "ProgramStatusType_ForceUpdate",
		7:  "ProgramStatusType_NessaryDataPreparing",
		8:  "ProgramStatusType_NessaryDataMissing",
		9:  "ProgramStatusType_UnAgreeDisclaimer",
		10: "ProgramStatusType_Ready",
		11: "ProgramStatusType_ForceLogout",
		12: "ProgramStatusType_DisclaimerPullFailed",
	}
	ProgramStatusType_value = map[string]int32{
		"ProgramStatusType_None":                 0,
		"ProgramStatusType_Loaded":               1,
		"ProgramStatusType_Loging":               2,
		"ProgramStatusType_NeedPicVerifyCode":    3,
		"ProgramStatusType_NeedPhoneVerifyCode":  4,
		"ProgramStatusType_LoginFailed":          5,
		"ProgramStatusType_ForceUpdate":          6,
		"ProgramStatusType_NessaryDataPreparing": 7,
		"ProgramStatusType_NessaryDataMissing":   8,
		"ProgramStatusType_UnAgreeDisclaimer":    9,
		"ProgramStatusType_Ready":                10,
		"ProgramStatusType_ForceLogout":          11,
		"ProgramStatusType_DisclaimerPullFailed": 12,
	}
)

func (x ProgramStatusType) Enum() *ProgramStatusType {
	p := new(ProgramStatusType)
	*p = x
	return p
}

func (x ProgramStatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProgramStatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[4].Descriptor()
}

func (ProgramStatusType) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[4]
}

func (x ProgramStatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProgramStatusType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProgramStatusType(num)
	return nil
}

// Deprecated: Use ProgramStatusType.Descriptor instead.
func (ProgramStatusType) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{4}
}

// 包的唯一标识，用于回放攻击的识别和保护
type PacketID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnID   *uint64 `protobuf:"varint,1,req,name=connID" json:"connID,omitempty"`     //当前TCP连接的连接ID，一条连接的唯一标识，InitConnect协议会返回
	SerialNo *uint32 `protobuf:"varint,2,req,name=serialNo" json:"serialNo,omitempty"` //自增序列号
}

func (x *PacketID) Reset() {
	*x = PacketID{}
	mi := &file_Common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketID) ProtoMessage() {}

func (x *PacketID) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketID.ProtoReflect.Descriptor instead.
func (*PacketID) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{0}
}

func (x *PacketID) GetConnID() uint64 {
	if x != nil && x.ConnID != nil {
		return *x.ConnID
	}
	return 0
}

func (x *PacketID) GetSerialNo() uint32 {
	if x != nil && x.SerialNo != nil {
		return *x.SerialNo
	}
	return 0
}

type ProgramStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       *ProgramStatusType `protobuf:"varint,1,req,name=type,enum=Common.ProgramStatusType" json:"type,omitempty"` //当前状态
	StrExtDesc *string            `protobuf:"bytes,2,opt,name=strExtDesc" json:"strExtDesc,omitempty"`                    // 额外描述
}

func (x *ProgramStatus) Reset() {
	*x = ProgramStatus{}
	mi := &file_Common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProgramStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramStatus) ProtoMessage() {}

func (x *ProgramStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramStatus.ProtoReflect.Descriptor instead.
func (*ProgramStatus) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{1}
}

func (x *ProgramStatus) GetType() ProgramStatusType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ProgramStatusType_ProgramStatusType_None
}

func (x *ProgramStatus) GetStrExtDesc() string {
	if x != nil && x.StrExtDesc != nil {
		return *x.StrExtDesc
	}
	return ""
}

var File_Common_proto protoreflect.FileDescriptor

var file_Common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x22, 0x5e, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x45, 0x78, 0x74,
	0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x45,
	0x78, 0x74, 0x44, 0x65, 0x73, 0x63, 0x2a, 0xb6, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x0e, 0x52, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x10, 0x9c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x12, 0x1f, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x44, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0xb8, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0xf0, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x10, 0x8c, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x2a,
	0x83, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x41, 0x6c, 0x67,
	0x6f, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x41, 0x6c,
	0x67, 0x6f, 0x5f, 0x46, 0x54, 0x41, 0x45, 0x53, 0x5f, 0x45, 0x43, 0x42, 0x10, 0x00, 0x12, 0x1f,
	0x0a, 0x12, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x41, 0x6c, 0x67, 0x6f, 0x5f,
	0x4e, 0x6f, 0x6e, 0x65, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x41, 0x6c, 0x67, 0x6f,
	0x5f, 0x41, 0x45, 0x53, 0x5f, 0x45, 0x43, 0x42, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x41, 0x6c, 0x67, 0x6f, 0x5f, 0x41, 0x45, 0x53, 0x5f,
	0x43, 0x42, 0x43, 0x10, 0x02, 0x2a, 0x34, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x6d,
	0x74, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x6d, 0x74, 0x5f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x46, 0x6d, 0x74, 0x5f, 0x4a, 0x73, 0x6f, 0x6e, 0x10, 0x01, 0x2a, 0xbe, 0x01, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x4e, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x4d, 0x4d, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x53, 0x47, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41, 0x55, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x4a, 0x50, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x48, 0x4b, 0x10, 0x06, 0x2a, 0xf0, 0x03, 0x0a,
	0x11, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x1c,
	0x0a, 0x18, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x4e, 0x65, 0x65, 0x64, 0x50, 0x69, 0x63, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x65, 0x65, 0x64, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x04, 0x12, 0x21,
	0x0a, 0x1d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10,
	0x05, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x10, 0x06, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x65, 0x73, 0x73, 0x61, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x07,
	0x12, 0x28, 0x0a, 0x24, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x65, 0x73, 0x73, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x08, 0x12, 0x27, 0x0a, 0x23, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x55, 0x6e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65,
	0x72, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x0a,
	0x12, 0x21, 0x0a, 0x1d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x10, 0x0b, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x44, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x65, 0x72, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x0c, 0x42,
	0x3b, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x66, 0x75,
	0x74, 0x75, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
}

var (
	file_Common_proto_rawDescOnce sync.Once
	file_Common_proto_rawDescData = file_Common_proto_rawDesc
)

func file_Common_proto_rawDescGZIP() []byte {
	file_Common_proto_rawDescOnce.Do(func() {
		file_Common_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_proto_rawDescData)
	})
	return file_Common_proto_rawDescData
}

var file_Common_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_Common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Common_proto_goTypes = []any{
	(RetType)(0),           // 0: Common.RetType
	(PacketEncAlgo)(0),     // 1: Common.PacketEncAlgo
	(ProtoFmt)(0),          // 2: Common.ProtoFmt
	(UserAttribution)(0),   // 3: Common.UserAttribution
	(ProgramStatusType)(0), // 4: Common.ProgramStatusType
	(*PacketID)(nil),       // 5: Common.PacketID
	(*ProgramStatus)(nil),  // 6: Common.ProgramStatus
}
var file_Common_proto_depIdxs = []int32{
	4, // 0: Common.ProgramStatus.type:type_name -> Common.ProgramStatusType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Common_proto_init() }
func file_Common_proto_init() {
	if File_Common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Common_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_proto_goTypes,
		DependencyIndexes: file_Common_proto_depIdxs,
		EnumInfos:         file_Common_proto_enumTypes,
		MessageInfos:      file_Common_proto_msgTypes,
	}.Build()
	File_Common_proto = out.File
	file_Common_proto_rawDesc = nil
	file_Common_proto_goTypes = nil
	file_Common_proto_depIdxs = nil
}
