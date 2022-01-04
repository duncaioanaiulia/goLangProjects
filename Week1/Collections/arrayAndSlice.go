package collections

import "fmt"

func ArraysAndSlice() {
	fmt.Println("TRAINING ARRAY:")
	DeclareArray()
	ArrayOfArray()
	AssingArrayToACopy()
	AssingArrayToAddress()
	LenghtAndCapacityArray()
	SlicesArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 6)
	collection := MakeAnCollection(3, 100)
	collection = AppendCollection(collection, []int{9, 8, 7}, 2, 3, 4, 5)
	ShiftOperator(collection)
}

func DeclareArray() {
	fmt.Println("TRAINING ARRAY: declare")
	fmt.Println("WARNING: An array keeps the value in memory and need a size")
	fmt.Println("WARNING: When decalrate with [...] is an array, when declarate with [] is a slice")
	grades := [3]int{97, 85, 93}
	grades = [...]int{97, 85, 93}
	fmt.Printf("Grades: [3]int %v\n", grades)

	var students [3]string
	fmt.Printf("Strudents: %v\n", students)

	students[0] = "A"
	students[1] = "B"
	students[2] = "C"

	fmt.Printf("Strudents: %v\n", students)
	fmt.Printf("Strudent 1: %v\n", students[1])

}
func ArrayOfArray() {
	fmt.Println("TRAINING ARRAYS: Array Of Array")
	fmt.Println("warning: redundant type from array, slice, or map composite  ")
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
}

func AssingArrayToACopy() {
	fmt.Println("TRAINING ARRAYS: Assing Array")
	a := [...]int{1, 2, 3}
	fmt.Println(a)
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}

func AssingArrayToAddress() {
	fmt.Println("TRAINING ARRAYS: Assing Array")
	a := [...]int{1, 2, 3}
	fmt.Println(a)
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(*b)
}

func LenghtAndCapacityArray() {
	fmt.Println("TRAINING ARRAYS: Lenght And Capacity")
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}

func SlicesArray(collection []int, begainSlice int, endSlice int) {
	fmt.Println("TRAINING ARRAYS: Slices")
	fmt.Println("WARNING: When is a slice it will point to a colection it will NOT be a copy")
	fmt.Println("WARNING: A slice is a pointer to an array and don't need a size")
	b := collection[:]
	c := collection[begainSlice:]
	d := collection[:endSlice]
	e := collection[begainSlice:endSlice]
	fmt.Println(collection)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func MakeAnCollection(lenght int, capacity int) []int {
	fmt.Println("TRAINING ARRAYS: Make a Slice")
	collection := make([]int, lenght, capacity)
	PrintDetailsCollection(collection)
	return collection
}

func AppendCollection(collection []int, addCol []int, prams ...int) []int {
	fmt.Println("TRAINING ARRAYS: Append , is for adding a new value")

	collection = append(collection, 1)
	PrintDetailsCollection(collection)

	collection = append(collection, prams...)
	PrintDetailsCollection(collection)

	collection = append(collection, addCol...)
	PrintDetailsCollection(collection)

	return collection
}

func ShiftOperator(col []int) {
	fmt.Println("TRAINING ARRAYS: Shift Operator")

	fmt.Printf("Collection is: %v\n", col)

	slice := col[:len(col)-1]
	fmt.Printf("Remove last element:: %v\n", slice)

	slice = append(col[:5], col[6:]...)
	fmt.Printf("Remove an element at position 5: %v\n", slice)
}
func PrintDetailsCollection(collection []int) {
	fmt.Println(collection)
	fmt.Printf("Length: %v\n", len(collection))
	fmt.Printf("Capacity: %v\n", cap(collection))
}
