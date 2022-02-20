package player

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/util"
	"github.com/lucas-s-work/warships/game/world"
	"github.com/lucas-s-work/warships/gui"
)

var (
	selectedShipElementDimensions = mgl32.Vec4{
		0, 0,
		32, 128,
	}
	selectedShipElementSize = 2
)

type selectedShipElement struct {
	*gui.BaseElement
	backgroundSprite *util.ListNode
	entityOverlay    *gui.Overlay
}

func CreateSelectedShipElement(ctx *graphics.Context, p mgl32.Vec2) *selectedShipElement {
	return &selectedShipElement{
		BaseElement:   gui.CreateBaseElement(ctx, selectedShipElementDimensions.Add(mgl32.Vec4{p.X(), p.Y(), 0, 0}), selectedShipElementSize, PLAYER_GUI_BASE_LAYER),
		entityOverlay: gui.CreateOverlay(ctx, p, PLAYER_GUI_BASE_LAYER),
	}
}

func (e *selectedShipElement) Init() {
	e.Context().AddJob(func() {
		v, t, err := graphics.Rectangle(
			selectedShipElementDimensions[0],
			selectedShipElementDimensions[1],
			selectedShipElementDimensions[2],
			selectedShipElementDimensions[3], 0, 0, 1, 1, e.Renderer().Texture())
		if err != nil {
			panic(err)
		}
		r := e.Renderer()
		a, err := r.AllocateAndSetVertices(v, t)
		if err != nil {
			panic(err)
		}
		e.backgroundSprite = a
		r.Update()
	})
}

func (e *selectedShipElement) SetSelectedEntity(entity world.Entity) {
	if entity == nil {
		e.entityOverlay.ClearOverlay()
		return
	}

	e.entityOverlay.SetOverlay(
		selectedShipElementDimensions,
		entity.Sprite(),
		entity.Renderer().Texture(),
	)
}

func (*selectedShipElement) Click() {
}
