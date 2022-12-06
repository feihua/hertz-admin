package entity

import "time"
import "hertz_admin/biz/dal"

type SysMenu struct {
	Id          int64     `json:"id"`           //主键
	GmtCreate   time.Time `json:"gmt_create"`   //创建时间
	GmtModified time.Time `json:"gmt_modified"` //修改时间
	StatusId    int32     `json:"status_id"`    //状态(1:正常，0:禁用)
	Sort        int32     `json:"sort"`         //排序
	ParentId    int64     `json:"parent_id"`    //父ID
	MenuName    string    `json:"menu_name"`    //菜单名称
	MenuUrl     string    `json:"menu_url"`     //路由路径
	ApiUrl      string    `json:"api_url"`      //接口URL
	MenuIcon    string    `json:"menu_icon"`    //菜单图标
	Remark      string    `json:"remark"`       //备注
	MenuType    int32     `json:"menu_type"`    //菜单类型(1：目录   2：菜单   3：按钮)
}

func (model SysMenu) TableName() string {
	return "sys_menu"
}

func CreateSysMenu(users []SysMenu) error {
	return dal.DB.Create(users).Error
}

func DeleteSysMenu(menuIds []int64) error {
	return dal.DB.Where("id in (?)", menuIds).Delete(&SysMenu{}).Error
}

func UpdateSysMenu(user SysMenu) error {
	return dal.DB.Updates(user).Error
}

func QuerySysMenu(keyword *string, page, pageSize int64) ([]SysMenu, int64, error) {
	db := dal.DB.Model(SysMenu{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(dal.DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []SysMenu
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
