package assert

import "testing"

func Equals(expected interface{}, result interface{}, t *testing.T) {
	if expected != result {
		t.Errorf("Error: expected value %s but get %s", expected, result)
	}
}

func NotEquals(expected interface{}, result interface{}, t *testing.T) {
	if expected == result {
		t.Errorf("Error: expected value %s but get %s", expected, result)
	}
}

func NotNil(result interface{}, t *testing.T) {
	if result == nil {
		t.Errorf("Error: expected value is not nil but get %v.", result)
	}
}

func Nil(result interface{}, t *testing.T) {
	if result != nil {
		t.Errorf("Error: expected value is nil but get %v.", result)
	}
}
