// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sys_post.proto

package post

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

// 添加岗位信息表请求参数
type AddPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remark   string `protobuf:"bytes,1,opt,name=remark,proto3" form:"remark" json:"remark,omitempty"`                      //备注
	PostCode string `protobuf:"bytes,2,opt,name=post_code,json=postCode,proto3" form:"postCode" json:"postCode,omitempty"` //岗位编码
	PostName string `protobuf:"bytes,3,opt,name=post_name,json=postName,proto3" form:"postName" json:"postName,omitempty"` //岗位名称
	Sort     int32  `protobuf:"varint,4,opt,name=sort,proto3" form:"sort" json:"sort,omitempty"`                           //显示顺序
	Status   int32  `protobuf:"varint,5,opt,name=status,proto3" form:"status" json:"status,omitempty"`                     //岗位状态（0：停用，1:正常）
}

func (x *AddPostReq) Reset() {
	*x = AddPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostReq) ProtoMessage() {}

func (x *AddPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostReq.ProtoReflect.Descriptor instead.
func (*AddPostReq) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{0}
}

func (x *AddPostReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *AddPostReq) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *AddPostReq) GetPostName() string {
	if x != nil {
		return x.PostName
	}
	return ""
}

func (x *AddPostReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AddPostReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 删除岗位信息表请求参数
type DeletePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" form:"ids" json:"ids,omitempty"`
}

func (x *DeletePostReq) Reset() {
	*x = DeletePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostReq) ProtoMessage() {}

func (x *DeletePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostReq.ProtoReflect.Descriptor instead.
func (*DeletePostReq) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{1}
}

func (x *DeletePostReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 更新岗位信息表请求参数
type UpdatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty"`                                 //岗位id
	PostCode string `protobuf:"bytes,2,opt,name=post_code,json=postCode,proto3" form:"postCode" json:"postCode,omitempty"` //岗位编码
	PostName string `protobuf:"bytes,3,opt,name=post_name,json=postName,proto3" form:"postName" json:"postName,omitempty"` //岗位名称
	Sort     int32  `protobuf:"varint,4,opt,name=sort,proto3" form:"sort" json:"sort,omitempty"`                           //显示顺序
	Status   int32  `protobuf:"varint,5,opt,name=status,proto3" form:"status" json:"status,omitempty"`                     //岗位状态（0：停用，1:正常）
	Remark   string `protobuf:"bytes,6,opt,name=remark,proto3" form:"remark" json:"remark,omitempty"`                      //备注
}

func (x *UpdatePostReq) Reset() {
	*x = UpdatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostReq) ProtoMessage() {}

func (x *UpdatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostReq.ProtoReflect.Descriptor instead.
func (*UpdatePostReq) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePostReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePostReq) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *UpdatePostReq) GetPostName() string {
	if x != nil {
		return x.PostName
	}
	return ""
}

func (x *UpdatePostReq) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *UpdatePostReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdatePostReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

// 更新岗位信息表状态请求参数
type UpdatePostStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids    []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" form:"ids" json:"ids,omitempty"`   //岗位id
	Status int32   `protobuf:"varint,5,opt,name=status,proto3" form:"status" json:"status,omitempty"` //岗位状态（0：停用，1:正常）
}

func (x *UpdatePostStatusReq) Reset() {
	*x = UpdatePostStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostStatusReq) ProtoMessage() {}

func (x *UpdatePostStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostStatusReq.ProtoReflect.Descriptor instead.
func (*UpdatePostStatusReq) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePostStatusReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *UpdatePostStatusReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 查询岗位信息表详情请求参数
type QueryPostDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty"` //岗位id
}

func (x *QueryPostDetailReq) Reset() {
	*x = QueryPostDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPostDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPostDetailReq) ProtoMessage() {}

func (x *QueryPostDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPostDetailReq.ProtoReflect.Descriptor instead.
func (*QueryPostDetailReq) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{4}
}

func (x *QueryPostDetailReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryPostDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty" query:"id"`                                                     //岗位id
	PostCode   string `protobuf:"bytes,2,opt,name=post_code,json=postCode,proto3" form:"post_code" json:"post_code,omitempty" query:"post_code"`            //岗位编码
	PostName   string `protobuf:"bytes,3,opt,name=post_name,json=postName,proto3" form:"post_name" json:"post_name,omitempty" query:"post_name"`            //岗位名称
	Sort       int32  `protobuf:"varint,4,opt,name=sort,proto3" form:"sort" json:"sort,omitempty" query:"sort"`                                             //显示顺序
	Status     int32  `protobuf:"varint,5,opt,name=status,proto3" form:"status" json:"status,omitempty" query:"status"`                                     //岗位状态（0：停用，1:正常）
	Remark     string `protobuf:"bytes,6,opt,name=remark,proto3" form:"remark" json:"remark,omitempty" query:"remark"`                                      //备注
	CreateBy   string `protobuf:"bytes,7,opt,name=create_by,json=createBy,proto3" form:"create_by" json:"create_by,omitempty" query:"create_by"`            //创建者
	CreateTime string `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" form:"create_time" json:"create_time,omitempty" query:"create_time"`  //创建时间
	UpdateBy   string `protobuf:"bytes,9,opt,name=update_by,json=updateBy,proto3" form:"update_by" json:"update_by,omitempty" query:"update_by"`            //更新者
	UpdateTime string `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" form:"update_time" json:"update_time,omitempty" query:"update_time"` //更新时间
}

