// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/interfaces/account_abstraction/v1/interface.proto

package v1

import (
	fmt "fmt"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgAuthenticate is a message that an x/account account abstraction implementer
// must handle to authenticate a transaction. Always ensure the caller is the Accounts module.
type MsgAuthenticate struct {
	// bundler defines the address of the bundler that sent the operation.
	// NOTE: in case the operation was sent directly by the user, this field will reflect
	// the user address.
	Bundler string `protobuf:"bytes,1,opt,name=bundler,proto3" json:"bundler,omitempty"`
	// raw_tx defines the raw version of the tx, this is useful to compute the signature quickly.
	RawTx *tx.TxRaw `protobuf:"bytes,2,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	// tx defines the decoded version of the tx, coming from raw_tx.
	Tx *tx.Tx `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	// signer_index defines the index of the signer in the tx.
	// Specifically this can be used to extract the signature at the correct
	// index.
	SignerIndex uint32 `protobuf:"varint,4,opt,name=signer_index,json=signerIndex,proto3" json:"signer_index,omitempty"`
}

func (m *MsgAuthenticate) Reset()         { *m = MsgAuthenticate{} }
func (m *MsgAuthenticate) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticate) ProtoMessage()    {}
func (*MsgAuthenticate) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{0}
}
func (m *MsgAuthenticate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticate.Merge(m, src)
}
func (m *MsgAuthenticate) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticate proto.InternalMessageInfo

func (m *MsgAuthenticate) GetBundler() string {
	if m != nil {
		return m.Bundler
	}
	return ""
}

func (m *MsgAuthenticate) GetRawTx() *tx.TxRaw {
	if m != nil {
		return m.RawTx
	}
	return nil
}

func (m *MsgAuthenticate) GetTx() *tx.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *MsgAuthenticate) GetSignerIndex() uint32 {
	if m != nil {
		return m.SignerIndex
	}
	return 0
}

// MsgAuthenticateResponse is the response to MsgAuthenticate.
// The authentication either fails or succeeds, this is why
// there are no auxiliary fields to the response.
type MsgAuthenticateResponse struct {
}

func (m *MsgAuthenticateResponse) Reset()         { *m = MsgAuthenticateResponse{} }
func (m *MsgAuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticateResponse) ProtoMessage()    {}
func (*MsgAuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{1}
}
func (m *MsgAuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticateResponse.Merge(m, src)
}
func (m *MsgAuthenticateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticateResponse proto.InternalMessageInfo

// QueryAuthenticationMethods is a query that an x/account account abstraction implementer
// must handle to return the authentication methods that the account supports.
type QueryAuthenticationMethods struct {
}

func (m *QueryAuthenticationMethods) Reset()         { *m = QueryAuthenticationMethods{} }
func (m *QueryAuthenticationMethods) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethods) ProtoMessage()    {}
func (*QueryAuthenticationMethods) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{2}
}
func (m *QueryAuthenticationMethods) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethods.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethods.Merge(m, src)
}
func (m *QueryAuthenticationMethods) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethods) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethods.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethods proto.InternalMessageInfo

// QueryAuthenticationMethodsResponse is the response to QueryAuthenticationMethods.
type QueryAuthenticationMethodsResponse struct {
	// authentication_methods are the authentication methods that the account supports.
	AuthenticationMethods []string `protobuf:"bytes,1,rep,name=authentication_methods,json=authenticationMethods,proto3" json:"authentication_methods,omitempty"`
}

