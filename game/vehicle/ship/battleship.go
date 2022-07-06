package ship

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/vehicle/module"
	"github.com/lucas-s-work/warships/game/world"
)

type Battleship struct {
	*BaseShip
}

const (
	BattleshipTexture    = "./textures/ships/ShipBattleshipHull.png"
	BattleshipWidth      = 31
	BattleshipHeight     = 209
	BattleshipComponents = 4

	BattleshipMaxSpeed = 1.5
	BattleshipTurnRate = 0.005

	BattleShipModuleSlots = 4

	BattleshipForwardModule1 = 56
	BattleshipForwardModule2 = 40
	BattleshipRearModule2    = -46
	BattleshipRearModule1    = -62
)

var (
	BattleshipSprite = mgl32.Vec4{
		0,
		0,
		BattleshipWidth,
		BattleshipHeight,
	}
)

func CreateBattleship(w world.World, position mgl32.Vec2) *Battleship {
	e := world.CreateBaseEntity(
		w,
		position,
		BattleshipTexture,
		world.SHIP_LAYER,
		BattleshipComponents,
		mgl32.Vec2{
			BattleshipWidth,
			BattleshipHeight,
		},
	)

	s := CreateBaseShip(e, BattleshipMaxSpeed, BattleshipTurnRate, BattleShipModuleSlots)
	b := &Battleship{
		BaseShip: s,
	}

	frontTurret1 := module.CreateSmallTurret(w, mgl32.Vec2{0, BattleshipForwardModule1}, b)
	b.AttachModule(0, frontTurret1)
	w.AttachEntity(frontTurret1)

	frontTurret2 := module.CreateSmallTurret(w, mgl32.Vec2{0, BattleshipForwardModule2}, b)
	b.AttachModule(1, frontTurret2)
	w.AttachEntity(frontTurret2)

	rearTurret1 := module.CreateSmallTurret(w, mgl32.Vec2{0, BattleshipRearModule1}, b)
	b.AttachModule(3, rearTurret1)
	w.AttachEntity(rearTurret1)

	rearTurret2 := module.CreateSmallTurret(w, mgl32.Vec2{0, BattleshipRearModule2}, b)
	b.AttachModule(2, rearTurret2)
	w.AttachEntity(rearTurret2)

	return b
}

func (b *Battleship) Sprite() mgl32.Vec4 {
	return BattleshipSprite
}
