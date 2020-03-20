package visitor

import "fmt"

type Visitor interface {
	Visit()
}

type ProductionVisitor struct{}

func (p *ProductionVisitor) Visit() {
	fmt.Println("Production environment.")
}

type TestingVisitor struct{}

func (p *TestingVisitor) Visit() {
	fmt.Println("Testing environment.")
}

type Environment interface {
	Accept(visitor Visitor)
}

type RealEnvironment struct{}

func (r *RealEnvironment) Accept(visitor Visitor) {
	visitor.Visit()
}

type Element struct {
	Env Environment
}

func (e *Element) VisitEnv(visitor Visitor) {
	e.Env.Accept(visitor)
}
