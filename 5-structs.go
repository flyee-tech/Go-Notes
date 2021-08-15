package main

import "fmt"

// Vertex 使用 struct 关键字创建结构体（一组字段）
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	// 结构体的字段用点号来访问
	fmt.Println(v.X)
	// 结构体也可以有指针
	p := &v
	// 可以通过指针访问
	fmt.Println((*p).X)
	// 以上有点啰嗦，可以使用隐式间接引用
	fmt.Println(p.X)
	// 结构体文法，可以灵活创建
	fmt.Println(Vertex{})
	v3 := Vertex{Y: 2}
	fmt.Println(v3)
	p2 := &Vertex{1, 2}
	fmt.Println(p2)
}
