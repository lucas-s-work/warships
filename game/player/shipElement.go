package player

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/warships/game/vehicle"
	"github.com/lucas-s-work/warships/game/vehicle/module"
	"github.com/lucas-s-work/warships/gui"
)

var (
	selectedShipElementDimensions = mgl32.Vec4{
		0, 0,
		32, 128,
	}
	selectedShipElementSize = 2
)

type shipElement struct {
	*gui.BaseElement
	parent        gui.GUI
	entityOverlay *gui.Overlay
	ship          vehicle.Ship

	modules *moduleElementGroup
}

func CreateSelectedShipElement(ctx *graphics.Context, parent gui.GUI, p mgl32.Vec2) *shipElement {
	return &shipElement{
		BaseElement:   gui.CreateBaseElement(ctx, p, selectedShipElementDimensions, selectedShipElementSize, PLAYER_GUI_BASE_LAYER),
		entityOverlay: gui.CreateOverlay(ctx, p, selectedShipElementDimensions, PLAYER_GUI_BASE_LAYER),
		modules:       CreateModuleElementGroup(ctx, parent, p.Add(mgl32.Vec2{64, 0})),
		parent:        parent,
	}
}

func (e *shipElement) SetSelectedShip(entity vehicle.Ship) {
	if entity == nil {
		e.ClearSelectedShip()

		return
	} else if entity != e.ship {
		e.ClearSelectedShip()
	}

	e.entityOverlay.SetOverlay(
		selectedShipElementDimensions,
		entity.Sprite(),
		entity.Renderer().Texture(),
	)

	e.modules.SetModules(entity.Modules())
}

func (e *shipElement) ClearSelectedShip() {
	e.entityOverlay.ClearOverlay()
	e.modules.ClearModules()
}

func (e *shipElement) SelectedShip() vehicle.Ship {
	return e.ship
}

func (e *shipElement) Modules() []module.Module {
	return e.modules.Modules()
}

func (e *shipElement) SelectedModule() module.Module {
	return e.modules.SelectedModule()
}

func (e *shipElement) Delete() {
	e.entityOverlay.Delete()
	e.BaseElement.Delete()
}
