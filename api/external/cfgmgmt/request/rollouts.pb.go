// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/request/rollouts.proto

package request

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

type SCMType int32

const (
	SCMType_UNKNOWN_SCM SCMType = 0
	SCMType_GIT         SCMType = 1
)

var SCMType_name = map[int32]string{
	0: "UNKNOWN_SCM",
	1: "GIT",
}

var SCMType_value = map[string]int32{
	"UNKNOWN_SCM": 0,
	"GIT":         1,
}

func (x SCMType) String() string {
	return proto.EnumName(SCMType_name, int32(x))
}

func (SCMType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{0}
}

type SCMWebType int32

const (
	SCMWebType_UNKNOWN_SCM_WEB SCMWebType = 0
	SCMWebType_GITHUB          SCMWebType = 1
)

var SCMWebType_name = map[int32]string{
	0: "UNKNOWN_SCM_WEB",
	1: "GITHUB",
}

var SCMWebType_value = map[string]int32{
	"UNKNOWN_SCM_WEB": 0,
	"GITHUB":          1,
}

func (x SCMWebType) String() string {
	return proto.EnumName(SCMWebType_name, int32(x))
}

func (SCMWebType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{1}
}

// CreateRollout is a request to create a new Rollout. All
// fields have the same meaning as with the response Rollout
// type.
type CreateRollout struct {
	PolicyName           string     `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyNodeGroup      string     `protobuf:"bytes,2,opt,name=policy_node_group,json=policyNodeGroup,proto3" json:"policy_node_group,omitempty"`
	PolicyRevisionId     string     `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty"`
	PolicyDomainUrl      string     `protobuf:"bytes,4,opt,name=policy_domain_url,json=policyDomainUrl,proto3" json:"policy_domain_url,omitempty"`
	ScmType              SCMType    `protobuf:"varint,5,opt,name=scm_type,json=scmType,proto3,enum=chef.automate.api.cfgmgmt.request.SCMType" json:"scm_type,omitempty"`
	ScmWebType           SCMWebType `protobuf:"varint,6,opt,name=scm_web_type,json=scmWebType,proto3,enum=chef.automate.api.cfgmgmt.request.SCMWebType" json:"scm_web_type,omitempty"`
	PolicyScmUrl         string     `protobuf:"bytes,7,opt,name=policy_scm_url,json=policyScmUrl,proto3" json:"policy_scm_url,omitempty"`
	PolicyScmWebUrl      string     `protobuf:"bytes,8,opt,name=policy_scm_web_url,json=policyScmWebUrl,proto3" json:"policy_scm_web_url,omitempty"`
	PolicyScmCommit      string     `protobuf:"bytes,9,opt,name=policy_scm_commit,json=policyScmCommit,proto3" json:"policy_scm_commit,omitempty"`
	Description          string     `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	CiJobUrl             string     `protobuf:"bytes,11,opt,name=ci_job_url,json=ciJobUrl,proto3" json:"ci_job_url,omitempty"`
	CiJobId              string     `protobuf:"bytes,12,opt,name=ci_job_id,json=ciJobId,proto3" json:"ci_job_id,omitempty"`
	ScmAuthorName        string     `protobuf:"bytes,13,opt,name=scm_author_name,json=scmAuthorName,proto3" json:"scm_author_name,omitempty"`
	ScmAuthorEmail       string     `protobuf:"bytes,14,opt,name=scm_author_email,json=scmAuthorEmail,proto3" json:"scm_author_email,omitempty"`
	PolicyDomainUsername string     `protobuf:"bytes,15,opt,name=policy_domain_username,json=policyDomainUsername,proto3" json:"policy_domain_username,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateRollout) Reset()         { *m = CreateRollout{} }
func (m *CreateRollout) String() string { return proto.CompactTextString(m) }
func (*CreateRollout) ProtoMessage()    {}
func (*CreateRollout) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{0}
}

