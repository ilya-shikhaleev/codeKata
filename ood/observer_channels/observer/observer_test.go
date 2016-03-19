package observer

import (
	"testing"
	"time"
)

var resultCh chan bool
var subject Subject

type MockObserver struct {
	inChan chan interface{}
}

func (this *MockObserver) Handle(data interface{}) {
	resultCh <- true
}

func NewMockObserver() *MockObserver {
	o := &MockObserver{}
	o.inChan = make(chan interface{})
	StartSubscribing(o.inChan, o)
	return o
}

type KillSelfObserver struct {
	inChan chan interface{}
}

func (this *KillSelfObserver) Handle(data interface{}) {
	resultCh <- true
	subject.RemoveObserver(this.inChan)
}

func NewKillSelfObserver() *KillSelfObserver {
	o := &KillSelfObserver{}
	o.inChan = make(chan interface{})
	StartSubscribing(o.inChan, o)
	return o
}

func TestObserver(t *testing.T) {
	resultCh = make(chan bool, 10)

	mockObserver := NewMockObserver()
	subject.RegisterObserver(mockObserver.inChan)
	subject.RegisterObserver(mockObserver.inChan)
	subject.RegisterObserver(mockObserver.inChan)

	killSelfObserver := NewKillSelfObserver()
	subject.RegisterObserver(killSelfObserver.inChan)
	subject.RegisterObserver(killSelfObserver.inChan)
	subject.RegisterObserver(killSelfObserver.inChan)

	subject.NotifyObservers(nil)
	timeout := 100 * time.Millisecond
	checkResultWithTimeout(6, timeout, t)

	subject.NotifyObservers(nil)
	checkResultWithTimeout(3, timeout, t)

	subject.RemoveObserver(mockObserver.inChan)
	subject.RegisterObserver(mockObserver.inChan)
	subject.NotifyObservers(nil)
	checkResultWithTimeout(1, timeout, t)
}

func checkResultWithTimeout(expectedCount int, timeout time.Duration, t *testing.T) {
	timeoutCh := time.After(timeout)
	for i := 0; i < expectedCount; i++ {
		select {
		case <-resultCh:
			if i == expectedCount {
				return
			}
		case <-timeoutCh:
			t.Errorf("After timeout got %v handles, expect %v", i, expectedCount)
			return
		}
	}
}
