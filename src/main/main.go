package main

import "fmt"

func main() {
	mark := make([]bool, 5)

	if mark[3] == false {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
