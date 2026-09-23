package query

type KeyvalStore struct {
	Container map[string]string
}

func (r *KeyvalStore) set(key string, value string) {
	r.Container[key] = value
}
