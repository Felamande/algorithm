package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestGetName(t *testing.T) {
	// ssa := "sddd"
	t.Log(getName(bubbleSort))
}

func TestSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	N := 500
	ss := makeInts(func(int) int { return rand.Intn(5 * N) }, N)

	for _, f := range []sortFn{sysSort, selSort, bubbleSort, insertSort, shSortd2, shSortAr1} {
		s := ss.Clone()
		f(s)
		if !sorted(s) {
			t.Fatalf("sort fail %s:%v", getName(f), s)
		}
	}

}

func TestTime(t *testing.T) {
	t1_ := time.Now()
	t1 := t1_.UnixNano()
	time.Sleep(time.Microsecond * 200)
	t2 := time.Now().UnixNano()
	t.Logf("%v, %v, %v\n", t1, t2, t2-t1)
}

func TestGenArr(t *testing.T) {
	el := 1
	for k := uint(1); el <= math.MaxInt32; k++ {
		t.Logf("%v, ", el)
		el = 9*(1<<(2*k)) - 9*(1<<k) + 1
	}
}

func TestRand(t *testing.T) {
	M := 300
	N := M * 10000
	rand.Seed(time.Now().UnixNano())
	cnt := make([]int, M)
	for i := 0; i < N; i++ {
		r := randn(float64(M), gauss(0.33, 0))
		cnt[int(r)]++
	}
	t.Log(cnt)

}
