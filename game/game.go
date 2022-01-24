package game

import "github.com/lucas-s-work/gopengl3/graphics"

type Game struct {
}

func CreateGame(_ *graphics.Context) *Game {
	return &Game{}
}

func (g *Game) Tick() {
}
