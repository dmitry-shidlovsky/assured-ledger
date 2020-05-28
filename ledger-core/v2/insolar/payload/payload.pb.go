// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: insolar/payload/payload.proto

package payload

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/insolar/assured-ledger/ledger-core/v2/insproto"
	github_com_insolar_assured_ledger_ledger_core_v2_pulse "github.com/insolar/assured-ledger/ledger-core/v2/pulse"
	github_com_insolar_assured_ledger_ledger_core_v2_reference "github.com/insolar/assured-ledger/ledger-core/v2/reference"
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

type Meta struct {
	Payload    []byte                                                            `protobuf:"bytes,20,opt,name=Payload,proto3" json:"Payload"`
	Sender     github_com_insolar_assured_ledger_ledger_core_v2_reference.Global `protobuf:"bytes,21,opt,name=Sender,proto3,customtype=github.com/insolar/assured-ledger/ledger-core/v2/reference.Global" json:"Sender"`
	Receiver   github_com_insolar_assured_ledger_ledger_core_v2_reference.Global `protobuf:"bytes,22,opt,name=Receiver,proto3,customtype=github.com/insolar/assured-ledger/ledger-core/v2/reference.Global" json:"Receiver"`
	Pulse      github_com_insolar_assured_ledger_ledger_core_v2_pulse.Number     `protobuf:"varint,23,opt,name=Pulse,proto3,casttype=github.com/insolar/assured-ledger/ledger-core/v2/pulse.Number" json:"Pulse"`
	ID         []byte                                                            `protobuf:"bytes,24,opt,name=ID,proto3" json:"ID"`
	OriginHash MessageHash                                                       `protobuf:"bytes,25,opt,name=OriginHash,proto3,customtype=MessageHash" json:"OriginHash"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_33334fec96407f54, []int{0}
}
func (m *Meta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Meta) GetPulse() github_com_insolar_assured_ledger_ledger_core_v2_pulse.Number {
	if m != nil {
		return m.Pulse
	}
	return 0
}

func (m *Meta) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func init() {
	proto.RegisterType((*Meta)(nil), "payload.Meta")
}

func init() { proto.RegisterFile("insolar/payload/payload.proto", fileDescriptor_33334fec96407f54) }

var fileDescriptor_33334fec96407f54 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0xa3, 0x77, 0x41, 0xe6, 0xcf, 0x10, 0x0e, 0x2e, 0x9c, 0x84, 0x5d, 0x31, 0xdd,
	0xd2, 0x44, 0x82, 0x01, 0x09, 0x89, 0x81, 0xde, 0x49, 0x5c, 0x87, 0x42, 0x15, 0xd8, 0x91, 0x93,
	0xbc, 0x4d, 0x2d, 0xa5, 0x71, 0x64, 0x27, 0x15, 0xfd, 0x0a, 0x4c, 0x4c, 0x48, 0xed, 0xc4, 0xcc,
	0x82, 0xc4, 0x27, 0x60, 0xec, 0xd8, 0xb1, 0x62, 0xa8, 0x50, 0xb2, 0x30, 0x76, 0xee, 0x84, 0xe2,
	0xa4, 0xa8, 0x6b, 0x87, 0x9b, 0x5e, 0x3f, 0x7e, 0x5e, 0xff, 0x9e, 0xd7, 0x96, 0xf1, 0x13, 0x9e,
	0x28, 0x11, 0x33, 0xe9, 0xa6, 0x6c, 0x1a, 0x0b, 0x16, 0xee, 0xaa, 0x93, 0x4a, 0x91, 0x09, 0xcb,
	0x6c, 0xe4, 0x79, 0x27, 0xe2, 0xd9, 0x28, 0xf7, 0x9d, 0x40, 0x8c, 0xdd, 0x48, 0x44, 0xc2, 0xd5,
	0xbe, 0x9f, 0x0f, 0xb5, 0xd2, 0x42, 0xaf, 0xea, 0x73, 0xe7, 0x97, 0x7b, 0xed, 0xbb, 0x04, 0xa6,
	0x54, 0x2e, 0x21, 0xec, 0xc4, 0x10, 0x46, 0x20, 0xdd, 0xba, 0x74, 0x02, 0x21, 0xc1, 0x9d, 0x3c,
	0xab, 0xba, 0x6a, 0x0a, 0x4f, 0x54, 0x0d, 0x79, 0xfa, 0xe3, 0x16, 0x6e, 0xf5, 0x21, 0x63, 0x16,
	0xc1, 0xe6, 0xa0, 0x9e, 0xc3, 0x3e, 0x6d, 0xa3, 0x8b, 0xbb, 0xdd, 0xd6, 0x62, 0x4d, 0x0d, 0x6f,
	0xb7, 0x69, 0x01, 0x3e, 0x79, 0x0f, 0x49, 0x08, 0xd2, 0x7e, 0xa8, 0xed, 0x7e, 0x65, 0xff, 0x5e,
	0xd3, 0xd7, 0x07, 0x4f, 0x21, 0x61, 0x08, 0x12, 0x92, 0x00, 0x9c, 0x37, 0xb1, 0xf0, 0x59, 0xbc,
	0xf9, 0xda, 0x36, 0xbc, 0x06, 0x6e, 0x71, 0x7c, 0xdb, 0x83, 0x00, 0xf8, 0x04, 0xa4, 0xfd, 0xe8,
	0x26, 0x82, 0xfe, 0xe3, 0xad, 0x8f, 0xf8, 0x78, 0x90, 0xc7, 0x0a, 0xec, 0xb3, 0x36, 0xba, 0xb8,
	0xd7, 0xed, 0x55, 0x39, 0xdb, 0x35, 0x7d, 0x75, 0x70, 0x4e, 0x5a, 0x51, 0x9c, 0xb7, 0xf9, 0xd8,
	0x07, 0xa9, 0x33, 0x6a, 0xae, 0x75, 0x8a, 0x8f, 0x7a, 0x57, 0xb6, 0xbd, 0xf7, 0x9a, 0x47, 0xbd,
	0x2b, 0xeb, 0x05, 0xc6, 0xef, 0x24, 0x8f, 0x78, 0x72, 0xcd, 0xd4, 0xc8, 0x7e, 0xac, 0xdd, 0xb3,
	0xe6, 0x8e, 0x77, 0xfa, 0xa0, 0x14, 0x8b, 0xa0, 0xb2, 0x34, 0x69, 0xaf, 0xf5, 0xe5, 0xf1, 0xaf,
	0x15, 0xdd, 0x98, 0xdd, 0x70, 0x51, 0x10, 0xb4, 0x2c, 0x08, 0x5a, 0x15, 0x04, 0xfd, 0x29, 0x88,
	0xf1, 0xa5, 0x24, 0xc6, 0xb7, 0x92, 0xa0, 0x65, 0x49, 0x8c, 0x55, 0x49, 0x8c, 0xbf, 0x33, 0x8a,
	0x36, 0x33, 0x8a, 0xb6, 0x33, 0x6a, 0x36, 0xbc, 0xcf, 0x73, 0xda, 0xba, 0x06, 0x16, 0x7e, 0x9f,
	0xd3, 0x07, 0x1e, 0x44, 0x5c, 0x65, 0x20, 0x1b, 0xeb, 0xc3, 0x34, 0x85, 0x9f, 0x73, 0x7a, 0xbf,
	0x91, 0x97, 0x22, 0xc9, 0xe0, 0x53, 0xe6, 0x9f, 0xe8, 0xef, 0xf1, 0xfc, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0xbb, 0x07, 0x35, 0xbc, 0x02, 0x00, 0x00,
}

func (m *Meta) SetupContext(ctx MessageContext) error {
	return ctx.Message(m, 1008)
}

const TypeMetaPolymorthID = 1008

func (*Meta) GetDefaultPolymorphID() uint64 {
	return 1008
}

func (m *Meta) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Meta) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Meta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l, fieldEnd int
	_, _ = l, fieldEnd
	{
		size := m.OriginHash.ProtoSize()
		i -= size
		if _, err := m.OriginHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPayload(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xca
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i--
		dAtA[i] = 132
		i = encodeVarintPayload(dAtA, i, uint64(len(m.ID)+1))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc2
	}
	if m.Pulse != 0 {
		i = encodeVarintPayload(dAtA, i, uint64(m.Pulse))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb8
	}
	{
		size := m.Receiver.ProtoSize()
		i -= size
		if _, err := m.Receiver.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPayload(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xb2
	{
		size := m.Sender.ProtoSize()
		i -= size
		if _, err := m.Sender.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPayload(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xaa
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i--
		dAtA[i] = 132
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Payload)+1))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	i = encodeVarintPayload(dAtA, i, uint64(1008))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x80
	return len(dAtA) - i, nil
}

func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	offset -= sovPayload(v)
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
	RegisterMessageType(1008, "", (*Meta)(nil))
}

func (m *Meta) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payload)
	if l > 0 {
		l++
		n += 2 + l + sovPayload(uint64(l))
	}
	l = m.Sender.ProtoSize()
	n += 2 + l + sovPayload(uint64(l))
	l = m.Receiver.ProtoSize()
	n += 2 + l + sovPayload(uint64(l))
	if m.Pulse != 0 {
		n += 2 + sovPayload(uint64(m.Pulse))
	}
	l = len(m.ID)
	if l > 0 {
		l++
		n += 2 + l + sovPayload(uint64(l))
	}
	l = m.OriginHash.ProtoSize()
	n += 2 + l + sovPayload(uint64(l))
	n += 2 + sovPayload(1008)
	return n
}

func sovPayload(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Meta) Unmarshal(dAtA []byte) error {
	return m.UnmarshalWithUnknownCallback(dAtA, skipPayload)
}
func (m *Meta) UnmarshalWithUnknownCallback(dAtA []byte, skipFn func([]byte) (int, error)) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
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
			return fmt.Errorf("proto: Meta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Meta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if byteLen > 0 {
				if dAtA[iNdEx] != 132 {
					return ErrExpectedBinaryMarkerPayload
				}
				iNdEx++
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Receiver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pulse", wireType)
			}
			m.Pulse = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pulse |= github_com_insolar_assured_ledger_ledger_core_v2_pulse.Number(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 24:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if byteLen > 0 {
				if dAtA[iNdEx] != 132 {
					return ErrExpectedBinaryMarkerPayload
				}
				iNdEx++
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 25:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OriginHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				if skippy, err = skipPayload(dAtA[iNdEx:]); err != nil {
					return err
				}
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPayload
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
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
				return 0, ErrInvalidLengthPayload
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPayload
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPayload
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPayload        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPayload = fmt.Errorf("proto: unexpected end of group")
	ErrExpectedBinaryMarkerPayload = fmt.Errorf("proto: binary marker was expected")
)
