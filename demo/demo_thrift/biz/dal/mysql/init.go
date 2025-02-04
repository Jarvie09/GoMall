package mysql

import (
	"GoMall/demo/demo_thrift/biz/model"
	"GoMall/demo/demo_thrift/conf"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DSN := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	type Version struct {
		Version string
	}
	var v Version
	err = DB.Raw("select version() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	fmt.Println("mysql version:", v.Version)
}
