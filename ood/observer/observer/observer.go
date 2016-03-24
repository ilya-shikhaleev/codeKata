package observer

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer/queue"
)

//Observers
type Observer interface {
	Update(data interface{})
}

type Subject struct {
	observers queue.PriorityQueue
}

func (this *Subject) RegisterObserver(observer Observer, priority float64) {
	this.observers.Push(observer, priority)
}

func (this *Subject) RemoveObserver(observer Observer) {
	this.observers.Remove(observer)
}

func (this *Subject) NotifyObservers(data interface{}) {
	copyObservers := this.observers.Copy()
	for _, item := range *copyObservers{
		observer := item.Value.(Observer)
		observer.Update(data)
	}
}
