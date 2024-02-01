// idl/hello/hello.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: menu.proto

package menu

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	api "hertz_admin/biz/model/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// menu params start -----------
type MenuListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MenuListReq) Reset() {
	*x = MenuListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuListReq) ProtoMessage() {}

func (x *MenuListReq) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuListReq.ProtoReflect.Descriptor instead.
func (*MenuListReq) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{0}
}

type MenuListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string          `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
	Code api.Code        `protobuf:"varint,2,opt,name=code,proto3,enum=api.Code" json:"code,omitempty" form:"code" query:"code"`
	Data []*MenuListData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *MenuListResp) Reset() {
	*x = MenuListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuListResp) ProtoMessage() {}

func (x *MenuListResp) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuListResp.ProtoReflect.Descriptor instead.
func (*MenuListResp) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{1}
}

func (x *MenuListResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MenuListResp) GetCode() api.Code {
	if x != nil {
		return x.Code
	}
	return api.Code(0)
}

func (x *MenuListResp) GetData() []*MenuListData {
	if x != nil {
		return x.Data
	}
	return nil
}

type MenuListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	CreateTime string `protobuf:"bytes,2,opt,name=createTime,proto3" json:"createTime,omitempty" form:"createTime" query:"createTime"`
	UpdateTime string `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty" form:"updateTime" query:"updateTime"`
	StatusId   int32  `protobuf:"varint,4,opt,name=statusId,proto3" json:"statusId,omitempty" form:"statusId" query:"statusId"`
	Sort       int32  `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty" form:"sort" query:"sort"`
	ParentId   int64  `protobuf:"varint,6,opt,name=parentId,proto3" json:"parentId,omitempty" form:"parentId" query:"parentId"`
	MenuName   string `protobuf:"bytes,7,opt,name=menuName,proto3" json:"menuName,omitempty" form:"menuName" query:"menuName"`
	MenuUrl    string `protobuf:"bytes,8,opt,name=menuUrl,proto3" json:"menuUrl,omitempty" form:"menuUrl" query:"menuUrl"`
	ApiUrl     string `protobuf:"bytes,9,opt,name=apiUrl,proto3" json:"apiUrl,omitempty" form:"apiUrl" query:"apiUrl"`
	MenuIcon   string `protobuf:"bytes,10,opt,name=menuIcon,proto3" json:"menuIcon,omitempty" form:"menuIcon" query:"menuIcon"`
	Remark     string `protobuf:"bytes,11,opt,name=remark,proto3" json:"remark,omitempty" form:"remark" query:"remark"`
	MenuType   int32  `protobuf:"varint,12,opt,name=menuType,proto3" json:"menuType,omitempty" form:"menuType" query:"menuType"`
}

func (x *MenuListData) Reset() {
	*x = MenuListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuListData) ProtoMessage() {}

func (x *MenuListData) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuListData.ProtoReflect.Descriptor instead.
func (*MenuListData) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{2}
}

func (x *MenuListData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuListData) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *MenuListData) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *MenuListData) GetStatusId() int32 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

func (x *MenuListData) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuListData) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MenuListData) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *MenuListData) GetMenuUrl() string {
	if x != nil {
		return x.MenuUrl
	}
	return ""
}

func (x *MenuListData) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *MenuListData) GetMenuIcon() string {
	if x != nil {
		return x.MenuIcon
	}
	return ""
}

func (x *MenuListData) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MenuListData) GetMenuType() int32 {
	if x != nil {
		return x.MenuType
	}
	return 0
}

type MenuSaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusId int32  `protobuf:"varint,1,opt,name=StatusId,proto3" form:"status_id" json:"status_id,omitempty"`
	Sort     int32  `protobuf:"varint,2,opt,name=Sort,proto3" form:"sort" json:"sort,omitempty"`
	ParentId int64  `protobuf:"varint,3,opt,name=ParentId,proto3" form:"parent_id" json:"parent_id,omitempty"`
	MenuName string `protobuf:"bytes,4,opt,name=MenuName,proto3" form:"menu_name" json:"menu_name,omitempty"`
	MenuUrl  string `protobuf:"bytes,5,opt,name=MenuUrl,proto3" form:"menu_url" json:"menu_url,omitempty"`
	ApiUrl   string `protobuf:"bytes,6,opt,name=ApiUrl,proto3" form:"api_url" json:"api_url,omitempty"`
	MenuIcon string `protobuf:"bytes,7,opt,name=MenuIcon,proto3" form:"menu_icon" json:"menu_icon,omitempty"`
	Remark   string `protobuf:"bytes,8,opt,name=Remark,proto3" form:"remark" json:"remark,omitempty"`
	MenuType int32  `protobuf:"varint,9,opt,name=MenuType,proto3" form:"menu_type" json:"menu_type,omitempty"`
}

func (x *MenuSaveReq) Reset() {
	*x = MenuSaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuSaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuSaveReq) ProtoMessage() {}

func (x *MenuSaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuSaveReq.ProtoReflect.Descriptor instead.
func (*MenuSaveReq) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{3}
}

func (x *MenuSaveReq) GetStatusId() int32 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

func (x *MenuSaveReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuSaveReq) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MenuSaveReq) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *MenuSaveReq) GetMenuUrl() string {
	if x != nil {
		return x.MenuUrl
	}
	return ""
}

func (x *MenuSaveReq) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *MenuSaveReq) GetMenuIcon() string {
	if x != nil {
		return x.MenuIcon
	}
	return ""
}

func (x *MenuSaveReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MenuSaveReq) GetMenuType() int32 {
	if x != nil {
		return x.MenuType
	}
	return 0
}

type MenuSaveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
	Code api.Code `protobuf:"varint,2,opt,name=code,proto3,enum=api.Code" json:"code,omitempty" form:"code" query:"code"`
}

func (x *MenuSaveResp) Reset() {
	*x = MenuSaveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuSaveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuSaveResp) ProtoMessage() {}

func (x *MenuSaveResp) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuSaveResp.ProtoReflect.Descriptor instead.
func (*MenuSaveResp) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{4}
}

func (x *MenuSaveResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MenuSaveResp) GetCode() api.Code {
	if x != nil {
		return x.Code
	}
	return api.Code(0)
}

type MenuUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" form:"id" json:"id,omitempty"`
	StatusId int32  `protobuf:"varint,2,opt,name=StatusId,proto3" form:"status_id" json:"status_id,omitempty"`
	Sort     int32  `protobuf:"varint,3,opt,name=Sort,proto3" form:"sort" json:"sort,omitempty"`
	ParentId int64  `protobuf:"varint,4,opt,name=ParentId,proto3" form:"parent_id" json:"parent_id,omitempty"`
	MenuName string `protobuf:"bytes,5,opt,name=MenuName,proto3" form:"menu_name" json:"menu_name,omitempty"`
	MenuUrl  string `protobuf:"bytes,6,opt,name=MenuUrl,proto3" form:"menu_url" json:"menu_url,omitempty"`
	ApiUrl   string `protobuf:"bytes,7,opt,name=ApiUrl,proto3" form:"api_url" json:"api_url,omitempty"`
	MenuIcon string `protobuf:"bytes,8,opt,name=MenuIcon,proto3" form:"menu_icon" json:"menu_icon,omitempty"`
	Remark   string `protobuf:"bytes,9,opt,name=Remark,proto3" form:"remark" json:"remark,omitempty"`
	MenuType int32  `protobuf:"varint,10,opt,name=MenuType,proto3" form:"menu_type" json:"menu_type,omitempty"`
}

func (x *MenuUpdateReq) Reset() {
	*x = MenuUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuUpdateReq) ProtoMessage() {}

func (x *MenuUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuUpdateReq.ProtoReflect.Descriptor instead.
func (*MenuUpdateReq) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{5}
}

func (x *MenuUpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuUpdateReq) GetStatusId() int32 {
	if x != nil {
		return x.StatusId
	}
	return 0
}

func (x *MenuUpdateReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuUpdateReq) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MenuUpdateReq) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *MenuUpdateReq) GetMenuUrl() string {
	if x != nil {
		return x.MenuUrl
	}
	return ""
}

