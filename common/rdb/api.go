package rdb

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// IRedis 定义了通用的Redis操作接口
type IRedis interface {
	Set(ctx context.Context, key string, value any, expiration int) error
	SetEx(ctx context.Context, key string, value any, expiration int) error
	SetNx(ctx context.Context, key string, value any) (bool, error)
	ExpireSec(ctx context.Context, key string, second int) error
	ExpireMill(ctx context.Context, key string, mill int) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) (int64, error)
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)

	HSet(ctx context.Context, key string, field string, value any) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) (int64, error)

	LPush(ctx context.Context, key string, values ...any) (int64, error)
	RPush(ctx context.Context, key string, values ...any) (int64, error)
	LPop(ctx context.Context, key string) (string, error)
	RPop(ctx context.Context, key string) (string, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)

	SAdd(ctx context.Context, key string, members ...any) (int64, error)
	SRem(ctx context.Context, key string, members ...any) (int64, error)
	SMembers(ctx context.Context, key string) ([]string, error)
	SIsMember(ctx context.Context, key string, member any) (bool, error)

	ZAdd(ctx context.Context, key string, members ...Z) (int64, error)
	ZRem(ctx context.Context, key string, members ...any) (int64, error)
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error)

	// Publish(channel string, message any) error
	// Subscribe(channels ...string) (Subscriber, error)

	BloomAdd(ctx context.Context, key string, item any) error
	BloomExists(ctx context.Context, key string, item any) (bool, error)
	BloomReserve(ctx context.Context, key string, capacity int, errorRate float64) error

	Close() error
}

// Z 表示有序集合中的成员
type Z struct {
	Score  float64
	Member any
}

// Subscriber 定义了订阅者的接口
type Subscriber interface {
	Receive() (string, error)
	Channel() <-chan string
	Close() error
}

// ClientWrapper redis client 包装器
type ClientWrapper[T any] struct {
	client T
}

func NewClientWrapper[T any](client T) *ClientWrapper[T] {
	return &ClientWrapper[T]{client: client}
}

// Set 设置键值对
func (c *ClientWrapper[T]) Set(ctx context.Context, key string, value any, expiration int) error {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err()
	case *redis.ClusterClient:
		return client.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err()
	default:
		return nil
	}
}

// SetEx 设置键值对并带有过期时间
func (c *ClientWrapper[T]) SetEx(ctx context.Context, key string, value any, expiration int) error {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.SetEx(ctx, key, value, time.Duration(expiration)*time.Second).Err()
	case *redis.ClusterClient:
		return client.SetEx(ctx, key, value, time.Duration(expiration)*time.Second).Err()
	default:
		return nil
	}
}

// SetNx 设置键值对并仅在键不存在时生效
func (c *ClientWrapper[T]) SetNx(ctx context.Context, key string, value any) (bool, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.SetNX(ctx, key, value, 0).Result()
	case *redis.ClusterClient:
		return client.SetNX(ctx, key, value, 0).Result()
	default:
		return false, nil
	}
}

// ExpireSec 设置键的过期时间（以秒为单位）
func (c *ClientWrapper[T]) ExpireSec(ctx context.Context, key string, second int) error {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Expire(ctx, key, time.Duration(second)*time.Second).Err()
	case *redis.ClusterClient:
		return client.Expire(ctx, key, time.Duration(second)*time.Second).Err()
	default:
		return nil
	}
}

// ExpireMill 设置键的过期时间（以毫秒为单位）
func (c *ClientWrapper[T]) ExpireMill(ctx context.Context, key string, mill int) error {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Expire(ctx, key, time.Duration(mill)*time.Millisecond).Err()
	case *redis.ClusterClient:
		return client.Expire(ctx, key, time.Duration(mill)*time.Millisecond).Err()
	default:
		return nil
	}
}

// Get 获取键对应的值
func (c *ClientWrapper[T]) Get(ctx context.Context, key string) (string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Get(ctx, key).Result()
	case *redis.ClusterClient:
		return client.Get(ctx, key).Result()
	default:
		return "", nil
	}
}

// Del 删除一个或多个键
func (c *ClientWrapper[T]) Del(ctx context.Context, keys ...string) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Del(ctx, keys...).Result()
	case *redis.ClusterClient:
		return client.Del(ctx, keys...).Result()
	default:
		return 0, nil
	}
}

