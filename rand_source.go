package qrng

import (
	"math/rand"

	"github.com/ashutoshgngwr/go-qrng/internal/api"
)

// Config defines the configuration options accepted by the `qrng` package.
type Config struct {
	// panic if a remote request to get random data fails.
	PanicOnError bool
}

// NewSource returns a new `rand.Source` implementation that queries QRNG API to
// get true random data.
func NewSource(cfg *Config) rand.Source {
	return &sourceImpl{cfg, api.New()}
}

type sourceImpl struct {
	cfg    *Config
	client *api.Client
}

// type cheking on `rand.Source64` since it is a superset of `rand.Source`.
var _ rand.Source64 = &sourceImpl{}

func (src *sourceImpl) Int63() int64 {
	return int64(src.Uint64() >> 1)
}

func (src *sourceImpl) Uint64() uint64 {
	u16s, err := src.client.GetUints(api.DataTypeUint16, 4)
	if err != nil && src.cfg.PanicOnError {
		panic(err)
	}

	var u64 uint64
	for _, u16 := range u16s {
		u64 = (u64 << 16) | uint64(u16)
	}

	return u64
}

// Seed need not be implemented since the remote source isn't a pseudo random
// number generator and doesn't need to be seeded.
func (src *sourceImpl) Seed(seed int64) {}
