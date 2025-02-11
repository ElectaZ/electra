// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meter/meterreadings.proto

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

type Meterreadings struct {
	DeviceID  string `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Phase     uint64 `protobuf:"varint,3,opt,name=phase,proto3" json:"phase,omitempty"`
	Whin      uint64 `protobuf:"varint,4,opt,name=whin,proto3" json:"whin,omitempty"`
	Whout     uint64 `protobuf:"varint,5,opt,name=whout,proto3" json:"whout,omitempty"`
	Mvolt     uint64 `protobuf:"varint,6,opt,name=mvolt,proto3" json:"mvolt,omitempty"`
	Mhertz    uint64 `protobuf:"varint,7,opt,name=mhertz,proto3" json:"mhertz,omitempty"`
	Mpf       uint64 `protobuf:"varint,8,opt,name=mpf,proto3" json:"mpf,omitempty"`
	Maxmi     uint64 `protobuf:"varint,9,opt,name=maxmi,proto3" json:"maxmi,omitempty"`
}

func (m *Meterreadings) Reset()         { *m = Meterreadings{} }
func (m *Meterreadings) String() string { return proto.CompactTextString(m) }
func (*Meterreadings) ProtoMessage()    {}
func (*Meterreadings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ed5ce77a048220, []int{0}
}
func (m *Meterreadings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Meterreadings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Meterreadings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Meterreadings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meterreadings.Merge(m, src)
}
func (m *Meterreadings) XXX_Size() int {
	return m.Size()
}
func (m *Meterreadings) XXX_DiscardUnknown() {
	xxx_messageInfo_Meterreadings.DiscardUnknown(m)
}

var xxx_messageInfo_Meterreadings proto.InternalMessageInfo

func (m *Meterreadings) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *Meterreadings) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Meterreadings) GetPhase() uint64 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *Meterreadings) GetWhin() uint64 {
	if m != nil {
		return m.Whin
	}
	return 0
}

func (m *Meterreadings) GetWhout() uint64 {
	if m != nil {
		return m.Whout
	}
	return 0
}

func (m *Meterreadings) GetMvolt() uint64 {
	if m != nil {
		return m.Mvolt
	}
	return 0
}

func (m *Meterreadings) GetMhertz() uint64 {
	if m != nil {
		return m.Mhertz
	}
	return 0
}

func (m *Meterreadings) GetMpf() uint64 {
	if m != nil {
		return m.Mpf
	}
	return 0
}

func (m *Meterreadings) GetMaxmi() uint64 {
	if m != nil {
		return m.Maxmi
	}
	return 0
}

func init() {
	proto.RegisterType((*Meterreadings)(nil), "electra.meter.Meterreadings")
}

func init() { proto.RegisterFile("meter/meterreadings.proto", fileDescriptor_c4ed5ce77a048220) }

var fileDescriptor_c4ed5ce77a048220 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x9a, 0x86, 0xc6, 0x52, 0x25, 0x64, 0x01, 0x3a, 0x10, 0xb2, 0x2a, 0xa6, 0x4e,
	0xcd, 0xc0, 0x1b, 0x20, 0x16, 0x06, 0x96, 0x8e, 0x6c, 0xa6, 0x3d, 0x88, 0xa5, 0xba, 0xb1, 0x9c,
	0xa3, 0x2d, 0x3c, 0x05, 0x8f, 0xc5, 0xd8, 0x91, 0x0d, 0x94, 0xbc, 0x08, 0xca, 0x39, 0x82, 0x2e,
	0xa7, 0xfb, 0xbe, 0xfb, 0x6f, 0xf9, 0xe5, 0x85, 0x43, 0xc2, 0x50, 0xf0, 0x0c, 0x68, 0x96, 0x76,
	0xfd, 0x52, 0xcf, 0x7c, 0xa8, 0xa8, 0x52, 0x63, 0x5c, 0xe1, 0x82, 0x82, 0x99, 0xf1, 0xf1, 0xfa,
	0x5b, 0xc8, 0xf1, 0xc3, 0x61, 0x4c, 0x5d, 0xca, 0xd1, 0x12, 0x37, 0x76, 0x81, 0xf7, 0x77, 0x20,
	0x26, 0x62, 0x9a, 0xcf, 0xff, 0x58, 0x5d, 0xc9, 0x9c, 0xac, 0xc3, 0x9a, 0x8c, 0xf3, 0x70, 0x34,
	0x11, 0xd3, 0x74, 0xfe, 0x2f, 0xd4, 0xa9, 0x1c, 0xfa, 0xd2, 0xd4, 0x08, 0x03, 0xbe, 0x44, 0x50,
	0x4a, 0xa6, 0xdb, 0xd2, 0xae, 0x21, 0x65, 0xc9, 0x7b, 0x97, 0xdc, 0x96, 0xd5, 0x2b, 0xc1, 0x30,
	0x26, 0x19, 0x3a, 0xeb, 0x36, 0xd5, 0x8a, 0x20, 0x8b, 0x96, 0x41, 0x9d, 0xcb, 0xcc, 0x95, 0x18,
	0xe8, 0x1d, 0x8e, 0x59, 0xf7, 0xa4, 0x4e, 0xe4, 0xc0, 0xf9, 0x67, 0x18, 0xb1, 0xec, 0x56, 0xfe,
	0x37, 0x3b, 0x67, 0x21, 0xef, 0xff, 0x3b, 0xb8, 0x2d, 0x3e, 0x1b, 0x2d, 0xf6, 0x8d, 0x16, 0x3f,
	0x8d, 0x16, 0x1f, 0xad, 0x4e, 0xf6, 0xad, 0x4e, 0xbe, 0x5a, 0x9d, 0x3c, 0x9e, 0xf5, 0x55, 0x14,
	0xbb, 0xd8, 0x54, 0x41, 0x6f, 0x1e, 0xeb, 0xa7, 0x8c, 0x8b, 0xba, 0xf9, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x84, 0x50, 0x21, 0x6e, 0x45, 0x01, 0x00, 0x00,
}

func (m *Meterreadings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Meterreadings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Meterreadings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Maxmi != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Maxmi))
		i--
		dAtA[i] = 0x48
	}
	if m.Mpf != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Mpf))
		i--
		dAtA[i] = 0x40
	}
	if m.Mhertz != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Mhertz))
		i--
		dAtA[i] = 0x38
	}
	if m.Mvolt != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Mvolt))
		i--
		dAtA[i] = 0x30
	}
	if m.Whout != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Whout))
		i--
		dAtA[i] = 0x28
	}
	if m.Whin != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Whin))
		i--
		dAtA[i] = 0x20
	}
	if m.Phase != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x18
	}
	if m.Timestamp != 0 {
		i = encodeVarintMeterreadings(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DeviceID) > 0 {
		i -= len(m.DeviceID)
		copy(dAtA[i:], m.DeviceID)
		i = encodeVarintMeterreadings(dAtA, i, uint64(len(m.DeviceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMeterreadings(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeterreadings(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Meterreadings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovMeterreadings(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovMeterreadings(uint64(m.Timestamp))
	}
	if m.Phase != 0 {
		n += 1 + sovMeterreadings(uint64(m.Phase))
	}
	if m.Whin != 0 {
		n += 1 + sovMeterreadings(uint64(m.Whin))
	}
	if m.Whout != 0 {
		n += 1 + sovMeterreadings(uint64(m.Whout))
	}
	if m.Mvolt != 0 {
		n += 1 + sovMeterreadings(uint64(m.Mvolt))
	}
	if m.Mhertz != 0 {
		n += 1 + sovMeterreadings(uint64(m.Mhertz))
	}
	if m.Mpf != 0 {
		n += 1 + sovMeterreadings(uint64(m.Mpf))
	}
	if m.Maxmi != 0 {
		n += 1 + sovMeterreadings(uint64(m.Maxmi))
	}
	return n
}

func sovMeterreadings(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeterreadings(x uint64) (n int) {
	return sovMeterreadings(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Meterreadings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeterreadings
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
			return fmt.Errorf("proto: Meterreadings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Meterreadings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
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
				return ErrInvalidLengthMeterreadings
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeterreadings
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			m.Phase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Phase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Whin", wireType)
			}
			m.Whin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Whin |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Whout", wireType)
			}
			m.Whout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Whout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mvolt", wireType)
			}
			m.Mvolt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mvolt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mhertz", wireType)
			}
			m.Mhertz = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mhertz |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mpf", wireType)
			}
			m.Mpf = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mpf |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maxmi", wireType)
			}
			m.Maxmi = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeterreadings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Maxmi |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeterreadings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMeterreadings
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
func skipMeterreadings(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeterreadings
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
					return 0, ErrIntOverflowMeterreadings
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
					return 0, ErrIntOverflowMeterreadings
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
				return 0, ErrInvalidLengthMeterreadings
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMeterreadings
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMeterreadings
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMeterreadings        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeterreadings          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMeterreadings = fmt.Errorf("proto: unexpected end of group")
)
