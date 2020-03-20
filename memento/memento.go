package memento

type Memento struct {
	state string
}

func (m *Memento) SetState(state string) {
	m.state = state
}

func (m *Memento) GetState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) WriteToMemento() *Memento {
	return &Memento{
		state: o.state,
	}
}

type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) SetMemento(memento *Memento) {
	c.memento = memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}
