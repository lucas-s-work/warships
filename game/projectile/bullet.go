package projectile

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/warships/game/world"
	"github.com/lucas-s-work/warships/renderers"
	"github.com/lucas-s-work/warships/util"
)

type Bullet struct {
	*BaseProjectile
}

const (
	BulletTexture       = "./textures/projectiles/Bullet.png"
	BulletSpeed         = 10.0
	BulletInitialHeight = 0 // Spawn above the projectile hit detection height to avoid colliding with parent
)

func CreateBullet(w world.World, target mgl32.Vec2, position mgl32.Vec2) *Bullet {
	// t refers to target vector here not time
	dt := target.Sub(position)
	vt, vz, ok := util.SolveQuadraticVelocity(BulletSpeed, float64(dt.Len()), Gravity)
	if !ok {
		return nil
	}

	theta := math.Atan2(float64(dt.Y()), float64(dt.X()))
	vx := vt * math.Cos(theta)
	vy := vt * math.Sin(theta)

	return &Bullet{
		BaseProjectile: CreateBaseProjectile(
			w,
			BulletTexture,
			mgl32.Vec3{float32(vx), float32(vy), float32(vz)},
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
	// e.OnCollision(15)

	return false
}
