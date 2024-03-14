// 放置ボーナススケジュール

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: master_idle_bonus_schedule_structure.proto

package masterIdleBonusSchedule

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MasterIdleBonusSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MasterIdleBonusId int64  `protobuf:"varint,2,opt,name=master_idle_bonus_id,json=masterIdleBonusId,proto3" json:"master_idle_bonus_id,omitempty"`
	Step              int32  `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	Name              string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MasterIdleBonusSchedule) Reset() {
	*x = MasterIdleBonusSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_idle_bonus_schedule_structure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterIdleBonusSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterIdleBonusSchedule) ProtoMessage() {}

func (x *MasterIdleBonusSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_master_idle_bonus_schedule_structure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterIdleBonusSchedule.ProtoReflect.Descriptor instead.
func (*MasterIdleBonusSchedule) Descriptor() ([]byte, []int) {
	return file_master_idle_bonus_schedule_structure_proto_rawDescGZIP(), []int{0}
}

func (x *MasterIdleBonusSchedule) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MasterIdleBonusSchedule) GetMasterIdleBonusId() int64 {
	if x != nil {
		return x.MasterIdleBonusId
	}
	return 0
}

func (x *MasterIdleBonusSchedule) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *MasterIdleBonusSchedule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_master_idle_bonus_schedule_structure_proto protoreflect.FileDescriptor

var file_master_idle_bonus_schedule_structure_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x17, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2f, 0x0a, 0x14, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_master_idle_bonus_schedule_structure_proto_rawDescOnce sync.Once
	file_master_idle_bonus_schedule_structure_proto_rawDescData = file_master_idle_bonus_schedule_structure_proto_rawDesc
)

func file_master_idle_bonus_schedule_structure_proto_rawDescGZIP() []byte {
	file_master_idle_bonus_schedule_structure_proto_rawDescOnce.Do(func() {
		file_master_idle_bonus_schedule_structure_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_idle_bonus_schedule_structure_proto_rawDescData)
	})
	return file_master_idle_bonus_schedule_structure_proto_rawDescData
}

var file_master_idle_bonus_schedule_structure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_master_idle_bonus_schedule_structure_proto_goTypes = []interface{}{
	(*MasterIdleBonusSchedule)(nil), // 0: proto.MasterIdleBonusSchedule
}
var file_master_idle_bonus_schedule_structure_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_master_idle_bonus_schedule_structure_proto_init() }
func file_master_idle_bonus_schedule_structure_proto_init() {
	if File_master_idle_bonus_schedule_structure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_master_idle_bonus_schedule_structure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterIdleBonusSchedule); i {
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
			RawDescriptor: file_master_idle_bonus_schedule_structure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_master_idle_bonus_schedule_structure_proto_goTypes,
		DependencyIndexes: file_master_idle_bonus_schedule_structure_proto_depIdxs,
		MessageInfos:      file_master_idle_bonus_schedule_structure_proto_msgTypes,
	}.Build()
	File_master_idle_bonus_schedule_structure_proto = out.File
	file_master_idle_bonus_schedule_structure_proto_rawDesc = nil
	file_master_idle_bonus_schedule_structure_proto_goTypes = nil
	file_master_idle_bonus_schedule_structure_proto_depIdxs = nil
}
