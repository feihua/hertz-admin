// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysMenu = "sys_menu"

// SysMenu 菜单信息
type SysMenu struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                          // 主键
	MenuName   string     `gorm:"column:menu_name;not null;comment:菜单名称" json:"menu_name"`                               // 菜单名称
	MenuType   int32      `gorm:"column:menu_type;not null;default:1;comment:菜单类型(1：目录   2：菜单   3：按钮)" json:"menu_type"` // 菜单类型(1：目录   2：菜单   3：按钮)
	Visible    int32      `gorm:"column:visible;not null;default:1;comment:显示状态（0:隐藏, 显示:1）" json:"visible"`             // 显示状态（0:隐藏, 显示:1）
	Status     int32      `gorm:"column:status;not null;default:1;comment:菜单状态(1:正常，0:禁用)" json:"status"`                // 菜单状态(1:正常，0:禁用)
	Sort       int32      `gorm:"column:sort;not null;default:1;comment:排序" json:"sort"`                                 // 排序
	ParentID   int64      `gorm:"column:parent_id;not null;comment:父ID" json:"parent_id"`                                // 父ID
	MenuURL    string     `gorm:"column:menu_url;not null;comment:路由路径" json:"menu_url"`                                 // 路由路径
	APIURL     string     `gorm:"column:api_url;not null;comment:接口URL" json:"api_url"`                                  // 接口URL
	MenuIcon   string     `gorm:"column:menu_icon;not null;comment:菜单图标" json:"menu_icon"`                               // 菜单图标
	Remark     string     `gorm:"column:remark;not null;comment:备注" json:"remark"`                                       // 备注
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                                // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by;not null;comment:更新者" json:"update_by"`                                // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName SysMenu's table name
func (*SysMenu) TableName() string {
	return TableNameSysMenu
}
