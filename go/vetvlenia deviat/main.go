package main

import "fmt"

func hello(name string) {
	fmt.Println("привет,", name)
}
func main() {
	fmt.Println("начало")
	hello("Дауд")
	fmt.Println("конец")
}
