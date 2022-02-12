package ship

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/util"
	"github.com/lucas-s-work/warships/game/world"
)

type Battleship struct {
	*BaseShip
	shipSprite *util.ListNode
	speed      float32
	angle      float64
}

const (
	BattleshipTexture    = "./textures/ships/ShipBattleshipHull.png"
	BattleshipWidth      = 31
	BattleshipHeight     = 209
	BattleshipComponents = 4

	BattleshipMaxSpeed = 1.5
	BattleshipTurnRate = 0.005
)

func CreateBattleship(w world.World) *Battleship {
	e := world.CreateBaseEntity(
		w,
		BattleshipTexture,
		world.SHIP_LAYER,
		BattleshipComponents,
		mgl32.Vec4{
			0,
			0,
			BattleshipWidth,
			BattleshipHeight,
		},
	)

	s := CreateBaseShip(e, BattleshipMaxSpeed, BattleshipTurnRate)

	return &Battleship{
		BaseShip: s,
		speed:    0,
		angle:    0,
	}
}

func (b *Battleship) Init() {
	w := b.World()
	w.Context().AddJob(func() {
		b.SetPosition(mgl32.Vec2{50, 50})

		v, t, err := graphics.Rectangle(0, 0, BattleshipWidth, BattleshipHeight, 0, 0, BattleshipWidth, BattleshipHeight, b.Renderer().Texture())
		if err != nil {
			panic(err)
		}
		r := b.Renderer()
		a, err := r.AllocateAndSetVertices(v, t)
		if err != nil {
			panic(err)
		}
		b.shipSprite = a
		r.Update()
	})
}

func (b *Battleship) Tick() {
	b.UpdatePosition()
}

func (b *Battleship) KeyPressed(key string) {
	switch key {
	case "w":
		b.IncreaseSpeed(world.UP)
	case "s":
		b.IncreaseSpeed(world.DOWN)
	case "d":
		b.DecreaseTurn()
	case "a":
		b.IncreaseTurn()
	}
}

func (b *Battleship) MousePressed(key string, pos mgl32.Vec2) {

}
