package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var keySlice []int
	for i, _ := range input {
		keySlice = append(keySlice, i)
	}
	sort.Ints(keySlice)
	for _, v := range keySlice {
		result = append(result, input[v])
	}
	return
}
