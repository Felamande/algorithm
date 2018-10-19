package main

import (
	"math"
	"math/rand"
)

func randn(max float64, f func(float64) float64) int {
	x := rand.Float64()
	return int(f(x) * max)
}

func gauss(sig float64, mu float64) func(float64) float64 {
	_1_d_sqrt_2pi_sig := float64(1) / (math.Sqrt(2*math.Pi) * sig)

	_2sig2 := 2 * sig * sig

	return func(x float64) float64 {
		if x < 0.0 || x > 1.0 {
			return 0.0
		}
		y := _1_d_sqrt_2pi_sig * math.Exp(-1*(x-mu)*(x-mu)/_2sig2)
		return y
	}
}
