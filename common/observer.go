// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

// ******************** Interface ********************
type IObserver interface {
	ObserverNotifyReceived(tag string, data interface{})
}

type ISubject interface {
	Register(observer IObserver)
	Unregister(observer IObserver)
	Notify(tag string, data interface{})
}

// ******************** Struct ********************
type TSubject struct {
	observers []IObserver
}

func (s *TSubject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *TSubject) Unregister(observer IObserver) {
	var observers []IObserver
	for _, obs := range s.observers {
		if obs != observer {
			observers = append(observers, obs)
		}
	}
	s.observers = observers
}

func (s *TSubject) Notify(tag string, data interface{}) {
	for _, observer := range s.observers {
		observer.ObserverNotifyReceived(tag, data)
	}
}

// ******************** Var ********************
var (
	subjects = make(map[string]*TSubject, 0)
)

func GetSubject(id string) *TSubject {
	subject := subjects[id]
	if subject == nil {
		subject = &TSubject{}
		subjects[id] = subject
	}
	return subject
}
