package cmd

func (this *cmd) cmdGet(key string) (string, error) {
	if value, ok := this.storage.Get(key); ok {
		return value, nil
	} else {
		return "", ErrKeyNotFound
	}
}

func (this *cmd) cmdSet(key string, value string) error {
	if v, ok := this.storage.Get(key); ok && v == value {
		return ErrKeyNotChanged
	}
	this.storage.Set(key, value)
	return nil
}
