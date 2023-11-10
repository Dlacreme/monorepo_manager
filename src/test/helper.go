package test

import "testing"

func AssertStringEq(t *testing.T, got string, expect string) {
	t.Helper()
	if expect != got {
		t.Errorf("Expect: [%s], Got: [%s]\n", expect, got)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error [%s]\n", err)
	}
}

func AssertErrorEq(t *testing.T, err error, expect string) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error but got none")
	}
	errStr := err.Error()
	if errStr != expect {
		t.Errorf("Expected error [%s] but got [%s]", expect, errStr)
	}
}

func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func AssertNumEq(t *testing.T, got int, expect int) {
	t.Helper()
	if expect != got {
		t.Errorf("Expect: [%d], Got: [%d]\n", expect, got)
	}
}
