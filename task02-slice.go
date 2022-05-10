package homework

func reverse(input []int64) (result []int64) {
	j := 0
	for i := 0; i <= len(input)-1; i++ {
		result[j] = input[len(input) - 1 - i]
		j = j + 1
	}
	return result
}
