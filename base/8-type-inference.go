package main

import "fmt"

// 类型推导，如果不指定类型，类型会由右表达式推导得出。

func main() {
	v := 1
	v2 := 1.1
	v3 := 0.876 + 0.5i
	fmt.Printf(" %T \n %T \n %T \n", v, v2, v3)
}
