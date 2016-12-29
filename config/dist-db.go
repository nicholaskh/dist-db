package config

import (
	"time"

	conf "github.com/nicholaskh/jsconf"
)

var (
	DistDb *DistDbConfig
)

type DistDbConfig struct {
	ListenAddr              string
	SessionTimeout          time.Duration
	ServInitialGoroutineNum int

	Storage *StorageConfig
}

func (this *DistDbConfig) LoadConfig(cf *conf.Conf) {
	this.ListenAddr = cf.String("listen_addr", ":8866")
	this.SessionTimeout = cf.Duration("session_timeout", time.Minute*2)
	this.ServInitialGoroutineNum = cf.Int("serv_initial_goroutine_num", 200)

	this.Storage = new(StorageConfig)
	section, err := cf.Section("storage")
	if err == nil {
		this.Storage.LoadConfig(section)
	}
}
