package model

// AccessLog 访问日志
type AccessLog struct {
	ID           string  // 主键
	Timestamp    int64   // 访问时间
	ShortUrl     string  // 短接
	LongUrl      string  // 长链
	ClientIp     string  // 客户端IP
	HttpMethod   string  // HTTP方法
	UserAgent    string  // User-Agent
	Referer      string  // Referer
	StatusCode   int     // 响应状态码
	ResponseTime float64 // 响应时间
	ErrorMessage string  // 错误信息
	Type         int8    // 类型 系统生成/自定义
}
