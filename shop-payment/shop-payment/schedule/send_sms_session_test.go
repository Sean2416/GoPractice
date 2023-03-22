package schedule

import (
	"fmt"
	"shop-payment/shop-payment/sutil"
	"testing"
)



func TestService_T(t *testing.T) {
	t.Run("T", func(t *testing.T) {
		contentFmt := "%s %s，快到上課時間囉！提醒您，開課前3分鐘，點擊以下連結進入教室。%s"
		courseUrl  := "https://www.gogoldtalk.com/course/%d"

		smsConent := fmt.Sprintf(contentFmt,
			sutil.TimeMin2Str(sutil.NowDateTime()),
			"info.Topic",
			fmt.Sprintf(courseUrl, 123),
		)
		fmt.Println(smsConent)
	})
}



