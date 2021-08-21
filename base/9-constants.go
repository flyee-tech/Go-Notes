package main

import "fmt"

// 使用 const 关键字，就可以定义常量
// 常量可以定义在包中，也可以定义在函数中
// 貌似常量的首字母是大写呀

const Name = "裴二龙"

func main() {
	const World = "世界"
	fmt.Println(Name)
	fmt.Println(World)
}
