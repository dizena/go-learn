package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	answer := myRand.Intn(1000)
	fmt.Println(answer)
	for {
		fmt.Println("输入猜测的数字：")
		var guess string
		_, _ = fmt.Scan(&guess)
		if guess == "exit" {
			break
		}
		guessNum, _ := strconv.Atoi(guess)
		switch {
		case guessNum > answer:
			fmt.Println("猜大了")

		case guessNum < answer:
			fmt.Println("猜小了")

		default:
			fmt.Println("猜对了")
			break
		}
	}
}
