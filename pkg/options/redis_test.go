package options

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestNewRedisOptions(t *testing.T) {
	opts := &RedisOptions{
		Addr:         "localhost:6379",
		Password:     "123456",
		Database:     0,
		DialTimeout:  1000000000,
		WriteTimeout: 400000000,
		ReadTimeout:  600000000,
	}
	rdb, err := NewRedisOptions(opts)
	if err != nil {
		log.Fatal("init redis fail! err:", err)
	}
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("ping redis fail! err:", err)
	}
	fmt.Println(rdb)
}
