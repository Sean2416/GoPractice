package client

import (
	"context"
	"errors"
	"fmt"
	"shop-payment/conf"
	"shop-payment/shop-payment/sutil"
	"net/http"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/go-redis"
	"scm.tutorabc.com/tgo-framework/go-resty"
)

var redisFactory *redis.Factory

/**
初始化 redis
*/
func InitRedis(config redis.Config) error {
	if conf.Conf.Cachecloud.AppID > 0 {
		d, err := GetCacheCloud(conf.Conf.Cachecloud.AppID)
		if err != nil {
			return err
		}
		url := fmt.Sprintf("redis://%s", d.Standalone)//d.standalone
		config.URL = url
	}

	redisFactory = redis.NewFactory(&config)
	if redisFactory == nil {
		log.Errorf("InitRedis factory failed||err=%+v", config)
		return errors.New("InitRedis failed")
	}
	log.Debugf("InitRedis done,config=%+v", config)
	return nil
}

/**
获取 redis client
*/
func RedisClient(ctx context.Context) redis.Redis {
	return redisFactory.New(ctx)
}

var urlFmt = "http://cachecloud.tutorabc.com/cache/client/redis/standalone/%d.json?clientVersion=1.0"

type CacheCloud struct {
	Standalone string `json:"standalone"`
	Message    string `json:"message"`
	Status     int    `json:"status"`
}

func GetCacheCloud(appID int64) (data *CacheCloud, err error) {
	url := fmt.Sprintf(urlFmt, appID)
	c := resty.New()
	resp, err := c.R().Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("cache cloud error")
	}
	err = sutil.JsonUnmarshalFromString(resp.String(), &data)
	if err != nil {
		return
	}
	return
}
