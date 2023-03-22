package client

import (
	"context"
	"scm.tutorabc.com/tgo-framework/go-resty"
)

var (
	rdaFactory *resty.Factory
)

func InitRDAAPI(config *resty.Config) *resty.Factory {
	rdaFactory = resty.NewFactory(config)
	return rdaFactory
}

func RDAAPI(ctx context.Context) *resty.Client {
	c := rdaFactory.NewTGOClient(ctx)
	c.SetHeader("Token", c.Token)
	c.SetHeader("Content-Type", "application/json")
	return c
}