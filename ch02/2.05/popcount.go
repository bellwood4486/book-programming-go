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

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	c := 0
	var i uint
	for i = 0; i < 8; i++ {
		c += int(pc[byte(x>>(i*8))])
	}
	return c
}

func PopCountBit(x uint64) int {
	c := 0
	for ; x > 0; x >>= 1 {
		c += int(x & 0x01)
	}
	return c
}

func PopCountClear(x uint64) int {
	c := 0
	for ; x > 0; x = x & (x - 1) {
		c++
	}
	return c
}

func main() {
	var x uint64 = 10000
	var start time.Time

	start = time.Now()
	fmt.Printf("pop count: %d\n", PopCount(x))
	fmt.Printf("PopCount: %v elapsed\n", time.Since(start))

	start = time.Now()
	fmt.Printf("pop count: %d\n", PopCountLoop(x))
	fmt.Printf("PopCountLoop: %v elapsed\n", time.Since(start))

	start = time.Now()
	fmt.Printf("pop count: %d\n", PopCountBit(x))
	fmt.Printf("PopCountBit: %v elapsed\n", time.Since(start))

	start = time.Now()
	fmt.Printf("pop count: %d\n", PopCountClear(x))
	fmt.Printf("PopCountClear: %v elapsed\n", time.Since(start))
}
