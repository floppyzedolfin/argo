// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: portdomain/portdomain.proto

package portdomain

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

type UpsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *UpsertRequest) Reset() {
	*x = UpsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portdomain_portdomain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRequest) ProtoMessage() {}

func (x *UpsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portdomain_portdomain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertRequest.ProtoReflect.Descriptor instead.
func (*UpsertRequest) Descriptor() ([]byte, []int) {
	return file_portdomain_portdomain_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

type UpsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpsertResponse) Reset() {
	*x = UpsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portdomain_portdomain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertResponse) ProtoMessage() {}

func (x *UpsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portdomain_portdomain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertResponse.ProtoReflect.Descriptor instead.
func (*UpsertResponse) Descriptor() ([]byte, []int) {
	return file_portdomain_portdomain_proto_rawDescGZIP(), []int{1}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portdomain_portdomain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portdomain_portdomain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_portdomain_portdomain_proto_rawDescGZIP(), []int{2}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portdomain_portdomain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portdomain_portdomain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_portdomain_portdomain_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string    `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,9,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	Id          string    `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portdomain_portdomain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_portdomain_portdomain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_portdomain_portdomain_proto_rawDescGZIP(), []int{4}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Port) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_portdomain_portdomain_proto protoreflect.FileDescriptor

var file_portdomain_portdomain_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x35, 0x0a, 0x0d, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x8e, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x89, 0x01, 0x0a, 0x0a, 0x70, 0x6f, 0x72,
	0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x6f, 0x70, 0x70, 0x79, 0x7a, 0x65, 0x64, 0x6f, 0x6c, 0x66, 0x69,
	0x6e, 0x2f, 0x61, 0x72, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_portdomain_portdomain_proto_rawDescOnce sync.Once
	file_portdomain_portdomain_proto_rawDescData = file_portdomain_portdomain_proto_rawDesc
)

func file_portdomain_portdomain_proto_rawDescGZIP() []byte {
	file_portdomain_portdomain_proto_rawDescOnce.Do(func() {
		file_portdomain_portdomain_proto_rawDescData = protoimpl.X.CompressGZIP(file_portdomain_portdomain_proto_rawDescData)
	})
	return file_portdomain_portdomain_proto_rawDescData
}

var file_portdomain_portdomain_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_portdomain_portdomain_proto_goTypes = []interface{}{
	(*UpsertRequest)(nil),  // 0: portdomain.UpsertRequest
	(*UpsertResponse)(nil), // 1: portdomain.UpsertResponse
	(*GetRequest)(nil),     // 2: portdomain.GetRequest
	(*GetResponse)(nil),    // 3: portdomain.GetResponse
	(*Port)(nil),           // 4: portdomain.Port
}
var file_portdomain_portdomain_proto_depIdxs = []int32{
	4, // 0: portdomain.UpsertRequest.port:type_name -> portdomain.Port
	4, // 1: portdomain.GetResponse.ports:type_name -> portdomain.Port
	0, // 2: portdomain.portdomain.Upsert:input_type -> portdomain.UpsertRequest
	2, // 3: portdomain.portdomain.Get:input_type -> portdomain.GetRequest
	1, // 4: portdomain.portdomain.Upsert:output_type -> portdomain.UpsertResponse
	3, // 5: portdomain.portdomain.Get:output_type -> portdomain.GetResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_portdomain_portdomain_proto_init() }
func file_portdomain_portdomain_proto_init() {
	if File_portdomain_portdomain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portdomain_portdomain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertRequest); i {
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
		file_portdomain_portdomain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertResponse); i {
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
		file_portdomain_portdomain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_portdomain_portdomain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_portdomain_portdomain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
			RawDescriptor: file_portdomain_portdomain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portdomain_portdomain_proto_goTypes,
		DependencyIndexes: file_portdomain_portdomain_proto_depIdxs,
		MessageInfos:      file_portdomain_portdomain_proto_msgTypes,
	}.Build()
	File_portdomain_portdomain_proto = out.File
	file_portdomain_portdomain_proto_rawDesc = nil
	file_portdomain_portdomain_proto_goTypes = nil
	file_portdomain_portdomain_proto_depIdxs = nil
}
