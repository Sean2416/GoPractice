package dao

import (
	"bufio"
	"fmt"
	"shop-payment/shop-payment/dao_model"
	"shop-payment/shop-payment/sutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestDao_GetSessionInfoByParams(t *testing.T) {
	t.Run("GetSessionInfoByParams", func(t *testing.T) {

		parmas := make(map[string]interface{})
		parmas["id"] = 16126

		dataList, err := getSessionInfoByParams(parmas)
		fmt.Println(err)
		fmt.Println(sutil.JsonString(dataList[0].SessionEndTime))

		//assert.Equal(t, err, nil)
		//assert.Equal(t, len(dataMap), 2)
	})
}

func TestDao_GetSessionInfoBySessionTime(t *testing.T) {
	t.Run("GetSessionInfoBySessionTime", func(t *testing.T) {

		dataList, err := GetSessionInfoAtTimeRange("2021-06-03 20:30", "2021-06-03 20:30")
		fmt.Println(err)
		for _, d := range dataList {
			fmt.Println(sutil.JsonString(d))
		}

		//assert.Equal(t, err, nil)
		//assert.Equal(t, len(dataMap), 2)
	})
}

func TestDao_GenSessioninfo(t *testing.T) {
	t.Run("GetSessionInfoBySessionTime", func(t *testing.T) {

		//d := &SessionInfo{
		//    Topic: "1234",
		//    Intro: "456",
		//}
		//err := AddSessionInfo(d)
		//fmt.Println(err)
		//fmt.Println(sutil.JsonString(d))

		dicArr, _ := readLines("/tmp/dicta")
		l := len(dicArr) - 1
		fmt.Println(GenerateRangeNum(0, l))

		for i := 0; i < 5; i++ {
			var topic, intro string

			//fmt.Println(GenerateRangeNum(0, l))

			for i := 0; i < 10; i++ {
				topic += dicArr[GenerateRangeNum(0, l)]
			}

			for i := 0; i < 66; i++ {
				intro += dicArr[GenerateRangeNum(0, l)]
			}

			d := &dao_model.SessionInfo{
				Topic: topic,
				Intro: intro,
			}
			AddSessionInfo(d)
		}

	})
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
