package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("2024.8.3 10:05")

	name := "test.png"
	name1 := "test_001.png"

	var err error

	img := robotgo.CaptureImg(10, 10, 30, 30)
	err = robotgo.Save(img, name)
	if err != nil {
		fmt.Println(err)
	}

	img1 := robotgo.CaptureImg()
	err = robotgo.Save(img1, name1)
	if err != nil {
		fmt.Println(err)
	}

	// robotgo.SaveCapture(name1, 10, 10, 30, 30)
	// robotgo.SaveCapture(name)

	robotgo.MouseSleep = 100

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)

	fmt.Println("Press enter to exit")
	// python input() to prevent the program from exiting immediately
	fmt.Scanln()
}
