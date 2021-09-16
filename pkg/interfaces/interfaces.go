package interfaces

type Cell struct {
	//Has two fields which allows to save elements in cache

	Key   interface{}
	Value interface{}
}

type ILRU interface {
	//Gets cell by the key
	Get(k interface{}) interface{} 

	//Sets cell
	Set(v Cell)

	//Returns the len of the container/list
	Len() int
}
