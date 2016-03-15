package observer

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer/queue"
)

//Observers
type Observer interface {
	Update(data interface{})
}

type PriorityObserver struct {
	Observer Observer
	Priority float64
}

func (this PriorityObserver) GetValue() interface{} {
	return this.Observer
}

func (this PriorityObserver) GetPriority() float64 {
	return this.Priority
}

//Observable subject
type Observable interface {
	RegisterObserver(observer PriorityObserver)
	RemoveObserver(observer PriorityObserver)
	NotifyObservers(data interface{})
}

type Subject struct {
	observers queue.PriorityQueue
}

func (this *Subject) RegisterObserver(observer *PriorityObserver) {
	this.observers.Push(observer)
}

func (this *Subject) RemoveObserver(observer *PriorityObserver) {
	this.observers.Remove(observer)
}

func (this *Subject) NotifyObservers(data interface{}) {
	for _, item := range this.observers {
		observer := item.GetValue().(Observer)
		observer.Update(data)
	}
}
