package lib

import (
	"ecommerce/config"
	"ecommerce/database/connection"
	"sync"
	"time"
)

type Cache struct {
	rc     *connection.RedisClient
	prefix string
}

var (
	instance *Cache
	once     sync.Once
)

func NewCache() *Cache {
	once.Do(func() {
		rc, _ := connection.ConnectRedis()
		instance = &Cache{
			rc:     rc,
			prefix: config.Cache.Prefix,
		}
	})

	return instance
}

func (c *Cache) addPrefix(key string) string {
	return c.prefix + key
}

func (c *Cache) Set(key string, value any, expiration time.Duration) error {
	return c.rc.Client.Set(c.addPrefix(key), value, expiration).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.rc.Client.Get(c.addPrefix(key)).Result()
}

func (c *Cache) Forget(key string) error {
	return c.rc.Client.Del(c.addPrefix(key)).Err()
}
