// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

// ******************** Interface ********************
type IObserver interface {
	ObserverNotifyReceived(tag string, data interface{})
}

type ITopic interface {
	AddObserver(observer IObserver)
	RemoveObserver(observer IObserver)
	Notify(tag string, data interface{})
}

// ******************** Struct ********************
type TTopic struct {
	observers []IObserver
}

func (tc *TTopic) AddObserver(observer IObserver) {
	tc.observers = append(tc.observers, observer)
}

func (tc *TTopic) RemoveObserver(observer IObserver) {
	var observers []IObserver
	for _, obs := range tc.observers {
		if obs != observer {
			observers = append(observers, obs)
		}
	}
	tc.observers = observers
}

func (tc *TTopic) Notify(tag string, data interface{}) {
	for _, observer := range tc.observers {
		observer.ObserverNotifyReceived(tag, data)
	}
}

// ******************** Var ********************
var (
	topicMap = make(map[string]*TTopic, 0)
)

func ObserverGetTopic(id string) *TTopic {
	topic := topicMap[id]
	if topic == nil {
		topic = &TTopic{}
		topicMap[id] = topic
	}
	return topic
}
