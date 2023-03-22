package schedule

import (
	"context"
	"fmt"
	"shop-payment/shop-payment/dao"
	"shop-payment/shop-payment/rpc"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/go-trace"
	"time"
)

const (
	querySize  = 100
	contentFmt = "%s %s，快到上課時間囉！提醒您，開課前10分鐘，點擊以下連結進入教室。%s"
	courseUrl  = "https://www.gogoldtalk.com/course/%d"
	signature  = "TutorABC"

	cronTimer    = time.Minute * 10
	preClassTime = time.Minute * 60
)

// 這邊Demo 怎麼定期執行一個script 邏輯
// ＊＊ 需要注意多台機器 會重複執行
// ＊＊ 要控制同一時間只有一個process執行, 分佈式鎖 or 冪等處理

// 呼叫方式 go SchSendSmsAtSessionTime()
// 就可以丟到背景執行

func SchSendSmsAtSessionTime() {
	var timer = time.NewTicker(cronTimer)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			ctx := trace.NewContext(context.Background(), nil)
			log.Infof("SchSendSmsAtSessionTime||%v||start at %s", trace.FromContext(ctx), sutil.NowDateTimeStr())
			sessST := sutil.NowDateTime().Add(preClassTime)
			sessET := sutil.NowDateTime().Add(preClassTime).Add(cronTimer)
			dataList, err := dao.GetSessionInfoAtTimeRange(sutil.TimeMin2Str(sessST), sutil.TimeMin2Str(sessET))
			log.Infof("SchSendSmsAtSessionTime||%v||dataList %s", trace.FromContext(ctx), sutil.JsonString(dataList))
			if err != nil {
				log.Errorf("_com_sch_error||%v||SchSendSmsAtSessionTime error||err=%v", trace.FromContext(ctx), err)
				break
			}

		}
	}
	return
}

func SendSmsBySessionInfo(ctx context.Context, info *dao.SessionInfo) {
	//避免分布式, 重複執行
	dLock := rpc.NewDLock(fmt.Sprintf("run_sendSMS_%d", info.Id), time.Hour*24)
	dLock.SetLock(ctx)
}

