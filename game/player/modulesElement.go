package player

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/warships/game/vehicle/module"
	"github.com/lucas-s-work/warships/gui"
)

var (
	moduleDimensions = mgl32.Vec4{
		0, 0,
		32, 32,
	}
	moduleSize = 5
)

type moduleElement struct {
	*gui.BaseElement
	mod     module.Module
	overlay *gui.Overlay
}

func CreateModuleElement(ctx *graphics.Context, p mgl32.Vec2) *moduleElement {
	return &moduleElement{
		BaseElement: gui.CreateBaseElement(ctx, p, moduleDimensions, 1, PLAYER_GUI_BASE_LAYER),
		overlay:     gui.CreateOverlay(ctx, p, moduleDimensions, PLAYER_GUI_BASE_LAYER),
	}
}

func (e *moduleElement) setModule(m module.Module) {
	if e.mod != nil {
		e.clearModule()
	}

	e.mod = m
	e.overlay.SetOverlay(
		moduleDimensions,
		m.Sprite(),
		m.Renderer().Texture(),
	)
}

func (e *moduleElement) clearModule() {
	if e.mod == nil {
		return
	}

	e.overlay.ClearOverlay()
	e.mod = nil
}

func (e *moduleElement) module() module.Module {
	return e.mod
}

func (e *moduleElement) Delete() {
	e.overlay.Delete()
	e.BaseElement.Delete()
}

type moduleElementGroup struct {
	ctx      *graphics.Context
	group    *gui.ElementGroup
	modules  []*moduleElement
	selected module.Module
}

func CreateModuleElementGroup(ctx *graphics.Context, parent gui.GUI, p mgl32.Vec2) *moduleElementGroup {
	return &moduleElementGroup{
		ctx:   ctx,
		group: gui.CreateElementGroup(parent, p),
	}
}

func (m *moduleElementGroup) SetModules(mods []module.Module) {
	if mods == nil {
		m.ClearModules()
	}

	modClick := func(mod module.Module) func() {
		return func() {
			m.selected = mod
		}
	}

	m.modules = make([]*moduleElement, len(mods))
	for i, mod := range mods {
		el := CreateModuleElement(m.ctx, mgl32.Vec2{float32(i * 32), 0})
		el.setModule(mod)
		el.SetClick(modClick(mod))
		m.modules[i] = el
		m.group.AttachElement(el)
	}
}

func (m *moduleElementGroup) ClearModules() {
	m.group.DeleteElements()
	m.modules = nil
	m.selected = nil
}

func (m *moduleElementGroup) Modules() []module.Module {
	if m.modules == nil {
		return nil
	}

	mods := make([]module.Module, len(m.modules))
	for i, m := range m.modules {
		mods[i] = m.module()
	}

	return mods
}

func (m *moduleElementGroup) SelectedModule() module.Module {
	return m.selected
}

func (m *moduleElementGroup) ClearSelectedModule() {
	m.selected = nil
}

func (m *moduleElementGroup) Delete() {
	m.group.Delete()
}
