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

func transform4x1(v *vec.VecN, m vec.Mat4) *vec.VecN {
	return v
}

func renderAsteroid(viewMatrix vec.Mat4, renderer *sdl.Renderer, object *world.WorldObject) {
	asteroid := []*vec.VecN{
		vec.NewVecNFromData([]float32{-10, 0, 0, 1}),
		vec.NewVecNFromData([]float32{0, -10, 0, 1}),
		vec.NewVecNFromData([]float32{0, 10, 0, 1}),
	}

	for _, v := range asteroid {
		transform4x1(v, viewMatrix)
	}

	// renderer.SetDrawColor(0, 255, 0, 255)
	// v0 := asteroid[0]
	// for i, v1 := asteroid[1:] {
	// 	v0r = v0.Raw()
	// 	v1r = v1.Raw()
	// 	renderer.DrawLine(v0r[0], v0r[1], v1r[0], v1r[1])
	// }

	// v0r := asteroid[0].Raw()
	// vnr := asteroid[len(asteroid)-1].Raw()

	// renderer.DrawLine(vnr[0], vnr[1], v0r[0], v0r[1])

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
