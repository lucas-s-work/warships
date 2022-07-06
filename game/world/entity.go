package world

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
	"github.com/lucas-s-work/gopengl3/util"
	custom_renderers "github.com/lucas-s-work/warships/renderers"
)

const (
	BACKGROUND_LAYER = 0
	SHIP_LAYER       = 1
	MODULE_LAYER     = 2
	PROJECTILE_LAYER = 3
	GUI_LAYER        = 10
)

type Dir int8

const (
	UP    Dir = 1
	DOWN  Dir = 2
	LEFT  Dir = 3
	RIGHT Dir = 4
)

type EntityType int8

const (
	UNDEFINED = 0
	SHIP      = 1
	MODULE    = 2
)

type KeyInputEvent struct {
	MousePos  mgl32.Vec2
	CameraPos mgl32.Vec2
	Key       glfw.Key
}

type Entity interface {
	Renderable
	Init()
	Delete()
	Tick()
	OnClick()
	SetID(int)
	GetID() int
	InBounds(mgl32.Vec2) bool
	BoundCenter() mgl32.Vec2
	Angle() float64
	SetAngle(float64)
	SetAngle2(float64)
	KeyPressed(KeyInputEvent)
	KeyTapped(KeyInputEvent)
	MousePressed(key glfw.MouseButton, pos mgl32.Vec2)
	Type() EntityType
	OnCollision(int) bool
}

type BaseEntity struct {
	renderer         *renderers.Rotational
	position, Camera mgl32.Vec2
	sprite           *util.ListNode
	world            World
	Bounds           mgl32.Vec2
	texture          string
	id               int
	rotation         float64
}

