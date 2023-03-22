package dao

import (
	"fmt"
	"shop-payment/shop-payment/dao_model"
	"shop-payment/shop-payment/sutil"
	"testing"
)



func TestDao_GetRelactionSiiSet(t *testing.T) {
	t.Run("GetRelactionSiiSet", func(t *testing.T) {


		//a ,err := GetRelactionSiiSet()
		//fmt.Println(sutil.JsonString(a))
		//fmt.Println(err)

		a ,err := GetRelactionSiiSetByID(11)
		fmt.Println(sutil.JsonString(a))
		fmt.Println(err)




	})
}

func TestDao_AddRelactionTag(t *testing.T) {
	t.Run("AddRelactionTag", func(t *testing.T) {
		//
		//b := &RelactionTag{
		//	Content: "jintest",
		//}
		//
		//err := AddRelactionTag(b)
		//fmt.Println(sutil.JsonString(b))
		//fmt.Println(err)

		b := &dao_model.RelactionSiiSet{
			Style: 1,
			Item: 1,
			Tag: 2,
		}

		err := AddRelactionSliSet(b)
		fmt.Println(sutil.JsonString(b))
		fmt.Println(err)



	})
}


