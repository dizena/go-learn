package main

import "fmt"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	BBA = (iota + 1) * 1000
	BBC
	BBD
	BBJ
)

func main() {
	var v8 = '五'
	fmt.Printf("五的类型%T值是%v字符是%c\n", v8, v8, v8)
	fmt.Println(Sunday, BBJ)

	var a, b = 8, 10
	fmt.Printf("a= %d 二进制%b\n", a, a)
	fmt.Printf("b=%d 二进制%b\n", b, b)
	fmt.Println("a <<1", a<<1)
	fmt.Println("a ^ b", a^b)
	fmt.Println("a & b", a&b)
	fmt.Println("a | b", a|b)
}
