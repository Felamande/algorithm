package main

import (
	"math"
	"math/rand"
)

func randNormInt(max int) int {
	u1 := rand.Float64()
	u2 := rand.Float64()

	R := math.Sqrt(2 * math.Log(1/u1))
	th := 2 * math.Pi * u2

	r := R * math.Cos(th)

	var f4Sigma float64 = 4

	mag := float64(max) / (2 * f4Sigma)

	ri := int((r + 4) * mag)
	if ri > max {
		return max
	}
	if ri < 0 {
		return 0
	}

	return ri
}

func randPoissonInt(lmd float64) int {
	expLmd := math.Exp(-lmd)

	PDF := func(k int) float64 {
		var pdf float64 = 1
		for i := 1; i <= k; i++ {
			pdf *= (lmd / float64(i))
		}
		return pdf * expLmd
	}

	u := rand.Float64()

	cdf := expLmd
	k := 0
	for u >= cdf {
		k++
		cdf += PDF(k)
	}
	return k
}

func randPoisson2Int(lmd float64) int {
	limit := math.Exp(-lmd)
	prod := rand.Float64()

	var n int
	for n = 0; prod >= limit; n++ {
		prod *= rand.Float64()
	}
	return n
}

func randGeoInt(p float64) int {
	if p > 1 || p < 0 {
		return 0
	}

	n := 0

	for {
		r := rand.Float64()
		n++
		if r <= p {
			break
		}
	}
	return n
}

func randBinaryInt(n int, p float64) int {
	return 0
}
