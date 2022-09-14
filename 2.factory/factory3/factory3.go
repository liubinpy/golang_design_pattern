package main

// import "fmt"

// // 抽象工厂方法模式
// // ==抽象层==
// type AbstractApple interface {
// 	ShowApple()
// }

// type AbstractBanana interface {
// 	ShowBanana()
// }

// type AbstractPear interface {
// 	ShowPear()
// }

// // 抽象工厂
// type AbstractFactory interface {
// 	CreateApple() AbstractApple
// 	CreateBanana() AbstractBanana
// 	CreatePear() AbstractPear
// }

// // == 实现层 ==
// // 中国产品族
// type ChinaApple struct{}

// func (ca *ChinaApple) ShowApple() {
// 	fmt.Println("中国苹果")
// }

// type ChinaBanana struct{}

// func (cb *ChinaBanana) ShowBanana() {
// 	fmt.Println("中国香蕉")
// }

// type ChinaPear struct{}

// func (cp *ChinaPear) ShowPear() {
// 	fmt.Println("中国梨")
// }

// type ChinaFactory struct{}

// func (cf *ChinaFactory) CreateApple() AbstractApple {
// 	var apple AbstractApple = new(ChinaApple)

// 	return apple
// }

// func (cf *ChinaFactory) CreateBanana() AbstractBanana {

// 	var banana AbstractBanana = new(ChinaBanana)

// 	return banana
// }

// func (cf *ChinaFactory) CreatePear() AbstractPear {

// 	var pear AbstractPear = new(ChinaPear)

// 	return pear
// }

// // 日本产品族
// type JapanApple struct{}

// func (ja *JapanApple) ShowApple() {
// 	fmt.Println("日本苹果")
// }

// type JapanBanana struct{}

// func (jb *JapanBanana) ShowBanana() {
// 	fmt.Println("日本香蕉")
// }

// type JapanPear struct{}

// func (cp *JapanPear) ShowPear() {
// 	fmt.Println("日本梨")
// }

// type JapanFactory struct{}

// func (jf *JapanFactory) CreateApple() AbstractApple {

// 	var apple AbstractApple = new(JapanApple)

// 	return apple
// }

// func (jf *JapanFactory) CreateBanana() AbstractBanana {

// 	var banana AbstractBanana = new(JapanBanana)

// 	return banana
// }

// func (cf *JapanFactory) CreatePear() AbstractPear {

// 	var pear AbstractPear = new(JapanPear)

// 	return pear
// }

// // 业务逻辑层
// func main() {
// 	//需要中国的苹果、香蕉
// 	//1-创建一个中国工厂
// 	cFac := new(ChinaFactory)

// 	//2-生产中国苹果
// 	cApple := cFac.CreateApple()
// 	cApple.ShowApple()

// 	//3-生产中国香蕉
// 	cBanana := cFac.CreateBanana()
// 	cBanana.ShowBanana()
// }
