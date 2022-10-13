package main

import "fmt"

func aaa(a int) int {
	return a
}
func main() {
	fmt.Println("开门")
	defer fmt.Println("关门")
	fmt.Println("开灯")
	defer fmt.Println("关灯")
	fmt.Println("脱衣服")
	fmt.Println("挂衣服")
	fmt.Println("开电视")
	fmt.Println("坐下")
	fmt.Println("喝水")
	fmt.Println(aaa(11))
}
