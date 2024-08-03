package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("2024.8.3 11:57")

	// robotgo.SaveCapture(name1, 10, 10, 30, 30)
	// robotgo.SaveCapture(name)

	robotgo.MouseSleep = 100

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	c := 0

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		c += 1
		name := fmt.Sprintf("%d.png", c)

		var err error

		// x, y, w, h
		img := robotgo.CaptureImg()
		err = robotgo.Save(img, name)
		if err = robotgo.Save(img, name); err != nil {
			panic(err)
		}
	})

	s := hook.Start()
	<-hook.Process(s)

	// fmt.Println("Press enter to exit")
	// // python input() to prevent the program from exiting immediately
	// fmt.Scanln()
}
