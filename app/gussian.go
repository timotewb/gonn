package app

import (
	"math"
	"math/rand"
	"time"
)

func GaussianNoise(mean float64, variance float64) float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	u1 := rand.Float64()
	u2 := rand.Float64()
	r := math.Sqrt(-2 * math.Log(u1))
	theta := 2 * math.Pi * u2
	z := r * math.Cos(theta)

	// Adjust the generated number to have the desired mean and variance.
	adjustedZ := z*math.Sqrt(variance) + mean
	return adjustedZ
}
