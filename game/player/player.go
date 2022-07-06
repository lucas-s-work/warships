package player

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/warships/game/vehicle"
	"github.com/lucas-s-work/warships/game/world"
)

type Player struct {
	selectedEntity world.Entity
	world          world.World
	window         *gl.Window
	camera         *Camera
	gui            *PlayerGUI
}

func CreatePlayer(window *gl.Window, w world.World) *Player {
	p := &Player{
		window: window,
		world:  w,
		camera: CreateCamera(w),
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
	if p.selectedEntity != nil {
		p.camera.SetBasePosition(p.selectedEntity.Position())
	}

	return p.camera.Position()
}

/*
Entity selection and interaction
*/

func (p *Player) selectEntities(pos mgl32.Vec2) {
	potentialEntities := p.world.EntitiesUnderPoint(pos.Add(p.camera.Position()))

	if len(potentialEntities) >= 1 {
		for _, e := range potentialEntities {
			p.selectedEntity = e
			if e.Type() == world.SHIP {
				s := e.(vehicle.Ship)
				p.gui.SetSelectedShip(s)

				return
			}
		}
	} else if len(potentialEntities) == 0 {
		p.gui.SetSelectedShip(nil)
		p.selectedEntity = nil
	}
}

func (p *Player) setSelectedEntity(e world.Entity) {
	if e == nil {
		p.gui.SetSelectedShip(nil)
		p.selectedEntity = nil
		// p.camera.ClearBasePosition()

		return
	}

	// Ideally this type determiniation should be performed in the GUI
	p.selectedEntity = e
	if e.Type() == world.SHIP {
		p.gui.SetSelectedShip(e.(vehicle.Ship))
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

func (p *Player) checkInputs(mousePos mgl32.Vec2) {
	cameraPos := p.CameraPosition()
	for _, k := range keys {
		if gl.CheckKeyPressed(k) {
			event := world.KeyInputEvent{
				MousePos:  mousePos,
				CameraPos: cameraPos,
				Key:       k,
			}
			p.keyPressSelectedEntities(event)
			p.handleKeyPress(event)
		}

		if gl.CheckKeyTapped(k) {
			event := world.KeyInputEvent{
				MousePos:  mousePos,
				CameraPos: cameraPos,
				Key:       k,
			}
			p.keyTapSelectedEntities(event)
		}
	}

	if gl.CheckMouseTapped(glfw.MouseButton1) {
		if !p.gui.Click(mousePos) {
			p.selectEntities(mousePos)
		}
	}
}

func (p *Player) keyPressSelectedEntities(event world.KeyInputEvent) {
	if p.selectedEntity != nil {
		p.selectedEntity.KeyPressed(event)
	}
}

func (p *Player) keyTapSelectedEntities(event world.KeyInputEvent) {
	if p.selectedEntity != nil {
		p.selectedEntity.KeyTapped(event)
	}

	// if mod := p.gui.SelectedModule(); mod != nil {
	// 	if event.Key == glfw.KeySpace {
	// 		if wep, ok := mod.(module.Weapon); ok {
	// 			wep.OnFire(p.world, event)
	// 		}
	// 	}
	// }
}

func (p *Player) handleKeyPress(event world.KeyInputEvent) {
	switch event.Key {
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
	mousePos, _ := gl.GetMouseInfo()
	p.checkInputs(mousePos)
	p.camera.Tick()
	if p.selectedEntity != nil {
		p.camera.Update()
	}

	p.gui.Tick()
	p.gui.UpdateMousePosition(mousePos, p.CameraPosition())
}
