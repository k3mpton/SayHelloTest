// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/auth.proto

package k3me

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PasswordHash  []byte                 `protobuf:"bytes,3,opt,name=passwordHash,proto3" json:"passwordHash,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegRequest) Reset() {
	*x = RegRequest{}
	mi := &file_proto_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegRequest) ProtoMessage() {}

func (x *RegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegRequest.ProtoReflect.Descriptor instead.
func (*RegRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegRequest) GetPasswordHash() []byte {
	if x != nil {
		return x.PasswordHash
	}
	return nil
}

func (x *RegRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type RegResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	HelloName     string                 `protobuf:"bytes,2,opt,name=HelloName,proto3" json:"HelloName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegResponse) Reset() {
	*x = RegResponse{}
	mi := &file_proto_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegResponse) ProtoMessage() {}

func (x *RegResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegResponse.ProtoReflect.Descriptor instead.
func (*RegResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegResponse) GetHelloName() string {
	if x != nil {
		return x.HelloName
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LogUserId     int64                  `protobuf:"varint,1,opt,name=LogUserId,proto3" json:"LogUserId,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_proto_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetLogUserId() int64 {
	if x != nil {
		return x.LogUserId
	}
	return 0
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LUserId       int64                  `protobuf:"varint,1,opt,name=LUserId,proto3" json:"LUserId,omitempty"`
	Svosrasheniem string                 `protobuf:"bytes,2,opt,name=Svosrasheniem,proto3" json:"Svosrasheniem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_proto_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetLUserId() int64 {
	if x != nil {
		return x.LUserId
	}
	return 0
}

func (x *LoginResponse) GetSvosrasheniem() string {
	if x != nil {
		return x.Svosrasheniem
	}
	return ""
}

var File_proto_auth_proto protoreflect.FileDescriptor

const file_proto_auth_proto_rawDesc = "" +
	"\n" +
	"\x10proto/auth.proto\x12\x04auth\"r\n" +
	"\n" +
	"RegRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\"\n" +
	"\fpasswordHash\x18\x03 \x01(\fR\fpasswordHash\x12\x16\n" +
	"\x06gender\x18\x04 \x01(\tR\x06gender\"C\n" +
	"\vRegResponse\x12\x16\n" +
	"\x06UserId\x18\x01 \x01(\x03R\x06UserId\x12\x1c\n" +
	"\tHelloName\x18\x02 \x01(\tR\tHelloName\"^\n" +
	"\fLoginRequest\x12\x1c\n" +
	"\tLogUserId\x18\x01 \x01(\x03R\tLogUserId\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"O\n" +
	"\rLoginResponse\x12\x18\n" +
	"\aLUserId\x18\x01 \x01(\x03R\aLUserId\x12$\n" +
	"\rSvosrasheniem\x18\x02 \x01(\tR\rSvosrasheniem2h\n" +
	"\x04Auth\x12,\n" +
	"\x03Reg\x12\x10.auth.RegRequest\x1a\x11.auth.RegResponse\"\x00\x122\n" +
	"\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponse\"\x00B\x1bZ\x19github.com/legi/auth;k3meb\x06proto3"

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData []byte
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)))
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_auth_proto_goTypes = []any{
	(*RegRequest)(nil),    // 0: auth.RegRequest
	(*RegResponse)(nil),   // 1: auth.RegResponse
	(*LoginRequest)(nil),  // 2: auth.LoginRequest
	(*LoginResponse)(nil), // 3: auth.LoginResponse
}
var file_proto_auth_proto_depIdxs = []int32{
	0, // 0: auth.Auth.Reg:input_type -> auth.RegRequest
	2, // 1: auth.Auth.Login:input_type -> auth.LoginRequest
	1, // 2: auth.Auth.Reg:output_type -> auth.RegResponse
	3, // 3: auth.Auth.Login:output_type -> auth.LoginResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}
