package main

import (
	"allthethings/astergoids/render"
	"allthethings/astergoids/world"
)

func main() {
	game := render.NewGame()
	world.AppendWorldObject(game.World, world.NewWorldObject(world.ASTEROID))

	render.Run(game)
}
