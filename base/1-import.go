package main

//如何导入包
import (
	"fmt"
	"math"
)

//另一种导入包的方式
//import "fmt"
//import "math"

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
