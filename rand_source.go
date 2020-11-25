package qrng

import (
	"math/rand"
	"sync"

	"github.com/ashutoshgngwr/go-qrng/internal/api"
)

// Config defines the configuration options accepted by the `qrng` package.
type Config struct {
	// Panic if a remote request to get random data fails.
	PanicOnError bool

	// Query maximum amount of random data allowed from QRNG API in a single
	// request. It will result in significant performance gain if generating
	// random numbers very frequently.
	EnableBuffer bool
}

type sourceImpl struct {
	cfg    *Config
	client *api.Client
	buffer []uint16
	lock   sync.Mutex
}

// NewSource returns a new thread-safe `rand.Source` implementation that queries
// QRNG API to get true random data.
func NewSource(cfg *Config) rand.Source64 {
	return &sourceImpl{
		cfg:    cfg,
		client: api.New(),
		buffer: make([]uint16, 0),
		lock:   sync.Mutex{},
	}
}

func (src *sourceImpl) Int63() int64 {
	return int64(src.Uint64() >> 1)
}

func (src *sourceImpl) Uint64() uint64 {
	src.lock.Lock()
	defer src.lock.Unlock()

	if len(src.buffer) < 4 { // gonna need more items.
		n := uint(4)
		if src.cfg.EnableBuffer {
			n = api.MaxLength
		}

		var err error
		src.buffer, err = src.client.GetUints(api.DataTypeUint16, n)
		if err != nil {
			if src.cfg.PanicOnError {
				panic(err)
			}

			// if it shouldn't panic, return 0 since the buffer would be empty
			return 0

		}
	}

	var u64 uint64
	for _, u16 := range src.buffer[:4] {
		u64 = (u64 << 16) | uint64(u16)
	}

	src.buffer = src.buffer[4:]
	return u64
}

// Seed need not be implemented since the remote source isn't a pseudo random
// number generator and doesn't need to be seeded.
func (src *sourceImpl) Seed(seed int64) {}
