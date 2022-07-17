package projectile

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/warships/game/world"
	"github.com/lucas-s-work/warships/renderers"
)

type MissileStage int8

const (
	MissileStageStart    = iota
	MissileStageCruise   = iota
	MissileStageTerminal = iota
)

type Missile struct {
	*BaseProjectile
	target       mgl32.Vec3
	MissileStage MissileStage
}

const (
	MissileTexture     = "./textures/projectiles/missiles.png"
	MissileSpeed       = 3.0
	MissileLaunchSpeed = 3.0
	MissileMaxHeight   = 125.0
	MissileDivestart   = 75.0
	MissileRange       = 250.0
)

func CreateMissile(w world.World, target mgl32.Vec3, position mgl32.Vec2) *Missile {
	return &Missile{
		BaseProjectile: CreateBaseProjectile(
			w,
			MissileTexture,
			mgl32.Vec3{0, 0, 1},
			mgl32.Vec3{position.X(), position.Y(), 0.5},
			missileOnCollision,
		),
		target:       target,
		MissileStage: MissileStageStart,
	}
}

func (b *Missile) Init() {
	w := b.World()

	startingPos := b.StartingPosition()
	w.Context().AddJob(func() {
		v, t, err := graphics.Rectangle(0, 0, 5, 19, 0, 0, 5, 19, b.Renderer().Texture())
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

func missileOnCollision(world.Entity) bool {
	return false
}

func (b *Missile) Velocity() mgl32.Vec3 {
	return b.velocity
}

func (b *Missile) updateVelocity() {
	if b.worldPosition.Z() <= 0 {
		return
	}

	toTarget := b.target.Vec2().Sub(b.Position()).Len()
	if toTarget <= MissileDivestart {
		b.terminalVelocity()
		return
	} else if b.worldPosition.Z() < MissileMaxHeight {
		b.startVelocity()
		return
	}

	b.cruiseVelocity()
}

func (b *Missile) startVelocity() {
	if b.MissileStage > MissileStageStart {
		return
	}

	dt := b.target.Vec2().Sub(b.worldPosition.Vec2()).Normalize().Mul(MissileSpeed)
	b.velocity = mgl32.Vec3{
		dt.X(),
		dt.Y(),
		MissileLaunchSpeed,
	}
}

func (b *Missile) cruiseVelocity() {
	if b.MissileStage > MissileStageCruise {
		return
	}
	b.MissileStage = MissileStageCruise

	b.velocity[2] = 0
}

func (b *Missile) terminalVelocity() {
	b.velocity = b.target.Sub(b.worldPosition).Normalize().Mul(MissileSpeed * 1.2)
	b.MissileStage = MissileStageTerminal
}

func (b *Missile) UpdatePosition() {
	b.updateVelocity()
	b.worldPosition = b.worldPosition.Add(b.velocity)
	b.SetPosition(b.worldPosition.Vec2())
	b.SetScale(float32(math.Max(float64(1.0-b.worldPosition.Z()/250.0), 0.5)))
}

func (b *Missile) Tick() {
	b.UpdatePosition()
	b.checkCollision()
}
