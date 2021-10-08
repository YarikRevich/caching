package zeroshifter

//Queue which shifts old values
//to left having been added new one
type IZeroShifter interface {
	Add(interface{})

	Get() []interface{} 

	Filter(f func(i interface{})bool)

	Clean()
}