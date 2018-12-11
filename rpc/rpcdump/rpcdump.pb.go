// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpcdump.proto

package rpcdump // import "github.com/hyperledger/burrow/rpc/rpcdump"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import dump "github.com/hyperledger/burrow/dump"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetDumpParam struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDumpParam) Reset()         { *m = GetDumpParam{} }
func (m *GetDumpParam) String() string { return proto.CompactTextString(m) }
func (*GetDumpParam) ProtoMessage()    {}
func (*GetDumpParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpcdump_a8b6aa4e2ca1c22a, []int{0}
}
func (m *GetDumpParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDumpParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDumpParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetDumpParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDumpParam.Merge(dst, src)
}
func (m *GetDumpParam) XXX_Size() int {
	return m.Size()
}
func (m *GetDumpParam) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDumpParam.DiscardUnknown(m)
}

var xxx_messageInfo_GetDumpParam proto.InternalMessageInfo

func (*GetDumpParam) XXX_MessageName() string {
	return "rpcdump.GetDumpParam"
}
func init() {
	proto.RegisterType((*GetDumpParam)(nil), "rpcdump.GetDumpParam")
	golang_proto.RegisterType((*GetDumpParam)(nil), "rpcdump.GetDumpParam")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DumpClient is the client API for Dump service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DumpClient interface {
	GetDump(ctx context.Context, in *GetDumpParam, opts ...grpc.CallOption) (Dump_GetDumpClient, error)
}

type dumpClient struct {
	cc *grpc.ClientConn
}

func NewDumpClient(cc *grpc.ClientConn) DumpClient {
	return &dumpClient{cc}
}

func (c *dumpClient) GetDump(ctx context.Context, in *GetDumpParam, opts ...grpc.CallOption) (Dump_GetDumpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dump_serviceDesc.Streams[0], "/rpcdump.Dump/GetDump", opts...)
	if err != nil {
		return nil, err
	}
	x := &dumpGetDumpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dump_GetDumpClient interface {
	Recv() (*dump.Dump, error)
	grpc.ClientStream
}

type dumpGetDumpClient struct {
	grpc.ClientStream
}

func (x *dumpGetDumpClient) Recv() (*dump.Dump, error) {
	m := new(dump.Dump)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DumpServer is the server API for Dump service.
type DumpServer interface {
	GetDump(*GetDumpParam, Dump_GetDumpServer) error
}

func RegisterDumpServer(s *grpc.Server, srv DumpServer) {
	s.RegisterService(&_Dump_serviceDesc, srv)
}

func _Dump_GetDump_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDumpParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DumpServer).GetDump(m, &dumpGetDumpServer{stream})
}

type Dump_GetDumpServer interface {
	Send(*dump.Dump) error
	grpc.ServerStream
}

type dumpGetDumpServer struct {
	grpc.ServerStream
}

func (x *dumpGetDumpServer) Send(m *dump.Dump) error {
	return x.ServerStream.SendMsg(m)
}

var _Dump_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcdump.Dump",
	HandlerType: (*DumpServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDump",
			Handler:       _Dump_GetDump_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpcdump.proto",
}

func (m *GetDumpParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDumpParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRpcdump(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetDumpParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRpcdump(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRpcdump(x uint64) (n int) {
	return sovRpcdump(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetDumpParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpcdump
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDumpParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDumpParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpcdump(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpcdump
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
func skipRpcdump(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpcdump
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
					return 0, ErrIntOverflowRpcdump
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpcdump
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRpcdump
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpcdump
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRpcdump(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRpcdump = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpcdump   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpcdump.proto", fileDescriptor_rpcdump_a8b6aa4e2ca1c22a) }
func init() { golang_proto.RegisterFile("rpcdump.proto", fileDescriptor_rpcdump_a8b6aa4e2ca1c22a) }

var fileDescriptor_rpcdump_a8b6aa4e2ca1c22a = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2a, 0x48, 0x4e,
	0x29, 0xcd, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x74, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1,
	0xf2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xf4, 0x49, 0x71, 0x21, 0xcc, 0x50,
	0xe2, 0xe3, 0xe2, 0x71, 0x4f, 0x2d, 0x71, 0x29, 0xcd, 0x2d, 0x08, 0x48, 0x2c, 0x4a, 0xcc, 0x35,
	0x32, 0xe3, 0x62, 0x01, 0x71, 0x84, 0xf4, 0xb8, 0xd8, 0xa1, 0xe2, 0x42, 0xa2, 0x7a, 0x30, 0x6b,
	0x91, 0x55, 0x4a, 0x71, 0xe9, 0x81, 0xc5, 0x40, 0x02, 0x06, 0x8c, 0x4e, 0xf6, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x81, 0xc7, 0x72, 0x8c, 0x27, 0x1e,
	0xcb, 0x31, 0x46, 0x69, 0x22, 0x39, 0x2c, 0xa3, 0xb2, 0x20, 0xb5, 0x28, 0x27, 0x35, 0x25, 0x3d,
	0xb5, 0x48, 0x3f, 0xa9, 0xb4, 0xa8, 0x28, 0xbf, 0x5c, 0xbf, 0xa8, 0x20, 0x59, 0x1f, 0x6a, 0x76,
	0x12, 0x1b, 0xd8, 0x3d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x2c, 0x52, 0x86, 0xe4,
	0x00, 0x00, 0x00,
}
