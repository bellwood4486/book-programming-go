package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	c := 0
	var i uint
	for i = 0; i < 8; i++ {
		c += int(pc[byte(x>>(i*8))])
	}
	return c
}

func main() {
	var x uint64 = 10000

	start := time.Now()
	fmt.Printf("pop count: %d\n", PopCount1(x))
	fmt.Printf("PopCount1: %v elapsed\n", time.Since(start))

	start = time.Now()
	fmt.Printf("pop count: %d\n", PopCount2(x))
	fmt.Printf("PopCount2: %v elapsed\n", time.Since(start))
}
