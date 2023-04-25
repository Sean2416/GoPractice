package main //main為保留字元，build時會產生可執行檔(其餘不會)
import (
	"encoding/json"
	"fmt"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

func main() {
	c := &config.AppConfig{
		AppID:         "Chinese.Dictionary",
		Cluster:       "stg",
		IP:            "http://apollo.tutorabc.com/meta/stg",
		NamespaceName: "application",
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println(client)

	//Use your apollo key to test
	cache := client.GetConfigCache(c.NamespaceName)
	value, _ := cache.Get("dic")

	jsonStr, ok := value.(string)

	if !ok {
		fmt.Println("error: data is not a string")
		return
	}
	data := make(map[string]map[string]string)

	// 使用Unmarshal函數將json轉換成map
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}

	for key, value := range data {
		fmt.Println("Key:", key)
		fmt.Println("value:", value["Zhuyin"])
	}
}
