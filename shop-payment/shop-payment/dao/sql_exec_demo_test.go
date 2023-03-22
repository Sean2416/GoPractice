package dao

import (
	"context"
	"fmt"
	"shop-payment/shop-payment/sutil"
	"testing"
)




func TestDao_DemoModelEasyQuery(t *testing.T) {
	t.Run("TestDao_DemoModelEasyQuery", func(t *testing.T) {

		result , err := DemoModelEasyQuery(context.Background())
		fmt.Println(err)
		fmt.Println(sutil.JsonString(result))

	})
}

func TestDao_DemoModelCondQuery(t *testing.T) {
	t.Run("TestDao_DemoModelCondQuery", func(t *testing.T) {

		result , err := DemoModelCondQuery(context.Background())
		fmt.Println(err)
		fmt.Println(sutil.JsonString(result))

	})
}

func TestDao_ExecSqlDemo(t *testing.T) {
	t.Run("TestDao_ExecSqlDemo", func(t *testing.T) {
		result , err := DemoSqlQuery(context.Background())
		fmt.Println(err)
		fmt.Println(sutil.JsonString(result))

		fmt.Println("====")
		c , err2 := DemoSqlCount(context.Background())
		fmt.Println(err2)
		fmt.Println(c)

	})
}


