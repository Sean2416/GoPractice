package rpc

import (
	"context"
	"errors"
	"fmt"
	"shop-payment/shop-payment/client"
	rpcmodel2 "shop-payment/shop-payment/rpc_model"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/go-trace"
)

const (
	consultantQueryUri = "/consultant/profile/queryByFields"
)

func QueryConsultant(ctx context.Context, req *rpcmodel2.GetConsultantRequest) (data []*rpcmodel2.GetConsultantResponse, err error) {
	defer func() {
		if err != nil {
			log.Errorf("CreateConsultantError||%v||err %s", trace.FromContext(ctx), err.Error())
			return
		}
	}()

	resp, err := client.ConAPI(ctx).R().SetBody(sutil.JsonString(req)).Post(consultantQueryUri)
	if err != nil {
		return
	}
	log.Infof("QueryConsultant||%v||resp %s", trace.FromContext(ctx), resp.String())
	if !resp.IsSuccess() {
		err = errors.New(fmt.Sprintf("ApiUnSuccess, code:%d", resp.StatusCode()))
		return
	}
	apiResp := &rpcmodel2.ConsultantResponse{}
	err = sutil.JsonUnmarshalFromString(resp.String(), apiResp)
	if err != nil {
		return
	}
	if !apiResp.Success {
		err = errors.New("Api error" + apiResp.Message)
		return
	}
	err = sutil.JsonUnmarshalFromString(sutil.JsonString(apiResp.Data), &data)
	if err != nil {
		return
	}
	return
}
