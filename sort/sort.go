package main

import (
	"fmt"
	"math"
	"math/rand"
	"path"
	"reflect"
	"runtime"
	"sort"
	"time"
)

type sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Count() (swpCnt, cmpCnt int)
	Clone() sortable
	At(i int) int
}

type sortFn func(s sortable)

func sorted(s sortable) bool {
	for i := 0; i < s.Len()-1; i++ {
		if s.Less(i+1, i) {
			return false
		}
	}
	return true
}
func sysSort(s sortable) {
	sort.Sort(s)
}

func selSort(s sortable) {
	for i := 0; i < s.Len(); i++ {
		min := i
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(j, min) {
				min = j
			}
		}
		s.Swap(i, min)
	}
}

//worst comp:O(n^2) swap:O(n^2)
//best  comp:O(n^2) swap:0
func bubbleSort(s sortable) {
	for i := 0; i < s.Len(); i++ {
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(j, i) {
				s.Swap(i, j)
			}
		}
	}
}

func insertSort(s sortable) {
	for i := 1; i < s.Len(); i++ {
		for j := i; j > 0 && s.Less(j, j-1); j-- {
			s.Swap(j-1, j)
		}
	}
}

func shSortd2(s sortable) {
	N := s.Len()
	for _, h := range arrd2 {
		for i := h; i < N; i++ {
			for j := i; j >= h && s.Less(j, j-h); j = j - h {
				s.Swap(j, j-h)
			}
		}
	}
}

func shSortAr1(s sortable) {
	N := s.Len()
	for _, h := range arrk {
		for i := h; i < N; i++ {
			for j := i; j >= h && s.Less(j, j-h); j = j - h {
				s.Swap(j, j-h)
			}
		}
	}
}

func shSort255(s sortable) {
	N := s.Len()
	for _, h := range arr225 {
		for i := h; i < N; i++ {
			for j := i; j >= h && s.Less(j, j-h); j = j - h {
				s.Swap(j, j-h)
			}
		}
	}
}

func shSort248(s sortable) {
	N := s.Len()
	for _, h := range arr2p48k {
		for i := h; i < N; i++ {
			for j := i; j >= h && s.Less(j, j-h); j = j - h {
				s.Swap(j, j-h)
			}
		}
	}
}

func shSortt(seq []int, s sortable) {

	N := s.Len()
	for _, h := range seq {
		for i := h; i < N; i++ {
			for j := i; j >= h && s.Less(j, j-h); j = j - h {
				s.Swap(j, j-h)
			}
		}
	}
}

func quickSort(s sortable) {

}

func main() {
	rand.Seed(time.Now().UnixNano())
	for N := 80000; N < 80000*1<<7; N *= 2 {
		ss := makeInts(func(int) int { return randNormInt(1000) }, N)

		// ss = makeInts(func(i int) int { return i }, N)
		// ss = makeInts(func(i int) int { return 1 }, N)
		// ss = makeInts(func(i int) int { return N - i }, N)
		// ss = makeInts2(1, 2, 4, 3, 5, 6, 8, 7)

		for _, f := range []sortFn{sysSort, shSortd2, shSortAr1, shSort255, shSort248} {
			s := ss.Clone()
			dt := benchmark(func() { f(s) }, 1)
			swp, cmp := s.Count()
			fmt.Printf("%s\tN=%d\tswp=%9d\tcmp=%10d\tt=%v\n", getName(f), N, swp, cmp, dt)
		}
		fmt.Println("----------------------------------------------------------------------")
	}

}

func main__() {
	rand.Seed(time.Now().UnixNano())
	for N := 20000; N < 20000*1<<3; N *= 2 {
		ss := makeInts(func(int) int { return rand.Intn(1000) }, N)
		// ss = makeInts(func(i int) int { return i }, N)
		// ss = makeInts(func(i int) int { return 1 }, N)
		// ss = makeInts(func(i int) int { return N - i }, N)
		// ss = makeInts2(1, 2, 4, 3, 5, 6, 8, 7)
		for t := float64(2.1); t <= 10.2; t += 1 {
			s := ss.Clone()
			seq := genSeq(seqt(t))
			dt := benchmark(func() { shSortt(seq, s) }, 1)
			swp, cmp := s.Count()
			fmt.Printf("t=%.3f\tN=%d\tswp=%9d\tcmp=%10d\tt=%v\n", t, N, swp, cmp, dt)
		}
		fmt.Println("----------------------------------------------------------------------")
	}
}

type Ints struct {
	ints   []int
	swpCnt int
	cmpCnt int
}

func (p *Ints) Len() int                    { return len(p.ints) }
func (p *Ints) Less(i, j int) bool          { p.cmpCnt++; return p.ints[i] < p.ints[j] }
func (p *Ints) Swap(i, j int)               { p.swpCnt++; p.ints[i], p.ints[j] = p.ints[j], p.ints[i] }
func (p *Ints) Count() (swpCnt, cmpCnt int) { return p.swpCnt, p.cmpCnt }
func (p *Ints) Clone() sortable {
	return makeInts(func(i int) int { return p.ints[i] }, len(p.ints))
}
func (p *Ints) At(i int) int { return p.ints[i] }

func makeInts(f func(idx int) int, N int) (a *Ints) {
	a = new(Ints)
	for i := 0; i < N; i++ {
		a.ints = append(a.ints, f(i))
	}
	return a
}
func makeInts2(ints ...int) *Ints {
	a := new(Ints)
	for i := 0; i < len(ints); i++ {
		a.ints = append(a.ints, ints[i])
	}
	return a
}
func benchmark(f func(), times int) time.Duration {
	var t time.Time
	var dt time.Duration

	t = time.Now()
	for i := 0; i < times; i++ {
		f()
	}
	dt = time.Since(t)

	return dt
}

func getName(f interface{}) string {
	full := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return path.Base(full)
}

var arrd2 = []int{1073741823, 536870911, 268435455, 134217727, 67108863, 33554431, 16777215,
	8388607, 4194303, 2097151, 1048575, 524287, 262143, 131071, 32767, 65535, 16383, 8191, 4095,
	2047, 1023, 511, 255, 127, 63, 31, 15, 7, 3, 1}

var arrk = []int{2415771649, 1073643521, 603906049, 268386305, 150958081, 67084289, 37730305,
	16764929, 9427969, 4188161, 2354689, 1045505, 587521, 260609, 146305, 64769, 36289, 16001,
	8929, 3905, 2161, 929, 505, 209, 109, 41, 19, 5, 1}

var arr225 = []int{1698453753, 754868335, 335497038, 149109795, 66271020, 29453787, 13090572,
	5818032, 1149241, 2585792, 510774, 227011, 100894, 44842, 19930, 8858, 3937, 1750, 701,
	301, 132, 57, 23, 10, 4, 1}

var arr2p48k = []int{2929815956, 1181377402, 476361856, 192081394, 77452175, 31230716, 12593031,
	5077835, 2047515, 825611, 332908, 134237, 54128, 21826, 8801, 3549, 1431, 577, 233, 94, 38, 16, 7, 3, 1}

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

func sortCompare() {}
