package main

import (
	"math"

	// "math/big"
	randc "crypto/rand"
	"fmt"
	"math/big"
	randm "math/rand"
	"time"
)

//intersec 1.4.12
func intersec(a, b []int) (r []int) {
	alen := len(a)
	blen := len(b)
	r_idx := 0
	a_idx := 0
	b_idx := 0
	last_val := 0
	for {
		if a_idx > alen-1 || b_idx > blen-1 {
			break
		}

		if a[a_idx] < b[b_idx] {
			a_idx++
			continue
		} else if a[a_idx] > b[b_idx] {
			b_idx++
			continue
		}

		// a[a_idx]==b[b_idx]
		dup_val := a[a_idx]
		if dup_val == last_val && r_idx != 0 {
			a_idx++
			b_idx++
			continue
		}
		r = append(r, dup_val)
		last_val = dup_val

		r_idx++
		a_idx++
		b_idx++
	}
	return
}

//BinSearch 1.1.10
func BinSearch(a []int, key int) int {
	center, left, right := 0, 0, len(a)-1

	for left < right {
		center = (left + right) / 2
		if a[center] > key {
			right = center - 1
		} else if a[center] < key {
			left = center + 1
		} else {
			return center
		}
	}
	return -1
}

var fibs = [48]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597,
	2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229,
	832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986,
	102334155, 165580141, 267914296, 433494437, 701408733, 1134903170, 1836311903, 1<<32 - 1}

//FiboSearch 1.4.22
func FiboSearch(a []int, key int) int {
	alen := len(a)
	k := 0
	start := 0

	for idx, v := range fibs {
		if alen > v {
			k = idx
		}
	}

	for k > 1 {
		fibkm1 := start + fibs[k-1]
		if a[fibkm1] > key {
			k = k - 2
		} else if a[fibkm1] < key {
			start = start + fibs[k-2]
		} else {
			return fibkm1
		}
	}
	return -1

}

func randcf(min, max int64) int64 {
	i, _ := randc.Int(randc.Reader, big.NewInt(max-min))
	r := i.Int64()
	return r + min
}

func randmf(min, max int64) int64 {
	randm.Seed(time.Now().UnixNano())
	a := randm.Int63n(max - min)
	return a + min
}
func main21() {
	randfunc := randcf

	min := int64(0)
	// max := int64(100)

	count := 4000
	total := 0
	var re []float64
	maxes := []int64{}
	for max := int64(10); max <= 3200; max *= 2 {
		maxes = append(maxes, max)
	}
	for _, max := range maxes {
		for c := 0; c < count; c++ {
			dup := make(map[int64]bool)

			for i := 0; i < 1<<32-1; i++ {
				r := randfunc(min, max)
				// println(r)
				if len(dup) == int(max) {
					// println(i)
					total += i
					break
				}
				dup[r] = true
			}
		}
		r := float64(total) / float64(count)
		re = append(re, r)
		fmt.Printf("max=%d n=%f\n", max, r)

	}
	for i := 0; i < len(re)-1; i++ {
		p := math.Log(re[i+1]) / math.Log(re[i])
		fmt.Print(p, " ")
	}
}

func benchmark2(f func(), use bool) time.Duration {
	var t time.Time
	var dt time.Duration
	if use {
		t = time.Now()
	}
	f()
	if use {
		dt = time.Since(t)
	}
	return dt
}
