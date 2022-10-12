package main

import "fmt"

func say() {
	fmt.Println("人生苦短须尽欢")
}

func say1(name string) {
	fmt.Println("人生苦短须尽欢,", name)
}

/**
变长的数组
*/
func say3(names ...string) {
	fmt.Println("人生苦短须尽欢,", names)
	for i, item := range names {
		fmt.Println(i, item)
	}
}

func judge(n int) (s string) {
	if n%2 == 0 {
		return "偶数"
	} else {
		return "奇数"
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) (m int) {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	say()
	say1("apple")
	say3("huawei", "apple")
	fmt.Println(judge(123))
	fmt.Println("阶乘：5!=", Factorial(5))
	fmt.Println("斐波那契数列")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
