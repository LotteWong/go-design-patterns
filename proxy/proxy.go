package proxy

import "fmt"

type Subject interface {
	Do()
}

type RealSubject struct{}

func (r *RealSubject) Do() {
	fmt.Printf("working...")
}

type Proxy struct {
	real RealSubject
}

func (p *Proxy) Do() {
	fmt.Printf("before working..")

	p.real.Do()

	fmt.Printf("after working...")
}
