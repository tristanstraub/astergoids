package render

import (
	"allthethings/astergoids/world"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	Quit  bool
	Keys  map[sdl.Keycode]bool
	World *world.World
}

func NewGame() *Game {
	return &Game{Quit: false, Keys: make(map[sdl.Keycode]bool), World: world.NewWorld()}
}

func Run() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Astergoids", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	game := NewGame()

	renderer, _ := sdl.CreateRenderer(window, -1, 0)

	lastTicks := sdl.GetTicks()

	for !game.Quit {
		ticks := sdl.GetTicks()
		delta := lastTicks - ticks
		lastTicks = ticks

		events(game)

		interactions := []world.Interaction{}

		if game.Keys[sdl.K_ESCAPE] {
			game.Quit = true
		}

		if game.Keys[sdl.K_w] {
			interactions = append(interactions, world.MOVE_RIGHT)
		}

		game.World = world.UpdateWorld(game.World, float32(delta), interactions)

		render(renderer, game)
		renderer.Present()
	}

	sdl.Quit()
}

func events(game *Game) {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.QuitEvent:
			game.Quit = true
		case *sdl.KeyDownEvent:
			k := e.(*sdl.KeyDownEvent)
			game.Keys[k.Keysym.Sym] = true
		case *sdl.KeyUpEvent:
			k := e.(*sdl.KeyUpEvent)
			game.Keys[k.Keysym.Sym] = false
		default:
		}
	}
}

func render(renderer *sdl.Renderer, game *Game) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.SetDrawColor(0, 255, 0, 255)

	//game.World.Objects

	renderer.DrawLine(0, 0, 200, 200)
}
