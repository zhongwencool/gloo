// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/filter/http/gzip/v2/gzip.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}

var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}

func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1affeaf387577671, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}

var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}

func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1affeaf387577671, []int{0, 0, 0}
}

// [#next-free-field: 10]
type Gzip struct {
	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. The default value is 5.
	MemoryLevel *types.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	// Minimum response length, in bytes, which will trigger compression. The default value is 30.
	ContentLength *types.UInt32Value `protobuf:"bytes,2,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	// A value used for selecting the zlib compression level. This setting will affect speed and
	// amount of compression applied to the content. "BEST" provides higher compression at the cost of
	// higher latency, "SPEED" provides lower compression with minimum impact on response time.
	// "DEFAULT" provides an optimal result between speed and compression. This field will be set to
	// "DEFAULT" if not specified.
	CompressionLevel Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	// A value used for selecting the zlib compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT" will be the best choice, though
	// there are situations which changing this parameter might produce better results. For example,
	// run-length encoding (RLE) is typically used when the content is known for having sequences
	// which same data occurs many consecutive times. For more information about each strategy, please
	// refer to zlib manual.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Set of strings that allows specifying which mime-types yield compression; e.g.,
	// application/json, text/html, etc. When this field is not defined, compression will be applied
	// to the following mime-types: "application/javascript", "application/json",
	// "application/xhtml+xml", "image/svg+xml", "text/css", "text/html", "text/plain", "text/xml".
	ContentType []string `protobuf:"bytes,6,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, disables compression when the response contains an etag header. When it is false, the
	// filter will preserve weak etags and remove the ones that require strong validation.
	DisableOnEtagHeader bool `protobuf:"varint,7,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	// If true, removes accept-encoding from the request headers before dispatching it to the upstream
	// so that responses do not get compressed before reaching the filter.
	RemoveAcceptEncodingHeader bool `protobuf:"varint,8,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits           *types.UInt32Value `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_1affeaf387577671, []int{0}
}
func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (m *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(m, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetMemoryLevel() *types.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetContentLength() *types.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Gzip) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Gzip) GetRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.RemoveAcceptEncodingHeader
	}
	return false
}

func (m *Gzip) GetWindowBits() *types.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

