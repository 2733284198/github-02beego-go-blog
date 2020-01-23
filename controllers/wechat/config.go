package wechat

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/sirupsen/logrus"
)

var redisCache = cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379"})
var token = "b8cf671eaa1a270a9b53ddb894dd9029"
var config = &wechat.Config{
	AppID:          "wx0e1f405eb29b0f5b",
	AppSecret:      "6ca99c66a90afbc6c6d536243083c7ff",
	Token:          token,
	EncodingAESKey: "gF0kTc4xNXTYESNxCeRcusV7Vnze52H4AIaC3RIy19u",
	Cache:			redisCache,
}

var log = logrus.New()
