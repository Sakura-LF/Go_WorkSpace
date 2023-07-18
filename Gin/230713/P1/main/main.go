package main

import "fmt"

type Student struct {
	Name string
	next *Student
}

func main() {
	fmt.Println("Test")
}
