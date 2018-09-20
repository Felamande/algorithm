package main

import (
	"fmt"
	"time"
)

func lg(n int64) (lg2n int) {
	cnt := 0
	if n == 0 {
		return 0
	}
	for {
		n = n >> 1
		if n == 0 {
			break
		}
		cnt++
	}
	return cnt
}

func main2() {
	nn := 0
	t := time.Now()
	for i := int64(1); i <= 1<<31; i++ {
		nn = lg2(i)
	}
	interval := time.Since(t).Nanoseconds()
	cnt := 1 << uint(nn)
	fmt.Printf("time=%vns, avg=%.3fns, n=%d", interval, float64(interval)/float64(cnt), nn)
}

func lg2(n int64) int {
	switch {
	case n < 2 && n >= 1:
		return 0
	case n < 4 && n >= 2:
		return 1
	case n < 8 && n >= 4:
		return 2
	case n < 16 && n >= 8:
		return 3
	case n < 32 && n >= 16:
		return 4
	case n < 64 && n >= 32:
		return 5
	case n < 128 && n >= 64:
		return 6
	case n < 256 && n >= 128:
		return 7
	case n < 512 && n >= 256:
		return 8
	case n < 1024 && n >= 512:
		return 9
	case n < 2048 && n >= 1024:
		return 10
	case n < 4096 && n >= 2048:
		return 11
	case n < 8192 && n >= 4096:
		return 12
	case n < 16384 && n >= 8192:
		return 13
	case n < 32768 && n >= 16384:
		return 14
	case n < 65536 && n >= 32768:
		return 15
	case n < 131072 && n >= 65536:
		return 16
	case n < 262144 && n >= 131072:
		return 17
	case n < 524288 && n >= 262144:
		return 18
	case n < 1048576 && n >= 524288:
		return 19
	case n < 2097152 && n >= 1048576:
		return 20
	case n < 4194304 && n >= 2097152:
		return 21
	case n < 8388608 && n >= 4194304:
		return 22
	case n < 16777216 && n >= 8388608:
		return 23
	case n < 33554432 && n >= 16777216:
		return 24
	case n < 67108864 && n >= 33554432:
		return 25
	case n < 134217728 && n >= 67108864:
		return 26
	case n < 268435456 && n >= 134217728:
		return 27
	case n < 536870912 && n >= 268435456:
		return 28
	case n < 1073741824 && n >= 536870912:
		return 29
	case n < 2147483648 && n >= 1073741824:
		return 30
	case n < 4294967296 && n >= 2147483648:
		return 31
	case n < 8589934592 && n >= 4294967296:
		return 32
	case n < 17179869184 && n >= 8589934592:
		return 33
	case n < 34359738368 && n >= 17179869184:
		return 34
	case n < 68719476736 && n >= 34359738368:
		return 35
	case n < 137438953472 && n >= 68719476736:
		return 36
	case n < 274877906944 && n >= 137438953472:
		return 37
	case n < 549755813888 && n >= 274877906944:
		return 38
	case n < 1099511627776 && n >= 549755813888:
		return 39
	case n < 2199023255552 && n >= 1099511627776:
		return 40
	case n < 4398046511104 && n >= 2199023255552:
		return 41
	case n < 8796093022208 && n >= 4398046511104:
		return 42
	case n < 17592186044416 && n >= 8796093022208:
		return 43
	case n < 35184372088832 && n >= 17592186044416:
		return 44
	case n < 70368744177664 && n >= 35184372088832:
		return 45
	case n < 140737488355328 && n >= 70368744177664:
		return 46
	case n < 281474976710656 && n >= 140737488355328:
		return 47
	case n < 562949953421312 && n >= 281474976710656:
		return 48
	case n < 1125899906842624 && n >= 562949953421312:
		return 49
	case n < 2251799813685248 && n >= 1125899906842624:
		return 50
	case n < 4503599627370496 && n >= 2251799813685248:
		return 51
	case n < 9007199254740992 && n >= 4503599627370496:
		return 52
	case n < 18014398509481984 && n >= 9007199254740992:
		return 53
	case n < 36028797018963968 && n >= 18014398509481984:
		return 54
	case n < 72057594037927936 && n >= 36028797018963968:
		return 55
	case n < 144115188075855872 && n >= 72057594037927936:
		return 56
	case n < 288230376151711744 && n >= 144115188075855872:
		return 57
	case n < 576460752303423488 && n >= 288230376151711744:
		return 58
	case n < 1152921504606846976 && n >= 576460752303423488:
		return 59
	case n < 2305843009213693952 && n >= 1152921504606846976:
		return 60
	case n < 4611686018427387904 && n >= 2305843009213693952:
		return 61
	case n <= 9223372036854775808-1 && n >= 4611686018427387904:
		return 62
	}

	return 0
}