// Incr 将键的值增加1
func (c *ClientWrapper[T]) Incr(ctx context.Context, key string) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Incr(ctx, key).Result()
	case *redis.ClusterClient:
		return client.Incr(ctx, key).Result()
	default:
		return 0, nil
	}
}

// Decr 将键的值减少1
func (c *ClientWrapper[T]) Decr(ctx context.Context, key string) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Decr(ctx, key).Result()
	case *redis.ClusterClient:
		return client.Decr(ctx, key).Result()
	default:
		return 0, nil
	}
}

// HSet 设置哈希表中的字段值
func (c *ClientWrapper[T]) HSet(ctx context.Context, key string, field string, value any) error {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.HSet(ctx, key, field, value).Err()
	case *redis.ClusterClient:
		return client.HSet(ctx, key, field, value).Err()
	default:
		return nil
	}
}

// HGet 获取哈希表中指定字段的值
func (c *ClientWrapper[T]) HGet(ctx context.Context, key string, field string) (string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.HGet(ctx, key, field).Result()
	case *redis.ClusterClient:
		return client.HGet(ctx, key, field).Result()
	default:
		return "", nil
	}
}

// HGetAll 获取哈希表中所有字段和值
func (c *ClientWrapper[T]) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.HGetAll(ctx, key).Result()
	case *redis.ClusterClient:
		return client.HGetAll(ctx, key).Result()
	default:
		return nil, nil
	}
}

// HDel 删除哈希表中的一个或多个字段
func (c *ClientWrapper[T]) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.HDel(ctx, key, fields...).Result()
	case *redis.ClusterClient:
		return client.HDel(ctx, key, fields...).Result()
	default:
		return 0, nil
	}
}

// LPush 将一个或多个值插入到列表头部
func (c *ClientWrapper[T]) LPush(ctx context.Context, key string, values ...any) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.LPush(ctx, key, values...).Result()
	case *redis.ClusterClient:
		return client.LPush(ctx, key, values...).Result()
	default:
		return 0, nil
	}
}

// RPush 将一个或多个值插入到列表尾部
func (c *ClientWrapper[T]) RPush(ctx context.Context, key string, values ...any) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.RPush(ctx, key, values...).Result()
	case *redis.ClusterClient:
		return client.RPush(ctx, key, values...).Result()
	default:
		return 0, nil
	}
}

// LPop 移除并返回列表的第一个元素
func (c *ClientWrapper[T]) LPop(ctx context.Context, key string) (string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.LPop(ctx, key).Result()
	case *redis.ClusterClient:
		return client.LPop(ctx, key).Result()
	default:
		return "", nil
	}
}

// RPop 移除并返回列表的最后一个元素
func (c *ClientWrapper[T]) RPop(ctx context.Context, key string) (string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.RPop(ctx, key).Result()
	case *redis.ClusterClient:
		return client.RPop(ctx, key).Result()
	default:
		return "", nil
	}
}

// LRange 获取列表中指定范围内的元素
func (c *ClientWrapper[T]) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.LRange(ctx, key, start, stop).Result()
	case *redis.ClusterClient:
		return client.LRange(ctx, key, start, stop).Result()
	default:
		return nil, nil
	}
}

// SAdd 向集合中添加一个或多个成员
func (c *ClientWrapper[T]) SAdd(ctx context.Context, key string, members ...any) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.SAdd(ctx, key, members...).Result()
	case *redis.ClusterClient:
		return client.SAdd(ctx, key, members...).Result()
	default:
		return 0, nil
	}
}

// SRem 从集合中移除一个或多个成员
func (c *ClientWrapper[T]) SRem(ctx context.Context, key string, members ...any) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.SRem(ctx, key, members...).Result()
	case *redis.ClusterClient:
		return client.SRem(ctx, key, members...).Result()
	default:
		return 0, nil
	}
}

// SMembers 获取集合中的所有成员
func (c *ClientWrapper[T]) SMembers(ctx context.Context, key string) ([]string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.SMembers(ctx, key).Result()
	case *redis.ClusterClient:
		return client.SMembers(ctx, key).Result()
	default:
		return nil, nil
	}
}

