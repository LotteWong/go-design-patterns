package adpater

import "fmt"

type Adapter interface {
	TargetRequest()
}

type Adaptee interface {
	AdapteeRequest()
}

type AdapteeImpl struct{}

func (a *AdapteeImpl) AdapteeRequest() {
	fmt.Println("Adaptee requesting...")
}

func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

type AdapterImpl struct {
	adaptee Adaptee
}

func (a *AdapterImpl) TargetRequest() {
	a.adaptee.AdapteeRequest()
}

func NewAdapter(adaptee Adaptee) Adapter {
	return &AdapterImpl{
		adaptee: adaptee,
	}
}
