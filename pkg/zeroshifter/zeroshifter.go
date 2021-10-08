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

func New(capacity int) IZeroShifter {
	return &zeroshifter{
		ring:     ring.New(capacity),
		capacity: capacity,
	}
}

// if l.list.Len() == 1 {
// 	f := l.list.Front()
// 	if f != nil {

// 		x, ok := f.Value.(interfaces.Cell)
// 		if !ok {
// 			return nil
// 		}

// 		if x.Key == k {
// 			l.list.Remove(f)
// 			return x.Value
// 		}
// 	}
// }

// func (l LRU) GetAllWithoutShift() []interface{} {
// 	if l.Len() != 0 {
// 		r := make([]interface{}, l.Len()-1)
// 		for f := l.list.Front(); f != nil; f = f.Next() {
// 			r = append(r, f.Value.(interfaces.Cell).Value)
// 		}
// 		return r
// 	}
// 	return nil
// }
