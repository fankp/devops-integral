// Code generated by protoc-gen-go. DO NOT EDIT.
// source: privilege.proto

package devops_integral_upm_srv_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Privilege struct {
	PrivilegeId          int32    `protobuf:"varint,1,opt,name=privilegeId,proto3" json:"privilegeId,omitempty"`
	PrivilegeGroupId     int32    `protobuf:"varint,2,opt,name=privilegeGroupId,proto3" json:"privilegeGroupId,omitempty"`
	PrivilegeCode        string   `protobuf:"bytes,3,opt,name=privilegeCode,proto3" json:"privilegeCode,omitempty"`
	PrivilegeName        string   `protobuf:"bytes,4,opt,name=privilegeName,proto3" json:"privilegeName,omitempty"`
	CreatedOn            int32    `protobuf:"varint,5,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	CreatedBy            string   `protobuf:"bytes,6,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedOn            int32    `protobuf:"varint,7,opt,name=updatedOn,proto3" json:"updatedOn,omitempty"`
	UpdatedBy            string   `protobuf:"bytes,8,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	DeletedOn            int32    `protobuf:"varint,9,opt,name=deletedOn,proto3" json:"deletedOn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Privilege) Reset()         { *m = Privilege{} }
func (m *Privilege) String() string { return proto.CompactTextString(m) }
func (*Privilege) ProtoMessage()    {}
func (*Privilege) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e900001d546c195, []int{0}
}

func (m *Privilege) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Privilege.Unmarshal(m, b)
}
func (m *Privilege) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Privilege.Marshal(b, m, deterministic)
}
func (m *Privilege) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Privilege.Merge(m, src)
}
func (m *Privilege) XXX_Size() int {
	return xxx_messageInfo_Privilege.Size(m)
}
func (m *Privilege) XXX_DiscardUnknown() {
	xxx_messageInfo_Privilege.DiscardUnknown(m)
}

var xxx_messageInfo_Privilege proto.InternalMessageInfo

func (m *Privilege) GetPrivilegeId() int32 {
	if m != nil {
		return m.PrivilegeId
	}
	return 0
}

func (m *Privilege) GetPrivilegeGroupId() int32 {
	if m != nil {
		return m.PrivilegeGroupId
	}
	return 0
}

func (m *Privilege) GetPrivilegeCode() string {
	if m != nil {
		return m.PrivilegeCode
	}
	return ""
}

func (m *Privilege) GetPrivilegeName() string {
	if m != nil {
		return m.PrivilegeName
	}
	return ""
}

func (m *Privilege) GetCreatedOn() int32 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *Privilege) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Privilege) GetUpdatedOn() int32 {
	if m != nil {
		return m.UpdatedOn
	}
	return 0
}

func (m *Privilege) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *Privilege) GetDeletedOn() int32 {
	if m != nil {
		return m.DeletedOn
	}
	return 0
}

