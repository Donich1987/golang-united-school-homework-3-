package homework
func reverse(input []int64) (result []int64) {
	slice := make([]int64, len(input))
	copy(slice, input)
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			slice[j], slice[j+1] = slice[j+1], slice[j]
		}
	} 
 	return slice
}
