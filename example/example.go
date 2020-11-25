package main

import (
	"fmt"
	"math/rand"

	"github.com/ashutoshgngwr/go-qrng"
)

func main() {
	// create a new `rand.Source` instance with QRNG implementation.
	s := qrng.NewSource(&qrng.Config{PanicOnError: true})

	// create a new `rand.Rand` instance
	r := rand.New(s)

	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}

	for i := 0; i < 10; i++ {
		fmt.Println(r.Float64())
	}
}
