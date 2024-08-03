package main

import (
	"fmt"
	"os"
	"time"

	"github.com/117503445/gorobot-demo/internal/cv"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("2024.8.3 11:57")

	robotgo.MouseSleep = 100
	robotgo.KeySleep = 100

	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		os.Exit(1)
		hook.End()
		// exit(0)
	})

	// c := 0

	fmt.Println("--- Please press alt+w---")
	hook.Register(hook.KeyDown, []string{"alt", "w"}, func(e hook.Event) {

		img := robotgo.CaptureImg()
		points := cv.GetUnlockedPoints(img)
		for _, p := range points {
			robotgo.Move(p.X/2, p.Y/2)
			time.Sleep(3 * time.Second)
			robotgo.Click()
			time.Sleep(3 * time.Second)
			robotgo.Move(980, 540)
			time.Sleep(3 * time.Second)
			robotgo.Click()
			time.Sleep(3 * time.Second)
			robotgo.Move(3726/2, 394/2)
			time.Sleep(3 * time.Second)
		}
		// 3726, 394

		// // 取消选择所有仪器
		// go func() {
		// 	robotgo.Move(3624/2, 546/2)
		// 	for {
		// 		// time.Sleep(3 * time.Second)
		// 		robotgo.Click()
		// 		time.Sleep(3 * time.Second)
		// 		// robotgo.KeyTap("d")
		// 		// time.Sleep(3 * time.Second)
		// 	}
		// }()

		// c += 1
		// name := fmt.Sprintf("%d.png", c)

		// var err error

		// // x, y, w, h
		// img := robotgo.CaptureImg()
		// err = robotgo.Save(img, name)
		// if err = robotgo.Save(img, name); err != nil {
		// 	panic(err)
		// }

	})

	s := hook.Start()
	<-hook.Process(s)

	// fmt.Println("Press enter to exit")
	// // python input() to prevent the program from exiting immediately
	// fmt.Scanln()
}
