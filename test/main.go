package main

import "fmt"

func add(a int, b int) int {
	var sum = a + b
	return sum
}

func main() {
	const pi = 3.14
	var r = 1.0
	area := pi * r * r
	// var area = ne  :=
	fmt.Println("Area = ", area)
	fmt.Println("hello world", add(3, 2))

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	/*
		常量测试
	*/
	const (
		lightSpeed  = 30 * 10000
		earthCircle = 50000
	)

	fmt.Println("光速 ", lightSpeed, "周长", earthCircle)

	/*
		变量测试
	*/
	var (
		age    = 21
		height = 179
		weight = 65
	)
	var married = false

	/*
		运算符
	*/
	age++
	weight = 72
	married = true
	fmt.Println(age, height, weight, married)

	if age < 18 {
		fmt.Println("未成年")
	} else if age == 22 {
		fmt.Println("毕业了")
	}

	/*
		switch
	*/
	var grade string = "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("优良")
	default:
		fmt.Println("一般")
	}

	var v1 = 'a'
	fmt.Printf("v1类型%T\n", v1)
	fmt.Printf("v1值是%v\n", v1)
	fmt.Printf("v1字符%c\n", v1)
}
