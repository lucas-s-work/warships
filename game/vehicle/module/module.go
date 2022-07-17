package module

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/world"
)

// Modules don't handle their own rendering, they are handled by their parent object
type Module interface {
	world.Entity
	Position() mgl32.Vec2
	RelativePosition() mgl32.Vec2
	Sprite() mgl32.Vec4
	Texture() string
	Health() int
	Repair(int)
	SetTarget(mgl32.Vec2)
	Destroyed() bool
}

type Weapon interface {
	Module
	OnFire(world.World, world.KeyInputEvent)
	AmmoCount() int
	Reload(int)
}

// Need to figure out what these should return
type Utility interface {
	Module
	OnTick() UtilityEvent
	OnFire(world.World, world.KeyInputEvent) UtilityEvent
}

type UtilityEvent struct {
}

type BaseModule struct {
	*world.BaseEntity
	position  mgl32.Vec2
	parent    world.Entity
	angle     float64
	aimCenter mgl32.Vec2
	health    int
}

func CreateBaseModule(w world.World, pos mgl32.Vec2, parent world.Entity, tex string, bounds, aimCenter mgl32.Vec2, health int) *BaseModule {
	entity := world.CreateBaseEntity(w, pos, tex, world.MODULE_LAYER, 12, bounds)
	return &BaseModule{
		BaseEntity: entity,
		position:   pos,
		parent:     parent,
		aimCenter:  aimCenter,
		health:     health,
	}
}

func (b *BaseModule) SetAngle(angle float64) {
	// Override the default behaviour of the base entity, we rotate relative to our parent's center (origin)
	// rather than relative to our center
	b.angle = angle
	b.SetRotation2(float32(b.Angle()), b.RelativePosition().Mul(-1).Add(b.BoundCenter()))
}

func (b *BaseModule) SetTarget(aim mgl32.Vec2) {
	d := aim.Sub(b.Position())
	b.SetRotation(-float32(math.Atan2(float64(d.X()), float64(d.Y()))+b.Angle()), b.aimCenter)
}

func (b *BaseModule) Angle() float64 {
	return b.angle
}

func (b *BaseModule) Position() mgl32.Vec2 {
	// Our "relative" position is the only thing we need to rotate
	// This represents the center of our module and we just rotate around 0 (the relative center of our parent)
	// Our true position is then this added to our parent entitites absolute position
	p := mgl32.Rotate2D(float32(b.Angle())).Mul2x1(b.position)
	p = p.Add(b.parent.Position())

	return p
}

func (b *BaseModule) RelativePosition() mgl32.Vec2 {
	return b.position
}

func (b *BaseModule) Health() int {
	return b.health
}

func (b *BaseModule) OnCollision(damage int) bool {
	b.health -= damage

	if b.health <= 0 {
		// Pass the damage onto the parent and destroy the module
		b.parent.OnCollision(damage)
		b.Delete()
	}

	return true
}

func (b *BaseModule) Destroyed() bool {
	return b.health <= 0
}

func (*BaseModule) Repair(int) {
	fmt.Println("Repaired")
}

func (*BaseModule) EntityType() world.EntityType {
	return world.MODULE
}

func (b *BaseModule) Parent() world.Entity {
	return b.parent
}
