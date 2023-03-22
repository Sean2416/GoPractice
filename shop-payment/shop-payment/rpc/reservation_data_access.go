package rpc

import (
	"context"
	"errors"
	"shop-payment/shop-payment/client"
	rpcmodel2 "shop-payment/shop-payment/rpc_model"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/tgo-framework/go-log"
)

const (
	GetClassInformationByLobbySnUri = "/ReservationDataAccess/Class/GetClassInformationByLobbySn"
)

func GetClassInformationByLobbySn(ctx context.Context, lobbySn int64) (data []*rpcmodel2.ClassInformationByLobbySn, err error) {
	req := &rpcmodel2.ClassInfoByLobbySnReq{
		Data: &rpcmodel2.ClassInfoByLobbySnData{
			Lobbysns:         []int64{lobbySn},
			HaveStrategyid:   true,
			IsonlyStrategyid: true,
		},
	}

	resp, err := client.RDAAPI(ctx).R().SetBody(sutil.JsonString(req)).Post(GetClassInformationByLobbySnUri)
	apiResp := &rpcmodel2.ReservationDataAccessResponse{}
	log.Infof("GetClassInformationByLobbySn||resp %s", resp.String())
	err = sutil.JsonUnmarshalFromString(resp.String(), apiResp)
	if err != nil {
		return
	}
	if !apiResp.Success {
		err = errors.New("Api error" + apiResp.Garbage)
		return
	}

	err = sutil.JsonUnmarshalFromString(sutil.JsonString(apiResp.Data), &data)
	if err != nil {
		return
	}
	return
}
