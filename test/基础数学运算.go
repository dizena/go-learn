package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 2
	var b = 3
	c := 5
	var d, e int = 10, 20
	fmt.Println(a, b, c, d, e)

	fmt.Println(10 + 3)
	fmt.Println(10 - 3)
	fmt.Println(10 * 3)
	fmt.Println(10 / 3)
	fmt.Println(10 % 3)
	fmt.Println(10.0 / 3)
	fmt.Println(10 / 3.0)

	//
	fmt.Println((1 + 2 + 3) * 4 / 6)

	fmt.Println("四舍五入", math.Round(3.5))
	fmt.Println("四舍五入", math.Round(-3.5))
	fmt.Println("绝对值", math.Abs(-15))
	fmt.Println("乘方", math.Pow(2, 3))
	fmt.Println("开平方", math.Sqrt(9))
	fmt.Println("开多方", math.Pow(8, 1.0/3))

	fmt.Println("三角函数", math.Sin((30.0/180)*math.Pi))
	fmt.Println("三角函数", math.Cos((30.0/180)*math.Pi))
	fmt.Println("三角函数", math.Tan((30.0/180)*math.Pi))

	fmt.Println("反三角函数", math.Asin(0.5))
	fmt.Println("反三角函数", math.Acos(0.86))
	fmt.Println("反三角函数", math.Atan(0.577))
}
