package ifswitch

import (
	"fmt"
	"math"
)

func main() {
	InitializerIf(45, 56)
	IfShortCircuiting(-1, 10)
	MathWithIf(0.123)
	LearnSwitchWithTag()
	LearnSwitchWithTagless(5)
	LearSwitchWithTypes()
}

func InitializerIf(chairs int, guess int) {
	fmt.Println("TRAINING IF: Training IF:")
	if guess < chairs {
		fmt.Println("Too low")
	} else if guess > chairs {
		fmt.Println("Too high")
	} else {
		fmt.Println("You got it!")
	}
	fmt.Println(guess <= chairs, guess >= chairs, guess != chairs)
}

func IfShortCircuiting(chairs int, guess int) {
	fmt.Println("TRAINING IF: with Short Circuiting:")
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("It is a SHORT Circuiting")
	} else {
		fmt.Println("It is a LONG Circuiting")
	}
}

func MathWithIf(number float64) {
	fmt.Println("TRAINING IF: Math IF:")
	if math.Abs(number/math.Pow(math.Sqrt(number), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}

func LearnSwitchWithTag() {
	fmt.Println("TRAINING SWITCH: with passing value/tag:")
	switch i := 2 + 3; i {
	case 1, 2, 3, 4:
		fmt.Println("First case")
	case 5, 6, 7, 8:
		fmt.Println("Second case")
	default:
		fmt.Println("Default")
	}
}

func LearnSwitchWithTagless(index int) {
	fmt.Println("TRAINING SWITCH: SWITCH with Tagless:")
	switch {
	case index <= 10:
		fmt.Println("First case")
		fallthrough
	case index >= 20:
		fmt.Println("Second case")
	default:
		fmt.Println("Default")
	}
}

func LearSwitchWithTypes() {
	fmt.Println("TRAINING SWITCH: when return a Type:")
	var i interface{} = [3]int{}
	var printText = "i is "
	switch i.(type) {
	case int:
		printText += "int"
	case float64:
		printText += "a float64"
	case string:
		printText += "string"
	case [3]int:
		printText += "[3]int"
	default:
		printText += "another type"
	}

	fmt.Println(printText)
}

func returnTrue() bool {
	fmt.Println("retrun true")
	return true
}
