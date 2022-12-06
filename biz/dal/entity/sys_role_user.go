package entity

import "time"
import "hertz_admin/biz/dal"

type SysRoleUser struct {
	Id            int64     `json:"id"`           //主键
	GmtCreate     time.Time `json:"gmt_create"`   //创建时间
	GmtModified   time.Time `json:"gmt_modified"` //修改时间
	StatusId      int8      `json:"status_id"`    //状态(1:正常，0:禁用)
	Sort          int32     `json:"sort"`         //排序
	RoleId        int64     `json:"role_id"`      //角色ID
	SysRoleUserId int64     `json:"user_id"`      //用户ID
}

func (model SysRoleUser) TableName() string {
	return "sys_role_user"
}
func CreateSysRoleUser(users []SysRoleUser) error {
	return dal.DB.Create(users).Error
}

func DeleteSysRoleUser(userId int64) error {
	return dal.DB.Where("id = ?", userId).Delete(&SysRoleUser{}).Error
}

func UpdateSysRoleUser(user SysRoleUser) error {
	return dal.DB.Updates(user).Error
}

func QuerySysRoleUser(keyword *string, page, pageSize int64) ([]SysRoleUser, int64, error) {
	db := dal.DB.Model(SysRoleUser{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(dal.DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []SysRoleUser
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
