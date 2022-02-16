package gui

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	GUI_BOTTOM_LAYER = iota
	GUI_BASE_LAYER   = iota
)

type GUI interface {
	Click(pos mgl32.Vec2)
	Tick()
	Delete()
	AttachElement(Element) error
	DetachElement(Element)
}

type BaseGUI struct {
	elements []Element
}

func CreateBaseGUI(size int) *BaseGUI {
	return &BaseGUI{
		elements: make([]Element, size),
	}
}

func (b *BaseGUI) Click(pos mgl32.Vec2) {
	for _, el := range b.elements {
		if el.InBounds(pos) {
			el.Click()
		}
	}
}

func (b *BaseGUI) AttachElement(el Element) error {
	for i, e := range b.elements {
		if e == nil {
			b.elements[i] = el
			el.SetID(i)
			return nil
		}
	}

	return fmt.Errorf("cannot attach element to GUI, no free slots")
}

func (b *BaseGUI) DetachElement(el Element) {
	for i, e := range b.elements {
		if e == el {
			b.elements[i] = nil
			return
		}
	}
}

func (b *BaseGUI) Delete() {
	for _, e := range b.elements {
		if e != nil {
			e.Delete()
		}
	}
}

func (*BaseGUI) Tick() {}
