package gui

import (
	"github.com/go-gl/mathgl/mgl32"
)

type ElementGroup struct {
	gui      GUI
	elements []Element
	position mgl32.Vec2
}

/*
An element group is not an element, it handles state management for a set of elements
Elements shouldn't be added or removed from an element group, if you need to change that you
should create a new group.
*/

func CreateElementGroup(gui GUI, position mgl32.Vec2) *ElementGroup {
	return &ElementGroup{
		gui:      gui,
		elements: make([]Element, 0),
		position: position,
	}
}

func (e *ElementGroup) AttachElement(el Element) error {
	if err := e.gui.AttachElement(el); err != nil {
		return err
	}
	e.elements = append(e.elements, el)

	return nil
}

func (e *ElementGroup) DeleteElements() {
	for _, el := range e.elements {
		if el == nil {
			continue
		}
		e.gui.DetachElement(el)
		el.Delete()
	}

	e.elements = make([]Element, 0)
}

func (e *ElementGroup) Element(id int) Element {
	return e.elements[id]
}

func (e *ElementGroup) Delete() {
	for _, el := range e.elements {
		if el == nil {
			continue
		}
		e.gui.DetachElement(el)
		el.Delete()
	}

	e.elements = make([]Element, 0)
}

func (e *ElementGroup) RelativeDimensions(dim mgl32.Vec4) mgl32.Vec4 {
	return mgl32.Vec4{
		dim.X() + e.position.X(),
		dim.Y() + e.position.Y(),
		dim.Z(),
		dim.W(),
	}
}

func (e *ElementGroup) SetPosition(pos mgl32.Vec2) {
	e.position = pos
	for _, el := range e.elements {
		el.Renderer().SetTranslation(pos)
	}
}

func (e *ElementGroup) SetActive(active bool) {
	for _, el := range e.elements {
		el.SetActive(active)
	}
}
