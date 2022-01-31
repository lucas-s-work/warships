package world

import (
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
)

type World interface {
	AttachEntity(Entity)
	Window() *gl.Window
	Context() *graphics.Context
}
