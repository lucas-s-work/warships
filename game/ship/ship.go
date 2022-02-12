package ship

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/world"
)

type Ship interface {
	Init()
	Tick()
	Position() mgl32.Vec2
	Velocity() mgl32.Vec2
	Status()
	UpdatePosition()
	IncreaseSpeed(world.Dir)
	Stop()
	DecreaseSpeed()
	TurnFactor() float64
	IncreaseTurn()
	DecreaseTurn()
}

type BaseShip struct {
	*world.BaseEntity
	speed float64
	angle float64

	MaxSpeed    float64
	MaxTurnRate float64
}

func CreateBaseShip(baseEntity *world.BaseEntity, maxSpeed float64, maxTurnRate float64) *BaseShip {
	return &BaseShip{
		BaseEntity: baseEntity,
		speed:      0,
		angle:      0,

		MaxSpeed:    maxSpeed,
		MaxTurnRate: maxTurnRate,
	}
}

func (b *BaseShip) Velocity() mgl32.Vec2 {
	return mgl32.Vec2{float32(b.speed * math.Cos(b.angle)), float32(b.speed * math.Sin(b.angle))}
}

func (b *BaseShip) UpdatePosition() {
	b.SetPosition(b.Position().Add(b.Velocity()))
	b.SetRotation(float32(b.angle-math.Pi/2), b.BoundCenter())
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

func (b *BaseShip) Stop() {
	b.speed = 0
}

func (b *BaseShip) DecreaseSpeed() {
	b.speed *= 0.95
	if math.Abs(float64(b.speed)) < 0.05 {
		b.speed = 0
	}
}

func (b *BaseShip) TurnFactor() float64 {
	absSpeed := math.Abs(float64(b.speed))
	// Inverted quadratic
	return -4 * (b.speed / BattleshipMaxSpeed) * (absSpeed - BattleshipMaxSpeed) / BattleshipMaxSpeed
}

func (b *BaseShip) IncreaseTurn() {
	b.angle += b.MaxTurnRate * b.TurnFactor()
}

func (b *BaseShip) DecreaseTurn() {
	b.angle -= b.MaxTurnRate * b.TurnFactor()
}
