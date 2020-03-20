package decorator

type Component interface {
	Calc() int
}

type RealComponent struct {
	num int
}

func NewRealComponent(num int) *RealComponent {
	return &RealComponent{
		num: num,
	}
}

func (r *RealComponent) Calc() int {
	return r.num
}

type AddDecorator struct {
	com Component
	num int
}

func WrapAddDecorator(com Component, num int) *AddDecorator {
	return &AddDecorator{
		com: com,
		num: num,
	}
}

func (a *AddDecorator) Calc() int {
	return a.com.Calc() + a.num
}
