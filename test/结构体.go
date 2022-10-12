package main

import "fmt"

type Book struct {
	title  string
	author string
	price  float32
}

func main() {
	fmt.Println(Book{"Go语言", "谷歌", 125.5})
	fmt.Println(Book{title: "Java", author: "Oracle", price: 99.5})

	var b1 Book
	b1.title = "Python数据分析"
	b1.author = "诸葛亮"
	b1.price = 159.9

	fmt.Println(b1)
}
