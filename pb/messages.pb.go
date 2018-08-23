// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateVMOpts struct {
	StemcellID           string   `protobuf:"bytes,1,opt,name=stemcellID,proto3" json:"stemcellID,omitempty"`
	AgentSettings        []byte   `protobuf:"bytes,2,opt,name=agentSettings,proto3" json:"agentSettings,omitempty"`
	DiskID               string   `protobuf:"bytes,3,opt,name=diskID,proto3" json:"diskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVMOpts) Reset()         { *m = CreateVMOpts{} }
func (m *CreateVMOpts) String() string { return proto.CompactTextString(m) }
func (*CreateVMOpts) ProtoMessage()    {}
func (*CreateVMOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{0}
}
func (m *CreateVMOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVMOpts.Unmarshal(m, b)
}
func (m *CreateVMOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVMOpts.Marshal(b, m, deterministic)
}
func (dst *CreateVMOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVMOpts.Merge(dst, src)
}
func (m *CreateVMOpts) XXX_Size() int {
	return xxx_messageInfo_CreateVMOpts.Size(m)
}
func (m *CreateVMOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVMOpts.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVMOpts proto.InternalMessageInfo

func (m *CreateVMOpts) GetStemcellID() string {
	if m != nil {
		return m.StemcellID
	}
	return ""
}

func (m *CreateVMOpts) GetAgentSettings() []byte {
	if m != nil {
		return m.AgentSettings
	}
	return nil
}

func (m *CreateVMOpts) GetDiskID() string {
	if m != nil {
		return m.DiskID
	}
	return ""
}

type DisksOpts struct {
	VmID                 string   `protobuf:"bytes,1,opt,name=vmID,proto3" json:"vmID,omitempty"`
	DiskID               string   `protobuf:"bytes,2,opt,name=diskID,proto3" json:"diskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisksOpts) Reset()         { *m = DisksOpts{} }
func (m *DisksOpts) String() string { return proto.CompactTextString(m) }
func (*DisksOpts) ProtoMessage()    {}
func (*DisksOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{1}
}
func (m *DisksOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisksOpts.Unmarshal(m, b)
}
func (m *DisksOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisksOpts.Marshal(b, m, deterministic)
}
func (dst *DisksOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisksOpts.Merge(dst, src)
}
func (m *DisksOpts) XXX_Size() int {
	return xxx_messageInfo_DisksOpts.Size(m)
}
func (m *DisksOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_DisksOpts.DiscardUnknown(m)
}

var xxx_messageInfo_DisksOpts proto.InternalMessageInfo

func (m *DisksOpts) GetVmID() string {
	if m != nil {
		return m.VmID
	}
	return ""
}

func (m *DisksOpts) GetDiskID() string {
	if m != nil {
		return m.DiskID
	}
	return ""
}

