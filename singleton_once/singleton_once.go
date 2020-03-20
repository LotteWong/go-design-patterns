package singleton_once

import "sync"

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetSingleton() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}
