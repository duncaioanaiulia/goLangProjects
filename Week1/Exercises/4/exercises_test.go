package ex4

import (
	"testing"
)

var pointerTest *int

func TestGetAddressPointer(t *testing.T) {

	var valTest int = 20

	got := GetAddressPointer(valTest)
	want := &valTest

	if *got != *want {
		t.Errorf("got %p want %p", got, want)
	}
	pointerTest = got
}

func TestGetDetailsPointer(t *testing.T) {

	if pointerTest == nil {
		t.Errorf("Pointer is empty \n")
	}

	gotAddress, gotValue := GetDetailsPointer(pointerTest)
	wantAddress := pointerTest
	wantValue := *pointerTest

	if *wantAddress != *gotAddress || wantValue != gotValue {
		t.Errorf("expected %p but got %p", gotAddress, wantAddress)
		t.Errorf("expected %d but got %d", gotValue, wantValue)
	}
}

func TestChangeDataToPhone(t *testing.T) {
	valStruct := developer{
		id:    1,
		name:  "Ioana",
		phone: "0000",
	}

	newPhone := "0001"

	got := valStruct.ChangeDataToPhone(newPhone)
	want := newPhone

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestWhatAChineseAndAnEnglishmanSayWhenTheyMeet(t *testing.T) {
	helloInEnglish := "Hello World"
	helloInChinese := "你好世界"
	eng := english{word: helloInEnglish}
	ch := chinese{word: helloInChinese}

	gotEng, gotCh := WhatAChineseAndAnEnglishmanSayWhenTheyMeet(eng, ch)
	wantEng := helloInEnglish
	wantCh := helloInChinese

	if wantEng != gotEng || wantCh != gotCh {
		t.Errorf("expected %v but got %v", gotEng, wantEng)
		t.Errorf("expected %v but got %v", gotCh, wantCh)
	}
}
