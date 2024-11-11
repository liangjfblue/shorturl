package component

import (
	"common/logger"
	"common/utils"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewLogger, NewSnowFlake)

func NewLogger(logConf *logger.LogConfig) *logger.Logger {
	return logger.NewLogger(logConf)
}

func NewSnowFlake(snowFlakeConf *utils.SnowFlakeConf) *utils.SnowFlake {
	return utils.NewSnowFlake(snowFlakeConf)
}
