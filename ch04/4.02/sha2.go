package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

const (
	KindSHA256 = "256"
	KindSHA384 = "384"
	KindSHA512 = "512"
)

var (
	kind = flag.String("k", KindSHA256, fmt.Sprintf("you can select: %s, %s, %s", KindSHA256, KindSHA384, KindSHA512))
)

func sha2String(in []byte, kind string) (string, error) {
	hash, err := sha2(in, kind)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash), nil
}

func sha2(in []byte, kind string) ([]byte, error) {
	var hash []byte
	switch kind {
	case KindSHA256:
		a := sha256.Sum256(in)
		hash = a[:]
	case KindSHA384:
		a := sha512.Sum384(in)
		hash = a[:]
	case KindSHA512:
		a := sha512.Sum512(in)
		hash = a[:]
	default:
		return nil, errors.Errorf("unknown kind: %s", kind)
	}

	return hash, nil
}

func main() {
	flag.Parse()

	fmt.Print("Input data: ")
	var in string
	_, _ = fmt.Scanln(&in)

	hash, err := sha2String([]byte(in), *kind)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "sha2 sum failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(hash)
}
