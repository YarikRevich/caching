package lru

import "container/list"

type lru struct {
	list.List
	capacity int //maximum capacity of lru
}

func (l *lru) Get(k interface{}) interface{} {
	for f := l.Front(); f != nil; f = f.Next() {
		if x := f.Value.(Cell); x.Key == k {
			l.MoveToFront(f)
			return x.Value
		}
	}

	return nil
}

func (l *lru) Set(v Cell) {
	if l.Len() >= l.capacity {
		l.Remove(l.Back())
	}
	l.PushFront(v)
}

func (l *lru) QueueLen() int{
	return l.Len()
}

func New(capacity int) ILRU {
	return &lru{
		capacity: capacity,
	}

}
