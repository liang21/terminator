package options

import (
	"context"
	"github.com/qiniu/qmgo"
	"log"
)

type MongoOptions struct {
	Uri              string  `yaml:"uri"`
	Database         string  `yaml:"database"`
	ConnectTimeoutMS *int64  `yaml:"connect_timeout_ms"`
	MaxPoolSize      *uint64 `yaml:"max_pool_size"`
	MinPoolSize      *uint64 `yaml:"min_pool_size"`
	SocketTimeoutMS  *int64  `yaml:"socket_timeout_ms"`
}

func NewMongoOptions(opts *MongoOptions) (*qmgo.Client, error) {
	ctx := context.Background()
	cli, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri:              opts.Uri,
		Database:         opts.Database,
		ConnectTimeoutMS: opts.ConnectTimeoutMS,
		MaxPoolSize:      opts.MinPoolSize,
		MinPoolSize:      opts.MinPoolSize,
		SocketTimeoutMS:  opts.SocketTimeoutMS,
	})
	if err != nil {
		log.Fatal("init mongo fail! err:", err)
	}
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	return cli, nil
}
