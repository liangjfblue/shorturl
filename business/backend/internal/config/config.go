package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"sync"
)

var ProviderSet = wire.NewSet(NewConfig, NewLogConf, NewSnowFlakeConf)

// Config ...
type Config struct {
	Server     ServerConfig   `mapstructure:"server" yaml:"server"`
	Log        Logger         `mapstructure:"log" yaml:"log"`
	SnowFlake  SnowFlake      `mapstructure:"snowflake" yaml:"snowflake"`
	LocalCache LocalCache     `mapstructure:"localCache" yaml:"localCache"`
	Redis      RedisConfig    `mapstructure:"redis" yaml:"redis"`
	Database   DatabaseConfig `mapstructure:"database" yaml:"database"`
}

type ServerConfig struct {
	Env  string `mapstructure:"env"`
	Type string `mapstructure:"type"` // http/https
	Http struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"http" yaml:"http"`
	Https struct {
		Port     int    `mapstructure:"port"`
		CertFile string `mapstructure:"certFile"`
		KeyFile  string `mapstructure:"keyFile"`
	} `mapstructure:"https" yaml:"https"`
}

type Logger struct {
	Level      string `mapstructure:"level" yaml:"level"`
	Path       string `mapstructure:"path" yaml:"path"`
	MaxSize    int    `mapstructure:"maxSize" yaml:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups" yaml:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge" yaml:"maxAge"`
}

type SnowFlake struct {
	DataCenterId int64 `mapstructure:"dataCenterId" yaml:"dataCenterId"`
	MachineId    int64 `mapstructure:"machineId" yaml:"machineId"`
}

type LocalCache struct {
	Memory int `mapstructure:"memory" yaml:"memory"`
	Expire int `mapstructure:"expire" yaml:"expire"`
}

type RedisConfig struct {
	Type   string `mapstructure:"type" yaml:"type"` // single/sentinel/cluster
	Single struct {
		Host     string `mapstructure:"host" yaml:"host"`
		Password string `mapstructure:"password" yaml:"password"`
	} `mapstructure:"single"`
	Sentinel struct {
		MasterName string   `mapstructure:"masterName"`
		Hosts      []string `mapstructure:"hosts"`
		Password   string   `mapstructure:"password"`
	} `mapstructure:"sentinel"`
	Cluster struct {
		Hosts    []string `mapstructure:"hosts"`
		Password string   `mapstructure:"password"`
	} `mapstructure:"cluster"`
}

type DatabaseConfig struct {
	Type  string `mapstructure:"type" yaml:"type"` // mongodb/mysql
	Mongo struct {
		Host   string `mapstructure:"host"`
		DBName string `mapstructure:"dbName"`
	} `mapstructure:"mongo"`
	Mysql struct {
		Host string `mapstructure:"host"`
		// Port     int    `mapstructure:"port"`
		// Username string `mapstructure:"username"`
		// Password string `mapstructure:"password"`
	} `mapstructure:"db"`
}

var (
	doOnce sync.Once
)

func NewConfig() *Config {
	c := new(Config)
	c.init()
	return c
}

func (c *Config) init() {
	doOnce.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(c); err != nil {
			panic(err)
		}
	})
}
