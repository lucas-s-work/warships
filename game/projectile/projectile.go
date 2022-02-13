package projectile

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/world"
)

type Projectile interface {
	world.Entity
	WorldPosition() mgl32.Vec3
	Velocity() mgl32.Vec3
}

const (
	Gravity          = 0.05
	ProjectileWidth  = 16
	ProjectileHeight = 16
)

type BaseProjectile struct {
	*world.BaseEntity
	world         world.World
	worldPosition mgl32.Vec3
	velocity      mgl32.Vec3
	onCollision   func(world.Entity)
}

func CreateBaseProjectile(w world.World, texture string, velocity, position mgl32.Vec3, onCollision func(world.Entity)) *BaseProjectile {
	return &BaseProjectile{
		BaseEntity:    world.CreateBaseEntity(w, position.Vec2(), texture, world.PROJECTILE_LAYER, 1, mgl32.Vec2{ProjectileWidth, ProjectileHeight}),
		world:         w,
		worldPosition: position,
		velocity:      velocity,
		onCollision:   onCollision,
	}
}

func (*BaseProjectile) Init() {}

func (b *BaseProjectile) checkCollision() {
	if b.worldPosition.Z() <= 0 {
		collidedEntites := b.world.EntitiesUnderPoint(b.worldPosition.Vec2())

		// First element is the projectile
		if len(collidedEntites) > 1 {
			b.onCollision(collidedEntites[1])
		}

		b.Delete()
	}
}

func (b *BaseProjectile) Velocity() mgl32.Vec3 {
	return b.velocity
}

func (b *BaseProjectile) WorldPosition() mgl32.Vec3 {
	return b.worldPosition
}

func (b *BaseProjectile) Delete() {
	b.world.DetachEntity(b)
	b.world.Context().AddJob(func() {
		b.world.Context().Detach(b.Renderer())
	})
}

func (b *BaseProjectile) UpdatePosition() {
	b.worldPosition = b.worldPosition.Add(b.velocity)
	b.velocity = b.velocity.Sub(mgl32.Vec3{0, 0, Gravity})
	b.SetPosition(b.worldPosition.Vec2())
}

func (b *BaseProjectile) Tick() {
	b.UpdatePosition()
	b.checkCollision()
}
