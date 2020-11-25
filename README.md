# go-qrng

![Go](https://img.shields.io/github/go-mod/go-version/ashutoshgngwr/go-qrng)
[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://pkg.go.dev/github.com/ashutoshgngwr/go-qrng)
[![Go](https://github.com/ashutoshgngwr/go-qrng/workflows/Go/badge.svg)](https://github.com/ashutoshgngwr/go-qrng/actions?query=workflow%3AGo)
[![codecov](https://codecov.io/gh/ashutoshgngwr/go-qrng/branch/master/graph/badge.svg?token=F6BRVRHY8M)](https://codecov.io/gh/ashutoshgngwr/go-qrng)

go-qrng is an extension for [`math/rand`](https://golang.org/pkg/math/rand/)
package to use [Australian National University](https://www.anu.edu.au/)'s
[Quantum Random Number Generator](https://qrng.anu.edu.au/) with the std
[`rand.Rand`](https://golang.org/pkg/math/rand/#Rand) API. ANU QRNG API provides
true random data generated in real-time in a lab by measuring the quantum
fluctuations of the vacuum.

## Example

```go
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
  // it requests for new data from the remote API to refill the buffer. Note:
  // A singe remote API request fetches up to 1024 uint16s.
  for i := 0; i < 128; i++ {
    fmt.Println(r.Int(), "\t", r.Uint32(), "\t", r.Float32(), "\t", r.Float64())
  }
}
```

## License

[Apache License Version 2.0](LICENSE)
