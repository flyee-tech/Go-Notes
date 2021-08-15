package main

import "fmt"

func main() {

	// if语句可以写前值表达式
	if v := 10; v == 10 {
		fmt.Println(v == 10)
	}

}
