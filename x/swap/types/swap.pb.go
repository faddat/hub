// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/swap/v1/swap.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Swap struct {
	TxHash   []byte     `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Receiver string     `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount   types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *Swap) Reset()         { *m = Swap{} }
func (m *Swap) String() string { return proto.CompactTextString(m) }
func (*Swap) ProtoMessage()    {}
func (*Swap) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c1c33cff7fdb4b2, []int{0}
}
func (m *Swap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Swap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Swap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Swap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Swap.Merge(m, src)
}
func (m *Swap) XXX_Size() int {
	return m.Size()
}
func (m *Swap) XXX_DiscardUnknown() {
	xxx_messageInfo_Swap.DiscardUnknown(m)
}

var xxx_messageInfo_Swap proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Swap)(nil), "sentinel.swap.v1.Swap")
}

func init() { proto.RegisterFile("sentinel/swap/v1/swap.proto", fileDescriptor_5c1c33cff7fdb4b2) }

var fileDescriptor_5c1c33cff7fdb4b2 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0x77, 0xf5, 0x82, 0x8a, 0x16, 0x86, 0x98, 0x88, 0x98, 0xac, 0xc4, 0x8a, 0xe6, 0x76,
	0x83, 0x16, 0xf6, 0x67, 0x63, 0x67, 0x82, 0x9d, 0x8d, 0x59, 0xc8, 0x1e, 0xbb, 0xc9, 0xb1, 0x43,
	0xd8, 0x85, 0xc3, 0xb7, 0xf0, 0x31, 0x7c, 0x14, 0xca, 0x2b, 0xad, 0x8c, 0xc2, 0x8b, 0x98, 0x83,
	0x3b, 0xab, 0xf9, 0xe7, 0x9f, 0x99, 0xe4, 0xfb, 0xc7, 0xbd, 0x36, 0x42, 0x5b, 0xa5, 0xc5, 0x8a,
	0x99, 0x35, 0x2f, 0x59, 0x13, 0x8f, 0x95, 0x96, 0x15, 0x58, 0xf0, 0xce, 0xf7, 0x43, 0x3a, 0x9a,
	0x4d, 0x1c, 0x90, 0x0c, 0x4c, 0x01, 0x86, 0xa5, 0xdc, 0x08, 0xd6, 0xc4, 0xa9, 0xb0, 0x3c, 0x66,
	0x19, 0x28, 0x3d, 0x5d, 0x04, 0x17, 0x39, 0xe4, 0x30, 0x4a, 0xb6, 0x55, 0x93, 0x7b, 0x6b, 0xdd,
	0xd9, 0xcb, 0x9a, 0x97, 0xde, 0xa5, 0x7b, 0x64, 0xdb, 0x37, 0xc9, 0x8d, 0xf4, 0x71, 0x88, 0xa3,
	0xb3, 0xc4, 0xb1, 0xed, 0x13, 0x37, 0xd2, 0x0b, 0xdc, 0xe3, 0x4a, 0x64, 0x42, 0x35, 0xa2, 0xf2,
	0x0f, 0x42, 0x1c, 0x9d, 0x24, 0xff, 0xbd, 0xf7, 0xe0, 0x3a, 0xbc, 0x80, 0x5a, 0x5b, 0xff, 0x30,
	0xc4, 0xd1, 0xe9, 0xdd, 0x15, 0x9d, 0x18, 0xe8, 0x96, 0x81, 0xee, 0x18, 0xe8, 0x23, 0x28, 0xbd,
	0x98, 0x75, 0xdf, 0x37, 0x28, 0xd9, 0xad, 0x2f, 0x9e, 0xbb, 0x5f, 0x82, 0x3e, 0x7b, 0x82, 0xba,
	0x9e, 0xe0, 0x4d, 0x4f, 0xf0, 0x4f, 0x4f, 0xf0, 0xc7, 0x40, 0xd0, 0x66, 0x20, 0xe8, 0x6b, 0x20,
	0xe8, 0x75, 0x9e, 0x2b, 0x2b, 0xeb, 0x94, 0x66, 0x50, 0xb0, 0x7d, 0xd4, 0x39, 0x2c, 0x97, 0x2a,
	0x53, 0x7c, 0xc5, 0x64, 0x9d, 0xb2, 0x76, 0x7a, 0x8b, 0x7d, 0x2f, 0x85, 0x49, 0x9d, 0x31, 0xcd,
	0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x12, 0x68, 0x52, 0x34, 0x01, 0x00, 0x00,
}

func (m *Swap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Swap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Swap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovSwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Swap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovSwap(uint64(l))
	return n
}

func sovSwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSwap(x uint64) (n int) {
	return sovSwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Swap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
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
			return fmt.Errorf("proto: Swap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Swap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
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
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
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
func skipSwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwap
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
					return 0, ErrIntOverflowSwap
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
					return 0, ErrIntOverflowSwap
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
				return 0, ErrInvalidLengthSwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSwap = fmt.Errorf("proto: unexpected end of group")
)
