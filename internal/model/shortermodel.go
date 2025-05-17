package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShorterModel = (*customShorterModel)(nil)

type (
	// ShorterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShorterModel.
	ShorterModel interface {
		shorterModel
	}

	customShorterModel struct {
		*defaultShorterModel
	}
)

// NewShorterModel returns a model for the database table.
func NewShorterModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ShorterModel {
	return &customShorterModel{
		defaultShorterModel: newShorterModel(conn, c, opts...),
	}
}
