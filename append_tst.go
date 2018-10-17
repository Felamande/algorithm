package main

import (
	"fmt"
	"time"
)

func main() {
	for _, N := range []int{100000, 2000000, 4000000, 8000000, 16000000, 32000000, 64000000} {
		a := make([]int, N)
		t1 := time.Now()
		for i := 0; i < N; i++ {
			a[i] = i
		}
		dt := time.Since(t1)

		fmt.Printf("N=%d t=%v\n", N, dt)
	}
}
