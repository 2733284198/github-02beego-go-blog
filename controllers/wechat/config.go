package wechat

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/sirupsen/logrus"
)

var redisCache = cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379"})

var config = &wechat.Config{
	AppID:          "wx115fb14a07bfd278",
	AppSecret:      "3d2fa6055b93a1bb196d271724a35df6",
	Token:          "b8cf671eaa1a270a9b53ddb894dd9029",
	EncodingAESKey: "5K3QpL541yBjGTyKG8l6nvMlyX7eFzbM7eJLRFXkKm3",
	Cache:			redisCache,
}

var log = logrus.New()