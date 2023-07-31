package options

import (
	"context"
	"github.com/qiniu/qmgo"
	"log"
)

type MongoOptions struct {
	Uri      string `json:"uri"`
	Database string `json:"database"`
	// ConnectTimeoutMS specifies a timeout that is used for creating connections to the server.
	//	If set to 0, no timeout will be used.
	//	The default is 30 seconds.
	ConnectTimeoutMS *int64 `json:"connect_timeout_ms"`
	// MaxPoolSize specifies that maximum number of connections allowed in the driver's connection pool to each server.
	// If this is 0, it will be set to math.MaxInt64,
	// The default is 100.
	MaxPoolSize *uint64 `json:"max_pool_size"`
	// MinPoolSize specifies the minimum number of connections allowed in the driver's connection pool to each server. If
	// this is non-zero, each server's pool will be maintained in the background to ensure that the size does not fall below
	// the minimum. This can also be set through the "minPoolSize" URI option (e.g. "minPoolSize=100"). The default is 0.
	MinPoolSize *uint64 `json:"min_pool_size"`
	// SocketTimeoutMS specifies how long the driver will wait for a socket read or write to return before returning a
	// network error. If this is 0 meaning no timeout is used and socket operations can block indefinitely.
	// The default is 300,000 ms.
	SocketTimeoutMS *int64 `json:"socket_timeout_ms"`
}

func NewMongoOptions(opts *MongoOptions) (*qmgo.Client, error) {
	mgo, err := qmgo.NewClient(context.Background(), &qmgo.Config{
		Uri:              opts.Uri,
		Database:         opts.Database,
		ConnectTimeoutMS: opts.ConnectTimeoutMS,
		MaxPoolSize:      opts.MinPoolSize,
		MinPoolSize:      opts.MinPoolSize,
		SocketTimeoutMS:  opts.SocketTimeoutMS,
	})
	if err != nil {
		log.Fatal("init mongo fail! err:%+v", err)
	}
	return mgo, nil
}
