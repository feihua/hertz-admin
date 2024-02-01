package dal

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hertz_admin/gen/query"
	"time"
)

var Dsn = "root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/rustdb?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		panic(err)
	}
	query.SetDefault(DB)
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	hlog.Debugf(format, args)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
