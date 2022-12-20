// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/setting/settingRes.proto

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

type SettingRes struct {
	EditName             bool     `protobuf:"varint,1,opt,name=editName,proto3" json:"editName,omitempty"`
	EditSex              bool     `protobuf:"varint,2,opt,name=editSex,proto3" json:"editSex,omitempty"`
	Bgm                  bool     `protobuf:"varint,3,opt,name=bgm,proto3" json:"bgm,omitempty"`
	MusicEffect          bool     `protobuf:"varint,4,opt,name=musicEffect,proto3" json:"musicEffect,omitempty"`
	Push                 bool     `protobuf:"varint,5,opt,name=push,proto3" json:"push,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingRes) Reset()         { *m = SettingRes{} }
func (m *SettingRes) String() string { return proto.CompactTextString(m) }
func (*SettingRes) ProtoMessage()    {}
func (*SettingRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_670a7c0f88074b82, []int{0}
}
func (m *SettingRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SettingRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SettingRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SettingRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingRes.Merge(m, src)
}
func (m *SettingRes) XXX_Size() int {
	return m.Size()
}
func (m *SettingRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingRes.DiscardUnknown(m)
}

var xxx_messageInfo_SettingRes proto.InternalMessageInfo

func (m *SettingRes) GetEditName() bool {
	if m != nil {
		return m.EditName
	}
	return false
}

func (m *SettingRes) GetEditSex() bool {
	if m != nil {
		return m.EditSex
	}
	return false
}

func (m *SettingRes) GetBgm() bool {
	if m != nil {
		return m.Bgm
	}
	return false
}

func (m *SettingRes) GetMusicEffect() bool {
	if m != nil {
		return m.MusicEffect
	}
	return false
}

func (m *SettingRes) GetPush() bool {
	if m != nil {
		return m.Push
	}
	return false
}

func init() {
	proto.RegisterType((*SettingRes)(nil), "pb.setting.SettingRes")
}

func init() { proto.RegisterFile("protocol/setting/settingRes.proto", fileDescriptor_670a7c0f88074b82) }

var fileDescriptor_670a7c0f88074b82 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4e, 0x2d, 0x29, 0xc9, 0xcc, 0x4b, 0x87, 0xd1, 0x41, 0xa9,
	0xc5, 0x7a, 0x60, 0x39, 0x21, 0xae, 0x82, 0x24, 0x3d, 0xa8, 0xa0, 0x52, 0x17, 0x23, 0x17, 0x57,
	0x30, 0x5c, 0x81, 0x90, 0x14, 0x17, 0x47, 0x6a, 0x4a, 0x66, 0x89, 0x5f, 0x62, 0x6e, 0xaa, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x9c, 0x2f, 0x24, 0xc1, 0xc5, 0x0e, 0x62, 0x07, 0xa7, 0x56,
	0x48, 0x30, 0x81, 0xa5, 0x60, 0x5c, 0x21, 0x01, 0x2e, 0xe6, 0xa4, 0xf4, 0x5c, 0x09, 0x66, 0xb0,
	0x28, 0x88, 0x29, 0xa4, 0xc0, 0xc5, 0x9d, 0x5b, 0x5a, 0x9c, 0x99, 0xec, 0x9a, 0x96, 0x96, 0x9a,
	0x5c, 0x22, 0xc1, 0x02, 0x96, 0x41, 0x16, 0x12, 0x12, 0xe2, 0x62, 0x29, 0x28, 0x2d, 0xce, 0x90,
	0x60, 0x05, 0x4b, 0x81, 0xd9, 0x4e, 0x72, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0x3c, 0xfa, 0x99, 0x79, 0x25, 0xa9, 0x45,
	0x79, 0x89, 0x39, 0xfa, 0x05, 0x49, 0x49, 0x6c, 0x60, 0xf7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0x1e, 0xf2, 0x29, 0xe4, 0x00, 0x00, 0x00,
}

func (m *SettingRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SettingRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SettingRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Push {
		i--
		if m.Push {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.MusicEffect {
		i--
		if m.MusicEffect {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Bgm {
		i--
		if m.Bgm {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.EditSex {
		i--
		if m.EditSex {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.EditName {
		i--
		if m.EditName {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSettingRes(dAtA []byte, offset int, v uint64) int {
	offset -= sovSettingRes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SettingRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EditName {
		n += 2
	}
	if m.EditSex {
		n += 2
	}
	if m.Bgm {
		n += 2
	}
	if m.MusicEffect {
		n += 2
	}
	if m.Push {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSettingRes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSettingRes(x uint64) (n int) {
	return sovSettingRes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SettingRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettingRes
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
			return fmt.Errorf("proto: SettingRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SettingRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EditName", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettingRes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EditName = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EditSex", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettingRes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EditSex = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bgm", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettingRes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Bgm = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MusicEffect", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettingRes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MusicEffect = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Push", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettingRes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Push = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSettingRes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSettingRes
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
func skipSettingRes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSettingRes
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
					return 0, ErrIntOverflowSettingRes
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
					return 0, ErrIntOverflowSettingRes
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
				return 0, ErrInvalidLengthSettingRes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSettingRes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSettingRes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSettingRes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSettingRes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSettingRes = fmt.Errorf("proto: unexpected end of group")
)
