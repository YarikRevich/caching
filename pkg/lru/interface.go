package lru

type Cell struct {
	Key   interface{}
	Value interface{}
}

type ILRU interface {
	Get(key interface{}) interface{} 

	Set(c Cell)

	QueueLen() int
}
