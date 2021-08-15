package main

import "fmt"

func main() {
	// defer 语句会将函数推迟到外层函数返回之后执行
	defer fmt.Println("world")
	fmt.Println("hello")
	DeferStack()
}

// DeferStack defer的函数会压入一个栈中
func DeferStack() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
