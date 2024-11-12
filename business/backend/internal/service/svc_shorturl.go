package service

import (
	"backend/internal/config"
	"backend/internal/dao/local"
	"backend/internal/dao/repo"
	"common/constinfo"
	"common/db"
	"common/logger"
	"common/rdb"
	"common/short"
	"context"
	"fmt"
	"github.com/pkg/errors"
)

type SvcShortUrl struct {
	log          *logger.Logger
	conf         *config.Config
	redisWrapper rdb.IRedis
	repoShortUrl *repo.RepoShortUrl
	short        *short.Short
	localCache   *local.LruCache
}

func NewSvcShort(
	log *logger.Logger,
	conf *config.Config,
	redisWrapper rdb.IRedis,
	repoShortUrl *repo.RepoShortUrl,
	short *short.Short,
	localCache *local.LruCache,
) *SvcShortUrl {
	return &SvcShortUrl{
		log:          log,
		conf:         conf,
		redisWrapper: redisWrapper,
		repoShortUrl: repoShortUrl,
		short:        short,
		localCache:   localCache,
	}
}

// GetLongUrl 根据短链找长链
// 请求短连接重定向流程:
// [x]1.http请求(参数短链)
// [x]2.布隆过滤器B判断短链是否存在
//
//	[x]2.1.不存在, 直接返回404
//	2.2.存在
//	    2.2.1.查本地缓存(热门LRU). 存在, 重定向对应短链,刷新LRU; 不存在, 下一步
//	    2.2.2.查redis(最近/热门). 存在, 重定向对应短链, 延长key过期时间; 不存在, 下一步
//	    2.2.3.查db. 存在, 重定向对应短链, 刷新本地缓存LRU和redis; 不存在, 返回404
//
// 3.找不到长链, 返回404页面
func (s *SvcShortUrl) GetLongUrl(ctx context.Context, shortUrl string) (longUrl string, err error) {
	// 布隆过滤器判断短链是否存在, 不存在, 快速返回
	exist, err := s.redisWrapper.BloomExists(ctx, constinfo.KeyShortBloom, shortUrl)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if !exist {
		s.log.WithField("shortUrl", shortUrl).Warn("short url not exist")
		return
	}

	// 查本地缓存lru, 存在, 快速返回
	long, err := s.localCache.Get([]byte(shortUrl))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if len(long) > 0 {
		longUrl = string(long)
		// 刷新lru
		s.localCache.Del([]byte(shortUrl))
		_ = s.localCache.Set([]byte(shortUrl), []byte(longUrl), s.conf.LocalCache.Expire)
		s.log.WithField("shortUrl", shortUrl).Info("short url exist cache")
		return
	}

	// 查redis缓存, 存在直接返回
	redisKey := fmt.Sprintf("%s%s", constinfo.KeyShortUrl, shortUrl)
	longUrl, err = s.redisWrapper.Get(ctx, redisKey)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if len(longUrl) > 0 {
		longUrl = string(long)
		// 续期
		_ = s.redisWrapper.SetEx(ctx, redisKey, longUrl, 3600)
		s.log.WithField("shortUrl", shortUrl).Info("short url exist cache")
		return
	}

	// 最后查db, 看布隆过滤器是否幻觉; 不存在, 证明短链真的不存在长链, 直接返回
	tb := db.GetTBShardingName(db.TBShortUrl, s.short.ToSnowFlakeID(shortUrl)%s.conf.Database.Sharding)
	longUrl, err = s.repoShortUrl.Get(ctx, tb, shortUrl)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if len(longUrl) == 0 {
		s.log.WithField("shortUrl", shortUrl).Debug("short had not found long")
		return
	}

	// 存在短链, 缓存
	go func() {
		_ = s.localCache.Set([]byte(shortUrl), []byte(longUrl), s.conf.LocalCache.Expire)
		_ = s.redisWrapper.SetEx(ctx, redisKey, longUrl, 3600)
	}()

	s.log.WithField("shortUrl", shortUrl).Info("short url exist db")
	return
}
