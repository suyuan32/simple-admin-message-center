package sms

import (
	"context"
	aliyun "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	smsprovider2 "github.com/suyuan32/simple-admin-message-center/ent/smsprovider"
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
	// If the provider is nil, use default
	if in.Provider == nil {
		defaultProvider, err := l.svcCtx.DB.SmsProvider.Query().Where(smsprovider2.IsDefaultEQ(true)).First(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}
		in.Provider = &defaultProvider.Name
	}

	// init group
	err := l.initProvider(in)
	if err != nil {
		return nil, err
	}

	// send message
	switch *in.Provider {
	case smsprovider.Tencent:
		request := sms.NewSendSmsRequest()
		request.TemplateId = in.TemplateId
		request.SmsSdkAppId = in.AppId
		request.PhoneNumberSet = pointy.GetSlicePointer(in.PhoneNumber)
		request.TemplateParamSet = pointy.GetSlicePointer(in.Params)
		request.SignName = in.SignName
		resp, err := l.svcCtx.SmsGroup.TencentSmsClient.SendSms(request)
		if err != nil {
			logx.Errorw("failed to send SMS", logx.Field("detail", err), logx.Field("data", in))

			err = l.svcCtx.DB.SmsLog.Create().
				SetSendStatus(2).
				SetContent(strings.Join(in.Params, ",")).
				SetPhoneNumber(strings.Join(in.PhoneNumber, ",")).
				SetProvider(*in.Provider).
				Exec(context.Background())

			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
			}

			return nil, errorx.NewInternalError(i18n.Failed)
		}
		logx.Infow("send SMS by Tencent", logx.Field("response", resp), logx.Field("phoneNumber", in.PhoneNumber))
	case smsprovider.Aliyun:
		request := aliyun.SendSmsRequest{}
		request.SignName = in.SignName
		request.TemplateCode = in.TemplateId
		request.PhoneNumbers = pointy.GetPointer(strings.Join(in.PhoneNumber, ","))
		if in.Params != nil {
			request.TemplateParam = pointy.GetPointer(strings.Join(in.Params, ""))
		}
		options := &util.RuntimeOptions{}
		resp, err := l.svcCtx.SmsGroup.AliyunSmsClient.SendSmsWithOptions(&request, options)
		if err != nil {
			logx.Errorw("failed to send SMS", logx.Field("detail", err), logx.Field("data", in))

			err = l.svcCtx.DB.SmsLog.Create().
				SetSendStatus(2).
				SetContent(strings.Join(in.Params, ",")).
				SetPhoneNumber(strings.Join(in.PhoneNumber, ",")).
				SetProvider(*in.Provider).
				Exec(context.Background())

			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
			}

			return nil, errorx.NewInternalError(i18n.Failed)
		}
		logx.Infow("send SMS by Aliyun", logx.Field("response", resp), logx.Field("phoneNumber", in.PhoneNumber))
	case smsprovider.Uni:
		request := unisms.BuildMessage()
		request.SetSignature(*in.SignName)
		request.SetTemplateId(*in.TemplateId)
		request.SetTo(in.PhoneNumber...)
		if in.Params != nil {
			paramsData := map[string]string{}
			for _, v := range in.Params {
				p := strings.Split(v, ":")
				if len(p) != 2 {
					logx.Errorw("wrong parameters in Uni SMS Request", logx.Field("param", in.Params))
					return nil, errorx.NewInvalidArgumentError(i18n.Failed)
				}
				paramsData[p[0]] = p[1]
			}
			request.SetTemplateData(paramsData)
		}

		resp, err := l.svcCtx.SmsGroup.UniSmsClient.Send(request)
		if err != nil {
			logx.Errorw("failed to send SMS", logx.Field("detail", err), logx.Field("data", in))

			err = l.svcCtx.DB.SmsLog.Create().
				SetSendStatus(2).
				SetContent(strings.Join(in.Params, ",")).
				SetPhoneNumber(strings.Join(in.PhoneNumber, ",")).
				SetProvider(*in.Provider).
				Exec(context.Background())

			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
			}

			return nil, errorx.NewInternalError(i18n.Failed)
		}
		logx.Infow("send SMS by Uni SMS", logx.Field("response", resp), logx.Field("phoneNumber", in.PhoneNumber))
	}

	logData, err := l.svcCtx.DB.SmsLog.Create().
		SetSendStatus(1).
		SetContent(strings.Join(in.Params, ",")).
		SetPhoneNumber(strings.Join(in.PhoneNumber, ",")).
		SetProvider(*in.Provider).
		Save(context.Background())

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseUUIDResp{Id: logData.ID.String(), Msg: i18n.Success}, nil
}