func (m *QueryAuthenticationMethodsResponse) Reset()         { *m = QueryAuthenticationMethodsResponse{} }
func (m *QueryAuthenticationMethodsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethodsResponse) ProtoMessage()    {}
func (*QueryAuthenticationMethodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{3}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethodsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.Merge(m, src)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethodsResponse proto.InternalMessageInfo

func (m *QueryAuthenticationMethodsResponse) GetAuthenticationMethods() []string {
	if m != nil {
		return m.AuthenticationMethods
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgAuthenticate)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticate")
	proto.RegisterType((*MsgAuthenticateResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticateResponse")
	proto.RegisterType((*QueryAuthenticationMethods)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethods")
	proto.RegisterType((*QueryAuthenticationMethodsResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethodsResponse")
}

func init() {
	proto.RegisterFile("cosmos/accounts/interfaces/account_abstraction/v1/interface.proto", fileDescriptor_56b360422260e9d1)
}

var fileDescriptor_56b360422260e9d1 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xeb, 0xf6, 0xde, 0xa2, 0xba, 0x20, 0xa4, 0x48, 0x85, 0x50, 0xa1, 0x28, 0x44, 0x42,
	0xca, 0x64, 0x2b, 0x20, 0x06, 0xc6, 0xb2, 0x31, 0x74, 0x20, 0x74, 0x82, 0x21, 0x72, 0x12, 0xd3,
	0x5a, 0x50, 0xbb, 0xb2, 0x4f, 0x5a, 0xf3, 0x16, 0x3c, 0x05, 0xcf, 0xc2, 0xd8, 0x91, 0x11, 0xb5,
	0x2f, 0x82, 0x68, 0x5a, 0x0a, 0xa8, 0x0c, 0x8c, 0xfe, 0xcf, 0xe7, 0xcf, 0xc7, 0xfa, 0x71, 0x27,
	0x53, 0x66, 0xa8, 0x0c, 0x65, 0x59, 0xa6, 0x0a, 0x09, 0x86, 0x0a, 0x09, 0x5c, 0xdf, 0xb1, 0x8c,
	0x7f, 0x66, 0x09, 0x4b, 0x0d, 0x68, 0x96, 0x81, 0x50, 0x92, 0x8e, 0xa3, 0x35, 0x41, 0x46, 0x5a,
	0x81, 0x72, 0xa2, 0x52, 0x41, 0x56, 0x0a, 0xb2, 0x56, 0x90, 0x0d, 0x0a, 0x32, 0x8e, 0xda, 0xed,
	0xe5, 0xab, 0x60, 0xe9, 0x38, 0x4a, 0x39, 0xb0, 0x88, 0x82, 0x2d, 0x75, 0xc1, 0x33, 0xc2, 0xbb,
	0x5d, 0xd3, 0xef, 0x14, 0x30, 0xe0, 0x12, 0x44, 0xc6, 0x80, 0x3b, 0x2e, 0xde, 0x4a, 0x0b, 0x99,
	0x3f, 0x70, 0xed, 0x22, 0x1f, 0x85, 0x8d, 0x78, 0x75, 0x74, 0x28, 0xae, 0x6b, 0x36, 0x49, 0xc0,
	0xba, 0x55, 0x1f, 0x85, 0xcd, 0x13, 0x97, 0x2c, 0xb7, 0x01, 0x4b, 0x96, 0x6a, 0xd2, 0xb3, 0x31,
	0x9b, 0xc4, 0xff, 0x35, 0x9b, 0xf4, 0xac, 0x73, 0x8c, 0xab, 0x60, 0xdd, 0xda, 0x02, 0x6e, 0x6d,
	0x86, 0xab, 0x60, 0x9d, 0x23, 0xbc, 0x6d, 0x44, 0x5f, 0x72, 0x9d, 0x08, 0x99, 0x73, 0xeb, 0xfe,
	0xf3, 0x51, 0xb8, 0x13, 0x37, 0xcb, 0xec, 0xf2, 0x23, 0x0a, 0x0e, 0xf0, 0xfe, 0x8f, 0x3d, 0x63,
	0x6e, 0x46, 0x4a, 0x1a, 0x1e, 0x1c, 0xe2, 0xf6, 0x55, 0xc1, 0xf5, 0xe3, 0x97, 0xa1, 0x50, 0xb2,
	0xcb, 0x61, 0xa0, 0x72, 0x13, 0xdc, 0xe2, 0xe0, 0xf7, 0xe9, 0xca, 0xe1, 0x9c, 0xe1, 0x3d, 0xf6,
	0x0d, 0x48, 0x86, 0x25, 0xe1, 0x22, 0xbf, 0x16, 0x36, 0xe2, 0x16, 0xdb, 0x74, 0xfd, 0xe2, 0xfa,
	0x65, 0xe6, 0xa1, 0xe9, 0xcc, 0x43, 0x6f, 0x33, 0x0f, 0x3d, 0xcd, 0xbd, 0xca, 0x74, 0xee, 0x55,
	0x5e, 0xe7, 0x5e, 0xe5, 0xe6, 0xbc, 0xfc, 0xac, 0xc9, 0xef, 0x89, 0x50, 0xd4, 0xfe, 0xa1, 0xf2,
	0xb4, 0xbe, 0xa8, 0xe6, 0xf4, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x1c, 0x3d, 0x0e, 0x2e, 0x02,
	0x00, 0x00,
}

func (m *MsgAuthenticate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignerIndex != 0 {
		i = encodeVarintInterface(dAtA, i, uint64(m.SignerIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.RawTx != nil {
		{
			size, err := m.RawTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Bundler) > 0 {
		i -= len(m.Bundler)
		copy(dAtA[i:], m.Bundler)
		i = encodeVarintInterface(dAtA, i, uint64(len(m.Bundler)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAuthenticateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethods) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethods) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethods) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethodsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethodsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethodsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for iNdEx := len(m.AuthenticationMethods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthenticationMethods[iNdEx])
			copy(dAtA[i:], m.AuthenticationMethods[iNdEx])
			i = encodeVarintInterface(dAtA, i, uint64(len(m.AuthenticationMethods[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterface(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterface(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAuthenticate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bundler)
	if l > 0 {
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.RawTx != nil {
		l = m.RawTx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.SignerIndex != 0 {
		n += 1 + sovInterface(uint64(m.SignerIndex))
	}
	return n
}

func (m *MsgAuthenticateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethods) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethodsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for _, s := range m.AuthenticationMethods {
			l = len(s)
			n += 1 + l + sovInterface(uint64(l))
		}
	}
	return n
}

func sovInterface(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterface(x uint64) (n int) {
	return sovInterface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAuthenticate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bundler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bundler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RawTx == nil {
				m.RawTx = &tx.TxRaw{}
			}
			if err := m.RawTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &tx.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerIndex", wireType)
			}
			m.SignerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignerIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *MsgAuthenticateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethods) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethods: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethods: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethodsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticationMethods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticationMethods = append(m.AuthenticationMethods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func skipInterface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
				return 0, ErrInvalidLengthInterface
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterface
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterface
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterface        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterface          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterface = fmt.Errorf("proto: unexpected end of group")
)
