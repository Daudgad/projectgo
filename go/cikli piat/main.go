package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}
