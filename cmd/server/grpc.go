package server

import (
	"context"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	grpc_middlewares "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/liang21/terminator/api/system/v1"
	"github.com/liang21/terminator/pkg/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func StartRPCService(bs config.Bootstrap) error {
	lis, err := net.Listen("tcp", bs.GRPC.Addr)
	if err != nil {
		log.Fatal("init listen grpc failed!", err)
		return nil
	}
	defer lis.Close()
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middlewares.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_validator.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(zap.NewNop()),
			)))
	// TODO: 注册服务
	v1.RegisterUserServiceServer(s, nil)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("start grpc service failed!", err)
		}
	}()
	mux := runtime.NewServeMux(
		// 返回错误结果处理
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			e := runtime.HTTPStatusError{
				HTTPStatus: 200,
				Err:        err,
			}
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &e)
		}),
		// 编码转换
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	rpcS := strings.Split(bs.GRPC.Addr, `:`)
	rpcPort := rpcS[1]
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// 注册服务
	if err := v1.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatal("register service failed!", err)
		return nil
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	pprof.Register(router)
	router.Use(GinLogger())
	return router.Run(bs.HTTP.Addr)

}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		log.Printf(
			"%d %s %s %s %s %s %s %s", c.Writer.Status(), c.Request.Method, path, query, c.ClientIP(), c.Request.UserAgent(), c.Errors.ByType(gin.ErrorTypePrivate).String(), cost,
		)
	}
}
