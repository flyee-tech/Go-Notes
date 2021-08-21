package main

import "fmt"

// Go 的返回值可以被命名，这种情况可以被视为在函数的顶部的变量

//TODO 这里没有看懂，何为命名？？？

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
