package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

// func CalcSquare(sideLen float64, sidesNum #yourTypeNameHere#) float64 {

const SidesCircle = 0
const SidesTriangle = 3
const SidesSquare = 4

type myint int

func CalcSquare(sideLen float64, sidesNum myint) float64 {
	var ans float64
	if sidesNum == 0 {
		ans = sideLen * sideLen * math.Pi
	} else if sidesNum == 3 {
		ans = math.Sqrt(3) / 4 * sideLen * sideLen
	} else if sidesNum == 4 {
		ans = sideLen * sideLen
	}
	return ans
}
