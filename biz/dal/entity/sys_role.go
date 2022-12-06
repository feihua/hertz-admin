package entity

import "time"
import "hertz_admin/biz/dal"

type SysRole struct {
	Id          int64     `json:"id"`           //主键
	GmtCreate   time.Time `json:"gmt_create"`   //创建时间
	GmtModified time.Time `json:"gmt_modified"` //修改时间
	StatusId    int32     `json:"status_id"`    //状态(1:正常，0:禁用)
	Sort        int32     `json:"sort"`         //排序
	RoleName    string    `json:"role_name"`    //名称
	Remark      string    `json:"remark"`       //备注
}

func (model SysRole) TableName() string {
	return "sys_role"
}
func CreateSysRole(users []SysRole) error {
	return dal.DB.Create(users).Error
}

func DeleteSysRole(roleIds []int64) error {
	return dal.DB.Where("id in (?)", roleIds).Delete(&SysRole{}).Error
}

func UpdateSysRole(user SysRole) error {
	return dal.DB.Updates(user).Error
}

func QuerySysRole(keyword *string, page, pageSize int64) ([]SysRole, int64, error) {
	db := dal.DB.Model(SysRole{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(dal.DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []SysRole
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
