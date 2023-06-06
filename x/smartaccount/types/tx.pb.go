// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: smartaccount/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_CosmWasm_wasmd_x_wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateAccount struct {
	// Sender is the actor who signs the message
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// CodeID indicates which wasm binary code is to be used for this contract
	CodeID uint64 `protobuf:"varint,2,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	// InitMsg is the JSON-encoded instantiate message for the contract
	InitMsg github_com_CosmWasm_wasmd_x_wasm_types.RawContractMessage `protobuf:"bytes,3,opt,name=init_msg,json=initMsg,proto3,casttype=github.com/CosmWasm/wasmd/x/wasm/types.RawContractMessage" json:"init_msg,omitempty"`
	// PubKey using for signature verification of this account
	PubKey secp256k1.PubKey `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key"`
	// Funds are coins to be deposited to the contract on instantiattion
	Funds github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=funds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"funds"`
	// Salt is an arbinary value to be used in deriving the account address.
	// Max 64 bytes.
	Salt []byte `protobuf:"bytes,6,opt,name=salt,proto3" json:"salt,omitempty"`
}

func (m *MsgCreateAccount) Reset()         { *m = MsgCreateAccount{} }
func (m *MsgCreateAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAccount) ProtoMessage()    {}
func (*MsgCreateAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92db8b72563f1f, []int{0}
}
func (m *MsgCreateAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAccount.Merge(m, src)
}
func (m *MsgCreateAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAccount proto.InternalMessageInfo

type MsgCreateAccountResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgCreateAccountResponse) Reset()         { *m = MsgCreateAccountResponse{} }
func (m *MsgCreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAccountResponse) ProtoMessage()    {}
func (*MsgCreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92db8b72563f1f, []int{1}
}
func (m *MsgCreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAccountResponse.Merge(m, src)
}
func (m *MsgCreateAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAccountResponse proto.InternalMessageInfo

func (m *MsgCreateAccountResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgCreateAccountResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MsgUpdateKey struct {
	// Sender is the actor who signs the message
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// smart-account address that need update
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// PubKey using for signature verification of this account
	PubKey secp256k1.PubKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key"`
}

func (m *MsgUpdateKey) Reset()         { *m = MsgUpdateKey{} }
func (m *MsgUpdateKey) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateKey) ProtoMessage()    {}
func (*MsgUpdateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92db8b72563f1f, []int{2}
}
func (m *MsgUpdateKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateKey.Merge(m, src)
}
func (m *MsgUpdateKey) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateKey proto.InternalMessageInfo

type MsgUpdateKeyResponse struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	NewPubKey string `protobuf:"bytes,2,opt,name=new_public_key,json=newPublicKey,proto3" json:"new_public_key,omitempty"`
}

func (m *MsgUpdateKeyResponse) Reset()         { *m = MsgUpdateKeyResponse{} }
func (m *MsgUpdateKeyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateKeyResponse) ProtoMessage()    {}
func (*MsgUpdateKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92db8b72563f1f, []int{3}
}
func (m *MsgUpdateKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateKeyResponse.Merge(m, src)
}
func (m *MsgUpdateKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateKeyResponse proto.InternalMessageInfo

func (m *MsgUpdateKeyResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgUpdateKeyResponse) GetNewPubKey() string {
	if m != nil {
		return m.NewPubKey
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCreateAccount)(nil), "auranw.aura.smartaccount.MsgCreateAccount")
	proto.RegisterType((*MsgCreateAccountResponse)(nil), "auranw.aura.smartaccount.MsgCreateAccountResponse")
	proto.RegisterType((*MsgUpdateKey)(nil), "auranw.aura.smartaccount.MsgUpdateKey")
	proto.RegisterType((*MsgUpdateKeyResponse)(nil), "auranw.aura.smartaccount.MsgUpdateKeyResponse")
}

func init() { proto.RegisterFile("smartaccount/tx.proto", fileDescriptor_3b92db8b72563f1f) }

