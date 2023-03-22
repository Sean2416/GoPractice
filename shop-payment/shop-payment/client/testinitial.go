package client

import (
	"shop-payment/conf"
	"scm.tutorabc.com/tgo-framework/go-log"
)

func InitClients(config conf.Config) {
	log.Init(&config.Log)
	InitRedis(config.Redis)

	apiConf := config.APIConf
	InitRDAAPI(&apiConf.RDA)
	InitConsultantAPI(&apiConf.Consultant)
	InitAserverAPI(&apiConf.Aserver)
	Setup(config.DB)
}
