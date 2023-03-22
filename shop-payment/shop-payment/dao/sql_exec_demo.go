package dao

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"shop-payment/shop-payment/dao_model"
	"shop-payment/shop-payment/sutil"
	"time"
)

//若有一些場景使用orm不方便, 可以使用sql獲取數據
//以下為Query / Count Demo

// ******** sqlbuilder ********
// github上面有很多 sqlbuilder 工具, xorm原生也有... 但星星很少
// 詳情: https://www.kancloud.cn/xormplus/xorm/212929

func DemoModelEasyQuery(ctx context.Context) (result []*dao_model.MemberInfo, err error) {
	memCond := &dao_model.MemberInfoCond{
		ID:           sutil.Int64Ptr(647),
		IDs:          []int64{1, 2, 647},
		CreatedEnd: sutil.TimePtr(time.Now()),
	}

	m := dao_model.MemberInfo{}
	sql, args, err := m.FindEasyToSQL(memCond, "id", "email")

	// select id, email from member_info where id = 647 and id in (1,2,647) and Now() < created_at

	if err != nil {
		return
	}
	err = Query(ctx, sql, args, &result)
	return
}

func DemoModelCondQuery(ctx context.Context) (result []*dao_model.MemberInfo, err error) {
	memCond := &dao_model.MemberInfoCond{
		ID:           sutil.Int64Ptr(647),
		IDs:          []int64{1, 2, 647},
		CreatedEnd: sutil.TimePtr(time.Now()),
	}
	// select * from member_info where id = 647 and id in (1, 2, 647) and created < now()

	m := dao_model.MemberInfo{}
	page := dao_model.NewDefPage()
	order := dao_model.OrderBy{"id": "desc", "client_sn": "asc"}
	groupby := []string{"phone_code"}

	sql, args, err := m.FindToSQL(memCond, page, order, groupby)
	if err != nil {
		return
	}
	err = Query(ctx, sql, args, &result)
	return
}

func DemoSqlQuery(ctx context.Context) (result []*dao_model.MemberInfo, err error) {
	sb := sqlbuilder.NewSelectBuilder()
	mInfo := new(dao_model.MemberInfo)
	sb.Select("id", "chn_name", "client_sn").From(mInfo.TableName()).Where(sb.GreaterThan("client_sn", 51017480))
	sql, args := sb.Build()

	err = Query(ctx, sql, args, &result)
	return
}

func DemoSqlCount(ctx context.Context) (counts int64, err error) {
	sb := sqlbuilder.NewSelectBuilder()
	mInfo := new(dao_model.MemberInfo)
	sb.Select("count(1)").From(mInfo.TableName()).Where(sb.GreaterThan("client_sn", 51017480))
	sql, args := sb.Build()
	return Counts(ctx, sql, args)
}
