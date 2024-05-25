package app

import (
	"math"
	"math/rand"
	"time"
)

// GaussianNoise generates a normally distributed random number with a specified mean and variance.
//
// This function implements the Box-Muller method to generate a normally distributed random number.
// It first generates two uniformly distributed random numbers, then transforms them into a pair of independent standard normally distributed (normal) random numbers.
// One of these numbers is then scaled by the square root of the variance and shifted by the mean to produce the final result.
// This method ensures that the generated number follows a normal distribution with the specified mean and variance.
//
// Parameters:
//
//	mean (float64): The mean of the normal distribution from which the random number is generated.
//	variance (float64): The variance of the normal distribution from which the random number is generated.
//
// Returns:
//
//	float64: The generated normally distributed random number.
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
