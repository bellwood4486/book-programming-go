package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for fileName, lines := range counts {
		for line, n := range lines {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", fileName, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		if counts[name] == nil {
			counts[name] = make(map[string]int)
		}
		counts[name][input.Text()]++
	}
	// 注意: input.Err() からのエラーの可能性を無視している
}
