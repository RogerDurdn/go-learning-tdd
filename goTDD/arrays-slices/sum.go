package sum

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(collections ...[]int) []int {
	var result []int
	for _, collection := range collections {
		result = append(result, sum(collection))
	}
	return result
}

func sumAllTails(collections ...[]int) []int {
	var result []int
	for _, collection := range collections {
		total := 0
		if len(collection) > 0 {
			total = sum(collection[1:])
		}
		result = append(result, total)
	}
	return result
}
