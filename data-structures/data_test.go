package dataStructures

import "testing"

func TestPerimeter(t *testing.T)  {
	
	verifyCorrectness := func(t *testing.T, result, expected float64){
		t.Helper()
		if result != expected {
			t.Errorf("got '%.2f', but expected '%.2f'", result, expected)
		}
	}
	
	t.Run("perimeter", func(t *testing.T) {
		result := Perimeter(Rectangle{10.0, 10.0})
		expected := 40.0

		verifyCorrectness(t, result, expected)
	})
}

func TestArea(t *testing.T) {
	verifyCorrectness := func(t *testing.T, result, expected float64){
		t.Helper()
		if result != expected {
			t.Errorf("got '%.2f', but expected '%.2f'", result, expected)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		result := rectangle.Area()
		expected := 72.0

		verifyCorrectness(t, result, expected)
	})
	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10.0}
		result := circle.Area()
		expected := 314.1592653589793

		verifyCorrectness(t, result, expected)
	})
	
}

func TestSearch(t *testing.T){
	dict := Dict{"test": "this is just a test"}

	t.Run("known entry", func (t *testing.T) {
		result, _ := dict.Search("test")
		expected := "this is just a test"
	
		compareStrings(result, expected, t)
	})

	t.Run("unknown entry", func (t *testing.T) {
		_, err := dict.Search("unknown")
	
		if err == nil{
			t.Fatal("Got error as expected")
		}
	})
}

func compareStrings(result string, expected string, t *testing.T) {
	t.Helper()
	if result != expected {
		t.Errorf("got '%s', but expected '%s'", result, expected)
	}
}