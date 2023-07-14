package smsprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-message-center/ent/predicate"
	"github.com/suyuan32/simple-admin-message-center/ent/smsprovider"
	"github.com/suyuan32/simple-admin-message-center/internal/svc"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderListLogic {
	return &GetSmsProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsProviderListLogic) GetSmsProviderList(in *mcms.SmsProviderListReq) (*mcms.SmsProviderListResp, error) {
	var predicates []predicate.SmsProvider
	if in.Name != nil {
		predicates = append(predicates, smsprovider.NameContains(*in.Name))
	}

	result, err := l.svcCtx.DB.SmsProvider.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mcms.SmsProviderListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mcms.SmsProviderInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:      &v.Name,
			SecretId:  &v.SecretID,
			SecretKey: &v.SecretKey,
			Region:    &v.Region,
			IsDefault: &v.IsDefault,
		})
	}

	return resp, nil
}
