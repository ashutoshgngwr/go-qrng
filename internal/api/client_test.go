package api

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const assertionMsgFmt = "expected '%v' to be '%v'\n"

func TestClient_GetUints(t *testing.T) {
	testCases := []struct {
		dataType DataType
		nums     uint
		httpResp string
		returns  []uint16
		errors   bool
	}{
		{DataTypeUint16, 4, `{"type":"uint16","length":4,"data":[0,1,0,1],"success":true}`, []uint16{0, 1, 0, 1}, false},
		{DataTypeUint16, 4, ``, nil, true},
		{DataTypeUint16, 4, `{"success":false}`, nil, true},
		{DataTypeUint16, MaxLength + 1, ``, nil, true},
		{"unsupported", 4, ``, nil, true},
	}

	for _, testCase := range testCases {
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte(testCase.httpResp)); err != nil {
				t.Errorf("unable to write response from mock server: %s", err)
			}
		}

		mockServer := httptest.NewServer(handler)
		client := clientImpl{mockServer.Client(), mockServer.URL}
		nums, err := client.GetUints(testCase.dataType, testCase.nums)
		if testCase.errors && err == nil {
			t.Fatalf(assertionMsgFmt, err, "non-nil")
		}

		if !testCase.errors && err != nil {
			t.Fatalf(assertionMsgFmt, err, nil)
		}

		if !reflect.DeepEqual(testCase.returns, nums) {
			t.Fatalf(assertionMsgFmt, nums, testCase.returns)
		}

		mockServer.Close()
	}
}
