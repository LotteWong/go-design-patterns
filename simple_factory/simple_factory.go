package simple_factory

import "fmt"

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Produce(name string) Product {
	switch name {
	case "p1":
		return Product1{}
	case "p2":
		return Product2{}
	default:
		return nil
	}
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

func TestCase() {
	factory := NewFactory()

	p1 := factory.Produce("p1")
	p1.create()

	p2 := factory.Produce("p2")
	p2.create()
}
