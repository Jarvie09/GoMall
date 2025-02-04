package main

import (
	"GoMall/demo/demo_thrift/biz/dal"
	"GoMall/demo/demo_thrift/biz/dal/mysql"
	"GoMall/demo/demo_thrift/biz/model"
	"fmt"
	"github.com/joho/godotenv"
)

// 测试数据库操作
func main() {
	//获取初始的连接，就需要环境变量，加载env
	err := godotenv.Load("./demo/demo_thrift/.env")
	if err != nil {
		panic(err)
	}
	dal.Init()
	//crud

	//增
	mysql.DB.Create(&model.User{Email: "2981306839@qq.com", Password: "123456"})
	//改
	mysql.DB.Model(&model.User{}).Where("email = ?", "2981306839@qq.com").Update("password", "11111111")

	//查
	//mysql.DB.First()   读取一行
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", "2981306839@qq.com").First(&row)
	fmt.Printf("row = %+v\n", row)
	//mysql.DB.Find()    读取多行

	//删
	//mysql.DB.Where("email = ?", "2981306839@qq.com").Delete(&model.User{})
	//强制删除
	mysql.DB.Unscoped().Where("email = ?", "2981306839@qq.com").Delete(&model.User{})

}
