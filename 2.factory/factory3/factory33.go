package main

import "fmt"

// 抽象层
type AbstractGraphicsCard interface {
	display()
}

type AbstractMemory interface {
	storage()
}

type AbstractCPU interface {
	calculate()
}

type AbstractFactory interface {
	CreateGraphicsCard() AbstractGraphicsCard
	CreateMemory() AbstractMemory
	CreateCPU() AbstractCPU
}

// 实现层
// intel厂
type IntelGraphicsCard struct {
}

func (i IntelGraphicsCard) display() {
	fmt.Println("intelGraphicsCard...")
}

type IntelMemory struct {
}

func (i IntelMemory) storage() {
	fmt.Println("intelmemory...")
}

type IntelCPU struct {
}

func (i IntelCPU) calculate() {
	fmt.Println("intelCPU...")
}

type IntelFactory struct{}

func (i IntelFactory) CreateGraphicsCard() AbstractGraphicsCard {
	var graphicsCard AbstractGraphicsCard = new(IntelGraphicsCard)
	return graphicsCard
}

func (i IntelFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory = new(IntelMemory)
	return memory
}

func (i IntelFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU = new(IntelCPU)
	return cpu
}

// nvidia厂
type NvidiaGraphicsCard struct {
}

func (i NvidiaGraphicsCard) display() {
	fmt.Println("nvidiaGraphicsCard...")
}

type NvidiaMemory struct {
}

func (i NvidiaMemory) storage() {
	fmt.Println("nvidiamemory...")
}

type NvidiaCPU struct {
}

func (i NvidiaCPU) calculate() {
	fmt.Println("nvidiaCPU...")
}

type NvidiaFactory struct{}

func (i NvidiaFactory) CreateGraphicsCard() AbstractGraphicsCard {
	var graphicsCard AbstractGraphicsCard = new(NvidiaGraphicsCard)
	return graphicsCard
}

func (i NvidiaFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory = new(NvidiaMemory)
	return memory
}

func (i NvidiaFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU = new(NvidiaCPU)
	return cpu
}

// Kingston厂
type KingstonGraphicsCard struct {
}

func (i KingstonGraphicsCard) display() {
	fmt.Println("KingstonGraphicsCard...")
}

type KingstonMemory struct {
}

func (i KingstonMemory) storage() {
	fmt.Println("Kingstonmemory...")
}

type KingstonCPU struct {
}

func (i KingstonCPU) calculate() {
	fmt.Println("KingstonCPU...")
}

type KingstonFactory struct{}

func (i KingstonFactory) CreateGraphicsCard() AbstractGraphicsCard {
	var graphicsCard AbstractGraphicsCard = new(KingstonGraphicsCard)
	return graphicsCard
}

func (i KingstonFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory = new(KingstonMemory)
	return memory
}

func (i KingstonFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU = new(KingstonCPU)
	return cpu
}

// 计算机
type AbstractComputer interface {
	Run()
}

type AbstractComputerFacory interface {
	CreateComputer(AbstractGraphicsCard, AbstractMemory, AbstractCPU) AbstractComputer
}

type Computer struct {
	graphicsCard AbstractGraphicsCard
	memory       AbstractMemory
	cpu          AbstractCPU
}

func (c Computer) Run() {
	c.graphicsCard.display()
	c.memory.storage()
	c.cpu.calculate()
}

type ComputerFactory struct {
}

func (cf ComputerFactory) CreateComputer(g AbstractGraphicsCard, a AbstractMemory, c AbstractCPU) AbstractComputer {
	var computer AbstractComputer = Computer{
		graphicsCard: g,
		memory:       a,
		cpu:          c,
	}
	return computer
}

// 逻辑层
func main() {
	// 1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	// 1-需要Inter的CPU
	// 首先需要一个Inter的工厂
	interlfactory := IntelFactory{}
	// 创建一个intel的cpu
	cpu := interlfactory.CreateCPU()
	// 2-需要一个nvidia的显卡
	nvidiaFactory := NvidiaFactory{}
	graphicsCard := nvidiaFactory.CreateGraphicsCard()
	// 3-需要i个Kingston的内存
	kinstonFactory := KingstonFactory{}
	memory := kinstonFactory.CreateMemory()
	// 4-需要一个工厂生产计算机
	computerfactory := ComputerFactory{}
	computer := computerfactory.CreateComputer(graphicsCard, memory, cpu)

	computer.Run()

}
