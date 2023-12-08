package ge_test

import (
	"errors"
	"github.com/CharLemAznable/ge"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	finished := make(chan error)
	panicked := make(ge.Panicked)

	go func() {
		defer panicked.Recover()
		finished <- errors.New("error")
	}()

	var actualError error
	select {
	case err := <-finished:
		actualError = err
	case v := <-panicked.Caught():
		actualError = ge.WrapPanic(v)
	}
	if actualError.Error() != "error" {
		t.Errorf("Expected error message 'error', but got '%s'", actualError.Error())
	}

	go func() {
		defer panicked.Recover()
		panic("panicked")
	}()

	select {
	case err := <-finished:
		actualError = err
	case v := <-panicked.Caught():
		actualError = ge.WrapPanic(v)
	}
	panicError, ok := (actualError).(*ge.PanicError)
	if !ok {
		t.Errorf("Expected error is common.PanicError, but got %T", actualError)
	}
	if panicError.Error() != "panicked with panicked" {
		t.Errorf("Expected error message 'panicked with panicked', but got '%s'", panicError.Error())
	}
}
