package observer

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer_channels/queue"
)

type Event int

type Handler interface {
	Handle(data interface{})
}

func StartSubscribing(inChan chan interface{}, handler Handler) {
	go func() {
		for {
			select {
			case data, ok := <-inChan:
				if !ok {
					inChan = nil
				}
				handler.Handle(data)
			}
			if inChan == nil {
				return
			}
		}
	}()
}

type Subject struct {
	observers queue.PriorityQueue
}

func (this *Subject) RegisterObserver(inChan chan interface{}, priority float64) {
	this.observers.Push(inChan, priority)
}

func (this *Subject) RemoveObserver(inChan chan interface{}) {
	this.observers.Remove(inChan)
}

func (this *Subject) NotifyObservers(data interface{}) {
	for _, item := range this.observers {
		inChan := item.Value.(chan interface{})
		inChan <- data
	}
}
