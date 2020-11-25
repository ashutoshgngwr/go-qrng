package main

import (
	"fmt"
	"math/rand"

	"github.com/ashutoshgngwr/go-qrng"
)

func main() {
	// Create a new `rand.Source` instance with QRNG implementation.
	s := qrng.NewSource(&qrng.Config{PanicOnError: true, EnableBuffer: true})

	// Create a new `rand.Rand` instance
	r := rand.New(s)

	// The following will trigger only two remote requests to the QRNG API since
	// buffering is enabled. Buffer-enabled `Source` fetches max allowed data in
	// a single remote request. It keeps serving future generate-number requests
	// from the local buffer until it is exhausted. Once local buffer is empty,
	// it requests for new data from the remote API to refill the buffer.
	for i := 0; i < 128; i++ {
		fmt.Println(r.Int(), "\t", r.Uint32(), "\t", r.Float32(), "\t", r.Float64())
	}
}
