package dal

import (
	"github.com/feihua/hertz-admin/biz/model/log"
	"github.com/feihua/hertz-admin/gen/model"
)

// QueryLoginLogList 查询登录日志列表
func QueryLoginLogList(userName, ip string, page, pageSize int64) ([]model.SysLoginLog, int64, error) {
	db := DB.Model(&model.SysLoginLog{})
	if len(userName) != 0 {
		db = db.Where("user_name like ?", "%"+userName+"%")
	}

	if len(ip) != 0 {
		db = db.Where("ip like ?", "%"+ip+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []model.SysLoginLog
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

// QueryStatisticsLoginLog 查询登录日志统计信息
func QueryStatisticsLoginLog() *log.StatisticsLoginLogData {
	resp := new(log.StatisticsLoginLogData)
	sql := `select count(distinct ip) current_day_login_count from sys_login_log where date(login_time) = curdate()`
	DB.Raw(sql).Scan(&resp.DayLoginCount)

	sql = `SELECT COUNT(DISTINCT ip) AS current_week_login_count FROM sys_login_log WHERE YEARWEEK(login_time, 1) = YEARWEEK(CURDATE(), 1)`
	DB.Raw(sql).Scan(&resp.WeekLoginCount)

	sql = `SELECT COUNT(DISTINCT ip) AS current_month_login_count FROM sys_login_log WHERE MONTH(login_time) = MONTH(CURDATE())   AND YEAR(login_time) = YEAR(CURDATE())`
	DB.Raw(sql).Scan(&resp.MonthLoginCount)

	return resp
}

// QueryOperateLogList 查询操作日志列表
func QueryOperateLogList(userName, ip, method string, page, pageSize int64) ([]model.SysOperateLog, int64, error) {
	db := DB.Model(&model.SysOperateLog{})
	if len(userName) != 0 {
		db = db.Where("user_name like ?", "%"+userName+"%")
	}

	if len(ip) != 0 {
		db = db.Where("ip like ?", "%"+ip+"%")
	}

	if len(method) != 0 {
		db = db.Where("method like ?", "%"+method+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []model.SysOperateLog
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
