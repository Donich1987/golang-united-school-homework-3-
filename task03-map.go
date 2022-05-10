package homework
import "sort"
func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	i := 0
	for kay := range input {
		keys[i] = kay
		i++
	}
	sort.Ints(keys)
	for k := range keys {
		result = append(result, input[k])
	}
	return result	
}
