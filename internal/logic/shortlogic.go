package logic

import (
	"context"

	"shorter/internal/svc"
	"shorter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortLogic {
	return &ShortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortLogic) Short(req *types.ReqUrl) (resp *types.RespShortUrl, err error) {
	resp = &types.RespShortUrl{
		ShortUrl: req.Url,
	}
	return
}
