// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation_ee/transformation.proto

package transformation_ee

import (
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	route "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/route"
	_type "github.com/solo-io/solo-kit/pkg/api/external/envoy/type"
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

type FilterTransformations struct {
	// Specifies transformations based on the route matches. The first matched transformation will be
	// applied. If there are overlapped match conditions, please put the most specific match first.
	Transformations      []*TransformationRule `protobuf:"bytes,1,rep,name=transformations,proto3" json:"transformations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FilterTransformations) Reset()         { *m = FilterTransformations{} }
func (m *FilterTransformations) String() string { return proto.CompactTextString(m) }
func (*FilterTransformations) ProtoMessage()    {}
func (*FilterTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e36764f5fd991d, []int{0}
}
func (m *FilterTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterTransformations.Unmarshal(m, b)
}
func (m *FilterTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterTransformations.Marshal(b, m, deterministic)
}
func (m *FilterTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterTransformations.Merge(m, src)
}
func (m *FilterTransformations) XXX_Size() int {
	return xxx_messageInfo_FilterTransformations.Size(m)
}
func (m *FilterTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_FilterTransformations proto.InternalMessageInfo

func (m *FilterTransformations) GetTransformations() []*TransformationRule {
	if m != nil {
		return m.Transformations
	}
	return nil
}

type TransformationRule struct {
	// The route matching parameter. Only when the match is satisfied, the "requires" field will
	// apply.
	//
	// For example: following match will match all requests.
	//
	// .. code-block:: yaml
	//
	//    match:
	//      prefix: /
	//
	Match *route.RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// transformation to perform
	RouteTransformations *RouteTransformations `protobuf:"bytes,2,opt,name=route_transformations,json=routeTransformations,proto3" json:"route_transformations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TransformationRule) Reset()         { *m = TransformationRule{} }
func (m *TransformationRule) String() string { return proto.CompactTextString(m) }
func (*TransformationRule) ProtoMessage()    {}
func (*TransformationRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e36764f5fd991d, []int{1}
}
func (m *TransformationRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformationRule.Unmarshal(m, b)
}
func (m *TransformationRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformationRule.Marshal(b, m, deterministic)
}
func (m *TransformationRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformationRule.Merge(m, src)
}
func (m *TransformationRule) XXX_Size() int {
	return xxx_messageInfo_TransformationRule.Size(m)
}
func (m *TransformationRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformationRule.DiscardUnknown(m)
}

var xxx_messageInfo_TransformationRule proto.InternalMessageInfo

func (m *TransformationRule) GetMatch() *route.RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *TransformationRule) GetRouteTransformations() *RouteTransformations {
	if m != nil {
		return m.RouteTransformations
	}
	return nil
}

type RouteTransformations struct {
	RequestTransformation *Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// clear the route cache if the request transformation was applied
	ClearRouteCache        bool            `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	ResponseTransformation *Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *RouteTransformations) Reset()         { *m = RouteTransformations{} }
func (m *RouteTransformations) String() string { return proto.CompactTextString(m) }
func (*RouteTransformations) ProtoMessage()    {}
func (*RouteTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e36764f5fd991d, []int{2}
}
func (m *RouteTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTransformations.Unmarshal(m, b)
}
func (m *RouteTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTransformations.Marshal(b, m, deterministic)
}
func (m *RouteTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTransformations.Merge(m, src)
}
func (m *RouteTransformations) XXX_Size() int {
	return xxx_messageInfo_RouteTransformations.Size(m)
}
func (m *RouteTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTransformations proto.InternalMessageInfo

func (m *RouteTransformations) GetRequestTransformation() *Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *RouteTransformations) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *RouteTransformations) GetResponseTransformation() *Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

type Transformation struct {
	// Template is in the transformed request language domain
	//
	// Types that are valid to be assigned to TransformationType:
	//	*Transformation_DlpTransformation
	TransformationType   isTransformation_TransformationType `protobuf_oneof:"transformation_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e36764f5fd991d, []int{3}
}
func (m *Transformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transformation.Unmarshal(m, b)
}
func (m *Transformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transformation.Marshal(b, m, deterministic)
}
func (m *Transformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transformation.Merge(m, src)
}
func (m *Transformation) XXX_Size() int {
	return xxx_messageInfo_Transformation.Size(m)
}
func (m *Transformation) XXX_DiscardUnknown() {
	xxx_messageInfo_Transformation.DiscardUnknown(m)
}

var xxx_messageInfo_Transformation proto.InternalMessageInfo

type isTransformation_TransformationType interface {
	isTransformation_TransformationType()
}

type Transformation_DlpTransformation struct {
	DlpTransformation *DlpTransformation `protobuf:"bytes,1,opt,name=dlp_transformation,json=dlpTransformation,proto3,oneof" json:"dlp_transformation,omitempty"`
}

func (*Transformation_DlpTransformation) isTransformation_TransformationType() {}

func (m *Transformation) GetTransformationType() isTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (m *Transformation) GetDlpTransformation() *DlpTransformation {
	if x, ok := m.GetTransformationType().(*Transformation_DlpTransformation); ok {
		return x.DlpTransformation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Transformation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Transformation_DlpTransformation)(nil),
	}
}

