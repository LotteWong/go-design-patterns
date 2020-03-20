package prototype

type Cloneable interface {
	Clone() Cloneable
}

type Example struct{}

func (e *Example) Clone() Cloneable {
	copy := *e
	return &copy
}
