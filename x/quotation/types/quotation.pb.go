// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quotationmvp/quotation/quotation.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Quotation struct {
	Id       string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tx       *Transaction `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	Envelope *Envelope    `protobuf:"bytes,3,opt,name=envelope,proto3" json:"envelope,omitempty"`
	Version  string       `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Creator  string       `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Quotation) Reset()         { *m = Quotation{} }
func (m *Quotation) String() string { return proto.CompactTextString(m) }
func (*Quotation) ProtoMessage()    {}
func (*Quotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d90be4b45d49a8, []int{0}
}
func (m *Quotation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Quotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Quotation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Quotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quotation.Merge(m, src)
}
func (m *Quotation) XXX_Size() int {
	return m.Size()
}
func (m *Quotation) XXX_DiscardUnknown() {
	xxx_messageInfo_Quotation.DiscardUnknown(m)
}

var xxx_messageInfo_Quotation proto.InternalMessageInfo

func (m *Quotation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Quotation) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *Quotation) GetEnvelope() *Envelope {
	if m != nil {
		return m.Envelope
	}
	return nil
}

func (m *Quotation) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Quotation) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Quotation)(nil), "quotationmvp.quotation.Quotation")
}

func init() {
	proto.RegisterFile("quotationmvp/quotation/quotation.proto", fileDescriptor_77d90be4b45d49a8)
}

var fileDescriptor_77d90be4b45d49a8 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x2c, 0xcd, 0x2f,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xcb, 0x2d, 0x2b, 0xd0, 0x87, 0x73, 0x10, 0x2c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x31, 0x64, 0x75, 0x7a, 0x70, 0x8e, 0x94, 0x2a, 0x0e, 0xfd, 0xa9, 0x79,
	0x65, 0xa9, 0x39, 0xf9, 0x05, 0xa9, 0x10, 0xed, 0x52, 0x1a, 0x38, 0x94, 0x95, 0x14, 0x25, 0xe6,
	0x15, 0x27, 0x26, 0x23, 0x2c, 0x52, 0x3a, 0xc4, 0xc8, 0xc5, 0x19, 0x08, 0x93, 0x17, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x32, 0xe6,
	0x62, 0x2a, 0xa9, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd6, 0xc3, 0xee, 0x26, 0xbd,
	0x10, 0x84, 0xa1, 0x41, 0x4c, 0x25, 0x15, 0x42, 0x36, 0x5c, 0x1c, 0x30, 0xe7, 0x48, 0x30, 0x83,
	0xb5, 0x2a, 0xe0, 0xd2, 0xea, 0x0a, 0x55, 0x17, 0x04, 0xd7, 0x21, 0x24, 0xc1, 0xc5, 0x5e, 0x96,
	0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x27, 0xc1, 0x02, 0x76, 0x07, 0x8c, 0x0b, 0x92, 0x49, 0x2e, 0x4a,
	0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x85, 0xc8, 0x40, 0xb9, 0x4e, 0x96, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x0f, 0xb7, 0x4b, 0x17, 0x14, 0x12, 0x15, 0xc8, 0x61,
	0x51, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0xe8, 0x9d, 0x59, 0x99, 0x01, 0x00, 0x00,
}

func (m *Quotation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Quotation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Quotation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuotation(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintQuotation(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x22
	}
	if m.Envelope != nil {
		{
			size, err := m.Envelope.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuotation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuotation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuotation(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuotation(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuotation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Quotation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuotation(uint64(l))
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovQuotation(uint64(l))
	}
	if m.Envelope != nil {
		l = m.Envelope.Size()
		n += 1 + l + sovQuotation(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovQuotation(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuotation(uint64(l))
	}
	return n
}

func sovQuotation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuotation(x uint64) (n int) {
	return sovQuotation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Quotation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuotation
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
			return fmt.Errorf("proto: Quotation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quotation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotation
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
				return ErrInvalidLengthQuotation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuotation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotation
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
				return ErrInvalidLengthQuotation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuotation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &Transaction{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelope", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotation
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
				return ErrInvalidLengthQuotation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuotation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Envelope == nil {
				m.Envelope = &Envelope{}
			}
			if err := m.Envelope.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotation
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
				return ErrInvalidLengthQuotation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuotation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuotation
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
				return ErrInvalidLengthQuotation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuotation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuotation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuotation
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
func skipQuotation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuotation
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
					return 0, ErrIntOverflowQuotation
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
					return 0, ErrIntOverflowQuotation
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
				return 0, ErrInvalidLengthQuotation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuotation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuotation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuotation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuotation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuotation = fmt.Errorf("proto: unexpected end of group")
)
