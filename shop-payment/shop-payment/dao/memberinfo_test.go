package dao

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"shop-payment/shop-payment/dao_model"
	"shop-payment/shop-payment/sutil"
	"testing"
)

func TestDao_UpdateMemberInfo(t *testing.T) {
	t.Run("GetMemberOAByConSn", func(t *testing.T) {

		data := &dao_model.MemberInfo{
			EnFname: "12134",
			ChnName: "34353",
		}
		c, err := UpdateMemberInfo(668, data)
		fmt.Println(c)
		fmt.Println(err)
		fmt.Println(fmt.Println(sutil.JsonString(data)))
		fmt.Println("#####")
		fmt.Println(data.GetAddressCity())
		fmt.Println(data.GetEnFname())

	})
}

func TestDao_getMemberInfoByParams(t *testing.T) {
	t.Run("GetMemberOAByConSn", func(t *testing.T) {

		data := make(map[string]interface{})
		data["id"] = 668

		c, err := getMemberInfoByParams(data)
		fmt.Println(c)
		fmt.Println(err)

	})
}


func TestDao_AddMemberInfo(t *testing.T) {
	t.Run("AddMemberInfo", func(t *testing.T) {

		data := &dao_model.MemberInfo{
			Email:     "1111111",
			EnFname:   "123",
			Type:      3,
			PhoneCode: "886",
			Phone:     "12345",
			ClientSn:  12345,
			NickName:  "886",
		}

		err := AddMemberInfo(data)
		fmt.Println(err)
		fmt.Println(sutil.JsonString(data))

	})
}
func TestDao_GetMemberInfoByID(t *testing.T) {
	t.Run("GetMemberInfoByID", func(t *testing.T) {

		data , err := GetMemberInfoByID(1900)
		fmt.Println(err)
		fmt.Println(sutil.JsonString(data))

	})
}

func TestDao_GetMemberInfo(t *testing.T) {
	t.Run("GetMemberInfo", func(t *testing.T) {

		page := dao_model.PageSet{
			No: 3,
			Size: 100,
		}

		data , err := GetMemberInfo(dao_model.NewPage(&page))
		fmt.Println(err)
		fmt.Println(sutil.JsonString(data))

	})
}


func TestDao_SQLBUILDER(t *testing.T) {
	t.Run("GetMemberInfo", func(t *testing.T) {
		sqlbuilder.DefaultFieldMapper = sqlbuilder.SnakeCaseMapper
		var userStruct = sqlbuilder.NewStruct(new(dao_model.MemberInfo))
		fmt.Println(userStruct.Columns())
	})
}


