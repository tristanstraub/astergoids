package render

import (
	"allthethings/astergoids/world"
	vec "github.com/go-gl/mathgl/mgl32"
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

func Run(game *Game) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Astergoids", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, _ := sdl.CreateRenderer(window, -1, 0)

	lastTicks := sdl.GetTicks() / 1000

	for !game.Quit {
		ticks := sdl.GetTicks() / 1000
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

func renderAsteroid(viewMatrix vec.Mat4, renderer *sdl.Renderer, object *world.WorldObject) {
	asteroid := []vec.Vec4{
		vec.Vec4{-100, 0, 0, 1},
		vec.Vec4{0, 100, 0, 1},
		vec.Vec4{100, 0, 0, 1},
	}

	transformed := make([]vec.Vec4, len(asteroid))

	for i, v := range asteroid {
		transformed[i] = viewMatrix.Mul4x1(v)
	}

	renderer.SetDrawColor(0, 255, 0, 255)
	v0 := transformed[0]
	for _, v1 := range transformed[1:] {
		renderer.DrawLine(int(v0[0]), int(v0[1]), int(v1[0]), int(v1[1]))
	}

	v0 = asteroid[0]
	vn := asteroid[len(asteroid)-1]

	renderer.DrawLine(int(vn[0]), int(vn[1]), int(v0[0]), int(v0[1]))
}

func renderShip(viewMatrix vec.Mat4, renderer *sdl.Renderer, object *world.WorldObject) {
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawLine(0, 0, 200, 200)
}

func renderWorldObjects(viewMatrix vec.Mat4, renderer *sdl.Renderer, objects []*world.WorldObject) {
	for _, o := range objects {
		switch o.ObjectType {
		case world.SHIP:
			renderShip(viewMatrix, renderer, o)
		case world.ASTEROID:
			renderAsteroid(viewMatrix, renderer, o)
		}
	}
}

func render(renderer *sdl.Renderer, game *Game) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	viewMatrix := vec.Ident2().Mat4()

	renderWorldObjects(viewMatrix, renderer, game.World.Objects)
}
