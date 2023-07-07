package email

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/errorx"
	"strings"

	"github.com/suyuan32/simple-admin-message-center/internal/svc"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *mcms.EmailInfo) (*mcms.BaseUUIDResp, error) {
	client := l.svcCtx.Config.EmailConf.NewClient(l.svcCtx.EmailAuth)

	// error handler
	emailErrHandler := func(err error) (*mcms.BaseUUIDResp, error) {
		l.Logger.Errorw("failed to send email", logx.Field("detail", err.Error()), logx.Field("data", in))

		dberr := l.svcCtx.DB.EmailLog.Create().
			SetTarget(strings.Join(in.Target, ",")).
			SetContent(in.Content).
			SetSubject(in.Subject).
			SetSendStatus(2).Exec(l.ctx)

		if dberr != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, dberr, in)
		}

		return nil, errorx.NewInternalError(i18n.Failed)
	}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = l.svcCtx.Config.EmailConf.EmailAddr
	headers["To"] = strings.Join(in.Target, ",")
	headers["Subject"] = in.Subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + in.Content

	err := client.Mail(l.svcCtx.Config.EmailConf.EmailAddr)
	if err != nil {
		l.Logger.Errorw("failed to set the from address in email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to set the from address in email"))
	}

	for _, v := range in.Target {
		err := client.Rcpt(v)
		if err != nil {
			l.Logger.Errorw("failed to set the to address in email", logx.Field("detail", err), logx.Field("data", in))
			return emailErrHandler(errors.Wrap(err, "failed to set the from address in email"))
		}
	}

	w, err := client.Data()
	if err != nil {
		l.Logger.Errorw("failed to create the writer for email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to create the writer for email"))
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		l.Logger.Errorw("failed to write the data to writer for email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to write the data to writer for email"))
	}

	err = w.Close()
	if err != nil {
		l.Logger.Errorw("failed to close the writer for email", logx.Field("detail", err), logx.Field("data", in))
		return emailErrHandler(errors.Wrap(err, "failed to close the writer for email"))
	}

	logData, dberr := l.svcCtx.DB.EmailLog.Create().
		SetTarget(strings.Join(in.Target, ",")).
		SetContent(in.Content).
		SetSubject(in.Subject).
		SetSendStatus(1).Save(l.ctx)

	if dberr != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, dberr, in)
	}

	return &mcms.BaseUUIDResp{Msg: i18n.Success, Id: logData.ID.String()}, nil
}
