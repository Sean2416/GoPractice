package dao

import (
	"context"
	"shop-payment/shop-payment/client"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/go-trace"
)


// sql類的原生 function
// 目前先實作 Query/Count


func Query(ctx context.Context, sql string, values []interface{}, list interface{}) (err error) {
	err = client.DB.SQL(sql, values...).Find(list)
	if err != nil {
		log.Errorf("_dao_query||%v||sql=%v||values=%v||err=%v", trace.ContextString(ctx), sql, values, err)
		return
	}
	log.Debugf("_dao_query||%v||sql=%v||values=%v||list=%v", trace.ContextString(ctx), sql, values, sutil.JsonString(list))
	return
}

// Counts 查询数量
func Counts(ctx context.Context, sql string, values []interface{}) (counts int64, err error) {
	counts, err = client.DB.SQL(sql, values...).Count()

	if err != nil {
		log.Errorf("_dao_counts||%v||sql=%v||counts=%v||err=%v", trace.ContextString(ctx), sql, values, err)
		return
	}

	log.Debugf("_dao_counts||%v||sql=%v||counts=%v||err=%v", trace.ContextString(ctx), sql, counts, err)
	return
}