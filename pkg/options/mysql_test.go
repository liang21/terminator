package options

import (
	"fmt"
	"log"
	"testing"
)

func TestNewMysqlOptions(t *testing.T) {
	opts := &MysqlOptions{
		Host:                  "localhost",
		Port:                  3306,
		Driver:                "mysql",
		Username:              "root",
		Password:              "123456",
		Database:              "terminator",
		MaxIdleConnections:    10,
		MaxOpenConnections:    10,
		MaxConnectionLifeTime: 10,
		ShowSQL:               true,
		LogLevel:              1,
	}
	engine, err := NewMysqlOptions(opts)
	if err != nil {
		log.Fatal("init mysql fail! err:", err)
	}
	err = engine.Ping()
	if err != nil {
		log.Fatal("ping mysql fail! err:", err)
	}
	fmt.Println(engine)
}
