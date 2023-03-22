package dao

import (
	"shop-payment/shop-payment/client"
	"shop-payment/shop-payment/dao_model"
	"shop-payment/shop-payment/sutil"
)

type MemberType int64

const (
	MemberType_NONE     MemberType = 0
	MemberType_USER     MemberType = 1
	MemberType_LECTURER MemberType = 2
	MemberType_FROMCON  MemberType = 3
)


func AddMemberInfo(data *dao_model.MemberInfo) (err error) {
	//now := sutil.NowDateTime()
	//data.ModifiedAt = now
	if _, err = client.DB.Insert(data); err != nil {
		return
	}
	return
}

func UpdateMemberInfo(id int64, data *dao_model.MemberInfo) (count int64, err error) {
	data.CreatedAt = sutil.NowDateTime()
	if count, err = client.DB.Where("id = ?", id).Update(data); err != nil {
		return
	}
	return
}

//Get 默認返回 Limit 1
func GetMemberInfoByID(id int64) (data *dao_model.MemberInfo, err error) {
	data = &dao_model.MemberInfo{}
	if _, err := client.DB.Where(
		"id = ?", id,
	).Get(data); err != nil {
		return nil, err
	}
	return
}

//Find 返回多條數據, 分頁範例
func GetMemberInfo(page *dao_model.Page) (data []*dao_model.MemberInfo, err error) {
	if err := client.DB.Limit(page.Limit, page.Offset).Find(&data); err != nil {
		return nil, err
	}
	return
}

func getMemberInfoByParams(params map[string]interface{}) (data []*dao_model.MemberInfo, err error) {
	if err = client.DB.Where(params).Find(&data); err != nil {
		return nil, err
	}
	return
}

