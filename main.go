package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MouseSleep = 100

	robotgo.ScrollDir(10, "up")
	robotgo.ScrollDir(20, "right")

	robotgo.Scroll(0, -10)
	robotgo.Scroll(100, 0)

	robotgo.MilliSleep(100)
	robotgo.ScrollSmooth(-10, 6)
	// robotgo.ScrollRelative(10, -100)

	robotgo.Move(10, 20)
	robotgo.MoveRelative(0, -10)
	robotgo.DragSmooth(10, 10)

	robotgo.Click("wheelRight")
	robotgo.Click("left", true)
	robotgo.MoveSmooth(100, 200, 1.0, 10.0)

	robotgo.Toggle("left")
	robotgo.Toggle("left", "up")

	fmt.Println("Press enter to exit")
	// python input() to prevent the program from exiting immediately
	fmt.Scanln()
}
