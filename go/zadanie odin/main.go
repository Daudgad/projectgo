package main

import "fmt"

func main() {
	number := 7
	for i := 7; i <= 10; i++ {
		fmt.Println(number, "x", i, "=", number*i)
	}
}
