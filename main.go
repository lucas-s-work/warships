package main

import (
	"runtime"
	"time"

	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/warships/game"
)

const (
	width          = 800
	height         = 600
	updateRate     = 60
	updateInterval = (1000 / updateRate) * time.Millisecond
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// Initialize the world, window and graphics context
	window, err := gl.CreateWindow(width, height, "Warships")
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	if err = gl.GlInit(); err != nil {
		panic(err)
	}

	ctx := graphics.CreateContext(window)
	ctxSync := ctx.GetSync()

	g := game.CreateGame(ctx)

	doneSync := startTick(g, ctxSync)
	ticker := time.NewTicker(updateInterval)

	for !window.ShouldClose() {
		<-ticker.C
		ctx.Render()
	}

	doneSync <- struct{}{}
}

func startTick(g *game.Game, ctxSync chan<- struct{}) chan<- struct{} {
	doneSync := make(chan struct{})

	go func() {
		for {
			select {
			case ctxSync <- struct{}{}:
				g.Tick()
			case <-doneSync:
				return
			}
		}
	}()

	return doneSync
}
