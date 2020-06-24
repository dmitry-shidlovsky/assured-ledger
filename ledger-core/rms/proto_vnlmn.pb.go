// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto_vnlmn.proto

package rms

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/insolar/assured-ledger/ledger-core/insproto"
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

type RegistrationFlags int32

const (
	RegistrationFlags_RFUndefined RegistrationFlags = 0
	RegistrationFlags_RFFast      RegistrationFlags = 1
	RegistrationFlags_RFSafe      RegistrationFlags = 2
	RegistrationFlags_RFFastSafe  RegistrationFlags = 3
)

var RegistrationFlags_name = map[int32]string{
	0: "RFUndefined",
	1: "RFFast",
	2: "RFSafe",
	3: "RFFastSafe",
}

var RegistrationFlags_value = map[string]int32{
	"RFUndefined": 0,
	"RFFast":      1,
	"RFSafe":      2,
	"RFFastSafe":  3,
}

func (x RegistrationFlags) String() string {
	return proto.EnumName(RegistrationFlags_name, int32(x))
}

func (RegistrationFlags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b0afd09ea099d9b9, []int{0}
}

type LRegisterRequest struct {
	AnyRecordLazy     `protobuf:"bytes,19,opt,name=Record,proto3,embedded=Record" json:"Record"`
	Flags             RegistrationFlags `protobuf:"varint,1800,opt,name=Flags,proto3,enum=rms.RegistrationFlags" json:"Flags"`
	AnticipatedRef    Reference         `protobuf:"bytes,1801,opt,name=AnticipatedRef,proto3" json:"AnticipatedRef"`
	ProducerSignature Binary            `protobuf:"bytes,1802,opt,name=ProducerSignature,proto3" json:"ProducerSignature"`
	ProducedBy        Reference         `protobuf:"bytes,1803,opt,name=ProducedBy,proto3" json:"ProducedBy"`
}

