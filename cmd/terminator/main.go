package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	grpc_middlewares "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	system_v1 "github.com/liang21/terminator/api/system/v1"
	test_v1 "github.com/liang21/terminator/api/test/v1"
	systemBiz "github.com/liang21/terminator/internal/system-service/biz"
	systemRepo "github.com/liang21/terminator/internal/system-service/repo"
	systemService "github.com/liang21/terminator/internal/system-service/service"
	testBiz "github.com/liang21/terminator/internal/test-service/biz"
	testRepo "github.com/liang21/terminator/internal/test-service/repo"
	testService "github.com/liang21/terminator/internal/test-service/service"
	"github.com/liang21/terminator/pkg/file"
	"github.com/liang21/terminator/pkg/log"
	"github.com/liang21/terminator/pkg/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

var (
	confPath = flag.String("conf", "./configs/config.yml", "config path, eg: -conf ./configs")
)

func main() {
	flag.Parse()
	// logger配置
	logOpts := &log.Options{
		Level:            "debug",
		Format:           "console",
		EnableColor:      true,
		DisableCaller:    true,
		OutputPaths:      []string{"D:\\terminator\\go\\terminator\\logs\\terminator.log", "stdout"},
		ErrorOutputPaths: []string{"D:\\terminator\\go\\terminator\\logs\\error.log"},
	}
	// 初始化全局logger
	log.Init(logOpts)
	defer log.Flush()
	bs, err := file.LoadFile(*confPath)
	if err != nil {
		log.Fatalf("load config failed!", err)
	}
	rdb, err := options.NewRedisOptions(&bs.Data.Redis)
	if err != nil {
		log.Fatalf("init redis failed!", err)
	}
	engine, err := options.NewMysqlOptions(&bs.Data.Mysql)
	if err != nil {
		log.Fatalf("init mysql failed!", err)
	}

	// grpc client
	cc, err := grpc.Dial(bs.GRPC.Addr, grpc.WithInsecure())

	lis, err := net.Listen("tcp", bs.GRPC.Addr)
	if err != nil {
		log.Fatalf("init listen grpc failed!", err)
	}
	defer lis.Close()
	s := grpc.NewServer(
		// add grpc interceptor
		grpc.UnaryInterceptor(
			grpc_middlewares.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_validator.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(zap.NewNop()),
			)))
	// TODO: 注册服务
	// user service
	userRepo := systemRepo.NewUserRepo(engine, rdb)
	userUsecase := systemBiz.NewUserUsecase(userRepo)
	userService := systemService.NewUserService(userUsecase)
	system_v1.RegisterUserServiceServer(s, userService)
	// product service
	productRepo := systemRepo.NewProductRepo(engine, rdb)
	productUsecase := systemBiz.NewProductUsecase(productRepo)
	productService := systemService.NewProductService(productUsecase)
	system_v1.RegisterProductServiceServer(s, productService)
	// project service
	projectRepo := systemRepo.NewProjectRepo(engine, rdb)
	projectUsecase := systemBiz.NewProjectUsecase(projectRepo, productRepo)
	projectService := systemService.NewProjectService(projectUsecase)
	system_v1.RegisterProjectServiceServer(s, projectService)
	//module service
	moduleRepo := testRepo.NewTestModuleRepo(engine, rdb)
	moduleUsecase := testBiz.NewTestModuleUsecase(moduleRepo)
	moduleService := testService.NewTestModuleService(cc, moduleUsecase)
	test_v1.RegisterTestModuleServiceServer(s, moduleService)
	// 启动grpc服务
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("start grpc service failed!", err)
		}
	}()
	mux := runtime.NewServeMux(
		// 返回错误结果处理
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaller runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			e := runtime.HTTPStatusError{HTTPStatus: 200, Err: err}
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaller, writer, request, &e)
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
	if err := system_v1.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	if err := system_v1.RegisterProductServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	if err := system_v1.RegisterProjectServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	if err := test_v1.RegisterTestModuleServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	pprof.Register(router)
	router.Use(GinLogger())
	router.Group("/api").Any("/*{gateway}", gin.WrapH(mux))
	if err = router.Run(bs.HTTP.Addr); err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s %v", debug.Stack(), err)
		}
	}()
}
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		log.Debugf(path,
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"cost", cost,
			"user-agent", c.Request.UserAgent(),
			"error", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
