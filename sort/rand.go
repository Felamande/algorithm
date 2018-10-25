package main

import (
	"math"
	"math/rand"
)

type rndFn func(max int) int

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

func randPoissonInt(lmd float64) func(max int) int {
	return func(int) int {
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
}

func randPoisson2Int(lmd float64) func(max int) int {
	limit := math.Exp(-lmd)

	return func(int) int {
		prod := rand.Float64()
		var n int
		for n = 0; prod >= limit; n++ {
			prod *= rand.Float64()
		}
		return n
	}
}

func randPoisson3Int(lmd float64) func(max int) int {

	pi := math.Pi
	sqrtLmd := math.Sqrt(lmd)
	logLmd := math.Log(lmd)
	randf := rand.Float64

	return func(int) int {
		var r, x, m float64
		var gx, fm float64

		for {
			for {
				rnd := randf()
				x = lmd + sqrtLmd*math.Tan(pi*(rnd-1.0/2.0))
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
}

func randGeoInt(p float64) func(max int) int {
	return func(int) int {
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
}

func randGeo2Int(p float64) int {
	ln := math.Log
	return int(math.Ceil(ln(1-rand.Float64()) / ln(1-p)))
}

func randBinaryInt(n int, p float64) int {
	return 0
}

func randMap(smpSize int, rndMax int, rf func(int) int) (m map[int]int) {
	m = make(map[int]int)
	for i := 0; i < smpSize; i++ {
		m[rf(rndMax)]++
	}
	return m
}

func ExpAndVar(smpSize int, m map[int]int) (Ex, Dx float64) {
	var Ex2 float64
	for x, n := range m {
		p := float64(n) / float64(smpSize)
		Ex += float64(x) * p
		Ex2 += float64(x) * float64(x) * p
	}
	return Ex, Ex2 - Ex*Ex
}

func expAndVar(arr []float64) (Ex, Dx float64) {
	n := len(arr)
	var Ex2 float64
	for _, v := range arr {
		p := 1 / float64(n)
		Ex += v * p
		Ex2 += v * v * p
	}
	return Ex, Ex2 - Ex*Ex
}

func randDiscreteInt(pdf []float64, rndf []rndFn) func(max int) int {

	cdf := make([]float64, len(pdf)+1)

	var cdfVal float64
	cdf[0] = cdfVal
	for idx, p := range pdf {
		cdfVal += p
		cdf[idx+1] = cdfVal
	}

	return func(max int) int {
		u := rand.Float64()
		for idx := 0; idx < len(pdf); idx++ {
			if cdf[idx] <= u && u < cdf[idx+1] {
				return rndf[idx](max)
			}
		}
		return -1
	}
}

func randAscendInt(step int, rf rndFn) func(idx int) int {
	r := 0

	return func(idx int) int {
		r += rf(step)
		return r
	}
}

func randNormZigguratInt() int {
	return 0
}

func randNorm() float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()

	R := math.Sqrt(2 * math.Log(1/u1))
	th := 2 * math.Pi * u2

	r := R * math.Cos(th)

	return r
}
