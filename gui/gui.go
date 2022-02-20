package gui

import (
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
	elements       []Element
	newElements    []Element
	elementsToInit []int
}

func CreateBaseGUI(size int) *BaseGUI {
	return &BaseGUI{
		elements:       make([]Element, size),
		newElements:    make([]Element, 0),
		elementsToInit: make([]int, 0),
	}
}

func (b *BaseGUI) Click(pos mgl32.Vec2) {
	for _, el := range b.elements {
		if el == nil || !el.Active() {
			continue
		}
		if el.InBounds(pos) {
			el.Click()
		}
	}
}

func (b *BaseGUI) AttachElement(el Element) {
	b.newElements = append(b.newElements, el)
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

func (b *BaseGUI) Tick() {
	if len(b.elementsToInit) > 0 {
		for _, i := range b.elementsToInit {
			b.elements[i].Init()
		}
		b.elementsToInit = []int{}
	}

	if len(b.newElements) > 0 {
		for _, e := range b.newElements {
			found := false
			for i, el := range b.elements {
				if el == nil {
					b.elements[i] = e
					e.SetID(i)
					b.elementsToInit = append(b.elementsToInit, i)
					found = true
					break
				}
			}

			if !found {
				panic("cannot attach element to gui, no free slots")
			}
		}
		b.newElements = []Element{}
	}

}
