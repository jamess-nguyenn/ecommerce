package lib

import (
	"ecommerce/config"
	"ecommerce/database/connection"
	"sync"
	"time"
)

type cache struct {
	rc     *connection.RedisClient
	prefix string
}

var (
	instance *cache
	once     sync.Once
)

func NewCache() *cache {
	once.Do(func() {
		rc, _ := connection.ConnectRedis()
		instance = &cache{
			rc:     rc,
			prefix: config.Cache.Prefix,
		}
	})

	return instance
}

func (c *cache) addPrefix(key string) string {
	return c.prefix + key
}

func (c *cache) Set(key string, value any, expiration time.Duration) error {
	return c.rc.Client.Set(c.addPrefix(key), value, expiration).Err()
}

func (c *cache) Get(key string) (string, error) {
	return c.rc.Client.Get(c.addPrefix(key)).Result()
}

func (c *cache) Forget(key string) error {
	return c.rc.Client.Del(c.addPrefix(key)).Err()
}
