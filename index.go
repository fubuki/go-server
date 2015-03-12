package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/test/:name", func(res http.ResponseWriter, req *http.Request) { // res 和 req 是由 Martini 注入
		res.WriteHeader(200) // HTTP 200
	})

	m.NotFound(func() {
		// handle 404
	})
	m.Run()
}

func test() {
	var a int
	var b int
	a = 20
	b = 16
	fmt.Println(a, b)

	if a > 15 {
		fmt.Println("ok")
	}
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	for sum > 0 {
		sum = sum - 1
	}
	fmt.Println(sum)

J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J
			}
		}
	}

	list := []string{"a", "b", "c", "d", "e"}
	for k, v := range list {
		fmt.Println(k, v)
	}

	var arr [10]int
	arr[0] = 10
	sl := make([]int, 10)
	sl[1] = 123

	fun := func(a int, b int) {
		fmt.Println(a + b)
	}

	fun(1, 3)

	callback(10, fun)
	fmt.Println(fab(5))

	// new make 差別
	/*
		type Sorter interface {
			Len() int
			Less(i, j int) bool
			Swap(i, j int)
		}

		type Xi []int
		type Xs []string
	*/

	/*
		ci := make(chan int)
		cs := make(chan string)
		cf := make(chan interface{})
	*/
}

type Person struct {
	name string "namestr"
	age  int
}

func ShowTag(i interface{}) {

}

func callback(y int, f func(int, int)) {
	f(y, y)
}

func myfunc() {
	i := 0
Here:
	println(i)
	i++
	goto Here
}

func fab(n int) int {
	first := 1
	second := 1
	result := 2
	if n <= 2 {
		return 1
	}
	for i := 3; i <= n; i++ {
		result = first + second
		second = first
		first = result
	}

	return result
}
