package sequences

func ArraySum(numbers []int) (sum int) {
	for _, number := range numbers{
		sum += number
	}

	return
}

func ArraysSum(arrays ...[]int) (sum []int)  {
	for _, array := range arrays{
		sum = append(sum, ArraySum(array))
	}

	return
}