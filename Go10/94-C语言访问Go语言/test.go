package main

import "C"
import "fmt"

//export say
func say() {
	fmt.Println("hello world")
}
