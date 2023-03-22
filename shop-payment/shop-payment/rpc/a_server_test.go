package rpc

import (
	"context"
	"fmt"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/tgo-framework/go-trace"
	"testing"
)

func TestService_BBB(t *testing.T) {
	t.Run("CallA", func(t *testing.T) {
		traceInfo := trace.FromContext(context.Background())
		ctx := trace.NewContext(context.Background(), traceInfo)
		fmt.Println("test ctx", trace.FromContext(ctx))
		result ,err := CallAServer2(ctx)
		fmt.Println("test: ", err)
		fmt.Println("r:", sutil.JsonString(result))
	})
}
