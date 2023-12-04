package main

import "fmt"

func main() {
	a := 1
	b := 5
	c := "let"
	d := false
	if d == true {
		fmt.Println("Sane", b-a, c)
	}
	fmt.Println("Sane", a+b, c)
}
