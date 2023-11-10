package test

import "testing"

func AssertStringEq(t *testing.T, expect string, got string) {
	if expect != got {
		t.Errorf("Expect: [%s], Got: [%s]\n", expect, got)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error [%s]\n", err)
	}
}

func AssertErrorEq(t *testing.T, err error, expect string) {
	if err == nil {
		t.Errorf("Expected error but got none")
	}
	errStr := err.Error()
	if errStr != expect {
		t.Errorf("Expected error [%s] but got [%s]", expect, errStr)
	}
}
