package lru

import (
	"container/list"

	"github.com/YarikRevich/lru/pkg/interfaces"
)

type LRU struct {
	list     *list.List
	capacity int //maximum capacity of lru
}

func (l LRU) Get(k interface{}) interface{} {
	if l.list.Len() > 1 {
		for f := l.list.Front(); f != nil; f = f.Next() {
			if x := f.Value.(interfaces.Cell); x.Key == k {
				l.list.MoveToFront(f)
				return x.Value
			}
		}
	}

	if l.list.Len() == 1 {
		f := l.list.Front()
		if f != nil {

			x, ok := f.Value.(interfaces.Cell)
			if !ok {
				return nil
			}

			if x.Key == k {
				l.list.Remove(f)
				return x.Value
			}
		}

	}

	return nil
}

func (l LRU) GetAllWithoutShift()[]interface{}{
	r := make([]interface{}, l.Len()-1)
	for f := l.list.Front(); f != nil; f = f.Next(){
		r = append(r, f.Value.(interfaces.Cell).Value)
	}
	return r
}

func (l LRU) Set(v interfaces.Cell) {
	if l.list.Len() >= l.capacity {
		l.list.Remove(l.list.Back())
	}
	l.list.PushFront(v)
}

func (l LRU) Len() int {
	return l.list.Len()
}

func New(capacity int) interfaces.ILRU {
	c := new(LRU)
	c.capacity = capacity
	c.list = list.New()
	return c
}
