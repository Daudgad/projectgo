package main

import "fmt"

func hello() {
	fmt.Println("привет эта функция вызовется")
}
func main() {
	fmt.Println("начало")
	hello()
	fmt.Println("конец")
}
