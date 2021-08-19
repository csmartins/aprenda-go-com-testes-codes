package sequences

import "testing"

func TestArrays(t *testing.T){

	verifyCorrectness := func(t *testing.T, result, expected int, caseData []int){
		t.Helper()
		if result != expected {
			t.Errorf("got '%d', but expected '%d'", result, expected)
		}
	}
	t.Run("sum array numbers", func(t *testing.T){
		numbers := []int{1,2,3,4,5}
		
		result := ArraySum(numbers)
		expected := 15
	
		verifyCorrectness(t, result, expected, numbers)
	})
	t.Run("sum array numbers no numbers", func(t *testing.T){
		numbers := []int{}
		
		result := ArraySum(numbers)
		expected := 0
	
		verifyCorrectness(t, result, expected, numbers)
	})
}