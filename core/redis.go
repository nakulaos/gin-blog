package core

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gvb_server/global"
)

func Addr() string {
	return fmt.Sprintf("%s:%s", global.Config.Redis.Ip, global.Config.Redis.Port)
}
func InitRedis() {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     Addr(),
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
		PoolSize: global.Config.Redis.PoolSize, //连接池大小
	})
	global.Redis = Rdb
}
