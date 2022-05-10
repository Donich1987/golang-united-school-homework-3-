package homework

func reverse(input []int64) (result []int64) {
	j := 0
	for i := len(input); i >= 0; i-- {
		result[j] = input[i]
		j = j + 1
	}
	return result
}
