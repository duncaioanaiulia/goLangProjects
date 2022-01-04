package looping

import (
	"fmt"
)

func main() {
	InitializeFor(0, 5)
	InitializeTwoLoops(0, 6)
	ForCheckIncrementer(1, 7)
	ForWhenIsPassedTheIndex(2, 6)
	ForWTFWorksWithoutSemicolons(1, 3)
	ForWithBreak(0, 7)
	ForWithContinue(0, 10)
	NestedLoops(1, 3)
	NestedLoopsWithLabel(1, 3)
	LoopToACollection([]int{1, 2, 3})
	devsEPAM := map[string]int{
		"Iasi":      45,
		"Bucurest":  35,
		"Timisoara": 3,
		"Cluj":      8,
	}
	LoopToMaps(devsEPAM)
	LoopToAString("Say Hello To EPAM")
}

func InitializeFor(startValue int, endValue int) {
	fmt.Println("TRAINING LOOPING: Initialize FOR:")
	for i := startValue; i < endValue; i++ {
		fmt.Println(i)
	}
}

func InitializeTwoLoops(startValue int, endValue int) {
	fmt.Println("TRAINING LOOPING: Two Loops:")
	fmt.Println("WARNING i++ is not a expression is a statement")
	for i, j := startValue, startValue; i < endValue; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}

func ForCheckIncrementer(startValue int, endValue int) {
	fmt.Println("TRAINING LOOPING: For Check Incrementer:")
	for i := startValue; i < endValue; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}
}

func ForWhenIsPassedTheIndex(index int, endValue int) {
	fmt.Println("TRAINING LOOPING: For When Is Passed The Index:")
	for index < endValue {
		fmt.Println(index)
		index++
	}
	fmt.Printf("Last value of the index: ")
	fmt.Println(index)
}

func ForWTFWorksWithoutSemicolons(index int, endValue int) {
	fmt.Println("TRAINING LOOPING: WTF Works Without Semicolons:")
	for index < endValue {
		fmt.Println(index)
		index++
	}
	fmt.Printf("Last value of the index: ")
	fmt.Println(index)
}

func ForWithBreak(index int, endValue int) {
	fmt.Println("TRAINING LOOPING: For With Break")
	for {
		fmt.Println(index)
		index++
		if index == 5 {
			break
		}
	}
}
func ForWithContinue(startValue int, endValue int) {
	fmt.Println("TRAINING LOOPING: For With Continue")
	for index := startValue; index < endValue; index++ {
		if index%2 == 0 {
			continue
		}
		fmt.Println(index)
	}
}

func NestedLoops(startValue int, endValue int) {
	fmt.Println("TRAINING LOOPING: Nested Loops")
	for i := startValue; i < endValue; i++ {
		for j := startValue; j < endValue; j++ {
			fmt.Println(i, j, i*j)
		}
	}
}

func NestedLoopsWithLabel(startValue int, endValue int) {
	fmt.Println("TRAINING LOOPING: Nested Loops With Label")
Loop:
	for i := startValue; i < endValue; i++ {
		for j := startValue; j < endValue; j++ {
			fmt.Println(i, j)
			if i*j >= endValue {
				fmt.Println("enter in label")
				break Loop
			}
		}
	}
}

func LoopToACollection(collection []int) {
	fmt.Println("TRAINING LOOPING: Loop To A Collection")
	for key, value := range collection {
		fmt.Println(key, value)
	}
}

func LoopToMaps(collection map[string]int) {

	fmt.Println("TRAINING LOOPING: Loop To Maps")
	fmt.Println("Print collection")
	for key, value := range collection {
		fmt.Println(key, value)
	}

	fmt.Println(" - Print keys from collection:")
	for key := range collection {
		fmt.Print(key)
		fmt.Print(" ")
	}

	fmt.Println()
	fmt.Println(" - Print values from collection:")
	for _, value := range collection {
		fmt.Print(value)
		fmt.Print(" ")
	}
	fmt.Println()
}

func LoopToAString(collection string) {
	fmt.Println("TRAINING LOOPING: Loop To A Strings")
	for key, value := range collection {
		fmt.Println(key, string(value))
	}
}