type VMFilterOpts struct {
	VmID                 string   `protobuf:"bytes,1,opt,name=vmID,proto3" json:"vmID,omitempty"`
	All                  bool     `protobuf:"varint,2,opt,name=all,proto3" json:"all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMFilterOpts) Reset()         { *m = VMFilterOpts{} }
func (m *VMFilterOpts) String() string { return proto.CompactTextString(m) }
func (*VMFilterOpts) ProtoMessage()    {}
func (*VMFilterOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{2}
}
func (m *VMFilterOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMFilterOpts.Unmarshal(m, b)
}
func (m *VMFilterOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMFilterOpts.Marshal(b, m, deterministic)
}
func (dst *VMFilterOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMFilterOpts.Merge(dst, src)
}
func (m *VMFilterOpts) XXX_Size() int {
	return xxx_messageInfo_VMFilterOpts.Size(m)
}
func (m *VMFilterOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_VMFilterOpts.DiscardUnknown(m)
}

var xxx_messageInfo_VMFilterOpts proto.InternalMessageInfo

func (m *VMFilterOpts) GetVmID() string {
	if m != nil {
		return m.VmID
	}
	return ""
}

func (m *VMFilterOpts) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

type DataParcel struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataParcel) Reset()         { *m = DataParcel{} }
func (m *DataParcel) String() string { return proto.CompactTextString(m) }
func (*DataParcel) ProtoMessage()    {}
func (*DataParcel) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{3}
}
func (m *DataParcel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataParcel.Unmarshal(m, b)
}
func (m *DataParcel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataParcel.Marshal(b, m, deterministic)
}
func (dst *DataParcel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataParcel.Merge(dst, src)
}
func (m *DataParcel) XXX_Size() int {
	return xxx_messageInfo_DataParcel.Size(m)
}
func (m *DataParcel) XXX_DiscardUnknown() {
	xxx_messageInfo_DataParcel.DiscardUnknown(m)
}

var xxx_messageInfo_DataParcel proto.InternalMessageInfo

func (m *DataParcel) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type TruthParcel struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TruthParcel) Reset()         { *m = TruthParcel{} }
func (m *TruthParcel) String() string { return proto.CompactTextString(m) }
func (*TruthParcel) ProtoMessage()    {}
func (*TruthParcel) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{4}
}
func (m *TruthParcel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TruthParcel.Unmarshal(m, b)
}
func (m *TruthParcel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TruthParcel.Marshal(b, m, deterministic)
}
func (dst *TruthParcel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TruthParcel.Merge(dst, src)
}
func (m *TruthParcel) XXX_Size() int {
	return xxx_messageInfo_TruthParcel.Size(m)
}
func (m *TruthParcel) XXX_DiscardUnknown() {
	xxx_messageInfo_TruthParcel.DiscardUnknown(m)
}

var xxx_messageInfo_TruthParcel proto.InternalMessageInfo

func (m *TruthParcel) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type NumberParcel struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberParcel) Reset()         { *m = NumberParcel{} }
func (m *NumberParcel) String() string { return proto.CompactTextString(m) }
func (*NumberParcel) ProtoMessage()    {}
func (*NumberParcel) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{5}
}
func (m *NumberParcel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberParcel.Unmarshal(m, b)
}
func (m *NumberParcel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberParcel.Marshal(b, m, deterministic)
}
func (dst *NumberParcel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberParcel.Merge(dst, src)
}
func (m *NumberParcel) XXX_Size() int {
	return xxx_messageInfo_NumberParcel.Size(m)
}
func (m *NumberParcel) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberParcel.DiscardUnknown(m)
}

var xxx_messageInfo_NumberParcel proto.InternalMessageInfo

func (m *NumberParcel) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TextParcel struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextParcel) Reset()         { *m = TextParcel{} }
func (m *TextParcel) String() string { return proto.CompactTextString(m) }
func (*TextParcel) ProtoMessage()    {}
func (*TextParcel) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{6}
}
func (m *TextParcel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextParcel.Unmarshal(m, b)
}
func (m *TextParcel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextParcel.Marshal(b, m, deterministic)
}
func (dst *TextParcel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextParcel.Merge(dst, src)
}
func (m *TextParcel) XXX_Size() int {
	return xxx_messageInfo_TextParcel.Size(m)
}
func (m *TextParcel) XXX_DiscardUnknown() {
	xxx_messageInfo_TextParcel.DiscardUnknown(m)
}

var xxx_messageInfo_TextParcel proto.InternalMessageInfo

func (m *TextParcel) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_65267586716775b5, []int{7}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (dst *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(dst, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateVMOpts)(nil), "pb.CreateVMOpts")
	proto.RegisterType((*DisksOpts)(nil), "pb.DisksOpts")
	proto.RegisterType((*VMFilterOpts)(nil), "pb.VMFilterOpts")
	proto.RegisterType((*DataParcel)(nil), "pb.DataParcel")
	proto.RegisterType((*TruthParcel)(nil), "pb.TruthParcel")
	proto.RegisterType((*NumberParcel)(nil), "pb.NumberParcel")
	proto.RegisterType((*TextParcel)(nil), "pb.TextParcel")
	proto.RegisterType((*Void)(nil), "pb.Void")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CPIDClient is the client API for CPID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CPIDClient interface {
	Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TextParcel, error)
	CreateStemcell(ctx context.Context, opts ...grpc.CallOption) (CPID_CreateStemcellClient, error)
	DeleteStemcell(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*Void, error)
	DeleteDisk(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*Void, error)
	CreateDisk(ctx context.Context, in *NumberParcel, opts ...grpc.CallOption) (*TextParcel, error)
	AttachDisk(ctx context.Context, in *DisksOpts, opts ...grpc.CallOption) (*Void, error)
	DetachDisk(ctx context.Context, in *DisksOpts, opts ...grpc.CallOption) (*Void, error)
	HasDisk(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*TruthParcel, error)
	CreateVM(ctx context.Context, in *CreateVMOpts, opts ...grpc.CallOption) (*TextParcel, error)
	DeleteVM(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*Void, error)
	CheckpointVM(ctx context.Context, in *VMFilterOpts, opts ...grpc.CallOption) (*Void, error)
	RestoreVM(ctx context.Context, in *VMFilterOpts, opts ...grpc.CallOption) (*Void, error)
	HasVM(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*TruthParcel, error)
	ShellExec(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*TextParcel, error)
}

type cPIDClient struct {
	cc *grpc.ClientConn
}

func NewCPIDClient(cc *grpc.ClientConn) CPIDClient {
	return &cPIDClient{cc}
}

func (c *cPIDClient) Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TextParcel, error) {
	out := new(TextParcel)
	err := c.cc.Invoke(ctx, "/pb.CPID/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) CreateStemcell(ctx context.Context, opts ...grpc.CallOption) (CPID_CreateStemcellClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CPID_serviceDesc.Streams[0], "/pb.CPID/CreateStemcell", opts...)
	if err != nil {
		return nil, err
	}
	x := &cPIDCreateStemcellClient{stream}
	return x, nil
}

type CPID_CreateStemcellClient interface {
	Send(*DataParcel) error
	CloseAndRecv() (*TextParcel, error)
	grpc.ClientStream
}

type cPIDCreateStemcellClient struct {
	grpc.ClientStream
}

func (x *cPIDCreateStemcellClient) Send(m *DataParcel) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cPIDCreateStemcellClient) CloseAndRecv() (*TextParcel, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TextParcel)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cPIDClient) DeleteStemcell(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.CPID/DeleteStemcell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) DeleteDisk(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.CPID/DeleteDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) CreateDisk(ctx context.Context, in *NumberParcel, opts ...grpc.CallOption) (*TextParcel, error) {
	out := new(TextParcel)
	err := c.cc.Invoke(ctx, "/pb.CPID/CreateDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) AttachDisk(ctx context.Context, in *DisksOpts, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.CPID/AttachDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) DetachDisk(ctx context.Context, in *DisksOpts, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.CPID/DetachDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) HasDisk(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*TruthParcel, error) {
	out := new(TruthParcel)
	err := c.cc.Invoke(ctx, "/pb.CPID/HasDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) CreateVM(ctx context.Context, in *CreateVMOpts, opts ...grpc.CallOption) (*TextParcel, error) {
	out := new(TextParcel)
	err := c.cc.Invoke(ctx, "/pb.CPID/CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) DeleteVM(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.CPID/DeleteVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) CheckpointVM(ctx context.Context, in *VMFilterOpts, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.CPID/CheckpointVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) RestoreVM(ctx context.Context, in *VMFilterOpts, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pb.CPID/RestoreVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) HasVM(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*TruthParcel, error) {
	out := new(TruthParcel)
	err := c.cc.Invoke(ctx, "/pb.CPID/HasVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPIDClient) ShellExec(ctx context.Context, in *TextParcel, opts ...grpc.CallOption) (*TextParcel, error) {
	out := new(TextParcel)
	err := c.cc.Invoke(ctx, "/pb.CPID/ShellExec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CPIDServer is the server API for CPID service.
type CPIDServer interface {
	Ping(context.Context, *Void) (*TextParcel, error)
	CreateStemcell(CPID_CreateStemcellServer) error
	DeleteStemcell(context.Context, *TextParcel) (*Void, error)
	DeleteDisk(context.Context, *TextParcel) (*Void, error)
	CreateDisk(context.Context, *NumberParcel) (*TextParcel, error)
	AttachDisk(context.Context, *DisksOpts) (*Void, error)
	DetachDisk(context.Context, *DisksOpts) (*Void, error)
	HasDisk(context.Context, *TextParcel) (*TruthParcel, error)
	CreateVM(context.Context, *CreateVMOpts) (*TextParcel, error)
	DeleteVM(context.Context, *TextParcel) (*Void, error)
	CheckpointVM(context.Context, *VMFilterOpts) (*Void, error)
	RestoreVM(context.Context, *VMFilterOpts) (*Void, error)
	HasVM(context.Context, *TextParcel) (*TruthParcel, error)
	ShellExec(context.Context, *TextParcel) (*TextParcel, error)
}

func RegisterCPIDServer(s *grpc.Server, srv CPIDServer) {
	s.RegisterService(&_CPID_serviceDesc, srv)
}

func _CPID_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).Ping(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_CreateStemcell_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CPIDServer).CreateStemcell(&cPIDCreateStemcellServer{stream})
}

type CPID_CreateStemcellServer interface {
	SendAndClose(*TextParcel) error
	Recv() (*DataParcel, error)
	grpc.ServerStream
}

type cPIDCreateStemcellServer struct {
	grpc.ServerStream
}

func (x *cPIDCreateStemcellServer) SendAndClose(m *TextParcel) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cPIDCreateStemcellServer) Recv() (*DataParcel, error) {
	m := new(DataParcel)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CPID_DeleteStemcell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextParcel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).DeleteStemcell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/DeleteStemcell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).DeleteStemcell(ctx, req.(*TextParcel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_DeleteDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextParcel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).DeleteDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/DeleteDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).DeleteDisk(ctx, req.(*TextParcel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_CreateDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberParcel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).CreateDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/CreateDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).CreateDisk(ctx, req.(*NumberParcel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_AttachDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisksOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).AttachDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/AttachDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).AttachDisk(ctx, req.(*DisksOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_DetachDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisksOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).DetachDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/DetachDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).DetachDisk(ctx, req.(*DisksOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_HasDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextParcel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).HasDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/HasDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).HasDisk(ctx, req.(*TextParcel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVMOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).CreateVM(ctx, req.(*CreateVMOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_DeleteVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextParcel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).DeleteVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/DeleteVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).DeleteVM(ctx, req.(*TextParcel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_CheckpointVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMFilterOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).CheckpointVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/CheckpointVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).CheckpointVM(ctx, req.(*VMFilterOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_RestoreVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMFilterOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).RestoreVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/RestoreVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).RestoreVM(ctx, req.(*VMFilterOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_HasVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextParcel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).HasVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/HasVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).HasVM(ctx, req.(*TextParcel))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPID_ShellExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextParcel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIDServer).ShellExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CPID/ShellExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIDServer).ShellExec(ctx, req.(*TextParcel))
	}
	return interceptor(ctx, in, info, handler)
}

var _CPID_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CPID",
	HandlerType: (*CPIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CPID_Ping_Handler,
		},
		{
			MethodName: "DeleteStemcell",
			Handler:    _CPID_DeleteStemcell_Handler,
		},
		{
			MethodName: "DeleteDisk",
			Handler:    _CPID_DeleteDisk_Handler,
		},
		{
			MethodName: "CreateDisk",
			Handler:    _CPID_CreateDisk_Handler,
		},
		{
			MethodName: "AttachDisk",
			Handler:    _CPID_AttachDisk_Handler,
		},
		{
			MethodName: "DetachDisk",
			Handler:    _CPID_DetachDisk_Handler,
		},
		{
			MethodName: "HasDisk",
			Handler:    _CPID_HasDisk_Handler,
		},
		{
			MethodName: "CreateVM",
			Handler:    _CPID_CreateVM_Handler,
		},
		{
			MethodName: "DeleteVM",
			Handler:    _CPID_DeleteVM_Handler,
		},
		{
			MethodName: "CheckpointVM",
			Handler:    _CPID_CheckpointVM_Handler,
		},
		{
			MethodName: "RestoreVM",
			Handler:    _CPID_RestoreVM_Handler,
		},
		{
			MethodName: "HasVM",
			Handler:    _CPID_HasVM_Handler,
		},
		{
			MethodName: "ShellExec",
			Handler:    _CPID_ShellExec_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateStemcell",
			Handler:       _CPID_CreateStemcell_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_messages_65267586716775b5) }

var fileDescriptor_messages_65267586716775b5 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x05, 0x31, 0xd4, 0x7e, 0x75, 0x68, 0xb4, 0xaa, 0x2a, 0x94, 0x43, 0x85, 0x5c, 0xd2,
	0x22, 0x5a, 0x71, 0x48, 0x2b, 0xf5, 0x5c, 0xe1, 0x56, 0xe1, 0x40, 0x8b, 0x4c, 0xc4, 0x7d, 0x71,
	0x46, 0xd8, 0x62, 0xfd, 0x47, 0xde, 0x25, 0xca, 0x67, 0xed, 0xa7, 0xa9, 0xbc, 0x8e, 0x93, 0x0d,
	0x76, 0x93, 0xdb, 0xce, 0xf0, 0x7b, 0x33, 0x6f, 0x98, 0x31, 0x06, 0x09, 0x49, 0xc9, 0x77, 0x24,
	0x67, 0x79, 0x91, 0xa9, 0x8c, 0x75, 0xf3, 0xad, 0x27, 0xe0, 0xce, 0x0b, 0xe2, 0x8a, 0x36, 0xcb,
	0x3f, 0xb9, 0x92, 0xec, 0x3d, 0x20, 0x15, 0x25, 0x21, 0x09, 0xb1, 0xf0, 0x87, 0x9d, 0x51, 0x67,
	0xe2, 0x04, 0x46, 0x86, 0x8d, 0x71, 0xca, 0x77, 0x94, 0xaa, 0x35, 0x29, 0x15, 0xa7, 0x3b, 0x39,
	0xec, 0x8e, 0x3a, 0x13, 0x37, 0x78, 0x9a, 0x64, 0xef, 0xd0, 0xbf, 0x89, 0xe5, 0x7e, 0xe1, 0x0f,
	0x4f, 0x74, 0x85, 0xfb, 0xc8, 0xfb, 0x0e, 0xc7, 0x8f, 0xe5, 0x5e, 0xea, 0x56, 0x0c, 0xd6, 0x6d,
	0xf2, 0xd0, 0x44, 0xbf, 0x0d, 0x61, 0xf7, 0x89, 0xf0, 0x1b, 0xdc, 0xcd, 0xf2, 0x57, 0x2c, 0x14,
	0x15, 0xff, 0xd5, 0x9e, 0xe1, 0x84, 0x0b, 0xa1, 0x85, 0x76, 0x50, 0x3e, 0x3d, 0x0f, 0xf0, 0xb9,
	0xe2, 0x2b, 0x5e, 0x84, 0x24, 0xd8, 0x5b, 0xf4, 0x6e, 0xb9, 0x38, 0x90, 0x16, 0xb9, 0x41, 0x15,
	0x78, 0x1f, 0xf0, 0xfa, 0xba, 0x38, 0xa8, 0xa8, 0x0d, 0xb2, 0x6b, 0x68, 0x0c, 0xf7, 0xf7, 0x21,
	0xd9, 0x52, 0xd1, 0x46, 0xf5, 0x6a, 0xca, 0x03, 0xae, 0xe9, 0x4e, 0xb5, 0x31, 0x4e, 0xcd, 0xf4,
	0x61, 0x6d, 0xb2, 0xf8, 0xe6, 0xf2, 0xaf, 0x05, 0x6b, 0xbe, 0x5a, 0xf8, 0x6c, 0x04, 0x6b, 0x15,
	0xa7, 0x3b, 0x66, 0xcf, 0xf2, 0xed, 0xac, 0xfc, 0xe9, 0x7c, 0x50, 0xbe, 0x8c, 0x42, 0x97, 0x18,
	0x54, 0x2b, 0x5a, 0xdf, 0xaf, 0x81, 0x69, 0xe2, 0x71, 0xb2, 0x63, 0xc5, 0xa4, 0xc3, 0xa6, 0x18,
	0xf8, 0x24, 0xe8, 0x58, 0xf3, 0xc8, 0x9c, 0x3f, 0xf4, 0x63, 0x1f, 0x81, 0x8a, 0x2d, 0x57, 0xf3,
	0x0c, 0x37, 0x03, 0x2a, 0x1f, 0x9a, 0x3b, 0x2b, 0xf3, 0xe6, 0x9f, 0xd2, 0xf0, 0x7d, 0x01, 0xfc,
	0x50, 0x8a, 0x87, 0x91, 0xe6, 0x4f, 0xb5, 0xe7, 0x7a, 0xf9, 0x46, 0xd9, 0x8b, 0xb2, 0xfd, 0xcb,
	0xd8, 0x14, 0xaf, 0xae, 0xb8, 0x6c, 0xb5, 0xf8, 0x46, 0xc7, 0xc6, 0x12, 0xbf, 0xc0, 0xae, 0x8f,
	0xba, 0xf2, 0x69, 0x9e, 0x78, 0xc3, 0xe7, 0x18, 0x76, 0x35, 0xff, 0x66, 0xf9, 0xcc, 0xf4, 0x53,
	0xb8, 0xf3, 0x88, 0xc2, 0x7d, 0x9e, 0xc5, 0xa9, 0xaa, 0xeb, 0x9a, 0x37, 0x69, 0xb0, 0x9f, 0xe0,
	0x04, 0x24, 0x55, 0x56, 0xd0, 0x0b, 0xe0, 0x04, 0xbd, 0x2b, 0x2e, 0x5b, 0xfa, 0x36, 0x46, 0xfa,
	0x0c, 0x67, 0x1d, 0x91, 0x10, 0x3f, 0xef, 0x28, 0x6c, 0xd0, 0x47, 0xf1, 0xb6, 0xaf, 0xbf, 0xef,
	0xaf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x4d, 0x75, 0x85, 0xf1, 0x03, 0x00, 0x00,
}
