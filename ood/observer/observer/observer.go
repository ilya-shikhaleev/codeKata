package observer

//Observable subject
type Observable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(data interface{})
}

type Subject struct {
	observers []Observer
}

func (this *Subject) RegisterObserver(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *Subject) RemoveObserver(observer Observer) {
	for index, currentObserver := range this.observers {
		if currentObserver == observer {
			this.observers = append(this.observers[:index], this.observers[index+1:]...)
		}
	}
}

func (this *Subject) NotifyObservers(data interface{}) {
	for _, observer := range this.observers {
		observer.Update(data)
	}
}

//Observers
type Observer interface {
	Update(data interface{})
}
