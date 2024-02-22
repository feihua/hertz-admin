package dal

/*
Author: LiuFeiHua
Date: 2024/2/2 下午5:08
*/
import (
	"errors"
	"hertz_admin/gen/model"
)

// QueryUserList 查询用户列表
func QueryUserList(userName, mobile string, page, pageSize int64) ([]model.SysUser, int64, error) {
	db := DB.Model(&model.SysUser{})
	if len(userName) != 0 {
		db = db.Where("role_name like ?", "%"+userName+"%")
	}
	if len(mobile) != 0 {
		db = db.Where("mobile like ?", "%"+mobile+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []model.SysUser
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

// QueryUserMenuList 查询用户菜单列表
func QueryUserMenuList(userId int64) ([]*model.SysMenu, error) {

	sql := `select u.*
from sys_user_role t
         left join sys_role usr on t.role_id = usr.id
         left join sys_role_menu srm on usr.id = srm.role_id
         left join sys_menu u on srm.menu_id = u.id
where t.user_id = ?`
	var res []*model.SysMenu
	DB.Raw(sql, userId).Scan(&res)

	if res == nil {
		return res, errors.New("用户还没有分配权限")
	}

	return res, nil
}
