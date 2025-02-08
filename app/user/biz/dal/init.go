package dal

import (
	"GoMall/app/user/biz/dal/mysql"
	"GoMall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
