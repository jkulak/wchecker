package main

import "fmt"

func main() {

	fmt.Println("Welcome in wchecker")
	s := "First test string"
	r := Check(&s)
	for err := range r {
		fmt.Println(err)
	}
}
