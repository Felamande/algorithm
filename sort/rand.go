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

func randPoisson3Int(lmd float64) int {
	var r, x, m float64

	pi := math.Pi
	sqrtLmd := math.Sqrt(lmd)
	logLmd := math.Log(lmd)

	var gx, fm float64

	randf := rand.Float64
	for {
		for {
			x = lmd + sqrtLmd*math.Tan(pi*(randf()-1.0/2.0))
			if x >= 0 {
				break
			}
		}
		gx = sqrtLmd / (pi * ((x-lmd)*(x-lmd) + lmd))
		m = math.Floor(x)
		lgammaMp1, sign := math.Lgamma(m + 1)
		fm = math.Exp(m*logLmd - lmd - lgammaMp1*float64(sign))
		r = fm / gx / 2.4
		if randf() <= r {
			break
		}
	}
	return int(m)
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

func randGeo2Int(p float64) int {
	ln := math.Log
	return int(math.Ceil(ln(1-rand.Float64()) / ln(1-p)))
}

func randBinaryInt(n int, p float64) int {
	return 0
}
