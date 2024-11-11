package model

// TotalVisits TotalVisits 总访问量
type TotalVisits struct {
	Total int64 `json:"totalVisits"`
}

// DailyHourlyVisits 每日/每小时访问量
type DailyHourlyVisits struct {
	Date        string `json:"date"`
	Hour        int    `json:"hour"`
	TotalVisits int64  `json:"totalVisits"`
}

// ShortLinkVisits 每个短链接的访问量
type ShortLinkVisits struct {
	ShortLink   string `json:"shortLink"`
	TotalVisits int64  `json:"totalVisits"`
}

// GeoLocationDistribution 用户地理位置分布
type GeoLocationDistribution struct {
	Country     string `json:"country"`
	TotalVisits int64  `json:"totalVisits"`
}

// DeviceTypeDistribution 设备类型分布
type DeviceTypeDistribution struct {
	DeviceType  string `json:"deviceType"`
	TotalVisits int64  `json:"totalVisits"`
}

// BrowserTypeDistribution 浏览器类型分布
type BrowserTypeDistribution struct {
	BrowserType string `json:"browserType"`
	TotalVisits int64  `json:"totalVisits"`
}

// AverageResponseTime 平均响应时间
type AverageResponseTime struct {
	AverageTime float64 `json:"averageResponseTime"`
}

// ResponseTimeRange 最大/最小响应时间
type ResponseTimeRange struct {
	MaxTime float64 `json:"maxResponseTime"`
	MinTime float64 `json:"minResponseTime"`
}

// ErrorRate 错误率
type ErrorRate struct {
	ErrorRate float64 `json:"errorRate"`
}

// RefererDistribution 来源URL分布
type RefererDistribution struct {
	Referer     string `json:"referer"`
	TotalVisits int64  `json:"totalVisits"`
}
