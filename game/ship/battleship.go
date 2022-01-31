package ship

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/util"
	"github.com/lucas-s-work/warships/game/world"
)

type Battleship struct {
	*world.BaseEntity
	shipSprite *util.ListNode
	speed      float32
	angle      float64
}

const (
	BattleshipTexture    = "./textures/ships/ShipBattleshipHull.png"
	BattleshipWidth      = 31
	BattleshipHeight     = 209
	BattleshipComponents = 4
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

	return &Battleship{
		BaseEntity: e,
		speed:      0.5,
		angle:      0,
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
	velocity := mgl32.Vec2{b.speed * float32(math.Cos(b.angle)), b.speed * float32(math.Sin(b.angle))}
	b.SetPosition(b.Position().Add(velocity))
	fmt.Println(b.Position())
	b.SetRotation(float32(b.angle-math.Pi/2), b.BoundCenter())
	b.angle += 0.005
	b.speed += 0.01
}
