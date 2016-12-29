package storage

type MemoryStorage map[string]string

func NewMemoryStorage() *MemoryStorage {
	this := make(MemoryStorage)
	return &this
}

func (this *MemoryStorage) Get(key string) (value string, exists bool) {
	value, exists = (*this)[key]
	return
}

func (this *MemoryStorage) Set(key string, value string) (succ bool) {
	(*this)[key] = value
	succ = true
	return
}
