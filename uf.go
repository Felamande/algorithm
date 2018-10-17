package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type UF struct {
	n   int
	cnt int
	id  []int
	sz  []int
}

func NewUF(N int) *UF {
	uf := &UF{
		n:   N,
		cnt: N,
		id:  make([]int, N),
		sz:  make([]int, N),
	}
	for i := 0; i < N; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return uf
}

func (uf *UF) PrintTree(w io.Writer) {
	depth := 0
	for i := 0; i < uf.n; i++ {
		if uf.id[i] == i {
			// fmt.Fprintf(w, "%d\t", i)
			uf.printChild(w, i, &depth)
		}
	}
}

func (uf *UF) isLeaf(n int) bool {
	return uf.sz[n] == 1
}

func (uf *UF) isFather(f, c int) bool {
	return uf.id[c] == f
}

func (uf *UF) isRoot(n int) bool {
	return uf.id[n] == n
}

func (uf *UF) printChild(w io.Writer, n int, depth *int) {
	d := *depth
	d = d
	for i := 0; i < *depth; i++ {
		fmt.Fprintf(w, "\t")
	}
	fmt.Fprintf(w, "%d", n)

	*depth++
	defer func() { *depth-- }()

	if uf.isLeaf(n) {
		fmt.Fprintf(w, "\n")
		return
	}

	for i := 0; i < uf.n; i++ {
		if uf.isFather(n, i) && !uf.isRoot(i) {
			uf.printChild(w, i, depth)
		}
	}
	fmt.Fprintf(w, "\n")

}
func (uf *UF) union(p, q int) {
	if p > uf.n-1 || q > uf.n-1 {
		return
	}
	idp := uf.find(p)
	idq := uf.find(q)

	if idp == idq {
		return
	}
	if uf.sz[idp] < uf.sz[idq] {
		uf.id[idp] = idq
		uf.sz[idq] += uf.sz[idp]
	} else {
		uf.id[idq] = idp
		uf.sz[idp] += uf.sz[idq]
	}
	// fmt.Printf("pair{%d,%d}:true,", p, q)
	uf.cnt--
}

func (uf *UF) find(p int) int {
	if p > uf.n-1 {
		return -1
	}
	root := p
	for uf.id[root] != root {
		root = uf.id[root]
	}

	// uf.id[p] = root

	return root
}

func (uf *UF) connected(p, q int) bool {
	if p > uf.n-1 || q > uf.n-1 {
		return false
	}
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	uf.id[p] = pRoot
	uf.id[q] = qRoot
	return pRoot == qRoot

}

func (uf *UF) count() int {
	return uf.cnt
}

func (uf *UF) String() string {
	s := ""
	for i := 0; i < uf.n; i++ {
		s += fmt.Sprintf("%d\t", i)
	}
	s += "\n"
	for i := 0; i < uf.n; i++ {
		s += fmt.Sprintf("%d\t", uf.id[i])
	}
	return s
}

func main000() {
	N := 320000
	uf := NewUF(N)
	var dt time.Duration
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1; i++ {
		p := rand.Intn(N)
		q := rand.Intn(N)
		dt += benchmark(func() { uf.union(p, q) }, 1000000)
	}
	// fmt.Println(uf)
	// uf.PrintTree(os.Stdout)
	fmt.Printf("t=%v", dt)
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
