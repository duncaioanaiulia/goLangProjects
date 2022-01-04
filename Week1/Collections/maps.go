package collections

import "fmt"

func Maps() {
	col := DeclareMaps()
	EmptyMap()
	ManipulateValueOnMap(col)
	col = AddKey(col)
	col = DeleteKey(col)
	CheckKey(col, "Cluj")
	LenghtMap(col)
}

func DeclareMaps() map[string]int {
	fmt.Println("TRAINING MAPS: declare")

	devsEPAM := map[string]int{
		"Iasi":      45,
		"Bucurest":  35,
		"Timisoara": 3,
		"Cluj":      8,
	}
	fmt.Printf("Map is %v:\n ", devsEPAM)
	return devsEPAM
}

func EmptyMap() map[string]int {
	fmt.Println("TRAINING MAPS: empty map:")

	emptyMap := make(map[string]int, 10)
	fmt.Println("WARNING: A slice can't be a key to the map. NO []int YES [...]int")
	fmt.Printf("Empty Map is %v:\n ", emptyMap)
	return emptyMap
}

func ManipulateValueOnMap(col map[string]int) {
	fmt.Println("TRAINING MAPS: data from map:")

	fmt.Println("At key Iasi: ", col["Iasi"])
}

func AddKey(col map[string]int) map[string]int {
	fmt.Println("TRAINING MAPS: add a key in map:")
	fmt.Println("WARNING: The order is not something specific")

	col["Galati"] = 2
	fmt.Println(col)
	return col
}

func DeleteKey(col map[string]int) map[string]int {
	fmt.Println("TRAINING MAPS: delete in map:")
	delete(col, "Galati")
	fmt.Println(col)
	return col
}

func CheckKey(col map[string]int, key string) {
	fmt.Println("TRAINING MAPS: check key:")
	val, hasKey := col[key]
	fmt.Printf("Exist key %v with value %v in the map? %v \n ", key, val, hasKey)
}

func LenghtMap(col map[string]int) {
	fmt.Println("TRAINING MAPS: Lenght:")
	fmt.Println(len(col))
}
