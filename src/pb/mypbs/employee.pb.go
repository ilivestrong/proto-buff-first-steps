// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: src/protos/crm/employee.proto

package mypbs

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32      `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Addresses []*Address `protobuf:"bytes,3,rep,name=Addresses,proto3" json:"Addresses,omitempty"`
	Project   *Project   `protobuf:"bytes,4,opt,name=Project,proto3" json:"Project,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_crm_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_crm_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_src_protos_crm_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *Employee) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street   string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	Postcode int32  `protobuf:"varint,2,opt,name=postcode,proto3" json:"postcode,omitempty"`
	Building string `protobuf:"bytes,3,opt,name=building,proto3" json:"building,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protos_crm_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_src_protos_crm_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_src_protos_crm_employee_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetPostcode() int32 {
	if x != nil {
		return x.Postcode
	}
	return 0
}

func (x *Address) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

var File_src_protos_crm_employee_proto protoreflect.FileDescriptor

var file_src_protos_crm_employee_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x6d,
	0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x63, 0x72, 0x6d, 0x1a, 0x20, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x72, 0x6d,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x59, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6d,
	0x79, 0x70, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_protos_crm_employee_proto_rawDescOnce sync.Once
	file_src_protos_crm_employee_proto_rawDescData = file_src_protos_crm_employee_proto_rawDesc
)

func file_src_protos_crm_employee_proto_rawDescGZIP() []byte {
	file_src_protos_crm_employee_proto_rawDescOnce.Do(func() {
		file_src_protos_crm_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_protos_crm_employee_proto_rawDescData)
	})
	return file_src_protos_crm_employee_proto_rawDescData
}

var file_src_protos_crm_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_protos_crm_employee_proto_goTypes = []interface{}{
	(*Employee)(nil), // 0: crm.Employee
	(*Address)(nil),  // 1: crm.Address
	(*Project)(nil),  // 2: profile.Project
}
var file_src_protos_crm_employee_proto_depIdxs = []int32{
	1, // 0: crm.Employee.Addresses:type_name -> crm.Address
	2, // 1: crm.Employee.Project:type_name -> profile.Project
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_src_protos_crm_employee_proto_init() }
func file_src_protos_crm_employee_proto_init() {
	if File_src_protos_crm_employee_proto != nil {
		return
	}
	file_src_protos_profile_project_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_src_protos_crm_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_src_protos_crm_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_src_protos_crm_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_protos_crm_employee_proto_goTypes,
		DependencyIndexes: file_src_protos_crm_employee_proto_depIdxs,
		MessageInfos:      file_src_protos_crm_employee_proto_msgTypes,
	}.Build()
	File_src_protos_crm_employee_proto = out.File
	file_src_protos_crm_employee_proto_rawDesc = nil
	file_src_protos_crm_employee_proto_goTypes = nil
	file_src_protos_crm_employee_proto_depIdxs = nil
}