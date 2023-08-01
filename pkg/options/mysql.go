package options

import (
	"fmt"
	"log"
	"time"
	"xorm.io/xorm"
	log2 "xorm.io/xorm/log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlOptions struct {
	Host                  string        `json:"host"`
	Port                  int           `json:"port"`
	Driver                string        `json:"driver"`
	Username              string        `json:"username"`
	Password              string        `json:"password"`
	Database              string        `json:"database"`
	MaxIdleConnections    int           `json:"max_idle_connections"`
	MaxOpenConnections    int           `json:"max_open_connections"`
	MaxConnectionLifeTime time.Duration `json:"max_connection_life_time"`
	ShowSQL               bool          `json:"show_sql"`
	LogLevel              int           `json:"log_level"`
}

// NewMysqlOptions New create a new gorm db instance with the given options.
func NewMysqlOptions(opts *MysqlOptions) (*xorm.Engine, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
		true,
		"Local")

	engine, err := xorm.NewEngine(opts.Driver, dsn)
	if err != nil {
		log.Fatalf("init engine fail! err:%+v", err)
	}

	// 连接池配置
	engine.SetMaxOpenConns(opts.MaxOpenConnections)       // 最大 db 连接
	engine.SetMaxIdleConns(opts.MaxIdleConnections)       // 最大 db 连接空闲数
	engine.SetConnMaxLifetime(opts.MaxConnectionLifeTime) // 超过空闲数连接存活时间

	// 日志相关配置
	engine.ShowSQL(opts.ShowSQL)                     // 打印日志
	engine.SetLogLevel(log2.LogLevel(opts.LogLevel)) // 打印日志级别
	return engine, nil
}
