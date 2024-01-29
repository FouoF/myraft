//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./election/election.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: election/election.proto

package election

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

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPre  bool  `protobuf:"varint,1,opt,name=isPre,proto3" json:"isPre,omitempty"`
	NodeId int32 `protobuf:"varint,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Term   int32 `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_election_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_election_election_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_election_election_proto_rawDescGZIP(), []int{0}
}

func (x *VoteRequest) GetIsPre() bool {
	if x != nil {
		return x.IsPre
	}
	return false
}

func (x *VoteRequest) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *VoteRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteGranted bool  `protobuf:"varint,1,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
	Term        int32 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_election_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_election_election_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_election_election_proto_rawDescGZIP(), []int{1}
}

func (x *VoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

func (x *VoteResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Term int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_election_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_election_election_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_election_election_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HeartbeatRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_election_election_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_election_election_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_election_election_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_election_election_proto protoreflect.FileDescriptor

var file_election_election_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x50, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x69, 0x73, 0x50, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x22, 0x44, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x36, 0x0a, 0x10, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x22, 0x2d, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x8b, 0x01, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35,
	0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0f, 0x5a, 0x0d, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_election_election_proto_rawDescOnce sync.Once
	file_election_election_proto_rawDescData = file_election_election_proto_rawDesc
)

func file_election_election_proto_rawDescGZIP() []byte {
	file_election_election_proto_rawDescOnce.Do(func() {
		file_election_election_proto_rawDescData = protoimpl.X.CompressGZIP(file_election_election_proto_rawDescData)
	})
	return file_election_election_proto_rawDescData
}

var file_election_election_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_election_election_proto_goTypes = []interface{}{
	(*VoteRequest)(nil),       // 0: election.VoteRequest
	(*VoteResponse)(nil),      // 1: election.VoteResponse
	(*HeartbeatRequest)(nil),  // 2: election.HeartbeatRequest
	(*HeartbeatResponse)(nil), // 3: election.HeartbeatResponse
}
var file_election_election_proto_depIdxs = []int32{
	0, // 0: election.election.Vote:input_type -> election.VoteRequest
	2, // 1: election.election.sendheartbeat:input_type -> election.HeartbeatRequest
	1, // 2: election.election.Vote:output_type -> election.VoteResponse
	3, // 3: election.election.sendheartbeat:output_type -> election.HeartbeatResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_election_election_proto_init() }
func file_election_election_proto_init() {
	if File_election_election_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_election_election_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_election_election_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteResponse); i {
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
		file_election_election_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_election_election_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
			RawDescriptor: file_election_election_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_election_election_proto_goTypes,
		DependencyIndexes: file_election_election_proto_depIdxs,
		MessageInfos:      file_election_election_proto_msgTypes,
	}.Build()
	File_election_election_proto = out.File
	file_election_election_proto_rawDesc = nil
	file_election_election_proto_goTypes = nil
	file_election_election_proto_depIdxs = nil
}