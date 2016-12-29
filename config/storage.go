package config

import conf "github.com/nicholaskh/jsconf"

const (
	StorageTypeMemory = "memory"
)

type StorageConfig struct {
	StorageType string
}

func (this *StorageConfig) LoadConfig(cf *conf.Conf) {
	this.StorageType = cf.String("storage_type", StorageTypeMemory)
}
