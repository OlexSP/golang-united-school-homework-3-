package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	cellCounter := 0
	for _, v := range input {
		result += v
		if v != 0 {
			cellCounter++
		}
	}
	result = result / float32(cellCounter)
	return
}
