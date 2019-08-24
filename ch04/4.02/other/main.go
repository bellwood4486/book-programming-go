package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func sha2Sum(kind string) (func([]byte) []byte, error) {
	switch kind {
	case "256":
		return func(in []byte) []byte {
			sum := sha256.Sum256(in)
			return sum[:]
		}, nil
	case "384":
		return func(in []byte) []byte {
			sum := sha512.Sum384(in)
			return sum[:]
		}, nil
	case "512":
		return func(in []byte) []byte {
			sum := sha512.Sum512(in)
			return sum[:]
		}, nil
	default:
		return nil, errors.Errorf("unknown kind: %s", kind)
	}
}

func main() {
	kind := flag.String("k", "256",
		fmt.Sprintf("you can select: %s, %s, %s", "256", "384", "512"))
	flag.Parse()

	fmt.Print("Input data: ")
	var in string
	_, _ = fmt.Scanln(&in)

	sumFunc, err := sha2Sum(*kind)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "sha2 sum failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", sumFunc([]byte(in)))
}