func (m *CreateRollout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRollout.Unmarshal(m, b)
}
func (m *CreateRollout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRollout.Marshal(b, m, deterministic)
}
func (m *CreateRollout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRollout.Merge(m, src)
}
func (m *CreateRollout) XXX_Size() int {
	return xxx_messageInfo_CreateRollout.Size(m)
}
func (m *CreateRollout) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRollout.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRollout proto.InternalMessageInfo

func (m *CreateRollout) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *CreateRollout) GetPolicyNodeGroup() string {
	if m != nil {
		return m.PolicyNodeGroup
	}
	return ""
}

func (m *CreateRollout) GetPolicyRevisionId() string {
	if m != nil {
		return m.PolicyRevisionId
	}
	return ""
}

func (m *CreateRollout) GetPolicyDomainUrl() string {
	if m != nil {
		return m.PolicyDomainUrl
	}
	return ""
}

func (m *CreateRollout) GetScmType() SCMType {
	if m != nil {
		return m.ScmType
	}
	return SCMType_UNKNOWN_SCM
}

func (m *CreateRollout) GetScmWebType() SCMWebType {
	if m != nil {
		return m.ScmWebType
	}
	return SCMWebType_UNKNOWN_SCM_WEB
}

func (m *CreateRollout) GetPolicyScmUrl() string {
	if m != nil {
		return m.PolicyScmUrl
	}
	return ""
}

func (m *CreateRollout) GetPolicyScmWebUrl() string {
	if m != nil {
		return m.PolicyScmWebUrl
	}
	return ""
}

func (m *CreateRollout) GetPolicyScmCommit() string {
	if m != nil {
		return m.PolicyScmCommit
	}
	return ""
}

func (m *CreateRollout) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateRollout) GetCiJobUrl() string {
	if m != nil {
		return m.CiJobUrl
	}
	return ""
}

func (m *CreateRollout) GetCiJobId() string {
	if m != nil {
		return m.CiJobId
	}
	return ""
}

func (m *CreateRollout) GetScmAuthorName() string {
	if m != nil {
		return m.ScmAuthorName
	}
	return ""
}

func (m *CreateRollout) GetScmAuthorEmail() string {
	if m != nil {
		return m.ScmAuthorEmail
	}
	return ""
}

func (m *CreateRollout) GetPolicyDomainUsername() string {
	if m != nil {
		return m.PolicyDomainUsername
	}
	return ""
}

type Rollouts struct {
	// Filters to apply to the request for the rollouts list.
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rollouts) Reset()         { *m = Rollouts{} }
func (m *Rollouts) String() string { return proto.CompactTextString(m) }
func (*Rollouts) ProtoMessage()    {}
func (*Rollouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{1}
}

func (m *Rollouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rollouts.Unmarshal(m, b)
}
func (m *Rollouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rollouts.Marshal(b, m, deterministic)
}
func (m *Rollouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rollouts.Merge(m, src)
}
func (m *Rollouts) XXX_Size() int {
	return xxx_messageInfo_Rollouts.Size(m)
}
func (m *Rollouts) XXX_DiscardUnknown() {
	xxx_messageInfo_Rollouts.DiscardUnknown(m)
}

var xxx_messageInfo_Rollouts proto.InternalMessageInfo

func (m *Rollouts) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type RolloutById struct {
	RolloutId            string   `protobuf:"bytes,1,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolloutById) Reset()         { *m = RolloutById{} }
func (m *RolloutById) String() string { return proto.CompactTextString(m) }
func (*RolloutById) ProtoMessage()    {}
func (*RolloutById) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{2}
}

func (m *RolloutById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolloutById.Unmarshal(m, b)
}
func (m *RolloutById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolloutById.Marshal(b, m, deterministic)
}
func (m *RolloutById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolloutById.Merge(m, src)
}
func (m *RolloutById) XXX_Size() int {
	return xxx_messageInfo_RolloutById.Size(m)
}
func (m *RolloutById) XXX_DiscardUnknown() {
	xxx_messageInfo_RolloutById.DiscardUnknown(m)
}

var xxx_messageInfo_RolloutById proto.InternalMessageInfo

func (m *RolloutById) GetRolloutId() string {
	if m != nil {
		return m.RolloutId
	}
	return ""
}

type RolloutForChefRun struct {
	PolicyName           string   `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyGroup          string   `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	PolicyRevisionId     string   `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolloutForChefRun) Reset()         { *m = RolloutForChefRun{} }
