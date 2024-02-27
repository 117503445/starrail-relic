package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Hello, World!")
	robotgo.MouseSleep = 100
	robotgo.Move(10, 20)
}
