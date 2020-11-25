package qrng

import (
	"testing"

	"github.com/ashutoshgngwr/go-qrng/internal/api"
)

const assertionMsgFmt = "expected '%v' to be '%v'\n"

func TestSource_Uint64(t *testing.T) {
	fakeAPIClient := api.NewFakeClient([]uint16{1, 0, 1, 0})
	s := NewSourceWithClient(&Config{}, fakeAPIClient)
	if got, expected := s.Uint64(), uint64(281474976776192); got != expected {
		t.Fatalf(assertionMsgFmt, got, expected)
	}
}

func TestSource_Int63(t *testing.T) {
	fakeAPIClient := api.NewFakeClient([]uint16{1, 0, 1, 0})
	s := NewSourceWithClient(&Config{}, fakeAPIClient)
	if got, expected := s.Int63(), int64(281474976776192>>1); got != expected {
		t.Fatalf(assertionMsgFmt, got, expected)
	}
}
