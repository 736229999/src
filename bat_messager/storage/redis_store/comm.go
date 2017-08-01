
package redis_store

type Store interface {
	StoreKey() string
	StoreData() interface{}
}
