package server

import (
	"fmt"
	"github.com/liang21/terminator/pkg/config"
	"runtime/debug"
)

func StartApiServer(bs config.Bootstrap) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s %v", debug.Stack(), err)
		}
	}()
	if err := StartRPCService(bs); err != nil {
		panic(err)
	}
	return nil
}
