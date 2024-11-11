package service

import (
	"backend/internal/dao/repo"
	"common/db"
	"common/logger"
	"common/model"
	"context"
	"github.com/pkg/errors"
)

type SvcAccessLog struct {
	log           *logger.Logger
	repoAccessLog *repo.Repo[model.AccessLog]
}

func NewSvcAccessLog(log *logger.Logger, repoAccessLog *repo.Repo[model.AccessLog]) *SvcAccessLog {
	return &SvcAccessLog{
		log:           log,
		repoAccessLog: repoAccessLog,
	}
}

func (s *SvcAccessLog) Record(ctx context.Context, item model.AccessLog) (err error) {
	err = s.repoAccessLog.Save(ctx, db.TBAccessLog, item)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	s.log.WithField("item", item).Info("record access log")
	return
}
