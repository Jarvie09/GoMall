package dal

import (
	"GoMall/demo/demo_thrift/biz/dal/mysql"
	//"GoMall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
