// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: go-grpc-gateway-sample.proto

package proto

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

type Device_DeviceType int32

const (
	Device_onOff  Device_DeviceType = 0
	Device_dimmer Device_DeviceType = 1
	Device_sensor Device_DeviceType = 2
)

// Enum value maps for Device_DeviceType.
var (
	Device_DeviceType_name = map[int32]string{
		0: "onOff",
		1: "dimmer",
		2: "sensor",
	}
	Device_DeviceType_value = map[string]int32{
		"onOff":  0,
		"dimmer": 1,
		"sensor": 2,
	}
)

func (x Device_DeviceType) Enum() *Device_DeviceType {
	p := new(Device_DeviceType)
	*p = x
	return p
}

func (x Device_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Device_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_go_grpc_gateway_sample_proto_enumTypes[0].Descriptor()
}

func (Device_DeviceType) Type() protoreflect.EnumType {
	return &file_go_grpc_gateway_sample_proto_enumTypes[0]
}

func (x Device_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Device_DeviceType.Descriptor instead.
func (Device_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_go_grpc_gateway_sample_proto_rawDescGZIP(), []int{2, 0}
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_grpc_gateway_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_go_grpc_gateway_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_go_grpc_gateway_sample_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value int32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateDevice) Reset() {
	*x = UpdateDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_grpc_gateway_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDevice) ProtoMessage() {}

func (x *UpdateDevice) ProtoReflect() protoreflect.Message {
	mi := &file_go_grpc_gateway_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDevice.ProtoReflect.Descriptor instead.
func (*UpdateDevice) Descriptor() ([]byte, []int) {
	return file_go_grpc_gateway_sample_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDevice) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDevice) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hardware string            `protobuf:"bytes,2,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Name     string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Location string            `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Type     Device_DeviceType `protobuf:"varint,5,opt,name=type,proto3,enum=proto.Device_DeviceType" json:"type,omitempty"`
	Unit     string            `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	State    int32             `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_grpc_gateway_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_go_grpc_gateway_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_go_grpc_gateway_sample_proto_rawDescGZIP(), []int{2}
}

func (x *Device) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Device) GetHardware() string {
	if x != nil {
		return x.Hardware
	}
	return ""
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Device) GetType() Device_DeviceType {
	if x != nil {
		return x.Type
	}
	return Device_onOff
}

func (x *Device) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Device) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type Devices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device []*Device `protobuf:"bytes,1,rep,name=device,proto3" json:"device,omitempty"`
}

func (x *Devices) Reset() {
	*x = Devices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_grpc_gateway_sample_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Devices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Devices) ProtoMessage() {}

func (x *Devices) ProtoReflect() protoreflect.Message {
	mi := &file_go_grpc_gateway_sample_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Devices.ProtoReflect.Descriptor instead.
func (*Devices) Descriptor() ([]byte, []int) {
	return file_go_grpc_gateway_sample_proto_rawDescGZIP(), []int{3}
}

func (x *Devices) GetDevice() []*Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_grpc_gateway_sample_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_go_grpc_gateway_sample_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_go_grpc_gateway_sample_proto_rawDescGZIP(), []int{4}
}

var File_go_grpc_gateway_sample_proto protoreflect.FileDescriptor

var file_go_grpc_gateway_sample_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xed, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x2f, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x64, 0x69, 0x6d,
	0x6d, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x10,
	0x02, 0x22, 0x30, 0x0a, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd5, 0x01, 0x0a,
	0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c,
	0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6e, 0x62, 0x72, 0x33, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_grpc_gateway_sample_proto_rawDescOnce sync.Once
	file_go_grpc_gateway_sample_proto_rawDescData = file_go_grpc_gateway_sample_proto_rawDesc
)

func file_go_grpc_gateway_sample_proto_rawDescGZIP() []byte {
	file_go_grpc_gateway_sample_proto_rawDescOnce.Do(func() {
		file_go_grpc_gateway_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_grpc_gateway_sample_proto_rawDescData)
	})
	return file_go_grpc_gateway_sample_proto_rawDescData
}

var file_go_grpc_gateway_sample_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_grpc_gateway_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_grpc_gateway_sample_proto_goTypes = []interface{}{
	(Device_DeviceType)(0), // 0: proto.Device.DeviceType
	(*ID)(nil),             // 1: proto.ID
	(*UpdateDevice)(nil),   // 2: proto.UpdateDevice
	(*Device)(nil),         // 3: proto.Device
	(*Devices)(nil),        // 4: proto.Devices
	(*Empty)(nil),          // 5: proto.Empty
}
var file_go_grpc_gateway_sample_proto_depIdxs = []int32{
	0, // 0: proto.Device.type:type_name -> proto.Device.DeviceType
	3, // 1: proto.Devices.device:type_name -> proto.Device
	5, // 2: proto.DeviceService.GetAllDevices:input_type -> proto.Empty
	1, // 3: proto.DeviceService.GetDeviceByID:input_type -> proto.ID
	2, // 4: proto.DeviceService.SwitchDevice:input_type -> proto.UpdateDevice
	3, // 5: proto.DeviceService.RegisterDevice:input_type -> proto.Device
	4, // 6: proto.DeviceService.GetAllDevices:output_type -> proto.Devices
	3, // 7: proto.DeviceService.GetDeviceByID:output_type -> proto.Device
	3, // 8: proto.DeviceService.SwitchDevice:output_type -> proto.Device
	3, // 9: proto.DeviceService.RegisterDevice:output_type -> proto.Device
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_grpc_gateway_sample_proto_init() }
func file_go_grpc_gateway_sample_proto_init() {
	if File_go_grpc_gateway_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_grpc_gateway_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_go_grpc_gateway_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDevice); i {
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
		file_go_grpc_gateway_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_go_grpc_gateway_sample_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Devices); i {
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
		file_go_grpc_gateway_sample_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_grpc_gateway_sample_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_grpc_gateway_sample_proto_goTypes,
		DependencyIndexes: file_go_grpc_gateway_sample_proto_depIdxs,
		EnumInfos:         file_go_grpc_gateway_sample_proto_enumTypes,
		MessageInfos:      file_go_grpc_gateway_sample_proto_msgTypes,
	}.Build()
	File_go_grpc_gateway_sample_proto = out.File
	file_go_grpc_gateway_sample_proto_rawDesc = nil
	file_go_grpc_gateway_sample_proto_goTypes = nil
	file_go_grpc_gateway_sample_proto_depIdxs = nil
}
