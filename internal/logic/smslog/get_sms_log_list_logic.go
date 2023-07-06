package smslog

import (
	"context"

	"github.com/suyuan32/simple-admin-message-center/ent/predicate"
	"github.com/suyuan32/simple-admin-message-center/ent/smslog"
	"github.com/suyuan32/simple-admin-message-center/internal/svc"
	"github.com/suyuan32/simple-admin-message-center/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogListLogic {
	return &GetSmsLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsLogListLogic) GetSmsLogList(in *mcms.SmsLogListReq) (*mcms.SmsLogListResp, error) {
	var predicates []predicate.SmsLog
	if in.PhoneNumber != nil {
		predicates = append(predicates, smslog.PhoneNumberContains(*in.PhoneNumber))
	}
	if in.Content != nil {
		predicates = append(predicates, smslog.ContentContains(*in.Content))
	}
	if in.Provider != nil {
		predicates = append(predicates, smslog.ProviderContains(*in.Provider))
	}
	result, err := l.svcCtx.DB.SmsLog.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mcms.SmsLogListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mcms.SmsLogInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			PhoneNumber: &v.PhoneNumber,
			Content:     &v.Content,
			SendStatus:  pointy.GetPointer(uint32(v.SendStatus)),
			Provider:    &v.Provider,
		})
	}

	return resp, nil
}
