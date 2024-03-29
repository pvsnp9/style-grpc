// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: pkg/proto/user/type/login_type.proto

package user

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input    string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetInput() string {
	if x != nil {
		return x.Input
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=user_uuid,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogoutResquest) Reset() {
	*x = LogoutResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResquest) ProtoMessage() {}

func (x *LogoutResquest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResquest.ProtoReflect.Descriptor instead.
func (*LogoutResquest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{2}
}

func (x *LogoutResquest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RenewTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RenewTokenRequest) Reset() {
	*x = RenewTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewTokenRequest) ProtoMessage() {}

func (x *RenewTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewTokenRequest.ProtoReflect.Descriptor instead.
func (*RenewTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{4}
}

func (x *RenewTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RenewTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RenewTokenResponse) Reset() {
	*x = RenewTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewTokenResponse) ProtoMessage() {}

func (x *RenewTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewTokenResponse.ProtoReflect.Descriptor instead.
func (*RenewTokenResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{5}
}

func (x *RenewTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PasswordResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *PasswordResetRequest) Reset() {
	*x = PasswordResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResetRequest) ProtoMessage() {}

func (x *PasswordResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResetRequest.ProtoReflect.Descriptor instead.
func (*PasswordResetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{6}
}

func (x *PasswordResetRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type PasswordResetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PasswordResetResponse) Reset() {
	*x = PasswordResetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordResetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResetResponse) ProtoMessage() {}

func (x *PasswordResetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResetResponse.ProtoReflect.Descriptor instead.
func (*PasswordResetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{7}
}

func (x *PasswordResetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CodeVerficationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *CodeVerficationRequest) Reset() {
	*x = CodeVerficationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeVerficationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeVerficationRequest) ProtoMessage() {}

func (x *CodeVerficationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeVerficationRequest.ProtoReflect.Descriptor instead.
func (*CodeVerficationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{8}
}

func (x *CodeVerficationRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type CodeVerficationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CodeVerficationResponse) Reset() {
	*x = CodeVerficationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeVerficationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeVerficationResponse) ProtoMessage() {}

func (x *CodeVerficationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_type_login_type_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeVerficationResponse.ProtoReflect.Descriptor instead.
func (*CodeVerficationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_type_login_type_proto_rawDescGZIP(), []int{9}
}

func (x *CodeVerficationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_proto_user_type_login_type_proto protoreflect.FileDescriptor

var file_pkg_proto_user_type_login_type_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x40, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5d,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x26, 0x0a,
	0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x29, 0x0a, 0x11, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x12,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x14, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x31, 0x0a, 0x15, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x16, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x65, 0x72,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x33, 0x0a, 0x17, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x65, 0x72, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x76, 0x73, 0x6e, 0x70, 0x39, 0x2f, 0x73, 0x74, 0x79, 0x6c,
	0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_user_type_login_type_proto_rawDescOnce sync.Once
	file_pkg_proto_user_type_login_type_proto_rawDescData = file_pkg_proto_user_type_login_type_proto_rawDesc
)

func file_pkg_proto_user_type_login_type_proto_rawDescGZIP() []byte {
	file_pkg_proto_user_type_login_type_proto_rawDescOnce.Do(func() {
		file_pkg_proto_user_type_login_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_user_type_login_type_proto_rawDescData)
	})
	return file_pkg_proto_user_type_login_type_proto_rawDescData
}

var file_pkg_proto_user_type_login_type_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_proto_user_type_login_type_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),            // 0: user.LoginRequest
	(*LoginResponse)(nil),           // 1: user.LoginResponse
	(*LogoutResquest)(nil),          // 2: user.LogoutResquest
	(*LogoutResponse)(nil),          // 3: user.LogoutResponse
	(*RenewTokenRequest)(nil),       // 4: user.RenewTokenRequest
	(*RenewTokenResponse)(nil),      // 5: user.RenewTokenResponse
	(*PasswordResetRequest)(nil),    // 6: user.PasswordResetRequest
	(*PasswordResetResponse)(nil),   // 7: user.PasswordResetResponse
	(*CodeVerficationRequest)(nil),  // 8: user.CodeVerficationRequest
	(*CodeVerficationResponse)(nil), // 9: user.CodeVerficationResponse
}
var file_pkg_proto_user_type_login_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_user_type_login_type_proto_init() }
func file_pkg_proto_user_type_login_type_proto_init() {
	if File_pkg_proto_user_type_login_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_user_type_login_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResquest); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewTokenRequest); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewTokenResponse); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordResetRequest); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordResetResponse); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeVerficationRequest); i {
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
		file_pkg_proto_user_type_login_type_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeVerficationResponse); i {
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
			RawDescriptor: file_pkg_proto_user_type_login_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_user_type_login_type_proto_goTypes,
		DependencyIndexes: file_pkg_proto_user_type_login_type_proto_depIdxs,
		MessageInfos:      file_pkg_proto_user_type_login_type_proto_msgTypes,
	}.Build()
	File_pkg_proto_user_type_login_type_proto = out.File
	file_pkg_proto_user_type_login_type_proto_rawDesc = nil
	file_pkg_proto_user_type_login_type_proto_goTypes = nil
	file_pkg_proto_user_type_login_type_proto_depIdxs = nil
}
