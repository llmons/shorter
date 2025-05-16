package logic

import (
	"context"

	"shorter/internal/svc"
	"shorter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortUrlLogic {
	return &GetShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShortUrlLogic) GetShortUrl(req *types.ReqShortUrlId) (resp *types.RespShortUrl, err error) {
	resp = &types.RespShortUrl{
		ShortUrl: `http://shorter.com/` + req.Id,
	}
	return
}
