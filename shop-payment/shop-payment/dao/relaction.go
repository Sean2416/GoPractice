package dao

import (
	"shop-payment/shop-payment/client"
	"shop-payment/shop-payment/dao_model"
)


//Find 返回多條數據
func GetRelactionSiiSet() (data []*dao_model.RelactionSiiSet, err error) {
	if err := client.DB.Where(
		"valid = ? and is_del = ?", true, false,
	).Find(&data); err != nil {
		return nil, err
	}
	return
}

//Get 默認返回 Limit 1
func GetRelactionSiiSetByID(id int64) (data *dao_model.RelactionSiiSet, err error) {
	data = &dao_model.RelactionSiiSet{}
	if _, err := client.DB.Where(
		"id = ?", id,
	).Get(&data); err != nil {
		return nil, err
	}
	return
}

func AddRelactionTag(data *dao_model.RelactionTag) (err error) {
	if _, err = client.DB.Insert(data); err != nil {
		return
	}
	return
}

func AddRelactionSliSet(data *dao_model.RelactionSiiSet) (err error) {
	if _, err = client.DB.Insert(data); err != nil {
		return
	}
	return
}

func GetRelactionSiiSetBySessionInfoID(id int64) (data []*dao_model.RelactionSiiSession, err error) {
	if err = client.DB.Where(
		"session_info_id = ? and valid = ? and is_del = ?", id, true, false,
	).Find(&data); err != nil {
		return nil, err
	}
	return
}

func AddRelactionSession(data *dao_model.RelactionSiiSession) (err error) {
	if _, err = client.DB.Insert(data); err != nil {
		return err
	}
	return
}

func GetRelactionStyleMap() (data map[int]*dao_model.RelactionStyle, err error) {
	var rows []*dao_model.RelactionStyle
	data = make(map[int]*dao_model.RelactionStyle)
	if err = client.DB.Where(
		"valid = ? and is_del = ?", true, false,
	).Find(&rows); err != nil {
		return nil, err
	}
	for _, d := range rows {
		data[d.Id] = d
	}
	return
}

func GetRelactionItemMap() (data map[int]*dao_model.RelactionItem, err error) {
	var rows []*dao_model.RelactionItem
	data = make(map[int]*dao_model.RelactionItem)
	if err = client.DB.Where(
		"valid = ? and is_del = ?", true, false,
	).Find(&rows); err != nil {
		return nil, err
	}
	for _, d := range rows {
		data[d.Id] = d
	}
	return
}

func GetRelactionTagMap() (data map[int]dao_model.RelactionTag, err error) {
	var rows []*dao_model.RelactionTag
	data = make(map[int]dao_model.RelactionTag)
	if err = client.DB.Where(
		"valid = ? and is_del = ?", true, false,
	).Find(&rows); err != nil {
		return nil, err
	}
	for _, d := range rows {
		data[d.Id] = *d
	}
	return
}
