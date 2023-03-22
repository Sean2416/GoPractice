package rpc

import (
	"context"
	"fmt"
	rpcmodel2 "shop-payment/shop-payment/rpc_model"
	"shop-payment/shop-payment/sutil"
	"testing"
)

func TestService_GetConsultant(t *testing.T) {
	t.Run("ConMember", func(t *testing.T) {

		ctx := context.Background()
		req := &rpcmodel2.GetConsultantRequest{
			Consns: []int64{2913,5763,5833,9964,122668,5773,122670,6063,1331,500,122682,9917,122686,9951},
			Fields: []string{"basicFname", "basicLname", "basicCname", "nationality", "workFor", "status", "location", "conName"},
		}
		r , e := QueryConsultant(ctx,req)
		fmt.Println( e, sutil.JsonString(r))
	})
}
