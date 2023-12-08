package ge_test

import (
	"github.com/CharLemAznable/ge"
	"testing"
)

func TestToBool(t *testing.T) {
	assertFalse(t, ge.ToBool(""))
	assertTrue(t, ge.ToBool("true"))
	assertTrue(t, ge.ToBool("TRUE"))
	assertTrue(t, ge.ToBool("tRUe"))
	assertTrue(t, ge.ToBool("on"))
	assertTrue(t, ge.ToBool("tRUe"))
	assertTrue(t, ge.ToBool("T"))
	assertFalse(t, ge.ToBool("false"))
	assertFalse(t, ge.ToBool("f"))
	assertFalse(t, ge.ToBool("No"))
	assertFalse(t, ge.ToBool("n"))
	assertTrue(t, ge.ToBool("on"))
	assertTrue(t, ge.ToBool("ON"))
	assertFalse(t, ge.ToBool("off"))
	assertFalse(t, ge.ToBool("oFf"))
	assertTrue(t, ge.ToBool("yes"))
	assertTrue(t, ge.ToBool("Y"))
	assertTrue(t, ge.ToBool("1"))
	assertFalse(t, ge.ToBool("0"))
	assertFalse(t, ge.ToBool("blue"))
	assertFalse(t, ge.ToBool("true "))
	assertFalse(t, ge.ToBool("ono"))
	assertFalse(t, ge.ToBool("oo"))
	assertFalse(t, ge.ToBool("o"))
	assertFalse(t, ge.ToBool("x gti"))
	assertFalse(t, ge.ToBool("x gti "))
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Error("Expected true, but got false")
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Error("Expected false, but got true")
	}
}
