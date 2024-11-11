package model

// ShortUrl 长短链映射
type ShortUrl struct {
	ID        string `json:"id" gorm:"column:id;primaryKey" bson:"_id"`          // 主键
	Short     string `json:"short" gorm:"column:short" bson:"short"`             // 短链
	Long      string `json:"long" gorm:"column:long" bson:"long"`                // 长链
	Type      int8   `json:"type" gorm:"column:type" bson:"type"`                // 类型
	Timestamp int64  `json:"timestamp" gorm:"column:timestamp" bson:"timestamp"` // 添加时间
}
