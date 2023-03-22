package dao

import (
	"shop-payment/shop-payment/client"
	"shop-payment/shop-payment/dao_model"
)

func GetSessionInfoAtTimeRange(start string, end string) (data []*dao_model.SessionInfo, err error) {
	if err = client.DB.Where(
		"session_start_time >= ? and session_start_time < ? and channel_id = 1 and valid = 1", start, end,
	).Find(&data); err != nil {
		return nil, err
	}
	return
}

func AddSessionInfo(data *dao_model.SessionInfo) (int64,  error) {
	return client.DB.Insert(data)
}

func getSessionInfoByParams(params map[string]interface{}) (data []*dao_model.SessionInfo, err error) {
	if err = client.DB.Where(params).Find(&data); err != nil {
		return nil, err
	}
	return
}
