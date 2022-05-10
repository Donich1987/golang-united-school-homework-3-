package homework

import "sort"


func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))
	value := make([]string, 0, len(input))
	for i := range input {
		keys = append(keys, i)

	}
	sort.Ints(keys)

	for _, j := range keys {
		value = append(value, input[j])
		j++

	}
	return value
}
