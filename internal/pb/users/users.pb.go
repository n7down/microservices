// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/users/pb/users.proto

package users_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UsersGetPasswordRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersGetPasswordRequest) Reset()         { *m = UsersGetPasswordRequest{} }
func (m *UsersGetPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UsersGetPasswordRequest) ProtoMessage()    {}
func (*UsersGetPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{0}
}

func (m *UsersGetPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersGetPasswordRequest.Unmarshal(m, b)
}
func (m *UsersGetPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersGetPasswordRequest.Marshal(b, m, deterministic)
}
func (m *UsersGetPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersGetPasswordRequest.Merge(m, src)
}
func (m *UsersGetPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UsersGetPasswordRequest.Size(m)
}
func (m *UsersGetPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersGetPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersGetPasswordRequest proto.InternalMessageInfo

func (m *UsersGetPasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UsersGetPasswordResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersGetPasswordResponse) Reset()         { *m = UsersGetPasswordResponse{} }
func (m *UsersGetPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*UsersGetPasswordResponse) ProtoMessage()    {}
func (*UsersGetPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{1}
}

func (m *UsersGetPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersGetPasswordResponse.Unmarshal(m, b)
}
func (m *UsersGetPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersGetPasswordResponse.Marshal(b, m, deterministic)
}
func (m *UsersGetPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersGetPasswordResponse.Merge(m, src)
}
func (m *UsersGetPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_UsersGetPasswordResponse.Size(m)
}
func (m *UsersGetPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersGetPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersGetPasswordResponse proto.InternalMessageInfo

func (m *UsersGetPasswordResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UsersGetPasswordResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UsersCreateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersCreateRequest) Reset()         { *m = UsersCreateRequest{} }
func (m *UsersCreateRequest) String() string { return proto.CompactTextString(m) }
func (*UsersCreateRequest) ProtoMessage()    {}
func (*UsersCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{2}
}

func (m *UsersCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersCreateRequest.Unmarshal(m, b)
}
func (m *UsersCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersCreateRequest.Marshal(b, m, deterministic)
}
func (m *UsersCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersCreateRequest.Merge(m, src)
}
func (m *UsersCreateRequest) XXX_Size() int {
	return xxx_messageInfo_UsersCreateRequest.Size(m)
}
func (m *UsersCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersCreateRequest proto.InternalMessageInfo

func (m *UsersCreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UsersCreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UsersCreateRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UsersCreateRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type UsersCreateResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersCreateResponse) Reset()         { *m = UsersCreateResponse{} }
func (m *UsersCreateResponse) String() string { return proto.CompactTextString(m) }
func (*UsersCreateResponse) ProtoMessage()    {}
func (*UsersCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{3}
}

func (m *UsersCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersCreateResponse.Unmarshal(m, b)
}
func (m *UsersCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersCreateResponse.Marshal(b, m, deterministic)
}
func (m *UsersCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersCreateResponse.Merge(m, src)
}
func (m *UsersCreateResponse) XXX_Size() int {
	return xxx_messageInfo_UsersCreateResponse.Size(m)
}
func (m *UsersCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersCreateResponse proto.InternalMessageInfo

func (m *UsersCreateResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UsersUpdateRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersUpdateRequest) Reset()         { *m = UsersUpdateRequest{} }
func (m *UsersUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UsersUpdateRequest) ProtoMessage()    {}
func (*UsersUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{4}
}

func (m *UsersUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersUpdateRequest.Unmarshal(m, b)
}
func (m *UsersUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersUpdateRequest.Marshal(b, m, deterministic)
}
func (m *UsersUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersUpdateRequest.Merge(m, src)
}
func (m *UsersUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UsersUpdateRequest.Size(m)
}
func (m *UsersUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersUpdateRequest proto.InternalMessageInfo

func (m *UsersUpdateRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UsersUpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UsersUpdateRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UsersUpdateRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type UsersUpdateResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersUpdateResponse) Reset()         { *m = UsersUpdateResponse{} }
func (m *UsersUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UsersUpdateResponse) ProtoMessage()    {}
func (*UsersUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{5}
}

func (m *UsersUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersUpdateResponse.Unmarshal(m, b)
}
func (m *UsersUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersUpdateResponse.Marshal(b, m, deterministic)
}
func (m *UsersUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersUpdateResponse.Merge(m, src)
}
func (m *UsersUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UsersUpdateResponse.Size(m)
}
func (m *UsersUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersUpdateResponse proto.InternalMessageInfo

func (m *UsersUpdateResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UsersListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersListRequest) Reset()         { *m = UsersListRequest{} }
func (m *UsersListRequest) String() string { return proto.CompactTextString(m) }
func (*UsersListRequest) ProtoMessage()    {}
func (*UsersListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{6}
}

func (m *UsersListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersListRequest.Unmarshal(m, b)
}
func (m *UsersListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersListRequest.Marshal(b, m, deterministic)
}
func (m *UsersListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersListRequest.Merge(m, src)
}
func (m *UsersListRequest) XXX_Size() int {
	return xxx_messageInfo_UsersListRequest.Size(m)
}
func (m *UsersListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersListRequest proto.InternalMessageInfo

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{7}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type UsersListResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersListResponse) Reset()         { *m = UsersListResponse{} }
func (m *UsersListResponse) String() string { return proto.CompactTextString(m) }
func (*UsersListResponse) ProtoMessage()    {}
func (*UsersListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{8}
}

func (m *UsersListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersListResponse.Unmarshal(m, b)
}
func (m *UsersListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersListResponse.Marshal(b, m, deterministic)
}
func (m *UsersListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersListResponse.Merge(m, src)
}
func (m *UsersListResponse) XXX_Size() int {
	return xxx_messageInfo_UsersListResponse.Size(m)
}
func (m *UsersListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersListResponse proto.InternalMessageInfo

func (m *UsersListResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UsersByIDRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersByIDRequest) Reset()         { *m = UsersByIDRequest{} }
func (m *UsersByIDRequest) String() string { return proto.CompactTextString(m) }
func (*UsersByIDRequest) ProtoMessage()    {}
func (*UsersByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{9}
}

func (m *UsersByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersByIDRequest.Unmarshal(m, b)
}
func (m *UsersByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersByIDRequest.Marshal(b, m, deterministic)
}
func (m *UsersByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersByIDRequest.Merge(m, src)
}
func (m *UsersByIDRequest) XXX_Size() int {
	return xxx_messageInfo_UsersByIDRequest.Size(m)
}
func (m *UsersByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersByIDRequest proto.InternalMessageInfo

func (m *UsersByIDRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UsersByIDResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersByIDResponse) Reset()         { *m = UsersByIDResponse{} }
func (m *UsersByIDResponse) String() string { return proto.CompactTextString(m) }
func (*UsersByIDResponse) ProtoMessage()    {}
func (*UsersByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{10}
}

func (m *UsersByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersByIDResponse.Unmarshal(m, b)
}
func (m *UsersByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersByIDResponse.Marshal(b, m, deterministic)
}
func (m *UsersByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersByIDResponse.Merge(m, src)
}
func (m *UsersByIDResponse) XXX_Size() int {
	return xxx_messageInfo_UsersByIDResponse.Size(m)
}
func (m *UsersByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersByIDResponse proto.InternalMessageInfo

func (m *UsersByIDResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UsersByIDResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UsersByIDResponse) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UsersByIDResponse) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type UsersDeleteRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersDeleteRequest) Reset()         { *m = UsersDeleteRequest{} }
func (m *UsersDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*UsersDeleteRequest) ProtoMessage()    {}
func (*UsersDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{11}
}

func (m *UsersDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersDeleteRequest.Unmarshal(m, b)
}
func (m *UsersDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersDeleteRequest.Marshal(b, m, deterministic)
}
func (m *UsersDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersDeleteRequest.Merge(m, src)
}
func (m *UsersDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_UsersDeleteRequest.Size(m)
}
func (m *UsersDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersDeleteRequest proto.InternalMessageInfo

func (m *UsersDeleteRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UsersDeleteResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersDeleteResponse) Reset()         { *m = UsersDeleteResponse{} }
func (m *UsersDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*UsersDeleteResponse) ProtoMessage()    {}
func (*UsersDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87300fcef84e6764, []int{12}
}

func (m *UsersDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersDeleteResponse.Unmarshal(m, b)
}
func (m *UsersDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersDeleteResponse.Marshal(b, m, deterministic)
}
func (m *UsersDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersDeleteResponse.Merge(m, src)
}
func (m *UsersDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_UsersDeleteResponse.Size(m)
}
func (m *UsersDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersDeleteResponse proto.InternalMessageInfo

func (m *UsersDeleteResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*UsersGetPasswordRequest)(nil), "users_pb.UsersGetPasswordRequest")
	proto.RegisterType((*UsersGetPasswordResponse)(nil), "users_pb.UsersGetPasswordResponse")
	proto.RegisterType((*UsersCreateRequest)(nil), "users_pb.UsersCreateRequest")
	proto.RegisterType((*UsersCreateResponse)(nil), "users_pb.UsersCreateResponse")
	proto.RegisterType((*UsersUpdateRequest)(nil), "users_pb.UsersUpdateRequest")
	proto.RegisterType((*UsersUpdateResponse)(nil), "users_pb.UsersUpdateResponse")
	proto.RegisterType((*UsersListRequest)(nil), "users_pb.UsersListRequest")
	proto.RegisterType((*User)(nil), "users_pb.User")
	proto.RegisterType((*UsersListResponse)(nil), "users_pb.UsersListResponse")
	proto.RegisterType((*UsersByIDRequest)(nil), "users_pb.UsersByIDRequest")
	proto.RegisterType((*UsersByIDResponse)(nil), "users_pb.UsersByIDResponse")
	proto.RegisterType((*UsersDeleteRequest)(nil), "users_pb.UsersDeleteRequest")
	proto.RegisterType((*UsersDeleteResponse)(nil), "users_pb.UsersDeleteResponse")
}

func init() { proto.RegisterFile("internal/users/pb/users.proto", fileDescriptor_87300fcef84e6764) }

var fileDescriptor_87300fcef84e6764 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcb, 0x4e, 0xc2, 0x40,
	0x14, 0xa5, 0x05, 0x0d, 0x5c, 0x0c, 0xd1, 0x71, 0x61, 0x53, 0x21, 0xc1, 0x09, 0x26, 0xac, 0x20,
	0xc1, 0xb8, 0x70, 0xab, 0x0d, 0x86, 0x84, 0x85, 0xc1, 0xb0, 0x72, 0x61, 0x8a, 0xcc, 0xa2, 0x09,
	0xb6, 0xb5, 0x53, 0x34, 0xf8, 0x01, 0xfe, 0xa9, 0xff, 0x61, 0xe6, 0xd1, 0x3a, 0xc3, 0xb4, 0xc4,
	0xc4, 0xe8, 0xae, 0x73, 0x1f, 0x67, 0xce, 0x99, 0x7b, 0x6e, 0xa1, 0x13, 0x84, 0x29, 0x49, 0x42,
	0x7f, 0x35, 0x5c, 0x53, 0x92, 0xd0, 0x61, 0xbc, 0x10, 0x1f, 0x83, 0x38, 0x89, 0xd2, 0x08, 0xd5,
	0xf9, 0xe1, 0x31, 0x5e, 0xe0, 0x4b, 0x38, 0x99, 0xb3, 0xef, 0x5b, 0x92, 0xde, 0xf9, 0x94, 0xbe,
	0x45, 0xc9, 0x72, 0x46, 0x5e, 0xd6, 0x84, 0xa6, 0xc8, 0x85, 0x3a, 0x4b, 0x85, 0xfe, 0x33, 0x71,
	0xac, 0xae, 0xd5, 0x6f, 0xcc, 0xf2, 0x33, 0x1e, 0x83, 0x63, 0xb6, 0xd1, 0x38, 0x0a, 0x29, 0x41,
	0x2d, 0xb0, 0x27, 0x9e, 0xec, 0xb0, 0x27, 0x1e, 0xc3, 0xc9, 0x6a, 0x1c, 0x5b, 0xe0, 0x64, 0x67,
	0xfc, 0x61, 0x01, 0xe2, 0x40, 0x37, 0x09, 0xf1, 0x53, 0xf2, 0x83, 0xab, 0x77, 0xc1, 0xa1, 0x36,
	0x34, 0xc6, 0x41, 0x42, 0x53, 0xde, 0x58, 0xe5, 0xc9, 0xef, 0x00, 0xeb, 0x9c, 0xfa, 0x32, 0x59,
	0x13, 0x9d, 0xd9, 0x19, 0x9f, 0xc3, 0xb1, 0xc6, 0xa3, 0x58, 0x0b, 0x7e, 0x97, 0x74, 0xe7, 0xf1,
	0x52, 0xa1, 0x5b, 0xa0, 0x38, 0xa7, 0x6f, 0x6f, 0xd1, 0xff, 0x3d, 0xc5, 0xec, 0xee, 0x12, 0x8a,
	0x08, 0x0e, 0x79, 0xd9, 0x34, 0xa0, 0xa9, 0x24, 0x88, 0x57, 0x50, 0x63, 0xb1, 0x7f, 0x22, 0x7a,
	0x05, 0x47, 0x0a, 0x03, 0x49, 0xb3, 0x07, 0x7b, 0x3c, 0xe8, 0x58, 0xdd, 0x6a, 0xbf, 0x39, 0x6a,
	0x0d, 0x32, 0x0b, 0x0e, 0x58, 0x78, 0x26, 0x92, 0x18, 0x4b, 0xf2, 0xd7, 0x9b, 0x89, 0x57, 0xf2,
	0xba, 0x78, 0x23, 0xe1, 0x45, 0x4d, 0xb9, 0xe9, 0xfe, 0x40, 0x59, 0x4f, 0x8e, 0xdf, 0x23, 0x2b,
	0x52, 0x3a, 0xfe, 0x7c, 0x50, 0x59, 0x55, 0x31, 0xc5, 0xd1, 0x67, 0x15, 0x0e, 0x78, 0xdd, 0x3d,
	0x49, 0x5e, 0x83, 0x27, 0x82, 0x1e, 0xa4, 0x78, 0x65, 0xa9, 0xd0, 0x99, 0xfe, 0x4e, 0x05, 0x7b,
	0xea, 0xe2, 0x5d, 0x25, 0xe2, 0x6e, 0x5c, 0x41, 0x53, 0x68, 0x2a, 0x06, 0x47, 0xed, 0xad, 0x26,
	0x6d, 0xff, 0xdc, 0x4e, 0x49, 0xd6, 0x40, 0x13, 0x5e, 0x34, 0xd0, 0xb4, 0xf5, 0x30, 0xd0, 0x74,
	0x03, 0xe3, 0x0a, 0x1a, 0x43, 0x23, 0x37, 0x0c, 0x72, 0xb7, 0xaa, 0x15, 0x1f, 0xbb, 0xa7, 0x85,
	0x39, 0x03, 0x87, 0x39, 0xc3, 0xc0, 0x51, 0x2c, 0x65, 0xe0, 0xa8, 0x56, 0x52, 0xd4, 0x89, 0x01,
	0x1a, 0xea, 0xb4, 0xe9, 0x1b, 0xea, 0xf4, 0xa9, 0xe3, 0xca, 0x62, 0x9f, 0xff, 0x73, 0x2f, 0xbe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x50, 0x0b, 0x97, 0x94, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersServiceClient interface {
	UsersGetPassword(ctx context.Context, in *UsersGetPasswordRequest, opts ...grpc.CallOption) (*UsersGetPasswordResponse, error)
	UsersCreate(ctx context.Context, in *UsersCreateRequest, opts ...grpc.CallOption) (*UsersCreateResponse, error)
	UsersUpdate(ctx context.Context, in *UsersUpdateRequest, opts ...grpc.CallOption) (*UsersUpdateResponse, error)
	UsersList(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error)
	UsersByID(ctx context.Context, in *UsersByIDRequest, opts ...grpc.CallOption) (*UsersByIDResponse, error)
	UsersDelete(ctx context.Context, in *UsersDeleteRequest, opts ...grpc.CallOption) (*UsersDeleteResponse, error)
}

type usersServiceClient struct {
	cc *grpc.ClientConn
}

func NewUsersServiceClient(cc *grpc.ClientConn) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) UsersGetPassword(ctx context.Context, in *UsersGetPasswordRequest, opts ...grpc.CallOption) (*UsersGetPasswordResponse, error) {
	out := new(UsersGetPasswordResponse)
	err := c.cc.Invoke(ctx, "/users_pb.UsersService/UsersGetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UsersCreate(ctx context.Context, in *UsersCreateRequest, opts ...grpc.CallOption) (*UsersCreateResponse, error) {
	out := new(UsersCreateResponse)
	err := c.cc.Invoke(ctx, "/users_pb.UsersService/UsersCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UsersUpdate(ctx context.Context, in *UsersUpdateRequest, opts ...grpc.CallOption) (*UsersUpdateResponse, error) {
	out := new(UsersUpdateResponse)
	err := c.cc.Invoke(ctx, "/users_pb.UsersService/UsersUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UsersList(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error) {
	out := new(UsersListResponse)
	err := c.cc.Invoke(ctx, "/users_pb.UsersService/UsersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UsersByID(ctx context.Context, in *UsersByIDRequest, opts ...grpc.CallOption) (*UsersByIDResponse, error) {
	out := new(UsersByIDResponse)
	err := c.cc.Invoke(ctx, "/users_pb.UsersService/UsersByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UsersDelete(ctx context.Context, in *UsersDeleteRequest, opts ...grpc.CallOption) (*UsersDeleteResponse, error) {
	out := new(UsersDeleteResponse)
	err := c.cc.Invoke(ctx, "/users_pb.UsersService/UsersDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
type UsersServiceServer interface {
	UsersGetPassword(context.Context, *UsersGetPasswordRequest) (*UsersGetPasswordResponse, error)
	UsersCreate(context.Context, *UsersCreateRequest) (*UsersCreateResponse, error)
	UsersUpdate(context.Context, *UsersUpdateRequest) (*UsersUpdateResponse, error)
	UsersList(context.Context, *UsersListRequest) (*UsersListResponse, error)
	UsersByID(context.Context, *UsersByIDRequest) (*UsersByIDResponse, error)
	UsersDelete(context.Context, *UsersDeleteRequest) (*UsersDeleteResponse, error)
}

// UnimplementedUsersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (*UnimplementedUsersServiceServer) UsersGetPassword(ctx context.Context, req *UsersGetPasswordRequest) (*UsersGetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersGetPassword not implemented")
}
func (*UnimplementedUsersServiceServer) UsersCreate(ctx context.Context, req *UsersCreateRequest) (*UsersCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersCreate not implemented")
}
func (*UnimplementedUsersServiceServer) UsersUpdate(ctx context.Context, req *UsersUpdateRequest) (*UsersUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersUpdate not implemented")
}
func (*UnimplementedUsersServiceServer) UsersList(ctx context.Context, req *UsersListRequest) (*UsersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersList not implemented")
}
func (*UnimplementedUsersServiceServer) UsersByID(ctx context.Context, req *UsersByIDRequest) (*UsersByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersByID not implemented")
}
func (*UnimplementedUsersServiceServer) UsersDelete(ctx context.Context, req *UsersDeleteRequest) (*UsersDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersDelete not implemented")
}

func RegisterUsersServiceServer(s *grpc.Server, srv UsersServiceServer) {
	s.RegisterService(&_UsersService_serviceDesc, srv)
}

func _UsersService_UsersGetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersGetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UsersGetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users_pb.UsersService/UsersGetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UsersGetPassword(ctx, req.(*UsersGetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UsersCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UsersCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users_pb.UsersService/UsersCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UsersCreate(ctx, req.(*UsersCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UsersUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UsersUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users_pb.UsersService/UsersUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UsersUpdate(ctx, req.(*UsersUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users_pb.UsersService/UsersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UsersList(ctx, req.(*UsersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UsersByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UsersByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users_pb.UsersService/UsersByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UsersByID(ctx, req.(*UsersByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UsersDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UsersDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users_pb.UsersService/UsersDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UsersDelete(ctx, req.(*UsersDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users_pb.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UsersGetPassword",
			Handler:    _UsersService_UsersGetPassword_Handler,
		},
		{
			MethodName: "UsersCreate",
			Handler:    _UsersService_UsersCreate_Handler,
		},
		{
			MethodName: "UsersUpdate",
			Handler:    _UsersService_UsersUpdate_Handler,
		},
		{
			MethodName: "UsersList",
			Handler:    _UsersService_UsersList_Handler,
		},
		{
			MethodName: "UsersByID",
			Handler:    _UsersService_UsersByID_Handler,
		},
		{
			MethodName: "UsersDelete",
			Handler:    _UsersService_UsersDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/users/pb/users.proto",
}