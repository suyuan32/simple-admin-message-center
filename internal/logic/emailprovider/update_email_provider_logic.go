package emailprovider

import (
	"context"
	"net/smtp"

	"github.com/suyuan32/simple-admin-message-center/internal/svc"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailProviderLogic {
	return &UpdateEmailProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmailProviderLogic) UpdateEmailProvider(in *mcms.EmailProviderInfo) (*mcms.BaseResp, error) {
	query := l.svcCtx.DB.EmailProvider.UpdateOneID(*in.Id).
		SetNotNilName(in.Name).
		SetNotNilEmailAddr(in.EmailAddr).
		SetNotNilPassword(in.Password).
		SetNotNilHostName(in.HostName).
		SetNotNilIdentify(in.Identify).
		SetNotNilSecret(in.Secret).
		SetNotNilPort(in.Port).
		SetNotNilTLS(in.Tls).
		SetNotNilIsDefault(in.IsDefault)

	if in.AuthType != nil {
		query.SetNotNilAuthType(pointy.GetPointer(uint8(*in.AuthType)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	l.svcCtx.EmailAddrGroup = map[string]string{}
	l.svcCtx.EmailClientGroup = map[string]*smtp.Client{}

	return &mcms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
