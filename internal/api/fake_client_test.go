package api

import (
	"reflect"
	"testing"
)

func TestFakeClient_GetUints(t *testing.T) {
	fakeClient := NewFakeClient(nil)
	got, err := fakeClient.GetUints("", 1)
	if err == nil {
		t.Fatalf(assertionMsgFmt, err, "non-nil")
	}

	if got != nil {
		t.Fatalf(assertionMsgFmt, got, nil)
	}

	expected := []uint16{12, 20, 30, 10}
	fakeClient = NewFakeClient(expected)
	got, err = fakeClient.GetUints("", 1)
	if err != nil {
		t.Fatalf(assertionMsgFmt, err, nil)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf(assertionMsgFmt, got, expected)
	}
}
