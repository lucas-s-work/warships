package ship

import (
	"fmt"
	"math"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/vehicle/module"
	"github.com/lucas-s-work/warships/game/world"
)

type BaseShip struct {
	*world.BaseEntity
	speed float64

	MaxSpeed    float64
	MaxTurnRate float64

	health int

	MaxModules int
	modules    []module.Module
}

func CreateBaseShip(baseEntity *world.BaseEntity, maxSpeed float64, maxTurnRate float64, maxModules int) *BaseShip {
	return &BaseShip{
		BaseEntity: baseEntity,
		speed:      0,

		MaxSpeed:    maxSpeed,
		MaxTurnRate: maxTurnRate,

		health: 10,

		MaxModules: maxModules,
		modules:    make([]module.Module, maxModules),
	}
}

func (b *BaseShip) AttachModule(slot int, mod module.Module) error {
	if b.modules[slot] != nil {
		return fmt.Errorf("cannot place into module slot: %v, already occupied", slot)
	}

	b.modules[slot] = mod

	return nil
}

func (b *BaseShip) DetachModule(slot int) {
	if mod := b.modules[slot]; mod != nil {
		b.World().DetachEntity(mod)
		mod.Delete()
	}

	b.modules[slot] = nil
}

func (b *BaseShip) Modules() []module.Module {
	return b.modules
}

func (b *BaseShip) SetAngle(angle float64) {
	b.BaseEntity.SetAngle(angle)

	for _, m := range b.modules {
		if m != nil {
			m.SetAngle(angle)
		}
	}
}

func (b *BaseShip) Velocity() mgl32.Vec2 {
	return mgl32.Vec2{-float32(b.speed * math.Sin(b.Angle())), float32(b.speed * math.Cos(b.Angle()))}
}

func (b *BaseShip) UpdatePosition() {
	b.SetPosition(b.Position().Add(b.Velocity()))

	for _, mod := range b.modules {
		if mod == nil {
			continue
		}

		mod.SetPosition(b.Position().Add(b.Velocity()).Add(mod.RelativePosition()).Add(mod.Offset()))
	}
}

func (b *BaseShip) IncreaseSpeed(dir world.Dir) {
	if b.speed == 0 {
		b.speed = 0.05
	} else {
		if dir == world.UP {
			b.speed += (b.MaxSpeed - math.Abs(b.speed)) / (100 * b.MaxSpeed)
		} else if dir == world.DOWN {
			b.speed -= (b.MaxSpeed - math.Abs(b.speed)) / (100 * b.MaxSpeed)
		}
	}
}

func (b *BaseShip) DecreaseSpeed() {
	b.speed *= 0.95
	if math.Abs(float64(b.speed)) < 0.05 {
		b.speed = 0
	}
}

func (b *BaseShip) turnFactor() float64 {
	absSpeed := math.Abs(float64(b.speed))
	// Inverted quadratic reflected in the y axis. Roots at 0 and MaxSpeed
	return -4 * (b.speed / b.MaxSpeed) * (absSpeed - b.MaxSpeed) / b.MaxSpeed
}

func (b *BaseShip) IncreaseTurn() {
	b.SetAngle(b.Angle() + b.MaxTurnRate*b.turnFactor())
}

func (b *BaseShip) DecreaseTurn() {
	b.SetAngle(b.Angle() - b.MaxTurnRate*b.turnFactor())
}

func (b *BaseShip) KeyPressed(event world.KeyInputEvent) {
	switch event.Key {
	case glfw.KeyW:
		b.IncreaseSpeed(world.UP)
	case glfw.KeyS:
		b.IncreaseSpeed(world.DOWN)
	case glfw.KeyD:
		b.DecreaseTurn()
	case glfw.KeyA:
		b.IncreaseTurn()
	}
}

func (b *BaseShip) Delete() {
	for _, m := range b.modules {
		if m != nil {
			m.Delete()
		}
	}

	b.BaseEntity.Delete()
}

func (b *BaseShip) Type() world.EntityType {
	return world.SHIP
}

func (b *BaseShip) OnCollision(damage int) bool {
	b.health -= damage

	if b.health <= 0 {
		b.Delete()
	}

	return true
}

func (b *BaseShip) Tick() {
	b.UpdatePosition()
}