func (x *QueryPostDetailResp) Reset() {
	*x = QueryPostDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPostDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPostDetailResp) ProtoMessage() {}

func (x *QueryPostDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPostDetailResp.ProtoReflect.Descriptor instead.
func (*QueryPostDetailResp) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{5}
}

func (x *QueryPostDetailResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryPostDetailResp) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *QueryPostDetailResp) GetPostName() string {
	if x != nil {
		return x.PostName
	}
	return ""
}

func (x *QueryPostDetailResp) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *QueryPostDetailResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryPostDetailResp) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *QueryPostDetailResp) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *QueryPostDetailResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *QueryPostDetailResp) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *QueryPostDetailResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

// 分页查询岗位信息表列表请求参数
type QueryPostListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostCode string `protobuf:"bytes,2,opt,name=post_code,json=postCode,proto3" form:"postCode" json:"postCode,omitempty"`  //岗位编码
	PostName string `protobuf:"bytes,3,opt,name=post_name,json=postName,proto3" form:"postName" json:"postName,omitempty"`  //岗位名称
	Status   int32  `protobuf:"varint,5,opt,name=status,proto3" form:"status" json:"status,omitempty"`                      //岗位状态（0：停用，1:正常）
	PageNum  int32  `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" form:"pageNo" json:"pageNo,omitempty"`       //第几页
	PageSize int32  `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" form:"pageSize" json:"pageSize,omitempty"` //每页的数量
}

func (x *QueryPostListReq) Reset() {
	*x = QueryPostListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPostListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPostListReq) ProtoMessage() {}

func (x *QueryPostListReq) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPostListReq.ProtoReflect.Descriptor instead.
func (*QueryPostListReq) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{6}
}

func (x *QueryPostListReq) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *QueryPostListReq) GetPostName() string {
	if x != nil {
		return x.PostName
	}
	return ""
}

func (x *QueryPostListReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryPostListReq) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *QueryPostListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryPostListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" form:"id" json:"id,omitempty" query:"id"`                                                     //岗位id
	PostCode   string `protobuf:"bytes,2,opt,name=post_code,json=postCode,proto3" form:"post_code" json:"post_code,omitempty" query:"post_code"`            //岗位编码
	PostName   string `protobuf:"bytes,3,opt,name=post_name,json=postName,proto3" form:"post_name" json:"post_name,omitempty" query:"post_name"`            //岗位名称
	Sort       int32  `protobuf:"varint,4,opt,name=sort,proto3" form:"sort" json:"sort,omitempty" query:"sort"`                                             //显示顺序
	Status     int32  `protobuf:"varint,5,opt,name=status,proto3" form:"status" json:"status,omitempty" query:"status"`                                     //岗位状态（0：停用，1:正常）
	Remark     string `protobuf:"bytes,6,opt,name=remark,proto3" form:"remark" json:"remark,omitempty" query:"remark"`                                      //备注
	CreateBy   string `protobuf:"bytes,7,opt,name=create_by,json=createBy,proto3" form:"create_by" json:"create_by,omitempty" query:"create_by"`            //创建者
	CreateTime string `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" form:"create_time" json:"create_time,omitempty" query:"create_time"`  //创建时间
	UpdateBy   string `protobuf:"bytes,9,opt,name=update_by,json=updateBy,proto3" form:"update_by" json:"update_by,omitempty" query:"update_by"`            //更新者
	UpdateTime string `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" form:"update_time" json:"update_time,omitempty" query:"update_time"` //更新时间
}

func (x *QueryPostListResp) Reset() {
	*x = QueryPostListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPostListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPostListResp) ProtoMessage() {}

func (x *QueryPostListResp) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPostListResp.ProtoReflect.Descriptor instead.
func (*QueryPostListResp) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{7}
}

func (x *QueryPostListResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryPostListResp) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *QueryPostListResp) GetPostName() string {
	if x != nil {
		return x.PostName
	}
	return ""
}

