package main

import "fmt"

func show(sk []string) {
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	skills := []string{"JS", "GO", "Python"}
	show(skills)

	show(skills[1:3])
}
