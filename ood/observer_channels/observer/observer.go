package observer

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
	observers []chan interface{}
}

func (this *Subject) RegisterObserver(inChan chan interface{}) {
	this.observers = append(this.observers, inChan)
}

func (this *Subject) RemoveObserver(inChan chan interface{}) {
	var newObservers []chan interface{}
	for _, ch := range this.observers {
		if ch != inChan {
			newObservers = append(newObservers, ch)
		}
	}
	this.observers = newObservers
}

func (this *Subject) NotifyObservers(data interface{}) {
	for _, ch := range this.observers {
		ch <- data
	}
}
