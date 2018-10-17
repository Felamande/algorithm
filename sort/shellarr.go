//+build ignore
package main

import (
	"math"
)

func main() {
	h := 0
	var a []int
	for i := 0; h < math.MaxInt32; i++ {
		h = int(math.Ceil(math.Pow(2.48, float64(i))))
		a = append(a, h)
	}
	invert(a)
	for _, v := range a {
		print(v, ",")
	}
}

func invert(a []int) {
	l := len(a)
	for i := 0; i <= l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
}
