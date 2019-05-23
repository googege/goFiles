package main

import "fmt"

// 循环队列
type LoopQueue struct {
	data  []interface{}
	front int
	trail int
	len   int
	cap   int
}

func NewLoopQueue(cap int) *LoopQueue {
	l  :=  &LoopQueue{
		data:  make([]interface{}, cap, cap),
		front: 0,
		trail: 0,
		cap:   cap,
	}
	for i := 0;i < cap;i++ {
		l.data = append(l.data,"")
	}
	return l
}

func (l *LoopQueue) EnQueue(t interface{}) error {
	if (l.trail+1)%l.cap == l.front%l.cap {
		return fmt.Errorf("满了")
	}
	l.data[l.trail] = t
	l.trail = (l.trail + 1) % l.cap
	return nil
}

func (l *LoopQueue) DeQueue() error {
	if l.front == l.trail {
		return fmt.Errorf("没有东西")
	}
	l.data[l.front] = 0
	l.front = (l.front + 1) % l.cap
	return nil
}

func main() {
	l := NewLoopQueue(10)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.DeQueue()
	l.EnQueue(1)
	l.EnQueue(1)
	l.DeQueue()
	l.DeQueue()
	l.DeQueue()
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)

	fmt.Println(l.data)
}