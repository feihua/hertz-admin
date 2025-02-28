package common

import (
	"context"
	"github.com/feihua/hertz-admin/gen/query"
)

// IsAdmin 判断是不是超级管理员
func IsAdmin(ctx context.Context, userId int64) (bool, error) {
	q := query.SysUserRole
	count, err := q.WithContext(ctx).Where(q.RoleID.Eq(1), q.UserID.Eq(userId)).Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
