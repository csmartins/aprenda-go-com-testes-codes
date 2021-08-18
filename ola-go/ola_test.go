package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola()
	expected := "Hey world"

	if result != expected {
		t.Errorf("got'%s', but expected '%s'", result, expected)
	}
}
