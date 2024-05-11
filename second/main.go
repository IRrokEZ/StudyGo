package main

import "fmt"

func main() {
	var num int32
	fmt.Println("Please enter number")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
	fmt.Scan(&num)
}
