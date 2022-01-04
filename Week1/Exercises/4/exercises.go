package ex4

import (
	"fmt"
)

func main() {
}

func GetAddressPointer(val int) *int {
	fmt.Println("--------- Ex 1")
	pointer := &val
	fmt.Printf("Value: %v\n", val)
	fmt.Printf("Address %p\n", pointer)
	return pointer
}

func GetDetailsPointer(pointer *int) (*int, int) {
	fmt.Println("--------- Ex 2")
	fmt.Printf("Address %p\n", pointer)
	fmt.Printf("Value that the pointer points to %v\n", *pointer)
	return pointer, *pointer
}

type developer struct {
	id    int
	name  string
	phone string
}

func (dev *developer) ChangeDataToPhone(newPhone string) string {
	fmt.Println("--------- Ex 3")
	fmt.Printf("Current value: %v\n", dev.phone)
	dev.phone = newPhone
	fmt.Printf("New value: %v\n", dev.phone)
	return dev.phone
}

type Speaker interface {
	speak() string
}

type english struct {
	word string
}

type chinese struct {
	word string
}

func (en english) speak() string {
	return en.word
}
func (ch chinese) speak() string {
	return ch.word
}

func SayHello(s Speaker) string {
	fmt.Println(s.speak())
	return s.speak()
}

func WhatAChineseAndAnEnglishmanSayWhenTheyMeet(eng english, ch chinese) (string, string) {
	fmt.Println("--------- Ex 4 and 5")
	return SayHello(eng), SayHello(ch)
}
