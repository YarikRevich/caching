package logic

import (
	"container/list"

	"github.com/YarikRevich/lru/internal/interfaces"
)

type LRU struct {
	list     *list.List
	capacity int //maximum capacity of lru
}

func (l LRU) Get(k interface{}) interface{} {
	for f := l.list.Front(); f != nil; f = f.Next() {
		if x := f.Value.(interfaces.Cell); x.Key == k {
			l.list.MoveToFront(f)
			return x.Value
		}
	}
	return nil
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
