package service

import (
	"backend/internal/dao/repo"
	"common/logger"
	"common/rdb"
	"context"
	"github.com/pkg/errors"
)

type SvcShort struct {
	log          *logger.Logger
	redisWrapper rdb.IRedis
	repoShortUrl *repo.RepoShortUrl
}

func NewSvcShort(
	log *logger.Logger,
	redisWrapper rdb.IRedis,
	repoShortUrl *repo.RepoShortUrl,
) *SvcShort {
	return &SvcShort{
		log:          log,
		redisWrapper: redisWrapper,
		repoShortUrl: repoShortUrl,
	}
}

func (s *SvcShort) GetLongUrl(ctx context.Context, short string) (long string, err error) {
	// 请求短连接重定向流程:
	// 1.http请求(参数短链)
	// 2.布隆过滤器B判断短链是否存在
	//    2.1.不存在, 直接返回404
	//    2.2.存在
	//        2.2.1.查本地缓存(热门LRU). 存在, 重定向对应短链,刷新LRU; 不存在, 下一步
	//        2.2.2.查redis(最近/热门). 存在, 重定向对应短链, 延长key过期时间; 不存在, 下一步
	//        2.2.3.查db. 存在, 重定向对应短链, 刷新本地缓存LRU和redis; 不存在, 返回404
	// 3.找不到长链, 返回404页面

	longUrl, err := s.dbShort.GetShortUrl(ctx, short)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if len(longUrl) == 0 {
		s.log.L.WithField("short", short).Debug("short had not found long")
		return
	}

	return
}
