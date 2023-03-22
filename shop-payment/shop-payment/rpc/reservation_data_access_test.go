package rpc

import (
    "context"
    "fmt"
	"shop-payment/shop-payment/sutil"
    "testing"
    "time"
)

func TestService_GetClassInformationByLobbySn(t *testing.T) {
    t.Run("A", func(t *testing.T) {
        ctx := context.Background()
        r , e := GetClassInformationByLobbySn(ctx,232020)
        fmt.Println("###")
        fmt.Println(e, sutil.JsonString(r))
        time.Sleep(10 * time.Second)
    })
}






