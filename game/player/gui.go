package player

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/warships/game/vehicle"
	"github.com/lucas-s-work/warships/game/world"
	"github.com/lucas-s-work/warships/gui"
)

const (
	maxPlayerGUIElements = 10
	playerGUISprite      = "./textures/gui/base.png"
)

const (
	PLAYER_GUI_BASE_LAYER = 0
)

type PlayerGUI struct {
	*gui.BaseGUI
	ctx    *graphics.Context
	window *gl.Window

	selectedShip *shipElement
}

func CreatePLayerGUI(player *Player, window *gl.Window, ctx *graphics.Context) *PlayerGUI {
	return &PlayerGUI{
		BaseGUI: gui.CreateBaseGUI(maxPlayerGUIElements),
		ctx:     player.world.Context(),
		window:  window,
	}
}

func (g *PlayerGUI) Init() {
	element := CreateSelectedShipElement(g.ctx, g, mgl32.Vec2{32, 32})
	g.AttachElement(element)
	g.selectedShip = element
}

func (g *PlayerGUI) Tick() {
	g.BaseGUI.Tick()
}

func (g *PlayerGUI) SetSelectedShip(s vehicle.Ship) {
	g.selectedShip.SetSelectedShip(s)
}

func (g *PlayerGUI) UpdateMousePosition(pos, camera mgl32.Vec2) {
	if mod := g.selectedShip.SelectedModule(); mod != nil {
		mod.SetTarget(pos.Add(camera))
	}
}

func (g *PlayerGUI) OnFire(w world.World, event world.KeyInputEvent) {
	g.selectedShip.OnFire(w, event)
}
