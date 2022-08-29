// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/activity/ActivityEverydaySignMsg.proto

package activity

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActivityEverydaySignMsg struct {
	ResActivityId        int32    `protobuf:"varint,1,opt,name=resActivityId,proto3" json:"resActivityId,omitempty"`
	ReceiveTime          int32    `protobuf:"varint,2,opt,name=receiveTime,proto3" json:"receiveTime,omitempty"`
	WeekNum              int32    `protobuf:"varint,3,opt,name=weekNum,proto3" json:"weekNum,omitempty"`
	Times                int32    `protobuf:"varint,4,opt,name=times,proto3" json:"times,omitempty"`
	ReceiveDoubleTime    int32    `protobuf:"varint,5,opt,name=receiveDoubleTime,proto3" json:"receiveDoubleTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityEverydaySignMsg) Reset()         { *m = ActivityEverydaySignMsg{} }
func (m *ActivityEverydaySignMsg) String() string { return proto.CompactTextString(m) }
func (*ActivityEverydaySignMsg) ProtoMessage()    {}
func (*ActivityEverydaySignMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_377422d962bec354, []int{0}
}
func (m *ActivityEverydaySignMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivityEverydaySignMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivityEverydaySignMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivityEverydaySignMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityEverydaySignMsg.Merge(m, src)
}
func (m *ActivityEverydaySignMsg) XXX_Size() int {
	return m.Size()
}
func (m *ActivityEverydaySignMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityEverydaySignMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityEverydaySignMsg proto.InternalMessageInfo

func (m *ActivityEverydaySignMsg) GetResActivityId() int32 {
	if m != nil {
		return m.ResActivityId
	}
	return 0
}

func (m *ActivityEverydaySignMsg) GetReceiveTime() int32 {
	if m != nil {
		return m.ReceiveTime
	}
	return 0
}

func (m *ActivityEverydaySignMsg) GetWeekNum() int32 {
	if m != nil {
		return m.WeekNum
	}
	return 0
}

func (m *ActivityEverydaySignMsg) GetTimes() int32 {
	if m != nil {
		return m.Times
	}
	return 0
}

func (m *ActivityEverydaySignMsg) GetReceiveDoubleTime() int32 {
	if m != nil {
		return m.ReceiveDoubleTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ActivityEverydaySignMsg)(nil), "pb.activity.ActivityEverydaySignMsg")
}

func init() {
	proto.RegisterFile("protocol/activity/ActivityEverydaySignMsg.proto", fileDescriptor_377422d962bec354)
}

var fileDescriptor_377422d962bec354 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x4c, 0x2e, 0xc9, 0x2c, 0xcb, 0x2c, 0xa9, 0xd4, 0x77, 0x84,
	0x32, 0x5c, 0xcb, 0x52, 0x8b, 0x2a, 0x53, 0x12, 0x2b, 0x83, 0x33, 0xd3, 0xf3, 0x7c, 0x8b, 0xd3,
	0xf5, 0xc0, 0x2a, 0x85, 0xb8, 0x0b, 0x92, 0xf4, 0x60, 0x4a, 0x95, 0xf6, 0x33, 0x72, 0x89, 0xe3,
	0x50, 0x2e, 0xa4, 0xc2, 0xc5, 0x5b, 0x94, 0x5a, 0x0c, 0x93, 0xf5, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x42, 0x15, 0x14, 0x52, 0xe0, 0xe2, 0x2e, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x4b,
	0x0d, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0xab, 0x41, 0x16, 0x12, 0x92, 0xe0, 0x62, 0x2f, 0x4f,
	0x4d, 0xcd, 0xf6, 0x2b, 0xcd, 0x95, 0x60, 0x06, 0xcb, 0xc2, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x25,
	0x99, 0xb9, 0xa9, 0xc5, 0x12, 0x2c, 0x60, 0x71, 0x08, 0x47, 0x48, 0x87, 0x4b, 0x10, 0xaa, 0xdd,
	0x25, 0xbf, 0x34, 0x29, 0x07, 0x62, 0x2e, 0x2b, 0x58, 0x05, 0xa6, 0x84, 0x93, 0xfa, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0xa8, 0x7e, 0x66, 0x5e, 0x49, 0x6a, 0x51, 0x5e, 0x62, 0x8e, 0x7e, 0x41, 0x12, 0x3c, 0x54, 0x92,
	0xd8, 0xc0, 0xde, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x19, 0xf5, 0xe5, 0x9a, 0x31, 0x01,
	0x00, 0x00,
}

func (m *ActivityEverydaySignMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivityEverydaySignMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityEverydaySignMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ReceiveDoubleTime != 0 {
		i = encodeVarintActivityEverydaySignMsg(dAtA, i, uint64(m.ReceiveDoubleTime))
		i--
		dAtA[i] = 0x28
	}
	if m.Times != 0 {
		i = encodeVarintActivityEverydaySignMsg(dAtA, i, uint64(m.Times))
		i--
		dAtA[i] = 0x20
	}
	if m.WeekNum != 0 {
		i = encodeVarintActivityEverydaySignMsg(dAtA, i, uint64(m.WeekNum))
		i--
		dAtA[i] = 0x18
	}
	if m.ReceiveTime != 0 {
		i = encodeVarintActivityEverydaySignMsg(dAtA, i, uint64(m.ReceiveTime))
		i--
		dAtA[i] = 0x10
	}
	if m.ResActivityId != 0 {
		i = encodeVarintActivityEverydaySignMsg(dAtA, i, uint64(m.ResActivityId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintActivityEverydaySignMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovActivityEverydaySignMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActivityEverydaySignMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResActivityId != 0 {
		n += 1 + sovActivityEverydaySignMsg(uint64(m.ResActivityId))
	}
	if m.ReceiveTime != 0 {
		n += 1 + sovActivityEverydaySignMsg(uint64(m.ReceiveTime))
	}
	if m.WeekNum != 0 {
		n += 1 + sovActivityEverydaySignMsg(uint64(m.WeekNum))
	}
	if m.Times != 0 {
		n += 1 + sovActivityEverydaySignMsg(uint64(m.Times))
	}
	if m.ReceiveDoubleTime != 0 {
		n += 1 + sovActivityEverydaySignMsg(uint64(m.ReceiveDoubleTime))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovActivityEverydaySignMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActivityEverydaySignMsg(x uint64) (n int) {
	return sovActivityEverydaySignMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActivityEverydaySignMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityEverydaySignMsg
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
			return fmt.Errorf("proto: ActivityEverydaySignMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivityEverydaySignMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResActivityId", wireType)
			}
			m.ResActivityId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityEverydaySignMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResActivityId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiveTime", wireType)
			}
			m.ReceiveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityEverydaySignMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceiveTime |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeekNum", wireType)
			}
			m.WeekNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityEverydaySignMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WeekNum |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Times", wireType)
			}
			m.Times = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityEverydaySignMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Times |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiveDoubleTime", wireType)
			}
			m.ReceiveDoubleTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityEverydaySignMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceiveDoubleTime |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActivityEverydaySignMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityEverydaySignMsg
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
func skipActivityEverydaySignMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActivityEverydaySignMsg
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
					return 0, ErrIntOverflowActivityEverydaySignMsg
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
					return 0, ErrIntOverflowActivityEverydaySignMsg
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
				return 0, ErrInvalidLengthActivityEverydaySignMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActivityEverydaySignMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActivityEverydaySignMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActivityEverydaySignMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActivityEverydaySignMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActivityEverydaySignMsg = fmt.Errorf("proto: unexpected end of group")
)
