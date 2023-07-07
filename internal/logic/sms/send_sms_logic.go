package sms

import (
	"context"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-message-center/internal/enum/smsprovider"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/dberrorhandler"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/zeromicro/go-zero/core/errorx"
	"strings"

	"github.com/suyuan32/simple-admin-message-center/internal/svc"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *mcms.SmsInfo) (*mcms.BaseUUIDResp, error) {
	switch l.svcCtx.Config.SmsConf.Provider {
	case smsprovider.Tencent:
		request := sms.NewSendSmsRequest()
		request.TemplateId = in.TemplateId
		request.SmsSdkAppId = in.AppId
		request.PhoneNumberSet = pointy.GetSlicePointer(in.PhoneNumber)
		request.TemplateParamSet = pointy.GetSlicePointer(in.Params)
		request.SignName = in.SignName
		resp, err := l.svcCtx.SmsGroup.TencentSmsClient.SendSms(request)
		if err != nil {
			logx.Errorw("failed to send sms", logx.Field("detail", err), logx.Field("data", in))

			err = l.svcCtx.DB.SmsLog.Create().
				SetSendStatus(2).
				SetContent(strings.Join(in.Params, ",")).
				SetPhoneNumber(strings.Join(in.PhoneNumber, ",")).
				SetProvider(l.svcCtx.Config.SmsConf.Provider).
				Exec(l.ctx)

			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
			}

			return nil, errorx.NewInternalError(i18n.Failed)
		}
		logx.Infow("send sms by tencent", logx.Field("response", resp), logx.Field("phoneNumber", in.PhoneNumber))
	}

	logData, err := l.svcCtx.DB.SmsLog.Create().
		SetSendStatus(1).
		SetContent(strings.Join(in.Params, ",")).
		SetPhoneNumber(strings.Join(in.PhoneNumber, ",")).
		SetProvider(l.svcCtx.Config.SmsConf.Provider).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseUUIDResp{Id: logData.ID.String(), Msg: i18n.Success}, nil
}
