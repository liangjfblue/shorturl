package config

import (
	"common/logger"
	"common/utils"
)

// NewLogConf ...
func NewLogConf(c *Config) *logger.LogConfig {
	return &logger.LogConfig{
		Level:      c.Log.Path,
		Filename:   c.Log.Path,
		MaxSize:    c.Log.MaxSize,
		MaxBackups: c.Log.MaxBackups,
		MaxAge:     c.Log.MaxAge,
	}
}

// NewSnowFlakeConf ...
func NewSnowFlakeConf(c *Config) *utils.SnowFlakeConf {
	return &utils.SnowFlakeConf{
		DataCenterId: c.SnowFlake.DataCenterId,
		MachineId:    c.SnowFlake.MachineId,
	}
}
