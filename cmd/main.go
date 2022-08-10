package main

import (
	"fmt"
	"log"
	"os"

	blgutils "blog/utils"

	"github.com/fasthttp/router"
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

func main() {
	appRouterInit := initRouter()
	appHandleFunc := func(ctx *fasthttp.RequestCtx) {
		appRouterInit.Handler(ctx)
	}

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
	
	if err := fasthttp.ListenAndServe(serverConfig.HttpServerSetting.Addr, appHandleFunc); err != nil {
		log.Fatal("err...", err)
	}

	fmt.Println("Server starting in port: ", serverConfig.HttpServerSetting.Addr)
}