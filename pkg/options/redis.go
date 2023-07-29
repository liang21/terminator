package options

import (
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisOptions struct {
	Addr         string `json:"addr"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Database     int    `json:"database"`
	MaxIdle      int    `json:"max_idle"`
	MaxActive    int    `json:"max_active"`
	DialTimeout  int    `json:"dial_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	ReadTimeout  int    `json:"read_timeout"`
}

func NewRedisOptions(opts *RedisOptions) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         opts.Addr,
		Username:     opts.Username,
		Password:     opts.Password,
		DB:           opts.Database,
		PoolSize:     opts.MaxActive,
		MaxIdleConns: opts.MaxIdle,
		DialTimeout:  time.Duration(opts.DialTimeout),
		WriteTimeout: time.Duration(opts.WriteTimeout),
		ReadTimeout:  time.Duration(opts.ReadTimeout),
	})
	if err := rdb.Close(); err != nil {
		log.Printf("init redis fail! err:%+v", err)
	}
	return rdb, nil
}
