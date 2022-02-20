package gui

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/util"
)

type Overlay struct {
	*BaseElement
	overlay *util.ListNode
}

func CreateOverlay(ctx *graphics.Context, pos mgl32.Vec2, layer int) *Overlay {
	b := CreateBaseElement(ctx, mgl32.Vec4{pos.X(), pos.Y(), 0, 0}, 1, layer)
	return &Overlay{
		BaseElement: b,
	}
}

func (o *Overlay) SetOverlay(vertLayout, texLayout mgl32.Vec4, tex *gl.Texture) {
	o.ctx.AddJob(func() {
		v, t, err := graphics.Rectangle(
			vertLayout.X(),
			vertLayout.Y(),
			vertLayout.Z(),
			vertLayout.W(),
			int(texLayout.X()),
			int(texLayout.Y()),
			int(texLayout.Z()),
			int(texLayout.W()),
			tex,
		)
		if err != nil {
			panic(err)
		}

		r := o.renderer
		r.SetTexture(tex)

		if o.overlay != nil {
			if err := r.SetVertices(v, t, o.overlay); err != nil {
				panic(err)
			}
		} else {
			o.overlay, err = r.AllocateAndSetVertices(v, t)
			if err != nil {
				panic(err)
			}
		}
		r.Update()
	})
}

func (o *Overlay) ClearOverlay() {
	if o.overlay == nil {
		return
	}

	o.ctx.AddJob(func() {
		r := o.renderer
		if err := r.ClearVertices(o.overlay); err != nil {
			panic(err)
		}
		o.overlay = nil
		r.Update()
	})
}
