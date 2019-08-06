package popcount

import (
	"strconv"
	"testing"
)

var benchSizes = []int{10, 100, 1000}

func BenchmarkPopCount(b *testing.B) {
	bench(b, PopCount)
}

func BenchmarkPopCountN(b *testing.B) {
	subBenchSizes(b, benchSizes, PopCount)
}

func BenchmarkPopCountTable(b *testing.B) {
	bench(b, PopCountTable)
}

func BenchmarkPopCountTableN(b *testing.B) {
	subBenchSizes(b, benchSizes, PopCountTable)
}

func BenchmarkPopCountBit(b *testing.B) {
	bench(b, PopCountBit)
}

func BenchmarkPopCountBitN(b *testing.B) {
	subBenchSizes(b, benchSizes, PopCountBit)
}
func BenchmarkPopCountClear(b *testing.B) {
	bench(b, PopCountClear)
}

func BenchmarkPopCountClearN(b *testing.B) {
	subBenchSizes(b, benchSizes, PopCountClear)
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func benchN(b *testing.B, n int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func subBenchSizes(b *testing.B, sizes []int, f func(uint64) int) {
	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			benchN(b, size, f)
		})
	}
}
