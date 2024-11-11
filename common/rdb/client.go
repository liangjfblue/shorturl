package rdb

import "github.com/redis/go-redis/v9"

// NewSingle 单节点
func NewSingle(addr string, password string) IRedis {
	return &ClientWrapper[*redis.Client]{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
		}),
	}
}

// NewSentinel sentinel架构
func NewSentinel(masterName string, sentinelAddress []string, password string) IRedis {
	return &ClientWrapper[*redis.Client]{
		client: redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    masterName,
			SentinelAddrs: sentinelAddress,
			Password:      password,
		}),
	}
}

// NewCluster cluster架构
func NewCluster(clusterAddress []string, password string) IRedis {
	return &ClientWrapper[*redis.ClusterClient]{
		client: redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    clusterAddress,
			Password: password,
		}),
	}
}