type DlpTransformation struct {
	// list of actions to apply
	Actions              []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DlpTransformation) Reset()         { *m = DlpTransformation{} }
func (m *DlpTransformation) String() string { return proto.CompactTextString(m) }
func (*DlpTransformation) ProtoMessage()    {}
func (*DlpTransformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e36764f5fd991d, []int{4}
}
func (m *DlpTransformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DlpTransformation.Unmarshal(m, b)
}
func (m *DlpTransformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DlpTransformation.Marshal(b, m, deterministic)
}
func (m *DlpTransformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DlpTransformation.Merge(m, src)
}
func (m *DlpTransformation) XXX_Size() int {
	return xxx_messageInfo_DlpTransformation.Size(m)
}
func (m *DlpTransformation) XXX_DiscardUnknown() {
	xxx_messageInfo_DlpTransformation.DiscardUnknown(m)
}

var xxx_messageInfo_DlpTransformation proto.InternalMessageInfo

func (m *DlpTransformation) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type Action struct {
	// Identifier for this action.
	// Used mostly to help ID specific actions in logs.
	// If left null will default to unknown
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of regexes to apply to the response body to match data which should be masked
	// They will be applied iteratively in the order which they are specified
	Regex []string `protobuf:"bytes,2,rep,name=regex,proto3" json:"regex,omitempty"`
	// If specified, this rule will not actually be applied, but only logged.
	Shadow bool `protobuf:"varint,3,opt,name=shadow,proto3" json:"shadow,omitempty"`
	// The percent of the string which should be masked.
	// If not set, defaults to 75%
	Percent *_type.Percent `protobuf:"bytes,4,opt,name=percent,proto3" json:"percent,omitempty"`
	// The character which should overwrite the masked data
	// If left empty, defaults to "X"
	MaskChar             string   `protobuf:"bytes,5,opt,name=mask_char,json=maskChar,proto3" json:"mask_char,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e36764f5fd991d, []int{5}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action) GetRegex() []string {
	if m != nil {
		return m.Regex
	}
	return nil
}

func (m *Action) GetShadow() bool {
	if m != nil {
		return m.Shadow
	}
	return false
}

func (m *Action) GetPercent() *_type.Percent {
	if m != nil {
		return m.Percent
	}
	return nil
}

func (m *Action) GetMaskChar() string {
	if m != nil {
		return m.MaskChar
	}
	return ""
}

func init() {
	proto.RegisterType((*FilterTransformations)(nil), "envoy.config.filter.http.transformation_ee.v2.FilterTransformations")
	proto.RegisterType((*TransformationRule)(nil), "envoy.config.filter.http.transformation_ee.v2.TransformationRule")
	proto.RegisterType((*RouteTransformations)(nil), "envoy.config.filter.http.transformation_ee.v2.RouteTransformations")
	proto.RegisterType((*Transformation)(nil), "envoy.config.filter.http.transformation_ee.v2.Transformation")
	proto.RegisterType((*DlpTransformation)(nil), "envoy.config.filter.http.transformation_ee.v2.DlpTransformation")
	proto.RegisterType((*Action)(nil), "envoy.config.filter.http.transformation_ee.v2.Action")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation_ee/transformation.proto", fileDescriptor_73e36764f5fd991d)
}

var fileDescriptor_73e36764f5fd991d = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xfd, 0x3c, 0xf9, 0x69, 0xe3, 0x4a, 0xed, 0x57, 0xb7, 0x69, 0x47, 0x45, 0xaa, 0xa2, 0x11,
	0x42, 0x11, 0x52, 0x67, 0xa4, 0x20, 0x56, 0x08, 0xa9, 0x9d, 0x50, 0xc4, 0x06, 0x11, 0x59, 0x6c,
	0x60, 0x13, 0xb9, 0x13, 0x27, 0x33, 0xcd, 0x64, 0xec, 0xda, 0x4e, 0x48, 0xf6, 0xac, 0x58, 0xb2,
	0xe4, 0x29, 0x60, 0x85, 0xc4, 0x8a, 0x47, 0xe0, 0x35, 0x78, 0x0b, 0x64, 0x7b, 0x22, 0x31, 0x93,
	0x2c, 0x18, 0x75, 0x33, 0xb2, 0xef, 0xb9, 0xf7, 0xdc, 0x73, 0x4f, 0xe2, 0x0b, 0x6f, 0x27, 0x89,
	0x8a, 0xe7, 0x37, 0x7e, 0xc4, 0x66, 0x81, 0x64, 0x29, 0xbb, 0x48, 0x58, 0x30, 0x49, 0x19, 0x0b,
	0xb8, 0x60, 0xb7, 0x34, 0x52, 0xd2, 0xde, 0x08, 0x4f, 0x02, 0xba, 0x54, 0x54, 0x64, 0x24, 0x0d,
	0x68, 0xb6, 0x60, 0x2b, 0x73, 0xcd, 0x64, 0xc2, 0x32, 0x19, 0x28, 0x41, 0x32, 0x39, 0x66, 0x62,
	0x46, 0x54, 0xc2, 0xb2, 0x21, 0xa5, 0xa5, 0x88, 0xcf, 0x05, 0x53, 0x0c, 0x5d, 0x98, 0x42, 0x3f,
	0x62, 0xd9, 0x38, 0x99, 0xf8, 0xe3, 0x24, 0x55, 0x54, 0xf8, 0xb1, 0x52, 0xdc, 0xdf, 0x20, 0xf0,
	0x17, 0xbd, 0xb3, 0xd3, 0x05, 0x49, 0x93, 0x11, 0x51, 0x34, 0x58, 0x1f, 0x2c, 0xcf, 0xd9, 0xb9,
	0x15, 0xa0, 0x35, 0x2d, 0x7a, 0x81, 0x60, 0x73, 0x45, 0xed, 0x37, 0xc7, 0x5d, 0x8b, 0xab, 0x15,
	0xa7, 0x01, 0xa7, 0x22, 0xa2, 0x99, 0xb2, 0x88, 0xf7, 0x11, 0xc0, 0xf6, 0x4b, 0xd3, 0xf7, 0x6d,
	0xa1, 0xa3, 0x44, 0x53, 0x78, 0x50, 0x14, 0x21, 0x5d, 0xd0, 0xa9, 0x75, 0xf7, 0x7a, 0x57, 0x7e,
	0x25, 0xd5, 0x7e, 0x91, 0x18, 0xcf, 0x53, 0x8a, 0xcb, 0xcc, 0xde, 0x2f, 0x00, 0xd1, 0x66, 0x1e,
	0xba, 0x84, 0x8d, 0x19, 0x51, 0x51, 0xec, 0x82, 0x0e, 0xe8, 0xee, 0xf5, 0xce, 0xf3, 0xce, 0x84,
	0x27, 0x9a, 0xd8, 0x4e, 0x88, 0xf5, 0xf7, 0xb5, 0xce, 0x0a, 0xe1, 0x8f, 0xdf, 0x3f, 0x6b, 0x8d,
	0x4f, 0xc0, 0xf9, 0x1f, 0x60, 0x5b, 0x88, 0x96, 0xb0, 0x6d, 0xd2, 0x86, 0xe5, 0x59, 0x1c, 0xc3,
	0xd8, 0xaf, 0x38, 0x8b, 0x69, 0x56, 0x72, 0x0a, 0x1f, 0x8b, 0x2d, 0x51, 0xef, 0xab, 0x03, 0x8f,
	0xb7, 0xa5, 0x23, 0x05, 0x4f, 0x04, 0xbd, 0x9b, 0x53, 0xa9, 0x4a, 0xa2, 0xf2, 0x29, 0x9f, 0xdf,
	0xcf, 0xdf, 0x76, 0x4e, 0x5e, 0x0c, 0xa3, 0xc7, 0xf0, 0x30, 0x4a, 0x29, 0x11, 0x43, 0x6b, 0x47,
	0x44, 0xa2, 0x98, 0xba, 0xb5, 0x0e, 0xe8, 0xee, 0xe2, 0x03, 0x03, 0x18, 0xad, 0x7d, 0x1d, 0x46,
	0x0b, 0x78, 0x2a, 0xa8, 0xe4, 0x2c, 0x93, 0x65, 0xdf, 0x72, 0xdb, 0xee, 0x29, 0xf1, 0x64, 0xcd,
	0x5e, 0x8c, 0x7b, 0x5f, 0x00, 0xdc, 0x2f, 0xc9, 0xbe, 0x83, 0x68, 0x94, 0xf2, 0xed, 0x46, 0x5d,
	0x56, 0x54, 0xf1, 0x22, 0xe5, 0x45, 0xf6, 0x57, 0xff, 0xe1, 0xc3, 0x51, 0x39, 0x18, 0xb6, 0xe1,
	0x51, 0xa9, 0x5c, 0xbf, 0x1b, 0x6f, 0x04, 0x0f, 0x37, 0x08, 0xd0, 0x1b, 0xb8, 0x43, 0xa2, 0xbf,
	0x1f, 0xc7, 0xd3, 0x8a, 0x9a, 0xae, 0x4c, 0x35, 0x5e, 0xb3, 0x78, 0xdf, 0x00, 0x6c, 0xda, 0x18,
	0x42, 0xb0, 0x9e, 0x91, 0x19, 0x35, 0xc3, 0xb6, 0xb0, 0x39, 0xa3, 0x87, 0xb0, 0x21, 0xe8, 0x84,
	0x2e, 0x5d, 0xa7, 0x53, 0xeb, 0xb6, 0xc2, 0x7d, 0xfd, 0x87, 0x6f, 0x7d, 0x06, 0x4d, 0xaf, 0x2e,
	0x9c, 0x0e, 0xc0, 0x16, 0x44, 0x27, 0xb0, 0x29, 0x63, 0x32, 0x62, 0x1f, 0xf2, 0x1f, 0x38, 0xbf,
	0xa1, 0x0b, 0xb8, 0x93, 0xbf, 0x7e, 0xb7, 0x6e, 0x1c, 0x3c, 0xca, 0xd5, 0xea, 0x01, 0xfd, 0x81,
	0x85, 0xf0, 0x3a, 0x07, 0x3d, 0x82, 0xad, 0x19, 0x91, 0xd3, 0x61, 0x14, 0x13, 0xe1, 0x36, 0xb4,
	0x8a, 0xb0, 0xa5, 0x1b, 0xd6, 0x85, 0xd3, 0x05, 0x78, 0x57, 0x63, 0xfd, 0x98, 0x88, 0xf0, 0x3b,
	0x80, 0xcf, 0x12, 0x66, 0xa9, 0xb8, 0x60, 0xcb, 0x55, 0x35, 0x0f, 0xc2, 0x07, 0x45, 0x53, 0xaf,
	0xaf, 0xed, 0x42, 0x1a, 0xe8, 0x05, 0x35, 0x00, 0xef, 0xdf, 0xfd, 0xdb, 0x42, 0xe6, 0xd3, 0x49,
	0xd5, 0xa5, 0x7c, 0xd3, 0x34, 0x4b, 0xf0, 0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x70,
	0x5d, 0x64, 0xf4, 0x05, 0x00, 0x00,
}
