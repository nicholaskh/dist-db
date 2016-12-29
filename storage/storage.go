package storage

import (
	"github.com/nicholaskh/dist-db/config"
)

var (
	memoryStorage *MemoryStorage
)

type Storage interface {
	Get(string) (string, bool)
	Set(string, string) bool
}

func Factory(storageType string) Storage {
	switch storageType {
	case config.StorageTypeMemory:
		if memoryStorage == nil {
			memoryStorage = NewMemoryStorage()
		}
		return memoryStorage
	default:
		return nil
	}
}
