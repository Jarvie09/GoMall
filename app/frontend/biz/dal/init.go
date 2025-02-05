package dal

import (
	"GoMall/app/frontend/biz/dal/mysql"
	"GoMall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
