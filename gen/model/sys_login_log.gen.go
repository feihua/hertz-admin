// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLoginLog = "sys_login_log"

// SysLoginLog 系统访问记录
type SysLoginLog struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:访问ID" json:"id"`                      // 访问ID
	LoginName     string    `gorm:"column:login_name;not null;comment:登录账号" json:"login_name"`                           // 登录账号
	Ipaddr        string    `gorm:"column:ipaddr;not null;comment:登录IP地址" json:"ipaddr"`                                 // 登录IP地址
	LoginLocation string    `gorm:"column:login_location;not null;comment:登录地点" json:"login_location"`                   // 登录地点
	Platform      string    `gorm:"column:platform;not null;comment:平台信息" json:"platform"`                               // 平台信息
	Browser       string    `gorm:"column:browser;not null;comment:浏览器类型" json:"browser"`                                // 浏览器类型
	Version       string    `gorm:"column:version;not null;comment:浏览器版本" json:"version"`                                // 浏览器版本
	Os            string    `gorm:"column:os;not null;comment:操作系统" json:"os"`                                           // 操作系统
	Arch          string    `gorm:"column:arch;not null;comment:体系结构信息" json:"arch"`                                     // 体系结构信息
	Engine        string    `gorm:"column:engine;not null;comment:渲染引擎信息" json:"engine"`                                 // 渲染引擎信息
	EngineDetails string    `gorm:"column:engine_details;not null;comment:渲染引擎详细信息" json:"engine_details"`               // 渲染引擎详细信息
	Extra         string    `gorm:"column:extra;not null;comment:其他信息（可选）" json:"extra"`                                 // 其他信息（可选）
	Status        int32     `gorm:"column:status;not null;comment:登录状态(0:失败,1:成功)" json:"status"`                        // 登录状态(0:失败,1:成功)
	Msg           string    `gorm:"column:msg;not null;comment:提示消息" json:"msg"`                                         // 提示消息
	LoginTime     time.Time `gorm:"column:login_time;not null;default:CURRENT_TIMESTAMP;comment:访问时间" json:"login_time"` // 访问时间
}

// TableName SysLoginLog's table name
func (*SysLoginLog) TableName() string {
	return TableNameSysLoginLog
}