var fileDescriptor_3b92db8b72563f1f = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6a, 0xdb, 0x4c,
	0x10, 0xb6, 0xe2, 0xc4, 0x8e, 0x37, 0x4e, 0xf8, 0x11, 0xf9, 0x41, 0xf5, 0x41, 0x32, 0x2e, 0x14,
	0x53, 0x9a, 0xdd, 0xc6, 0xa1, 0x85, 0xf6, 0x56, 0xbb, 0x94, 0x84, 0xe0, 0x50, 0x04, 0xa5, 0xd0,
	0x8b, 0x59, 0xaf, 0xb6, 0xaa, 0x70, 0xb4, 0x2b, 0x34, 0xeb, 0x3a, 0x7a, 0x82, 0xf6, 0xd8, 0x07,
	0xe8, 0x21, 0xe7, 0x3e, 0x49, 0x8e, 0x39, 0x16, 0x0a, 0x6e, 0x51, 0x2e, 0x7d, 0x86, 0x9e, 0xca,
	0x4a, 0x72, 0x2a, 0x07, 0x52, 0x0c, 0x3d, 0xcd, 0xac, 0xf7, 0xf3, 0x37, 0xf3, 0x7d, 0x33, 0x5a,
	0xf4, 0x3f, 0x84, 0x34, 0x56, 0x94, 0x31, 0x39, 0x15, 0x8a, 0xa8, 0x33, 0x1c, 0xc5, 0x52, 0x49,
	0xd3, 0xa2, 0xd3, 0x98, 0x8a, 0x19, 0xd6, 0x01, 0x97, 0x21, 0xad, 0x5d, 0x5f, 0xfa, 0x32, 0x03,
	0x11, 0x9d, 0xe5, 0xf8, 0x96, 0xcd, 0x24, 0x84, 0x12, 0xc8, 0x98, 0x02, 0x27, 0xef, 0xf7, 0xc7,
	0x5c, 0xd1, 0x7d, 0xc2, 0x64, 0x20, 0x8a, 0xfb, 0x4e, 0x71, 0xcf, 0xe2, 0x24, 0x52, 0x92, 0x00,
	0x67, 0x51, 0xef, 0xd1, 0xe3, 0xc9, 0x3e, 0x99, 0xf0, 0x04, 0x72, 0x4c, 0xe7, 0x43, 0x15, 0xfd,
	0x37, 0x04, 0x7f, 0x10, 0x73, 0xaa, 0xf8, 0xb3, 0xbc, 0x9c, 0x69, 0xa1, 0x3a, 0xd3, 0x3f, 0xc8,
	0xd8, 0x32, 0xda, 0x46, 0xb7, 0xe1, 0x2e, 0x8e, 0xe6, 0x5d, 0x54, 0x67, 0xd2, 0xe3, 0xa3, 0xc0,
	0xb3, 0xd6, 0xda, 0x46, 0x77, 0xbd, 0x8f, 0xd2, 0xb9, 0x53, 0x1b, 0x48, 0x8f, 0x1f, 0x3d, 0x77,
	0x6b, 0xfa, 0xea, 0xc8, 0x33, 0x19, 0xda, 0x0c, 0x44, 0xa0, 0x46, 0x21, 0xf8, 0x56, 0xb5, 0x6d,
	0x74, 0x9b, 0xfd, 0xc3, 0x74, 0xee, 0xd4, 0x8f, 0x44, 0xa0, 0x86, 0xe0, 0xff, 0x9a, 0x3b, 0x4f,
	0xfc, 0x40, 0xbd, 0x9b, 0x8e, 0x31, 0x93, 0x21, 0x19, 0x48, 0x08, 0x5f, 0x53, 0x08, 0xc9, 0x8c,
	0x42, 0xe8, 0x91, 0xb3, 0x2c, 0x12, 0x95, 0x44, 0x1c, 0xb0, 0x4b, 0x67, 0x03, 0x29, 0x54, 0x4c,
	0x99, 0x1a, 0x72, 0x00, 0xea, 0x73, 0xb7, 0x1e, 0xe4, 0x2c, 0xe6, 0x09, 0x42, 0xd1, 0x74, 0x7c,
	0x1a, 0xb0, 0xd1, 0x84, 0x27, 0xd6, 0x7a, 0xdb, 0xe8, 0x6e, 0xf5, 0x1c, 0x9c, 0x2b, 0xc6, 0xb9,
	0x62, 0x7c, 0xad, 0x18, 0xbf, 0x9c, 0x8e, 0x8f, 0x79, 0xd2, 0xdf, 0xb9, 0x98, 0x3b, 0x15, 0xdd,
	0x71, 0x7e, 0x76, 0x1b, 0x39, 0xc5, 0x31, 0x4f, 0x4c, 0x8a, 0x36, 0xde, 0x4e, 0x85, 0x07, 0xd6,
	0x46, 0xbb, 0xda, 0xdd, 0xea, 0xdd, 0x59, 0x50, 0x69, 0x73, 0x71, 0x61, 0x2e, 0x1e, 0xc8, 0x40,
	0xf4, 0x1f, 0x6a, 0x92, 0x2f, 0xdf, 0x9d, 0x6e, 0x49, 0xc5, 0xc2, 0xe9, 0x2c, 0xec, 0x81, 0x37,
	0x29, 0x14, 0xe8, 0x3f, 0x80, 0x9b, 0x33, 0x9b, 0x26, 0x5a, 0x07, 0x7a, 0xaa, 0xac, 0x9a, 0xf6,
	0xc4, 0xcd, 0xf2, 0xa7, 0x9b, 0x1f, 0xcf, 0x9d, 0xca, 0xcf, 0x73, 0xa7, 0xd2, 0x39, 0x44, 0xd6,
	0xcd, 0x41, 0xb8, 0x1c, 0x22, 0x29, 0x80, 0xeb, 0x81, 0x50, 0xcf, 0x8b, 0x39, 0xc0, 0x62, 0x20,
	0xc5, 0x51, 0x73, 0x7a, 0x54, 0xd1, 0x6c, 0x1a, 0x4d, 0x37, 0xcb, 0x3b, 0x9f, 0x0d, 0xd4, 0x1c,
	0x82, 0xff, 0x2a, 0xf2, 0xa8, 0xe2, 0x5a, 0xdb, 0xed, 0xf3, 0x2c, 0x11, 0xaf, 0x2d, 0x13, 0x2f,
	0xfb, 0x5b, 0xfd, 0x57, 0x7f, 0x4b, 0x42, 0x39, 0xda, 0x2d, 0x77, 0xb7, 0x82, 0xc8, 0x03, 0xb4,
	0x23, 0xf8, 0x6c, 0x54, 0xea, 0x27, 0x6b, 0xb6, 0xbf, 0x9d, 0xce, 0x9d, 0xc6, 0x09, 0x9f, 0x15,
	0xd5, 0x9a, 0x22, 0x4b, 0xf3, 0x82, 0xbd, 0x6f, 0x06, 0xaa, 0xea, 0x45, 0x91, 0x68, 0x7b, 0x79,
	0xbb, 0xef, 0xe3, 0xdb, 0xbe, 0x33, 0x7c, 0x73, 0x00, 0xad, 0xde, 0xea, 0xd8, 0x6b, 0x1d, 0x0c,
	0x35, 0xfe, 0x58, 0x7f, 0xef, 0xaf, 0x04, 0xd7, 0xb8, 0x16, 0x5e, 0x0d, 0xb7, 0x28, 0xd2, 0x7f,
	0x71, 0x91, 0xda, 0xc6, 0x65, 0x6a, 0x1b, 0x3f, 0x52, 0xdb, 0xf8, 0x74, 0x65, 0x57, 0x2e, 0xaf,
	0xec, 0xca, 0xd7, 0x2b, 0xbb, 0xf2, 0xe6, 0x41, 0x69, 0x2d, 0x35, 0xd9, 0x9e, 0x98, 0x65, 0x91,
	0x9c, 0x91, 0xe5, 0x67, 0x47, 0x2f, 0xe8, 0xb8, 0x96, 0x3d, 0x03, 0x07, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x85, 0x03, 0xe6, 0x93, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateAccount(ctx context.Context, in *MsgCreateAccount, opts ...grpc.CallOption) (*MsgCreateAccountResponse, error)
	UpdateKey(ctx context.Context, in *MsgUpdateKey, opts ...grpc.CallOption) (*MsgUpdateKeyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateAccount(ctx context.Context, in *MsgCreateAccount, opts ...grpc.CallOption) (*MsgCreateAccountResponse, error) {
	out := new(MsgCreateAccountResponse)
	err := c.cc.Invoke(ctx, "/auranw.aura.smartaccount.Msg/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateKey(ctx context.Context, in *MsgUpdateKey, opts ...grpc.CallOption) (*MsgUpdateKeyResponse, error) {
	out := new(MsgUpdateKeyResponse)
	err := c.cc.Invoke(ctx, "/auranw.aura.smartaccount.Msg/UpdateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateAccount(context.Context, *MsgCreateAccount) (*MsgCreateAccountResponse, error)
	UpdateKey(context.Context, *MsgUpdateKey) (*MsgUpdateKeyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateAccount(ctx context.Context, req *MsgCreateAccount) (*MsgCreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (*UnimplementedMsgServer) UpdateKey(ctx context.Context, req *MsgUpdateKey) (*MsgUpdateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKey not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auranw.aura.smartaccount.Msg/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAccount(ctx, req.(*MsgCreateAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auranw.aura.smartaccount.Msg/UpdateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateKey(ctx, req.(*MsgUpdateKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auranw.aura.smartaccount.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Msg_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateKey",
			Handler:    _Msg_UpdateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smartaccount/tx.proto",
}

func (m *MsgCreateAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.InitMsg) > 0 {
		i -= len(m.InitMsg)
		copy(dAtA[i:], m.InitMsg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InitMsg)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CodeID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewPubKey) > 0 {
		i -= len(m.NewPubKey)
		copy(dAtA[i:], m.NewPubKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewPubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CodeID != 0 {
		n += 1 + sovTx(uint64(m.CodeID))
	}
	l = len(m.InitMsg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.PubKey.Size()
	n += 1 + l + sovTx(uint64(l))
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.PubKey.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewPubKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitMsg = append(m.InitMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.InitMsg == nil {
				m.InitMsg = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funds = append(m.Funds, types.Coin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = append(m.Salt[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt == nil {
				m.Salt = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewPubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
