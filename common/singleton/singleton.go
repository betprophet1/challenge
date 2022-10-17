package singleton

import (
	"errors"
	"sync"
)

type Singleton struct {
	mux      sync.Mutex
	instance interface{}
	loader   func() interface{}
}

func New(loader func() interface{}) *Singleton {
	return &Singleton{
		loader: loader,
	}
}

func (s *Singleton) Get() interface{} {
	if s.instance != nil {
		return s.instance
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.instance == nil {
		value := s.loader()
		if value == nil {
			panic(errors.New("singleton: loader returns nil value"))
		}
		s.instance = value
	}
	return s.instance
}
