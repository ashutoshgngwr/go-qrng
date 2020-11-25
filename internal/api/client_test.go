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
	}{
		{DataTypeUint16, 4, `{"type":"uint16","length":4,"data":[0,1,0,1],"success":true}`, []uint16{0, 1, 0, 1}},
		{DataTypeUint16, 4, ``, nil},
		{DataTypeUint16, 4, `{"success":false}`, nil},
		{DataTypeUint16, MaxLength + 1, ``, nil},
		{"unsupported", 4, ``, nil},
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
		if testCase.returns == nil && err == nil {
			t.Fatalf(assertionMsgFmt, err, "non-nil")
		}

		if testCase.returns != nil && err != nil {
			t.Fatalf(assertionMsgFmt, err, nil)
		}

		if !reflect.DeepEqual(testCase.returns, nums) {
			t.Fatalf(assertionMsgFmt, nums, testCase.returns)
		}

		mockServer.Close()
	}
}
