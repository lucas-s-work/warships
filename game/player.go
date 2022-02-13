package game

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/warships/game/world"
)

const MaxPlayerEntites = 64

type Player struct {
	selectedEntities []world.Entity
	game             *Game
	window           *gl.Window
}

func CreatePlayer(window *gl.Window, game *Game) *Player {
	return &Player{
		selectedEntities: make([]world.Entity, MaxPlayerEntites),
		window:           window,
		game:             game,
	}
}

func (p *Player) checkInputs() {
	keys := []glfw.Key{glfw.KeyW, glfw.KeyA, glfw.KeyD, glfw.KeyS}
	for _, k := range keys {
		if gl.CheckKeyPressed(k) {
			p.keyPressSelectedEntities(k)
		}
	}

	mousePos, _ := gl.GetMouseInfo()
	if gl.CheckMouseTapped(glfw.MouseButton1) {
		fmt.Println(p.game.EntitiesUnderPoint(mousePos))
	}
}

func (p *Player) keyPressSelectedEntities(key glfw.Key) {
	for _, e := range p.selectedEntities {
		if e != nil {
			e.KeyPressed(key)
		}
	}
}

func (p *Player) Tick() {
	p.checkInputs()
}
