package arraysslices

import "fmt"

func main() {
	SumAll([]int{1, 2}, []int{0, 9})
	SumAllTails([]int{1, 2}, []int{0, 9})
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	fmt.Println(lengthOfNumbers)
	sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	fmt.Println(sums)
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	fmt.Println("The tail of a collection is all items in the collection except the first one")
	var sums []int
	for _, numbers := range numbersToSum {

		fmt.Println(numbers)
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	fmt.Println(sums)
	return sums
}
