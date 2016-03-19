package observer

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer/queue"
)

//Observers
type Observer interface {
	Update(data interface{})
}

//Observable subject
type Observable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(data interface{})
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
	for _, item := range this.observers {
		observer := item.Value.(Observer)
		observer.Update(data)
	}
}