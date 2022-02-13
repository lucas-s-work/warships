package game

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
	"github.com/lucas-s-work/warships/game/ship"
	"github.com/lucas-s-work/warships/game/world"
)

const (
	backgroundLocation = "./textures/background.png"
)

type Game struct {
	ctx         *graphics.Context
	window      *gl.Window
	background  *renderers.Translational
	entities    []*entityState
	newEntities []*entityState

	player *Player
}

type entityState struct {
	e           world.Entity
	initialized bool
}

func CreateGame(ctx *graphics.Context, window *gl.Window) *Game {
	g := &Game{
		ctx:      ctx,
		window:   window,
		entities: make([]*entityState, 0),
	}
	g.createBackground()

	b := ship.CreateBattleship(g, mgl32.Vec2{50, 50})
	b2 := ship.CreateBattleship(g, mgl32.Vec2{100, 200})
	g.AttachEntity(b)
	g.AttachEntity(b2)

	p := CreatePlayer(window, g)
	g.player = p

	return g
}

func (g *Game) Window() *gl.Window {
	return g.window
}

func (g *Game) Context() *graphics.Context {
	return g.ctx
}

func (g *Game) AttachEntity(e world.Entity) {
	g.newEntities = append(g.newEntities, &entityState{
		initialized: false,
		e:           e,
	})
}

func (g *Game) EntitiesUnderPoint(point mgl32.Vec2) []world.Entity {
	entities := make([]world.Entity, 0)

	for _, e := range g.entities {
		if e.e.InBounds(point) {
			entities = append(entities, e.e)
		}
	}

	return entities
}

func (g *Game) SetCamera(p mgl32.Vec2) {
	for _, e := range g.entities {
		e.e.SetCamera(p)
	}
}

func (g *Game) createBackground() {
	if g.background != nil {
		g.background.Delete()
	}

	g.ctx.AddJob(func() {
		background, err := renderers.CreateTranslationalRenderer(g.window, backgroundLocation, 6)
		g.ctx.Attach(background, world.BACKGROUND_LAYER)
		if err != nil {
			panic(err)
		}

		v, t, _ := graphics.Rectangle(0, 0, g.window.Width, g.window.Height, 1, 0, 1, 1, background.Texture())
		background.AllocateAndSetVertices(v, t)
		background.Update()

		g.background = background
	})
}

func (g *Game) Tick() {
	// Entities are given the first tick to initialize
	// This allows graphics synchronization and any interactinos
	// Post world loading
	for _, s := range g.entities {
		if s == nil {
			continue
		}
		if !s.initialized {
			s.initialized = true
			s.e.Init()
		} else {
			s.e.Tick()
		}
	}

	g.player.Tick()

	g.entities = append(g.entities, g.newEntities...)
	g.newEntities = []*entityState{}
}
