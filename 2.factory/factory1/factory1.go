package factory1

import "fmt"

// 简单工厂
// == 抽象层 ==
type Fruit interface {
	show()
}

// == 基础类模块 ==
type Apple struct{}

func (a Apple) show() {
	fmt.Println("我是苹果")
}

type Banana struct{}

func (b Banana) show() {
	fmt.Println("我是香蕉")
}

type Pear struct{}

func (p Pear) show() {
	fmt.Println("我是梨子")
}

// == 工厂模块 ==
//一个工厂， 有一个生产水果的机器，返回一个抽象水果的指针
type factory struct{}

func (f factory) CreateFruit(kind string) Fruit {
	switch kind {
	case "apple":
		return &Apple{}
	case "banana":
		return &Banana{}
	case "pear":
		return &Pear{}
	}
	return nil
}
