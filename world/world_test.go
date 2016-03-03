package world

import (
	vec "github.com/go-gl/mathgl/mgl32"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateWorld(t *testing.T) {
	func() {
		// Test - UpdateWorld updates objects X1 = X0 + V0 * delta
		w := NewWorld()
		asteroid := NewWorldObject(ASTEROID)
		asteroid.X = vec.Vec4{10, 20, 0, 1}
		asteroid.V = vec.Vec4{1, 0.5, 0, 1}
		AppendWorldObject(w, asteroid)
		newWorld := UpdateWorld(w, 0.5, []Interaction{})

		assert.Equal(t, newWorld.Objects[0].X, vec.Vec4{10.5, 20.25, 0, 1},
			"Position should be updated")
	}()

	func() {
		// Test - UpdateWorld updates objects R1 = R0 + W0 * delta
		w := NewWorld()
		asteroid := NewWorldObject(ASTEROID)
		asteroid.R = vec.Vec4{10, 20, 0, 1}
		asteroid.W = vec.Vec4{1, 0.5, 0, 1}
		AppendWorldObject(w, asteroid)
		newWorld := UpdateWorld(w, 0.5, []Interaction{})

		assert.Equal(t, newWorld.Objects[0].R, vec.Vec4{10.5, 20.25, 0, 1},
			"Position should be updated")
	}()
}
