package inteiros

import "testing"

func TestAdd(t *testing.T){
	result := Add(2,2)
	expected := 4

	if expected != result{
		t.Errorf("got '%d', but expected '%d'", result, expected)
	}
}