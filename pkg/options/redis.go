package options

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisOptions struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	Database     int    `yaml:"database"`
	DialTimeout  int    `yaml:"dial_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
}

func NewRedisOptions(opts *RedisOptions) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         opts.Addr,
		Password:     opts.Password,
		DB:           opts.Database,
		DialTimeout:  time.Duration(opts.DialTimeout),
		WriteTimeout: time.Duration(opts.WriteTimeout),
		ReadTimeout:  time.Duration(opts.ReadTimeout),
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("ping redis fail! err:", err)
	}
	//defer func() {
	//	if err = rdb.Close(); err != nil {
	//		panic(err)
	//	}
	//}()
	return rdb, nil
}
