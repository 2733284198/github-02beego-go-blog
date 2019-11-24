package wechat

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/sirupsen/logrus"
)

var redisCache = cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379"})

var config = &wechat.Config{
	AppID:          "xxxxxxxxxxxxxxxxxxxxxx",
	AppSecret:      "xxxxxxxxxxxxxxxxxxxxxx",
	Token:          "xxxxxxxxxxxxxxxxxxxxxx",
	EncodingAESKey: "xxxxxxxxxxxxxxxxxxxxxx",
	Cache:			redisCache,
}

var log = logrus.New()
