package ge_test

import (
	"github.com/CharLemAznable/ge"
	"testing"
)

type testFn func()

func TestRemoveElementByValue_Func(t *testing.T) {
	fn1 := func() {}
	fn2 := func() {}
	slice := []testFn{fn1, fn2, fn2, fn1}
	result := ge.RemoveElementByValue(slice, fn2)
	expected := []testFn{fn1, fn1}
	if !slicesEqualFunc(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	fn3 := func() {}
	result = ge.RemoveElementByValue(slice, fn3)
	expected = []testFn{fn1, fn2, fn2, fn1}
	if !slicesEqualFunc(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	result = ge.AppendElementUnique(slice, fn2)
	expected = []testFn{fn1, fn1, fn2}
	if !slicesEqualFunc(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func slicesEqualFunc(a, b []testFn) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !ge.EqualsPointer(v, b[i]) {
			return false
		}
	}
	return true
}
