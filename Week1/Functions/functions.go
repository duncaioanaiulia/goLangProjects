package functions

import "fmt"

func main() {
	fmt.Println("----- TRAINING FUNCTIONS:")
	BasicSyntax()

	pointer := "pointer"
	BasicSyntaxWithParameters("First Param", 2, &pointer, "Param 3", "Param 4")

	sum := ReturnValue(1, 2, 3, 4, 5)
	fmt.Printf("Sum is: %v\n", *sum)

	FunctionWithPanic()
	AnonymousFunc()
	FunctionAsType()

	g := greeter{
		id:   5,
		name: "Name",
	}
	g.Methods()

	fmt.Println(g.name)
}

func BasicSyntax() {
	fmt.Println("an function start with func keyword, name of the function, (), after the body is put in {}")
}

func BasicSyntaxWithParameters(parameter string, secondParam int, pointer *string, slice ...string) {
	fmt.Printf("The param is: %v , %v , %v , %v \n", parameter, secondParam, *pointer, slice)
}

func ReturnValue(values ...int) *int {
	fmt.Println("----- Return values:")
	result := 0
	for _, v := range values {
		fmt.Printf("%v ,", v)
		result += v
	}
	return &result
}

func FunctionWithPanic() {
	d, err := divider(5.0, 0.0)
	if (err) != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divider(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func AnonymousFunc() {
	func() {
		fmt.Println("------- Anonymous Functions")
	}()
}

func FunctionAsType() {
	fmt.Println("------- Function As Type")
	var f func(int) = func(index int) {
		fmt.Println(index)
	}

	for index := 0; index < 5; index++ {
		f(index)
	}
}

type greeter struct {
	id   int
	name string
}

func (g greeter) Methods() {
	fmt.Println("------- Methods or Extension to a variable")
	fmt.Println(g.id, g.name)

	fmt.Println("WARNING!!! Is just a copy of the struct")
	g.name = "test"
}
