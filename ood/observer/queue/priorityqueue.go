package queue

type QueueItem interface {
	GetValue() interface{}
	GetPriority() float64
}

type PriorityQueue []QueueItem

func (this *PriorityQueue) Push(item QueueItem) {
	var result PriorityQueue
	copy := *this
	var inserted bool

	for index := 0; index < len(copy); index++ {
		currentItem := copy[index]
		if !inserted && currentItem.GetPriority() < item.GetPriority() {
			result = append(result, item)
			inserted = true
		}
		result = append(result, currentItem)
	}
	if !inserted {
		result = append(result, item)
	}

	*this = result
}

func (this *PriorityQueue) Remove(item QueueItem) {
	var result PriorityQueue
	copy := *this

	for index := 0; index < len(*this); index++ {
		currentItem := copy[index]
		if currentItem != item {
			result = append(result, currentItem)
		}
	}

	*this = result
}
