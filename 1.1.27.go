package main

func binomial(N int, k int, p float64, cnt *int) float64 {
	*cnt = *cnt + 1
	if N == 0 && k == 0 {
		return 1.0
	}
	if N < 0 || k < 0 {
		return 0.0
	}

	return (1.0-p)*binomial(N-1, k, p, cnt) + p*binomial(N-1, k-1, p, cnt)
}

func main3() {
	var cnt int
	println(binomial(9, 10, 2, &cnt))

	println(cnt)
}
