package arr_sum

func main() {
}

func Sum(numbers [5]int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbersToSum... []int) (sums []int) {
	return
}