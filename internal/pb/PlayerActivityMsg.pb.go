// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/activity/PlayerActivityMsg.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
// A compilation merror at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PlayerActivityMsg struct {
	OpenTime             int32                     `protobuf:"varint,1,opt,name=openTime,proto3" json:"openTime,omitempty"`
	RegisterTime         int32                     `protobuf:"varint,2,opt,name=registerTime,proto3" json:"registerTime,omitempty"`
	FirstRecharge        *ActivityFirstRechargeMsg `protobuf:"bytes,4,opt,name=firstRecharge,proto3" json:"firstRecharge,omitempty"`
	EverydaySign         *ActivityEverydaySignMsg  `protobuf:"bytes,5,opt,name=everydaySign,proto3" json:"everydaySign,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PlayerActivityMsg) Reset()         { *m = PlayerActivityMsg{} }
func (m *PlayerActivityMsg) String() string { return proto.CompactTextString(m) }
func (*PlayerActivityMsg) ProtoMessage()    {}
func (*PlayerActivityMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e71fe9125809bc60, []int{0}
}
func (m *PlayerActivityMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlayerActivityMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlayerActivityMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlayerActivityMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerActivityMsg.Merge(m, src)
}
func (m *PlayerActivityMsg) XXX_Size() int {
	return m.Size()
}
func (m *PlayerActivityMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerActivityMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerActivityMsg proto.InternalMessageInfo

func (m *PlayerActivityMsg) GetOpenTime() int32 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *PlayerActivityMsg) GetRegisterTime() int32 {
	if m != nil {
		return m.RegisterTime
	}
	return 0
}

func (m *PlayerActivityMsg) GetFirstRecharge() *ActivityFirstRechargeMsg {
	if m != nil {
		return m.FirstRecharge
	}
	return nil
}

func (m *PlayerActivityMsg) GetEverydaySign() *ActivityEverydaySignMsg {
	if m != nil {
		return m.EverydaySign
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerActivityMsg)(nil), "pb.activity.PlayerActivityMsg")
}

func init() {
	proto.RegisterFile("protocol/activity/PlayerActivityMsg.proto", fileDescriptor_e71fe9125809bc60)
}

var fileDescriptor_e71fe9125809bc60 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x4c, 0x2e, 0xc9, 0x2c, 0xcb, 0x2c, 0xa9, 0xd4, 0x0f, 0xc8,
	0x49, 0xac, 0x4c, 0x2d, 0x72, 0x84, 0x72, 0x7d, 0x8b, 0xd3, 0xf5, 0xc0, 0x6a, 0x84, 0xb8, 0x0b,
	0x92, 0xf4, 0x60, 0x8a, 0xa4, 0xf4, 0x31, 0xf5, 0xc1, 0x74, 0xb8, 0x96, 0xa5, 0x16, 0x55, 0xa6,
	0x24, 0x56, 0x06, 0x67, 0xa6, 0xe7, 0xc1, 0x75, 0x4b, 0x19, 0xe0, 0xd6, 0xe0, 0x96, 0x59, 0x54,
	0x5c, 0x12, 0x94, 0x9a, 0x9c, 0x91, 0x58, 0x94, 0x9e, 0x0a, 0xd7, 0xa1, 0xf4, 0x8a, 0x91, 0x4b,
	0x10, 0xc3, 0x2d, 0x42, 0x52, 0x5c, 0x1c, 0xf9, 0x05, 0xa9, 0x79, 0x21, 0x99, 0xb9, 0xa9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x70, 0xbe, 0x90, 0x12, 0x17, 0x4f, 0x51, 0x6a, 0x7a, 0x66,
	0x71, 0x49, 0x6a, 0x11, 0x58, 0x9e, 0x09, 0x2c, 0x8f, 0x22, 0x26, 0xe4, 0xcd, 0xc5, 0x9b, 0x86,
	0x6c, 0x9f, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xaa, 0x1e, 0x92, 0xef, 0xf4, 0x70, 0xb9,
	0x2c, 0x08, 0x55, 0xaf, 0x90, 0x07, 0x17, 0x4f, 0x2a, 0x92, 0x6f, 0x25, 0x58, 0xc1, 0x66, 0xa9,
	0x60, 0x35, 0x0b, 0x2d, 0x58, 0x82, 0x50, 0x74, 0x3a, 0xc9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xf1, 0xe8, 0x67, 0xe6,
	0x95, 0xa4, 0x16, 0xe5, 0x25, 0xe6, 0xe8, 0x17, 0x24, 0x25, 0xb1, 0x81, 0xc3, 0xc4, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x55, 0x54, 0xdc, 0x6b, 0xb0, 0x01, 0x00, 0x00,
}

func (m *PlayerActivityMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerActivityMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlayerActivityMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EverydaySign != nil {
		{
			size, err := m.EverydaySign.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlayerActivityMsg(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.FirstRecharge != nil {
		{
			size, err := m.FirstRecharge.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPlayerActivityMsg(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.RegisterTime != 0 {
		i = encodeVarintPlayerActivityMsg(dAtA, i, uint64(m.RegisterTime))
		i--
		dAtA[i] = 0x10
	}
	if m.OpenTime != 0 {
		i = encodeVarintPlayerActivityMsg(dAtA, i, uint64(m.OpenTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlayerActivityMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlayerActivityMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PlayerActivityMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OpenTime != 0 {
		n += 1 + sovPlayerActivityMsg(uint64(m.OpenTime))
	}
	if m.RegisterTime != 0 {
		n += 1 + sovPlayerActivityMsg(uint64(m.RegisterTime))
	}
	if m.FirstRecharge != nil {
		l = m.FirstRecharge.Size()
		n += 1 + l + sovPlayerActivityMsg(uint64(l))
	}
	if m.EverydaySign != nil {
		l = m.EverydaySign.Size()
		n += 1 + l + sovPlayerActivityMsg(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPlayerActivityMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlayerActivityMsg(x uint64) (n int) {
	return sovPlayerActivityMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerActivityMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayerActivityMsg
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
			return fmt.Errorf("proto: PlayerActivityMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerActivityMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenTime", wireType)
			}
			m.OpenTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerActivityMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpenTime |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterTime", wireType)
			}
			m.RegisterTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerActivityMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegisterTime |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstRecharge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerActivityMsg
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
				return ErrInvalidLengthPlayerActivityMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerActivityMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FirstRecharge == nil {
				m.FirstRecharge = &ActivityFirstRechargeMsg{}
			}
			if err := m.FirstRecharge.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EverydaySign", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayerActivityMsg
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
				return ErrInvalidLengthPlayerActivityMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayerActivityMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EverydaySign == nil {
				m.EverydaySign = &ActivityEverydaySignMsg{}
			}
			if err := m.EverydaySign.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlayerActivityMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlayerActivityMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPlayerActivityMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlayerActivityMsg
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
					return 0, ErrIntOverflowPlayerActivityMsg
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
					return 0, ErrIntOverflowPlayerActivityMsg
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
				return 0, ErrInvalidLengthPlayerActivityMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlayerActivityMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlayerActivityMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlayerActivityMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlayerActivityMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlayerActivityMsg = fmt.Errorf("proto: unexpected end of group")
)
