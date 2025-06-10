package geometry

import (
	"fmt"
)

// Point:
// Represents a point in 3D space (x, y, z). The w value (always 1.0) is included for interoperability with Vectors and using homogeneous coordinates (x, y, z, w).
type Point struct {
	x, y, z float64
}

func NewPoint(x, y, z float64) Point {
	return Point{
		x: x,
		y: y,
		z: z,
	}
}

func ToPoint(t HomogeneousTuple) Point {
	return NewPoint(t.X(), t.Y(), t.Z())
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) Z() float64 {
	return p.z
}

func (p Point) W() float64 {
	return 1.0
}

func (p Point) ToVector() Vector {
	return NewVector(p.X(), p.Y(), p.Z())
}

func (p Point) ToTuple() Tuple {
	return NewTuple(p.X(), p.Y(), p.Z(), p.W())
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%f, %f, %f)", p.X(), p.Y(), p.Z())
}

func (p Point) Equals(other HomogeneousTuple, epsilon ...float64) bool {
	return IsEqual(p, other, epsilon...)
}
