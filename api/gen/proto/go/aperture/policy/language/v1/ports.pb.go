// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: aperture/policy/language/v1/ports.proto

package languagev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Components receive input from other components via InPorts
type InPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*InPort_SignalName
	//	*InPort_ConstantSignal
	Value isInPort_Value `protobuf_oneof:"value"`
}

func (x *InPort) Reset() {
	*x = InPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_language_v1_ports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InPort) ProtoMessage() {}

func (x *InPort) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_language_v1_ports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InPort.ProtoReflect.Descriptor instead.
func (*InPort) Descriptor() ([]byte, []int) {
	return file_aperture_policy_language_v1_ports_proto_rawDescGZIP(), []int{0}
}

func (m *InPort) GetValue() isInPort_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *InPort) GetSignalName() string {
	if x, ok := x.GetValue().(*InPort_SignalName); ok {
		return x.SignalName
	}
	return ""
}

func (x *InPort) GetConstantSignal() *ConstantSignal {
	if x, ok := x.GetValue().(*InPort_ConstantSignal); ok {
		return x.ConstantSignal
	}
	return nil
}

type isInPort_Value interface {
	isInPort_Value()
}

type InPort_SignalName struct {
	// Name of the incoming Signal on the InPort.
	SignalName string `protobuf:"bytes,1,opt,name=signal_name,json=signalName,proto3,oneof"`
}

type InPort_ConstantSignal struct {
	// Constant value to be used for this InPort instead of a signal.
	ConstantSignal *ConstantSignal `protobuf:"bytes,2,opt,name=constant_signal,json=constantSignal,proto3,oneof"`
}

func (*InPort_SignalName) isInPort_Value() {}

func (*InPort_ConstantSignal) isInPort_Value() {}

// Components produce output for other components via OutPorts
type OutPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the outgoing Signal on the OutPort.
	SignalName string `protobuf:"bytes,1,opt,name=signal_name,json=signalName,proto3" json:"signal_name,omitempty"`
}

func (x *OutPort) Reset() {
	*x = OutPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_language_v1_ports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutPort) ProtoMessage() {}

func (x *OutPort) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_language_v1_ports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutPort.ProtoReflect.Descriptor instead.
func (*OutPort) Descriptor() ([]byte, []int) {
	return file_aperture_policy_language_v1_ports_proto_rawDescGZIP(), []int{1}
}

func (x *OutPort) GetSignalName() string {
	if x != nil {
		return x.SignalName
	}
	return ""
}

// Special constant input for ports and Variable component. Can provide either a constant value or special Nan/+-Inf value.
type ConstantSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Const:
	//
	//	*ConstantSignal_SpecialValue
	//	*ConstantSignal_Value
	Const isConstantSignal_Const `protobuf_oneof:"const"`
}

func (x *ConstantSignal) Reset() {
	*x = ConstantSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_language_v1_ports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstantSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstantSignal) ProtoMessage() {}

func (x *ConstantSignal) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_language_v1_ports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstantSignal.ProtoReflect.Descriptor instead.
func (*ConstantSignal) Descriptor() ([]byte, []int) {
	return file_aperture_policy_language_v1_ports_proto_rawDescGZIP(), []int{2}
}

func (m *ConstantSignal) GetConst() isConstantSignal_Const {
	if m != nil {
		return m.Const
	}
	return nil
}

func (x *ConstantSignal) GetSpecialValue() string {
	if x, ok := x.GetConst().(*ConstantSignal_SpecialValue); ok {
		return x.SpecialValue
	}
	return ""
}

func (x *ConstantSignal) GetValue() float64 {
	if x, ok := x.GetConst().(*ConstantSignal_Value); ok {
		return x.Value
	}
	return 0
}

type isConstantSignal_Const interface {
	isConstantSignal_Const()
}

type ConstantSignal_SpecialValue struct {
	// A special value such as NaN, +Inf, -Inf.
	SpecialValue string `protobuf:"bytes,1,opt,name=special_value,json=specialValue,proto3,oneof" validate:"oneof=NaN +Inf -Inf"` // @gotags: validate:"oneof=NaN +Inf -Inf"
}

type ConstantSignal_Value struct {
	// A constant value.
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3,oneof"`
}

func (*ConstantSignal_SpecialValue) isConstantSignal_Const() {}

func (*ConstantSignal_Value) isConstantSignal_Const() {}

var File_aperture_policy_language_v1_ports_proto protoreflect.FileDescriptor

var file_aperture_policy_language_v1_ports_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x8c, 0x01, 0x0a, 0x06, 0x49, 0x6e, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2a, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x58, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x42, 0xa7, 0x02, 0x0a, 0x33,
	0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c,
	0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x50, 0x4c, 0xaa, 0x02,
	0x1b, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x41,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x41, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x3a,
	0x3a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x3a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aperture_policy_language_v1_ports_proto_rawDescOnce sync.Once
	file_aperture_policy_language_v1_ports_proto_rawDescData = file_aperture_policy_language_v1_ports_proto_rawDesc
)

func file_aperture_policy_language_v1_ports_proto_rawDescGZIP() []byte {
	file_aperture_policy_language_v1_ports_proto_rawDescOnce.Do(func() {
		file_aperture_policy_language_v1_ports_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_policy_language_v1_ports_proto_rawDescData)
	})
	return file_aperture_policy_language_v1_ports_proto_rawDescData
}

var file_aperture_policy_language_v1_ports_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aperture_policy_language_v1_ports_proto_goTypes = []interface{}{
	(*InPort)(nil),         // 0: aperture.policy.language.v1.InPort
	(*OutPort)(nil),        // 1: aperture.policy.language.v1.OutPort
	(*ConstantSignal)(nil), // 2: aperture.policy.language.v1.ConstantSignal
}
var file_aperture_policy_language_v1_ports_proto_depIdxs = []int32{
	2, // 0: aperture.policy.language.v1.InPort.constant_signal:type_name -> aperture.policy.language.v1.ConstantSignal
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aperture_policy_language_v1_ports_proto_init() }
func file_aperture_policy_language_v1_ports_proto_init() {
	if File_aperture_policy_language_v1_ports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aperture_policy_language_v1_ports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InPort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aperture_policy_language_v1_ports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutPort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aperture_policy_language_v1_ports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstantSignal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_aperture_policy_language_v1_ports_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*InPort_SignalName)(nil),
		(*InPort_ConstantSignal)(nil),
	}
	file_aperture_policy_language_v1_ports_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ConstantSignal_SpecialValue)(nil),
		(*ConstantSignal_Value)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aperture_policy_language_v1_ports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aperture_policy_language_v1_ports_proto_goTypes,
		DependencyIndexes: file_aperture_policy_language_v1_ports_proto_depIdxs,
		MessageInfos:      file_aperture_policy_language_v1_ports_proto_msgTypes,
	}.Build()
	File_aperture_policy_language_v1_ports_proto = out.File
	file_aperture_policy_language_v1_ports_proto_rawDesc = nil
	file_aperture_policy_language_v1_ports_proto_goTypes = nil
	file_aperture_policy_language_v1_ports_proto_depIdxs = nil
}
