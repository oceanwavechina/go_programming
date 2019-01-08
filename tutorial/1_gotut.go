package main

import (
	"fmt"
	"math/rand"
)

func add(x, y float64) float64 {
	return x + y
}

const gNum1 = 5.6

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	fmt.Println("A number from 1-100", rand.Intn(100))

	// var num1, num2 float64 = 5.6, 9.5
	num1, num2 := 5.6, 9.5 // 此时，num1, num2的类型都是float64
	fmt.Println("add return : ", add(num1, num2))

	w1, w2 := "Hey", "There"
	// fmt.Println("multiple: ", multiple(w1, w2))
	fmt.Println(multiple(w1, w2))

	var a int = 62
	var b float64 = float64(a)

	x := a

	fmt.Println(x, b, a)
}
