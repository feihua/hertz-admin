package entity

import (
	"errors"
	"gorm.io/gorm"
	"time"
)
import "hertz_admin/biz/dal"

type SysUser struct {
	Id          int64     `json:"id"`           //主键
	GmtCreate   time.Time `json:"gmt_create"`   //创建时间
	GmtModified time.Time `json:"gmt_modified"` //修改时间
	StatusId    int32     `json:"status_id"`    //状态(1:正常，0:禁用)
	Sort        int32     `json:"sort"`         //排序
	UserNo      int64     `json:"user_no"`      //用户编号
	Mobile      string    `json:"mobile"`       //手机
	RealName    string    `json:"real_name"`    //真实姓名
	Remark      string    `json:"remark"`       //备注
	Password    string    `json:"password"`     //密码
}

func (model SysUser) TableName() string {
	return "sys_user"
}
func CreateSysUser(users []SysUser) error {
	return dal.DB.Create(users).Error
}

func DeleteSysUser(userId []int64) error {
	return dal.DB.Where("id in (?)", userId).Delete(&SysUser{}).Error
}

func UpdateSysUser(user SysUser) error {
	return dal.DB.Updates(user).Error
}

func QuerySysUser(keyword *string, page, pageSize int64) ([]SysUser, int64, error) {
	db := dal.DB.Model(SysUser{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(dal.DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []SysUser
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func QuerySysUserByMobile(mobile string) (*SysUser, error) {
	var res SysUser
	err := dal.DB.Where("mobile = ?", mobile).First(&res).Error

	switch err {
	case nil:
		return &res, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("用户不存在")
	default:
		return nil, err
	}
}
