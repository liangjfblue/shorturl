package constinfo

const (
	// 环境变量
	EnvDev = "dev"
	EnvPro = "pro"

	// 短链类型
	TypeSystem = "system"
	TypeCustom = "custom"
)

const (
	KeyShortUrl   = "short:url:"  // 热门短链
	KeyShortBloom = "short:bloom" // 短链集合, 用于快速判断短链是否存在
)
