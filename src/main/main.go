package main

import "fmt"

func main() {

	marry := human{"marry"}
	//marry := &human{"marry"}
	Speak(marry)
}

type Speaker interface {
	speak() string
}

type human struct {
	Name string
}

func (h *human) speak() string {
	return fmt.Sprintf("my name is %s\n", h.Name)
}

func Speak(s Speaker) {
	fmt.Println(s.speak())
}
