package player

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/warships/game/world"
)

type Player struct {
	selectedEntities []world.Entity
	world            world.World
	window           *gl.Window
	camera           *Camera
	gui              *PlayerGUI
}

func CreatePlayer(window *gl.Window, w world.World) *Player {
	p := &Player{
		selectedEntities: make([]world.Entity, 0),
		window:           window,
		world:            w,
		camera:           CreateCamera(w),
	}
	gui := CreatePLayerGUI(p, window, w.Context())
	p.gui = gui
	p.gui.Init()

	return p
}

/*
Camera handling
*/

func (p *Player) CameraPosition() mgl32.Vec2 {
	return p.camera.Position()
}

/*
Entity selection and interaction
*/

func (p *Player) selectEntities(pos mgl32.Vec2) {
	p.selectedEntities = p.world.EntitiesUnderPoint(pos.Add(p.camera.Position()))

	if len(p.selectedEntities) == 1 {
		p.gui.SetSelectedEntity(p.selectedEntities[0])
	} else if len(p.selectedEntities) == 0 {
		p.gui.SetSelectedEntity(nil)
	}
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
	glfw.KeySpace,
}

func (p *Player) checkInputs() {
	for _, k := range keys {
		if gl.CheckKeyPressed(k) {
			p.keyPressSelectedEntities(k)
			p.handleKeyPress(k)
		}

		if gl.CheckKeyTapped(k) {
			p.keyTapSelectedEntities(k)
		}
	}

	mousePos, _ := gl.GetMouseInfo()
	if gl.CheckMouseTapped(glfw.MouseButton1) {
		p.selectEntities(mousePos)
		p.gui.Click(mousePos)
	}
}

func (p *Player) keyPressSelectedEntities(key glfw.Key) {
	for _, e := range p.selectedEntities {
		if e != nil {
			e.KeyPressed(key)
		}
	}
}

func (p *Player) keyTapSelectedEntities(key glfw.Key) {
	for _, e := range p.selectedEntities {
		if e != nil {
			e.KeyTapped(key)
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
	p.gui.Tick()
}
