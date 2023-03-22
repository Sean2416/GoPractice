package sutil

import "time"

//todo 默認拿+0時間
//有func轉TPE時間

const (
	TaipeiTimeZone = "Asia/Taipei"
)

func NowDateTimeStr() string {
	loc, _ := time.LoadLocation(TaipeiTimeZone)
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func NowDateTime() time.Time {
	loc, _ := time.LoadLocation(TaipeiTimeZone)
	return time.Now().In(loc)
}

func NowDateTimePoint() *time.Time {
	loc, _ := time.LoadLocation(TaipeiTimeZone)
	t := time.Now().In(loc)
	return &t
}

func TimeDay2Str(t time.Time) string {
	loc, _ := time.LoadLocation(TaipeiTimeZone)
	return t.In(loc).Format("2006-01-02")
}

func TimeStrByTs(ts int64) string {
	loc, _ := time.LoadLocation(TaipeiTimeZone)
	return time.Unix(ts, 0).In(loc).Format("2006-01-02 15:04:05")
}
func TimeMin2Str(t time.Time) string {
	loc, _ := time.LoadLocation(TaipeiTimeZone)
	return t.In(loc).Format("2006-01-02 15:04")
}

func TimeMin2StrOrinig(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func NowMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func NowNanosecond() int64 {
	return time.Now().UnixNano()
}

func CalcLatency(start time.Time) float64 {
	return time.Since(start).Seconds() * 1000
}
