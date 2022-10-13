package main

import "fmt"

func main() {
	str := "hello你好吗"
	fmt.Println("str len=", len(str))

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}
}
