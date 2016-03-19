package observer

type Event struct {
	Data interface{}
	Type string
}

type Handler interface {
	Handle(event *Event)
}

func StartSubscribing(inChan chan *Event, handler Handler) {
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
	Observers map[string][]chan *Event
}

func (this *Subject) RegisterObserver(events []string, inChan chan *Event) {
	for _, event := range events {
		this.Observers[event] = append(this.Observers[event], inChan)
	}
}

func (this *Subject) RemoveObserver(inChan chan *Event) {
	for event := range this.Observers {
		var newObservers []chan *Event
		for _, ch := range this.Observers[event] {
			if ch != inChan {
				newObservers = append(newObservers, ch)
			}
		}
		this.Observers[event] = newObservers
	}
}

func (this *Subject) RemoveObserverFromEvents(events []string, inChan chan *Event) {
	for _, event := range events {
		var newObservers []chan *Event
		for _, ch := range this.Observers[event] {
			if ch != inChan {
				newObservers = append(newObservers, ch)
			}
		}
		this.Observers[event] = newObservers
	}
}

func (this *Subject) NotifyObservers(events []string, data interface{}) {
	for _, event := range events {
		e := &Event{data, event}
		for _, ch := range this.Observers[event] {
			ch <- e
		}
	}
}

func NewSubject() *Subject {
	s := &Subject{}
	s.Observers = make(map[string][]chan *Event)
	return s
}
