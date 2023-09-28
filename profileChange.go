package logger

import (
	"sync"
)

var globalProfileChangeSubject *profileChangeSubject

func init() {
	globalProfileChangeSubject = NewProfileChangeSubject()
}

// SubscribeToProfileChange subscribes an observer
func SubscribeToProfileChange(observer ProfileChangeObserver) {
	globalProfileChangeSubject.subscribe(observer)
}

// UnsubscribeFromProfileChange unsubscribes an observer
func UnsubscribeFromProfileChange(observer ProfileChangeObserver) {
	globalProfileChangeSubject.unsubscribe(observer)
}

// NotifyProfileChange notifies observers about a profile change
func NotifyProfileChange() {
	globalProfileChangeSubject.NotifyAll()
}

type profileChangeSubject struct {
	observers []ProfileChangeObserver
	mutex     sync.RWMutex
}

// NewProfileChangeSubject -
func NewProfileChangeSubject() *profileChangeSubject {
	return &profileChangeSubject{
		observers: make([]ProfileChangeObserver, 0),
	}
}

func (subject *profileChangeSubject) subscribe(observer ProfileChangeObserver) {
	subject.mutex.Lock()
	subject.observers = append(subject.observers, observer)
	subject.mutex.Unlock()
}

func (subject *profileChangeSubject) unsubscribe(observer ProfileChangeObserver) {
	subject.mutex.Lock()
	defer subject.mutex.Unlock()

	for i := 0; i < len(subject.observers); i++ {
		if subject.observers[i] == observer {
			subject.observers = append(subject.observers[0:i], subject.observers[i+1:]...)
		}
	}
}

// Notifies all observers that the profile changed
func (subject *profileChangeSubject) NotifyAll() {
	subject.mutex.RLock()
	defer subject.mutex.RUnlock()

	for _, observer := range subject.observers {
		observer.OnProfileChanged()
	}
}
