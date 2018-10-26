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
		_ = randPoissonInt(10)(0)
	}
}

func BenchmarkRandPoissonL100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoissonInt(100)(0)
	}
}

func BenchmarkRandPoissonL200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoissonInt(200)(0)
	}
}

func TestRandPoiss(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	lmd := float64(20)
	m := make([]int, 100)
	for i := 0; i < 100000; i++ {
		r := randPoissonInt(lmd)(0)
		m[r]++
	}
	t.Log(m)
}

func BenchmarkRandPoisson2L10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(10)(0)
	}
}

func BenchmarkRandPoisson2L100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(100)(0)
	}
}

func BenchmarkRandPoisson2L200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(200)(0)
	}
}

func BenchmarkRandPoisson2L500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson2Int(500)(0)
	}
}

func TestRandPoiss2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	lmd := float64(20)
	m := make([]int, 100)
	for i := 0; i < 100000; i++ {
		r := randPoisson2Int(lmd)(0)
		m[r]++
	}
	t.Log(m)
}

func BenchmarkRandGeoP01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randGeoInt(0.1)(0)
	}
}
func BenchmarkRandGeoP001(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randGeoInt(0.01)(0)
	}
}

func BenchmarkRandGeoP099(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randGeoInt(0.99)(0)
	}
}

func TestRandGeo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	p := 0.9
	m := make([]int, 100)
	for i := 0; i < 100000000; i++ {
		r := randGeoInt(p)(0)
		m[r]++
	}
	t.Log(m)
}

func BenchmarkRandPoisson3L10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson3Int(10)(0)
	}
}

func BenchmarkRandPoisson3L100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson3Int(100)(0)
	}
}

func BenchmarkRandPoisson3L200(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson3Int(200)(0)
	}
}

func BenchmarkRandPoisson3L500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randPoisson3Int(500)(0)
	}
}

func TestRandPoiss3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	lmd := float64(20)
	m := make([]int, 100)
	for i := 0; i < 100000; i++ {
		r := randPoisson3Int(lmd)(0)
		m[r]++
	}
	t.Log(m)
}

func BenchmarkRandGeo2P01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randGeo2Int(0.1)
	}
}
func BenchmarkRandGeo2P001(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randGeo2Int(0.01)
	}
}

func BenchmarkRandGeo2P099(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randGeo2Int(0.99)
	}
}

func TestRandGeo2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	p := 0.9
	m := make([]int, 100)
	for i := 0; i < 100000000; i++ {
		r := randGeo2Int(p)
		m[r]++
	}
	t.Log(m)
}
func TestExpAndVar(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	smpSize := 10000
	rndf := randPoisson2Int(20)

	m := randMap(smpSize, 1000, rndf)
	ex, dx := ExpAndVar(smpSize, m)
	t.Logf("Ex=%f, Dx=%f\n", ex, dx)

	m2 := randMap(smpSize, 1000, rndf)
	ex2, dx2 := ExpAndVar(smpSize, m2)
	t.Logf("Ex=%f, Dx=%f\n", ex2, dx2)
}

func BenchmarkRandDescreteInt(b *testing.B) {
	rf := randDiscreteInt([]float64{0.5, 0.25, 0.25}, []rndFn{N(0), N(1), N(3)})

	for i := 0; i < b.N; i++ {
		_ = rf(10)
	}
}

func TestRandDescreteInt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	rf := randDiscreteInt([]float64{0.5, 0.25, 0.25}, []rndFn{N(0), N(1), N(3)})

	m := make([]int, 10)
	for i := 0; i < 100000; i++ {
		r := rf(10)
		m[r]++
	}
	t.Log(m)
}

func TestRandAscendInt(t *testing.T) {
	N := 30
	s := make([]int, N)
	rf := randAscendInt(3, rand.Intn)
	for i := 0; i < N; i++ {
		s[i] = rf(0)
	}
	t.Log(s)
}

func N(n int) func(int) int {
	return func(int) int {
		return n
	}
}
