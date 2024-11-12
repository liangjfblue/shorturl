package local

import (
	"backend/internal/config"
	"github.com/coocood/freecache"
)

type LruCache struct {
	conf *config.Config
	*freecache.Cache
}

func NewLruCache(conf *config.Config) *LruCache {
	r := &LruCache{
		conf: conf,
	}
	r.Cache = freecache.NewCache(conf.LocalCache.Memory * 1024 * 1024)
	return r
}
