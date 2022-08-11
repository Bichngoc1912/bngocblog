package main

import (
	"fmt"
	"log"
	"os"

	// "blog/pkg/app/config"
	"blog/pkg/app/config"
	blgutils "blog/utils"

	"github.com/antigloss/go/logger"
	"github.com/fasthttp/router"
	"github.com/linuxpham/fasthttpsession"
	"github.com/linuxpham/fasthttpsession/redis"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

func HelloWorld(ctx *fasthttp.RequestCtx) {
	fmt.Println("helloworld ne")
}

func initRouter() *router.Router{
	router := router.New()

	router.GET("/hello", HelloWorld)

	return router
}

func initSession(sessionStore *fasthttpsession.Session, serverConfig *config.ServerSetting) error {
	//init session store
	err := sessionStore.SetProvider(serverConfig.SessionServer.Type, &redis.Config{
		Host: serverConfig.SessionServer.StorageServer.Host,
		Port: serverConfig.SessionServer.StorageServer.Port,
		MaxIdle:     8,
		IdleTimeout: 300,
		Password:    serverConfig.SessionServer.StorageServer.Password,
		KeyPrefix:   serverConfig.SessionServer.StorageServer.Prefix,
	})

	if err != nil {
		return errors.Errorf(fmt.Sprintf("Session store [%v] with error: %v", serverConfig.SessionServer.Type, err.Error()))
	}

	return nil
}

func main() {
	var rootPath string
	rootPathDefault, err := os.Getwd()
	if err != nil {
		log.Fatalf("%vCannot get rootpath default", err.Error())
		os.Exit(1)
	}

	var configDefautlPath = fmt.Sprintf("%s%s%s", rootPath, "/configs", "/development.json")
	var configPath =rootPathDefault + configDefautlPath
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	serverConfig, err := blgutils.LoadConfig(configPath)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	
	var sessionConfig = fasthttpsession.NewDefaultConfig()
	var sessionStore = fasthttpsession.NewSession(sessionConfig)

	//init session
	err = initSession(sessionStore, &serverConfig)
	if err != nil {
		if serverConfig.DebugMode {
			log.Println(fmt.Sprintf("Session store [%v] with error: %v", serverConfig.SessionServer.Type, err.Error()))
		} else {
			logger.Error(err.Error())
		}
		os.Exit(1)
	}

	// init router
	appRouterInit := initRouter()
	appHandleFunc := func(ctx *fasthttp.RequestCtx) {
		appRouterInit.Handler(ctx)
	}

	if err := fasthttp.ListenAndServe(serverConfig.HttpServerSetting.Addr, appHandleFunc); err != nil {
		log.Fatal("err...", err)
	}

	fmt.Println("Server starting in port: ", serverConfig.HttpServerSetting.Addr)
}