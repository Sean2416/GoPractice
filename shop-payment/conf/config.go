package conf

import (
	"fmt"
	"path/filepath"
	gin_middleware "scm.tutorabc.com/tgo-framework/gin-middleware"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/go-redis"
	"scm.tutorabc.com/tgo-framework/go-resty"
	"scm.tutorabc.com/tgo-framework/toml"
	"time"
)

// Config 表示一个 http 服务器的配置。
var (
	Conf Config
)

type ServerConfig struct {
	RunMode      string        `toml:"run_mode"`
	HttpPort     int           `toml:"http_port"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Schema      string
	TablePrefix string
	IdleConns   int    `toml:"idle_conns"`
	OpenConns   int    `toml:"open_conns"`
	FilePath    string `toml:"file_path"`
}

type SMS struct {
	AppSn string `toml:"appsn"`
}

type Cachecloud struct {
	AppID int64 `toml:"appid"`
}

type APIConf struct {
	RDA        resty.Config `toml:"rda"`
	Consultant resty.Config `toml:"consultant"`
	Aserver    resty.Config `toml:"aserver"`
}

// 对应 conf/service.conf 的结构。
type Config struct {
	Server     ServerConfig          `toml:"server"`
	Middleware gin_middleware.Config `toml:"middleware"`
	Log        log.Config            `toml:"log"`
	APIConf    APIConf               `toml:"apiconf"`
	DB         Database              `toml:"database"`
	Cachecloud Cachecloud            `toml:"cachecloud"`
	Redis      redis.Config          `toml:"redis"`
}

func LoadConfigFile(confFilePath string) (err error) {
	if confFilePath == "" {
		confFilePath = "../conf/controller.toml"
	}

	absConfFilePath, _ := filepath.Abs(confFilePath)
	fmt.Println("load config from " + absConfFilePath)
	_, err = toml.DecodeFile(confFilePath, &Conf)
	return err
}
