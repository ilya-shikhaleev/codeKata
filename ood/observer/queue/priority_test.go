package queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	var queue PriorityQueue

	for i := 0; i < 10; i++ {
		queue.Push(i, float64(i%2))
	}

	checkMultipleOfTwo(queue, t)

	checkRemoveItem(queue, t)
}

func checkMultipleOfTwo(queue PriorityQueue, t *testing.T) {
	for i := 0; i < 5; i++ {
		item := queue[i]
		intValue, ok := item.Value.(int)
		if !ok || intValue%2 == 0 {
			t.Errorf("Push invalid into queue. Value %v is a multiple of two. %v", intValue, queue)
		}
	}
}

func checkRemoveItem(queue PriorityQueue, t *testing.T) {
	removedItemValue := 1
	queue.Remove(removedItemValue)
	for _, item := range queue {
		if item.Value == removedItemValue {
			t.Errorf("Item %v isn't removed from queue", removedItemValue)
		}
	}

	insertingCount := 10
	for i := 0; i < insertingCount; i++ {
		queue.Push(removedItemValue, float64(insertingCount))
	}
	queue.Remove(removedItemValue)

	survivedCount := 0
	for _, item := range queue {
		if item.Value == removedItemValue {
			survivedCount++
		}
	}

	if survivedCount != insertingCount-1 {
		t.Errorf("After remove expect %v items, got %v items", insertingCount-1, survivedCount)
	}
}
