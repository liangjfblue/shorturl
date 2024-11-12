package db

import "fmt"

const (
	// DBShortUrl 数据库
	// DBShortUrl = "shorturl"

	// TBShortUrl 长短链映射
	TBShortUrl = "tb_shortUrl_"
	// TBAccessLog 访问日志
	TBAccessLog = "tb_accessLog_"
)

func GetTBShardingName(tb string, num int64) string {
	return fmt.Sprintf("%s_%d", tb, num)
}
