package query

import "errors"

type Memtable struct {
	container map[string]string
}

func NewMemtable() *Memtable {
	return &Memtable{
		container: make(map[string]string),
	}
}

func (r *Memtable) get(key string) (string, error) {
	val, exists := r.container[key]
	if !exists {
		return "", errors.New("Key does not exist")
	}
	return val, nil
}

func (r *Memtable) set(key string, value string) {
	r.container[key] = value
}

func (r *Memtable) del(key string) {
	delete(r.container, key)
}