type PrivilegeGroup struct {
	PrivilegeGroupId     int32        `protobuf:"varint,1,opt,name=privilegeGroupId,proto3" json:"privilegeGroupId,omitempty"`
	PrivilegeGroupName   string       `protobuf:"bytes,2,opt,name=privilegeGroupName,proto3" json:"privilegeGroupName,omitempty"`
	Privileges           []*Privilege `protobuf:"bytes,3,rep,name=privileges,proto3" json:"privileges,omitempty"`
	CreatedOn            int32        `protobuf:"varint,4,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	CreatedBy            string       `protobuf:"bytes,5,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedOn            int32        `protobuf:"varint,6,opt,name=updatedOn,proto3" json:"updatedOn,omitempty"`
	UpdatedBy            string       `protobuf:"bytes,7,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	DeletedOn            int32        `protobuf:"varint,8,opt,name=deletedOn,proto3" json:"deletedOn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrivilegeGroup) Reset()         { *m = PrivilegeGroup{} }
func (m *PrivilegeGroup) String() string { return proto.CompactTextString(m) }
func (*PrivilegeGroup) ProtoMessage()    {}
func (*PrivilegeGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e900001d546c195, []int{1}
}

func (m *PrivilegeGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivilegeGroup.Unmarshal(m, b)
}
func (m *PrivilegeGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivilegeGroup.Marshal(b, m, deterministic)
}
func (m *PrivilegeGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivilegeGroup.Merge(m, src)
}
func (m *PrivilegeGroup) XXX_Size() int {
	return xxx_messageInfo_PrivilegeGroup.Size(m)
}
func (m *PrivilegeGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivilegeGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PrivilegeGroup proto.InternalMessageInfo

func (m *PrivilegeGroup) GetPrivilegeGroupId() int32 {
	if m != nil {
		return m.PrivilegeGroupId
	}
	return 0
}

func (m *PrivilegeGroup) GetPrivilegeGroupName() string {
	if m != nil {
		return m.PrivilegeGroupName
	}
	return ""
}

func (m *PrivilegeGroup) GetPrivileges() []*Privilege {
	if m != nil {
		return m.Privileges
	}
	return nil
}

func (m *PrivilegeGroup) GetCreatedOn() int32 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *PrivilegeGroup) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *PrivilegeGroup) GetUpdatedOn() int32 {
	if m != nil {
		return m.UpdatedOn
	}
	return 0
}

func (m *PrivilegeGroup) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

func (m *PrivilegeGroup) GetDeletedOn() int32 {
	if m != nil {
		return m.DeletedOn
	}
	return 0
}

type SelectPrivilegesReq struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProjectId            int32    `protobuf:"varint,2,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Admin                bool     `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectPrivilegesReq) Reset()         { *m = SelectPrivilegesReq{} }
func (m *SelectPrivilegesReq) String() string { return proto.CompactTextString(m) }
func (*SelectPrivilegesReq) ProtoMessage()    {}
func (*SelectPrivilegesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e900001d546c195, []int{2}
}

func (m *SelectPrivilegesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectPrivilegesReq.Unmarshal(m, b)
}
func (m *SelectPrivilegesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectPrivilegesReq.Marshal(b, m, deterministic)
}
func (m *SelectPrivilegesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectPrivilegesReq.Merge(m, src)
}
func (m *SelectPrivilegesReq) XXX_Size() int {
	return xxx_messageInfo_SelectPrivilegesReq.Size(m)
}
func (m *SelectPrivilegesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectPrivilegesReq.DiscardUnknown(m)
}

var xxx_messageInfo_SelectPrivilegesReq proto.InternalMessageInfo

func (m *SelectPrivilegesReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SelectPrivilegesReq) GetProjectId() int32 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *SelectPrivilegesReq) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

type SelectPrivilegeCodesResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	PrivilegeCodes       []string `protobuf:"bytes,3,rep,name=privilegeCodes,proto3" json:"privilegeCodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectPrivilegeCodesResp) Reset()         { *m = SelectPrivilegeCodesResp{} }
func (m *SelectPrivilegeCodesResp) String() string { return proto.CompactTextString(m) }
func (*SelectPrivilegeCodesResp) ProtoMessage()    {}
func (*SelectPrivilegeCodesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e900001d546c195, []int{3}
}

func (m *SelectPrivilegeCodesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectPrivilegeCodesResp.Unmarshal(m, b)
}
func (m *SelectPrivilegeCodesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectPrivilegeCodesResp.Marshal(b, m, deterministic)
}
func (m *SelectPrivilegeCodesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectPrivilegeCodesResp.Merge(m, src)
}
func (m *SelectPrivilegeCodesResp) XXX_Size() int {
	return xxx_messageInfo_SelectPrivilegeCodesResp.Size(m)
}
func (m *SelectPrivilegeCodesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectPrivilegeCodesResp.DiscardUnknown(m)
}

var xxx_messageInfo_SelectPrivilegeCodesResp proto.InternalMessageInfo

func (m *SelectPrivilegeCodesResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SelectPrivilegeCodesResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SelectPrivilegeCodesResp) GetPrivilegeCodes() []string {
	if m != nil {
		return m.PrivilegeCodes
	}
	return nil
}

type SelectPrivilegeGroupsResp struct {
	Success              bool              `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	PrivilegeGroups      []*PrivilegeGroup `protobuf:"bytes,3,rep,name=privilegeGroups,proto3" json:"privilegeGroups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SelectPrivilegeGroupsResp) Reset()         { *m = SelectPrivilegeGroupsResp{} }
func (m *SelectPrivilegeGroupsResp) String() string { return proto.CompactTextString(m) }
func (*SelectPrivilegeGroupsResp) ProtoMessage()    {}
func (*SelectPrivilegeGroupsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e900001d546c195, []int{4}
}

func (m *SelectPrivilegeGroupsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectPrivilegeGroupsResp.Unmarshal(m, b)
}
func (m *SelectPrivilegeGroupsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectPrivilegeGroupsResp.Marshal(b, m, deterministic)
}
func (m *SelectPrivilegeGroupsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectPrivilegeGroupsResp.Merge(m, src)
}
func (m *SelectPrivilegeGroupsResp) XXX_Size() int {
	return xxx_messageInfo_SelectPrivilegeGroupsResp.Size(m)
}
func (m *SelectPrivilegeGroupsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectPrivilegeGroupsResp.DiscardUnknown(m)
}

var xxx_messageInfo_SelectPrivilegeGroupsResp proto.InternalMessageInfo

func (m *SelectPrivilegeGroupsResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SelectPrivilegeGroupsResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SelectPrivilegeGroupsResp) GetPrivilegeGroups() []*PrivilegeGroup {
	if m != nil {
		return m.PrivilegeGroups
	}
	return nil
}

type CheckPrivilegeReq struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProjectId            int32    `protobuf:"varint,2,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Admin                bool     `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
	PrivilegeCode        string   `protobuf:"bytes,4,opt,name=privilegeCode,proto3" json:"privilegeCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPrivilegeReq) Reset()         { *m = CheckPrivilegeReq{} }
func (m *CheckPrivilegeReq) String() string { return proto.CompactTextString(m) }
func (*CheckPrivilegeReq) ProtoMessage()    {}
func (*CheckPrivilegeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e900001d546c195, []int{5}
}

func (m *CheckPrivilegeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPrivilegeReq.Unmarshal(m, b)
}
func (m *CheckPrivilegeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPrivilegeReq.Marshal(b, m, deterministic)
}
func (m *CheckPrivilegeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPrivilegeReq.Merge(m, src)
}
func (m *CheckPrivilegeReq) XXX_Size() int {
	return xxx_messageInfo_CheckPrivilegeReq.Size(m)
}
func (m *CheckPrivilegeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPrivilegeReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPrivilegeReq proto.InternalMessageInfo

func (m *CheckPrivilegeReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CheckPrivilegeReq) GetProjectId() int32 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *CheckPrivilegeReq) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func (m *CheckPrivilegeReq) GetPrivilegeCode() string {
	if m != nil {
		return m.PrivilegeCode
	}
	return ""
}

type CheckPrivilegeResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Passed               bool     `protobuf:"varint,3,opt,name=passed,proto3" json:"passed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPrivilegeResp) Reset()         { *m = CheckPrivilegeResp{} }
func (m *CheckPrivilegeResp) String() string { return proto.CompactTextString(m) }
func (*CheckPrivilegeResp) ProtoMessage()    {}
func (*CheckPrivilegeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e900001d546c195, []int{6}
}

