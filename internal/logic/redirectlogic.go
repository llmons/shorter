package logic

import (
	"context"
	"database/sql"

	"shorter/internal/model"
	"shorter/internal/svc"
	"shorter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectLogic {
	return &RedirectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectLogic) Redirect(req *types.ReqShortCode) (resp *types.RespRedirect, err error) {
	shorter, err := l.svcCtx.Model.FindOneByShortCode(l.ctx, sql.NullString{String: req.ShortCode, Valid: true})
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("GetUrl err: %v", err)
		return nil, err
	}
	if shorter == nil {
		return nil, model.ErrNotFound
	}
	resp = &types.RespRedirect{
		Url: shorter.RawUrl,
	}

	return
}
