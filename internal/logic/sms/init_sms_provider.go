package sms

import (
	"github.com/pkg/errors"
	smsprovider2 "github.com/suyuan32/simple-admin-message-center/ent/smsprovider"
	"github.com/suyuan32/simple-admin-message-center/internal/enum/smsprovider"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/smssdk"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"
	"github.com/zeromicro/go-zero/core/errorx"
)

func (l *SendSmsLogic) initProvider(in *mcms.SmsInfo) error {
	switch *in.Provider {
	case smsprovider.Tencent:
		if l.svcCtx.SmsGroup.TencentSmsClient == nil {
			data, err := l.svcCtx.DB.SmsProvider.Query().Where(smsprovider2.NameEQ(*in.Provider)).First(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, in)
			}
			clientConf := &smssdk.SmsConf{
				SecretId:  data.SecretID,
				SecretKey: data.SecretKey,
				Provider:  *in.Provider,
				Region:    data.Region,
			}
			l.svcCtx.SmsGroup.TencentSmsClient = clientConf.NewTencentClient()
		}
	case smsprovider.Aliyun:
		if l.svcCtx.SmsGroup.AliyunSmsClient == nil {
			data, err := l.svcCtx.DB.SmsProvider.Query().Where(smsprovider2.NameEQ(*in.Provider)).First(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, in)
			}
			clientConf := &smssdk.SmsConf{
				SecretId:  data.SecretID,
				SecretKey: data.SecretKey,
				Provider:  *in.Provider,
				Region:    data.Region,
			}
			l.svcCtx.SmsGroup.AliyunSmsClient, err = clientConf.NewAliyunClient()
			if err != nil {
				return errors.Wrap(err, "failed to initialize Aliyun SMS client")
			}
		}
	case smsprovider.Uni:
		if l.svcCtx.SmsGroup.UniSmsClient == nil {
			data, err := l.svcCtx.DB.SmsProvider.Query().Where(smsprovider2.NameEQ(*in.Provider)).First(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, in)
			}
			clientConf := &smssdk.SmsConf{
				SecretId:  data.SecretID,
				SecretKey: data.SecretKey,
				Provider:  *in.Provider,
				Region:    data.Region,
			}
			l.svcCtx.SmsGroup.UniSmsClient, err = clientConf.NewUniClient()
			if err != nil {
				return errors.Wrap(err, "failed to initialize Uni SMS client")
			}
		}
	default:
		return errorx.NewInvalidArgumentError("provider not found")
	}

	return nil
}
