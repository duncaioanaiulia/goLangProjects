package ex3

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	CountOddAndEven()
	fmt.Println(FizzBuzz(6))
	CreateAMap()
}

func CountOddAndEven() {
	rand.Seed(time.Now().Unix())
	slice := rand.Perm(50)
	fmt.Println(slice)
	oddCount := 0
	evenCount := 0
	for _, value := range slice {
		if value%2 == 0 {
			oddCount++
		} else {
			evenCount++
		}
	}
	fmt.Printf("Odd Count: %v and Even Count: %v \n", oddCount, evenCount)
}

func FizzBuzz(input int) string {
	if input%15 == 0 {
		return "FizzBuzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(input)
}

func CreateAMap() {
	cities := map[string]int{
		"Iasi":      452312,
		"Bucurest":  35432423,
		"Timisoara": 343323,
		"Cluj":      823424,
		"Brasov":    232443,
	}

	for key, value := range cities {
		fmt.Print(key)
		if value > 1000000 {
			fmt.Printf(" has  population greater than a million")
		}
		fmt.Println()
	}
}
