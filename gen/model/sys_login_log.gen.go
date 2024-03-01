// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLoginLog = "sys_login_log"

// SysLoginLog 系统登录日志
type SysLoginLog struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                                                       // 编号
	UserName  string    `gorm:"column:user_name;not null;comment:用户名" json:"user_name"`                                                             // 用户名
	Status    string    `gorm:"column:status;not null;comment:登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）" json:"status"` // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
	IP        string    `gorm:"column:ip;not null;comment:IP地址" json:"ip"`                                                                          // IP地址
	LoginTime time.Time `gorm:"column:login_time;not null;default:CURRENT_TIMESTAMP;comment:登陆时间" json:"login_time"`                                // 登陆时间
}

// TableName SysLoginLog's table name
func (*SysLoginLog) TableName() string {
	return TableNameSysLoginLog
}
