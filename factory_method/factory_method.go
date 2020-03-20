package factory_method

import "fmt"

type Factory interface {
	Produce(p string) Product
}

type Product interface {
	create()
}

type Factory1 struct{}

func NewFactory1() *Factory1 {
	return &Factory1{}
}

func (f Factory1) Produce(p string) Product {
	switch p {
	case "p1":
		return Product1{}
	default:
		return nil
	}
}

type Factory2 struct{}

func NewFactory2() *Factory2 {
	return &Factory2{}
}

func (f Factory2) Produce(p string) Product {
	switch p {
	case "p2":
		return Product2{}
	default:
		return nil
	}
}

type Product1 struct{}

func (p Product1) create() {
	fmt.Printf("create p1...")
}

type Product2 struct{}

func (p Product2) create() {
	fmt.Printf("create p2...")
}

func TestCase() {
	factory1 := NewFactory1()
	factory2 := NewFactory2()

	p1 := factory1.Produce("p1")
	p1.create()

	p2 := factory2.Produce("p2")
	p2.create()
}
