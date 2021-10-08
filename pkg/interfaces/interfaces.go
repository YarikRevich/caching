package interfaces

type Cell struct {
	//Has two fields which allows to save elements in cache

	Key   interface{}
	Value interface{}
}

type ILRU interface {
	//Returns value gotten by the key shifting 
	//other elements
	Get(k interface{}) interface{} 

	//Returns the list of all values, but
	//does not shift
	GetAllWithoutShift() []interface{} 

	Set(v Cell)

	//Returns the len of the queue
	Len() int
}