func (m *CheckPrivilegeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPrivilegeResp.Unmarshal(m, b)
}
func (m *CheckPrivilegeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPrivilegeResp.Marshal(b, m, deterministic)
}
func (m *CheckPrivilegeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPrivilegeResp.Merge(m, src)
}
func (m *CheckPrivilegeResp) XXX_Size() int {
	return xxx_messageInfo_CheckPrivilegeResp.Size(m)
}
func (m *CheckPrivilegeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPrivilegeResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPrivilegeResp proto.InternalMessageInfo

func (m *CheckPrivilegeResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CheckPrivilegeResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CheckPrivilegeResp) GetPassed() bool {
	if m != nil {
		return m.Passed
	}
	return false
}

func init() {
	proto.RegisterType((*Privilege)(nil), "devops.integral.upm.srv.service.Privilege")
	proto.RegisterType((*PrivilegeGroup)(nil), "devops.integral.upm.srv.service.PrivilegeGroup")
	proto.RegisterType((*SelectPrivilegesReq)(nil), "devops.integral.upm.srv.service.SelectPrivilegesReq")
	proto.RegisterType((*SelectPrivilegeCodesResp)(nil), "devops.integral.upm.srv.service.SelectPrivilegeCodesResp")
	proto.RegisterType((*SelectPrivilegeGroupsResp)(nil), "devops.integral.upm.srv.service.SelectPrivilegeGroupsResp")
	proto.RegisterType((*CheckPrivilegeReq)(nil), "devops.integral.upm.srv.service.CheckPrivilegeReq")
	proto.RegisterType((*CheckPrivilegeResp)(nil), "devops.integral.upm.srv.service.CheckPrivilegeResp")
}

func init() { proto.RegisterFile("privilege.proto", fileDescriptor_9e900001d546c195) }

var fileDescriptor_9e900001d546c195 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xf3, 0xe3, 0xc4, 0x53, 0x11, 0xca, 0x52, 0xaa, 0xa5, 0x42, 0x22, 0xb2, 0x10, 0x8a,
	0x7a, 0x30, 0x52, 0xca, 0x05, 0x8e, 0xed, 0x01, 0x95, 0x03, 0x54, 0xdb, 0x13, 0x37, 0x8c, 0x77,
	0x14, 0x0c, 0x49, 0xbc, 0xec, 0xda, 0x96, 0xf2, 0x00, 0x08, 0xf1, 0x1c, 0xdc, 0x78, 0x05, 0xde,
	0x81, 0x67, 0x42, 0x5e, 0x3b, 0xeb, 0x9f, 0x98, 0x84, 0x44, 0x1c, 0xe7, 0x9b, 0xfd, 0x66, 0xc6,
	0xdf, 0x7c, 0xbb, 0x86, 0xbb, 0x42, 0x86, 0x69, 0x38, 0xc7, 0x19, 0x7a, 0x42, 0x46, 0x71, 0x44,
	0x1e, 0x73, 0x4c, 0x23, 0xa1, 0xbc, 0x70, 0x19, 0xe3, 0x4c, 0xfa, 0x73, 0x2f, 0x11, 0x0b, 0x4f,
	0xc9, 0xd4, 0x53, 0x28, 0xd3, 0x30, 0x40, 0xf7, 0x57, 0x07, 0x9c, 0x9b, 0x35, 0x89, 0x8c, 0xe1,
	0xc8, 0x54, 0xb8, 0xe6, 0xd4, 0x1a, 0x5b, 0x93, 0x3e, 0xab, 0x42, 0xe4, 0x1c, 0x8e, 0x4d, 0xf8,
	0x4a, 0x46, 0x89, 0xb8, 0xe6, 0xb4, 0xa3, 0x8f, 0x6d, 0xe0, 0xe4, 0x09, 0xdc, 0x31, 0xd8, 0x55,
	0xc4, 0x91, 0x76, 0xc7, 0xd6, 0xc4, 0x61, 0x75, 0xb0, 0x76, 0xea, 0x8d, 0xbf, 0x40, 0xda, 0x6b,
	0x9c, 0xca, 0x40, 0xf2, 0x08, 0x9c, 0x40, 0xa2, 0x1f, 0x23, 0x7f, 0xbb, 0xa4, 0x7d, 0xdd, 0xb0,
	0x04, 0x2a, 0xd9, 0xcb, 0x15, 0xb5, 0x35, 0xbf, 0x04, 0xb2, 0x6c, 0x22, 0x78, 0xc1, 0x1d, 0xe4,
	0x5c, 0x03, 0x54, 0xb2, 0x97, 0x2b, 0x3a, 0xcc, 0xb9, 0x06, 0xc8, 0xb2, 0x1c, 0xe7, 0x98, 0x73,
	0x9d, 0x9c, 0x6b, 0x00, 0xf7, 0x77, 0x07, 0x46, 0x37, 0xb5, 0xcf, 0x6e, 0x15, 0xc8, 0xfa, 0x8b,
	0x40, 0x1e, 0x90, 0x3a, 0xa6, 0xbf, 0xbf, 0xa3, 0x67, 0x68, 0xc9, 0x90, 0xd7, 0x00, 0x06, 0x55,
	0xb4, 0x3b, 0xee, 0x4e, 0x8e, 0xa6, 0xe7, 0xde, 0x8e, 0x15, 0x7b, 0x66, 0x40, 0x56, 0x61, 0xd7,
	0x05, 0xed, 0x6d, 0x15, 0xb4, 0xbf, 0x55, 0x50, 0x7b, 0xab, 0xa0, 0x83, 0xad, 0x82, 0x0e, 0x9b,
	0x82, 0xfa, 0x70, 0xff, 0x16, 0xe7, 0x18, 0xc4, 0x66, 0x68, 0xc5, 0xf0, 0x0b, 0x39, 0x05, 0x3b,
	0x51, 0x28, 0x8d, 0x94, 0x45, 0x94, 0x15, 0x13, 0x32, 0xfa, 0x84, 0x41, 0x6c, 0x6c, 0x58, 0x02,
	0xe4, 0x04, 0xfa, 0x3e, 0x5f, 0x84, 0x4b, 0xed, 0xbb, 0x21, 0xcb, 0x03, 0x37, 0x05, 0xda, 0x68,
	0x91, 0xd9, 0x50, 0x31, 0x54, 0x82, 0x50, 0x18, 0xa8, 0x24, 0x08, 0x50, 0x29, 0xdd, 0x68, 0xc8,
	0xd6, 0x61, 0x96, 0x59, 0xa0, 0x52, 0xfe, 0x6c, 0xbd, 0x9f, 0x75, 0x48, 0x9e, 0xc2, 0xa8, 0x66,
	0xe8, 0x7c, 0x31, 0x0e, 0x6b, 0xa0, 0xee, 0x4f, 0x0b, 0x1e, 0x36, 0x1a, 0xeb, 0xcd, 0x1e, 0xde,
	0xf9, 0x5d, 0xe5, 0xbe, 0xe7, 0xa5, 0x0a, 0x4f, 0x3c, 0xfb, 0x77, 0x4f, 0x68, 0x1e, 0x6b, 0xd6,
	0x71, 0xbf, 0x5a, 0x70, 0xef, 0xea, 0x23, 0x06, 0x9f, 0x4b, 0xf3, 0xfc, 0xdf, 0x35, 0x6c, 0x3e,
	0x0e, 0xbd, 0x96, 0xc7, 0xc1, 0x7d, 0x0f, 0xa4, 0x39, 0xc6, 0x81, 0x62, 0x9d, 0x82, 0x2d, 0x7c,
	0xa5, 0x90, 0x17, 0x63, 0x14, 0xd1, 0xf4, 0x47, 0x17, 0x8e, 0x4d, 0xf5, 0xdb, 0x5c, 0x1e, 0xf2,
	0xcd, 0x82, 0x93, 0x36, 0x93, 0x90, 0xe7, 0x3b, 0x95, 0x6d, 0xb1, 0xef, 0xd9, 0x8b, 0x7d, 0x59,
	0xa5, 0x23, 0x57, 0x30, 0xaa, 0x0b, 0x40, 0xa6, 0x3b, 0x8b, 0x6d, 0x2c, 0xee, 0xec, 0x62, 0x6f,
	0x8e, 0x12, 0xe4, 0xbb, 0x05, 0x0f, 0x5a, 0x0d, 0x7b, 0xa0, 0x0a, 0x2f, 0xf7, 0x65, 0x95, 0xd7,
	0xe3, 0x83, 0xad, 0x7f, 0x67, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x57, 0x72, 0xcf, 0x43,
	0xe1, 0x06, 0x00, 0x00,
}