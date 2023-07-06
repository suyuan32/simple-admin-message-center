package svc

import (
	"github.com/suyuan32/simple-admin-message-center/ent"
	"github.com/suyuan32/simple-admin-message-center/internal/config"
	"github.com/suyuan32/simple-admin-message-center/internal/enum/smsprovider"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/smssdk"
	"net/smtp"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config    config.Config
	DB        *ent.Client
	Redis     *redis.Redis
	EmailAuth *smtp.Auth
	SmsGroup  *smssdk.SmsGroup
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	smsGroup := &smssdk.SmsGroup{}
	switch c.SmsConf.Provider {
	case smsprovider.Tencent:
		smsGroup.TencentSmsClient = c.SmsConf.NewTencentClient()
	default:
		smsGroup.TencentSmsClient = c.SmsConf.NewTencentClient()
	}

	return &ServiceContext{
		Config:    c,
		DB:        db,
		Redis:     redis.MustNewRedis(c.RedisConf),
		EmailAuth: c.EmailConf.NewAuth(),
		SmsGroup:  smsGroup,
	}
}
