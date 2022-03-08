// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"gohub/pkg/cache"
	"gohub/pkg/config"
)

// SetupCache 缓存
func SetupCache() {
	instance := initWith("default")
	cache.InitWithCacheStore(instance)
}

func initWith(cacheType string) cache.Store {
	var instance cache.Store
	switch cacheType {
	case "redis":
		instance = cache.NewRedisStore(
			fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
			config.GetString("redis.username"),
			config.GetString("redis.password"),
			config.GetInt("redis.database_cache"),
		)
	default:
		instance = cache.NewRistrettoStore()
	}
	return instance
}
