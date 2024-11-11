package repo

import (
	"common/db"
	"common/model"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type RepoShortUrl struct {
	db db.DB
}

func NewRepoShortUrl(db db.DB) *RepoShortUrl {
	return &RepoShortUrl{db: db}
}

func (d *RepoShortUrl) Get(ctx context.Context, tb string, shortUrl string) (longUrl string, err error) {
	query := map[string]interface{}{"short": shortUrl}
	var item model.ShortUrl
	err = d.db.Find(ctx, tb, query, &item)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	longUrl = item.Long
	return
}
