// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.18.1
// source: bright/overview/overview.proto

package overview

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "yun.tea/block/bright/proto/bright"
	_ "yun.tea/block/bright/proto/bright/basetype"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TimeNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeAt uint32 `protobuf:"varint,10,opt,name=TimeAt,proto3" json:"TimeAt,omitempty"`
	Num    uint64 `protobuf:"varint,20,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *TimeNum) Reset() {
	*x = TimeNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bright_overview_overview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeNum) ProtoMessage() {}

func (x *TimeNum) ProtoReflect() protoreflect.Message {
	mi := &file_bright_overview_overview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeNum.ProtoReflect.Descriptor instead.
func (*TimeNum) Descriptor() ([]byte, []int) {
	return file_bright_overview_overview_proto_rawDescGZIP(), []int{0}
}

func (x *TimeNum) GetTimeAt() uint32 {
	if x != nil {
		return x.TimeAt
	}
	return 0
}

func (x *TimeNum) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Overview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverviewAt        uint32            `protobuf:"varint,10,opt,name=OverviewAt,proto3" json:"OverviewAt,omitempty"`
	ChainName         string            `protobuf:"bytes,20,opt,name=ChainName,proto3" json:"ChainName,omitempty"`
	ChainID           string            `protobuf:"bytes,30,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	ChainExplore      string            `protobuf:"bytes,40,opt,name=ChainExplore,proto3" json:"ChainExplore,omitempty"`
	ContractLang      string            `protobuf:"bytes,50,opt,name=ContractLang,proto3" json:"ContractLang,omitempty"`
	EndpointNum       uint32            `protobuf:"varint,60,opt,name=EndpointNum,proto3" json:"EndpointNum,omitempty"`
	EndpointStatesNum map[string]uint32 `protobuf:"bytes,70,rep,name=EndpointStatesNum,proto3" json:"EndpointStatesNum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	AccountNum        uint32            `protobuf:"varint,80,opt,name=AccountNum,proto3" json:"AccountNum,omitempty"`
	AccountStatesNum  map[string]uint32 `protobuf:"bytes,90,rep,name=AccountStatesNum,proto3" json:"AccountStatesNum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ContractTopicNum  uint32            `protobuf:"varint,100,opt,name=ContractTopicNum,proto3" json:"ContractTopicNum,omitempty"`
	BlockNums         []*TimeNum        `protobuf:"bytes,110,rep,name=BlockNums,proto3" json:"BlockNums,omitempty"`
	TxNums            []*TimeNum        `protobuf:"bytes,120,rep,name=TxNums,proto3" json:"TxNums,omitempty"`
}

func (x *Overview) Reset() {
	*x = Overview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bright_overview_overview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Overview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Overview) ProtoMessage() {}

func (x *Overview) ProtoReflect() protoreflect.Message {
	mi := &file_bright_overview_overview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Overview.ProtoReflect.Descriptor instead.
func (*Overview) Descriptor() ([]byte, []int) {
	return file_bright_overview_overview_proto_rawDescGZIP(), []int{1}
}

func (x *Overview) GetOverviewAt() uint32 {
	if x != nil {
		return x.OverviewAt
	}
	return 0
}

func (x *Overview) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *Overview) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *Overview) GetChainExplore() string {
	if x != nil {
		return x.ChainExplore
	}
	return ""
}

func (x *Overview) GetContractLang() string {
	if x != nil {
		return x.ContractLang
	}
	return ""
}

func (x *Overview) GetEndpointNum() uint32 {
	if x != nil {
		return x.EndpointNum
	}
	return 0
}

func (x *Overview) GetEndpointStatesNum() map[string]uint32 {
	if x != nil {
		return x.EndpointStatesNum
	}
	return nil
}

func (x *Overview) GetAccountNum() uint32 {
	if x != nil {
		return x.AccountNum
	}
	return 0
}

func (x *Overview) GetAccountStatesNum() map[string]uint32 {
	if x != nil {
		return x.AccountStatesNum
	}
	return nil
}

func (x *Overview) GetContractTopicNum() uint32 {
	if x != nil {
		return x.ContractTopicNum
	}
	return 0
}

func (x *Overview) GetBlockNums() []*TimeNum {
	if x != nil {
		return x.BlockNums
	}
	return nil
}

func (x *Overview) GetTxNums() []*TimeNum {
	if x != nil {
		return x.TxNums
	}
	return nil
}

type GetOverviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOverviewRequest) Reset() {
	*x = GetOverviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bright_overview_overview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverviewRequest) ProtoMessage() {}

