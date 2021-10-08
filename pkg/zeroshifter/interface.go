package zeroshifter

//Queue which shifts old values
//to left having been added new one
type IZeroShifter interface {
	Add(interface{})

	Get() []interface{} 

	Filter(func(interface{})bool)

	IsExist(func(interface{})bool)(interface{}, bool)

	Clean()
}