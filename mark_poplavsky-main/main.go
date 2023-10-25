package main

import (
	"fmt"
	"upi/first"
	"upi/second"
)

func main() {
	fmt.Println("введите задание 1 или 2")
	var input int
	fmt.Scan(&input)
	if input == 1 {
		first.First()
	}
	if input == 2 {
		second.Second()
	}
	if input != 1 && input != 2 {
		fmt.Println("неправильный ввод")
	}
}
