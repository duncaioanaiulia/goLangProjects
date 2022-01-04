package collections

import (
	"fmt"
	"reflect"
)

func Struct() {
	DeclareStruct()
	CreateAnonymousStruct()
	CompositionStruct()
	GetTagInStruct("Name")
}

type Developer struct {
	Id         int
	Name       string `TAG`
	Companions []string
}

type TeamLead struct {
	Developer
	ageExperience int
}

func DeclareStruct() {
	fmt.Println("TRAINING STRUCT: Declare:")
	dev := Developer{
		Id:   3,
		Name: "Ioana Dunca",
		Companions: []string{
			"Roxana",
			"Elisa",
			"Mariana",
			"Monica",
		},
	}
	fmt.Println(dev)
}

func CreateAnonymousStruct() {
	fmt.Println("TRAINING STRUCT: Anonymous Struct:")
	bestDeveloper := struct{ name string }{"Ioana Dunca"}
	anotherDeveloper := bestDeveloper
	anotherDeveloper.name = "Roxana Apostol"
	fmt.Println(bestDeveloper)
	fmt.Println(anotherDeveloper)
}

func CompositionStruct() {
	fmt.Println("TRAINING STRUCT: Composition Struct:")
	fmt.Println("WARNING: Golang does not support inheritance, use a concept called composition")

	teamLead := TeamLead{}
	teamLead.ageExperience = 6
	teamLead.Developer.Id = 24
	teamLead.Companions = []string{"Ioana", "Raluca"}
	teamLead.Name = "Fluture"

	fmt.Println(teamLead)
}

func GetTagInStruct(tag string) {
	fmt.Println("TRAINING STRUCT: Tag:")
	t := reflect.TypeOf(Developer{})
	field, _ := t.FieldByName(tag)
	fmt.Println(field.Tag)
}
