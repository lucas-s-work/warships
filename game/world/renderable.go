package world

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
)

type Renderable interface {
	Renderer() graphics.Renderer
	Position() mgl32.Vec2
	SetPosition(mgl32.Vec2)
	SetRotation(angle float32, center mgl32.Vec2)
	SetCamera(mgl32.Vec2)
	Sprite() mgl32.Vec4
}
