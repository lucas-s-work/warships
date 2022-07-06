package gui

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
	"github.com/lucas-s-work/gopengl3/util"
	"github.com/lucas-s-work/warships/game/world"
)

type Element interface {
	Renderer() *renderers.Translational
	Click()
	InBounds(mgl32.Vec2) bool
	SetID(int)
	GetID() int
	SetActive(bool)
	SetClick(func())
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
	position   mgl32.Vec2
	dimensions mgl32.Vec4
	active     bool
	id         int
	sprite     *util.ListNode
	clickFunc  func()
}

func CreateBaseElement(ctx *graphics.Context, position mgl32.Vec2, dimensions mgl32.Vec4, size, layer int) *BaseElement {
	element := &BaseElement{
		ctx:        ctx,
		position:   position,
		dimensions: dimensions,
		active:     true,
	}

	ctx.AddJob(func() {
		r, err := renderers.CreateTranslationalRenderer(ctx.Window(), BASE_ELEMENT_TEXTURE, int32(size)*6)
		if err != nil {
			panic(err)
		}
		ctx.Attach(r, world.GUI_LAYER+layer)
		r.SetTranslation(position)
		element.renderer = r
	})

	return element
}

func (e *BaseElement) Init() {
	e.Context().AddJob(func() {
		v, t, err := graphics.Rectangle(
			e.dimensions[0], e.dimensions[1], e.dimensions[2], e.dimensions[3], 0, 0, 1, 1, e.Renderer().Texture())
		if err != nil {
			panic(err)
		}
		r := e.Renderer()
		a, err := r.AllocateAndSetVertices(v, t)
		if err != nil {
			panic(err)
		}
		e.sprite = a
		r.Update()
	})
}
func (b *BaseElement) Delete() {
	b.ctx.AddJob(func() {
		b.ctx.Detach(b.renderer)
		b.renderer = nil
	})
}

func (b *BaseElement) SetClick(f func()) {
	b.clickFunc = f
}

func (b *BaseElement) Click() {
	if b.clickFunc != nil {
		b.clickFunc()
	}
}

func (b *BaseElement) Renderer() *renderers.Translational {
	return b.renderer
}

func (b *BaseElement) SetActive(a bool) {
	b.active = a
	b.renderer.SetActive(a)
}

func (b *BaseElement) InBounds(pos mgl32.Vec2) bool {
	pos = pos.Sub(b.position)
	return (pos.X() >= b.dimensions[0] && pos.X() <= b.dimensions.X()+b.dimensions.Z()) && (pos.Y() >= b.dimensions.Y() && pos.Y() <= b.dimensions.Y()+b.dimensions.W())
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