func (m *LRegisterRequest) Reset()         { *m = LRegisterRequest{} }
func (m *LRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*LRegisterRequest) ProtoMessage()    {}
func (*LRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0afd09ea099d9b9, []int{0}
}
func (m *LRegisterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LRegisterRequest.Merge(m, src)
}
func (m *LRegisterRequest) XXX_Size() int {
	return m.ProtoSize()
}
func (m *LRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LRegisterRequest proto.InternalMessageInfo

func (m *LRegisterRequest) GetFlags() RegistrationFlags {
	if m != nil {
		return m.Flags
	}
	return RegistrationFlags_RFUndefined
}

func (m *LRegisterRequest) GetAnticipatedRef() Reference {
	if m != nil {
		return m.AnticipatedRef
	}
	return Reference{}
}

func (m *LRegisterRequest) GetProducerSignature() Binary {
	if m != nil {
		return m.ProducerSignature
	}
	return Binary{}
}

func (m *LRegisterRequest) GetProducedBy() Reference {
	if m != nil {
		return m.ProducedBy
	}
	return Reference{}
}

type LRegisterResponse struct {
	Flags              RegistrationFlags `protobuf:"varint,1800,opt,name=Flags,proto3,enum=rms.RegistrationFlags" json:"Flags"`
	AnticipatedRef     Reference         `protobuf:"bytes,1801,opt,name=AnticipatedRef,proto3" json:"AnticipatedRef"`
	RegistrarSignature Binary            `protobuf:"bytes,1805,opt,name=RegistrarSignature,proto3" json:"RegistrarSignature"`
}

func (m *LRegisterResponse) Reset()         { *m = LRegisterResponse{} }
func (m *LRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*LRegisterResponse) ProtoMessage()    {}
func (*LRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0afd09ea099d9b9, []int{1}
}
func (m *LRegisterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LRegisterResponse.Merge(m, src)
}
func (m *LRegisterResponse) XXX_Size() int {
	return m.ProtoSize()
}
func (m *LRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LRegisterResponse proto.InternalMessageInfo

func (m *LRegisterResponse) GetFlags() RegistrationFlags {
	if m != nil {
		return m.Flags
	}
	return RegistrationFlags_RFUndefined
}

func (m *LRegisterResponse) GetAnticipatedRef() Reference {
	if m != nil {
		return m.AnticipatedRef
	}
	return Reference{}
}

func (m *LRegisterResponse) GetRegistrarSignature() Binary {
	if m != nil {
		return m.RegistrarSignature
	}
	return Binary{}
}

func init() {
	proto.RegisterEnum("rms.RegistrationFlags", RegistrationFlags_name, RegistrationFlags_value)
	proto.RegisterType((*LRegisterRequest)(nil), "rms.LRegisterRequest")
	proto.RegisterType((*LRegisterResponse)(nil), "rms.LRegisterResponse")
}

func init() { proto.RegisterFile("proto_vnlmn.proto", fileDescriptor_b0afd09ea099d9b9) }

var fileDescriptor_b0afd09ea099d9b9 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x3f, 0x8b, 0x13, 0x41,
	0x18, 0xc6, 0x77, 0xee, 0x92, 0x98, 0x4c, 0x60, 0xdd, 0x4c, 0x54, 0xc2, 0x15, 0xbb, 0xe1, 0xaa,
	0x43, 0x48, 0x02, 0x77, 0xd8, 0x08, 0xc2, 0x65, 0x8b, 0x20, 0x21, 0x07, 0xb2, 0xa7, 0xb6, 0x32,
	0xd9, 0x7d, 0xb3, 0x0e, 0x24, 0x33, 0xb9, 0x99, 0x5d, 0x21, 0x7e, 0x82, 0x78, 0x22, 0x88, 0x85,
	0x70, 0xa9, 0x2c, 0xc5, 0x4a, 0xf4, 0x0b, 0x58, 0xa6, 0x4c, 0x25, 0xb1, 0x39, 0x74, 0xd3, 0xf8,
	0x19, 0xae, 0x92, 0xcc, 0xee, 0x41, 0xf0, 0xd4, 0xd6, 0x2a, 0xef, 0x9f, 0xdf, 0xf3, 0xbc, 0x93,
	0x87, 0xc5, 0x95, 0xb1, 0x14, 0x91, 0x78, 0xf2, 0x8c, 0x0f, 0x47, 0xbc, 0xa9, 0x6b, 0xb2, 0x2d,
	0x47, 0x6a, 0xa7, 0x11, 0xb2, 0xe8, 0x69, 0xdc, 0x6f, 0xfa, 0x62, 0xd4, 0x0a, 0x45, 0x28, 0x5a,
	0x7a, 0xd7, 0x8f, 0x07, 0xba, 0xd3, 0x8d, 0xae, 0x52, 0xcd, 0xce, 0xe1, 0x06, 0xce, 0xb8, 0x12,
	0x43, 0x2a, 0x5b, 0x54, 0xa9, 0x58, 0x42, 0xd0, 0x18, 0x42, 0x10, 0x82, 0x6c, 0xa5, 0x3f, 0x0d,
	0x5f, 0x48, 0x58, 0x23, 0xa9, 0x05, 0xe3, 0x2a, 0x73, 0x28, 0xc9, 0x51, 0x56, 0xee, 0x7e, 0xdd,
	0xc2, 0x56, 0xcf, 0x83, 0x90, 0xa9, 0x08, 0xa4, 0x07, 0x27, 0x31, 0xa8, 0x88, 0xb4, 0x71, 0xc1,
	0x03, 0x5f, 0xc8, 0xa0, 0x56, 0xad, 0xa3, 0xbd, 0xf2, 0x3e, 0x69, 0xae, 0x05, 0x6d, 0x3e, 0x49,
	0xa7, 0x3d, 0xfa, 0x7c, 0xe2, 0xde, 0x9c, 0x9f, 0x3b, 0xc6, 0xe2, 0xdc, 0x41, 0x17, 0x6f, 0xeb,
	0xa5, 0x23, 0x15, 0xa6, 0x2b, 0x2f, 0x13, 0x92, 0x03, 0x9c, 0xef, 0x0c, 0x69, 0xa8, 0x6a, 0x53,
	0xb3, 0x8e, 0xf6, 0xcc, 0xfd, 0x5b, 0xda, 0x22, 0x3d, 0x24, 0x69, 0xc4, 0x04, 0xd7, 0x6b, 0x37,
	0xb7, 0xb6, 0xf1, 0x52, 0x96, 0xdc, 0xc3, 0x66, 0x9b, 0x47, 0xcc, 0x67, 0x63, 0x1a, 0x41, 0xe0,
	0xc1, 0xa0, 0xf6, 0xc2, 0xd4, 0x0f, 0x30, 0x33, 0xf5, 0x00, 0x24, 0x70, 0x1f, 0x32, 0xd5, 0x6f,
	0x30, 0x39, 0xc4, 0x95, 0x07, 0x52, 0x04, 0xb1, 0x0f, 0xf2, 0x98, 0x85, 0x9c, 0x46, 0xb1, 0x84,
	0xda, 0x69, 0xea, 0x50, 0xd6, 0x0e, 0x2e, 0xe3, 0x54, 0x4e, 0x32, 0xf9, 0x55, 0x98, 0xdc, 0xc1,
	0x38, 0x1b, 0x06, 0xee, 0xa4, 0xf6, 0xf2, 0x5f, 0xc7, 0x37, 0xc0, 0xbb, 0xf9, 0x2f, 0x4b, 0xe7,
	0x63, 0xa9, 0x9b, 0x2b, 0x22, 0xab, 0xda, 0xcd, 0x17, 0x6f, 0x58, 0x53, 0x73, 0xf7, 0x1b, 0xc2,
	0x95, 0x8d, 0x60, 0xd5, 0x58, 0x70, 0x05, 0xff, 0x25, 0x16, 0x17, 0x93, 0xcb, 0x03, 0x1b, 0xb9,
	0xbc, 0xfa, 0x6b, 0x2e, 0x7f, 0xa0, 0xf5, 0x3f, 0xfc, 0x5c, 0xea, 0xe6, 0x8b, 0xc8, 0x9a, 0x9a,
	0xb7, 0x7b, 0xb8, 0x72, 0xe5, 0xc9, 0xe4, 0x3a, 0x2e, 0x7b, 0x9d, 0x47, 0x3c, 0x80, 0x01, 0xe3,
	0x10, 0x58, 0x06, 0xc1, 0xb8, 0xe0, 0x75, 0x3a, 0x54, 0x45, 0x16, 0x4a, 0xeb, 0x63, 0x3a, 0x00,
	0x6b, 0x8b, 0x98, 0x18, 0xa7, 0x73, 0xdd, 0x6f, 0xbb, 0x27, 0xf3, 0x1f, 0x36, 0x7a, 0x9f, 0xd8,
	0x68, 0x9e, 0xd8, 0x68, 0x91, 0xd8, 0x68, 0x99, 0xd8, 0xe8, 0x7b, 0x62, 0x1b, 0xaf, 0x57, 0xb6,
	0xf1, 0x6e, 0x65, 0xa3, 0xc5, 0xca, 0x36, 0x96, 0x2b, 0xdb, 0xf8, 0x79, 0xe6, 0xa0, 0x8b, 0x33,
	0xe7, 0xda, 0x11, 0x28, 0x45, 0x43, 0x38, 0x9d, 0x39, 0xb9, 0xfb, 0x40, 0x83, 0x37, 0x33, 0x27,
	0xff, 0x98, 0x29, 0x16, 0x7d, 0x98, 0x39, 0xd5, 0xcb, 0xc4, 0x33, 0xe6, 0xe1, 0x64, 0x0c, 0x9f,
	0x66, 0x8e, 0x99, 0xb5, 0x1a, 0x13, 0xb2, 0x5f, 0xd0, 0x1f, 0xff, 0xc1, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7c, 0x3d, 0x72, 0x8b, 0x92, 0x03, 0x00, 0x00,
}

func (this *LRegisterRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LRegisterRequest)
	if !ok {
		that2, ok := that.(LRegisterRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AnyRecordLazy.Equal(&that1.AnyRecordLazy) {
		return false
	}
	if this.Flags != that1.Flags {
		return false
	}
	if !this.AnticipatedRef.Equal(&that1.AnticipatedRef) {
		return false
	}
	if !this.ProducerSignature.Equal(&that1.ProducerSignature) {
		return false
	}
	if !this.ProducedBy.Equal(&that1.ProducedBy) {
		return false
	}
	return true
}
func (this *LRegisterResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LRegisterResponse)
	if !ok {
		that2, ok := that.(LRegisterResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Flags != that1.Flags {
		return false
	}
	if !this.AnticipatedRef.Equal(&that1.AnticipatedRef) {
		return false
	}
	if !this.RegistrarSignature.Equal(&that1.RegistrarSignature) {
		return false
	}
	return true
}
func (m *LRegisterRequest) Visit(ctx MessageVisitor) error {
	if err := ctx.MsgRecord(m, 19, &m.AnyRecordLazy); err != nil {
		return err
	}
	return ctx.Message(m, 1200)
}

const TypeLRegisterRequestPolymorthID uint64 = 1200

func (*LRegisterRequest) GetDefaultPolymorphID() uint64 {
	return 1200
}

func (m *LRegisterRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LRegisterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LRegisterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l, fieldEnd int
	_, _ = l, fieldEnd
	{
		size, err := m.ProducedBy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		if size > 0 {
			i -= size
			i = encodeVarintProtoVnlmn(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x70
			i--
			dAtA[i] = 0xda
		}
	}
	{
		size, err := m.ProducerSignature.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		if size > 0 {
			i -= size
			i = encodeVarintProtoVnlmn(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x70
			i--
			dAtA[i] = 0xd2
		}
	}
	{
		size, err := m.AnticipatedRef.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		if size > 0 {
			i -= size
			i = encodeVarintProtoVnlmn(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x70
			i--
			dAtA[i] = 0xca
		}
	}
	if m.Flags != 0 {
		i = encodeVarintProtoVnlmn(dAtA, i, uint64(m.Flags))
		i--
		dAtA[i] = 0x70
		i--
		dAtA[i] = 0xc0
	}
	{
		size, err := m.AnyRecordLazy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		if size > 0 {
			i -= size
			i = encodeVarintProtoVnlmn(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x9a
		}
	}
	i = encodeVarintProtoVnlmn(dAtA, i, uint64(1200))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x80
	return len(dAtA) - i, nil
}

func (m *LRegisterResponse) Visit(ctx MessageVisitor) error {
	return ctx.Message(m, 1203)
}

const TypeLRegisterResponsePolymorthID uint64 = 1203

func (*LRegisterResponse) GetDefaultPolymorphID() uint64 {
	return 1203
}

func (m *LRegisterResponse) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LRegisterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LRegisterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l, fieldEnd int
	_, _ = l, fieldEnd
	{
		size, err := m.RegistrarSignature.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		if size > 0 {
			i -= size
			i = encodeVarintProtoVnlmn(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x70
			i--
			dAtA[i] = 0xea
		}
	}
	{
		size, err := m.AnticipatedRef.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		if size > 0 {
			i -= size
			i = encodeVarintProtoVnlmn(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x70
			i--
			dAtA[i] = 0xca
		}
	}
	if m.Flags != 0 {
		i = encodeVarintProtoVnlmn(dAtA, i, uint64(m.Flags))
		i--
		dAtA[i] = 0x70
		i--
		dAtA[i] = 0xc0
	}
	i = encodeVarintProtoVnlmn(dAtA, i, uint64(1203))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x80
	return len(dAtA) - i, nil
}

func encodeVarintProtoVnlmn(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtoVnlmn(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func init() {
	RegisterMessageType(1200, "", (*LRegisterRequest)(nil))
	RegisterMessageType(1203, "", (*LRegisterResponse)(nil))
}

func (m *LRegisterRequest) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if l = m.AnyRecordLazy.ProtoSize(); l > 0 {
		n += 2 + l + sovProtoVnlmn(uint64(l))
	}
	if m.Flags != 0 {
		n += 2 + sovProtoVnlmn(uint64(m.Flags))
	}
	if l = m.AnticipatedRef.ProtoSize(); l > 0 {
		n += 2 + l + sovProtoVnlmn(uint64(l))
	}
	if l = m.ProducerSignature.ProtoSize(); l > 0 {
		n += 2 + l + sovProtoVnlmn(uint64(l))
	}
	if l = m.ProducedBy.ProtoSize(); l > 0 {
		n += 2 + l + sovProtoVnlmn(uint64(l))
	}
	n += 2 + sovProtoVnlmn(1200)
	return n
}

func (m *LRegisterResponse) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Flags != 0 {
		n += 2 + sovProtoVnlmn(uint64(m.Flags))
	}
	if l = m.AnticipatedRef.ProtoSize(); l > 0 {
		n += 2 + l + sovProtoVnlmn(uint64(l))
	}
	if l = m.RegistrarSignature.ProtoSize(); l > 0 {
		n += 2 + l + sovProtoVnlmn(uint64(l))
	}
	n += 2 + sovProtoVnlmn(1203)
	return n
}

func sovProtoVnlmn(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtoVnlmn(x uint64) (n int) {
	return sovProtoVnlmn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LRegisterRequest) Unmarshal(dAtA []byte) error {
	return m.UnmarshalWithUnknownCallback(dAtA, skipProtoVnlmn)
}
func (m *LRegisterRequest) UnmarshalWithUnknownCallback(dAtA []byte, skipFn func([]byte) (int, error)) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtoVnlmn
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
			return fmt.Errorf("proto: LRegisterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LRegisterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnyRecordLazy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
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
				return ErrInvalidLengthProtoVnlmn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtoVnlmn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnyRecordLazy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1800:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			m.Flags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flags |= RegistrationFlags(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 1801:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnticipatedRef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
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
				return ErrInvalidLengthProtoVnlmn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtoVnlmn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnticipatedRef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1802:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerSignature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
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
				return ErrInvalidLengthProtoVnlmn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtoVnlmn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProducerSignature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1803:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducedBy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
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
				return ErrInvalidLengthProtoVnlmn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtoVnlmn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProducedBy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				l = iNdEx
				break
			}
			if skippy == 0 {
				if skippy, err = skipProtoVnlmn(dAtA[iNdEx:]); err != nil {
					return err
				}
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtoVnlmn
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
func (m *LRegisterResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalWithUnknownCallback(dAtA, skipProtoVnlmn)
}
func (m *LRegisterResponse) UnmarshalWithUnknownCallback(dAtA []byte, skipFn func([]byte) (int, error)) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtoVnlmn
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
			return fmt.Errorf("proto: LRegisterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LRegisterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1800:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			m.Flags = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flags |= RegistrationFlags(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 1801:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnticipatedRef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
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
				return ErrInvalidLengthProtoVnlmn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtoVnlmn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnticipatedRef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1805:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistrarSignature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoVnlmn
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
				return ErrInvalidLengthProtoVnlmn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtoVnlmn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RegistrarSignature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				l = iNdEx
				break
			}
			if skippy == 0 {
				if skippy, err = skipProtoVnlmn(dAtA[iNdEx:]); err != nil {
					return err
				}
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtoVnlmn
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
func skipProtoVnlmn(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtoVnlmn
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
					return 0, ErrIntOverflowProtoVnlmn
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
					return 0, ErrIntOverflowProtoVnlmn
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
				return 0, ErrInvalidLengthProtoVnlmn
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtoVnlmn
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtoVnlmn
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtoVnlmn        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtoVnlmn          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtoVnlmn = fmt.Errorf("proto: unexpected end of group")
	ErrExpectedBinaryMarkerProtoVnlmn = fmt.Errorf("proto: binary marker was expected")
)