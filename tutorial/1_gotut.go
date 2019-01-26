package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
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

func testMap() {
	grades := make(map[string]float32)

	grades["timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	fmt.Println(grades)

	TimsGrades := grades["timmy"]
	fmt.Println("timmy's grade:", TimsGrades)

	delete(grades, "timmy")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recoverd in cleanup! ", r)
	}
	// defer wg.Done()
}

func say(s string) {
	defer wg.Done()
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println("say", s)
		time.Sleep(time.Millisecond * 1000)
		if i == 2 {
			panic("oh dear, a 2")
		}
	}
	// wg.Done()
}

func testGoroutine() {
	// say("Hey")
	// say("there")

	// go say("Hey")
	// say("there")

	// 这两个都不能按预期执行，因为gorutine是异步执行的，程序结束后gorutiine还没有执行完
	// say("Hey")
	// go say("there")

	// go say("Hey")
	// go say("there")

	// 这个执行的并不完整，因为主线程的阻塞时间比gorutine的时间还要短
	// go say("Hey")
	// go say("there")
	// time.Sleep(time.Second)

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait() // 阻塞

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}

func foo() {
	// defer 类似一个stack的方式执行的
	defer fmt.Println("Done !")
	defer fmt.Println("Done again!")
	fmt.Println("Doing Some Stufff!")

	for i := 0; i < 5; i++ {
		defer fmt.Println("reverse output:", i)
	}
}

var wg2 sync.WaitGroup

func fooChan(c chan int, someValue int) {
	defer wg2.Done()
	c <- someValue * 5
	// 把值输出到channel中了，所以就不用return了

}

func testChannel() {
	fooVal := make(chan int)
	// fooVal := make(chan int, 10)
	/*
		go fooChan(fooVal, 5)
		go fooChan(fooVal, 3)

		v1 := <-fooVal
		v2 := <-fooVal

		// channel 会阻塞到gorutine返回结果
		v1, v2 := <-fooVal, <-fooVal
		fmt.Println("channel Values:", v1, v2)
	*/

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go fooChan(fooVal, i)
	}

	wg2.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
}

// func main() {
// 	/*
// 		testBasic()
// 		testForLoop()
// 		testReflect()
// 		testMap()
// 		testGoroutine()
// 		foo()
// 	*/
// 	testChannel()
// }
