package main

import "fmt"

func main() {
	var (
		m1 = "A:黄蓉"
		m2 = "B:程灵素"
	)
	fmt.Println(m1, m2, "你选择一个?")
	var name string
	var _, err = fmt.Scan(&name)
	if err != nil {
		return
	}
	if name == "A" {
		fmt.Println("选择了黄蓉")
	} else if name == "B" {
		fmt.Println("选择了程灵素")
	} else {
		fmt.Println("都不选")
	}

}
