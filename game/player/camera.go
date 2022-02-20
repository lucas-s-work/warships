package player

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/warships/game/world"
)

const (
	MaxCameraSpeed     = 10
	CameraVelocityMult = 1.01
)

type Camera struct {
	position mgl32.Vec2
	holdTime int
	moved    bool
	world    world.World
}

func CreateCamera(world world.World) *Camera {
	return &Camera{
		world: world,
	}
}

func (c *Camera) Position() mgl32.Vec2 {
	return c.position
}

func (c *Camera) Speed() float32 {
	return float32(math.Min(math.Pow(CameraVelocityMult, float64(c.holdTime)), MaxCameraSpeed))
}

func (c *Camera) Move(dir world.Dir) {
	s := c.Speed()
	var v mgl32.Vec2
	switch dir {
	case world.UP:
		v = mgl32.Vec2{0, s}
	case world.DOWN:
		v = mgl32.Vec2{0, -s}
	case world.LEFT:
		v = mgl32.Vec2{-s, 0}
	case world.RIGHT:
		v = mgl32.Vec2{s, 0}
	default:
		return
	}
	c.moved = true
	c.holdTime++
	c.position = c.position.Add(v)

	c.world.SetCamera(c.position)
}

func (c *Camera) Tick() {
	if !c.moved {
		c.holdTime = 0
	}
	c.moved = false
}
