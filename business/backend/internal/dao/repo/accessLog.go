package repo

import (
	"common/db"
	"common/model"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type RepoAccessLog struct {
	db db.DB
}

func NewRepoAccessLog(db db.DB) *RepoAccessLog {
	return &RepoAccessLog{db: db}
}

func (d *RepoAccessLog) Save(ctx context.Context, tb string, item model.AccessLog) (err error) {
	err = d.db.Save(ctx, tb, item)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
