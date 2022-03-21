package hashtable

type HashTableInterface interface {
	Set(k interface{}, value interface{})
	Remove(k interface{}) error
	Get(k interface{}) interface{}
}