// SIsMember 判断成员是否在集合中
func (c *ClientWrapper[T]) SIsMember(ctx context.Context, key string, member any) (bool, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.SIsMember(ctx, key, member).Result()
	case *redis.ClusterClient:
		return client.SIsMember(ctx, key, member).Result()
	default:
		return false, nil
	}
}

// ZAdd 向有序集合中添加一个或多个成员
func (c *ClientWrapper[T]) ZAdd(ctx context.Context, key string, members ...Z) (int64, error) {
	var zMembers []redis.Z
	for _, member := range members {
		zMembers = append(zMembers, redis.Z{
			Score:  member.Score,
			Member: member.Member,
		})
	}
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.ZAdd(ctx, key, zMembers...).Result()
	case *redis.ClusterClient:
		return client.ZAdd(ctx, key, zMembers...).Result()
	default:
		return 0, nil
	}
}

// ZRem 从有序集合中移除一个或多个成员
func (c *ClientWrapper[T]) ZRem(ctx context.Context, key string, members ...any) (int64, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.ZRem(ctx, key, members...).Result()
	case *redis.ClusterClient:
		return client.ZRem(ctx, key, members...).Result()
	default:
		return 0, nil
	}
}

// ZRange 获取有序集合中指定范围内的成员
func (c *ClientWrapper[T]) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.ZRange(ctx, key, start, stop).Result()
	case *redis.ClusterClient:
		return client.ZRange(ctx, key, start, stop).Result()
	default:
		return nil, nil
	}
}

// ZRangeWithScores 获取有序集合中指定范围内的成员及其分数
func (c *ClientWrapper[T]) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	switch client := any(c.client).(type) {
	case *redis.Client:
		results, err := client.ZRangeWithScores(ctx, key, start, stop).Result()
		if err != nil {
			return nil, err
		}
		var zs []Z
		for _, result := range results {
			zs = append(zs, Z{
				Score:  result.Score,
				Member: result.Member,
			})
		}
		return zs, nil
	case *redis.ClusterClient:
		results, err := client.ZRangeWithScores(ctx, key, start, stop).Result()
		if err != nil {
			return nil, err
		}
		var zs []Z
		for _, result := range results {
			zs = append(zs, Z{
				Score:  result.Score,
				Member: result.Member,
			})
		}
		return zs, nil
	default:
		return nil, nil
	}
}

// BloomAdd 添加元素到布隆过滤器
func (c *ClientWrapper[T]) BloomAdd(ctx context.Context, key string, item any) error {
	itemStr, ok := item.(string)
	if !ok {
		return fmt.Errorf("item must be a string")
	}

	switch client := any(c.client).(type) {
	case *redis.Client:
		_, err := client.Do(ctx, "BF.ADD", key, itemStr).Result()
		return err
	case *redis.ClusterClient:
		_, err := client.Do(ctx, "BF.ADD", key, itemStr).Result()
		return err
	default:
		return nil
	}
}

// BloomExists checks if an item might be in the Bloom filter
func (c *ClientWrapper[T]) BloomExists(ctx context.Context, key string, item any) (bool, error) {
	itemStr, ok := item.(string)
	if !ok {
		return false, fmt.Errorf("item must be a string")
	}

	var err error
	var result any
	switch client := any(c.client).(type) {
	case *redis.Client:
		result, err = client.Do(ctx, "BF.EXISTS", key, itemStr).Result()
	case *redis.ClusterClient:
		result, err = client.Do(ctx, "BF.EXISTS", key, itemStr).Result()
	}
	if err != nil {
		return false, err
	}
	exists, ok := result.(int64)
	if !ok {
		return false, fmt.Errorf("unexpected result type")
	}

	return exists == 1, nil
}

// BloomReserve 从布隆过滤器删除元素
func (c *ClientWrapper[T]) BloomReserve(ctx context.Context, key string, capacity int, errorRate float64) error {
	var err error
	switch client := any(c.client).(type) {
	case *redis.Client:
		_, err = client.Do(ctx, "BF.RESERVE", key, capacity, errorRate).Result()
	case *redis.ClusterClient:
		_, err = client.Do(ctx, "BF.RESERVE", key, capacity, errorRate).Result()
	}
	return err
}

// Close 关闭Redis连接
func (c *ClientWrapper[T]) Close() error {
	switch client := any(c.client).(type) {
	case *redis.Client:
		return client.Close()
	case *redis.ClusterClient:
		return client.Close()
	default:
		return nil
	}
}
