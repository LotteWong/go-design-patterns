package facade

import (
	"fmt"
)

type Facade interface {
	Start()
	End()
}

type RealFacade1 struct {
	Subsystem1
	Subsystem2
	Subsystem3
}

func (r *RealFacade1) Start() {
	fmt.Println("RealFacade1 starting...")
	r.Subsystem1.Up()
	r.Subsystem2.Up()
	r.Subsystem3.Up()
}

func (r *RealFacade1) End() {
	fmt.Println("RealFacade1 ending...")
	r.Subsystem1.Down()
	r.Subsystem2.Down()
	r.Subsystem3.Down()
}

type RealFacade2 struct {
	Subsystem1
	Subsystem2
	Subsystem3
}

func (r *RealFacade2) Start() {
	fmt.Println("RealFacade2 starting...")
	r.Subsystem1.Up()
	r.Subsystem2.Up()
	r.Subsystem3.Up()
}

func (r *RealFacade2) End() {
	fmt.Println("RealFacade2 ending...")
	r.Subsystem1.Down()
	r.Subsystem2.Down()
	r.Subsystem3.Down()
}

type Subsystem1 struct{}

func (s *Subsystem1) Up() {
	fmt.Println("Subsystem1 uping...")
}

func (s *Subsystem1) Down() {
	fmt.Println("Subsystem1 downing...")
}

type Subsystem2 struct{}

func (s *Subsystem2) Up() {
	fmt.Println("Subsystem2 uping...")
}

func (s *Subsystem2) Down() {
	fmt.Println("Subsystem2 downing...")
}

type Subsystem3 struct{}

func (s *Subsystem3) Up() {
	fmt.Println("Subsystem3 uping...")
}

func (s *Subsystem3) Down() {
	fmt.Println("Subsystem3 downing...")
}
