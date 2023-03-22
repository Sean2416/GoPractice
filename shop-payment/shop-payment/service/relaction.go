package service

import (
	"context"
	"fmt"

	"shop-payment/shop-payment/dao"
	"shop-payment/shop-payment/dao_model"
	"shop-payment/shop-payment/rpc"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/jimzhong/grpc-idl-templete/api"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/go-trace"
)

type RelactionSiiSetListDescp struct {
	Descp map[int][]string
}

// service層.... 做Business Access Layer....
// 看是要call rpc or dao... 搞一搞再返回
// 順便展示下 log 用法
func CallAServer(ctx context.Context) (result *api.UserResponseData, err error){
	lb := sutil.NewLogBuild(ctx, "CallAServer")
	defer func() {
		if lb.IsError() {
			log.Error(lb)
			return
		}
		log.Info(lb)
	}()

	lb.Set("A", 1) //set一個東西, 會先cache起來, 在log.Info(lb)才打印
	lb.Set("B", 2) //裡面是一個map, Key不能重複
	result, err = rpc.CallAServer(ctx)
	if err != nil{
		lb.SetError(err) //如果有error 就返回, log.Error(lb)會記錄下來
		return
	}

	lb.SetDescp("安安你好") //如果有想要紀錄啥的.. 可以用Descp, 在log.Info(lb)才打印
	return
}


//微職人的某個邏輯... 正好留下當範例
//這邊log是留單行的
func RelactionSiiSetList(ctx context.Context) (res *RelactionSiiSetListDescp, err error) {
	res = &RelactionSiiSetListDescp{}
	res.Descp = make(map[int][]string)

	var relConf []*dao_model.RelactionSiiSet
	relConf, err = dao.GetRelactionSiiSet()
	if err != nil {
		return
	}

	var styleMap map[int]*dao_model.RelactionStyle
	styleMap, err = dao.GetRelactionStyleMap()
	log.Infof("styleMap||%v||data: %s", trace.FromContext(ctx), sutil.JsonString(styleMap))

	var itemMap map[int]*dao_model.RelactionItem
	itemMap, err = dao.GetRelactionItemMap()
	log.Infof("itemMap||%v||data: %s", trace.FromContext(ctx), sutil.JsonString(itemMap))

	var TagMap map[int]dao_model.RelactionTag
	TagMap, err = dao.GetRelactionTagMap()

	for _, rel := range relConf {
		s := rel.Style
		i := rel.Item
		t := rel.Tag

		var sDscp, iDescp, tDescp string
		if val, ok := styleMap[s]; ok {
			sDscp = fmt.Sprintf("%s.%d", val.Title, s)
		}
		if val, ok := itemMap[i]; ok {
			iDescp = fmt.Sprintf("%s.%d", val.Title, i)
		}
		if val, ok := TagMap[t]; ok {
			tDescp =fmt.Sprintf("%s.%d", val.Content, t)
		}

		if len(sDscp) == 0 || len(iDescp) == 0 || len(tDescp) == 0 {
			continue
		}
		res.Descp[rel.Id] = []string{sDscp, iDescp, tDescp}
	}
	return
}

