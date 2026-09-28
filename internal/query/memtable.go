package query

import (
	"errors"

	"github.com/tomciusromcius/distributed-store/internal/query/types"
)

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

func (r *Memtable) PopulateFromChannel(channel chan types.LogEntry) {
	for msg := range channel {
		if msg.Operation == types.OperationAdd {
			r.set(msg.Key, msg.Val)
		}
		if msg.Operation == types.OperationRemove {
			r.del(msg.Key)
		}
	}
}

func (r *Memtable) DumpToStream(channel chan *types.Pair) {
	defer close(channel)
	for key, val := range r.container {
		channel <- &types.Pair{Key: key, Val: val}
	}
}
