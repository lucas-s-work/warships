package ship

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/vehicle/module"
	"github.com/lucas-s-work/warships/game/world"
)

type Destroyer struct {
	*BaseShip
}

const (
	DestroyerTexture    = "./textures/ships/ShipDestroyerHull.png"
	DestroyerWidth      = 20
	DestroyerHeight     = 100
	DestroyerComponents = 2

	DestroyerMaxSpeed = 2.4
	DestroyerTurnRate = 0.009

	DestroyerForwardModuleY = 30
	DestroyerRearModuleY    = -30
	DestroyerModuleSlots    = 2
)

var (
	DestroyerSprite = mgl32.Vec4{
		0,
		0,
		DestroyerWidth,
		DestroyerHeight,
	}
)

func CreateDestroyer(w world.World, position mgl32.Vec2) *Destroyer {
	e := world.CreateBaseEntity(
		w,
		position,
		DestroyerTexture,
		world.SHIP_LAYER,
		DestroyerComponents,
		mgl32.Vec2{
			DestroyerWidth,
			DestroyerHeight,
		},
	)

	s := CreateBaseShip(e, DestroyerMaxSpeed, DestroyerTurnRate, DestroyerModuleSlots)
	d := &Destroyer{
		BaseShip: s,
	}

	frontTurret := module.CreateSmallTurret(w, mgl32.Vec2{0, DestroyerForwardModuleY}, d)
	d.AttachModule(0, frontTurret)
	w.AttachEntity(frontTurret)

	rearTurret := module.CreateSmallTurret(w, mgl32.Vec2{0, DestroyerRearModuleY}, d)
	d.AttachModule(1, rearTurret)
	w.AttachEntity(rearTurret)

	return d
}

func (b *Destroyer) Sprite() mgl32.Vec4 {
	return DestroyerSprite
}