func (x *GetOverviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bright_overview_overview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverviewRequest.ProtoReflect.Descriptor instead.
func (*GetOverviewRequest) Descriptor() ([]byte, []int) {
	return file_bright_overview_overview_proto_rawDescGZIP(), []int{2}
}

type GetOverviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Overview `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetOverviewResponse) Reset() {
	*x = GetOverviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bright_overview_overview_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverviewResponse) ProtoMessage() {}

func (x *GetOverviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bright_overview_overview_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverviewResponse.ProtoReflect.Descriptor instead.
func (*GetOverviewResponse) Descriptor() ([]byte, []int) {
	return file_bright_overview_overview_proto_rawDescGZIP(), []int{3}
}

func (x *GetOverviewResponse) GetInfo() *Overview {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_bright_overview_overview_proto protoreflect.FileDescriptor

var file_bright_overview_overview_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x1a, 0x13, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x22, 0xca, 0x05, 0x0a, 0x08, 0x4f, 0x76,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4c, 0x61,
	0x6e, 0x67, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x5e, 0x0a, 0x11, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x46, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x6f, 0x76, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x4e, 0x75, 0x6d,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x5b, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x5a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x75,
	0x6d, 0x12, 0x36, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x6e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x6f, 0x76,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x52, 0x09,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x54, 0x78, 0x4e,
	0x75, 0x6d, 0x73, 0x18, 0x78, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x52, 0x06, 0x54, 0x78, 0x4e, 0x75, 0x6d, 0x73, 0x1a, 0x44, 0x0a, 0x16, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x4e, 0x75, 0x6d,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x43, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x4e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0x7b, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x70, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x23, 0x2e, 0x62,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a,
	0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x42,
	0x1f, 0x5a, 0x1d, 0x79, 0x75, 0x6e, 0x2e, 0x74, 0x65, 0x61, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2f, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bright_overview_overview_proto_rawDescOnce sync.Once
	file_bright_overview_overview_proto_rawDescData = file_bright_overview_overview_proto_rawDesc
)

func file_bright_overview_overview_proto_rawDescGZIP() []byte {
	file_bright_overview_overview_proto_rawDescOnce.Do(func() {
		file_bright_overview_overview_proto_rawDescData = protoimpl.X.CompressGZIP(file_bright_overview_overview_proto_rawDescData)
	})
	return file_bright_overview_overview_proto_rawDescData
}

var file_bright_overview_overview_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bright_overview_overview_proto_goTypes = []interface{}{
	(*TimeNum)(nil),             // 0: bright.overview.TimeNum
	(*Overview)(nil),            // 1: bright.overview.Overview
	(*GetOverviewRequest)(nil),  // 2: bright.overview.GetOverviewRequest
	(*GetOverviewResponse)(nil), // 3: bright.overview.GetOverviewResponse
	nil,                         // 4: bright.overview.Overview.EndpointStatesNumEntry
	nil,                         // 5: bright.overview.Overview.AccountStatesNumEntry
}
var file_bright_overview_overview_proto_depIdxs = []int32{
	4, // 0: bright.overview.Overview.EndpointStatesNum:type_name -> bright.overview.Overview.EndpointStatesNumEntry
	5, // 1: bright.overview.Overview.AccountStatesNum:type_name -> bright.overview.Overview.AccountStatesNumEntry
	0, // 2: bright.overview.Overview.BlockNums:type_name -> bright.overview.TimeNum
	0, // 3: bright.overview.Overview.TxNums:type_name -> bright.overview.TimeNum
	1, // 4: bright.overview.GetOverviewResponse.Info:type_name -> bright.overview.Overview
	2, // 5: bright.overview.Manager.GetOverview:input_type -> bright.overview.GetOverviewRequest
	3, // 6: bright.overview.Manager.GetOverview:output_type -> bright.overview.GetOverviewResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bright_overview_overview_proto_init() }
func file_bright_overview_overview_proto_init() {
	if File_bright_overview_overview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bright_overview_overview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeNum); i {
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
		file_bright_overview_overview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Overview); i {
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
		file_bright_overview_overview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOverviewRequest); i {
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
		file_bright_overview_overview_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOverviewResponse); i {
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
			RawDescriptor: file_bright_overview_overview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bright_overview_overview_proto_goTypes,
		DependencyIndexes: file_bright_overview_overview_proto_depIdxs,
		MessageInfos:      file_bright_overview_overview_proto_msgTypes,
	}.Build()
	File_bright_overview_overview_proto = out.File
	file_bright_overview_overview_proto_rawDesc = nil
	file_bright_overview_overview_proto_goTypes = nil
	file_bright_overview_overview_proto_depIdxs = nil
}