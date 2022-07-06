package projectile

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/warships/game/world"
	"github.com/lucas-s-work/warships/renderers"
)

type Bullet struct {
	*BaseProjectile
}

const (
	BulletTexture       = "./textures/projectiles/Bullet.png"
	BulletSpeed         = 5.0
	BulletInitialHeight = 25 // Spawn above the projectile hit detection height to avoid colliding with parent
	BulletInitialZVel   = 3
)

func CreateBullet(w world.World, dir mgl32.Vec2, position mgl32.Vec2) *Bullet {
	velocity := dir.Normalize().Mul(BulletSpeed)

	return &Bullet{
		BaseProjectile: CreateBaseProjectile(
			w,
			BulletTexture,
			mgl32.Vec3{velocity.X(), velocity.Y(), BulletInitialZVel},
			mgl32.Vec3{position.X(), position.Y(), BulletInitialHeight},
			bulletOnCollision,
		),
	}
}

func (b *Bullet) Init() {
	w := b.World()

	startingPos := b.StartingPosition()
	w.Context().AddJob(func() {
		v, t, err := graphics.Rectangle(0, 0, 13, 5, 0, 0, 13, 5, b.Renderer().Texture())
		if err != nil {
			panic(err)
		}
		r := b.Renderer().(*renderers.Scaled)
		_, err = r.AllocateAndSetVertices(v, t)
		if err != nil {
			panic(err)
		}
		r.SetTranslation(startingPos)
		r.Update()
	})
}

func bulletOnCollision(e world.Entity) bool {
	e.OnCollision(15)

	return true
}
