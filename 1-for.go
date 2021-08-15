package main

import "fmt"

func main() {
	fmt.Println("test")
	NormalFor()
	IgnoreFor()
	//ForeverFor()
}

// NormalFor 最普通的For循环的写法
func NormalFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("normal for output : %d \n", sum)
}

// IgnoreFor 初始化语句和后置语句可以省略
func IgnoreFor() {
	sum := 1
	for sum < 1000 { // 也可以加上分号 for ; sum < 1000 ; {
		sum += sum
	}
	fmt.Printf("ignore for output : %d \n", sum)
}

// ForeverFor 死循环
func ForeverFor() {
	for {

	}
}
