package ship

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Ship interface {
	Init()
	Tick()
	Position() mgl32.Vec2
	Velocity() mgl32.Vec2
	Status()
}
