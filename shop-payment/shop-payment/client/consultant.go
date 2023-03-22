package client

import (
    "context"
    "scm.tutorabc.com/tgo-framework/go-resty"
)

var (
    consultantFactory *resty.Factory
)

func InitConsultantAPI(config *resty.Config) *resty.Factory {
    consultantFactory = resty.NewFactory(config)
    return consultantFactory
}

func ConAPI(ctx context.Context) *resty.Client {
    c := consultantFactory.NewTGOClient(ctx)
    c.SetHeader("Content-Type", "application/json")
    return c
}