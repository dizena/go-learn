package main

import "fmt"

func test(a, b int) (int, int, int, float32) {
	var c = a + b
	var d = a - b
	var e = a * b
	var f = float32(a * 1.0 / b)
	return c, d, e, f
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("For Over")

	var a = 10
LOOP:
	for a < 20 {
		a++
		if a == 15 {
			a = a + 2
			goto LOOP
		}
		fmt.Println("a=", a)
	}

	var c, d, e, _ = test(5, 10)
	var _, _, _, f = test(5, 10)
	fmt.Println("a,b的和差倍分：", c, d, e, f)
}
