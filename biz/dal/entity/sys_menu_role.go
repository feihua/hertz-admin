package entity

import "time"
import "hertz_admin/biz/dal"

type SysMenuRole struct {
	Id          int64     `json:"id"`           //主键
	GmtCreate   time.Time `json:"gmt_create"`   //创建时间
	GmtModified time.Time `json:"gmt_modified"` //修改时间
	StatusId    int8      `json:"status_id"`    //状态(1:正常，0:禁用)
	Sort        int32     `json:"sort"`         //排序
	MenuId      int64     `json:"menu_id"`      //菜单ID
	RoleId      int64     `json:"role_id"`      //角色ID
}

func (model SysMenuRole) TableName() string {
	return "sys_menu_role"
}

func CreateSysMenuRole(users []SysMenuRole) error {
	return dal.DB.Create(users).Error
}

func DeleteSysMenuRole(userId int64) error {
	return dal.DB.Where("id = ?", userId).Delete(&SysMenuRole{}).Error
}

func UpdateSysMenuRole(user SysMenuRole) error {
	return dal.DB.Updates(user).Error
}

func QuerySysMenuRole(keyword *string, page, pageSize int64) ([]SysMenuRole, int64, error) {
	db := dal.DB.Model(SysMenuRole{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(dal.DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []SysMenuRole
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
