package world

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
)

const (
	BACKGROUND_LAYER = 0
	SHIP_LAYER       = 1
)

type Dir int8

const (
	UP    Dir = 1
	DOWN  Dir = 2
	LEFT  Dir = 3
	RIGHT Dir = 4
)

type Entity interface {
	Init()
	Delete()
	Tick()
	OnClick()
	Renderer() *renderers.Rotational
	Position() mgl32.Vec2
	SetPosition(mgl32.Vec2)
	InBounds(mgl32.Vec2) bool
	KeyPressed(key glfw.Key)
	MousePressed(key glfw.MouseButton, pos mgl32.Vec2)
}

type BaseEntity struct {
	renderer *renderers.Rotational
	position mgl32.Vec2
	world    World
	bounds   mgl32.Vec2
}

func CreateBaseEntity(w World, position mgl32.Vec2, texture string, layer int, size int, bounds mgl32.Vec2) *BaseEntity {
	entity := &BaseEntity{
		world:    w,
		bounds:   bounds,
		position: position,
	}
	w.Context().AddJob(func() {
		r, err := renderers.CreateRotationalRenderer(w.Window(), texture, int32(size)*6)
		if err != nil {
			panic(err)
		}
		w.Context().Attach(r, layer)
		r.SetTranslation(position)

		entity.renderer = r
	})

	return entity
}

func (e *BaseEntity) Init() {}

func (e *BaseEntity) Tick() {
}

func (e *BaseEntity) Delete() {
	e.renderer.Delete()
}

func (e *BaseEntity) OnClick() {
}

func (e *BaseEntity) InBounds(v mgl32.Vec2) bool {
	b := e.bounds
	av := v.Sub(e.position.Sub(e.BoundCenter()))
	inX := (av.X() >= 0) && (av.X() <= b.X())
	inY := (av.Y() >= 0) && (av.Y() <= b.Y())

	return inX && inY
}

func (e *BaseEntity) SetPosition(p mgl32.Vec2) {
	e.position = p
	e.renderer.SetTranslation(p.Sub(e.BoundCenter()))
}

func (e *BaseEntity) SetRotation(angle float32, center mgl32.Vec2) {
	e.renderer.SetRotation1(angle, center)
}

func (e *BaseEntity) Renderer() *renderers.Rotational {
	return e.renderer
}

func (e *BaseEntity) Position() mgl32.Vec2 {
	return e.position
}

func (e *BaseEntity) World() World {
	return e.world
}

func (e *BaseEntity) BoundCenter() mgl32.Vec2 {
	x := 0.5 * (e.bounds.X())
	y := 0.5 * (e.bounds.Y())

	return mgl32.Vec2{x, y}
}

func (e *BaseEntity) Center() mgl32.Vec2 {
	return e.position.Add(e.BoundCenter())
}

func (*BaseEntity) KeyPressed(key glfw.Key) {

}

func (*BaseEntity) MousePressed(key glfw.MouseButton, pos mgl32.Vec2) {

}
