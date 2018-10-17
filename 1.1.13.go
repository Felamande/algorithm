package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main9() {
	m1 := makeMatrix(2, 3)
	rp := newRandProvider(3)
	m1.fill(rp)
	fmt.Println(m1)
	m2 := makeMatrix(3, 1)
	m2.fill(rp)
	fmt.Println("----------------------")
	fmt.Println(m2)
	fmt.Println("----------------------")
	fmt.Println(m1.Mult(m2))
	fmt.Println("----------------------")
	fmt.Println(m1.transpose())

}

type matrix2 struct {
	m   [][]int
	row int
	col int
}

func (m *matrix2) Mult(m1 *matrix2) (m2 *matrix2, err error) {
	if m1 == nil {
		return nil, errors.New("input matrix is nil")
	}
	if m.col != m1.row {
		return nil, fmt.Errorf("matrix size not match:(%d,%d)*(%d,%d)", m.row, m.col, m1.row, m1.col)
	}

	m2 = makeMatrix(m.row, m1.col)

	for row := 0; row < m.row; row++ {
		for col1 := 0; col1 < m1.col; col1++ {
			re := 0
			for x := 0; x < m.col; x++ {
				re += m.m[row][x] * m1.m[x][col1]
			}
			m2.m[row][col1] = re
		}
	}
	return
}

func (m *matrix2) transpose() (m1 *matrix2) {
	m1 = makeMatrix(m.col, m.row)
	for row := 0; row < m.row; row++ {
		for col := 0; col < row; col++ {
			m1.m[row][col], m1.m[col][row] = m.m[col][row], m.m[row][col]
		}
	}
	return
}

func (m *matrix2) size() (row, col int) {
	return m.row, m.col
}

func (m *matrix2) fill(p elemProvider) {
	for row := 0; row < m.row; row++ {
		for col := 0; col < m.col; col++ {
			m.m[row][col] = p.Provide(row, col)
		}
	}
}

func (m *matrix2) String() string {
	str := ""
	for row := 0; row < m.row; row++ {
		for col := 0; col < m.col; col++ {
			str += strconv.Itoa(m.m[row][col]) + "\t" + ""
		}
		str += "\n"
	}
	return strings.Trim(str, "\t\n")
}

func makeMatrix(row, col int) *matrix2 {
	m := new(matrix2)
	m.row = row
	m.col = col

	dimm := 0
	if row >= col {
		dimm = row
	} else {
		dimm = col
	}

	for i := 0; i < dimm; i++ {
		m.m = append(m.m, make([]int, dimm))
	}
	return m
}

type elemProvider interface {
	Provide(row, col int) int
}

type randProvider struct {
	max int64
}

func newRandProvider(max int64) randProvider {
	return randProvider{max: max}
}

func (rp randProvider) Provide(row, col int) int {
	i, _ := rand.Int(rand.Reader, big.NewInt(rp.max))
	return int(i.Int64())
}

func zeros(int, int) int {
	return 0
}

func ones(int, int) int {
	return 1
}

func all(i int) func(int, int) int {
	return func(int, int) int {
		return i
	}
}

type funcProvider func(row, col int) int

func (pf funcProvider) Provide(row, col int) int {
	return pf(row, col)
}
