// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sys_dict_type.proto

package dicttype

import (
	_ "github.com/feihua/hertz-admin/biz/model/system/api"
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

// 添加字典类型表请求参数
type AddDictTypeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remark   string `protobuf:"bytes,1,opt,name=remark,proto3" form:"remark" json:"remark,omitempty"`                      //备注
	DictName string `protobuf:"bytes,2,opt,name=dict_name,json=dictName,proto3" form:"dictName" json:"dictName,omitempty"` //字典名称
	DictType string `protobuf:"bytes,3,opt,name=dict_type,json=dictType,proto3" form:"dictType" json:"dictType,omitempty"` //字典类型
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" form:"status" json:"status,omitempty"`                     //状态（0：停用，1:正常）
}

func (x *AddDictTypeReq) Reset() {
	*x = AddDictTypeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDictTypeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDictTypeReq) ProtoMessage() {}

func (x *AddDictTypeReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDictTypeReq.ProtoReflect.Descriptor instead.
func (*AddDictTypeReq) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{0}
}

func (x *AddDictTypeReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *AddDictTypeReq) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *AddDictTypeReq) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *AddDictTypeReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 删除字典类型表请求参数
type DeleteDictTypeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" form:"ids" json:"ids,omitempty"`
}

func (x *DeleteDictTypeReq) Reset() {
	*x = DeleteDictTypeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDictTypeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDictTypeReq) ProtoMessage() {}

func (x *DeleteDictTypeReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDictTypeReq.ProtoReflect.Descriptor instead.
func (*DeleteDictTypeReq) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteDictTypeReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 更新字典类型表请求参数
type UpdateDictTypeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty"`                                 //字典主键
	DictName string `protobuf:"bytes,2,opt,name=dict_name,json=dictName,proto3" form:"dictName" json:"dictName,omitempty"` //字典名称
	DictType string `protobuf:"bytes,3,opt,name=dict_type,json=dictType,proto3" form:"dictType" json:"dictType,omitempty"` //字典类型
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" form:"status" json:"status,omitempty"`                     //状态（0：停用，1:正常）
	Remark   string `protobuf:"bytes,5,opt,name=remark,proto3" form:"remark" json:"remark,omitempty"`                      //备注
}

func (x *UpdateDictTypeReq) Reset() {
	*x = UpdateDictTypeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDictTypeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDictTypeReq) ProtoMessage() {}

func (x *UpdateDictTypeReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDictTypeReq.ProtoReflect.Descriptor instead.
func (*UpdateDictTypeReq) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDictTypeReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDictTypeReq) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *UpdateDictTypeReq) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *UpdateDictTypeReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateDictTypeReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

// 更新字典类型表状态请求参数
type UpdateDictTypeStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" form:"id" json:"id,omitempty"`      //字典主键
	Status int32   `protobuf:"varint,4,opt,name=status,proto3" form:"status" json:"status,omitempty"` //状态（0：停用，1:正常）
}

func (x *UpdateDictTypeStatusReq) Reset() {
	*x = UpdateDictTypeStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDictTypeStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDictTypeStatusReq) ProtoMessage() {}

func (x *UpdateDictTypeStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDictTypeStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateDictTypeStatusReq) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDictTypeStatusReq) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdateDictTypeStatusReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 查询字典类型表详情请求参数
type QueryDictTypeDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty"` //字典主键
}

func (x *QueryDictTypeDetailReq) Reset() {
	*x = QueryDictTypeDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDictTypeDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDictTypeDetailReq) ProtoMessage() {}

func (x *QueryDictTypeDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDictTypeDetailReq.ProtoReflect.Descriptor instead.
func (*QueryDictTypeDetailReq) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{4}
}

func (x *QueryDictTypeDetailReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryDictTypeDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty" query:"id"`                                                    //字典主键
	DictName   string `protobuf:"bytes,2,opt,name=dict_name,json=dictName,proto3" form:"dict_name" json:"dict_name,omitempty" query:"dict_name"`           //字典名称
	DictType   string `protobuf:"bytes,3,opt,name=dict_type,json=dictType,proto3" form:"dict_type" json:"dict_type,omitempty" query:"dict_type"`           //字典类型
	Status     int32  `protobuf:"varint,4,opt,name=status,proto3" form:"status" json:"status,omitempty" query:"status"`                                    //状态（0：停用，1:正常）
	Remark     string `protobuf:"bytes,5,opt,name=remark,proto3" form:"remark" json:"remark,omitempty" query:"remark"`                                     //备注
	CreateBy   string `protobuf:"bytes,6,opt,name=create_by,json=createBy,proto3" form:"create_by" json:"create_by,omitempty" query:"create_by"`           //创建者
	CreateTime string `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" form:"create_time" json:"create_time,omitempty" query:"create_time"` //创建时间
	UpdateBy   string `protobuf:"bytes,8,opt,name=update_by,json=updateBy,proto3" form:"update_by" json:"update_by,omitempty" query:"update_by"`           //更新者
	UpdateTime string `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" form:"update_time" json:"update_time,omitempty" query:"update_time"` //更新时间
}

