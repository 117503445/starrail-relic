package main

import (
	"fmt"
	// "image"
	"os"
	"time"

	"github.com/117503445/gorobot-demo/internal/cv"
	"github.com/117503445/goutils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/rs/zerolog/log"
)

var logsDir string

func AltWCallback() {
	img := robotgo.CaptureImg()

	imgFile := fmt.Sprintf("%s/%s.png", logsDir, time.Now().Format("20060102.150405"))
	if err := robotgo.Save(img, imgFile); err != nil {
		log.Fatal().Err(err).Msg("robotgo.Save")
	}

	// 轮换遗器槽
	points := cv.GetRelicPoints()
	for _, p := range points {
		robotgo.Move(p.X/2, p.Y/2)
		robotgo.Click()
	}

	// 选中当前未选中的遗器
	// points := cv.GetUnlockedPoints(img)
	// log.Debug().Interface("points", points).Msg("GetUnlockedPoints")
	// for _, p := range points {
	// 	robotgo.Move(p.X/2, p.Y/2)
	// 	robotgo.Click()
	// 	robotgo.Move(3726/2, 394/2)
	// 	robotgo.Click()
	// }

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
}

func main() {
	goutils.InitZeroLog()
	// runID as a unique identifier for the current run, filename
	// 20240803.203942
	runID := time.Now().Format("20060102.150405")

	logsDir = fmt.Sprintf("logs/%s", runID)
	if err := os.MkdirAll(logsDir, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("os.MkdirAll")
	}

	robotgo.MouseSleep = 1000
	robotgo.KeySleep = 1000

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
		go AltWCallback()
	})

	s := hook.Start()
	<-hook.Process(s)

	// fmt.Println("Press enter to exit")
	// // python input() to prevent the program from exiting immediately
	// fmt.Scanln()
}
