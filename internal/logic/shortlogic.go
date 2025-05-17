package logic

import (
	"context"
	"database/sql"

	"shorter/internal/model"
	"shorter/internal/svc"
	"shorter/internal/types"
	"shorter/pkg/utils"

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
	const BaseShortDomain = "http://localhost:8888"

	// Check if the URL is already in the database
	existShorter, err := l.svcCtx.Model.FindOneByRawUrl(l.ctx, req.Url)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("GetUrl err: %v", err)
		return nil, err
	}

	if existShorter != nil {
		resp = &types.RespShortUrl{
			ShortUrl: BaseShortDomain + "/s/" + existShorter.ShortCode.String,
		}
		return
	}

	// Insert the new URL
	l.svcCtx.Model.Insert(l.ctx, &model.Shorter{
		RawUrl: req.Url,
	})
	insertShorter, err := l.svcCtx.Model.FindOneByRawUrl(l.ctx, req.Url)
	if err != nil {
		l.Logger.Errorf("GetUrlId err: %v", err)
		return nil, err
	}

	shortCode := utils.HashEncode(insertShorter.Id)

	if err := l.svcCtx.Model.Update(l.ctx, &model.Shorter{
		Id:        insertShorter.Id,
		RawUrl:    req.Url,
		ShortCode: sql.NullString{String: shortCode, Valid: true},
	}); err != nil {
		l.Logger.Errorf("UpdateShortCode err: %v", err)
		return nil, err
	}

	resp = &types.RespShortUrl{
		ShortUrl: BaseShortDomain + "/s/" + shortCode,
	}
	return
}
