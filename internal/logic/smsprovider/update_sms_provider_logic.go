package smsprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-message-center/internal/svc"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSmsProviderLogic {
	return &UpdateSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSmsProviderLogic) UpdateSmsProvider(in *mcms.SmsProviderInfo) (*mcms.BaseResp, error) {
	err := l.svcCtx.DB.SmsProvider.UpdateOneID(*in.Id).
		SetNotNilName(in.Name).
		SetNotNilSecretID(in.SecretId).
		SetNotNilSecretKey(in.SecretKey).
		SetNotNilRegion(in.Region).
		SetNotNilIsDefault(in.IsDefault).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
