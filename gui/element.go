package gui

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
	"github.com/lucas-s-work/warships/game/world"
)

type Element interface {
	Renderer() *renderers.Translational
	Click()
	InBounds(mgl32.Vec2) bool
	SetID(int)
	GetID() int
	SetActive(bool)
	Active() bool
	Delete()
	Init()
	Context() *graphics.Context
}

const (
	BASE_ELEMENT_TEXTURE = "textures/gui/base.png"
)

type BaseElement struct {
	ctx        *graphics.Context
	renderer   *renderers.Translational
	dimensions mgl32.Vec4
	active     bool
	id         int
}

func CreateBaseElement(ctx *graphics.Context, dimensions mgl32.Vec4, size, layer int) *BaseElement {
	element := &BaseElement{
		ctx:        ctx,
		dimensions: dimensions,
		active:     true,
	}

	ctx.AddJob(func() {
		r, err := renderers.CreateTranslationalRenderer(ctx.Window(), BASE_ELEMENT_TEXTURE, int32(size)*6)
		if err != nil {
			panic(err)
		}
		ctx.Attach(r, world.GUI_LAYER+layer)
		r.SetTranslation(dimensions.Vec2())
		element.renderer = r
	})

	return element
}

func (*BaseElement) Init() {}
func (b *BaseElement) Delete() {
	b.ctx.AddJob(func() {
		b.ctx.Detach(b.renderer)
		b.renderer = nil
	})
}

func (b *BaseElement) Renderer() *renderers.Translational {
	return b.renderer
}

func (b *BaseElement) SetActive(a bool) {
	b.active = a
	b.renderer.SetActive(a)
}

func (b *BaseElement) InBounds(pos mgl32.Vec2) bool {
	return (pos.X() >= b.dimensions.X() && pos.X() <= b.dimensions.X()+b.dimensions.Z()) && (pos.Y() >= b.dimensions.Y() && pos.Y() <= b.dimensions.Y()+b.dimensions.W())
}

func (b *BaseElement) Active() bool {
	return b.active
}

func (b *BaseElement) SetID(id int) {
	b.id = id
}

func (b *BaseElement) GetID() int {
	return b.id
}

func (b *BaseElement) Context() *graphics.Context {
	return b.ctx
}
