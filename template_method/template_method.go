package template_method

import "fmt"

type Template interface {
	CommonStep()
	SpecificStep()
}

type Parent struct{}

func (p *Parent) CommonStep() {
	fmt.Println("Begin common step...")
}

func (p *Parent) SpecificStep() {

}

func Excute(template Template) {
	template.CommonStep()
	template.SpecificStep()
}

type Child1 struct {
	Parent
}

func (c *Child1) SpecificStep() {
	fmt.Println("Begin specific step...")
}

type Child2 struct {
	Parent
}

func (c *Child2) SpecificStep() {
	fmt.Println("Begin specific step...")
}
