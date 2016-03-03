package world

import (
	world "allthethings/astergoids/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateWorld(t *testing.T) {
	func() {
		// Test - UpdateWorld updates objects X1 = X0 + V0 * delta
		w := world.NewWorld()
		asteroid := world.NewWorldObject(world.ASTEROID)
		asteroid.X = world.VecFromData([]float32{10, 20})
		asteroid.V = world.VecFromData([]float32{1, 0.5})
		world.AppendWorldObject(w, asteroid)
		newworld := world.UpdateWorld(w, 0.5, []world.Interaction{})

		assert.Equal(t, newworld.Objects[0].X, world.VecFromData([]float32{10.5, 20.25}),
			"Position should be updated")
	}()

	func() {
		// Test - UpdateWorld updates objects R1 = R0 + W0 * delta
		w := world.NewWorld()
		asteroid := world.NewWorldObject(world.ASTEROID)
		asteroid.R = world.VecFromData([]float32{10, 20})
		asteroid.W = world.VecFromData([]float32{1, 0.5})
		world.AppendWorldObject(w, asteroid)
		newworld := world.UpdateWorld(w, 0.5, []world.Interaction{})

		assert.Equal(t, newworld.Objects[0].R, world.VecFromData([]float32{10.5, 20.25}),
			"Position should be updated")
	}()
}
