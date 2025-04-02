package main

import "fmt"

func inc() int {
	return 1
}

func curr() int {
	return 2
}

func adder() (func() int, func() int) {
	return inc, curr
}

func main() {
	i, c := adder()
	fmt.Println(i)
	fmt.Println(c())
}

//★
