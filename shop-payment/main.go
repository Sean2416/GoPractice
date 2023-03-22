package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"shop-payment/conf"
	"shop-payment/shop-payment/client"
	"shop-payment/shop-payment/sutil"
	"shop-payment/routers"
	"net/http"
	"os"
	"runtime"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/toml"
	"time"
)

func initHttpSvr(config conf.Config) *http.Server {
	//init gin framwork
	gin.SetMode(config.Server.RunMode)
	//InitRouter
	routersInit := routers.InitRouter(config.Middleware)
	readTimeout := config.Server.ReadTimeout
	writeTimeout := config.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Infof("start http server listening %s\n", endPoint)
	return server
}

func main() {
	// 命令行参数。
	var configPath string
	flag.StringVar(&configPath, "config", "conf/service.conf", "server config.")
	flag.Parse()

	// 解析配置。
	if _, err := toml.DecodeFile(configPath, &conf.Conf); err != nil {
		fmt.Printf("fail to read config.||err=%+v||config=%+v", err, configPath)
		os.Exit(1)
		return
	}
	config := conf.Conf
	fmt.Println("load config:" + sutil.JsonString(config))
	log.Init(&config.Log)
	defer log.Close()

	client.InitRedis(config.Redis)
	initHttpAPI(config.APIConf)

	client.Setup(config.DB)
	defer client.CloseDB()
	//統計NumGoroutine
	go monitorInfo()

	svr := initHttpSvr(config)
	svr.ListenAndServe()
}

func monitorInfo() {
	var timer *time.Ticker = time.NewTicker(2 * time.Minute)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			log.Infof("mainNumGoroutine=%d", runtime.NumGoroutine())
		}
	}
	return
}

func initHttpAPI(apiConf conf.APIConf) {
	log.Info("init webapi")
	client.InitRDAAPI(&apiConf.RDA)
	client.InitConsultantAPI(&apiConf.Consultant)
	client.InitAserverAPI(&apiConf.Aserver)
}
