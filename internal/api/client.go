package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// APIEndpointFmt declares the format string to fetch 16-bit uints from the
	// QRNG API endpoint.
	APIEndpointFmt = "https://qrng.anu.edu.au/API/jsonI.php?type=%s&length=%d"

	// MaxLength declares the maximum number of uints that can queried in a
	// single request.
	MaxLength = 1024

	// DataTypeUint8 declares the type literal for querying 8-bit uints from the
	// API endpoint.
	DataTypeUint8 DataType = "uint8"

	// DataTypeUint16 declares the type literal for querying 16-bit uints from
	// the API endpoint.
	DataTypeUint16 DataType = "uint16"
)

// DataType declares a named type for declaring type literals accepted by the
// API Endpoint.
type DataType string

// QRNGResponse declares the JSON schema of the response returned by QRNG API
// endpoint.
type QRNGResponse struct {
	Success bool
	Type    DataType
	Length  uint
	Data    []uint16
}

// Client implements functions to interact with the QRNG API.
type Client struct {
	*http.Client
}

// New returns a new instance of `QRNGClient`.
func New() *Client {
	return &Client{http.DefaultClient}
}

// GetUints queries `n` integers of requested data type `t` from the QRNG API.
// `n` should not be more than `MaxLength`.
func (client *Client) GetUints(t DataType, n uint) ([]uint16, error) {
	if t != DataTypeUint8 && t != DataTypeUint16 {
		return nil, fmt.Errorf("t should be one of %q or %q", DataTypeUint8, DataTypeUint16)
	}

	if n > MaxLength { // since both have equal maxLen
		return nil, fmt.Errorf("n should be < %d, given: %d", MaxLength, n)
	}

	resp, err := client.Get(fmt.Sprintf(APIEndpointFmt, DataTypeUint16, n))
	if err != nil {
		return nil, fmt.Errorf("remote request failed: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %s", err)
	}

	qrngResp := QRNGResponse{}
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
