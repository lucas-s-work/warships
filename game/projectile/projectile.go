package projectile

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/world"
)

type Projectile interface {
	world.Entity
}

const (
	Gravity          = 0.05
	ProjectileWidth  = 16
	ProjectileHeight = 16
	CollisionHeight  = 24 // Height at which objects start getting hit
)

type BaseProjectile struct {
	*world.ScaledBaseEntity
	world         world.World
	worldPosition mgl32.Vec3
	velocity      mgl32.Vec3
	onCollision   func(world.Entity) bool
}

func CreateBaseProjectile(w world.World, texture string, velocity, position mgl32.Vec3, onCollision func(world.Entity) bool) *BaseProjectile {
	p := &BaseProjectile{
		ScaledBaseEntity: world.CreateScaledBaseEntity(w, position.Vec2(), texture, world.PROJECTILE_LAYER, 1, mgl32.Vec2{ProjectileWidth, ProjectileHeight}),
		world:            w,
		worldPosition:    position,
		velocity:         velocity,
		onCollision:      onCollision,
	}

	return p
}

func (b *BaseProjectile) checkCollision() {
	if height := b.worldPosition.Z(); height <= CollisionHeight {
		collidedEntites := b.world.EntitiesUnderPoint(b.worldPosition.Vec2())

		// First element is the projectile
		if len(collidedEntites) > 1 {
			for i, entity := range collidedEntites {
				// The projectile shouldn't collide with itself or with it's parent
				if i == 0 {
					continue
				}

				if b.onCollision(entity) {
					b.Delete()

					return
				}
			}
		}

		if height <= 0 {
			b.Delete()
		}
	}
}

func (b *BaseProjectile) Velocity() mgl32.Vec3 {
	return b.velocity
}

func (b *BaseProjectile) UpdatePosition() {
	b.worldPosition = b.worldPosition.Add(b.velocity)
	b.velocity = b.velocity.Sub(mgl32.Vec3{0, 0, Gravity})
	b.SetPosition(b.worldPosition.Vec2())
	b.SetScale(float32(math.Max(float64(1.0-b.worldPosition.Z()/250.0), 0.3)))
}

func (b *BaseProjectile) Tick() {
	b.UpdatePosition()
	b.checkCollision()
}
