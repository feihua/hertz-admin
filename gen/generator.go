// configuration.go
package main

import (
	"github.com/feihua/hertz-admin/biz/dal"
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

	//var dsn = "root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/rustdb?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dal.Dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("sys_user"),
		g.GenerateModel("sys_user_role"),
		g.GenerateModel("sys_role"),
		g.GenerateModel("sys_role_menu"),
		g.GenerateModel("sys_menu"),
	)

	// Execute the generator
	g.Execute()
}
