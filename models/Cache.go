package models

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"time"
)

var redisClient cache.Cache
var enableRedis, _ = beego.AppConfig.Bool("enableRedis")
var redisTime, _ = beego.AppConfig.Int("redisTime")

func init() {
	if enableRedis {
		config := map[string]string{
			"key":      beego.AppConfig.String("redisKey"),
			"conn":     beego.AppConfig.String("redisConn"),
			"dbNum":    beego.AppConfig.String("redisDbNum"),
			"password": beego.AppConfig.String("redisPwd"),
		}
		bytes, _ := json.Marshal(config)
		redisClient, err = cache.NewCache("redis", string(bytes))
		if err != nil {
			logs.Error("Redis 连接失败")
		} else {
			logs.Error("Redis 连接成功")
		}
	}
}

type cacheDb struct{}

var CacheDb = &cacheDb{}

// 写入缓存
func (c *cacheDb) Set(key string, value interface{}) {
	if enableRedis {
		bytes, _ := json.Marshal(value)
		redisClient.Put(key, string(bytes), time.Second*time.Duration(redisTime))
	}
}

// 读取缓存
func (c cacheDb) Get(key string, obj interface{}) bool {
	if enableRedis {
		if redisStr := redisClient.Get(key); redisStr != nil {
			redisValue, ok := redisStr.([]uint8)
			if !ok {
				logs.Error("从 Redis 获取数据失败")
				return false
			}
			json.Unmarshal([]byte(redisValue), obj)
			return true
		}
		return false
	}
	return false
}
