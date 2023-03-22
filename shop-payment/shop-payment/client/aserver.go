package client

import (
    "context"
    "scm.tutorabc.com/tgo-framework/go-resty"
)

var (
    aserverFactory *resty.Factory
)

func InitAserverAPI(config *resty.Config) *resty.Factory {
    aserverFactory = resty.NewFactory(config)
    return aserverFactory
}

func AserverAPI(ctx context.Context) *resty.Client {
    c := aserverFactory.NewTGOClient(ctx)
    c.SetHeader("Content-Type", "application/json")
    return c
}

