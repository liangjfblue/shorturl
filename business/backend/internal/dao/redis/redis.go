package redis

import (
	"backend/internal/config"
	"common/rdb"
)

// NewRedis init redis
func NewRedis(c *config.Config) (rdb.IRedis, func(), error) {
	var r rdb.IRedis
	switch c.Redis.Type {
	case "single":
		if len(c.Redis.Single.Host) == 0 {
			panic("redis host is empty")
		}
		r = rdb.NewSingle(c.Redis.Single.Host, c.Redis.Single.Password)
	case "sentinel":
		if len(c.Redis.Sentinel.Hosts) == 0 {
			panic("redis host is empty")
		}
		r = rdb.NewSentinel(c.Redis.Sentinel.MasterName, c.Redis.Sentinel.Hosts, c.Redis.Sentinel.Password)
	case "cluster":
		if len(c.Redis.Cluster.Hosts) == 0 {
			panic("redis cluster host is empty")
		}
		r = rdb.NewCluster(c.Redis.Cluster.Hosts, c.Redis.Cluster.Password)
	}
	return r, func() { _ = r.Close() }, nil
}
