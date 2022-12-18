// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/login/loginReq.proto

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LoginReq struct {
	Passport  string `protobuf:"bytes,1,opt,name=passport,proto3" json:"passport,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ChannelId int32  `protobuf:"varint,3,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Sex       int32  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Sign      string `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign,omitempty"`
	Timestamp int32  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	LoginType int32  `protobuf:"varint,7,opt,name=loginType,proto3" json:"loginType,omitempty"`
	FromData  string `protobuf:"bytes,8,opt,name=fromData,proto3" json:"fromData,omitempty"`
	Subscribe bool   `protobuf:"varint,9,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	// 统计相关
	Imei         string `protobuf:"bytes,20,opt,name=imei,proto3" json:"imei,omitempty"`
	DeviceModel  string `protobuf:"bytes,21,opt,name=deviceModel,proto3" json:"deviceModel,omitempty"`
	OsName       string `protobuf:"bytes,22,opt,name=osName,proto3" json:"osName,omitempty"`
	OsVer        string `protobuf:"bytes,23,opt,name=osVer,proto3" json:"osVer,omitempty"`
	MacAddr      string `protobuf:"bytes,24,opt,name=macAddr,proto3" json:"macAddr,omitempty"`
	DeviceHeight int32  `protobuf:"varint,25,opt,name=deviceHeight,proto3" json:"deviceHeight,omitempty"`
	DeviceWidth  int32  `protobuf:"varint,26,opt,name=deviceWidth,proto3" json:"deviceWidth,omitempty"`
	Isp          string `protobuf:"bytes,27,opt,name=isp,proto3" json:"isp,omitempty"`
	Network      string `protobuf:"bytes,28,opt,name=network,proto3" json:"network,omitempty"`
	Root         string `protobuf:"bytes,29,opt,name=root,proto3" json:"root,omitempty"`
	// 邀请进入
	FriendOpenId string `protobuf:"bytes,30,opt,name=friendOpenId,proto3" json:"friendOpenId,omitempty"`
	ShareZone    int32  `protobuf:"varint,33,opt,name=shareZone,proto3" json:"shareZone,omitempty"`
	// loadingGame
	LoadingGameMoney     int32    `protobuf:"varint,34,opt,name=loadingGameMoney,proto3" json:"loadingGameMoney,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a86780ef438e533, []int{0}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return m.Size()
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetPassport() string {
	if m != nil {
		return m.Passport
	}
	return ""
}

func (m *LoginReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginReq) GetChannelId() int32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *LoginReq) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *LoginReq) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *LoginReq) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *LoginReq) GetLoginType() int32 {
	if m != nil {
		return m.LoginType
	}
	return 0
}

func (m *LoginReq) GetFromData() string {
	if m != nil {
		return m.FromData
	}
	return ""
}

func (m *LoginReq) GetSubscribe() bool {
	if m != nil {
		return m.Subscribe
	}
	return false
}

func (m *LoginReq) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *LoginReq) GetDeviceModel() string {
	if m != nil {
		return m.DeviceModel
	}
	return ""
}

func (m *LoginReq) GetOsName() string {
	if m != nil {
		return m.OsName
	}
	return ""
}

func (m *LoginReq) GetOsVer() string {
	if m != nil {
		return m.OsVer
	}
	return ""
}

func (m *LoginReq) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func (m *LoginReq) GetDeviceHeight() int32 {
	if m != nil {
		return m.DeviceHeight
	}
	return 0
}

func (m *LoginReq) GetDeviceWidth() int32 {
	if m != nil {
		return m.DeviceWidth
	}
	return 0
}

func (m *LoginReq) GetIsp() string {
	if m != nil {
		return m.Isp
	}
	return ""
}

func (m *LoginReq) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *LoginReq) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func (m *LoginReq) GetFriendOpenId() string {
	if m != nil {
		return m.FriendOpenId
	}
	return ""
}

func (m *LoginReq) GetShareZone() int32 {
	if m != nil {
		return m.ShareZone
	}
	return 0
}

func (m *LoginReq) GetLoadingGameMoney() int32 {
	if m != nil {
		return m.LoadingGameMoney
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "pb.login.LoginReq")
}

func init() { proto.RegisterFile("protocol/login/loginReq.proto", fileDescriptor_2a86780ef438e533) }

var fileDescriptor_2a86780ef438e533 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x59, 0x9a, 0xa4, 0x1b, 0xd3, 0x43, 0x65, 0x95, 0x32, 0x94, 0x76, 0x15, 0x72, 0x8a,
	0x38, 0x34, 0x07, 0x9e, 0x00, 0x84, 0x04, 0x95, 0x28, 0x48, 0x11, 0x02, 0xa9, 0x37, 0xef, 0x7a,
	0x9a, 0x58, 0xec, 0xda, 0xc6, 0x36, 0x7f, 0xfa, 0x26, 0x3c, 0x12, 0x47, 0x1e, 0x01, 0x05, 0x89,
	0xe7, 0x40, 0x33, 0x4e, 0x37, 0xad, 0x7a, 0x89, 0xe6, 0xfb, 0x7d, 0xce, 0xe7, 0xcf, 0xa3, 0x15,
	0x27, 0x3e, 0xb8, 0xe4, 0x1a, 0xd7, 0xce, 0x5b, 0xb7, 0x34, 0x36, 0xff, 0x2e, 0xf0, 0xcb, 0x29,
	0x73, 0x59, 0xfa, 0xfa, 0x94, 0xd1, 0xf4, 0xdf, 0x40, 0x94, 0x6f, 0x37, 0xa6, 0x3c, 0x12, 0xa5,
	0x57, 0x31, 0x7a, 0x17, 0x12, 0x14, 0x93, 0x62, 0x36, 0x5e, 0xf4, 0x5a, 0x4a, 0x31, 0xb0, 0xaa,
	0x43, 0xb8, 0xcf, 0x9c, 0x67, 0x79, 0x2c, 0xc6, 0xcd, 0x4a, 0x59, 0x8b, 0xed, 0x99, 0x86, 0x9d,
	0x49, 0x31, 0x1b, 0x2e, 0xb6, 0x40, 0xee, 0x8b, 0x9d, 0x88, 0x3f, 0x60, 0xc0, 0x9c, 0x46, 0xca,
	0x88, 0x66, 0x69, 0x61, 0x98, 0x33, 0x68, 0xa6, 0x8c, 0x64, 0x3a, 0x8c, 0x49, 0x75, 0x1e, 0x46,
	0x39, 0xa3, 0x07, 0xe4, 0x72, 0xcf, 0x0f, 0x57, 0x1e, 0x61, 0x37, 0xbb, 0x3d, 0xa0, 0xbe, 0x97,
	0xc1, 0x75, 0xaf, 0x54, 0x52, 0x50, 0xe6, 0xbe, 0xd7, 0x9a, 0xfe, 0x19, 0xbf, 0xd6, 0xb1, 0x09,
	0xa6, 0x46, 0x18, 0x4f, 0x8a, 0x59, 0xb9, 0xd8, 0x02, 0x6a, 0x62, 0x3a, 0x34, 0x70, 0x90, 0x9b,
	0xd0, 0x2c, 0x27, 0xe2, 0x81, 0xc6, 0x6f, 0xa6, 0xc1, 0x73, 0xa7, 0xb1, 0x85, 0x87, 0x6c, 0xdd,
	0x44, 0xf2, 0x50, 0x8c, 0x5c, 0x7c, 0x47, 0x5b, 0x38, 0x64, 0x73, 0xa3, 0xe4, 0x81, 0x18, 0xba,
	0xf8, 0x11, 0x03, 0x3c, 0x62, 0x9c, 0x85, 0x04, 0xb1, 0xdb, 0xa9, 0xe6, 0x85, 0xd6, 0x01, 0x80,
	0xf9, 0xb5, 0x94, 0x53, 0xb1, 0x97, 0x63, 0xdf, 0xa0, 0x59, 0xae, 0x12, 0x3c, 0xe6, 0x87, 0xdd,
	0x62, 0xdb, 0x36, 0x9f, 0x8c, 0x4e, 0x2b, 0x38, 0xe2, 0x23, 0x37, 0x11, 0xed, 0xd7, 0x44, 0x0f,
	0x4f, 0x38, 0x9b, 0x46, 0xba, 0xd1, 0x62, 0xfa, 0xee, 0xc2, 0x67, 0x38, 0xce, 0x37, 0x6e, 0x24,
	0xbd, 0x37, 0x38, 0x97, 0xe0, 0x24, 0xbf, 0x97, 0x66, 0x6a, 0x71, 0x19, 0x0c, 0x5a, 0xfd, 0xde,
	0xa3, 0x3d, 0xd3, 0x50, 0xb1, 0x77, 0x8b, 0xf1, 0x16, 0x57, 0x2a, 0xe0, 0x85, 0xb3, 0x08, 0x4f,
	0xf3, 0xfe, 0x7b, 0x20, 0x9f, 0x89, 0xfd, 0xd6, 0x29, 0x6d, 0xec, 0xf2, 0xb5, 0xea, 0xf0, 0xdc,
	0x59, 0xbc, 0x82, 0x29, 0x1f, 0xba, 0xc3, 0x5f, 0x56, 0xbf, 0xd6, 0x55, 0xf1, 0x7b, 0x5d, 0x15,
	0x7f, 0xd6, 0x55, 0xf1, 0xf3, 0x6f, 0x75, 0xef, 0x62, 0x6f, 0x6e, 0x6c, 0xc2, 0x60, 0x55, 0x3b,
	0xf7, 0x75, 0x3d, 0xe2, 0x2f, 0xf3, 0xf9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xf3, 0x1d,
	0xc8, 0xba, 0x02, 0x00, 0x00,
}

func (m *LoginReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LoadingGameMoney != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.LoadingGameMoney))
		i--
		dAtA[i] = 0x2
		i--
		dAtA[i] = 0x90
	}
	if m.ShareZone != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.ShareZone))
		i--
		dAtA[i] = 0x2
		i--
		dAtA[i] = 0x88
	}
	if len(m.FriendOpenId) > 0 {
		i -= len(m.FriendOpenId)
		copy(dAtA[i:], m.FriendOpenId)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.FriendOpenId)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xf2
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xea
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xe2
	}
	if len(m.Isp) > 0 {
		i -= len(m.Isp)
		copy(dAtA[i:], m.Isp)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.Isp)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xda
	}
	if m.DeviceWidth != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.DeviceWidth))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xd0
	}
	if m.DeviceHeight != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.DeviceHeight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc8
	}
	if len(m.MacAddr) > 0 {
		i -= len(m.MacAddr)
		copy(dAtA[i:], m.MacAddr)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.MacAddr)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc2
	}
	if len(m.OsVer) > 0 {
		i -= len(m.OsVer)
		copy(dAtA[i:], m.OsVer)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.OsVer)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xba
	}
	if len(m.OsName) > 0 {
		i -= len(m.OsName)
		copy(dAtA[i:], m.OsName)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.OsName)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb2
	}
	if len(m.DeviceModel) > 0 {
		i -= len(m.DeviceModel)
		copy(dAtA[i:], m.DeviceModel)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.DeviceModel)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if len(m.Imei) > 0 {
		i -= len(m.Imei)
		copy(dAtA[i:], m.Imei)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.Imei)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if m.Subscribe {
		i--
		if m.Subscribe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.FromData) > 0 {
		i -= len(m.FromData)
		copy(dAtA[i:], m.FromData)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.FromData)))
		i--
		dAtA[i] = 0x42
	}
	if m.LoginType != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.LoginType))
		i--
		dAtA[i] = 0x38
	}
	if m.Timestamp != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Sex != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.Sex))
		i--
		dAtA[i] = 0x20
	}
	if m.ChannelId != 0 {
		i = encodeVarintLoginReq(dAtA, i, uint64(m.ChannelId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Passport) > 0 {
		i -= len(m.Passport)
		copy(dAtA[i:], m.Passport)
		i = encodeVarintLoginReq(dAtA, i, uint64(len(m.Passport)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoginReq(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoginReq(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoginReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Passport)
	if l > 0 {
		n += 1 + l + sovLoginReq(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLoginReq(uint64(l))
	}
	if m.ChannelId != 0 {
		n += 1 + sovLoginReq(uint64(m.ChannelId))
	}
	if m.Sex != 0 {
		n += 1 + sovLoginReq(uint64(m.Sex))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovLoginReq(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovLoginReq(uint64(m.Timestamp))
	}
	if m.LoginType != 0 {
		n += 1 + sovLoginReq(uint64(m.LoginType))
	}
	l = len(m.FromData)
	if l > 0 {
		n += 1 + l + sovLoginReq(uint64(l))
	}
	if m.Subscribe {
		n += 2
	}
	l = len(m.Imei)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	l = len(m.DeviceModel)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	l = len(m.OsName)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	l = len(m.OsVer)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	l = len(m.MacAddr)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	if m.DeviceHeight != 0 {
		n += 2 + sovLoginReq(uint64(m.DeviceHeight))
	}
	if m.DeviceWidth != 0 {
		n += 2 + sovLoginReq(uint64(m.DeviceWidth))
	}
	l = len(m.Isp)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	l = len(m.Root)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	l = len(m.FriendOpenId)
	if l > 0 {
		n += 2 + l + sovLoginReq(uint64(l))
	}
	if m.ShareZone != 0 {
		n += 2 + sovLoginReq(uint64(m.ShareZone))
	}
	if m.LoadingGameMoney != 0 {
		n += 2 + sovLoginReq(uint64(m.LoadingGameMoney))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLoginReq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoginReq(x uint64) (n int) {
	return sovLoginReq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoginReq
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
			return fmt.Errorf("proto: LoginReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Passport", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Passport = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			m.ChannelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sex", wireType)
			}
			m.Sex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginType", wireType)
			}
			m.LoginType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoginType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
			m.Subscribe = bool(v != 0)
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Imei", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Imei = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceModel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceModel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OsName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsVer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OsVer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 24:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MacAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MacAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 25:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceHeight", wireType)
			}
			m.DeviceHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 26:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceWidth", wireType)
			}
			m.DeviceWidth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeviceWidth |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 27:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Isp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Isp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 28:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 29:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 30:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FriendOpenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
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
				return ErrInvalidLengthLoginReq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoginReq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FriendOpenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 33:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareZone", wireType)
			}
			m.ShareZone = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShareZone |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 34:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadingGameMoney", wireType)
			}
			m.LoadingGameMoney = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginReq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoadingGameMoney |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLoginReq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoginReq
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
func skipLoginReq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoginReq
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
					return 0, ErrIntOverflowLoginReq
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
					return 0, ErrIntOverflowLoginReq
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
				return 0, ErrInvalidLengthLoginReq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoginReq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoginReq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoginReq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoginReq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoginReq = fmt.Errorf("proto: unexpected end of group")
)