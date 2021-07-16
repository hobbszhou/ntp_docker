package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
)

type NtpConfig struct {
	NtpConfigFile string
	NtpScript     string
}
type Config struct {
	rest.RestConf
	LogConf *logx.LogConf
	Ntp     *NtpConfig
}
