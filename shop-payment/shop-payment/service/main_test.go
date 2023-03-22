package service

import (
	"shop-payment/shop-payment/client"

	"shop-payment/conf"
	"testing"
)

func TestMain(m *testing.M) {
	conf.LoadConfigFile("../../conf/service.conf")
	client.InitClients(conf.Conf)
	m.Run()
}