func (x *MenuUpdateReq) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *MenuUpdateReq) GetMenuIcon() string {
	if x != nil {
		return x.MenuIcon
	}
	return ""
}

func (x *MenuUpdateReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *MenuUpdateReq) GetMenuType() int32 {
	if x != nil {
		return x.MenuType
	}
	return 0
}

type MenuUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string   `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
	Code api.Code `protobuf:"varint,2,opt,name=code,proto3,enum=api.Code" json:"code,omitempty" form:"code" query:"code"`
}

func (x *MenuUpdateResp) Reset() {
	*x = MenuUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuUpdateResp) ProtoMessage() {}

func (x *MenuUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuUpdateResp.ProtoReflect.Descriptor instead.
func (*MenuUpdateResp) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{6}
}

func (x *MenuUpdateResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MenuUpdateResp) GetCode() api.Code {
	if x != nil {
		return x.Code
	}
	return api.Code(0)
}

type MenuDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" form:"ids" json:"ids,omitempty"`
}

func (x *MenuDeleteReq) Reset() {
	*x = MenuDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuDeleteReq) ProtoMessage() {}

func (x *MenuDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuDeleteReq.ProtoReflect.Descriptor instead.
func (*MenuDeleteReq) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{7}
}

func (x *MenuDeleteReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MenuDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string   `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty" form:"Msg" query:"Msg"`
	Code api.Code `protobuf:"varint,2,opt,name=code,proto3,enum=api.Code" json:"code,omitempty" form:"code" query:"code"`
}

func (x *MenuDeleteResp) Reset() {
	*x = MenuDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuDeleteResp) ProtoMessage() {}

func (x *MenuDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuDeleteResp.ProtoReflect.Descriptor instead.
func (*MenuDeleteResp) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{8}
}

func (x *MenuDeleteResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MenuDeleteResp) GetCode() api.Code {
	if x != nil {
		return x.Code
	}
	return api.Code(0)
}

var File_menu_proto protoreflect.FileDescriptor

var file_menu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65,
	0x6e, 0x75, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a,
	0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x67, 0x0a, 0x0c,
	0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1d,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65,
	0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc8, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x6e, 0x75, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x55,
	0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x55, 0x72, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xf3, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x29, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0d, 0xca, 0xbb, 0x18, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69,
	0x64, 0x52, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x53,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xca, 0xbb, 0x18, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0d, 0xca, 0xbb, 0x18,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca, 0xbb, 0x18, 0x09, 0x6d, 0x65, 0x6e, 0x75,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x75, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x75, 0x72, 0x6c, 0x52, 0x07,
	0x4d, 0x65, 0x6e, 0x75, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x55, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07, 0x61, 0x70, 0x69,
	0x5f, 0x75, 0x72, 0x6c, 0x52, 0x06, 0x41, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x12, 0x29, 0x0a, 0x08,
	0x4d, 0x65, 0x6e, 0x75, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d,
	0xca, 0xbb, 0x18, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x52, 0x08, 0x4d,
	0x65, 0x6e, 0x75, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x29, 0x0a, 0x08, 0x4d,
	0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0d, 0xca,
	0xbb, 0x18, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x08, 0x4d, 0x65,
	0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3f, 0x0a, 0x0c, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x8d, 0x03, 0x0a, 0x0d, 0x4d, 0x65, 0x6e, 0x75,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xca, 0xbb, 0x18, 0x02, 0x69, 0x64, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x0d, 0xca, 0xbb, 0x18, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x69, 0x64, 0x52, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x04,
	0x53, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xca, 0xbb, 0x18, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0d, 0xca, 0xbb,
	0x18, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x08, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca, 0xbb, 0x18, 0x09, 0x6d, 0x65, 0x6e,
	0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x75, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x75, 0x72, 0x6c, 0x52,
	0x07, 0x4d, 0x65, 0x6e, 0x75, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x55,
	0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07, 0x61, 0x70,
	0x69, 0x5f, 0x75, 0x72, 0x6c, 0x52, 0x06, 0x41, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x12, 0x29, 0x0a,
	0x08, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0d, 0xca, 0xbb, 0x18, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x52, 0x08,
	0x4d, 0x65, 0x6e, 0x75, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x29, 0x0a, 0x08,
	0x4d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0d,
	0xca, 0xbb, 0x18, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x08, 0x4d,
	0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x22, 0x41, 0x0a, 0x0e, 0x4d, 0x65, 0x6e, 0x75, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2a, 0x0a, 0x0d, 0x4d, 0x65,
	0x6e, 0x75, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x69, 0x64,
	0x73, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x41, 0x0a, 0x0e, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xb9, 0x02, 0x0a, 0x0b, 0x4d, 0x65,
	0x6e, 0x75, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x4d, 0x65, 0x6e,
	0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e,
	0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0xca, 0xc1,
	0x18, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x61, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x12, 0xd2, 0xc1, 0x18, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e,
	0x75, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d,
	0x65, 0x6e, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_proto_rawDescOnce sync.Once
	file_menu_proto_rawDescData = file_menu_proto_rawDesc
)

