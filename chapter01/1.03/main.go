package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 非効率
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("inefficiency:\t%v elapsed\n", time.Since(start))

	// strings.Join
	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("strings.Join:\t%v elapsed\n", time.Since(start))
}
