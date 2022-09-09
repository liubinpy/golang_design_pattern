package singleton

// 懒汉式单例模式

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazySingleton() *Singleton {
	if lazySingleton != nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}

	return lazySingleton
}
