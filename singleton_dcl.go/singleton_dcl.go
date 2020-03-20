package singleton_dcl

import "sync"

type Singleton struct{}

var instance *Singleton
var mux sync.Mutex

func GetSingleton() *Singleton {
	// 单例没被初始化才加锁
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()

		// 单例没被初始化才创建
		if instance == nil {
			instance = &Singleton{}
		}
	}

	return instance
}
