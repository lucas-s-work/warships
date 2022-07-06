package vehicle

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/warships/game/vehicle/module"
	"github.com/lucas-s-work/warships/game/world"
)

type Vehicle interface {
	Position() mgl32.Vec2
	Velocity() mgl32.Vec2
	IncreaseSpeed(world.Dir)
	DecreaseSpeed()
	IncreaseTurn()
	DecreaseTurn()
	Sprite() mgl32.Vec4
	Renderer() graphics.Renderer
	AttachModule(int, module.Module) error
	DetachModule(int)
	Modules() []module.Module
}

type Ship interface {
	Vehicle
}