func (m *RolloutForChefRun) String() string { return proto.CompactTextString(m) }
func (*RolloutForChefRun) ProtoMessage()    {}
func (*RolloutForChefRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{3}
}

func (m *RolloutForChefRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolloutForChefRun.Unmarshal(m, b)
}
func (m *RolloutForChefRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolloutForChefRun.Marshal(b, m, deterministic)
}
func (m *RolloutForChefRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolloutForChefRun.Merge(m, src)
}
func (m *RolloutForChefRun) XXX_Size() int {
	return xxx_messageInfo_RolloutForChefRun.Size(m)
}
func (m *RolloutForChefRun) XXX_DiscardUnknown() {
	xxx_messageInfo_RolloutForChefRun.DiscardUnknown(m)
}

var xxx_messageInfo_RolloutForChefRun proto.InternalMessageInfo

func (m *RolloutForChefRun) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *RolloutForChefRun) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *RolloutForChefRun) GetPolicyRevisionId() string {
	if m != nil {
		return m.PolicyRevisionId
	}
	return ""
}

type ListNodeSegmentsWithRolloutProgress struct {
	// Filters to apply to the request for the node segments list.
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNodeSegmentsWithRolloutProgress) Reset()         { *m = ListNodeSegmentsWithRolloutProgress{} }
func (m *ListNodeSegmentsWithRolloutProgress) String() string { return proto.CompactTextString(m) }
func (*ListNodeSegmentsWithRolloutProgress) ProtoMessage()    {}
func (*ListNodeSegmentsWithRolloutProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{4}
}

func (m *ListNodeSegmentsWithRolloutProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodeSegmentsWithRolloutProgress.Unmarshal(m, b)
}
func (m *ListNodeSegmentsWithRolloutProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodeSegmentsWithRolloutProgress.Marshal(b, m, deterministic)
}
func (m *ListNodeSegmentsWithRolloutProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodeSegmentsWithRolloutProgress.Merge(m, src)
}
func (m *ListNodeSegmentsWithRolloutProgress) XXX_Size() int {
	return xxx_messageInfo_ListNodeSegmentsWithRolloutProgress.Size(m)
}
func (m *ListNodeSegmentsWithRolloutProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodeSegmentsWithRolloutProgress.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodeSegmentsWithRolloutProgress proto.InternalMessageInfo

func (m *ListNodeSegmentsWithRolloutProgress) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterEnum("chef.automate.api.cfgmgmt.request.SCMType", SCMType_name, SCMType_value)
	proto.RegisterEnum("chef.automate.api.cfgmgmt.request.SCMWebType", SCMWebType_name, SCMWebType_value)
	proto.RegisterType((*CreateRollout)(nil), "chef.automate.api.cfgmgmt.request.CreateRollout")
	proto.RegisterType((*Rollouts)(nil), "chef.automate.api.cfgmgmt.request.Rollouts")
	proto.RegisterType((*RolloutById)(nil), "chef.automate.api.cfgmgmt.request.RolloutById")
	proto.RegisterType((*RolloutForChefRun)(nil), "chef.automate.api.cfgmgmt.request.RolloutForChefRun")
	proto.RegisterType((*ListNodeSegmentsWithRolloutProgress)(nil), "chef.automate.api.cfgmgmt.request.ListNodeSegmentsWithRolloutProgress")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/request/rollouts.proto", fileDescriptor_8cacd55819479c9d)
}