type Gzip_CompressionLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gzip_CompressionLevel) Reset()         { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()    {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1affeaf387577671, []int{0, 0}
}
func (m *Gzip_CompressionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip_CompressionLevel.Unmarshal(m, b)
}
func (m *Gzip_CompressionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip_CompressionLevel.Marshal(b, m, deterministic)
}
func (m *Gzip_CompressionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip_CompressionLevel.Merge(m, src)
}
func (m *Gzip_CompressionLevel) XXX_Size() int {
	return xxx_messageInfo_Gzip_CompressionLevel.Size(m)
}
func (m *Gzip_CompressionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip_CompressionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip_CompressionLevel proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
	proto.RegisterType((*Gzip)(nil), "envoy.config.filter.http.gzip.v2.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.config.filter.http.gzip.v2.Gzip.CompressionLevel")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/filter/http/gzip/v2/gzip.proto", fileDescriptor_1affeaf387577671)
}

var fileDescriptor_1affeaf387577671 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x49, 0x9b, 0x75, 0xad, 0x3b, 0x46, 0xf0, 0x26, 0x88, 0x2a, 0x98, 0xaa, 0x9d, 0xa2,
	0x49, 0xb3, 0xa5, 0xee, 0x86, 0x76, 0x69, 0x59, 0xca, 0x86, 0xca, 0x40, 0x59, 0x07, 0x12, 0x07,
	0x22, 0x37, 0x7d, 0xe7, 0x1a, 0x52, 0xdb, 0x4a, 0xdc, 0x6e, 0xdd, 0x91, 0x2f, 0x80, 0xc4, 0xa7,
	0xe0, 0x23, 0x20, 0x4e, 0x7c, 0x16, 0x6e, 0x7c, 0x00, 0xee, 0x28, 0x7f, 0x06, 0x6c, 0x20, 0xb1,
	0x9d, 0x62, 0xfb, 0x79, 0x7e, 0xcf, 0xeb, 0xbc, 0xb6, 0x8c, 0xde, 0x70, 0x61, 0x26, 0xb3, 0x11,
	0x89, 0xd4, 0x94, 0xa6, 0x2a, 0x56, 0xdb, 0x42, 0x51, 0x1e, 0x2b, 0x45, 0x75, 0xa2, 0xde, 0x42,
	0x64, 0xd2, 0x62, 0xc6, 0xb4, 0xa0, 0x70, 0x66, 0x20, 0x91, 0x2c, 0xa6, 0x20, 0xe7, 0x6a, 0x41,
	0x23, 0x25, 0x4f, 0x04, 0xa7, 0x27, 0x22, 0x36, 0x90, 0xd0, 0x89, 0x31, 0x9a, 0xf2, 0x73, 0xa1,
	0xe9, 0xbc, 0x93, 0x7f, 0x89, 0x4e, 0x94, 0x51, 0xb8, 0x9d, 0x9b, 0x49, 0x61, 0x26, 0x85, 0x99,
	0x64, 0x66, 0x92, 0x9b, 0xe6, 0x9d, 0xd6, 0x3a, 0x57, 0x5c, 0xe5, 0x66, 0x9a, 0x8d, 0x0a, 0xae,
	0x85, 0xe1, 0xcc, 0x14, 0x8b, 0x70, 0x66, 0xca, 0xb5, 0x0d, 0xae, 0x14, 0x8f, 0x81, 0xe6, 0xb3,
	0xd1, 0xec, 0x84, 0x9e, 0x26, 0x4c, 0x6b, 0x48, 0xd2, 0x52, 0xbf, 0x3f, 0x67, 0xb1, 0x18, 0x33,
	0x03, 0xf4, 0x62, 0x50, 0x08, 0x9b, 0x1f, 0x6a, 0xc8, 0x7e, 0x72, 0x2e, 0x34, 0x1e, 0xa0, 0x95,
	0x29, 0x4c, 0x55, 0xb2, 0x08, 0x63, 0x98, 0x43, 0xec, 0x5a, 0x6d, 0xcb, 0x6b, 0x76, 0x1e, 0x90,
	0x22, 0x98, 0x5c, 0x04, 0x93, 0xe3, 0x03, 0x69, 0x76, 0x3a, 0x2f, 0x59, 0x3c, 0x83, 0x5e, 0xf3,
	0xcb, 0xf7, 0xaf, 0xd5, 0xda, 0x96, 0xed, 0x36, 0x3c, 0x2b, 0x68, 0x16, 0xf8, 0x20, 0xa3, 0xf1,
	0x21, 0x5a, 0x8d, 0x94, 0x34, 0x20, 0x4d, 0x18, 0x83, 0xe4, 0x66, 0xe2, 0x56, 0xae, 0x91, 0xd7,
	0xc8, 0xf2, 0xec, 0xad, 0x8a, 0xb7, 0x11, 0xdc, 0x2e, 0xf1, 0x41, 0x4e, 0xe3, 0x19, 0xba, 0x1b,
	0xa9, 0xa9, 0x4e, 0x20, 0x4d, 0x85, 0x92, 0xe5, 0x16, 0xab, 0x6d, 0xcb, 0x5b, 0xed, 0xec, 0x92,
	0xff, 0xf5, 0x91, 0x64, 0x3f, 0x48, 0x1e, 0xff, 0xe6, 0xf3, 0x3d, 0x12, 0x5f, 0xce, 0xa6, 0x3d,
	0x94, 0x95, 0x5c, 0x7a, 0x6f, 0x55, 0x1c, 0x2b, 0x70, 0xa2, 0x2b, 0x16, 0xbc, 0x40, 0xeb, 0x7f,
	0x96, 0x4d, 0x4d, 0xc2, 0x0c, 0xf0, 0x85, 0x6b, 0xe7, 0x95, 0x1f, 0xdd, 0xbc, 0xf2, 0x51, 0x99,
	0x70, 0xa9, 0xee, 0x5a, 0xf4, 0xb7, 0x01, 0x6f, 0xa3, 0x95, 0x8b, 0x0e, 0x9a, 0x85, 0x06, 0xb7,
	0xd6, 0xae, 0x7a, 0x8d, 0x12, 0xfb, 0x68, 0x55, 0x9c, 0x4e, 0xd0, 0x2c, 0xf5, 0xe1, 0x42, 0x03,
	0xde, 0x41, 0xf7, 0xc6, 0x22, 0x65, 0xa3, 0x18, 0x42, 0x25, 0x43, 0x30, 0x8c, 0x87, 0x13, 0x60,
	0x63, 0x48, 0xdc, 0xe5, 0xb6, 0xe5, 0xd5, 0x83, 0xb5, 0x52, 0x7d, 0x2e, 0x7d, 0xc3, 0xf8, 0x7e,
	0x2e, 0xe1, 0x2e, 0x7a, 0x98, 0xc0, 0x54, 0xcd, 0x21, 0x64, 0x51, 0x04, 0xda, 0x84, 0x20, 0x23,
	0x35, 0x16, 0xf2, 0x17, 0x5b, 0xcf, 0xd9, 0x56, 0x61, 0xea, 0xe6, 0x1e, 0xbf, 0xb4, 0x94, 0x11,
	0x4f, 0x51, 0xf3, 0x54, 0xc8, 0xb1, 0x3a, 0x0d, 0x47, 0xc2, 0xa4, 0x6e, 0xe3, 0x26, 0xb7, 0xe6,
	0x8e, 0xd7, 0x08, 0x50, 0x41, 0xf7, 0x84, 0x49, 0x5b, 0xbb, 0xc8, 0xb9, 0x7a, 0x48, 0x9b, 0x1e,
	0xb2, 0xb3, 0x73, 0xc2, 0x4d, 0xb4, 0xbc, 0xe7, 0xf7, 0xbb, 0xc7, 0x83, 0xa1, 0x73, 0x0b, 0xd7,
	0x91, 0xdd, 0xf3, 0x8f, 0x86, 0x8e, 0x85, 0x1b, 0x68, 0xe9, 0xe8, 0x85, 0xef, 0xef, 0x39, 0x95,
	0xcd, 0x3e, 0x5a, 0xfb, 0x47, 0xa3, 0x2f, 0x83, 0x2b, 0xa8, 0xde, 0x3f, 0x18, 0x0c, 0xfd, 0xc0,
	0xdf, 0x73, 0xac, 0x4c, 0xda, 0x3f, 0xee, 0xf7, 0x9f, 0x75, 0x0f, 0x9d, 0x0a, 0x5e, 0x46, 0xd5,
	0x60, 0xe0, 0x3b, 0xd5, 0x1e, 0xfb, 0xfc, 0xc3, 0xb6, 0x3e, 0x7d, 0xdb, 0xb0, 0x5e, 0xbf, 0xba,
	0xde, 0x03, 0xa0, 0xdf, 0xf1, 0x9b, 0x3d, 0x02, 0xa3, 0x5a, 0xde, 0x97, 0x9d, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0xfe, 0x5a, 0x23, 0x62, 0x04, 0x00, 0x00,
}

func (this *Gzip) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gzip)
	if !ok {
		that2, ok := that.(Gzip)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MemoryLevel.Equal(that1.MemoryLevel) {
		return false
	}
	if !this.ContentLength.Equal(that1.ContentLength) {
		return false
	}
	if this.CompressionLevel != that1.CompressionLevel {
		return false
	}
	if this.CompressionStrategy != that1.CompressionStrategy {
		return false
	}
	if len(this.ContentType) != len(that1.ContentType) {
		return false
	}
	for i := range this.ContentType {
		if this.ContentType[i] != that1.ContentType[i] {
			return false
		}
	}
	if this.DisableOnEtagHeader != that1.DisableOnEtagHeader {
		return false
	}
	if this.RemoveAcceptEncodingHeader != that1.RemoveAcceptEncodingHeader {
		return false
	}
	if !this.WindowBits.Equal(that1.WindowBits) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Gzip_CompressionLevel) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gzip_CompressionLevel)
	if !ok {
		that2, ok := that.(Gzip_CompressionLevel)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
