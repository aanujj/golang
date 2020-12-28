package main

import "fmt"

func calc(x int, y int) (int, int) {

	add := x + y
	sub := x - y

	return add, sub

}

func main() {

	num1 := 10
	num2 := 20

	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)

}
