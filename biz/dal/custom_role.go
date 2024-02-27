package dal

/*
Author: LiuFeiHua
Date: 2024/2/2 下午5:07
*/
import (
	"github.com/feihua/hertz_admin/gen/model"
)

// QueryRoleList 查询角色列表
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
