package main

import (
	"allthethings/astergoids/render"
	"allthethings/astergoids/world"
	vec "github.com/go-gl/mathgl/mgl32"
)

func main() {
	game := render.NewGame()
	asteroid := world.NewWorldObject(world.ASTEROID)
	asteroid.V = vec.Vec4{1, 1, 0, 1}

	world.AppendWorldObject(game.World, asteroid)

	render.Run(game)
}
