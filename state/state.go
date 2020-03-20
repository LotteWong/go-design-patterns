package state

import "fmt"

type State interface {
	Update()
	Next() State
}

type GameStartState struct{}

func (g *GameStartState) Update() {
	fmt.Println("Game starting...")
}

func (g *GameStartState) Next() State {
	return &GameRunState{}
}

type GameRunState struct{}

func (g *GameRunState) Update() {
	fmt.Println("Game running...")
}

func (g *GameRunState) Next() State {
	return &GameEndState{}
}

type GameEndState struct{}

func (g *GameEndState) Update() {
	fmt.Println("Game ending...")
}

func (g *GameEndState) Next() State {
	return &GameStartState{}
}
