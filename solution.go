package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	var result float64
	if sidesNum == SidesTriangle {
		result = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		result = math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		result = math.Pi * math.Pow(sideLen, 2)
	} else {
		result = 0
	}
	return result
}
