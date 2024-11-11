package dao

import (
	"backend/internal/dao/redis"
	"backend/internal/dao/repo"
	"common/db"
	"common/logger"
	"common/rdb"
	"common/utils"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo.NewDB,
	repo.NewRepoShortUrl,
	repo.NewRepoAccessLog,
	redis.NewRedis,
)

type Data struct {
	db  db.DB
	rds rdb.IRedis
	sf  *utils.SnowFlake
}

func NewData(
	l *logger.Logger,
	db db.DB,
	rds rdb.IRedis,
	sf *utils.SnowFlake,
) (*Data, func(), error) {
	clean := func() {
		l.Info("close data resource")
	}
	return &Data{
		db:  db,
		rds: rds,
		sf:  sf,
	}, clean, nil
}
