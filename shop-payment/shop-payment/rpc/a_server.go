package rpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"shop-payment/shop-payment/client"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/jimzhong/grpc-idl-templete/api"
	"scm.tutorabc.com/tgo-framework/go-trace"
	"time"
)

func CallAServer(ctx context.Context) (result *api.UserResponseData, err error) {
	result = &api.UserResponseData{
		LastUpdated: timestamppb.New(time.Now()),
	}

	//
	//a := api.UserResponse{
	//	Header: &include.ResponseApiMeta{
	//		Code:    sutil.Int32Ptr(int32(constants.ErrNo_NotExist)),
	//		Success: sutil.BoolPtr(false),
	//	},
	//	Data: result,
	//}

	//fmt.Println(sutil.JsonString(a))
	//orig := &api.UserResponse{}
	//sutil.JsonUnmarshalFromString(sutil.JsonString(a), orig)
	//fmt.Println(sutil.JsonString(orig))

	return
}

func CallAServer2(ctx context.Context) (result *api.UserResponseData, err error) {
	userResponse := &api.UserResponse{}

	resp, err := client.AserverAPI(ctx).R().Get("/v1/callA")
	if err != nil {
		return
	}
	fmt.Println(err, trace.FromContext(ctx))
	fmt.Println(resp.String())
	fmt.Println("跑 unit tset 測試看看...")

	err = sutil.JsonUnmarshalFromString(resp.String(), userResponse)
	if err != nil {
		return
	}
	if userResponse.GetApiMeta().GetSuccess() == false || userResponse.GetApiMeta().GetCode() < 0 {
		err = errors.New("api error: " + userResponse.GetApiMeta().GetMsg())
		return
	}

	return userResponse.GetData(), nil
}
