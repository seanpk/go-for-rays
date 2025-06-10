package geometry

import (
	"fmt"
)

// Vector:
// Represents a vector in 3D space (x, y, z). The w value (always 0.0) is included for interoperability with Points and using homogeneous coordinates (x, y, z, w).
type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{
		x: x,
		y: y,
		z: z,
	}
}

func ToVector(t HomogeneousTuple) Vector {
	return NewVector(t.X(), t.Y(), t.Z())
}

func (v Vector) X() float64 {
	return v.x
}

func (v Vector) Y() float64 {
	return v.y
}

func (v Vector) Z() float64 {
	return v.z
}

func (v Vector) W() float64 {
	return 0.0
}

func (v Vector) ToPoint() Point {
	return NewPoint(v.X(), v.Y(), v.Z())
}

func (v Vector) ToTuple() Tuple {
	return NewTuple(v.X(), v.Y(), v.Z(), v.W())
}

func (v Vector) Add(other HomogeneousTuple) Tuple {
	return Add(v, other)
}

func (v Vector) Subtract(other HomogeneousTuple) Tuple {
	return Subtract(v, other)
}

func (v Vector) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	return IsEqual(v, other, epsilon...)
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector(%f, %f, %f)", v.X(), v.Y(), v.Z())
}
