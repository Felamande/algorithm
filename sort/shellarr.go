//+build ignore
package main

import (
	"math"
)

func invert(a []int) {
	l := len(a)
	for i := 0; i <= l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
}

func printA(a []int) {
	for _, v := range a {
		print(v, ",")
	}
}

func genSeq(f func(k int) int) (a []int) {
	h := 0
	for k := 0; h <= math.MaxInt32; k++ {
		h = f(k)
		a = append(a, h)
	}
	invert(a)
	return a
}

func seqt(t float64) func(k int) int {
	return func(k int) int {
		return int(math.Floor(math.Pow(t, float64(k))))
	}
}
