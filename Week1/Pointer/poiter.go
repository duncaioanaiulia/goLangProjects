package pointer

import "fmt"

type Wallet struct{}

func main() {
	fmt.Println("----- TRAINING POINTER:")
	CreatePointer(34)
	PointerToArray()
	CreatePointerToStruct()
	ArrayMemoryAndSlicePointer()
	MapInPointerContext()
}

func CreatePointer(val int) {
	fmt.Printf("Value is: %v\n", val)
	var pointer *int = &val
	fmt.Printf("Numerical representation in the memory/ Address : %v\n", pointer)
	fmt.Printf("Reference pointer %v\n", *pointer)

	val = 14
	*pointer = 45
	fmt.Printf("Change value: %v\n", val)
	fmt.Printf("Reference pointer %v\n", *pointer)
}

func PointerToArray() {
	fmt.Println("----- Pointer To Array:")
	fmt.Println("WARNING!! If you want to use point arithmetic use packege 'unsafe'")
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
}

type pointerStruct struct {
	id int
}

func CreatePointerToStruct() {
	fmt.Println("----- Pointer To Struct:")
	fmt.Println("WARNING!! Check pointer for <nil>")
	fmt.Println("WARNING!! Struct it is a address, so is not need for reference (* )")
	var ps *pointerStruct
	fmt.Println(ps)
	ps = &pointerStruct{id: 45}
	ps = new(pointerStruct)
	ps.id = 45
	fmt.Println(ps)
	fmt.Printf("Value from Struct %v\n", ps.id)
}

func ArrayMemoryAndSlicePointer() {
	fmt.Println("----- Array is Memory And Slice ia a Pointer to an array:")
	array1 := [3]int{0, 1, 2}
	array2 := array1
	fmt.Printf("Array %v %v\n", array1, array2)
	array1[2] = 1
	fmt.Printf("Array change %v %v\n", array1, array2)
	slice1 := []int{0, 1, 2}
	slice2 := slice1
	fmt.Printf("Slice %v %v\n", slice1, slice2)
	slice1[2] = 1
	fmt.Printf("Slice change %v %v\n", slice1, slice2)
}

func MapInPointerContext() {
	fmt.Println("----- Map In Pointer Context:")
	map1 := map[string]string{"false": "0", "true": "1"}
	map2 := map1
	fmt.Printf("Map %v %v\n", map1, map2)
	map1["true"] = "2"
	fmt.Printf("Map change %v %v\n", map1, map2)
}

func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	return 0
}
