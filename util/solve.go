package util

import "math"

// Given a target point assumed to be on the same plane this finds the correct velocity to reach that point
func SolveQuadraticVelocity(v, dx, g float64) (float64, float64, bool) {
	v2 := v * v
	disc := (v2 * v2) - 4*(dx*dx)*(g*g)
	if disc < 0 {
		return 0, 0, false
	}
	vy := math.Sqrt(0.5 * (v2 + math.Sqrt(disc)))
	vx := (dx / math.Abs(dx)) * math.Sqrt(v2-vy*vy)

	return vx, vy, true
}
