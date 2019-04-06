package cache

//PkgMemCache implementation of Cache Interface
type PkgMemCache struct {
}

// NewMemCache returns memcache
func NewMemCache() Cache {
	obj := &PkgMemCache{}
	return obj
}

//Get value from key
func (p *PkgMemCache) Get(key string) (value string, err error) {
	return
}

//Set value with key
func (p *PkgMemCache) Set(key string, value string, ttl ...int) (err error) {
	return
}

// Ping will health check
func (p *PkgMemCache) Ping() (err error) {
	return
}
