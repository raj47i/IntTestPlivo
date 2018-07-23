package config

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

var cachePool *redis.Pool

// Initialized global redis connection pool
func initCache() {
	log.Debug("initializing redis cache connection pool ..")
	dsn := fmt.Sprintf("%s:%d", cfg.CacheHost, cfg.CachePort)
	cachePool = &redis.Pool{
		MaxIdle:   5,
		MaxActive: 20, // max number of connections
		Dial: func() (redis.Conn, error) {
			log.Debug("open new redis-db connection")
			c, err := redis.Dial("tcp", dsn)
			if err != nil {
				log.Fatalf("error connecting to redis: %v", err)
			}
			log.Debugf("selecting redis-db: %d", cfg.CacheDb)
			c.Do("SELECT", cfg.CacheDb)
			return c, err
		},
	}
}

// GetCache returns the redis connection
func GetCache() redis.Conn {
	return cachePool.Get()
}

// FlushCache flushes the current redisDb
func FlushCache() error {
	cache := GetCache()
	_, err := cache.Do("FLUSHDB")
	log.Warn("redis-db cleard")
	return err
}
