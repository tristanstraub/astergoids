package world

import (
	"fmt"
	vec "github.com/go-gl/mathgl/mgl32"
)

type ObjectType int

const (
	SHIP     ObjectType = iota
	BULLET              = iota
	ASTEROID            = iota
)

type Interaction int

const (
	MOVE_RIGHT Interaction = iota
)

type WorldObject struct {
	ObjectType ObjectType
	// position
	X vec.Vec4
	// angular rotation
	R vec.Vec4
	// angular velocity (always points outwards)
	W vec.Vec4
	// velocity
	V vec.Vec4
}

type World struct {
	Objects []*WorldObject
}

func Vec() vec.Vec4 {
	return vec.Vec4{0, 0, 0, 1}
}

func NewWorldObject(otype ObjectType) *WorldObject {
	return &WorldObject{ObjectType: otype, X: Vec(), R: Vec(), W: Vec(), V: Vec()}
}

func AppendWorldObject(world *World, obj *WorldObject) {
	world.Objects = append(world.Objects, obj)
}

func clone(obj *WorldObject) *WorldObject {
	return &WorldObject{ObjectType: obj.ObjectType, X: obj.X, R: obj.R, W: obj.W, V: obj.V}
}

func moveWorldObject(obj *WorldObject, delta float32, acceleration vec.Vec4, angularAcceleration vec.Vec4) *WorldObject {
	out := clone(obj)

	v := obj.V.Mul(delta)
	w := obj.W.Mul(delta)

	mv := vec.Translate3D(v[0], v[1], v[2])
	mw := vec.Translate3D(w[0], w[1], w[2])

	out.X = vec.TransformCoordinate(obj.X.Vec3(), mv).Vec4(1)
	out.R = vec.TransformCoordinate(obj.R.Vec3(), mw).Vec4(1)

	fmt.Println(obj.X, obj.V, delta, out.X)

	return out
}

func NewWorld() *World {
	world := &World{}

	return world
}

func interactionType(interaction Interaction) ObjectType {
	var itype ObjectType
	switch interaction {
	case MOVE_RIGHT:
		itype = SHIP
	}
	return itype
}

func groupInteractionsByType(interactions []Interaction) map[ObjectType][]Interaction {
	grouped := make(map[ObjectType][]Interaction)
	for _, interaction := range interactions {
		itype := interactionType(interaction)
		_, exists := grouped[itype]

		if !exists {
			grouped[itype] = []Interaction{}
		}

		grouped[itype] = append(grouped[itype], interaction)
	}

	return grouped
}

func UpdateWorld(W *World, delta float32, interactions []Interaction) *World {
	out := World{}
	//groupedInteractions := groupInteractionsByType(interactions)

	for _, obj := range W.Objects {
		acceleration := vec.Vec4{0, 0, 0, 1}
		angularAcceleration := Vec()
		out.Objects = append(out.Objects, moveWorldObject(obj, 100, acceleration, angularAcceleration))
	}

	return &out
}
