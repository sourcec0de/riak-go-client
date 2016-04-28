// Code generated by protoc-gen-go.
// source: riak.proto
// DO NOT EDIT!

/*
Package riak is a generated protocol buffer package.

It is generated from these files:
	riak.proto

It has these top-level messages:
	RpbErrorResp
	RpbGetServerInfoResp
	RpbPair
	RpbGetBucketReq
	RpbGetBucketResp
	RpbSetBucketReq
	RpbResetBucketReq
	RpbGetBucketTypeReq
	RpbSetBucketTypeReq
	RpbModFun
	RpbCommitHook
	RpbBucketProps
	RpbAuthReq
	RpbToggleEncodingReq
	RpbToggleEncodingResp
*/
package riak

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Used by riak_repl bucket fixup
type RpbBucketProps_RpbReplMode int32

const (
	RpbBucketProps_FALSE    RpbBucketProps_RpbReplMode = 0
	RpbBucketProps_REALTIME RpbBucketProps_RpbReplMode = 1
	RpbBucketProps_FULLSYNC RpbBucketProps_RpbReplMode = 2
	RpbBucketProps_TRUE     RpbBucketProps_RpbReplMode = 3
)

var RpbBucketProps_RpbReplMode_name = map[int32]string{
	0: "FALSE",
	1: "REALTIME",
	2: "FULLSYNC",
	3: "TRUE",
}
var RpbBucketProps_RpbReplMode_value = map[string]int32{
	"FALSE":    0,
	"REALTIME": 1,
	"FULLSYNC": 2,
	"TRUE":     3,
}

func (x RpbBucketProps_RpbReplMode) Enum() *RpbBucketProps_RpbReplMode {
	p := new(RpbBucketProps_RpbReplMode)
	*p = x
	return p
}
func (x RpbBucketProps_RpbReplMode) String() string {
	return proto.EnumName(RpbBucketProps_RpbReplMode_name, int32(x))
}
func (x *RpbBucketProps_RpbReplMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpbBucketProps_RpbReplMode_value, data, "RpbBucketProps_RpbReplMode")
	if err != nil {
		return err
	}
	*x = RpbBucketProps_RpbReplMode(value)
	return nil
}
func (RpbBucketProps_RpbReplMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11, 0}
}

