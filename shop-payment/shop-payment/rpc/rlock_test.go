package rpc

import (
	"context"
	"fmt"
	"testing"
	"time"
)


func Test_Lock(t *testing.T) {
	t.Run("Test_Lock", func(t *testing.T) {
		dLock := NewDLock("aaa", time.Second*100)

		ctx := context.Background()
		l, err := dLock.SetLock(ctx)
		fmt.Println("dLock:", l, err)
		fmt.Println(dLock.UnLock(ctx))

		fmt.Println(dLock.SetLock(ctx))

	})
}
