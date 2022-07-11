package main

// Refactor
func SumAll(numbersToSum ...[]int) []int {
	
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums,Sum(numbers))
	}

	return sums
}

// func SumAll(numbersToSum ...[]int) (sums []int) {

// 	lengthOfNumbers := len(numbersToSum)

// 	sums = make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return	sums
// }