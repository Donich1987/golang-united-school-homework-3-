package homework

func average(input [15]float32) (result float32) {
//Place your code here
var sum float32 = 0
	for _, number := range input {
		sum += number
	}	
	N := float32(len(input))
	av := sum / N
	return av
}
