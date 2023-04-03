package main //main為保留字元，build時會產生可執行檔(其餘不會)
import (
	"fmt"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

func main() {
	c := &config.AppConfig{
		AppID:         "Payment.PaymentExchange",
		Cluster:       "stg",
		IP:            "http://apollo.tutorabc.com/meta/prd",
		NamespaceName: "OMS.Share",
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println("初始化Apollo配置成功")

	//Use your apollo key to test
	cache := client.GetConfigCache(c.NamespaceName)
	value, _ := cache.Get("OmsAPI")
	fmt.Println(value)

}
