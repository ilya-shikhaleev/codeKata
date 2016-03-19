package observer

import (
	"testing"
)

var countUpdates int
var subject Subject

type MockObserver struct{}

func (this *MockObserver) Update(data interface{}) {
	countUpdates++
}

type KillSelfObserver struct{}

func (this *KillSelfObserver) Update(data interface{}) {
	countUpdates++
	subject.RemoveObserver(this)
}

func TestObserver(t *testing.T) {
	mockObserver := &MockObserver{}
	subject.RegisterObserver(mockObserver, 1.0)
	subject.RegisterObserver(mockObserver, 1.0)
	subject.RegisterObserver(mockObserver, 1.0)

	killSelfObserver := &KillSelfObserver{}
	subject.RegisterObserver(killSelfObserver, 1.0)
	subject.RegisterObserver(killSelfObserver, 1.0)
	subject.RegisterObserver(killSelfObserver, 1.0)

	subject.NotifyObservers(nil)
	expectedCount := 6
	if countUpdates != expectedCount {
		t.Errorf("Expected %v observer calls, got %v", expectedCount, countUpdates)
	}

	subject.NotifyObservers(nil)
	expectedCount = 9
	if countUpdates != expectedCount {
		t.Errorf("Expected %v observer calls, got %v", expectedCount, countUpdates)
	}

	subject.RemoveObserver(mockObserver)
	subject.NotifyObservers(nil)
	expectedCount = 11
	if countUpdates != expectedCount {
		t.Errorf("Expected %v observer calls, got %v", expectedCount, countUpdates)
	}
}
