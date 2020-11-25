package api

import "fmt"

// NewFakeClient returns a new `Client` instance with a fake implementation that
// returns the `returns` array alongwith a nil error. If `returns` is nil, it
// returns a nil array alongwith a non-nil error.
func NewFakeClient(returns []uint16) Client {
	return &fakeClientImpl{returns}
}

type fakeClientImpl struct {
	returns []uint16
}

func (c *fakeClientImpl) GetUints(_ DataType, _ uint) ([]uint16, error) {
	if c.returns == nil {
		return nil, fmt.Errorf("fake-error")
	}

	return c.returns, nil
}
