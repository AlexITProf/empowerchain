// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/plasticcredit/credit_batch.proto

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

type CreditBatch struct {
	BatchDenom     string        `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	ProjectId      uint64        `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	TotalAmount    *CreditAmount `protobuf:"bytes,3,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	CreditData     []*ProvenData `protobuf:"bytes,4,rep,name=credit_data,json=creditData,proto3" json:"credit_data,omitempty"`
	AdditionalData []*ProvenData `protobuf:"bytes,5,rep,name=additional_data,json=additionalData,proto3" json:"additional_data,omitempty"`
}

func (m *CreditBatch) Reset()         { *m = CreditBatch{} }
func (m *CreditBatch) String() string { return proto.CompactTextString(m) }
func (*CreditBatch) ProtoMessage()    {}
func (*CreditBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_013616f774bc0fd4, []int{0}
}
func (m *CreditBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreditBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreditBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreditBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditBatch.Merge(m, src)
}
func (m *CreditBatch) XXX_Size() int {
	return m.Size()
}
func (m *CreditBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditBatch.DiscardUnknown(m)
}

var xxx_messageInfo_CreditBatch proto.InternalMessageInfo

func (m *CreditBatch) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *CreditBatch) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *CreditBatch) GetTotalAmount() *CreditAmount {
	if m != nil {
		return m.TotalAmount
	}
	return nil
}

func (m *CreditBatch) GetCreditData() []*ProvenData {
	if m != nil {
		return m.CreditData
	}
	return nil
}

func (m *CreditBatch) GetAdditionalData() []*ProvenData {
	if m != nil {
		return m.AdditionalData
	}
	return nil
}

func init() {
	proto.RegisterType((*CreditBatch)(nil), "empowerchain.empowerchain.plasticcredit.CreditBatch")
}

func init() {
	proto.RegisterFile("empowerchain/plasticcredit/credit_batch.proto", fileDescriptor_013616f774bc0fd4)
}

var fileDescriptor_013616f774bc0fd4 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x6d, 0x0a, 0x4b, 0x45, 0x21, 0xa7, 0x32, 0x30, 0x16, 0x2f, 0xf6, 0xa0, 0x19,
	0x6c, 0x78, 0xf1, 0xe6, 0xdc, 0xc5, 0x9b, 0x94, 0x1d, 0x44, 0x84, 0x92, 0x25, 0xc1, 0x45, 0xd6,
	0x26, 0x74, 0xef, 0xfc, 0xf3, 0x2d, 0xfc, 0x58, 0x1e, 0x77, 0x14, 0xbc, 0xc8, 0xf6, 0x45, 0xa4,
	0x49, 0x41, 0xeb, 0x41, 0x87, 0xa7, 0xf6, 0x79, 0xf2, 0xbe, 0x3f, 0x9e, 0x97, 0x07, 0x9f, 0xa8,
	0xcc, 0x9a, 0x47, 0x55, 0x88, 0x29, 0xd7, 0x79, 0xcf, 0xce, 0xf8, 0x1c, 0xb4, 0x10, 0x85, 0x92,
	0x1a, 0x7a, 0xfe, 0x93, 0x4e, 0x38, 0x88, 0x29, 0xb3, 0x85, 0x01, 0x43, 0x8e, 0xbe, 0x8f, 0xb3,
	0x9a, 0xa8, 0xed, 0x76, 0xd9, 0xdf, 0x5c, 0x9e, 0x99, 0x45, 0x0e, 0x1e, 0xdc, 0x3d, 0xfe, 0x65,
	0xde, 0x16, 0xe6, 0x41, 0xe5, 0xa9, 0xe4, 0xc0, 0xfd, 0xf4, 0xe1, 0x7b, 0x13, 0x07, 0x17, 0xee,
	0x71, 0x58, 0x86, 0x23, 0x07, 0x38, 0x70, 0x29, 0x53, 0xa9, 0x72, 0x93, 0x85, 0x28, 0x42, 0x71,
	0x27, 0xc1, 0xce, 0x1a, 0x95, 0x0e, 0xd9, 0xc7, 0xd8, 0x16, 0xe6, 0x5e, 0x09, 0x48, 0xb5, 0x0c,
	0x9b, 0x11, 0x8a, 0xdb, 0x49, 0xa7, 0x72, 0x2e, 0x25, 0xb9, 0xc6, 0x3b, 0x60, 0x80, 0xcf, 0xaa,
	0x4c, 0x61, 0x2b, 0x42, 0x71, 0xd0, 0x3f, 0x65, 0x1b, 0x5e, 0xcb, 0x7c, 0x96, 0x73, 0xb7, 0x9c,
	0x04, 0x0e, 0xe5, 0x05, 0x19, 0xe3, 0xa0, 0x3a, 0xb7, 0x8c, 0x1f, 0xb6, 0xa3, 0x56, 0x1c, 0xf4,
	0x07, 0x1b, 0x83, 0xaf, 0xdc, 0xe9, 0x23, 0x0e, 0x3c, 0xc1, 0xde, 0x2a, 0xff, 0xc9, 0x2d, 0xde,
	0xe3, 0x52, 0x6a, 0xd0, 0x26, 0xe7, 0x33, 0x4f, 0xde, 0xfa, 0x3f, 0x79, 0xf7, 0x8b, 0x55, 0xea,
	0xe1, 0xf8, 0x75, 0x45, 0xd1, 0x72, 0x45, 0xd1, 0xc7, 0x8a, 0xa2, 0x97, 0x35, 0x6d, 0x2c, 0xd7,
	0xb4, 0xf1, 0xb6, 0xa6, 0x8d, 0x9b, 0xb3, 0x3b, 0x0d, 0xd3, 0xc5, 0x84, 0x09, 0x93, 0xf5, 0x6a,
	0x85, 0xd5, 0xc4, 0xd3, 0x8f, 0xfe, 0xe0, 0xd9, 0xaa, 0xf9, 0x64, 0xdb, 0x55, 0x37, 0xf8, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x81, 0x4b, 0x00, 0x16, 0x72, 0x02, 0x00, 0x00,
}

func (m *CreditBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreditBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreditBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdditionalData) > 0 {
		for iNdEx := len(m.AdditionalData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdditionalData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCreditBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.CreditData) > 0 {
		for iNdEx := len(m.CreditData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCreditBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.TotalAmount != nil {
		{
			size, err := m.TotalAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCreditBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ProjectId != 0 {
		i = encodeVarintCreditBatch(dAtA, i, uint64(m.ProjectId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintCreditBatch(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCreditBatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovCreditBatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreditBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovCreditBatch(uint64(l))
	}
	if m.ProjectId != 0 {
		n += 1 + sovCreditBatch(uint64(m.ProjectId))
	}
	if m.TotalAmount != nil {
		l = m.TotalAmount.Size()
		n += 1 + l + sovCreditBatch(uint64(l))
	}
	if len(m.CreditData) > 0 {
		for _, e := range m.CreditData {
			l = e.Size()
			n += 1 + l + sovCreditBatch(uint64(l))
		}
	}
	if len(m.AdditionalData) > 0 {
		for _, e := range m.AdditionalData {
			l = e.Size()
			n += 1 + l + sovCreditBatch(uint64(l))
		}
	}
	return n
}

func sovCreditBatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCreditBatch(x uint64) (n int) {
	return sovCreditBatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreditBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCreditBatch
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
			return fmt.Errorf("proto: CreditBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreditBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			m.ProjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TotalAmount == nil {
				m.TotalAmount = &CreditAmount{}
			}
			if err := m.TotalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditData = append(m.CreditData, &ProvenData{})
			if err := m.CreditData[len(m.CreditData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreditBatch
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
				return ErrInvalidLengthCreditBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCreditBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalData = append(m.AdditionalData, &ProvenData{})
			if err := m.AdditionalData[len(m.AdditionalData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCreditBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCreditBatch
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
func skipCreditBatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCreditBatch
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
					return 0, ErrIntOverflowCreditBatch
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
					return 0, ErrIntOverflowCreditBatch
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
				return 0, ErrInvalidLengthCreditBatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCreditBatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCreditBatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCreditBatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCreditBatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCreditBatch = fmt.Errorf("proto: unexpected end of group")
)