var fileDescriptor_8cacd55819479c9d = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0xc7, 0x97, 0x67, 0xcf, 0xfa, 0x72, 0xba, 0xb5, 0x9d, 0x41, 0x53, 0x84, 0x40, 0x74, 0x1d,
	0x42, 0x55, 0xb7, 0x25, 0x12, 0x2f, 0xe2, 0x8a, 0x0b, 0x5a, 0xc6, 0x28, 0xb0, 0x0e, 0xb5, 0x9b,
	0x2a, 0x71, 0x13, 0xa5, 0xc9, 0x69, 0x6b, 0x14, 0xc7, 0xc1, 0x76, 0x80, 0x7e, 0x00, 0x3e, 0x10,
	0xdf, 0x10, 0xd9, 0xf1, 0xb6, 0x6e, 0x12, 0x0c, 0x71, 0x97, 0xfc, 0xcf, 0xef, 0xfc, 0xed, 0x63,
	0x9f, 0x63, 0xd8, 0x0f, 0x33, 0xea, 0xe3, 0x77, 0x85, 0x22, 0x0d, 0x13, 0x3f, 0x9a, 0xcd, 0xd9,
	0x9c, 0x29, 0x5f, 0xe0, 0x97, 0x1c, 0xa5, 0xf2, 0x05, 0x4f, 0x12, 0x9e, 0x2b, 0xe9, 0x65, 0x82,
	0x2b, 0x4e, 0x76, 0xa3, 0x05, 0xce, 0xbc, 0x30, 0x57, 0x9c, 0x85, 0x0a, 0xbd, 0x30, 0xa3, 0x9e,
	0xcd, 0xf0, 0x6c, 0x46, 0xfb, 0xe7, 0x06, 0x6c, 0xf5, 0x05, 0x86, 0x0a, 0x47, 0x45, 0x2e, 0x79,
	0x08, 0xb5, 0x8c, 0x27, 0x34, 0x5a, 0x06, 0x69, 0xc8, 0xd0, 0x75, 0x5a, 0x4e, 0xa7, 0x3a, 0x82,
	0x42, 0x1a, 0x86, 0x0c, 0x49, 0x17, 0xb6, 0x2f, 0x00, 0x1e, 0x63, 0x30, 0x17, 0x3c, 0xcf, 0xdc,
	0xff, 0x0c, 0xd6, 0xb0, 0x18, 0x8f, 0xf1, 0x58, 0xcb, 0xe4, 0x00, 0x88, 0x65, 0x05, 0x7e, 0xa5,
	0x92, 0xf2, 0x34, 0xa0, 0xb1, 0xbb, 0x6e, 0xe0, 0x66, 0x11, 0x19, 0xd9, 0xc0, 0x20, 0x5e, 0x71,
	0x8e, 0x39, 0x0b, 0x69, 0x1a, 0xe4, 0x22, 0x71, 0xff, 0x5f, 0x75, 0x7e, 0x6d, 0xf4, 0x73, 0x91,
	0x90, 0x23, 0xa8, 0xc8, 0x88, 0x05, 0x6a, 0x99, 0xa1, 0xbb, 0xd1, 0x72, 0x3a, 0xf5, 0x27, 0x5d,
	0xef, 0xd6, 0x72, 0xbd, 0x71, 0xff, 0xe4, 0x6c, 0x99, 0xe1, 0xa8, 0x2c, 0x23, 0xa6, 0x3f, 0xc8,
	0x29, 0x6c, 0x6a, 0x9b, 0x6f, 0x38, 0x2d, 0xac, 0x4a, 0xc6, 0xea, 0xf0, 0xef, 0xac, 0x26, 0x38,
	0x35, 0x6e, 0x20, 0x23, 0x66, 0xbf, 0xc9, 0x23, 0xa8, 0xdb, 0x1a, 0xb4, 0xaf, 0x2e, 0xa0, 0x6c,
	0x0a, 0xd8, 0x2c, 0xd4, 0x71, 0xc4, 0xf4, 0xee, 0xf7, 0x2f, 0xcf, 0xe5, 0x62, 0x75, 0x4d, 0x56,
	0x56, 0x4b, 0x1d, 0x1b, 0x4f, 0x0d, 0x5f, 0x1d, 0x8b, 0x86, 0x23, 0xce, 0x18, 0x55, 0x6e, 0xf5,
	0x06, 0xdb, 0x37, 0x32, 0x69, 0x41, 0x2d, 0x46, 0x19, 0x09, 0x9a, 0x29, 0xca, 0x53, 0x17, 0x0c,
	0xb5, 0x2a, 0x91, 0xfb, 0x00, 0x11, 0x0d, 0x3e, 0xf3, 0x62, 0xc9, 0x9a, 0x01, 0x2a, 0x11, 0x7d,
	0xc7, 0xcd, 0x5a, 0xf7, 0xa0, 0x6a, 0xa3, 0x34, 0x76, 0x37, 0x4d, 0xb0, 0x6c, 0x82, 0x83, 0x98,
	0x3c, 0x86, 0x86, 0xde, 0x40, 0x98, 0xab, 0x05, 0x17, 0x45, 0x77, 0x6c, 0x19, 0x62, 0x4b, 0x46,
	0xec, 0x95, 0x51, 0x4d, 0x83, 0x74, 0xa0, 0xb9, 0xc2, 0x21, 0x0b, 0x69, 0xe2, 0xd6, 0x0d, 0x58,
	0xbf, 0x04, 0x8f, 0xb4, 0x4a, 0x9e, 0xc1, 0xce, 0x8d, 0x0b, 0x97, 0xba, 0xb1, 0x19, 0xba, 0x0d,
	0xc3, 0xdf, 0xbd, 0x76, 0xeb, 0x36, 0xd6, 0x6e, 0x43, 0xc5, 0x36, 0xab, 0x24, 0x3b, 0x50, 0x9a,
	0xd1, 0x44, 0xa1, 0x70, 0x9d, 0xd6, 0x7a, 0xa7, 0x3a, 0xb2, 0x7f, 0xed, 0x03, 0xa8, 0x59, 0xa6,
	0xb7, 0x1c, 0xc4, 0xe4, 0x01, 0x80, 0x9d, 0x0d, 0x5d, 0x57, 0xd1, 0xd3, 0x55, 0xab, 0x0c, 0xe2,
	0xf6, 0x0f, 0x07, 0xb6, 0x2d, 0xfe, 0x86, 0x8b, 0xfe, 0x02, 0x67, 0xa3, 0x3c, 0xbd, 0x7d, 0x12,
	0x76, 0xc1, 0xde, 0xea, 0xb5, 0x21, 0xb0, 0x49, 0xff, 0x30, 0x00, 0xed, 0x97, 0xb0, 0xf7, 0x81,
	0x4a, 0xa5, 0xe7, 0x67, 0x8c, 0x73, 0x86, 0xa9, 0x92, 0x13, 0xaa, 0x16, 0x76, 0x6b, 0x1f, 0x05,
	0x9f, 0x0b, 0x94, 0xbf, 0x2d, 0xba, 0xbb, 0x07, 0x65, 0xdb, 0xe0, 0xa4, 0x01, 0xb5, 0xf3, 0xe1,
	0xfb, 0xe1, 0xe9, 0x64, 0x18, 0x8c, 0xfb, 0x27, 0xcd, 0x35, 0x52, 0x86, 0xf5, 0xe3, 0xc1, 0x59,
	0xd3, 0xe9, 0x1e, 0x02, 0x5c, 0xb5, 0x2e, 0xb9, 0x03, 0x8d, 0x15, 0x2e, 0x98, 0x1c, 0xf5, 0x9a,
	0x6b, 0x04, 0xa0, 0x74, 0x3c, 0x38, 0x7b, 0x7b, 0xde, 0x6b, 0x3a, 0xbd, 0x17, 0x9f, 0x9e, 0xcf,
	0xa9, 0x5a, 0xe4, 0x53, 0x2f, 0xe2, 0xcc, 0xd7, 0x63, 0xe1, 0x5f, 0x8c, 0x85, 0xff, 0xa7, 0xb7,
	0x68, 0x5a, 0x32, 0x6f, 0xd0, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4a, 0x51, 0xd9,
	0xb2, 0x04, 0x00, 0x00,
}