func (x *QueryPostListResp) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *QueryPostListResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryPostListResp) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *QueryPostListResp) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *QueryPostListResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *QueryPostListResp) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *QueryPostListResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type PostEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostEmpty) Reset() {
	*x = PostEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEmpty) ProtoMessage() {}

func (x *PostEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_sys_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEmpty.ProtoReflect.Descriptor instead.
func (*PostEmpty) Descriptor() ([]byte, []int) {
	return file_sys_post_proto_rawDescGZIP(), []int{8}
}

var File_sys_post_proto protoreflect.FileDescriptor

var file_sys_post_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x79, 0x73, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x29, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x29, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xca, 0xbb, 0x18, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2a,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x07, 0xca, 0xbb,
	0x18, 0x03, 0x69, 0x64, 0x73, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xe3, 0x01, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xca, 0xbb, 0x18, 0x02, 0x69, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x70, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x29, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xca, 0xbb, 0x18, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0xbb,
	0x18, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x22, 0x54, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x69, 0x64, 0x73, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xca, 0xbb, 0x18, 0x02, 0x69, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x9f, 0x02, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f,
	0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x09, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xca, 0xbb, 0x18, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x6f, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x29, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0c, 0xca, 0xbb, 0x18, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x32, 0xd6, 0x04, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x61, 0x64, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0xd2, 0xc1,
	0x18, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x57, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0xd2, 0xc1, 0x18, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0xd2, 0xc1, 0x18, 0x21,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x70, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0xd2,
	0xc1, 0x18, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x68, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x22, 0xd2, 0xc1, 0x18, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x69, 0x68,
	0x75, 0x61, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_post_proto_rawDescOnce sync.Once
	file_sys_post_proto_rawDescData = file_sys_post_proto_rawDesc
)

func file_sys_post_proto_rawDescGZIP() []byte {
	file_sys_post_proto_rawDescOnce.Do(func() {
		file_sys_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_post_proto_rawDescData)
	})
	return file_sys_post_proto_rawDescData
}

var file_sys_post_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sys_post_proto_goTypes = []interface{}{
	(*AddPostReq)(nil),          // 0: system.AddPostReq
	(*DeletePostReq)(nil),       // 1: system.DeletePostReq
	(*UpdatePostReq)(nil),       // 2: system.UpdatePostReq
	(*UpdatePostStatusReq)(nil), // 3: system.UpdatePostStatusReq
	(*QueryPostDetailReq)(nil),  // 4: system.QueryPostDetailReq
	(*QueryPostDetailResp)(nil), // 5: system.QueryPostDetailResp
	(*QueryPostListReq)(nil),    // 6: system.QueryPostListReq
	(*QueryPostListResp)(nil),   // 7: system.QueryPostListResp
	(*PostEmpty)(nil),           // 8: system.PostEmpty
}
var file_sys_post_proto_depIdxs = []int32{
	0, // 0: system.PostHandler.AddPost:input_type -> system.AddPostReq
	1, // 1: system.PostHandler.DeletePost:input_type -> system.DeletePostReq
	2, // 2: system.PostHandler.UpdatePost:input_type -> system.UpdatePostReq
	3, // 3: system.PostHandler.UpdatePostStatus:input_type -> system.UpdatePostStatusReq
	4, // 4: system.PostHandler.QueryPostDetail:input_type -> system.QueryPostDetailReq
	6, // 5: system.PostHandler.QueryPostList:input_type -> system.QueryPostListReq
	8, // 6: system.PostHandler.AddPost:output_type -> system.PostEmpty
	8, // 7: system.PostHandler.DeletePost:output_type -> system.PostEmpty
	8, // 8: system.PostHandler.UpdatePost:output_type -> system.PostEmpty
	8, // 9: system.PostHandler.UpdatePostStatus:output_type -> system.PostEmpty
	5, // 10: system.PostHandler.QueryPostDetail:output_type -> system.QueryPostDetailResp
	7, // 11: system.PostHandler.QueryPostList:output_type -> system.QueryPostListResp
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sys_post_proto_init() }
func file_sys_post_proto_init() {
	if File_sys_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostReq); i {
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
		file_sys_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostReq); i {
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
		file_sys_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostReq); i {
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
		file_sys_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostStatusReq); i {
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
		file_sys_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPostDetailReq); i {
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
		file_sys_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPostDetailResp); i {
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
		file_sys_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPostListReq); i {
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
		file_sys_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPostListResp); i {
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
		file_sys_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEmpty); i {
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
			RawDescriptor: file_sys_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_post_proto_goTypes,
		DependencyIndexes: file_sys_post_proto_depIdxs,
		MessageInfos:      file_sys_post_proto_msgTypes,
	}.Build()
	File_sys_post_proto = out.File
	file_sys_post_proto_rawDesc = nil
	file_sys_post_proto_goTypes = nil
	file_sys_post_proto_depIdxs = nil
}
