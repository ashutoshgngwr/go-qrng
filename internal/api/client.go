package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// defaultBaseURL where the API apiEndpoint can be reached.
	defaultBaseURL = "https://qrng.anu.edu.au"

	// apiEndpointFmt declares the format string to fetch uints from the QRNG
	// API endpoint. The following values need to be specified in-order: base
	// URL, data type and output length.
	apiEndpointFmt = "%s/API/jsonI.php?type=%s&length=%d"

	// MaxLength declares the maximum number of uints that can queried in a
	// single request.
	MaxLength = 1024

	// DataTypeUint8 declares the `DataType` literal for querying 8-bit uints from the
	// API endpoint.
	DataTypeUint8 DataType = "uint8"

	// DataTypeUint16 declares the `DataType` literal for querying 16-bit uints from
	// the API endpoint.
	DataTypeUint16 DataType = "uint16"
)

// DataType declares a named type for declaring type literals accepted by the
// API Endpoint.
type DataType string

// Client declares functions to interact with the QRNG API.
type Client interface {
	// GetUints queries `n` integers of requested data type `t` from the QRNG API.
	// `n` should not be more than `MaxLength`.
	GetUints(t DataType, n uint) ([]uint16, error)
}

// NewClient returns a new instance of `Client`.
func NewClient() Client {
	return &clientImpl{http.DefaultClient, defaultBaseURL}
}

type clientImpl struct {
	*http.Client
	baseURL string
}

// qrngResponse declares the JSON schema of the response returned by QRNG API.
type qrngResponse struct {
	Success bool
	Type    DataType
	Length  uint
	Data    []uint16
}

func (c *clientImpl) GetUints(t DataType, n uint) ([]uint16, error) {
	if t != DataTypeUint8 && t != DataTypeUint16 {
		return nil, fmt.Errorf("t should be one of %q or %q", DataTypeUint8, DataTypeUint16)
	}

	if n > MaxLength { // since both have equal maxLen
		return nil, fmt.Errorf("n should be <= %d, given: %d", MaxLength, n)
	}

	resp, err := c.Get(fmt.Sprintf(apiEndpointFmt, c.baseURL, t, n))
	if err != nil {
		return nil, fmt.Errorf("remote request failed: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %s", err)
	}

	qrngResp := qrngResponse{}
	if err = json.Unmarshal(body, &qrngResp); err != nil {
		return nil, fmt.Errorf("error while decoding JSON response: %s", err)
	}

	if !qrngResp.Success {
		return nil, fmt.Errorf("QRNG API returned an unsuccessful response")
	}

	if qrngResp.Length != n {
		return nil, fmt.Errorf("requested length didn't match the returned response length")
	}

	return qrngResp.Data, nil
}
