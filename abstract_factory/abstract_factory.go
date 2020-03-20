package abstract_factory

import "fmt"

type Factory interface {
	ProduceOdd() Product
	ProduceEven() Product
}

type Factory1 struct{}

func NewFactory1() *Factory1 {
	return &Factory1{}
}

func (f Factory1) ProduceOdd() Product {
	return Product1{}
}

func (f Factory1) ProduceEven() Product {
	return Product2{}
}

type Factory2 struct{}

func NewFactory2() *Factory2 {
	return &Factory2{}
}

func (f Factory2) ProduceOdd() Product {
	return Product3{}
}

func (f Factory2) ProduceEven() Product {
	return Product4{}
}

type Product interface {
	create()
}

type Product1 struct{}

func (p Product1) create() {
	fmt.Printf("create p1...")
}

type Product2 struct{}

func (p Product2) create() {
	fmt.Printf("create p2...")
}

type Product3 struct{}

func (p Product3) create() {
	fmt.Printf("create p3...")
}

type Product4 struct{}

func (p Product4) create() {
	fmt.Printf("create p4...")
}

func TestCase() {
	factory1 := NewFactory1()
	factory2 := NewFactory2()

	p1 := factory1.ProduceOdd()
	p1.create()
	p2 := factory1.ProduceEven()
	p2.create()

	p3 := factory2.ProduceOdd()
	p3.create()
	p4 := factory2.ProduceEven()
	p4.create()
}
