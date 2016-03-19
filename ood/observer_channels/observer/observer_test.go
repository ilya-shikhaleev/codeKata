package observer

import (
	"testing"
	"time"
)

var resultCh chan bool
var subject *Subject

type MockObserver struct {
	inChan chan *Event
}

func (this *MockObserver) Handle(data *Event) {
	resultCh <- true
}

func NewMockObserver() *MockObserver {
	o := &MockObserver{}
	o.inChan = make(chan *Event)
	StartSubscribing(o.inChan, o)
	return o
}

type KillSelfObserver struct {
	inChan chan *Event
}

func (this *KillSelfObserver) Handle(data *Event) {
	resultCh <- true
	subject.RemoveObserver(this.inChan)
}

func NewKillSelfObserver() *KillSelfObserver {
	o := &KillSelfObserver{}
	o.inChan = make(chan *Event)
	StartSubscribing(o.inChan, o)
	return o
}

func TestObserver(t *testing.T) {
	subject = NewSubject()
	resultCh = make(chan bool, 10)

	mockObserver := NewMockObserver()
	subject.RegisterObserver([]string{"first_event", "second_event"}, mockObserver.inChan)

	killSelfObserver := NewKillSelfObserver()
	subject.RegisterObserver([]string{"third_event", "second_event"}, killSelfObserver.inChan)

	timeout := 100 * time.Millisecond

	subject.NotifyObservers([]string{"first_event"}, nil)
	checkResultWithTimeout(1, timeout, t)

	subject.NotifyObservers([]string{"first_event", "second_event"}, nil)
	checkResultWithTimeout(3, timeout, t)

	subject.RegisterObserver([]string{"third_event", "second_event"}, killSelfObserver.inChan)
	subject.NotifyObservers([]string{"first_event", "second_event"}, nil)
	checkResultWithTimeout(3, timeout, t)

	subject.RegisterObserver([]string{"third_event", "second_event"}, killSelfObserver.inChan)
	subject.NotifyObservers([]string{"first_event", "second_event", "third_event"}, nil)
	checkResultWithTimeout(4, timeout, t)

	subject.RegisterObserver([]string{"third_event", "second_event"}, killSelfObserver.inChan)
	subject.RemoveObserverFromEvents([]string{"second_event"}, mockObserver.inChan)
	subject.NotifyObservers([]string{"first_event", "second_event", "third_event"}, nil)
	checkResultWithTimeout(3, timeout, t)
}

func checkResultWithTimeout(expectedCount int, timeout time.Duration, t *testing.T) {
	timeoutCh := time.After(timeout)
	var i int
	for i = 0; i < expectedCount; i++ {
		select {
		case <-resultCh:
			if i == expectedCount-1 {
				return
			}
		case <-timeoutCh:
			t.Errorf("After timeout got %v handles, expect %v", i, expectedCount)
			return
		}
	}
	t.Errorf("Got %v handles, expect %v", i, expectedCount)
}
