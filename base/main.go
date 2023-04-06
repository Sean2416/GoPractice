package main //main為保留字元，build時會產生可執行檔(其餘不會)
import (
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
	namespaceConfig := client.GetConfigAndInit("application")

	fmt.Println(namespaceConfig)
}
