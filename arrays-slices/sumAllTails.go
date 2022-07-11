package main

func SumAllTails(numbersToSum ...[]int) []int {
	
	var sums []int

	for _, numbers := range numbersToSum{
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}