package world

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
)

type World interface {
	AttachEntity(Entity)
	DetachEntity(Entity)
	CameraPosition() mgl32.Vec2
	Window() *gl.Window
	Context() *graphics.Context
	EntitiesUnderPoint(mgl32.Vec2) []Entity
	SetCamera(mgl32.Vec2)
}
