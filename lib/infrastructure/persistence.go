package infrastructure

type KeyValueStorage interface {
	Put(key string, value interface{}) error
	Has(key string) bool
	Get(key string) interface{}
	Delete(key string) error
	GetDefault(key string, value interface{}) interface{}
}
