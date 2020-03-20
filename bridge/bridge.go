package bridge

import "fmt"

type Operation interface {
	Boot()
	Excute()
	Exit()
}

type OperationImpl struct{}

func (o *OperationImpl) Boot() {
	fmt.Println("Booting...")
}

func (o *OperationImpl) Excute() {
	fmt.Println("Excuting...")
}

func (o *OperationImpl) Exit() {
	fmt.Println("Exit...")
}

type Abstract struct {
	Operation
	id string
}

func (a *Abstract) Boot() {
	a.Operation.Boot()
	a.id = "abstract"
}

func (a *Abstract) Excute() {
	a.Operation.Excute()
	a.id = "abstract"
}

func (a *Abstract) Exit() {
	a.Operation.Exit()
	a.id = "abstract"
}
