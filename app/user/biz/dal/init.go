package dal

import (
	"GoMall/biz/dal/mysql"
	"GoMall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
