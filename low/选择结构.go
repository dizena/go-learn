package main

import "fmt"

func main() {
	fmt.Println("输入数字：")
	var today int
	var _, err = fmt.Scan(&today)
	if err != nil {
		return
	}
	switch today {
	case 1:
		fmt.Println("今天是周一")
	case 2:
		fmt.Println("今天是周二")
	case 3:
		fmt.Println("今天是周三")
	case 4:
		fmt.Println("今天是周四")
	case 5:
		fmt.Println("今天是周五")
	case 6:
		fmt.Println("今天是周六")
	case 7:
		fmt.Println("今天是周日")
	default:
		fmt.Println("输入错误")
	}

	switch {
	case today < 10:
		fmt.Println("个位数")
	case today > 10 && today < 100:
		fmt.Println("两位数")
	}
}
