package observer

import "fmt"

type Subject struct {
	observers []Observer // 观察者
	content   string     // 内容
}

func NewSubject(content string) *Subject {
	return &Subject{
		observers: make([]Observer, 0),
		content:   content,
	}
}

func (s *Subject) Attach(o Observer) { // 绑定观察者
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Subscribe(s)
	}
}

func (s *Subject) Publish(content string) { // 发布内容
	s.content = content
	s.notify()
}

type Observer interface {
	Subscribe(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Subscribe(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.content)
}