func file_menu_proto_rawDescGZIP() []byte {
	file_menu_proto_rawDescOnce.Do(func() {
		file_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_proto_rawDescData)
	})
	return file_menu_proto_rawDescData
}

var file_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_menu_proto_goTypes = []interface{}{
	(*MenuListReq)(nil),    // 0: menu.MenuListReq
	(*MenuListResp)(nil),   // 1: menu.MenuListResp
	(*MenuListData)(nil),   // 2: menu.MenuListData
	(*MenuSaveReq)(nil),    // 3: menu.MenuSaveReq
	(*MenuSaveResp)(nil),   // 4: menu.MenuSaveResp
	(*MenuUpdateReq)(nil),  // 5: menu.MenuUpdateReq
	(*MenuUpdateResp)(nil), // 6: menu.MenuUpdateResp
	(*MenuDeleteReq)(nil),  // 7: menu.MenuDeleteReq
	(*MenuDeleteResp)(nil), // 8: menu.MenuDeleteResp
	(api.Code)(0),          // 9: api.Code
}
var file_menu_proto_depIdxs = []int32{
	9, // 0: menu.MenuListResp.code:type_name -> api.Code
	2, // 1: menu.MenuListResp.data:type_name -> menu.MenuListData
	9, // 2: menu.MenuSaveResp.code:type_name -> api.Code
	9, // 3: menu.MenuUpdateResp.code:type_name -> api.Code
	9, // 4: menu.MenuDeleteResp.code:type_name -> api.Code
	0, // 5: menu.MenuHandler.MenuList:input_type -> menu.MenuListReq
	3, // 6: menu.MenuHandler.MenuSave:input_type -> menu.MenuSaveReq
	5, // 7: menu.MenuHandler.MenuUpdate:input_type -> menu.MenuUpdateReq
	7, // 8: menu.MenuHandler.MenuDelete:input_type -> menu.MenuDeleteReq
	1, // 9: menu.MenuHandler.MenuList:output_type -> menu.MenuListResp
	4, // 10: menu.MenuHandler.MenuSave:output_type -> menu.MenuSaveResp
	6, // 11: menu.MenuHandler.MenuUpdate:output_type -> menu.MenuUpdateResp
	8, // 12: menu.MenuHandler.MenuDelete:output_type -> menu.MenuDeleteResp
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_menu_proto_init() }
func file_menu_proto_init() {
	if File_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuListReq); i {
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
		file_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuListResp); i {
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
		file_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuListData); i {
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
		file_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuSaveReq); i {
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
		file_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuSaveResp); i {
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
		file_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuUpdateReq); i {
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
		file_menu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuUpdateResp); i {
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
		file_menu_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuDeleteReq); i {
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
		file_menu_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuDeleteResp); i {
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
			RawDescriptor: file_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_proto_goTypes,
		DependencyIndexes: file_menu_proto_depIdxs,
		MessageInfos:      file_menu_proto_msgTypes,
	}.Build()
	File_menu_proto = out.File
	file_menu_proto_rawDesc = nil
	file_menu_proto_goTypes = nil
	file_menu_proto_depIdxs = nil
}
