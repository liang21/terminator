package config

import (
	"github.com/liang21/terminator/pkg/options"
)

type Bootstrap struct {
	Server
	Data
}
type Server struct {
	GRPC Addr
	HTTP Addr
}
type Addr struct {
	Addr string
}
type Data struct {
	options.MysqlOptions
	options.RedisOptions
	options.MongoOptions
}