// Error response - may be generated for any Req
type RpbErrorResp struct {
	Errmsg           []byte  `protobuf:"bytes,1,req,name=errmsg" json:"errmsg,omitempty"`
	Errcode          *uint32 `protobuf:"varint,2,req,name=errcode" json:"errcode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbErrorResp) Reset()                    { *m = RpbErrorResp{} }
func (m *RpbErrorResp) String() string            { return proto.CompactTextString(m) }
func (*RpbErrorResp) ProtoMessage()               {}
func (*RpbErrorResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpbErrorResp) GetErrmsg() []byte {
	if m != nil {
		return m.Errmsg
	}
	return nil
}

func (m *RpbErrorResp) GetErrcode() uint32 {
	if m != nil && m.Errcode != nil {
		return *m.Errcode
	}
	return 0
}

// Get server info request - no message defined, just send RpbGetServerInfoReq message code
type RpbGetServerInfoResp struct {
	Node             []byte `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	ServerVersion    []byte `protobuf:"bytes,2,opt,name=server_version" json:"server_version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetServerInfoResp) Reset()                    { *m = RpbGetServerInfoResp{} }
func (m *RpbGetServerInfoResp) String() string            { return proto.CompactTextString(m) }
func (*RpbGetServerInfoResp) ProtoMessage()               {}
func (*RpbGetServerInfoResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RpbGetServerInfoResp) GetNode() []byte {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *RpbGetServerInfoResp) GetServerVersion() []byte {
	if m != nil {
		return m.ServerVersion
	}
	return nil
}

// Key/value pair - used for user metadata, indexes, search doc fields
type RpbPair struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbPair) Reset()                    { *m = RpbPair{} }
func (m *RpbPair) String() string            { return proto.CompactTextString(m) }
func (*RpbPair) ProtoMessage()               {}
func (*RpbPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RpbPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Get bucket properties request
type RpbGetBucketReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Type             []byte `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetBucketReq) Reset()                    { *m = RpbGetBucketReq{} }
func (m *RpbGetBucketReq) String() string            { return proto.CompactTextString(m) }
func (*RpbGetBucketReq) ProtoMessage()               {}
func (*RpbGetBucketReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RpbGetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbGetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Get bucket properties response
type RpbGetBucketResp struct {
	Props            *RpbBucketProps `protobuf:"bytes,1,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbGetBucketResp) Reset()                    { *m = RpbGetBucketResp{} }
func (m *RpbGetBucketResp) String() string            { return proto.CompactTextString(m) }
func (*RpbGetBucketResp) ProtoMessage()               {}
func (*RpbGetBucketResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RpbGetBucketResp) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

// Set bucket properties request
type RpbSetBucketReq struct {
	Bucket           []byte          `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Props            *RpbBucketProps `protobuf:"bytes,2,req,name=props" json:"props,omitempty"`
	Type             []byte          `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSetBucketReq) Reset()                    { *m = RpbSetBucketReq{} }
func (m *RpbSetBucketReq) String() string            { return proto.CompactTextString(m) }
func (*RpbSetBucketReq) ProtoMessage()               {}
func (*RpbSetBucketReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RpbSetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbSetBucketReq) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *RpbSetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Reset bucket properties request
type RpbResetBucketReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Type             []byte `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbResetBucketReq) Reset()                    { *m = RpbResetBucketReq{} }
func (m *RpbResetBucketReq) String() string            { return proto.CompactTextString(m) }
func (*RpbResetBucketReq) ProtoMessage()               {}
func (*RpbResetBucketReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RpbResetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbResetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Get bucket properties request
type RpbGetBucketTypeReq struct {
	Type             []byte `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetBucketTypeReq) Reset()                    { *m = RpbGetBucketTypeReq{} }
func (m *RpbGetBucketTypeReq) String() string            { return proto.CompactTextString(m) }
func (*RpbGetBucketTypeReq) ProtoMessage()               {}
func (*RpbGetBucketTypeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RpbGetBucketTypeReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Set bucket properties request
type RpbSetBucketTypeReq struct {
	Type             []byte          `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Props            *RpbBucketProps `protobuf:"bytes,2,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSetBucketTypeReq) Reset()                    { *m = RpbSetBucketTypeReq{} }
func (m *RpbSetBucketTypeReq) String() string            { return proto.CompactTextString(m) }
func (*RpbSetBucketTypeReq) ProtoMessage()               {}
func (*RpbSetBucketTypeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RpbSetBucketTypeReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *RpbSetBucketTypeReq) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

// Module-Function pairs for commit hooks and other bucket properties
// that take functions
type RpbModFun struct {
	Module           []byte `protobuf:"bytes,1,req,name=module" json:"module,omitempty"`
	Function         []byte `protobuf:"bytes,2,req,name=function" json:"function,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbModFun) Reset()                    { *m = RpbModFun{} }
func (m *RpbModFun) String() string            { return proto.CompactTextString(m) }
func (*RpbModFun) ProtoMessage()               {}
func (*RpbModFun) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RpbModFun) GetModule() []byte {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *RpbModFun) GetFunction() []byte {
	if m != nil {
		return m.Function
	}
	return nil
}

// A commit hook, which may either be a modfun or a JavaScript named
// function
type RpbCommitHook struct {
	Modfun           *RpbModFun `protobuf:"bytes,1,opt,name=modfun" json:"modfun,omitempty"`
	Name             []byte     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RpbCommitHook) Reset()                    { *m = RpbCommitHook{} }
func (m *RpbCommitHook) String() string            { return proto.CompactTextString(m) }
func (*RpbCommitHook) ProtoMessage()               {}
func (*RpbCommitHook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RpbCommitHook) GetModfun() *RpbModFun {
	if m != nil {
		return m.Modfun
	}
	return nil
}

func (m *RpbCommitHook) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

// Bucket properties
type RpbBucketProps struct {
	// Declared in riak_core_app
	NVal          *uint32          `protobuf:"varint,1,opt,name=n_val" json:"n_val,omitempty"`
	AllowMult     *bool            `protobuf:"varint,2,opt,name=allow_mult" json:"allow_mult,omitempty"`
	LastWriteWins *bool            `protobuf:"varint,3,opt,name=last_write_wins" json:"last_write_wins,omitempty"`
	Precommit     []*RpbCommitHook `protobuf:"bytes,4,rep,name=precommit" json:"precommit,omitempty"`
	HasPrecommit  *bool            `protobuf:"varint,5,opt,name=has_precommit,def=0" json:"has_precommit,omitempty"`
	Postcommit    []*RpbCommitHook `protobuf:"bytes,6,rep,name=postcommit" json:"postcommit,omitempty"`
	HasPostcommit *bool            `protobuf:"varint,7,opt,name=has_postcommit,def=0" json:"has_postcommit,omitempty"`
	ChashKeyfun   *RpbModFun       `protobuf:"bytes,8,opt,name=chash_keyfun" json:"chash_keyfun,omitempty"`
	// Declared in riak_kv_app
	Linkfun     *RpbModFun `protobuf:"bytes,9,opt,name=linkfun" json:"linkfun,omitempty"`
	OldVclock   *uint32    `protobuf:"varint,10,opt,name=old_vclock" json:"old_vclock,omitempty"`
	YoungVclock *uint32    `protobuf:"varint,11,opt,name=young_vclock" json:"young_vclock,omitempty"`
	BigVclock   *uint32    `protobuf:"varint,12,opt,name=big_vclock" json:"big_vclock,omitempty"`
	SmallVclock *uint32    `protobuf:"varint,13,opt,name=small_vclock" json:"small_vclock,omitempty"`
	Pr          *uint32    `protobuf:"varint,14,opt,name=pr" json:"pr,omitempty"`
	R           *uint32    `protobuf:"varint,15,opt,name=r" json:"r,omitempty"`
	W           *uint32    `protobuf:"varint,16,opt,name=w" json:"w,omitempty"`
	Pw          *uint32    `protobuf:"varint,17,opt,name=pw" json:"pw,omitempty"`
	Dw          *uint32    `protobuf:"varint,18,opt,name=dw" json:"dw,omitempty"`
	Rw          *uint32    `protobuf:"varint,19,opt,name=rw" json:"rw,omitempty"`
	BasicQuorum *bool      `protobuf:"varint,20,opt,name=basic_quorum" json:"basic_quorum,omitempty"`
	NotfoundOk  *bool      `protobuf:"varint,21,opt,name=notfound_ok" json:"notfound_ok,omitempty"`
	// Used by riak_kv_multi_backend
	Backend []byte `protobuf:"bytes,22,opt,name=backend" json:"backend,omitempty"`
	// Used by riak_search bucket fixup
	Search *bool                       `protobuf:"varint,23,opt,name=search" json:"search,omitempty"`
	Repl   *RpbBucketProps_RpbReplMode `protobuf:"varint,24,opt,name=repl,enum=RpbBucketProps_RpbReplMode" json:"repl,omitempty"`
	// Search index
	SearchIndex []byte `protobuf:"bytes,25,opt,name=search_index" json:"search_index,omitempty"`
	// KV Datatypes
	Datatype []byte `protobuf:"bytes,26,opt,name=datatype" json:"datatype,omitempty"`
	// KV strong consistency
	Consistent *bool `protobuf:"varint,27,opt,name=consistent" json:"consistent,omitempty"`
	// KV fast path
	WriteOnce        *bool  `protobuf:"varint,28,opt,name=write_once" json:"write_once,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbBucketProps) Reset()                    { *m = RpbBucketProps{} }
func (m *RpbBucketProps) String() string            { return proto.CompactTextString(m) }
func (*RpbBucketProps) ProtoMessage()               {}
func (*RpbBucketProps) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

const Default_RpbBucketProps_HasPrecommit bool = false
const Default_RpbBucketProps_HasPostcommit bool = false

func (m *RpbBucketProps) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

func (m *RpbBucketProps) GetAllowMult() bool {
	if m != nil && m.AllowMult != nil {
		return *m.AllowMult
	}
	return false
}

func (m *RpbBucketProps) GetLastWriteWins() bool {
	if m != nil && m.LastWriteWins != nil {
		return *m.LastWriteWins
	}
	return false
}

func (m *RpbBucketProps) GetPrecommit() []*RpbCommitHook {
	if m != nil {
		return m.Precommit
	}
	return nil
}

func (m *RpbBucketProps) GetHasPrecommit() bool {
	if m != nil && m.HasPrecommit != nil {
		return *m.HasPrecommit
	}
	return Default_RpbBucketProps_HasPrecommit
}

func (m *RpbBucketProps) GetPostcommit() []*RpbCommitHook {
	if m != nil {
		return m.Postcommit
	}
	return nil
}

func (m *RpbBucketProps) GetHasPostcommit() bool {
	if m != nil && m.HasPostcommit != nil {
		return *m.HasPostcommit
	}
	return Default_RpbBucketProps_HasPostcommit
}

func (m *RpbBucketProps) GetChashKeyfun() *RpbModFun {
	if m != nil {
		return m.ChashKeyfun
	}
	return nil
}

func (m *RpbBucketProps) GetLinkfun() *RpbModFun {
	if m != nil {
		return m.Linkfun
	}
	return nil
}

func (m *RpbBucketProps) GetOldVclock() uint32 {
	if m != nil && m.OldVclock != nil {
		return *m.OldVclock
	}
	return 0
}

func (m *RpbBucketProps) GetYoungVclock() uint32 {
	if m != nil && m.YoungVclock != nil {
		return *m.YoungVclock
	}
	return 0
}

func (m *RpbBucketProps) GetBigVclock() uint32 {
	if m != nil && m.BigVclock != nil {
		return *m.BigVclock
	}
	return 0
}

func (m *RpbBucketProps) GetSmallVclock() uint32 {
	if m != nil && m.SmallVclock != nil {
		return *m.SmallVclock
	}
	return 0
}

func (m *RpbBucketProps) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbBucketProps) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbBucketProps) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbBucketProps) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbBucketProps) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbBucketProps) GetRw() uint32 {
	if m != nil && m.Rw != nil {
		return *m.Rw
	}
	return 0
}

func (m *RpbBucketProps) GetBasicQuorum() bool {
	if m != nil && m.BasicQuorum != nil {
		return *m.BasicQuorum
	}
	return false
}

func (m *RpbBucketProps) GetNotfoundOk() bool {
	if m != nil && m.NotfoundOk != nil {
		return *m.NotfoundOk
	}
	return false
}

func (m *RpbBucketProps) GetBackend() []byte {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *RpbBucketProps) GetSearch() bool {
	if m != nil && m.Search != nil {
		return *m.Search
	}
	return false
}

func (m *RpbBucketProps) GetRepl() RpbBucketProps_RpbReplMode {
	if m != nil && m.Repl != nil {
		return *m.Repl
	}
	return RpbBucketProps_FALSE
}

func (m *RpbBucketProps) GetSearchIndex() []byte {
	if m != nil {
		return m.SearchIndex
	}
	return nil
}

func (m *RpbBucketProps) GetDatatype() []byte {
	if m != nil {
		return m.Datatype
	}
	return nil
}

func (m *RpbBucketProps) GetConsistent() bool {
	if m != nil && m.Consistent != nil {
		return *m.Consistent
	}
	return false
}

func (m *RpbBucketProps) GetWriteOnce() bool {
	if m != nil && m.WriteOnce != nil {
		return *m.WriteOnce
	}
	return false
}

// Authentication request
type RpbAuthReq struct {
	User             []byte `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Password         []byte `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbAuthReq) Reset()                    { *m = RpbAuthReq{} }
func (m *RpbAuthReq) String() string            { return proto.CompactTextString(m) }
func (*RpbAuthReq) ProtoMessage()               {}
func (*RpbAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *RpbAuthReq) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RpbAuthReq) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

// Erlang term encoding
type RpbToggleEncodingReq struct {
	UseNative        *bool  `protobuf:"varint,1,req,name=use_native" json:"use_native,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbToggleEncodingReq) Reset()                    { *m = RpbToggleEncodingReq{} }
func (m *RpbToggleEncodingReq) String() string            { return proto.CompactTextString(m) }
func (*RpbToggleEncodingReq) ProtoMessage()               {}
func (*RpbToggleEncodingReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RpbToggleEncodingReq) GetUseNative() bool {
	if m != nil && m.UseNative != nil {
		return *m.UseNative
	}
	return false
}

type RpbToggleEncodingResp struct {
	UseNative        *bool  `protobuf:"varint,1,req,name=use_native" json:"use_native,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbToggleEncodingResp) Reset()                    { *m = RpbToggleEncodingResp{} }
func (m *RpbToggleEncodingResp) String() string            { return proto.CompactTextString(m) }
func (*RpbToggleEncodingResp) ProtoMessage()               {}
func (*RpbToggleEncodingResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *RpbToggleEncodingResp) GetUseNative() bool {
	if m != nil && m.UseNative != nil {
		return *m.UseNative
	}
	return false
}

func init() {
	proto.RegisterType((*RpbErrorResp)(nil), "RpbErrorResp")
	proto.RegisterType((*RpbGetServerInfoResp)(nil), "RpbGetServerInfoResp")
	proto.RegisterType((*RpbPair)(nil), "RpbPair")
	proto.RegisterType((*RpbGetBucketReq)(nil), "RpbGetBucketReq")
	proto.RegisterType((*RpbGetBucketResp)(nil), "RpbGetBucketResp")
	proto.RegisterType((*RpbSetBucketReq)(nil), "RpbSetBucketReq")
	proto.RegisterType((*RpbResetBucketReq)(nil), "RpbResetBucketReq")
	proto.RegisterType((*RpbGetBucketTypeReq)(nil), "RpbGetBucketTypeReq")
	proto.RegisterType((*RpbSetBucketTypeReq)(nil), "RpbSetBucketTypeReq")
	proto.RegisterType((*RpbModFun)(nil), "RpbModFun")
	proto.RegisterType((*RpbCommitHook)(nil), "RpbCommitHook")
	proto.RegisterType((*RpbBucketProps)(nil), "RpbBucketProps")
	proto.RegisterType((*RpbAuthReq)(nil), "RpbAuthReq")
	proto.RegisterType((*RpbToggleEncodingReq)(nil), "RpbToggleEncodingReq")
	proto.RegisterType((*RpbToggleEncodingResp)(nil), "RpbToggleEncodingResp")
	proto.RegisterEnum("RpbBucketProps_RpbReplMode", RpbBucketProps_RpbReplMode_name, RpbBucketProps_RpbReplMode_value)
}

var fileDescriptor0 = []byte{
	// 761 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x59, 0x6f, 0xe2, 0x48,
	0x10, 0x5e, 0xae, 0x00, 0xc5, 0x19, 0xe7, 0xea, 0x4d, 0xb2, 0x2b, 0xe2, 0xd5, 0x4a, 0x7b, 0x12,
	0x2d, 0x6f, 0x3b, 0x1a, 0x8d, 0x14, 0x22, 0x32, 0x13, 0x89, 0xcc, 0x44, 0x40, 0x1e, 0xe6, 0xc9,
	0x32, 0x76, 0x03, 0x16, 0x76, 0xb7, 0xd3, 0xdd, 0xc6, 0x93, 0x7f, 0x35, 0x3f, 0x71, 0xca, 0x0d,
	0x86, 0x1c, 0x28, 0x33, 0x0f, 0x96, 0xdc, 0x55, 0xdf, 0x57, 0xf5, 0x75, 0x55, 0x75, 0x01, 0x08,
	0xcf, 0x9e, 0xb7, 0x43, 0xc1, 0x15, 0x37, 0xcf, 0xa1, 0x3a, 0x08, 0xc7, 0x3d, 0x21, 0xb8, 0x18,
	0x50, 0x19, 0x1a, 0x75, 0xd8, 0xa1, 0x42, 0x04, 0x72, 0x4a, 0x32, 0xad, 0xec, 0x1f, 0x55, 0xa3,
	0x01, 0x45, 0x3c, 0x3b, 0xdc, 0xa5, 0x24, 0x8b, 0x86, 0x9a, 0xf9, 0x16, 0xf6, 0x91, 0xf0, 0x9e,
	0xaa, 0x21, 0x15, 0x0b, 0x2a, 0xae, 0xd9, 0x84, 0x6b, 0x62, 0x15, 0xf2, 0x2c, 0x41, 0x65, 0x5a,
	0x19, 0xa4, 0x1d, 0x42, 0x5d, 0x6a, 0xbf, 0x85, 0x9f, 0xf4, 0x38, 0x43, 0x36, 0xda, 0xcd, 0xdf,
	0xa1, 0x88, 0xec, 0x5b, 0xdb, 0x13, 0x46, 0x05, 0x72, 0x73, 0xfa, 0xb0, 0x4a, 0x53, 0x83, 0xc2,
	0xc2, 0xf6, 0x23, 0xba, 0x82, 0x9d, 0x43, 0x63, 0x99, 0xa4, 0x1b, 0x39, 0x73, 0xaa, 0x06, 0xf4,
	0x3e, 0x11, 0x36, 0xd6, 0x87, 0x15, 0x03, 0xf3, 0xa9, 0x87, 0x30, 0x25, 0x74, 0xa0, 0xf9, 0x94,
	0x80, 0x8a, 0x7e, 0x85, 0x02, 0xde, 0x31, 0x94, 0x9a, 0x50, 0xe9, 0x34, 0xda, 0x88, 0x58, 0xba,
	0x6f, 0x13, 0xb3, 0xf9, 0x49, 0x27, 0x19, 0xbe, 0x96, 0x64, 0x1d, 0x22, 0xbb, 0x35, 0xc4, 0x5a,
	0x44, 0x4e, 0x8b, 0xf8, 0x0f, 0x76, 0xd1, 0x8f, 0xb9, 0x7f, 0x5c, 0xf7, 0x6f, 0xb0, 0xf7, 0x58,
	0xf7, 0x08, 0x3d, 0x09, 0x29, 0x05, 0x69, 0x8a, 0x79, 0xa9, 0x41, 0xc3, 0x57, 0x41, 0xdf, 0x93,
	0x6a, 0xfe, 0x0b, 0x65, 0xb4, 0xdc, 0x70, 0xf7, 0x2a, 0x62, 0x89, 0xa8, 0x80, 0xbb, 0x91, 0x9f,
	0x92, 0x9b, 0x50, 0x9a, 0x44, 0xcc, 0x51, 0xcb, 0x46, 0x25, 0x39, 0xff, 0x87, 0x1a, 0xc2, 0x2f,
	0x79, 0x10, 0x78, 0xea, 0x03, 0xe7, 0x73, 0xe3, 0x58, 0x53, 0x10, 0xa5, 0x3b, 0x5c, 0xe9, 0x40,
	0x7b, 0x13, 0x2e, 0xe9, 0xbd, 0x1d, 0xa4, 0x77, 0xfa, 0x5a, 0x80, 0xfa, 0xb3, 0x3a, 0x61, 0x7b,
	0x99, 0x85, 0x0d, 0xd6, 0xdc, 0x9a, 0x61, 0x00, 0xd8, 0xbe, 0xcf, 0x63, 0x2b, 0x88, 0x7c, 0xa5,
	0x59, 0x25, 0xe3, 0x08, 0x1a, 0xbe, 0x2d, 0x95, 0x15, 0x0b, 0x4f, 0x51, 0x2b, 0xf6, 0x98, 0xd4,
	0x55, 0x2d, 0x19, 0x67, 0x50, 0x0e, 0x05, 0x75, 0xb4, 0x12, 0x92, 0x6f, 0xe5, 0x30, 0x77, 0xbd,
	0xfd, 0x54, 0xdb, 0x29, 0xd4, 0x66, 0xb6, 0xb4, 0x36, 0xb0, 0x42, 0xc2, 0x7c, 0x53, 0x98, 0xd8,
	0xbe, 0xa4, 0x86, 0x09, 0x10, 0x72, 0xa9, 0x56, 0xae, 0x9d, 0xad, 0x11, 0x7e, 0x81, 0xba, 0x8e,
	0xb0, 0xc1, 0x15, 0x1f, 0x87, 0x68, 0x41, 0xd5, 0x41, 0xff, 0xcc, 0xc2, 0x89, 0x4d, 0x4a, 0x50,
	0x7a, 0x51, 0x82, 0x13, 0x28, 0xfa, 0x1e, 0x9b, 0x27, 0xce, 0xf2, 0x0b, 0x27, 0xde, 0x97, 0xfb,
	0xae, 0xb5, 0x70, 0x7c, 0xee, 0xcc, 0x09, 0xe8, 0x1a, 0xec, 0x43, 0xf5, 0x81, 0x47, 0x6c, 0x9a,
	0x5a, 0x2b, 0x69, 0x65, 0xc6, 0xde, 0xda, 0x56, 0x4d, 0x91, 0x32, 0xc0, 0x7a, 0xa5, 0xd6, 0x9a,
	0xb6, 0x02, 0x64, 0x43, 0x41, 0xea, 0xfa, 0xbf, 0x0c, 0x19, 0x41, 0x1a, 0xe9, 0x6f, 0x4c, 0x9a,
	0x6b, 0x44, 0x4c, 0x76, 0xd3, 0x7f, 0x37, 0x26, 0x46, 0xfa, 0x2f, 0x62, 0xb2, 0x97, 0xc6, 0x1e,
	0xdb, 0xd2, 0x73, 0xac, 0xfb, 0x88, 0x8b, 0x28, 0x20, 0xfb, 0xba, 0xe4, 0x7b, 0x50, 0x61, 0x5c,
	0x4d, 0x50, 0x9e, 0x6b, 0xf1, 0x39, 0x39, 0xd0, 0x46, 0xdc, 0x04, 0x63, 0x1b, 0x5b, 0xca, 0x5c,
	0x72, 0xa8, 0xdf, 0x38, 0x0e, 0x91, 0xa4, 0xb6, 0x70, 0x66, 0xe4, 0x48, 0x03, 0xfe, 0x84, 0xbc,
	0xa0, 0xa1, 0x4f, 0x08, 0x9e, 0xea, 0x9d, 0x93, 0x67, 0x03, 0xd8, 0xd6, 0x4f, 0x23, 0xf4, 0xb1,
	0x24, 0x54, 0x5f, 0x49, 0x53, 0x2d, 0x8f, 0xb9, 0xf4, 0x0b, 0xf9, 0x59, 0x07, 0xc4, 0x29, 0x74,
	0x6d, 0x65, 0xeb, 0xa1, 0x3e, 0xd6, 0x16, 0xd4, 0xea, 0x70, 0x26, 0x3d, 0xa9, 0x28, 0x53, 0xe4,
	0x44, 0xa7, 0x41, 0xdb, 0x72, 0x46, 0x38, 0x73, 0x28, 0x39, 0x4d, 0x6c, 0xe6, 0x3b, 0xa8, 0x3c,
	0x0e, 0x5f, 0x86, 0xc2, 0xd5, 0x45, 0x7f, 0xd8, 0x6b, 0xfe, 0x84, 0xa3, 0x59, 0x1a, 0xf4, 0x2e,
	0xfa, 0xa3, 0xeb, 0x9b, 0x5e, 0x33, 0x93, 0x9c, 0xae, 0xee, 0xfa, 0xfd, 0xe1, 0xe7, 0x8f, 0x97,
	0xcd, 0xac, 0x51, 0x82, 0xfc, 0x68, 0x70, 0xd7, 0x6b, 0xe6, 0xcc, 0x7f, 0x00, 0x90, 0x7f, 0x11,
	0xa9, 0xd9, 0xea, 0x61, 0x45, 0xb8, 0xbd, 0x36, 0x6f, 0x23, 0xb4, 0xa5, 0x8c, 0xb9, 0x70, 0x57,
	0x6f, 0xe3, 0x2f, 0xbd, 0x02, 0x47, 0x7c, 0x3a, 0xf5, 0x69, 0x8f, 0xe1, 0x6e, 0xf4, 0xd8, 0x34,
	0xe1, 0xa1, 0x32, 0xe4, 0x59, 0xcc, 0x56, 0xde, 0x62, 0xf9, 0xb2, 0x4a, 0xe6, 0xdf, 0x70, 0xb0,
	0x05, 0x8b, 0xdb, 0x69, 0x0b, 0xb8, 0x7b, 0x06, 0x47, 0x38, 0x7d, 0x6d, 0xec, 0xc8, 0x8c, 0xb7,
	0x37, 0x4b, 0x7a, 0x1c, 0x4d, 0xba, 0x3b, 0x03, 0x3c, 0xde, 0x76, 0xbf, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0x77, 0x00, 0x9e, 0xbc, 0x05, 0x00, 0x00,
}
