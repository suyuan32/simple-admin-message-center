package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-message-center/ent"
	"github.com/suyuan32/simple-admin-message-center/internal/config"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/smssdk"
	"net/smtp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config           config.Config
	DB               *ent.Client
	Redis            *redis.Client
	EmailAuth        *smtp.Auth
	SmsGroup         *smssdk.SmsGroup
	EmailClientGroup map[string]*smtp.Client
	EmailAddrGroup   map[string]string
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Error), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		//ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config:           c,
		DB:               db,
		Redis:            c.RedisConf.MustNewRedis(),
		SmsGroup:         &smssdk.SmsGroup{},
		EmailAddrGroup:   map[string]string{},
		EmailClientGroup: map[string]*smtp.Client{},
	}
}
