package projectile

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/warships/game/world"
)

type Bullet struct {
	*BaseProjectile
}

const (
	BulletTexture       = "./textures/projectiles/BulletProjectile.png"
	BulletSpeed         = 100
	BulletInitialHeight = 10
)

func CreateBullet(w world.World, velocity mgl32.Vec3, position mgl32.Vec2) *Bullet {
	return &Bullet{
		BaseProjectile: CreateBaseProjectile(
			w,
			BulletTexture,
			velocity,
			mgl32.Vec3{position.X(), position.Y(), BulletInitialHeight},
			bulletOnCollision,
		),
	}
}

func (b *Bullet) Init() {
	w := b.World()

	startingPos := b.StartingPosition()
	w.Context().AddJob(func() {
		v, t, err := graphics.Rectangle(0, 0, 13, 21, 0, 0, 13, 21, b.Renderer().Texture())
		if err != nil {
			panic(err)
		}
		r := b.Renderer()
		_, err = r.AllocateAndSetVertices(v, t)
		if err != nil {
			panic(err)
		}
		r.SetTranslation(startingPos)
		r.Update()
	})
}

func bulletOnCollision(e world.Entity) {
	fmt.Println(e)
}