func CreateBaseEntity(w World, position mgl32.Vec2, texture string, layer int, size int, bounds mgl32.Vec2) *BaseEntity {
	entity := &BaseEntity{
		world:    w,
		Bounds:   bounds,
		position: position,
		Camera:   w.CameraPosition(),
		texture:  texture,
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

func (e *BaseEntity) Init() {
	/*
		Objects have two concepts of position, a relative position and an absolute position. The relative position for base entities is the origin
		the absolute position is wherever this object exists relative to the origin.
		Whenever performing rotations they are always performed on the relative position. For base entities this means rotating the sprite about
		its center.
	*/
	startingPos := e.StartingPosition()
	e.world.Context().AddJob(func() {
		r := e.renderer
		v, t, err := graphics.Rectangle(0, 0, e.Bounds.X(), e.Bounds.Y(), 0, 0, int(e.Bounds.X()), int(e.Bounds.Y()), r.Texture())
		if err != nil {
			panic(err)
		}
		a, err := r.AllocateAndSetVertices(v, t)
		if err != nil {
			panic(err)
		}
		r.SetTranslation(startingPos)
		e.sprite = a
		r.Update()
	})
}

func (e *BaseEntity) StartingPosition() mgl32.Vec2 {
	return e.position.Sub(e.BoundCenter()).Sub(e.Camera)
}

func (e *BaseEntity) SetID(id int) {
	e.id = id
}

func (e *BaseEntity) GetID() int {
	return e.id
}

func (e *BaseEntity) Angle() float64 {
	return e.rotation
}

func (e *BaseEntity) SetAngle(rot float64) {
	e.rotation = rot
	e.SetRotation(float32(e.rotation), e.BoundCenter())
}

func (e *BaseEntity) SetAngle2(rot float64) {
	e.SetRotation2(float32(rot), e.BoundCenter())
}

func (e *BaseEntity) Delete() {
	e.world.DetachEntity(e)
	e.world.Context().AddJob(func() {
		e.world.Context().Detach(e.renderer)
	})
}

func (e *BaseEntity) World() World {
	return e.world
}

func (b *BaseEntity) InBounds(p mgl32.Vec2) bool {
	// Rotate our bounds about our angle
	bc := b.Position()
	p = p.Sub(bc)
	p = mgl32.Rotate2D(-float32(b.Angle())).Mul2x1(p)
	p = p.Add(bc)

	return b.inNonRotBounds(p)
}

func (e *BaseEntity) inNonRotBounds(v mgl32.Vec2) bool {
	// Used once we have rotated into a frame of reference where our bounds are along basis vectors
	b := e.Bounds
	av := v.Sub(e.position.Sub(e.BoundCenter()))
	inX := (av.X() >= 0) && (av.X() <= b.X())
	inY := (av.Y() >= 0) && (av.Y() <= b.Y())

	return inX && inY
}

func (e *BaseEntity) SetPosition(p mgl32.Vec2) {
	e.position = p
	e.setTranslation()
}

func (e *BaseEntity) SetCamera(p mgl32.Vec2) {
	e.Camera = p
	e.setTranslation()
}

func (e *BaseEntity) setTranslation() {
	e.renderer.SetTranslation(e.position.Sub(e.BoundCenter()).Sub(e.Camera))
}

func (e *BaseEntity) SetRotation(angle float32, center mgl32.Vec2) {
	e.renderer.SetRotation1(angle, center)
}

func (e *BaseEntity) SetRotation2(angle float32, center mgl32.Vec2) {
	e.renderer.SetRotation2(angle, center)
}

func (e *BaseEntity) Renderer() graphics.Renderer {
	return e.renderer
}

func (e *BaseEntity) Position() mgl32.Vec2 {
	return e.position
}

func (e *BaseEntity) BoundCenter() mgl32.Vec2 {
	return e.Bounds.Mul(0.5)
}

func (e *BaseEntity) Center() mgl32.Vec2 {
	return e.Bounds.Add(e.BoundCenter())
}

func (e *BaseEntity) Type() EntityType                                { return UNDEFINED }
func (e *BaseEntity) Tick()                                           {}
func (*BaseEntity) KeyPressed(KeyInputEvent)                          {}
func (*BaseEntity) KeyTapped(KeyInputEvent)                           {}
func (*BaseEntity) MousePressed(key glfw.MouseButton, pos mgl32.Vec2) {}
func (*BaseEntity) OnClick()                                          {}
func (b *BaseEntity) Sprite() mgl32.Vec4                              { return mgl32.Vec4{} }
func (b *BaseEntity) TextureFile() string {
	return b.texture
}
func (b *BaseEntity) OnCollision(d int) bool { return false }

/*
Scaled base entity is a special entity that supports size scaling
*/

type ScaledBaseEntity struct {
	*BaseEntity
	renderer *custom_renderers.Scaled
}

func CreateScaledBaseEntity(w World, position mgl32.Vec2, texture string, layer int, size int, bounds mgl32.Vec2) *ScaledBaseEntity {
	entity := &BaseEntity{
		world:    w,
		Bounds:   bounds,
		position: position,
		Camera:   w.CameraPosition(),
		texture:  texture,
	}
	scaled := &ScaledBaseEntity{
		BaseEntity: entity,
	}

	w.Context().AddJob(func() {
		r, err := custom_renderers.CreateScaledRenderer(w.Window(), texture, int32(size)*6)
		if err != nil {
			panic(err)
		}
		w.Context().Attach(r, layer)

		r.SetTranslation(position)

		scaled.renderer = r
	})

	return scaled
}

func (b *ScaledBaseEntity) Delete() {
	b.World().DetachEntity(b)
	b.World().Context().AddJob(func() {
		b.World().Context().Detach(b.renderer)
	})
}

func (b *ScaledBaseEntity) SetPosition(p mgl32.Vec2) {
	b.position = p
	b.setTranslation()
}

func (b *ScaledBaseEntity) SetCamera(p mgl32.Vec2) {
	b.Camera = p
	b.setTranslation()
}

func (b *ScaledBaseEntity) setTranslation() {
	b.renderer.SetTranslation(b.position.Sub(b.BoundCenter()).Sub(b.Camera))
}

func (b *ScaledBaseEntity) SetRotation(angle float32, center mgl32.Vec2) {
	b.renderer.SetRotation1(angle, center)
}

func (b *ScaledBaseEntity) Renderer() graphics.Renderer {
	return b.renderer
}

func (b *ScaledBaseEntity) SetScale(scale float32) {
	b.renderer.SetScale(scale)
}
