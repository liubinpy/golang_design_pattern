package singleton

// 饿汉式单例模式

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
