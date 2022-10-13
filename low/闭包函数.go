package main

import "fmt"

func biBao(name string, age int) func() string {
	fmt.Println(name, age, "岁")
	return func() string {
		return "干就完了 " + name
	}
}

func main() {
	var s1 = biBao("张三", 26)
	fmt.Println(s1())

	go func() {
		var s2 = biBao("李四", 29)
		fmt.Println(s2())
	}()
}
