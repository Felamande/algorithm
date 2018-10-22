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

func TestRandNorm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	max := 100
	m := make([]int, max)
	for i := 0; i < 100000; i++ {
		r := randNormInt(max - 1)
		m[r]++
	}
	t.Log(m)
}

func BenchmarkRandNorm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randNormInt(10000)
	}
}

func BenchmarkRandPoissonL10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoissonInt(10)
	}
}

func BenchmarkRandPoissonL100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoissonInt(100)
	}
}

func BenchmarkRandPoissonL200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoissonInt(200)
	}
}

func TestRandPoiss(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	lmd := float64(20)
	m := make([]int, 100)
	for i := 0; i < 100000; i++ {
		r := randPoissonInt(lmd)
		m[r]++
	}
	t.Log(m)
}

func BenchmarkRandPoisson2L10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(10)
	}
}

func BenchmarkRandPoisson2L100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(100)
	}
}

func BenchmarkRandPoisson2L200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(200)
	}
}

func BenchmarkRandPoisson2L1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(730)
	}
}

func TestRandPoiss2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	lmd := float64(20)
	m := make([]int, 100)
	for i := 0; i < 100000; i++ {
		r := randPoisson2Int(lmd)
		m[r]++
	}
	t.Log(m)
}
