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