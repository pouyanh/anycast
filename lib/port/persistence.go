package port

// Key/Value Storage
type Dictionary interface {
	Put(key string, value interface{}) error
	Has(key string) bool
	Get(key string) interface{}
	Delete(key string) error
	GetDefault(key string, value interface{}) interface{}
}
