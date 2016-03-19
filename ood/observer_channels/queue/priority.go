package queue

import (
	"bytes"
	"fmt"
)

type QueueItem struct {
	Value    interface{}
	Priority float64
}

type PriorityQueue []*QueueItem

func (this *PriorityQueue) Push(value interface{}, priority float64) {
	copy := *this
	item := &QueueItem{value, priority}
	index := this.findIndexForInsert(priority)

	*this = append(copy[:index], append(PriorityQueue{item}, copy[index:]...)...)
}

//Removes first found item.
func (this *PriorityQueue) Remove(value interface{}) {
	index := this.firstPosition(value)
	if index >= 0 {
		copy := *this
		*this = append(copy[:index], copy[index+1:]...)
	}
}

func (this PriorityQueue) String() string {
	var buffer bytes.Buffer
	for _, item := range this {
		buffer.WriteString(fmt.Sprintf("%v|", item.Value))
	}
	return buffer.String()
}

func (this PriorityQueue) findIndexForInsert(priority float64) int {
	for index, item := range this {
		if item.Priority < priority {
			return index
		}
	}
	return len(this)
}

func (this PriorityQueue) firstPosition(value interface{}) int {
	for index, item := range this {
		if item.Value == value {
			return index
		}
	}
	return -1
}
