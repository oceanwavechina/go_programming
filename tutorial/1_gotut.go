package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func add(x, y float64) float64 {
	return x + y
}

const gNum1 = 5.6

func multiple(a, b string) (string, string) {
	return a, b
}

func testBasic() {
	fmt.Println("A number from 1-100", rand.Intn(100))

	// var num1, num2 float64 = 5.6, 9.5
	num1, num2 := 5.6, 9.5 // 此时，num1, num2的类型都是float64
	fmt.Println("add return : ", add(num1, num2))

	w1, w2 := "Hey", "There"
	// fmt.Println("multiple: ", multiple(w1, w2))
	fmt.Println(multiple(w1, w2))

	var a = 62
	var b = float64(a)
	x := a
	fmt.Println(x, b, a)

	addr := &a

	fmt.Println("address:", addr, ", value:", *addr)
	fmt.Println("*addr:", *addr, ", *addr**addr:", *addr**addr, ", a:", a)
}

func testForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}

	// := 初始化并赋值
	ii := -1
	for ; ii < 1; ii++ {
		fmt.Println("ii=", ii)
	}

	// = 只赋值
	ii = -1
	for ii < 1 {
		fmt.Println("idx=", ii)
		ii++
	}

	for {
		fmt.Println("infinite loop")
		break
	}
}

func testReflect() {
	// Tag最常用的大概就是在marshaling
	type User struct {
		Name   string `tag_1:"User Name"`
		Passwd string `tag_2:"User Password"`
	}

	user := &User{"liuyanan", "pass"}
	s := reflect.TypeOf(user).Elem() // 通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)
	}
}

// func main() {
// 	testBasic()
// 	testForLoop()
// 	testReflect()
// }
