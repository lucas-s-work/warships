package game

import (
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/warships/game/world"
)

const MaxPlayerEntites = 64

type Player struct {
	selectedEntities []world.Entity
	window           *gl.Window
}

func CreatePlayer(window *gl.Window) *Player {
	return &Player{
		selectedEntities: make([]world.Entity, MaxPlayerEntites),
		window:           window,
	}
}

func (p *Player) checkInputs() {
	keys := gl.CheckKeys([]string{"w", "a", "s", "d"})
	if keys[0] {
		p.keyPressSelectedEntities("w")
	}
	if keys[1] {
		p.keyPressSelectedEntities("a")
	}
	if keys[2] {
		p.keyPressSelectedEntities("s")
	}
	if keys[3] {
		p.keyPressSelectedEntities("d")
	}
}

func (p *Player) keyPressSelectedEntities(key string) {
	for _, e := range p.selectedEntities {
		if e != nil {
			e.KeyPressed(key)
		}
	}
}

func (p *Player) Tick() {
	p.checkInputs()
}