func (x *QueryDictTypeDetailResp) Reset() {
	*x = QueryDictTypeDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDictTypeDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDictTypeDetailResp) ProtoMessage() {}

func (x *QueryDictTypeDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDictTypeDetailResp.ProtoReflect.Descriptor instead.
func (*QueryDictTypeDetailResp) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{5}
}

func (x *QueryDictTypeDetailResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryDictTypeDetailResp) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *QueryDictTypeDetailResp) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *QueryDictTypeDetailResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryDictTypeDetailResp) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *QueryDictTypeDetailResp) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *QueryDictTypeDetailResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *QueryDictTypeDetailResp) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *QueryDictTypeDetailResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// 分页查询字典类型表列表请求参数
type QueryDictTypeListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictName string `protobuf:"bytes,2,opt,name=dict_name,json=dictName,proto3" form:"dictName" json:"dictName,omitempty"`  //字典名称
	DictType string `protobuf:"bytes,3,opt,name=dict_type,json=dictType,proto3" form:"dictType" json:"dictType,omitempty"`  //字典类型
	Status   int32  `protobuf:"varint,4,opt,name=status,proto3" form:"status" json:"status,omitempty"`                      //状态（0：停用，1:正常）
	PageNum  int32  `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" form:"pageNo" json:"pageNo,omitempty"`       //第几页
	PageSize int32  `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" form:"pageSize" json:"pageSize,omitempty"` //每页的数量
}

func (x *QueryDictTypeListReq) Reset() {
	*x = QueryDictTypeListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDictTypeListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDictTypeListReq) ProtoMessage() {}

func (x *QueryDictTypeListReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDictTypeListReq.ProtoReflect.Descriptor instead.
func (*QueryDictTypeListReq) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{6}
}

func (x *QueryDictTypeListReq) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *QueryDictTypeListReq) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *QueryDictTypeListReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryDictTypeListReq) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *QueryDictTypeListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryDictTypeListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty" query:"id"`                                                    //字典主键
	DictName   string `protobuf:"bytes,2,opt,name=dict_name,json=dictName,proto3" form:"dict_name" json:"dict_name,omitempty" query:"dict_name"`           //字典名称
	DictType   string `protobuf:"bytes,3,opt,name=dict_type,json=dictType,proto3" form:"dict_type" json:"dict_type,omitempty" query:"dict_type"`           //字典类型
	Status     int32  `protobuf:"varint,4,opt,name=status,proto3" form:"status" json:"status,omitempty" query:"status"`                                    //状态（0：停用，1:正常）
	Remark     string `protobuf:"bytes,5,opt,name=remark,proto3" form:"remark" json:"remark,omitempty" query:"remark"`                                     //备注
	CreateBy   string `protobuf:"bytes,6,opt,name=create_by,json=createBy,proto3" form:"create_by" json:"create_by,omitempty" query:"create_by"`           //创建者
	CreateTime string `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" form:"create_time" json:"create_time,omitempty" query:"create_time"` //创建时间
	UpdateBy   string `protobuf:"bytes,8,opt,name=update_by,json=updateBy,proto3" form:"update_by" json:"update_by,omitempty" query:"update_by"`           //更新者
	UpdateTime string `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" form:"update_time" json:"update_time,omitempty" query:"update_time"` //更新时间
}

func (x *QueryDictTypeListResp) Reset() {
	*x = QueryDictTypeListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDictTypeListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDictTypeListResp) ProtoMessage() {}

func (x *QueryDictTypeListResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDictTypeListResp.ProtoReflect.Descriptor instead.
func (*QueryDictTypeListResp) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{7}
}

func (x *QueryDictTypeListResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryDictTypeListResp) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *QueryDictTypeListResp) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *QueryDictTypeListResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryDictTypeListResp) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *QueryDictTypeListResp) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *QueryDictTypeListResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *QueryDictTypeListResp) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *QueryDictTypeListResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type DictTypeEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DictTypeEmpty) Reset() {
	*x = DictTypeEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictTypeEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictTypeEmpty) ProtoMessage() {}

func (x *DictTypeEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictTypeEmpty.ProtoReflect.Descriptor instead.
func (*DictTypeEmpty) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_proto_rawDescGZIP(), []int{8}
}

var File_sys_dict_type_proto protoreflect.FileDescriptor

