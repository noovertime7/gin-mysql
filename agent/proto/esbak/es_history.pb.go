// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: es_history.proto

package esbak

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

type EsHistoryEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EsHistoryEmpty) Reset() {
	*x = EsHistoryEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsHistoryEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsHistoryEmpty) ProtoMessage() {}

func (x *EsHistoryEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsHistoryEmpty.ProtoReflect.Descriptor instead.
func (*EsHistoryEmpty) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{0}
}

type GetEsHistoryListInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      string `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
	PageNo    int64  `protobuf:"varint,2,opt,name=PageNo,proto3" json:"PageNo,omitempty"`
	PageSize  int64  `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	SortField string `protobuf:"bytes,5,opt,name=SortField,proto3" json:"SortField,omitempty"`
	SortOrder string `protobuf:"bytes,6,opt,name=SortOrder,proto3" json:"SortOrder,omitempty"`
	Status    string `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GetEsHistoryListInput) Reset() {
	*x = GetEsHistoryListInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEsHistoryListInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEsHistoryListInput) ProtoMessage() {}

func (x *GetEsHistoryListInput) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEsHistoryListInput.ProtoReflect.Descriptor instead.
func (*GetEsHistoryListInput) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{1}
}

func (x *GetEsHistoryListInput) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *GetEsHistoryListInput) GetPageNo() int64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *GetEsHistoryListInput) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetEsHistoryListInput) GetSortField() string {
	if x != nil {
		return x.SortField
	}
	return ""
}

func (x *GetEsHistoryListInput) GetSortOrder() string {
	if x != nil {
		return x.SortOrder
	}
	return ""
}

func (x *GetEsHistoryListInput) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ESHistoryIDInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ESHistoryIDInput) Reset() {
	*x = ESHistoryIDInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryIDInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryIDInput) ProtoMessage() {}

func (x *ESHistoryIDInput) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryIDInput.ProtoReflect.Descriptor instead.
func (*ESHistoryIDInput) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{2}
}

func (x *ESHistoryIDInput) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ESHistoryOneMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	OK      bool   `protobuf:"varint,2,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *ESHistoryOneMessage) Reset() {
	*x = ESHistoryOneMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryOneMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryOneMessage) ProtoMessage() {}

func (x *ESHistoryOneMessage) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryOneMessage.ProtoReflect.Descriptor instead.
func (*ESHistoryOneMessage) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{3}
}

func (x *ESHistoryOneMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ESHistoryOneMessage) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

type ESHistoryListOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"total"`
	// @inject_tag: json:"list"
	EsHistoryListOutItem []*ESHistoryListOutItem `protobuf:"bytes,2,rep,name=esHistoryListOutItem,proto3" json:"list"`
	// @inject_tag: json:"page_no"
	PageNo int64 `protobuf:"varint,3,opt,name=PageNo,proto3" json:"page_no"`
	// @inject_tag: json:"page_size"
	PageSize int64 `protobuf:"varint,4,opt,name=PageSize,proto3" json:"page_size"`
}

func (x *ESHistoryListOutput) Reset() {
	*x = ESHistoryListOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryListOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryListOutput) ProtoMessage() {}

func (x *ESHistoryListOutput) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryListOutput.ProtoReflect.Descriptor instead.
func (*ESHistoryListOutput) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{4}
}

func (x *ESHistoryListOutput) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ESHistoryListOutput) GetEsHistoryListOutItem() []*ESHistoryListOutItem {
	if x != nil {
		return x.EsHistoryListOutItem
	}
	return nil
}

func (x *ESHistoryListOutput) GetPageNo() int64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *ESHistoryListOutput) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ESHistoryListOutItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"task_id"
	TaskID int64 `protobuf:"varint,9,opt,name=TaskID,proto3" json:"task_id"`
	// @inject_tag: json:"uuid"
	UUID string `protobuf:"bytes,10,opt,name=UUID,proto3" json:"uuid"`
	// @inject_tag: json:"duration_in_millis"
	DurationInMillis int64 `protobuf:"varint,11,opt,name=DurationInMillis,proto3" json:"duration_in_millis"`
	// @inject_tag: json:"host"
	Host string `protobuf:"bytes,12,opt,name=Host,proto3" json:"host"`
	// @inject_tag: json:"snapshot"
	Snapshot string `protobuf:"bytes,2,opt,name=Snapshot,proto3" json:"snapshot"`
	// @inject_tag: json:"repository"
	Repository string `protobuf:"bytes,3,opt,name=Repository,proto3" json:"repository"`
	// @inject_tag: json:"indices"
	Indices string `protobuf:"bytes,4,opt,name=Indices,proto3" json:"indices"`
	// @inject_tag: json:"state"
	State string `protobuf:"bytes,5,opt,name=State,proto3" json:"state"`
	// @inject_tag: json:"start_time"
	StartTime string `protobuf:"bytes,6,opt,name=StartTime,proto3" json:"start_time"`
	// @inject_tag: json:"end_time"
	EndTime string `protobuf:"bytes,7,opt,name=EndTime,proto3" json:"end_time"`
	// @inject_tag: json:"message"
	Message string `protobuf:"bytes,8,opt,name=Message,proto3" json:"message"`
	// @inject_tag: json:"status"
	Status int64 `protobuf:"varint,13,opt,name=Status,proto3" json:"status"`
}

func (x *ESHistoryListOutItem) Reset() {
	*x = ESHistoryListOutItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryListOutItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryListOutItem) ProtoMessage() {}

func (x *ESHistoryListOutItem) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryListOutItem.ProtoReflect.Descriptor instead.
func (*ESHistoryListOutItem) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{5}
}

func (x *ESHistoryListOutItem) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ESHistoryListOutItem) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *ESHistoryListOutItem) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *ESHistoryListOutItem) GetDurationInMillis() int64 {
	if x != nil {
		return x.DurationInMillis
	}
	return 0
}

func (x *ESHistoryListOutItem) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ESHistoryListOutItem) GetSnapshot() string {
	if x != nil {
		return x.Snapshot
	}
	return ""
}

func (x *ESHistoryListOutItem) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ESHistoryListOutItem) GetIndices() string {
	if x != nil {
		return x.Indices
	}
	return ""
}

func (x *ESHistoryListOutItem) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ESHistoryListOutItem) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ESHistoryListOutItem) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *ESHistoryListOutItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ESHistoryListOutItem) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 获取历史记录顶部数量信息
type EsHistoryNumInfoOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"all_nums"
	AllNums int64 `protobuf:"varint,1,opt,name=AllNums,proto3" json:"all_nums"`
	// @inject_tag: json:"week_nums"
	WeekNums int64 `protobuf:"varint,2,opt,name=WeekNums,proto3" json:"week_nums"`
	// @inject_tag: json:"fail_nums"
	FailNums int64 `protobuf:"varint,3,opt,name=FailNums,proto3" json:"fail_nums"`
}

func (x *EsHistoryNumInfoOut) Reset() {
	*x = EsHistoryNumInfoOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsHistoryNumInfoOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsHistoryNumInfoOut) ProtoMessage() {}

func (x *EsHistoryNumInfoOut) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsHistoryNumInfoOut.ProtoReflect.Descriptor instead.
func (*EsHistoryNumInfoOut) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{6}
}

func (x *EsHistoryNumInfoOut) GetAllNums() int64 {
	if x != nil {
		return x.AllNums
	}
	return 0
}

func (x *EsHistoryNumInfoOut) GetWeekNums() int64 {
	if x != nil {
		return x.WeekNums
	}
	return 0
}

func (x *EsHistoryNumInfoOut) GetFailNums() int64 {
	if x != nil {
		return x.FailNums
	}
	return 0
}

// 查看历史记录详情
type EsHostDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"host_id"
	HostID int64 `protobuf:"varint,4,opt,name=HostID,proto3" json:"host_id"`
	// @inject_tag: json:"host"
	Host string `protobuf:"bytes,1,opt,name=Host,proto3" json:"host"`
	// @inject_tag: json:"create_at"
	CreateAt string `protobuf:"bytes,2,opt,name=CreateAt,proto3" json:"create_at"`
	// @inject_tag: json:"update_at"
	UpdateAt string `protobuf:"bytes,3,opt,name=UpdateAt,proto3" json:"update_at"`
	// @inject_tag: json:"status"
	Status int64 `protobuf:"varint,5,opt,name=Status,proto3" json:"status"`
}

func (x *EsHostDetail) Reset() {
	*x = EsHostDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsHostDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsHostDetail) ProtoMessage() {}

func (x *EsHostDetail) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsHostDetail.ProtoReflect.Descriptor instead.
func (*EsHostDetail) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{7}
}

func (x *EsHostDetail) GetHostID() int64 {
	if x != nil {
		return x.HostID
	}
	return 0
}

func (x *EsHostDetail) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *EsHostDetail) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *EsHostDetail) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *EsHostDetail) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ESTaskDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"backup_cycle"
	BackupCycle string `protobuf:"bytes,1,opt,name=BackupCycle,proto3" json:"backup_cycle"`
	// @inject_tag: json:"keep_number"
	KeepNumber int64 `protobuf:"varint,2,opt,name=KeepNumber,proto3" json:"keep_number"`
	// @inject_tag: json:"status"
	Status int64 `protobuf:"varint,3,opt,name=Status,proto3" json:"status"`
	// @inject_tag: json:"create_at"
	CreateAt string `protobuf:"bytes,4,opt,name=CreateAt,proto3" json:"create_at"`
}

func (x *ESTaskDetail) Reset() {
	*x = ESTaskDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESTaskDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESTaskDetail) ProtoMessage() {}

func (x *ESTaskDetail) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESTaskDetail.ProtoReflect.Descriptor instead.
func (*ESTaskDetail) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{8}
}

func (x *ESTaskDetail) GetBackupCycle() string {
	if x != nil {
		return x.BackupCycle
	}
	return ""
}

func (x *ESTaskDetail) GetKeepNumber() int64 {
	if x != nil {
		return x.KeepNumber
	}
	return 0
}

func (x *ESTaskDetail) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ESTaskDetail) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type EsHistoryDetailOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"es_host_detail"
	EsHostDetail *EsHostDetail `protobuf:"bytes,1,opt,name=EsHostDetail,proto3" json:"es_host_detail"`
	// @inject_tag: json:"es_task_detail"
	ESTaskDetail *ESTaskDetail `protobuf:"bytes,2,opt,name=ESTaskDetail,proto3" json:"es_task_detail"`
	// @inject_tag: json:"es_history_detail"
	EsHistoryDetail *ESHistoryListOutItem `protobuf:"bytes,3,opt,name=EsHistoryDetail,proto3" json:"es_history_detail"`
}

func (x *EsHistoryDetailOut) Reset() {
	*x = EsHistoryDetailOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_es_history_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsHistoryDetailOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsHistoryDetailOut) ProtoMessage() {}

func (x *EsHistoryDetailOut) ProtoReflect() protoreflect.Message {
	mi := &file_es_history_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsHistoryDetailOut.ProtoReflect.Descriptor instead.
func (*EsHistoryDetailOut) Descriptor() ([]byte, []int) {
	return file_es_history_proto_rawDescGZIP(), []int{9}
}

func (x *EsHistoryDetailOut) GetEsHostDetail() *EsHostDetail {
	if x != nil {
		return x.EsHostDetail
	}
	return nil
}

func (x *EsHistoryDetailOut) GetESTaskDetail() *ESTaskDetail {
	if x != nil {
		return x.ESTaskDetail
	}
	return nil
}

func (x *EsHistoryDetailOut) GetEsHistoryDetail() *ESHistoryListOutItem {
	if x != nil {
		return x.EsHistoryDetail
	}
	return nil
}

var File_es_history_proto protoreflect.FileDescriptor

var file_es_history_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x73, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0xb3, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x45, 0x53, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x13,
	0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x4f, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x4b, 0x22, 0xc7, 0x01,
	0x0a, 0x13, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x66, 0x0a, 0x14, 0x65,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x14, 0x65,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x14, 0x45, 0x53, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6e, 0x64, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x67, 0x0a, 0x13, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4e,
	0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x6c, 0x6c,
	0x4e, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x6c, 0x6c, 0x4e,
	0x75, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x65, 0x65, 0x6b, 0x4e, 0x75, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x57, 0x65, 0x65, 0x6b, 0x4e, 0x75, 0x6d, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x75, 0x6d, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0c,
	0x45, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x6f,
	0x73, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x45, 0x53, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4b,
	0x65, 0x65, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x4b, 0x65, 0x65, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22,
	0x92, 0x02, 0x0a, 0x12, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x4e, 0x0a, 0x0c, 0x45, 0x73, 0x48, 0x6f, 0x73, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x73, 0x48, 0x6f,
	0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x45, 0x73, 0x48, 0x6f, 0x73, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x4e, 0x0a, 0x0c, 0x45, 0x53, 0x54, 0x61, 0x73, 0x6b,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x45, 0x53, 0x54, 0x61, 0x73, 0x6b,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x5c, 0x0a, 0x0f, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45,
	0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x0f, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x32, 0xfc, 0x03, 0x0a, 0x10, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x78, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x31,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x73,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75,
	0x74, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x65, 0x73, 0x62, 0x61, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_es_history_proto_rawDescOnce sync.Once
	file_es_history_proto_rawDescData = file_es_history_proto_rawDesc
)

func file_es_history_proto_rawDescGZIP() []byte {
	file_es_history_proto_rawDescOnce.Do(func() {
		file_es_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_es_history_proto_rawDescData)
	})
	return file_es_history_proto_rawDescData
}

var file_es_history_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_es_history_proto_goTypes = []interface{}{
	(*EsHistoryEmpty)(nil),        // 0: go.micro.service.backupAgent.EsHistoryEmpty
	(*GetEsHistoryListInput)(nil), // 1: go.micro.service.backupAgent.GetEsHistoryListInput
	(*ESHistoryIDInput)(nil),      // 2: go.micro.service.backupAgent.ESHistoryIDInput
	(*ESHistoryOneMessage)(nil),   // 3: go.micro.service.backupAgent.ESHistoryOneMessage
	(*ESHistoryListOutput)(nil),   // 4: go.micro.service.backupAgent.ESHistoryListOutput
	(*ESHistoryListOutItem)(nil),  // 5: go.micro.service.backupAgent.ESHistoryListOutItem
	(*EsHistoryNumInfoOut)(nil),   // 6: go.micro.service.backupAgent.EsHistoryNumInfoOut
	(*EsHostDetail)(nil),          // 7: go.micro.service.backupAgent.EsHostDetail
	(*ESTaskDetail)(nil),          // 8: go.micro.service.backupAgent.ESTaskDetail
	(*EsHistoryDetailOut)(nil),    // 9: go.micro.service.backupAgent.EsHistoryDetailOut
}
var file_es_history_proto_depIdxs = []int32{
	5, // 0: go.micro.service.backupAgent.ESHistoryListOutput.esHistoryListOutItem:type_name -> go.micro.service.backupAgent.ESHistoryListOutItem
	7, // 1: go.micro.service.backupAgent.EsHistoryDetailOut.EsHostDetail:type_name -> go.micro.service.backupAgent.EsHostDetail
	8, // 2: go.micro.service.backupAgent.EsHistoryDetailOut.ESTaskDetail:type_name -> go.micro.service.backupAgent.ESTaskDetail
	5, // 3: go.micro.service.backupAgent.EsHistoryDetailOut.EsHistoryDetail:type_name -> go.micro.service.backupAgent.ESHistoryListOutItem
	1, // 4: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryList:input_type -> go.micro.service.backupAgent.GetEsHistoryListInput
	2, // 5: go.micro.service.backupAgent.EsHistoryService.DeleteESHistory:input_type -> go.micro.service.backupAgent.ESHistoryIDInput
	2, // 6: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryDetail:input_type -> go.micro.service.backupAgent.ESHistoryIDInput
	0, // 7: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryNumInfo:input_type -> go.micro.service.backupAgent.EsHistoryEmpty
	4, // 8: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryList:output_type -> go.micro.service.backupAgent.ESHistoryListOutput
	3, // 9: go.micro.service.backupAgent.EsHistoryService.DeleteESHistory:output_type -> go.micro.service.backupAgent.ESHistoryOneMessage
	9, // 10: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryDetail:output_type -> go.micro.service.backupAgent.EsHistoryDetailOut
	6, // 11: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryNumInfo:output_type -> go.micro.service.backupAgent.EsHistoryNumInfoOut
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_es_history_proto_init() }
func file_es_history_proto_init() {
	if File_es_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_es_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsHistoryEmpty); i {
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
		file_es_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEsHistoryListInput); i {
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
		file_es_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryIDInput); i {
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
		file_es_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryOneMessage); i {
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
		file_es_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryListOutput); i {
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
		file_es_history_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryListOutItem); i {
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
		file_es_history_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsHistoryNumInfoOut); i {
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
		file_es_history_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsHostDetail); i {
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
		file_es_history_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESTaskDetail); i {
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
		file_es_history_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsHistoryDetailOut); i {
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
			RawDescriptor: file_es_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_es_history_proto_goTypes,
		DependencyIndexes: file_es_history_proto_depIdxs,
		MessageInfos:      file_es_history_proto_msgTypes,
	}.Build()
	File_es_history_proto = out.File
	file_es_history_proto_rawDesc = nil
	file_es_history_proto_goTypes = nil
	file_es_history_proto_depIdxs = nil
}
