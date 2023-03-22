package service

import (
	"context"
	"errors"
	"shop-payment/shop-payment/rpc"
	rpcmodel2 "shop-payment/shop-payment/rpc_model"
)

var (
	ErrorNotFoundConsultant = errors.New("NotFound Consultant")
)

func GetConsultant(ctx context.Context, conSn int64) (data *rpcmodel2.GetConsultantResponse, err error) {
	var req = &rpcmodel2.GetConsultantRequest{
		Consns: []int64{conSn},
		Fields: []string{"basicFname", "basicLname", "basicCname", "nationality", "workFor", "status", "location", "conName"},
	}
	var conInfos []*rpcmodel2.GetConsultantResponse
	conInfos, err = rpc.QueryConsultant(ctx, req)
	if err != nil {
		err = errors.New("get member error" + err.Error())
		return
	}

	if len(conInfos) == 0 {
		err = ErrorNotFoundConsultant
		return
	}
	return conInfos[0], nil

}
