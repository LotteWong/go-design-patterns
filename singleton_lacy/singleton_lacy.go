package singleton_lacy

type Singleton struct{}

// 私有且单例
var instance *Singleton

func GetSingleton() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
