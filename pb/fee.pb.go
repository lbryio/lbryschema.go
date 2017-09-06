// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/fee.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Fee_Version int32

const (
	Fee_UNKNOWN_VERSION Fee_Version = 0
	Fee__0_0_1          Fee_Version = 1
)

var Fee_Version_name = map[int32]string{
	0: "UNKNOWN_VERSION",
	1: "_0_0_1",
}
var Fee_Version_value = map[string]int32{
	"UNKNOWN_VERSION": 0,
	"_0_0_1":          1,
}

func (x Fee_Version) Enum() *Fee_Version {
	p := new(Fee_Version)
	*p = x
	return p
}
func (x Fee_Version) String() string {
	return proto.EnumName(Fee_Version_name, int32(x))
}
func (x *Fee_Version) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Fee_Version_value, data, "Fee_Version")
	if err != nil {
		return err
	}
	*x = Fee_Version(value)
	return nil
}
func (Fee_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Fee_Currency int32

const (
	Fee_UNKNOWN_CURRENCY Fee_Currency = 0
	Fee_LBC              Fee_Currency = 1
	Fee_BTC              Fee_Currency = 2
	Fee_USD              Fee_Currency = 3
)

var Fee_Currency_name = map[int32]string{
	0: "UNKNOWN_CURRENCY",
	1: "LBC",
	2: "BTC",
	3: "USD",
}
var Fee_Currency_value = map[string]int32{
	"UNKNOWN_CURRENCY": 0,
	"LBC":              1,
	"BTC":              2,
	"USD":              3,
}

func (x Fee_Currency) Enum() *Fee_Currency {
	p := new(Fee_Currency)
	*p = x
	return p
}
func (x Fee_Currency) String() string {
	return proto.EnumName(Fee_Currency_name, int32(x))
}
func (x *Fee_Currency) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Fee_Currency_value, data, "Fee_Currency")
	if err != nil {
		return err
	}
	*x = Fee_Currency(value)
	return nil
}
func (Fee_Currency) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

type Fee struct {
	Version          *Fee_Version  `protobuf:"varint,1,req,name=version,enum=pb.Fee_Version" json:"version,omitempty"`
	Currency         *Fee_Currency `protobuf:"varint,2,req,name=currency,enum=pb.Fee_Currency" json:"currency,omitempty"`
	Address          []byte        `protobuf:"bytes,3,req,name=address" json:"address,omitempty"`
	Amount           *float32      `protobuf:"fixed32,4,req,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Fee) Reset()                    { *m = Fee{} }
func (m *Fee) String() string            { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()               {}
func (*Fee) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Fee) GetVersion() Fee_Version {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Fee_UNKNOWN_VERSION
}

func (m *Fee) GetCurrency() Fee_Currency {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return Fee_UNKNOWN_CURRENCY
}

func (m *Fee) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Fee) GetAmount() float32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Fee)(nil), "pb.Fee")
	proto.RegisterEnum("pb.Fee_Version", Fee_Version_name, Fee_Version_value)
	proto.RegisterEnum("pb.Fee_Currency", Fee_Currency_name, Fee_Currency_value)
}

func init() { proto.RegisterFile("pb/fee.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8d, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd7, 0x54, 0xd6, 0xf1, 0x18, 0xee, 0xf1, 0x14, 0xc9, 0xb1, 0xf4, 0x54, 0x45, 0xea,
	0xf4, 0xea, 0x6d, 0x71, 0x03, 0x51, 0x32, 0xc8, 0xec, 0xc4, 0x53, 0x59, 0xb7, 0x27, 0x78, 0xb0,
	0x09, 0xe9, 0x26, 0xf8, 0xa7, 0x7b, 0x93, 0x6a, 0xb3, 0xdb, 0xf7, 0xc7, 0x07, 0x3e, 0x30, 0x76,
	0xf5, 0xcd, 0x3b, 0x73, 0xe1, 0xbc, 0xdd, 0x5b, 0x12, 0xae, 0xce, 0x7e, 0x22, 0x88, 0x17, 0xcc,
	0x74, 0x09, 0xc9, 0x17, 0xfb, 0xf6, 0xc3, 0x36, 0x32, 0x4a, 0x45, 0x7e, 0x7a, 0x37, 0x29, 0x5c,
	0x5d, 0x2c, 0x98, 0x8b, 0xf5, 0xff, 0x6c, 0xc2, 0x4f, 0xd7, 0x30, 0xda, 0x1e, 0xbc, 0xe7, 0x66,
	0xfb, 0x2d, 0xc5, 0x1f, 0x8b, 0x81, 0x55, 0xfd, 0x6e, 0x8e, 0x04, 0x49, 0x48, 0x36, 0xbb, 0x9d,
	0xe7, 0xb6, 0x95, 0x71, 0x2a, 0xf2, 0xb1, 0x09, 0x95, 0x2e, 0x60, 0xb8, 0xf9, 0xb4, 0x87, 0x66,
	0x2f, 0x4f, 0x52, 0x91, 0x0b, 0xd3, 0xb7, 0xec, 0x0a, 0x92, 0xde, 0x49, 0x67, 0x30, 0x29, 0xf5,
	0x93, 0x5e, 0xbe, 0xea, 0x6a, 0x3d, 0x37, 0xab, 0xc7, 0xa5, 0xc6, 0x01, 0x01, 0x0c, 0xab, 0x69,
	0x35, 0xad, 0x6e, 0x31, 0xca, 0xee, 0x61, 0x14, 0x9c, 0x74, 0x0e, 0x18, 0x60, 0x55, 0x1a, 0x33,
	0xd7, 0xea, 0x0d, 0x07, 0x94, 0x40, 0xfc, 0x3c, 0x53, 0x18, 0x75, 0x61, 0xf6, 0xa2, 0x50, 0x74,
	0xa1, 0x5c, 0x3d, 0x60, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x47, 0x34, 0x9b, 0x9c, 0x0e, 0x01,
	0x00, 0x00,
}