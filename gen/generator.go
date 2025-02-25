// configuration.go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./gen/query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// var dsn = "root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/rustdb?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	var dsn = "root:123456@tcp(127.0.0.1:3306)/hertz-admin?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("sys_dept"),
		g.GenerateModelAs("sys_dict_data", "SysDictData"),
		g.GenerateModel("sys_dict_type"),
		g.GenerateModel("sys_login_log"),
		g.GenerateModel("sys_menu"),
		g.GenerateModel("sys_notice"),
		g.GenerateModel("sys_operate_log"),
		g.GenerateModel("sys_post"),
		g.GenerateModel("sys_role"),
		g.GenerateModel("sys_role_dept"),
		g.GenerateModel("sys_role_menu"),
		g.GenerateModel("sys_user"),
		g.GenerateModel("sys_user_post"),
		g.GenerateModel("sys_user_role"),
	)

	// Execute the generator
	g.Execute()
}
