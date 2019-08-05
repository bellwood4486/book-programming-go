package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	var argsF []float64
	if len(os.Args) >= 2 {
		argsF = convToFloat64(os.Args[1:])
	} else {
		var v float64
		fmt.Print("Enter number: ")
		_, _ = fmt.Scan(&v)
		argsF = append(argsF, v)
	}

	printTemps(argsF)
}

func convToFloat64(inputs []string) []float64 {
	fs := make([]float64, len(inputs))
	for i, in := range inputs {
		f, err := strconv.ParseFloat(in, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fs[i] = f
	}
	return fs
}

func printTemps(temps []float64) {
	for _, t := range temps {
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
