package main

import "testing"

func TestOla(t *testing.T) {

	verifyCorrectness := func(t *testing.T, result, expected string){
		t.Helper()
		if result != expected {
			t.Errorf("got '%s', but expected '%s'", result, expected)
		}
	}
	t.Run("say hey to the world", func(t *testing.T){
		result := Ola("world")
		expected := "Hey world"
	
		verifyCorrectness(t, result, expected)
	})

	t.Run("default hey", func(t *testing.T){
		result := Ola("")
		expected := "Hey you"
		
		verifyCorrectness(t, result, expected)
	})
}