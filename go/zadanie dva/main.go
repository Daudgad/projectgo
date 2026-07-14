package main

import "fmt"

func main() {
	phones := map[string]string{
		"Салман": "89995337179",
		"Даниял": "89884588009",
	}
	fmt.Println("телефонная книга")
	fmt.Println("...................")
	for name, phone := range phones {
		fmt.Println(name, ".", phone)
	}
}
