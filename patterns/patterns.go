package main

import (
	"fmt"
	"sync"
)

// === Singleton ===
type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

// === Factory ===
type Shape interface {
	Draw()
}

type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Рисую круг")
}

type Square struct{}

func (s Square) Draw() {
	fmt.Println("Рисую квадрат")
}

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}

// === Strategy ===
type Strategy interface {
	Execute(a, b int) int
}

type Add struct{}

func (Add) Execute(a, b int) int {
	return a + b
}

type Multiply struct{}

func (Multiply) Execute(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

// === Observer ===
type Observer interface {
	Update(data string)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type ConcreteSubject struct {
	observers []Observer
	data      string
}

func (s *ConcreteSubject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) RemoveObserver(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.data)
	}
}

func (s *ConcreteSubject) SetData(data string) {
	s.data = data
	s.NotifyObservers()
}

type ConcreteObserver struct {
	id string
}

func (o ConcreteObserver) Update(data string) {
	fmt.Printf("Observer %s получил обновление: %s\n", o.id, data)
}

// === Builder ===
type Product struct {
	Part1 string
	Part2 string
}

type Builder interface {
	SetPart1()
	SetPart2()
	GetResult() Product
}

type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) SetPart1() {
	b.product.Part1 = "Часть 1"
}

func (b *ConcreteBuilder) SetPart2() {
	b.product.Part2 = "Часть 2"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.SetPart1()
	d.builder.SetPart2()
}

// === Main ===
func main() {
	// Singleton
	fmt.Println("=== Singleton ===")
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println("Одинаковые ли ссылки?", s1 == s2)

	// Factory
	fmt.Println("\n=== Factory ===")
	shape1 := ShapeFactory("circle")
	shape1.Draw()
	shape2 := ShapeFactory("square")
	shape2.Draw()

	// Strategy
	fmt.Println("\n=== Strategy ===")
	add := Add{}
	multiply := Multiply{}
	context := Context{}
	context.SetStrategy(add)
	fmt.Println("Сложение:", context.ExecuteStrategy(10, 5))
	context.SetStrategy(multiply)
	fmt.Println("Умножение:", context.ExecuteStrategy(10, 5))

	// Observer
	fmt.Println("\n=== Observer ===")
	subject := &ConcreteSubject{}
	observer1 := &ConcreteObserver{id: "1"}
	observer2 := &ConcreteObserver{id: "2"}
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)
	subject.SetData("Новое состояние 1")
	subject.SetData("Новое состояние 2")

	// Builder
	fmt.Println("\n=== Builder ===")
	builder := &ConcreteBuilder{}
	director := &Director{builder: builder}
	director.Construct()
	product := builder.GetResult()
	fmt.Println("Продукт:", product)
}
