package rpc

import (
	"HQ-Backend/common"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	DB struct {
		DSN  string
		Type string
	}
}

func MustInitConfig() *Config {
	var c Config
	conf.MustLoad(common.ConfigFile, &c)
	return &c
}
