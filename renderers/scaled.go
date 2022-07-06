package renderers

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
)

const (
	scaledVertShader = "./shaders/rotational/scaled/vertex.vert"
	scaledFragShader = "./shaders/rotational/frag.frag"
)

const (
	dimensionUniform       = "dimension"
	tranlsationUniform     = "trans"
	rotation1AngleUniform  = "rot1angle"
	rotation1CenterUniform = "rot1center"
	scaleUniform           = "scale"
)

type Scaled struct {
	*renderers.Renderer2D
	shader *shader.Program
}

func CreateScaledRenderer(window *ggl.Window, texture string, size int32) (*Scaled, error) {
	p := shader.CreateProgram(0)
	if err := p.LoadShader(scaledVertShader, gl.VERTEX_SHADER); err != nil {
		return nil, err
	}
	if err := p.LoadShader(scaledFragShader, gl.FRAGMENT_SHADER); err != nil {
		return nil, err
	}
	if err := p.Link(); err != nil {
		return nil, err
	}

	t := mgl32.Vec2{}
	if err := p.AttachUniform(tranlsationUniform, t); err != nil {
		return nil, err
	}
	var rAngle float32 = 0
	if err := p.AttachUniform(rotation1AngleUniform, rAngle); err != nil {
		return nil, err
	}
	rCenter := mgl32.Vec2{}
	if err := p.AttachUniform(rotation1CenterUniform, rCenter); err != nil {
		return nil, err
	}
	if err := p.AttachUniform(dimensionUniform, mgl32.Vec2{window.Width, window.Height}); err != nil {
		return nil, err
	}
	var scale float32 = 1
	if err := p.AttachUniform(scaleUniform, scale); err != nil {
		return nil, err
	}

	r, err := renderers.CreateRenderer2D(window, texture, size, p)
	if err != nil {
		return nil, err
	}

	return &Scaled{
		Renderer2D: r,
		shader:     p,
	}, nil
}

func (r *Scaled) SetTranslation(translation mgl32.Vec2) {
	r.shader.SetUniform(tranlsationUniform, translation)
}

func (r *Scaled) SetRotation1(angle float32, center mgl32.Vec2) {
	r.shader.SetUniform(rotation1AngleUniform, angle)
	r.shader.SetUniform(rotation1CenterUniform, center)
}

func (r *Scaled) SetScale(scale float32) {
	r.shader.SetUniform(scaleUniform, scale)
}
