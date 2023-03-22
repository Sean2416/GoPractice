package client

import (
	"shop-payment/conf"
	"testing"
)

func TestMain(m *testing.M) {
	conf.LoadConfigFile("../../conf/service.conf")
	InitClients(conf.Conf)
	m.Run()
}
