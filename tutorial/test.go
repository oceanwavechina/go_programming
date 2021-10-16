package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	代码执行顺序
	1. 先初始化全局变量
	2. 执行 init 函数， init 函数是内置函数，与其位置无关,
		并且可以有多个init 函数(按声明的先后顺序执行), 只有init 函数才能有多个，其他的都不行
	3. 执行 main 函数
*/

var (
	GVar int = 3
	b byte
)

const (
	TYPE = 10
	TYPE2 = 0
	TYPE3		// 如果没有初始化的话，会使用上一个的值？
)

type Book struct {
	title string
	id int
}

// Gvsal := 100 只能在函数中出现



func sum(a int, b int) int {
	a = 2
	return a+b
}

func task(s string) {
	for i:=0; i<5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go task("task1")
	go task("task2")
	fmt.Println("hello, world" + GFlag)

	var balance = [8]int{0}

	var ret string
	for _, val := range balance {
		ret += strconv.Itoa(val)
	}
	fmt.Println("balance =", ret)

	for i:=0; i<3; i++ {
		fmt.Println(i)
	}

	a := 1
	b := 2
	fmt.Println("a+b =", sum(a, b), ", a =", a)

	book := Book{title:"go programming", id:1}
	fmt.Println((&book).title)		//  & 和 . 的优先级是: .会优先执行

	time.Sleep(100 * time.Second)
}

func init() {
	fmt.Println("init2 g_var=", GVar, ", b=", b, &b, &GVar)
}

func init() {
	fmt.Println("init g_var=", GVar, TYPE, TYPE2, TYPE3)
}