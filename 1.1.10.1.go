package main

import (
	"fmt"
)

func main4() {
	t1129()
}

func t1129() {
	N := 9
	var a [][]bool
	for i := 0; i < N; i++ {
		a = append(a, make([]bool, N))
	}
	fmt.Println(a)

	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if gcd(i, j) {
				a[i][j] = true
				a[j][i] = true
			}
		}
	}
}

func gcd(a, b int) bool {
	return true
}
