package svc

import (
	"shorter/internal/config"
	"shorter/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShorterModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewShorterModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
