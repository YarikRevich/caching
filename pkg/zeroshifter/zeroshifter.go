package zeroshifter

import "container/ring"

type zeroshifter struct {
	ring     *ring.Ring
	capacity int
}

func (z *zeroshifter) Add(v interface{}) {
	z.ring.Value = v
	z.ring = z.ring.Next()
}

func (z *zeroshifter) Get() []interface{} {
	r := []interface{}{}
	z.ring.Do(func(i interface{}) {
		if i != nil {
			r = append(r, i)
		}
	})
	return r
}

func (z *zeroshifter) Filter(f func(i interface{}) bool) {
	for i := 0; i < z.capacity; i++ {
		v := z.ring.Value
		if v != nil {
			if !f(v) {
				z.ring.Value = nil
			}
		}
		z.ring = z.ring.Next()
	}
}

func (z *zeroshifter) Clean() {
	z.ring = ring.New(z.capacity)
}

func (z *zeroshifter) IsExist(c func(i interface{})bool)bool{
	var r bool
	z.ring.Do(func(i interface{}) {
		if i != nil{
			if c(i){
				r = true
			}
		}
	})
	return r
}

func New(capacity int) IZeroShifter {
	return &zeroshifter{
		ring:     ring.New(capacity),
		capacity: capacity,
	}
}
