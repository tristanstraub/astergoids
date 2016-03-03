package world

import (
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
	ObjectType int
	// position
	X *vec.VecN
	// angular rotation
	R *vec.VecN
	// angular velocity (always points outwards)
	W *vec.VecN
	// velocity
	V *vec.VecN
}

type World struct {
	Objects []*WorldObject
}

func Vec() *vec.VecN {
	// TODO change to vec.Vec2()
	return vec.NewVecNFromData([]float32{0, 0, 0})
}

func VecFromData(data []float32) *vec.VecN {
	// TODO change to vec.Vec2()
	return vec.NewVecNFromData(data)
}

func cloneVec(v *vec.VecN) *vec.VecN {
	return vec.NewVecNFromData(v.Raw())
}

func NewWorldObject(otype int) *WorldObject {
	return &WorldObject{ObjectType: otype, X: Vec(), R: Vec(), W: Vec(), V: Vec()}
}

func AppendWorldObject(world *World, obj *WorldObject) {
	world.Objects = append(world.Objects, obj)
}

func clone(obj *WorldObject) *WorldObject {
	return &WorldObject{ObjectType: obj.ObjectType, X: cloneVec(obj.X), R: cloneVec(obj.R), W: cloneVec(obj.W), V: cloneVec(obj.V)}
}

func moveWorldObject(obj *WorldObject, delta float32, acceleration *vec.VecN, angularAcceleration *vec.VecN) *WorldObject {
	out := clone(obj)
	impulse := Vec()
	obj.V.Mul(impulse, float32(delta))
	obj.X.Add(out.X, impulse)

	angularImpulse := Vec()
	obj.W.Mul(angularImpulse, float32(delta))
	obj.R.Add(out.R, angularImpulse)

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
		acceleration := VecFromData([]float32{0, 0, 0})
		angularAcceleration := VecFromData([]float32{0, 0, 0})
		out.Objects = append(out.Objects, moveWorldObject(obj, delta, acceleration, angularAcceleration))
	}

	return &out
}