var file_sys_dict_type_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0xbb, 0x18,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x29, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca,
	0xbb, 0x18, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x07, 0xca, 0xbb, 0x18,
	0x03, 0x69, 0x64, 0x73, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xca, 0xbb, 0x18,
	0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08,
	0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca,
	0xbb, 0x18, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x55, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x06, 0xca, 0xbb,
	0x18, 0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x16,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x06, 0xca, 0xbb, 0x18, 0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8f,
	0x02, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xe2, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x09, 0x64, 0x69, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb,
	0x18, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x64, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0xca, 0xbb, 0x18, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x29, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c, 0xca,
	0xbb, 0x18, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x8d, 0x02, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44,
	0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd3, 0x05, 0x0a, 0x0f, 0x44, 0x69, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x0b, 0x41, 0x64,
	0x64, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0xd2, 0xc1, 0x18, 0x20, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6b,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x27, 0xd2, 0xc1, 0x18, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6b, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x27, 0xd2, 0xc1, 0x18, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x7d, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0xd2, 0xc1, 0x18, 0x29, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x1e, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x1f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x2c, 0xd2, 0xc1, 0x18, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x7c,
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x2a, 0xd2, 0xc1, 0x18, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x69, 0x68, 0x75,
	0x61, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x69,
	0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x64,
	0x69, 0x63, 0x74, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_dict_type_proto_rawDescOnce sync.Once
	file_sys_dict_type_proto_rawDescData = file_sys_dict_type_proto_rawDesc
)

func file_sys_dict_type_proto_rawDescGZIP() []byte {
	file_sys_dict_type_proto_rawDescOnce.Do(func() {
		file_sys_dict_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_dict_type_proto_rawDescData)
	})
	return file_sys_dict_type_proto_rawDescData
}

var file_sys_dict_type_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sys_dict_type_proto_goTypes = []interface{}{
	(*AddDictTypeReq)(nil),          // 0: system.AddDictTypeReq
	(*DeleteDictTypeReq)(nil),       // 1: system.DeleteDictTypeReq
	(*UpdateDictTypeReq)(nil),       // 2: system.UpdateDictTypeReq
	(*UpdateDictTypeStatusReq)(nil), // 3: system.UpdateDictTypeStatusReq
	(*QueryDictTypeDetailReq)(nil),  // 4: system.QueryDictTypeDetailReq
	(*QueryDictTypeDetailResp)(nil), // 5: system.QueryDictTypeDetailResp
	(*QueryDictTypeListReq)(nil),    // 6: system.QueryDictTypeListReq
	(*QueryDictTypeListResp)(nil),   // 7: system.QueryDictTypeListResp
	(*DictTypeEmpty)(nil),           // 8: system.DictTypeEmpty
}
var file_sys_dict_type_proto_depIdxs = []int32{
	0, // 0: system.DictTypeHandler.AddDictType:input_type -> system.AddDictTypeReq
	1, // 1: system.DictTypeHandler.DeleteDictType:input_type -> system.DeleteDictTypeReq
	2, // 2: system.DictTypeHandler.UpdateDictType:input_type -> system.UpdateDictTypeReq
	3, // 3: system.DictTypeHandler.UpdateDictTypeStatus:input_type -> system.UpdateDictTypeStatusReq
	4, // 4: system.DictTypeHandler.QueryDictTypeDetail:input_type -> system.QueryDictTypeDetailReq
	6, // 5: system.DictTypeHandler.QueryDictTypeList:input_type -> system.QueryDictTypeListReq
	8, // 6: system.DictTypeHandler.AddDictType:output_type -> system.DictTypeEmpty
	8, // 7: system.DictTypeHandler.DeleteDictType:output_type -> system.DictTypeEmpty
	8, // 8: system.DictTypeHandler.UpdateDictType:output_type -> system.DictTypeEmpty
	8, // 9: system.DictTypeHandler.UpdateDictTypeStatus:output_type -> system.DictTypeEmpty
	5, // 10: system.DictTypeHandler.QueryDictTypeDetail:output_type -> system.QueryDictTypeDetailResp
	7, // 11: system.DictTypeHandler.QueryDictTypeList:output_type -> system.QueryDictTypeListResp
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sys_dict_type_proto_init() }
func file_sys_dict_type_proto_init() {
	if File_sys_dict_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_dict_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDictTypeReq); i {
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
		file_sys_dict_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDictTypeReq); i {
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
		file_sys_dict_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDictTypeReq); i {
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
		file_sys_dict_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDictTypeStatusReq); i {
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
		file_sys_dict_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDictTypeDetailReq); i {
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
		file_sys_dict_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDictTypeDetailResp); i {
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
		file_sys_dict_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDictTypeListReq); i {
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
		file_sys_dict_type_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDictTypeListResp); i {
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
		file_sys_dict_type_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictTypeEmpty); i {
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
			RawDescriptor: file_sys_dict_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_dict_type_proto_goTypes,
		DependencyIndexes: file_sys_dict_type_proto_depIdxs,
		MessageInfos:      file_sys_dict_type_proto_msgTypes,
	}.Build()
	File_sys_dict_type_proto = out.File
	file_sys_dict_type_proto_rawDesc = nil
	file_sys_dict_type_proto_goTypes = nil
	file_sys_dict_type_proto_depIdxs = nil
}
