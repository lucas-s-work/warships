package game

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/warships/game/world"
)

type Player struct {
	selectedEntities []world.Entity
	game             *Game
	window           *gl.Window
	camera           *Camera
}

func CreatePlayer(window *gl.Window, game *Game) *Player {
	return &Player{
		selectedEntities: make([]world.Entity, 0),
		window:           window,
		game:             game,
		camera:           CreateCamera(game),
	}
}

/*
Entity selection and interaction
*/

func (p *Player) selectEntities(pos mgl32.Vec2) {
	p.selectedEntities = p.game.EntitiesUnderPoint(pos.Add(p.camera.Position()))
}

/*
Input handling
*/

var keys = []glfw.Key{
	glfw.KeyW,
	glfw.KeyA,
	glfw.KeyD,
	glfw.KeyS,
	glfw.KeyUp,
	glfw.KeyDown,
	glfw.KeyLeft,
	glfw.KeyRight,
}

func (p *Player) checkInputs() {
	for _, k := range keys {
		if gl.CheckKeyPressed(k) {
			p.keyPressSelectedEntities(k)
			p.handleKeyPress(k)
		}
	}

	mousePos, _ := gl.GetMouseInfo()
	if gl.CheckMouseTapped(glfw.MouseButton1) {
		p.selectEntities(mousePos)
	}
}

func (p *Player) keyPressSelectedEntities(key glfw.Key) {
	for _, e := range p.selectedEntities {
		if e != nil {
			e.KeyPressed(key)
		}
	}
}

func (p *Player) handleKeyPress(key glfw.Key) {
	switch key {
	case glfw.KeyUp:
		p.camera.Move(world.UP)
	case glfw.KeyDown:
		p.camera.Move(world.DOWN)
	case glfw.KeyLeft:
		p.camera.Move(world.LEFT)
	case glfw.KeyRight:
		p.camera.Move(world.RIGHT)
	}
}

func (p *Player) Tick() {
	p.checkInputs()
	p.camera.Tick()
}
