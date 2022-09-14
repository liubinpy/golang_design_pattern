package main

import "fmt"

// 工厂方法

// ==抽象层==

// 水果类，抽象层
type Fruit interface {
	Show()
}

// 工厂类，抽象工厂
type AbstractFactory interface {
	CreateFruit() Fruit
}

// ==基础类==
type Apple struct {
}

func (a Apple) Show() {
	fmt.Println("i am apple")
}

type Banana struct {
}

func (b Banana) Show() {
	fmt.Println("i am banana")
}

type Pear struct {
}

func (p Pear) Show() {
	fmt.Println("i am pear")
}

// == 工厂实现 ==
// 具体的苹果工厂
type AppleFactory struct{}

func (a AppleFactory) CreateFruit() Fruit {
	// 生产一个苹果
	var fruit Fruit = new(Apple)
	return fruit
}

// 具体的香蕉工厂
type BananaFactory struct{}

func (a BananaFactory) CreateFruit() Fruit {
	// 生产一个香蕉
	var fruit Fruit = new(Banana)
	return fruit
}

// 具体的梨子工厂
type PearFactory struct{}

func (a PearFactory) CreateFruit() Fruit {
	// 生产一个梨子
	var fruit Fruit = new(Pear)
	return fruit
}

// 业务实现层
func main() {
	// 需要一个苹果
	// 需要先创建一个苹果工厂
	appleFactory := AppleFactory{}
	// 苹果工厂生产一个苹果
	apple := appleFactory.CreateFruit()
	apple.Show()

	// 需要一个香蕉
	// 需要先创建一个香蕉工厂
	bananaFacory := BananaFactory{}
	// 香蕉工厂生产一个香蕉
	banana := bananaFacory.CreateFruit()
	banana.Show()
}
