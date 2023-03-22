package client

import (
	"context"
	"fmt"
	"testing"
)

func Test_SetGet(t *testing.T) {
	t.Run("Test_SetGet", func(t *testing.T) {
		ctx := context.Background()
		err := RedisClient(ctx).Set("scm1", "hello scm")
		fmt.Println(err)

		// get
		v, err := RedisClient(ctx).Get("scm1")
		fmt.Printf("get=======v=%v,%T,err=%+v=====\n", v, v, err)
	})
}
