package dal

import (
	"hertz_admin/gen/model"
)

func QueryRoleList(roleName string, page, pageSize int64) ([]model.SysRole, int64, error) {
	db := DB.Model(&model.SysRole{})
	if len(roleName) != 0 {
		db = db.Where("role_name like ?", "%"+roleName+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []model.SysRole
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
